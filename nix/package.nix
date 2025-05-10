{ pkgs, version, ... }:
pkgs.buildGoModule {
  pname = "spar-api-cli";
  version = version;
  src = ./..;
  vendorHash = null;
  doCheck = true;
  subPackages = [ "cmd/example" ];
}
