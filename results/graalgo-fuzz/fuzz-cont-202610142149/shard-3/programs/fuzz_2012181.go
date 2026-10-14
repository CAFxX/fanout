package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int32; F2 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := []uint32{4294967295, 2147483648, 4294967295}
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 691
    var v6 uint8 = 96
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 bool = (("" >= ("" + ("a" + ""))) && true)
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := v1["\n\t"]
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v6 -= uint8(v5)
        if true {
            var v12 int = fi11
            if false {
                v13 := GenS1{F0: (v5 | (v5 << ((111) & 63))), F1: 2, F2: 0}
                _ = v13
            } else {
                var v14 uint16 = 65535
                _ = len(v2)
                _ = v14
            }
            v15 := GenS2{F0: 0}
            _ = v12
            _ = v15
        } else {
            fmt.Println(fi11)
            v16 := GenS2{F0: (v6 ^ (v6 << ((551) & 63)))}
            v6 = (v6 % ((v6) | 1))
            _ = v16
        }
        _ = fi11
    }
    fmt.Println(v5, v6)
    var v17 uint = 225
    if true {
        v2[("\n\t" + "foo bar")] = 2147483648
        v18 := GenS1{F0: 1, F1: (v17 << ((4294967295) & 63)), F2: 2147483647}
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 uint32 = v5
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        fmt.Println(v5, v5)
        if (!(3.14159 == 0.1)) {
            var v21 int8 = -20
            _ = v21
        } else {
            switch ("\n\t" + ("a" + "hello")) {
            case "foo bar":
                for fi22 := 0; fi22 < 7; fi22++ {
                    v17 -= v17
                    var v23 string = ""
                    _ = fi22
                    _ = v23
                }
            case "\n\t":
                var v24 int = 4294967295
                _ = v24
            case "hello":
                v25 := GenS1{F0: 690, F1: 18446744073709551615, F2: -2147483648}
                _ = v25
            default:
                for fi26 := 0; fi26 < 1; fi26++ {
                    v5 = v5
                    delete(v4, (^(v5)))
                    v6 += (^(v6))
                    _ = fi26
                }
            }
            if (0 <= 9223372036854775807) {
                switch -1 {
                case -2147483648:
                    switch -2147483648 {
                    case -2147483648:
                        for fi27 := 0; fi27 < 12; fi27++ {
                            switch (("\n\t" + (((("a" + "hello") + (("" + ("a" + "foo bar")) + ("a" + ("\n\t" + "hello")))) + "a") + ("" + ("hello" + "foo bar")))) + (("a" + ("foo bar" + (("foo bar" + "hello") + "hello"))) + "")) {
                            case "\n\t":
                                fmt.Println(v5, v6, v5)
                            default:
                                v28 := GenS1{F0: 42, F1: (v17 & (^(v17))), F2: 884}
                                _ = v28
                            }
                            _ = fi27
                        }
                    case -9223372036854775808:
                        if true {
                            v29 := GenS2{F0: v6}
                            v30 := GenS2{F0: uint8(v5)}
                            for fi31 := 0; fi31 < 19; fi31++ {
                                var v32 int8 = 68
                                _ = fi31
                                _ = v32
                            }
                            _ = v29
                            _ = v30
                        }
                    case 625:
                        var v33 float32 = 0.0
                        _ = v33
                    }
                case 2147483647:
                    if (-2.25 != 0.0) {
                        var v34 float64 = 3.14159
                        _ = v34
                    }
                }
            }
        }
        fmt.Println(v6, v5, v6)
    }
    v35 := GenS3{F0: ((127 < -9223372036854775808) || false), F1: -707, F2: 632}
    var v36 int16 = -32768
    for fi37 := 0; fi37 < 8; fi37++ {
        if ((v36 - v36) != (v36 & -32768)) {
            for fi38 := 0; fi38 < 5; fi38++ {
                switch "" {
                case "hello":
                    switch fi38 {
                    case -945:
                        var v39 int = fi37
                        _ = v39
                    default:
                        var v40 uint16 = 68
                        _ = v40
                    }
                case "foo bar":
                    if (float32(0.0) == (0.0 - ((3.14159 * -2.25) - (3.14159 / 1.5)))) {
                        var v41 int = (fi37 * -1)
                        v42 := v1[""]
                        var v43 float64 = 1.5
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                case "a":
                    delete(v2, ("hello" + "a"))
                }
                fmt.Println(fi38, v6)
                _ = fi38
            }
        } else {
            var v44 string = "a"
            _ = v44
        }
        v45 := GenS2{F0: (v6 % ((v6) | 1))}
        _ = fi37
        _ = v45
    }
    fmt.Println(v17, v36)
    v17 = v17
    var v46 float32 = ((float32(-2.25) * 0.1) * 1e10)
    fmt.Println(v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
