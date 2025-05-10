{
  description = "Spar REST API Client";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

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
      pre-commit-hooks,
      nix-github-actions,
      ...
    }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "aarch64-darwin"
      ];

      forAllSystems =
        f:
        builtins.listToAttrs (
          map (system: {
            name = system;
            value = f system;
          }) supportedSystems
        );

      perSystem =
        system:
        let
          pkgs = import nixpkgs {
            inherit system;
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
            preCommitHooks = pre-commit-hooks.lib.${system};
          };

          package = pkgs.buildGoModule {
            pname = "spar-api-cli";
            version = "0.1.0";
            src = ./.;
            vendorHash = null;
            doCheck = true;
            subPackages = [ "cmd/example" ];
          };
        in
        {
          packages.default = package;
          devShells.default = import ./nix/dev-shell.nix {
            inherit pkgs;
            deps = devDeps;
            preCommitCheck = preCommitCheck;
          };
          checks.pre-commit-check = preCommitCheck;
        };

    in
    {
      packages = forAllSystems (system: (perSystem system).packages);
      devShells = forAllSystems (system: (perSystem system).devShells);
      checks = forAllSystems (system: (perSystem system).checks);

      githubActions = nix-github-actions.lib.mkGithubMatrix {
        checks = {
          x86_64-linux = self.checks.x86_64-linux;
          aarch64-linux = self.checks.aarch64-linux;
          aarch64-darwin = self.checks.aarch64-darwin;
        };
      };
    };
}
