package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 bool, p1 int32) int64 {
    p0 = p0
    p1 = p1
    if (95 != 32767) {
        p1 -= -984
        v7 := GenS1{F0: 18446744073709551615}
        fmt.Println(p1)
        _ = v7
    }
    _ = p0
    _ = p1
    return 332
}
func genFn1() bool {
    v8 := []int{0, 32767, 42}
    var v9 float32 = 0.0
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 4294967295
    var v3 int8 = 120
    var v4 int16 = -1
    v5 := make(map[string]int)
    var v6 int16 = (v4 >> ((605) & 63))
    v10 := []int{-15, -517}
    v2 -= v2
    v6 -= (^(v4))
    v3 += 127
    v3 = (v3 ^ 100)
    fmt.Println(v4, v6, v2)
    switch -723 {
    case -9223372036854775808:
        switch 926 {
        case 0:
            fmt.Println(v6)
        case -708:
            v11 := GenS1{F0: 4294967295}
            _ = v11
        case -863:
            if (449 >= uint64(v2)) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    for fi13 := 0; fi13 < 15; fi13++ {
                        var v14 uint16 = 1
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS2{F0: 42}
                    var v16 uint8 = 128
                    _ = fi12
                    _ = v15
                    _ = v16
                }
            } else {
                fmt.Println(v4, v4, v3)
                if (int32(v2) != 37) {
                    fmt.Println(v6)
                } else {
                    v2 = (v2 >> ((255) & 63))
                    var v17 uint16 = 255
                    v18 := GenS3{F0: (v6 | (v6 - v6)), F1: (-(v4))}
                    _ = v17
                    _ = v18
                }
                var v19 int8 = -128
                _ = v19
            }
        }
    }
    if true {
        v2 = v2
    } else {
        v20 := GenS3{F0: (^(v6)), F1: (v6 << ((18446744073709551615) & 63))}
        _ = v20
    }
    switch "a" {
    case "foo bar":
        var v21 int = 0
        _ = v21
    case "hello":
        v2 -= (v2 >> ((255) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
