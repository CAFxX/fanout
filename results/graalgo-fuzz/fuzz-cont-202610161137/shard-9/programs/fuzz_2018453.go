package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int16) int32 {
    var v6 uint16 = 816
    v7 := []int{-828, 42}
    v8 := GenS1{F0: 42}
    var v9 string = "foo bar"
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int8, p1 uint8) uint32 {
    if (false || true) {
        if (!true) {
            if (-1 <= genFn0(-32768)) {
                p0 = p0
            }
            p1 = p1
        }
        p1 += (^(p1))
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{(fi10 - 9223372036854775807), 505, 42}
            _ = fi10
            _ = v11
        }
    } else {
        if (127 != (p0 >> ((65535) & 63))) {
            var v12 float64 = 1e10
            var v13 int32 = genFn0(int16(p1))
            fmt.Println(p0, v13, p0)
            _ = v12
            _ = v13
        }
    }
    p0 -= (p0 << ((806) & 63))
    var v14 uint = 65535
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint64 = 115
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v14
    return 751
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 1000
    v3 := []int{-1, 1, 9223372036854775807}
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = v4
    fmt.Println(v2)
    v4 = v5
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            v2 = -1
            v5 += (v4 | v5)
            _ = fi18
        }
        v5 = (v5 ^ 0)
        for fi19 := 0; fi19 < 10; fi19++ {
            v4 = (v5 * v4)
            _ = fi19
        }
        _ = fi17
    }
    v1[65535] = 4294967295
    v4 += (v5 << ((2) & 63))
    v1[903] = 65535
    var v20 uint8 = 2
    if true {
        _ = len(v3)
        v4 += 4294967295
    }
    v20 = (v20 * (v20 * v20))
    switch (("" + "") + ("a" + "hello")) {
    case "foo bar":
        var v21 int32 = -2147483648
        _ = v21
    case "hello":
        if (127 >= 0) {
            v5 = (v5 << ((4294967295) & 63))
        }
    default:
        v22 := GenS1{F0: 4294967295}
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint8 = (v20 >> ((2) & 63))
        fmt.Println(v4)
        _ = fi23
        _ = v24
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
