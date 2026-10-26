package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
type GenS3 struct { F0 int64; F1 uint }
func genFn0() uint32 {
    v6 := GenS1{F0: 233, F1: 18446744073709551615, F2: 65535}
    var v7 uint16 = 985
    v7 = (v7 >> ((4294967295) & 63))
    if (false && false) {
        v8 := []int{}
        if true {
            var v9 int = 32767
            switch (v9 - -715) {
            case 32767:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            default:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = v9
        }
        v7 -= (v7 / ((255) | 1))
        _ = v8
    }
    switch "a" {
    case "a":
        v7 = v7
    case "":
        for fi10 := 0; fi10 < 7; fi10++ {
            if (2 <= 153) {
                if false {
                    v11 := GenS2{F0: 0, F1: 1, F2: 107}
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := []int{}
                        if true {
                            v7 = (v7 << ((0) & 63))
                            v7 -= (v7 * v7)
                        }
                        var v14 uint16 = (^(v7))
                        v14 -= (^(v14))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    var v15 uint64 = 138
                    _ = v11
                    _ = v15
                }
                switch ("foo bar" + ("foo bar" + "a")) {
                case "foo bar":
                    var v16 bool = true
                    _ = v16
                case "":
                    v17 := []int{(fi10 ^ (fi10 - fi10)), (fi10 + -366)}
                    _ = v17
                case "a":
                    if false {
                        v18 := []int{fi10, (fi10 / ((fi10) | 1)), (^(fi10)), (fi10 | 4294967295)}
                        _ = v18
                    }
                }
                v19 := []int{}
                _ = v19
            } else {
                fmt.Println(v7, v7)
            }
            _ = fi10
        }
    }
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 913
    var v3 uint8 = 1
    var v4 float64 = (0.0 - (((0.0 - float64(float32(0.0))) * 0.0) / 1e10))
    var v5 int32 = 42
    v20 := GenS2{F0: -9223372036854775808, F1: 42, F2: 33}
    if ((float64(((0.0 / (3.14159 + 3.14159)) / -2.25)) - (v4 + v4)) == 3.14159) {
        fmt.Println(v5)
        fmt.Println(v2)
        v2 -= v2
    } else {
        v21 := GenS1{F0: 42, F1: v2, F2: 2}
        var v22 uint8 = v3
        v23 := []int{-449}
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        if (!((!false) && false)) {
            var v25 string = ("a" + "")
            var v26 int = (fi24 - -9223372036854775808)
            _ = len(v1)
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS1{F0: 0, F1: v2, F2: uint16(v2)}
            fmt.Println(v4, v3)
            _ = fi27
            _ = v28
        }
        _ = fi24
    }
    v3 += v3
    v29 := GenS3{F0: -32, F1: 2}
    var v30 uint64 = v2
    switch 32767 {
    case -507:
        var v31 float32 = 0.1
        _ = v31
    case -991:
        v32 := GenS3{F0: 0, F1: 183}
        _ = v32
    default:
        v33 := []int{int(v3), 2147483647, 32768}
        _ = v33
    }
    v30 += (v30 - 504)
    var v34 uint64 = (v2 - (v30 - (v2 << ((18446744073709551615) & 63))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
