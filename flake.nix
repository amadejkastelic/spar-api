{
  description = "Spar REST API Client";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      nixpkgs,
      flake-utils,
      ...
    }:
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
        ];
      in
      {
        devShell = pkgs.mkShell {
          nativeBuildInputs = devDeps;

          hardeningDisable = [ "fortify" ];
        };
      }
    );
}
