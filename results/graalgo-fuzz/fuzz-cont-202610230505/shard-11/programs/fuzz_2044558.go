package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{"hello", "\n\t", "", "a"}
    var v3 int8 = -128
    var v4 float64 = 0.0
    v5 := []string{"hello", "\n\t", ("" + "foo bar"), "a"}
    var v6 float32 = ((0.0 + 1e10) / 1.5)
    fmt.Println(v1, v3)
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7, v6)
        var v8 int8 = (v3 + 1)
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint32 = 2147483647
            if true {
                var v11 float64 = 1e10
                _ = v11
            } else {
                switch ("hello" + "") {
                case "hello":
                    v8 -= -128
                }
            }
            v12 := GenS1{F0: 127}
            v6 -= -2.25
            _ = fi9
            _ = v10
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v3)
    v4 = (v4 + (1e10 + (v4 + v4)))
    var v13 int32 = -57
    if (!false) {
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS2{F0: 255}
            _ = fi14
            _ = v15
        }
        if len(v2) > 0 {
            v2[0] = "a"
        }
    }
    v16 := GenS1{F0: 0}
    if (v4 <= (0.0 - v4)) {
        if (false && (2 > 676)) {
            var v17 uint64 = uint64(v1)
            if (522 > 9223372036854775807) {
                var v18 int8 = (-(v3))
                if (true || ("a" <= "foo bar")) {
                    v17 = (v17 + v17)
                    if false {
                        fmt.Println(v1)
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                    }
                }
                v19 := GenS3{F0: 32767, F1: 0}
                _ = v18
                _ = v19
            } else {
                fmt.Println(v17, v13)
            }
            fmt.Println(v1)
            _ = v17
        } else {
            _ = len(v5)
        }
    }
    v4 = float64(0.1)
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS4{F0: 4294967295, F1: 4294967295}
            v6 = (float32(v4) / ((((v6 / 1.0) + v6) / v6) / (v6 + v6)))
            switch ("a" + "\n\t") {
            case "":
                fmt.Println(v6, v3)
            default:
                if len(v5) > 0 {
                    v5[0] = ""
                }
            }
            v13 -= v13
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v4 = 1.5
    fmt.Println(v1, v13, v4)
    v6 = v6
    switch 32768 {
    case -2147483648:
        v13 = (v13 << ((495) & 63))
    case 32768:
        for fi23 := 0; fi23 < 20; fi23++ {
            if len(v5) > 0 {
                v5[0] = "a"
            }
            _ = fi23
        }
    case 0:
        v24 := GenS2{F0: 18446744073709551615}
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 float64 = v4
            if false {
                var v28 int16 = 42
                v13 += v13
                v3 -= v3
                _ = v28
            } else {
                if false {
                    var v29 int64 = 42
                    _ = v29
                } else {
                    var v30 int64 = -267
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            var v31 uint = 255
                            _ = v31
                        }
                        if false {
                            v1 = (v1 * 42)
                            var v32 uint8 = 255
                            _ = v32
                        }
                    }
                    v30 = 42
                    _ = v30
                }
            }
            v5 = append(v5, "hello")
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
