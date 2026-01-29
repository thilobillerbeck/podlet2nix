{
  description = "Nix flake for podlet2nix";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
      ];
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "aarch64-darwin"
        "x86_64-darwin"
      ];
      perSystem =
        {
          config,
          self',
          inputs',
          pkgs,
          system,
          ...
        }:
        let
          nativeBuildInputs = with pkgs; [
            go
            gopls
            cobra-cli
            podlet
            (writeShellScriptBin "podlet2nix" ''
              go run main.go "$@"
            '')
          ];
        in
        {
          devShells.default = pkgs.mkShell { inherit nativeBuildInputs; };

          packages.default = pkgs.buildGoModule {
            name = "podlet2nix";
            src = ./.;

            cgo_enabled = false;

            ldflags = [
              "-s"
              "-w"
            ];

            vendorHash = "sha256-wBdQIjRlA6Xr5o/ejmMO+/NWCGNKTApexb5z+6L7wbE=";
          };
        };
      flake = {
      };
    };
}
