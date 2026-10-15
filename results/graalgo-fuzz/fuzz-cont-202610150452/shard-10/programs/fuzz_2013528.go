package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []string{"", ("a" + "a"), (("" + "hello") + "a")}
    var v2 bool = true
    v3 := []string{"a"}
    v4 := make(map[uint32]int)
    delete(v4, 2)
    if (65535 == 255) {
        var v5 uint8 = 127
        if (true && v2) {
            if ((((("\n\t" + (("a" + "") + "foo bar")) + "") + ("hello" + "foo bar")) + "") == "foo bar") {
                fmt.Println(v2, v5)
                if (!false) {
                    v4[684] = -473
                } else {
                    v5 += (v5 / (((v5 * v5)) | 1))
                    v6 := v4[2147483648]
                    for fi7 := 0; fi7 < 13; fi7++ {
                        var v8 uint16 = 2
                        v8 += v8
                        _ = fi7
                        _ = v8
                    }
                    _ = v6
                }
            }
            var v9 int32 = 42
            _ = v9
        } else {
            var v10 float32 = 3.14159
            _ = v10
        }
        v5 = v5
        _ = v5
    }
    var v11 string = ("" + ("a" + "a"))
    for fi12 := 0; fi12 < 12; fi12++ {
        delete(v4, 4294967295)
        for fi13 := 0; fi13 < 14; fi13++ {
            if (!v2) {
                v11 = ((((("foo bar" + ((v11 + "\n\t") + "\n\t")) + (((v11 + v11) + "hello") + v11)) + "foo bar") + "") + v11)
            } else {
                v11 = (("\n\t" + v11) + "a")
                if (100 == 96) {
                    if len(v3) > 0 {
                        v3[0] = v11
                    }
                }
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi12, fi12, v11)
                case "":
                    v14 := GenS1{F0: 243, F1: 32767, F2: 492}
                    _ = v14
                }
            }
            var v15 uint = 65535
            var v16 bool = ((((((("foo bar" + v11) + (v11 + v11)) + v11) + v11) + v11) + v11) <= ((v11 + v11) + v11))
            _ = fi13
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    if false {
        if true {
            v17 := GenS1{F0: 0, F1: 9223372036854775807, F2: 4294967295}
            for fi18 := 0; fi18 < 4; fi18++ {
                if ((!true) && (false || (("hello" + v11) > "\n\t"))) {
                    v19 := GenS1{F0: 255, F1: (fi18 << ((255) & 63)), F2: 4294967295}
                    var v20 int = fi18
                    v20 = (-(v20))
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
            _ = v17
        }
    } else {
        v11 = v11
    }
    v4[2] = -1
    fmt.Println(v2, v2)
    switch 137 {
    case 793:
        v11 = v11
    case 2147483647:
        v21 := GenS2{F0: (!v2), F1: -9223372036854775808, F2: true}
        _ = v21
    case 4294967295:
        for fi22 := 0; fi22 < 13; fi22++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi22
        }
    default:
        for fi23 := 0; fi23 < 6; fi23++ {
            if (!true) {
                v24 := v4[4294967295]
                v25 := GenS3{F0: 65535, F1: (!v2)}
                v11 = (v11 + (v11 + v11))
                _ = v24
                _ = v25
            } else {
                var v26 uint8 = 0
                v27 := GenS3{F0: 127, F1: (uint16(v26) == 178)}
                _ = v26
                _ = v27
            }
            v28 := GenS2{F0: v2, F1: 0, F2: false}
            v11 = "hello"
            _ = fi23
            _ = v28
        }
    }
    v29 := GenS3{F0: 32767, F1: v2}
    var v30 int8 = 100
    for fi31 := 0; fi31 < 10; fi31++ {
        v11 = v11
        _ = fi31
    }
    if false {
        var v32 uint64 = 42
        var v33 uint16 = 290
        v32 += (^(v32))
        _ = v32
        _ = v33
    } else {
        if (1 <= 32767) {
            v34 := GenS4{F0: 255, F1: 42}
            var v35 uint = 255
            _ = v34
            _ = v35
        }
        var v36 float32 = (float32(1e10) / 0.1)
        var v37 uint16 = 65535
        _ = v36
        _ = v37
    }
    v2 = (255 > 152)
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 bool = true
        v40 := GenS1{F0: 255, F1: (fi38 * fi38), F2: 14}
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS5{F0: 255, F1: 128}
    var v42 int = 0
    for fi43 := 0; fi43 < 1; fi43++ {
        v44 := GenS1{F0: 241, F1: (^(fi43)), F2: 4294967295}
        var v45 float64 = 0.1
        var v46 uint16 = 0
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    fmt.Println(v11, v42, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v42
}
