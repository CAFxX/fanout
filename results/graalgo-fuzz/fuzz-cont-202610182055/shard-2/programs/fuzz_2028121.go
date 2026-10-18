package main

import (
    "fmt"
)

func genFn0(p0 string) int {
    fmt.Println(p0)
    p0 = p0
    var v4 uint8 = 42
    _ = p0
    _ = v4
    return 2147483647
}
func genFn1(p0 int16, p1 int16) int {
    if (true || true) {
        var v5 int32 = -2147483648
        fmt.Println(p0)
        v6 := []int{0, 32768}
        _ = v5
        _ = v6
    } else {
        var v7 uint32 = 891
        _ = v7
    }
    p1 = p0
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{fi8}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 int32 = 42
    var v2 int16 = 1000
    v3 := make(map[uint32]uint32)
    if false {
        switch ("a" + "foo bar") {
        case "foo bar":
            var v10 int8 = -70
            _ = v10
        }
        v11 := []int{-373, 32768}
        _ = v11
    }
    v2 = v2
    var v12 uint32 = 65535
    v1 = (v1 + v1)
    v12 = (^(v12))
    var v13 int64 = 42
    fmt.Println(v12, v1, v13)
    v12 = v12
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
