package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0() uint {
    var v6 int = 32768
    v7 := []int{v6, 128}
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 uint32, p2 uint64) uint32 {
    p0 -= 686
    switch 128 {
    case 835:
        v8 := []int{-9223372036854775808, -2147483648, 127, 2147483647}
        _ = v8
    case -620:
        p0 -= 327
    case -702:
        v9 := GenS1{F0: 1, F1: 1}
        _ = v9
    default:
        for fi10 := 0; fi10 < 18; fi10++ {
            if false {
                switch (("" + "a") + "a") {
                case "\n\t":
                    for fi11 := 0; fi11 < 10; fi11++ {
                        p1 = p1
                        _ = fi11
                    }
                case "a":
                    var v12 string = "foo bar"
                    _ = v12
                case "foo bar":
                    if (-2.25 == 1.5) {
                        var v13 int16 = 206
                        var v14 int = (-(fi10))
                        v13 -= (-(v13))
                        _ = v13
                        _ = v14
                    }
                }
                p0 -= 42
                if (!true) {
                    v15 := GenS1{F0: -9223372036854775808, F1: 0}
                    _ = v15
                }
            } else {
                fmt.Println(fi10)
                v16 := []int{fi10, 127, fi10, fi10}
                v17 := GenS2{F0: 42, F1: -891, F2: p2}
                _ = v16
                _ = v17
            }
            _ = fi10
        }
    }
    var v18 uint32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return uint32(p0)
}

func main() {
    var v1 uint8 = 42
    var v2 uint = 2
    v3 := []string{("foo bar" + (("hello" + "\n\t") + "hello")), "", "hello", "\n\t"}
    v4 := make(map[int]uint32)
    var v5 uint64 = 18446744073709551615
    v5 += v5
    v19 := GenS3{F0: 0, F1: ("\n\t" + "a")}
    v20 := GenS3{F0: v5, F1: "\n\t"}
    var v21 uint32 = 2147483648
    for fi22 := 0; fi22 < 3; fi22++ {
        if ((1.5 + 0.0) > -2.25) {
            var v23 uint8 = 42
            _ = v23
        }
        fmt.Println(v5, v2, fi22)
        _ = fi22
    }
    v2 = (v2 * (v2 << (((v2 % (((v2 % ((v2) | 1))) | 1))) & 63)))
    var v24 int64 = -9223372036854775808
    v1 -= (v1 | (^(v1)))
    v25 := GenS1{F0: int64(v1), F1: 0}
    v5 += (v5 + v5)
    v5 -= (v5 + (v5 * v5))
    var v26 uint = 667
    v27 := GenS2{F0: 1000, F1: -1, F2: (^(v5))}
    fmt.Println(v5, v5)
    v5 = (v5 - (v5 % (((v5 & v5)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
