package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint32; F2 int16 }
func genFn0() int64 {
    var v6 int64 = 0
    switch ("" + "foo bar") {
    case "hello":
        v7 := []int{128, -677}
        _ = v7
    case "":
        fmt.Println(v6)
    default:
        v6 += (v6 ^ v6)
    }
    var v8 int16 = 32767
    v9 := []int{-9223372036854775808, 32768, -135, 0}
    _ = v6
    _ = v8
    _ = v9
    return v6
}
func genFn1(p0 uint32, p1 int64, p2 int64) int16 {
    p0 += (p0 % ((205) | 1))
    v10 := []int{42, 3, 32767}
    _ = len(v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1000
}

func main() {
    var v1 uint = 238
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v4 := []string{(("foo bar" + ("\n\t" + ("hello" + ""))) + ("hello" + ("\n\t" + "")))}
    var v5 int = 127
    v1 = v1
    var v11 uint16 = 192
    if len(v4) > 0 {
        v4[0] = "a"
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if ((-9223372036854775808 != 42) && false) {
            var v13 float64 = 0.1
            if false {
                fmt.Println(v13, v11)
                v1 += v1
                v1 = v1
            }
            v13 += 1e10
            _ = v13
        } else {
            _ = len(v4)
        }
        _ = fi12
    }
    fmt.Println(v11, v5)
    var v14 uint = 4294967295
    for fi15 := 0; fi15 < 17; fi15++ {
        switch 9223372036854775807 {
        case -268:
            fmt.Println(v11)
        case 1:
            v11 = 263
        case -905:
            v16 := GenS1{F0: -9223372036854775808, F1: 0}
            _ = v16
        }
        if (false && ((128 >= 1) || (923 == -217))) {
            _ = len(v4)
        } else {
            switch v5 {
            case 0:
                if (-277 == -523) {
                    if false {
                        delete(v2, 2147483647)
                        v14 -= v1
                    } else {
                        v2[1] = (v5 * int(v1))
                        switch "foo bar" {
                        case "\n\t":
                            var v17 float32 = 1e10
                            _ = v17
                        case "hello":
                            var v18 uint8 = 42
                            _ = v18
                        }
                        if true {
                            fmt.Println(fi15)
                        }
                    }
                    v19 := GenS2{F0: (v14 ^ v14), F1: 835, F2: 0}
                    _ = v19
                }
            default:
                v11 = v11
            }
            _ = len(v3)
            var v20 bool = (0.0 == 1.5)
            _ = v20
        }
        _ = fi15
    }
    var v21 uint8 = 255
    var v22 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
