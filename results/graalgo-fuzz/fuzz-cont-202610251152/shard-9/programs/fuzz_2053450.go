package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32) int16 {
    v5 := []int{-1, -1}
    var v6 int = 155
    fmt.Println(p0)
    v5 = append(v5, v6)
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := []int{}
    var v2 uint8 = 107
    var v3 uint = 255
    v4 := []uint32{4294967295}
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: uint64(v2)}
        if len(v4) > 0 {
            v4[0] = 0
        }
        _ = fi7
        _ = v8
    }
    if true {
        _ = len(v4)
    } else {
        var v9 uint = (^(v3))
        _ = v9
    }
    if (v3 > v3) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v2 -= (v2 << ((v3) & 63))
        var v10 uint = v3
        fmt.Println(v2, v10)
        _ = v10
    }
    v4 = append(v4, 2)
    var v11 int16 = genFn0(0)
    v11 = 802
    switch -1 {
    case 0:
        fmt.Println(v3, v2, v3)
    }
    v2 = 128
    switch 32767 {
    case 1:
        v12 := GenS1{F0: 0}
        _ = v12
    default:
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v11, v3)
            _ = fi13
        }
    }
    v3 -= (v3 + v3)
    var v14 uint64 = 483
    var v15 int16 = -1
    fmt.Println(v14)
    fmt.Println(v2)
    fmt.Println(v14, v3, v14)
    switch (("" + ("hello" + "")) + (("foo bar" + (("" + "") + "")) + ("hello" + (("foo bar" + (("\n\t" + "foo bar") + ("hello" + ("foo bar" + ("a" + "\n\t"))))) + "")))) {
    case "":
        for fi16 := 0; fi16 < 12; fi16++ {
            if (!((^(v3)) < (v3 >> ((v3) & 63)))) {
                v15 -= v11
                if (fi16 < (-(fi16))) {
                    var v17 int32 = 42
                    v18 := GenS1{F0: v14}
                    var v19 bool = (!true)
                    _ = v17
                    _ = v18
                    _ = v19
                }
                var v20 int64 = -1
                _ = v20
            } else {
                var v21 int8 = 0
                if (false || true) {
                    v1 = append(v1, fi16)
                }
                v1 = append(v1, (fi16 + (fi16 / (((fi16 & (fi16 * fi16))) | 1))))
                _ = v21
            }
            _ = fi16
        }
    case "hello":
        fmt.Println(v14)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
