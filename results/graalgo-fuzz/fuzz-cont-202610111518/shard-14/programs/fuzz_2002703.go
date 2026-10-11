package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint16 }
func genFn0(p0 uint32, p1 uint, p2 uint) bool {
    var v4 int = 4294967295
    var v5 int64 = 482
    p0 = 745
    var v6 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (!((0.1 < 1.5) || true))
}
func genFn1() int64 {
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 2; fi8++ {
        v7 -= v7
        _ = fi8
    }
    if (-125 != -128) {
        if false {
            v7 = v7
        }
    }
    v9 := []int{128, 127}
    for fi10 := 0; fi10 < 18; fi10++ {
        if true {
            v7 -= v7
            v9 = append(v9, 9223372036854775807)
            for fi11 := 0; fi11 < 9; fi11++ {
                switch ("" + ((("" + "a") + ("" + ((("\n\t" + "foo bar") + "\n\t") + ("a" + ("foo bar" + (("foo bar" + ("\n\t" + "foo bar")) + "\n\t")))))) + ("a" + ("foo bar" + ("foo bar" + ""))))) {
                case "":
                    switch (fi11 >> ((42) & 63)) {
                    case 217:
                        v7 = v7
                    }
                case "\n\t":
                    v12 := GenS1{F0: true, F1: 0, F2: v7}
                    _ = v12
                default:
                    var v13 int = -528
                    _ = v13
                }
                _ = fi11
            }
        } else {
            v14 := GenS2{F0: 2147483647, F1: 100}
            _ = v14
        }
        if (0 != 78) {
            v15 := GenS2{F0: 1, F1: -128}
            _ = v15
        }
        if (!((0.1 - float32(3.14159)) >= 0.0)) {
            v16 := GenS3{F0: 3, F1: -2, F2: 2}
            _ = v16
        }
        _ = fi10
    }
    _ = v7
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{2, 42, 1, 65535}
    var v2 int64 = 42
    var v3 uint16 = 819
    v17 := GenS2{F0: 2, F1: 0}
    var v18 float64 = 3.14159
    fmt.Println(v18, v3, v18)
    v1 = append(v1, 713)
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v3 = 80
    if (!(!false)) {
        v3 = v3
    } else {
        if false {
            var v19 bool = false
            v2 = (^(v2))
            _ = v19
        }
        v3 = v3
        v18 += (v18 + v18)
    }
    if ((18446744073709551615 != 65535) || (v2 >= (v2 ^ v2))) {
        var v20 float64 = 0.0
        v3 = (v3 % (((^(v3))) | 1))
        if (0 < 42) {
            _ = len(v1)
        }
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
