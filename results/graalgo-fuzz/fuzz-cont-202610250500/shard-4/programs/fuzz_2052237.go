package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 42
    v3 := make(map[uint32]int)
    var v4 int32 = -2147483648
    v3[(v2 * (v2 * (v2 >> ((101) & 63))))] = -146
    var v5 float32 = (-2.25 + 0.1)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(v2, fi6)
        _ = fi6
    }
    v7 := []int{32768, 1, 32767}
    delete(v1, "foo bar")
    if true {
        if false {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int64 = -133
                fmt.Println(v2)
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint16 = 700
                v12 := v1[("foo bar" + "\n\t")]
                v11 += v11
                _ = fi10
                _ = v11
                _ = v12
            }
            v5 = float32(0.0)
        }
        fmt.Println(v2, v5)
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        delete(v3, v2)
        v14 := GenS1{F0: -1}
        v5 = v5
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: v2}
    v7 = append(v7, -9223372036854775808)
    v16 := GenS2{F0: (v2 & (v2 << ((686) & 63)))}
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint64 = 41
        switch "hello" {
        case "foo bar":
            v1["\n\t"] = "foo bar"
        default:
            for fi19 := 0; fi19 < 12; fi19++ {
                if true {
                    if len(v7) > 0 {
                        v7[0] = (^(fi17))
                    }
                } else {
                    fmt.Println(v18, v18)
                }
                v4 = (v4 ^ 0)
                if (-128 > 42) {
                    v5 += 0.1
                }
                _ = fi19
            }
        }
        v20 := GenS3{F0: 881, F1: 28, F2: 32767}
        _ = fi17
        _ = v18
        _ = v20
    }
    fmt.Println(v2)
    v4 = v4
    var v21 bool = (false && true)
    _ = len(v7)
    v2 = (^(v2))
    v5 = float32(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
