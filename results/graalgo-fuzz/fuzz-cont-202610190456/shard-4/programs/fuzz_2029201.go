package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint32; F2 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 uint }
type GenS4 struct { F0 int64; F1 int32; F2 uint16 }

func main() {
    var v1 int32 = 42
    var v2 int16 = -32768
    var v3 string = ("a" + "foo bar")
    var v4 uint64 = 4294967295
    v5 := []int{1, 966, 1}
    var v6 uint8 = 2
    switch (v3 + "foo bar") {
    case "foo bar":
        v6 -= v6
    case "a":
        for fi7 := 0; fi7 < 16; fi7++ {
            if ((v4 ^ 1) == v4) {
                for fi8 := 0; fi8 < 10; fi8++ {
                    v9 := GenS1{F0: 101}
                    switch fi7 {
                    case -2147483648:
                        switch "a" {
                        case "foo bar":
                            if (((0.0 + (((((3.14159 + 3.14159) / -2.25) / 1.5) + float32(3.14159)) + (float32(float64(0.1)) - 0.1))) + 1.5) == 0.0) {
                                fmt.Println(fi7)
                                v10 := GenS1{F0: 449}
                                v6 = v6
                                _ = v10
                            }
                        case "\n\t":
                            v6 -= (v6 >> ((42) & 63))
                        }
                    case 4294967295:
                        var v11 uint32 = 0
                        _ = v11
                    case 398:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = fi8
                    _ = v9
                }
                if false {
                    var v12 int16 = -405
                    _ = v12
                } else {
                    v13 := GenS2{F0: (^(v2)), F1: 2147483648, F2: -70}
                    _ = v13
                }
                var v14 uint8 = 127
                _ = v14
            }
            v15 := GenS1{F0: 554}
            _ = fi7
            _ = v15
        }
    case "hello":
        var v16 string = "a"
        _ = v16
    }
    if false {
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint32 = 0
            v3 = ("hello" + ((v3 + "a") + v3))
            if ((18446744073709551615 != 706) || false) {
                switch "\n\t" {
                case "hello":
                    v5 = append(v5, 2147483647)
                case "foo bar":
                    var v19 uint32 = (v18 ^ 0)
                    _ = v19
                case "a":
                    v18 = (v18 - (v18 >> ((255) & 63)))
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch 9223372036854775807 {
                    case 0:
                        var v21 int = fi17
                        _ = v21
                    case 938:
                        switch fi20 {
                        case 2147483647:
                            v22 := GenS3{F0: (-(v2)), F1: 42, F2: 18446744073709551615}
                            _ = v22
                        default:
                            var v23 int8 = 1
                            _ = v23
                        }
                    }
                    _ = fi20
                }
            }
            _ = fi17
            _ = v18
        }
        v4 = v4
    } else {
        var v24 int32 = v1
        _ = v24
    }
    if true {
        var v25 string = "a"
        _ = v25
    }
    v26 := GenS1{F0: 0}
    _ = len(v5)
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 float64 = ((1e10 * (-2.25 / 1.0)) * 3.14159)
        v5 = append(v5, (fi27 * fi27))
        v29 := GenS3{F0: (v2 & v2), F1: 42, F2: 65535}
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint8 = (v6 | v6)
    if len(v5) > 0 {
        v5[0] = 32768
    }
    _ = len(v5)
    v31 := GenS4{F0: -9223372036854775808, F1: (v1 + v1), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
