{
  description = "A Flake for Pollux's accompanying LaTeX report";

  inputs = {
    nixpkgs.url = "nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
        };
      in {
        devShells.default = with pkgs;
          mkShell {
            buildInputs = [
              (texlive.combine {
                inherit
                  (texlive)
                  scheme-medium
                  geometry
                  newtx
                  fontaxes
                  xstring
                  bigfoot
                  microtype
                  titlesec
                  hyperref
                  amsmath
                  mathtools
                  enumitem
                  mathpartir
                  xcolor
                  ;
              })
            ];

            shellHook = ''
            '';
          };
      }
    );
}
