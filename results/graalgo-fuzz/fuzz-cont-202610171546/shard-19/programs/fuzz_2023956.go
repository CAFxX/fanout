package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint, p1 int) int16 {
    var v5 uint8 = 36
    switch (^(p1)) {
    case 32768:
        v6 := []int{(p1 ^ 595)}
        _ = v6
    case 127:
        p0 = p0
    default:
        p1 = p1
    }
    p0 = (p0 << ((p0) & 63))
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint8) int32 {
    if false {
        var v7 uint8 = p0
        _ = v7
    } else {
        var v8 int16 = 42
        fmt.Println(p0, p0)
        _ = v8
    }
    p0 = 2
    if false {
        if false {
            if true {
                v9 := []int{-35, 32767, 127, -2147483648}
                _ = v9
            }
            v10 := []int{9223372036854775807, 606}
            _ = v10
        }
    } else {
        p0 -= p0
        p0 -= 140
        var v11 int = 2147483647
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            p0 += p0
            var v14 int64 = -9223372036854775808
            var v15 int64 = (v14 + v14)
            _ = fi13
            _ = v14
            _ = v15
        }
        p0 = p0
        _ = fi12
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 int8 = 127
    var v2 float32 = 1.5
    v3 := make(map[int]bool)
    var v4 uint16 = 65535
    v4 = (v4 % ((65535) | 1))
    for fi16 := 0; fi16 < 7; fi16++ {
        switch "foo bar" {
        case "hello":
            var v17 int32 = -2147483648
            _ = v17
        case "":
            v18 := []int{fi16, fi16, (fi16 >> ((65535) & 63))}
            _ = v18
        }
        var v19 uint64 = 703
        v2 += v2
        _ = fi16
        _ = v19
    }
    var v20 int32 = -1
    v21 := []int{-2147483648, 32767, 1}
    v22 := GenS1{F0: "\n\t"}
    switch "foo bar" {
    case "foo bar":
        delete(v3, -784)
    case "\n\t":
        for fi23 := 0; fi23 < 9; fi23++ {
            _ = len(v21)
            _ = fi23
        }
    default:
        v24 := GenS1{F0: "\n\t"}
        _ = v24
    }
    if false {
        var v25 int8 = (v1 * -1)
        _ = v25
    }
    v26 := GenS1{F0: (("hello" + "a") + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
