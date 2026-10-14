package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
type GenS2 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 int16) uint8 {
    fmt.Println(p0)
    v5 := GenS1{F0: 82, F1: "foo bar", F2: 18446744073709551615}
    v6 := []int{104, 0, 127}
    fmt.Println(p0)
    if (!((p0 >> ((711) & 63)) <= (^(p0)))) {
        fmt.Println(p0)
        var v7 bool = false
        fmt.Println(p0, v7)
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return 128
}
func genFn1() bool {
    var v8 uint = 2
    var v9 float32 = 3.14159
    v10 := []int{0, 122, -1}
    _ = v10
    _ = v8
    _ = v9
    return (!((true && true) || true))
}

func main() {
    v1 := []int{0}
    var v2 float32 = 1e10
    var v3 int32 = -392
    v4 := make(map[uint32]string)
    v3 -= (v3 / (((^(v3))) | 1))
    v3 -= v3
    v11 := GenS2{F0: -1, F1: 1, F2: 398}
    var v12 string = "a"
    var v13 uint32 = 997
    fmt.Println(v2, v13, v3)
    v14 := GenS1{F0: 18446744073709551615, F1: "hello", F2: 65535}
    var v15 int16 = 32767
    var v16 uint32 = (^(v13))
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 bool = false
        for fi19 := 0; fi19 < 19; fi19++ {
            if (v18 && v18) {
                fmt.Println(v18)
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v15 = v15
    v1 = append(v1, -9223372036854775808)
    v20 := v4[(v16 + 65535)]
    v21 := GenS3{F0: "\n\t", F1: 1}
    v22 := GenS1{F0: 759, F1: "foo bar", F2: 604}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
