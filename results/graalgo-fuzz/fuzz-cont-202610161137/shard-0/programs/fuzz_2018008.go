package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{(^(v1)), v1, v1}
    var v3 int = 4294967295
    var v4 int64 = -1
    var v5 uint16 = 858
    fmt.Println(v4)
    if false {
        for fi6 := 0; fi6 < 13; fi6++ {
            v1 = v1
            if (true || true) {
                var v7 float64 = ((((3.14159 - -2.25) + 0.1) * 0.0) / 3.14159)
                _ = v7
            }
            _ = fi6
        }
        v5 = v5
        v4 -= (v4 / ((v4) | 1))
    }
    v8 := GenS1{F0: -1, F1: ("" + "")}
    v4 = v4
    v3 -= v3
    switch (v3 / (((^(v3))) | 1)) {
    case 4294967295:
        var v9 uint8 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 float64 = ((3.14159 + -2.25) / 1.0)
        var v12 string = "hello"
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = len(v2)
    switch v3 {
    case 429:
        if ((1000 != 32767) || false) {
            v13 := GenS1{F0: 881, F1: "hello"}
            switch v3 {
            case -9223372036854775808:
                var v14 bool = (276 <= (v3 & v3))
                _ = v14
            case -208:
                v15 := GenS1{F0: int16(v5), F1: ""}
                _ = v15
            default:
                if (3.14159 >= (0.1 / 3.14159)) {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        if (v1 < (^(v1))) {
                            fmt.Println(v1, v4)
                        } else {
                            v3 += 4294967295
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                        }
                        _ = fi16
                    }
                    fmt.Println(v5)
                } else {
                    var v17 uint = 2
                    var v18 bool = (!(v5 == (^(v5))))
                    var v19 uint8 = uint8(v17)
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
            var v20 int8 = 127
            _ = v13
            _ = v20
        } else {
            var v21 uint16 = (v5 % ((v5) | 1))
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 14; fi22++ {
            switch (v3 / ((698) | 1)) {
            case 32768:
                v3 = (v3 >> ((18446744073709551615) & 63))
            case 32767:
                v23 := GenS1{F0: int16(fi22), F1: ("a" + "a")}
                _ = v23
            case -321:
                if (false && (42 > 2147483647)) {
                    var v24 uint32 = v1
                    _ = v24
                }
            }
            v4 -= v4
            if true {
                v3 = -1
                if (65535 == v5) {
                    var v25 uint32 = (v1 * 42)
                    _ = v25
                } else {
                    var v26 float64 = float64(1.5)
                    fmt.Println(v4, v26)
                    _ = v26
                }
                v3 = 1
            } else {
                v27 := GenS1{F0: 748, F1: "\n\t"}
                _ = v27
            }
            _ = fi22
        }
    }
    if (!(false || false)) {
        switch v3 {
        case 0:
            v28 := GenS1{F0: int16(v5), F1: ("\n\t" + "a")}
            _ = v28
        case 2147483647:
            v4 = v4
        }
    }
    v3 = v3
    var v29 int = (v3 / ((23) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
