package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 int32, p1 uint, p2 uint8) int64 {
    p1 = (p1 | p1)
    v6 := []int{9223372036854775807, -833}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return int64(p2)
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    var v5 int = 42
    delete(v4, ("a" + "a"))
    v5 += v5
    for fi7 := 0; fi7 < 3; fi7++ {
        v1 = v1
        if true {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: 9223372036854775807, F1: 17, F2: 1}
                v1 -= (v1 * v1)
                var v10 uint64 = 321
                _ = fi8
                _ = v10
                _ = v9
            }
            v1 -= v1
        } else {
            switch ("a" + ("\n\t" + ("a" + ("foo bar" + "foo bar")))) {
            case "hello":
                if false {
                    v11 := v3[fi7]
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 float64 = 3.14159
                        _ = fi12
                        _ = v13
                    }
                    v11 = (("" + v11) + (("hello" + v11) + v11))
                    _ = v11
                } else {
                    var v14 int8 = -1
                    var v15 int32 = 2147483647
                    _ = v14
                    _ = v15
                }
            case "":
                v1 = 65535
            case "a":
                if ((true && (false && false)) && true) {
                    fmt.Println(v1, v1)
                    v16 := GenS2{F0: 1000, F1: v1, F2: 0}
                    if false {
                        var v17 int = (^(v5))
                        v17 -= 127
                        _ = v17
                    } else {
                        fmt.Println(v5, v1)
                        v18 := []int{(v5 ^ (v5 | (^(fi7))))}
                        v5 += (fi7 & (^(v5)))
                        _ = v18
                    }
                    _ = v16
                }
            }
            var v19 bool = (-2.25 > 1e10)
            v1 = v1
            _ = v19
        }
        _ = fi7
    }
    var v20 float64 = -2.25
    v20 = v20
    v1 = 4294967295
    fmt.Println(v1, v5)
    switch v5 {
    case 608:
        fmt.Println(v5)
    default:
        var v21 uint = 967
        _ = v21
    }
    var v22 uint32 = v1
    for fi23 := 0; fi23 < 14; fi23++ {
        switch ("\n\t" + "") {
        case "foo bar":
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS1{F0: fi24, F1: 255, F2: 2}
                _ = fi24
                _ = v25
            }
        case "a":
            var v26 uint8 = 42
            _ = v26
        case "":
            if false {
                v5 -= 32767
                v5 -= (-(fi23))
            } else {
                var v27 bool = (!((4294967295 != 18446744073709551615) || true))
                v22 -= 2147483647
                _ = v27
            }
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
