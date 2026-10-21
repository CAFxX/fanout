package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string, p1 int) uint32 {
    v7 := GenS1{F0: -529}
    var v8 uint16 = 255
    if (!true) {
        var v9 int32 = 0
        var v10 int32 = (v9 + v9)
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{(("" + "a") + "\n\t"), ("a" + ("a" + "a")), "foo bar", "a"}
    v3 := []uint32{714}
    v4 := make(map[int]uint32)
    var v5 int = 1
    var v6 uint = 65535
    v11 := GenS2{F0: 951}
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v12 uint8 = 122
    switch (-(v5)) {
    case 0:
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13)
            _ = fi13
        }
    case 4294967295:
        if false {
            var v14 uint64 = 42
            _ = len(v2)
            _ = v14
        } else {
            var v15 bool = false
            if (0 > int8(v12)) {
                v12 = 2
                v12 = (v12 | (v12 - v12))
                var v16 uint32 = 0
                _ = v16
            }
            _ = v15
        }
    case 128:
        v17 := GenS1{F0: -2147483648}
        _ = v17
    }
    v6 = (^(v6))
    v18 := GenS1{F0: 42}
    if false {
        v5 = (v5 ^ v5)
        fmt.Println(v12)
    } else {
        var v19 int = (v5 / (((v5 * 196)) | 1))
        v5 = 0
        _ = v19
    }
    v20 := GenS1{F0: 0}
    var v21 uint32 = 738
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 string = "foo bar"
        var v24 int16 = 1
        _ = fi22
        _ = v23
        _ = v24
    }
    if false {
        v5 = v5
        fmt.Println(v12, v12)
    } else {
        switch "a" {
        case "":
            v4[v5] = genFn0("hello", v5)
        case "hello":
            var v25 uint8 = v12
            _ = v25
        default:
            v21 = v21
        }
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 uint64 = 1
        if true {
            v28 := GenS1{F0: -2147483648}
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 uint64 = (v27 % ((0) | 1))
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        delete(v1, v21)
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
