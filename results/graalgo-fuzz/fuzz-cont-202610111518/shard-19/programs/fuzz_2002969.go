package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 = p0
        fmt.Println(p0)
        var v5 int16 = (p0 >> ((618) & 63))
        _ = fi4
        _ = v5
    }
    if (0 < 32767) {
        v6 := GenS1{F0: (false || ((("\n\t" + "") + "foo bar") < ""))}
        var v7 uint = 2
        _ = v6
        _ = v7
    }
    var v8 float32 = (-2.25 + float32(((1e10 * 1e10) - 1e10)))
    _ = p0
    _ = v8
    return "a"
}
func genFn1(p0 uint, p1 uint32, p2 uint) int8 {
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int32 = 42
        v11 := GenS2{F0: 42}
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 = (p2 / ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    var v1 int64 = 42
    var v2 int32 = -885
    v3 := make(map[uint32]int)
    for fi12 := 0; fi12 < 3; fi12++ {
        if (!(!(-128 > genFn1(42, 42, 2)))) {
            var v13 bool = false
            _ = v13
        } else {
            v14 := GenS3{F0: 128}
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint8 = 42
            v17 := GenS1{F0: ((!(!(!(fi15 == (fi12 << ((42) & 63)))))) && false)}
            var v18 int32 = (^(v2))
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi12
    }
    var v19 uint64 = 1
    fmt.Println(v19)
    fmt.Println(v19)
    v20 := []int{1, 181, -9223372036854775808}
    var v21 int64 = (v1 ^ (v1 * v1))
    v22 := GenS4{F0: 2}
    switch 350 {
    case -2147483648:
        v23 := GenS5{F0: -128}
        _ = v23
    case -9223372036854775808:
        fmt.Println(v2, v21)
    case 502:
        if (!false) {
            v19 += v19
            v24 := v3[1]
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
