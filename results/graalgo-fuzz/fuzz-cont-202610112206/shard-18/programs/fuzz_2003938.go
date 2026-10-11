package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint }

func main() {
    var v1 uint = 42
    var v2 int32 = -940
    var v3 float64 = float64(1.5)
    v4 := []string{}
    v2 += (v2 / ((v2) | 1))
    var v5 uint = 18446744073709551615
    var v6 int = 9223372036854775807
    var v7 float32 = 1e10
    if ((127 < -12) || true) {
        if false {
            var v8 uint16 = 1
            var v9 int16 = 42
            switch "\n\t" {
            case "":
                switch "foo bar" {
                case "foo bar":
                    v10 := GenS1{F0: 65535, F1: 110, F2: 1}
                    _ = v10
                case "\n\t":
                    var v11 int = v6
                    _ = v11
                default:
                    v12 := GenS1{F0: 65535, F1: 2147483647, F2: 69}
                    _ = v12
                }
            default:
                var v13 bool = false
                _ = v13
            }
            _ = v8
            _ = v9
        }
        if false {
            v14 := GenS1{F0: 2147483647, F1: 1, F2: -1}
            var v15 int64 = 42
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS2{F0: -1, F1: -9223372036854775808, F2: v5}
                _ = fi16
                _ = v17
            }
            v5 = v1
            var v18 uint32 = 2147483648
            _ = v18
        }
    }
    if len(v4) > 0 {
        v4[0] = ("\n\t" + (("" + ("\n\t" + ("foo bar" + "foo bar"))) + ("hello" + ("hello" + "\n\t"))))
    }
    fmt.Println(v2, v1)
    v19 := GenS1{F0: 718, F1: 65535, F2: -572}
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(v5, v7, v2)
        _ = fi20
    }
    if (42 != (v6 + (v6 * (v6 ^ (v6 / ((v6) | 1)))))) {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 uint16 = 981
            v23 := GenS2{F0: -1, F1: -9223372036854775808, F2: (^(v5))}
            _ = fi21
            _ = v22
            _ = v23
        }
        fmt.Println(v6)
    } else {
        v5 -= v1
        v24 := GenS2{F0: -9223372036854775808, F1: int64(v2), F2: 0}
        for fi25 := 0; fi25 < 6; fi25++ {
            if true {
                if false {
                    var v26 bool = false
                    for fi27 := 0; fi27 < 20; fi27++ {
                        fmt.Println(v5)
                        _ = len(v4)
                        _ = fi27
                    }
                    _ = v26
                } else {
                    var v28 int8 = 74
                    _ = v28
                }
                v1 = v1
                v29 := GenS2{F0: -1, F1: 42, F2: 4294967295}
                _ = v29
            }
            _ = len(v4)
            _ = fi25
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
