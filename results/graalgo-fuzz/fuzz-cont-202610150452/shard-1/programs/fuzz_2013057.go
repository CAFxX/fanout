package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint; F1 int; F2 uint }
type GenS5 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 uint, p1 uint, p2 uint) uint {
    v4 := []int{2147483647, -1, -87, 32767}
    p1 += 1
    switch ((("" + ("\n\t" + "\n\t")) + "a") + ("hello" + ("\n\t" + "hello"))) {
    case "a":
        if len(v4) > 0 {
            v4[0] = -712
        }
    case "\n\t":
        for fi5 := 0; fi5 < 13; fi5++ {
            if false {
                for fi6 := 0; fi6 < 4; fi6++ {
                    p1 = (p1 << ((p1) & 63))
                    _ = fi6
                }
            }
            _ = fi5
        }
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            p2 -= (p1 - (^(p0)))
            fmt.Println(p1, fi7, fi7)
            if false {
                var v8 int16 = 293
                _ = v8
            }
            _ = fi7
        }
    }
    v9 := GenS1{F0: uint16(p2), F1: -10}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return (p2 ^ p2)
}
func genFn1(p0 int32) uint32 {
    p0 = (p0 & -1)
    p0 += (p0 / (((p0 * 0)) | 1))
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: 65535, F1: 116}
        var v12 int32 = p0
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    return 42
}

func main() {
    var v1 bool = (3.14159 < float32(float64(0.0)))
    var v2 uint8 = 127
    var v3 bool = (!(32767 >= 32767))
    v3 = (0 == 0)
    var v13 int = 127
    if (42 > 18446744073709551615) {
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS2{F0: 127}
            var v16 int16 = 1000
            var v17 uint = 2
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    } else {
        v1 = (1 == -1)
        if ((!(655 < 485)) || v1) {
            for fi18 := 0; fi18 < 9; fi18++ {
                switch "\n\t" {
                case "foo bar":
                    v19 := GenS3{F0: v2, F1: 246, F2: 2}
                    _ = v19
                case "\n\t":
                    if (18446744073709551615 == 125) {
                        var v20 int = (^(fi18))
                        _ = v20
                    }
                case "a":
                    var v21 int8 = -1
                    _ = v21
                default:
                    var v22 string = "foo bar"
                    _ = v22
                }
                v3 = v3
                var v23 uint64 = 1
                _ = fi18
                _ = v23
            }
        } else {
            v2 = 113
            fmt.Println(v1)
            v24 := GenS4{F0: 255, F1: (v13 >> ((510) & 63)), F2: 1}
            _ = v24
        }
    }
    var v25 float64 = 0.1
    v3 = ((true && (!v1)) && v1)
    switch ("\n\t" + ("a" + "a")) {
    case "hello":
        var v26 int64 = 972
        _ = v26
    case "":
        var v27 uint = genFn0(65535, 65535, 255)
        _ = v27
    case "foo bar":
        v13 = (v13 | v13)
    default:
        if ((("hello" == "a") || false) && (!v1)) {
            v28 := []int{(v13 / ((v13) | 1)), (-(v13)), 9223372036854775807}
            _ = v28
        }
    }
    v29 := GenS5{F0: 1000, F1: 1, F2: 255}
    v1 = true
    for fi30 := 0; fi30 < 12; fi30++ {
        v13 = v13
        var v31 uint16 = 1
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        if (v3 && (1 <= 2147483648)) {
            var v33 string = ("a" + ("foo bar" + ("a" + ("hello" + "foo bar"))))
            _ = v33
        } else {
            fmt.Println(v1)
        }
        _ = fi32
    }
    v13 += (v13 + -804)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v29
    _ = v3
}
