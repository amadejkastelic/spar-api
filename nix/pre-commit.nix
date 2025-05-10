{ pkgs, preCommitHooks }:

preCommitHooks.run {
  src = ./.;
  hooks = {
    nixfmt-rfc-style.enable = true;
    gofmt.enable = true;
    golangci-lint = {
      enable = true;
      extraPackages = [ pkgs.go ];
    };
    golines = {
      enable = true;
      settings.flags = "-m 100 --dry-run";
    };
    gotest.enable = true;
    govet.enable = true;
  };
}
