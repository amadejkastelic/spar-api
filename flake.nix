{
  description = "Spar REST API Client";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

    flake-utils.url = "github:numtide/flake-utils";

    pre-commit-hooks = {
      url = "github:cachix/git-hooks.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      nixpkgs,
      flake-utils,
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

        devShells.default = import ./nix/dev-shell.nix {
          inherit pkgs;
          deps = devDeps;
          preCommitCheck = preCommitCheck;
        };
      }
    );
}
