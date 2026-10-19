package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 string, p1 uint16) int8 {
    if false {
        var v7 bool = true
        _ = v7
    } else {
        v8 := GenS1{F0: 65535}
        var v9 uint32 = 1
        var v10 string = p0
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int32 = 1
    v12 := []int{-104, -178, 1, 32768}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 105
}
func genFn1(p0 uint, p1 uint32, p2 int8) int16 {
    fmt.Println(p2, p1)
    var v13 int8 = genFn0("foo bar", 449)
    if (!true) {
        fmt.Println(p1, p2, v13)
    } else {
        var v14 uint32 = (p1 * p1)
        v15 := GenS1{F0: 976}
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS1{F0: 65535}
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    switch (("a" + "foo bar") + "") {
    case "":
        v18 := GenS1{F0: uint16(p0)}
        _ = v18
    }
    v19 := GenS1{F0: 780}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v19
    return 1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = (0.0 / (1e10 / ((-2.25 * 1.5) / 1.0)))
    var v3 int32 = -1
    var v4 uint8 = 0
    v5 := make(map[uint32]uint32)
    var v6 bool = true
    v6 = false
    v20 := []int{-2147483648, -91, 32768}
    var v21 float64 = 1e10
    for fi22 := 0; fi22 < 3; fi22++ {
        v21 = v21
        _ = fi22
    }
    var v23 float64 = v21
    v1 += v1
    if (2147483647 <= 2147483647) {
        switch -9223372036854775808 {
        case -890:
            v3 += -658
        default:
            switch "foo bar" {
            case "":
                v1 = 1
            case "hello":
                var v24 float32 = 1.5
                _ = v24
            case "foo bar":
                var v25 int64 = int64(v3)
                _ = v25
            }
        }
    } else {
        var v26 bool = true
        var v27 uint = 255
        _ = v26
        _ = v27
    }
    v23 += v23
    var v28 int32 = v3
    v29 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
