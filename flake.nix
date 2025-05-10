{
  description = "Spar REST API Client";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

    flake-utils.url = "github:numtide/flake-utils";

    pre-commit-hooks = {
      url = "github:cachix/git-hooks.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    nix-github-actions = {
      url = "github:nix-community/nix-github-actions";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      nix-github-actions,
      ...
    }@inputs:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ ];
        };

        buildDeps = [
          pkgs.git
          pkgs.go
          pkgs.gnumake
        ];
        devDeps = buildDeps ++ [
          pkgs.golangci-lint
          pkgs.delve
          pkgs.gopls
          pkgs.golines
        ];

        preCommitCheck = import ./nix/pre-commit.nix {
          inherit pkgs;
          preCommitHooks = inputs.pre-commit-hooks.lib.${system};
        };
      in
      {
        checks = {
          pre-commit-check = preCommitCheck;
        };

        githubActions = nix-github-actions.lib.mkGithubMatrix {
          inherit (self) checks;
        };

        packages.default = pkgs.buildGoModule {
          pname = "spar-api-cli";
          version = "0.1.0";

          src = ./.;

          vendorHash = null;

          doCheck = true;

          subPackages = [ "cmd/example" ];
        };

        devShells.default = import ./nix/dev-shell.nix {
          inherit pkgs;
          deps = devDeps;
          preCommitCheck = preCommitCheck;
        };
      }
    );
}
