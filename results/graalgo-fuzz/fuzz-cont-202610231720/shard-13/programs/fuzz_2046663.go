package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0() uint32 {
    var v7 float64 = -2.25
    var v8 uint = 18446744073709551615
    if (!false) {
        var v9 uint = (^(v8))
        _ = v9
    } else {
        var v10 int = 4294967295
        v10 = 1
        _ = v10
    }
    fmt.Println(v8, v7)
    v11 := GenS1{F0: 128, F1: 0, F2: (743 != 0)}
    _ = v11
    _ = v7
    _ = v8
    return 951
}
func genFn1(p0 bool, p1 uint8) uint16 {
    switch "foo bar" {
    case "\n\t":
        v12 := []int{-2147483648, 4294967295}
        _ = v12
    }
    fmt.Println(p0, p1)
    fmt.Println(p1, p0)
    fmt.Println(p0)
    if p0 {
        p0 = (255 >= 42)
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 float64 = float64(-2.25)
    var v2 bool = true
    var v3 uint32 = 0
    var v4 uint8 = 1
    v5 := make(map[int]bool)
    var v6 string = "foo bar"
    var v13 uint = 4294967295
    v3 = genFn0()
    v1 = 3.14159
    v6 = (v6 + (v6 + (v6 + v6)))
    v2 = (!((-1 <= 32767) || (v2 && (0 != 1000))))
    fmt.Println(v1, v4)
    fmt.Println(v1, v4, v3)
    v13 += (^(v13))
    v14 := v5[506]
    if ((0 > -1) || v2) {
        if (4294967295 != 4294967295) {
            if v2 {
                var v15 uint16 = 1
                if (42 >= 0) {
                    v5[42] = (-508 == -310)
                    v16 := GenS1{F0: (v4 << (((v13 - 4294967295)) & 63)), F1: 0, F2: v14}
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 uint64 = 1
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
                _ = v15
            }
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 uint64 = 0
            for fi21 := 0; fi21 < 12; fi21++ {
                v20 += uint64(v3)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
    }
    _ = len(v5)
    switch v6 {
    case "a":
        var v22 uint64 = 1
        _ = v22
    case "hello":
        var v23 string = (v6 + (v6 + (v6 + v6)))
        _ = v23
    }
    if (!(255 >= (^(v13)))) {
        v4 = v4
        v24 := v5[9223372036854775807]
        _ = v24
    } else {
        if (226 > uint16(v3)) {
            var v25 string = "foo bar"
            _ = v25
        } else {
            for fi26 := 0; fi26 < 6; fi26++ {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := GenS1{F0: (v4 - (v4 - 2)), F1: 1, F2: (!(((1e10 + 0.0) - ((1.5 - ((-2.25 / 1.0) / 1.0)) + 3.14159)) != -2.25))}
                    if (v13 == v13) {
                        v29 := GenS1{F0: v4, F1: 2147483647, F2: (true || (!v14))}
                        var v30 float32 = (0.1 / 1e10)
                        v31 := []int{fi27}
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            var v32 int = 9223372036854775807
            _ = v32
        }
        v33 := GenS2{F0: (v3 << ((4294967295) & 63)), F1: 0, F2: (v3 / ((2) | 1))}
        _ = v33
    }
    v1 = (0.1 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
