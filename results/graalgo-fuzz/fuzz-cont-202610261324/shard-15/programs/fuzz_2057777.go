package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
func genFn0(p0 int) uint16 {
    p0 = p0
    switch ("a" + ("" + "a")) {
    case "":
        switch ("a" + ("hello" + "foo bar")) {
        case "":
            fmt.Println(p0)
        case "\n\t":
            p0 = p0
        case "hello":
            if true {
                var v7 float32 = 1.5
                _ = v7
            }
        }
    }
    var v8 float32 = 0.1
    p0 -= p0
    if (p0 < (p0 & (p0 & (p0 * 4294967295)))) {
        switch "hello" {
        case "hello":
            v9 := GenS1{F0: 1, F1: 412}
            _ = v9
        case "foo bar":
            v10 := []int{(p0 - p0), 127, (p0 & p0), (p0 ^ (^(p0)))}
            _ = v10
        default:
            v8 += float32(float64(v8))
        }
        v11 := []int{(p0 * (p0 & p0))}
        if false {
            if (1 > 2) {
                if len(v11) > 0 {
                    v11[0] = (p0 >> ((1) & 63))
                }
                if len(v11) > 0 {
                    v11[0] = (-(p0))
                }
            }
        }
        _ = v11
    }
    _ = p0
    _ = v8
    return 77
}

func main() {
    v1 := []uint32{}
    var v2 int32 = 2147483647
    var v3 uint64 = 42
    var v4 uint32 = 988
    var v5 string = "hello"
    v6 := make(map[uint32]string)
    v2 = v2
    var v12 int64 = -1
    if true {
        v2 -= 2147483647
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint64 = (v3 + v3)
            switch (fi13 % ((750) | 1)) {
            case 128:
                var v15 uint32 = 254
                _ = v15
            case -9223372036854775808:
                v3 = 0
            }
            _ = fi13
            _ = v14
        }
        v2 = 42
    }
    if false {
        switch -50 {
        case 2147483647:
            var v16 uint8 = 2
            _ = v16
        default:
            v5 = (v5 + (v5 + ("a" + "")))
        }
        var v17 uint8 = 2
        v4 -= 42
        _ = v17
    }
    v6[(v4 & v4)] = "foo bar"
    v2 = 507
    v12 = (v12 | (v12 << ((42) & 63)))
    var v18 uint = 65535
    var v19 int64 = (v12 >> ((337) & 63))
    v5 = (v5 + "")
    v19 += v19
    v4 = v4
    v2 = v2
    var v20 uint64 = v3
    fmt.Println(v19, v18, v5)
    v21 := GenS1{F0: 67, F1: -32768}
    var v22 string = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
