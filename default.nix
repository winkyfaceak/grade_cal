{ lib, buildGoModule }:
buildGoModule {
  pname = "grade_cal";
  version = "0.0.1";

  src = ./.;

  vendorHash = null;

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "grade calulator";
    homepage = "https://github.com/winkyfaceak/grade_cal";
    license = lib.licenses.mit;
    maintainers = [ ];
    mainPackage = "grade_cal";
  };
}
