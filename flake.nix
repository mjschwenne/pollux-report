{
  description = "A Flake for Pollux's accompanying LaTeX report";

  inputs = {
    nixpkgs.url = "nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    cddl.url = "github:anweiss/cddl";
    rust-overlay.url = "github:oxalica/rust-overlay";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    cddl,
    rust-overlay,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        overlays = [(import rust-overlay)];
        pkgs = import nixpkgs {
          inherit system overlays;
        };
      in {
        devShells.default = with pkgs;
          mkShell {
            buildInputs = [
              # For protobuf experiments
              protobuf
              protoc-gen-go
              protoscope
              go
              gopls
              xxd
              cddl.packages.${system}.default
              rust-bin.stable.latest.default

              # LaTeX
              (texlive.combine {
                inherit
                  (texlive)
                  scheme-medium
                  biblatex
                  biber
                  geometry
                  wrapfig
                  multirow
                  booktabs
                  capt-of
                  titlesec
                  hyperref
                  amsmath
                  mathtools
                  enumitem
                  mathpartir
                  xcolor
                  tcolorbox
                  bbold
                  ;
              })
            ];

            shellHook = ''
            '';
          };
      }
    );
}
