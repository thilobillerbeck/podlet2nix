{
  description = "Nix flake for podlet2nix";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
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
        pkgs = nixpkgs.legacyPackages.${system};

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
      }
    );
}
