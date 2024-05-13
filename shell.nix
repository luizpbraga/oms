{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = with pkgs.buildPackages; [ 
      go
      air
      mongodb
      protobuf
      grpcurl
      docker
    ];
}
