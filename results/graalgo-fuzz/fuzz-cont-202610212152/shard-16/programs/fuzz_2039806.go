package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string) uint16 {
    v4 := []int{-2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 856
}
func genFn1(p0 int8, p1 uint32) uint8 {
    var v5 uint8 = 0
    if true {
        if (!(((false || false) && (p0 == -62)) && (p0 < -44))) {
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 int8 = p0
                var v8 int32 = 503
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        var v9 int32 = -2147483648
        _ = v9
    }
    v10 := []int{-244}
    if (!true) {
        _ = len(v10)
        var v11 uint64 = 18446744073709551615
        _ = v11
    }
    var v12 int32 = -1
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    _ = v5
    return v5
}

func main() {
    v1 := []uint32{2147483647, 472}
    var v2 uint32 = 2
    v3 := make(map[uint32]bool)
    switch 132 {
    case 2147483647:
        fmt.Println(v2)
    case 127:
        v2 = 4294967295
    default:
        if len(v1) > 0 {
            v1[0] = (v2 - (^(v2)))
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint8 = 255
        _ = fi13
        _ = v14
    }
    v2 = (v2 << ((4294967295) & 63))
    switch ("a" + "hello") {
    case "hello":
        v15 := GenS1{F0: 130}
        _ = v15
    default:
        v2 += v2
    }
    _ = len(v1)
    delete(v3, v2)
    if (0.1 != float32(float64((0.1 / 1.0)))) {
        v16 := GenS1{F0: 65535}
        fmt.Println(v2)
        v2 = (^(v2))
        _ = v16
    }
    fmt.Println(v2)
    switch (("foo bar" + "\n\t") + "hello") {
    case "foo bar":
        switch 1 {
        case -670:
            var v17 bool = true
            _ = v17
        }
    default:
        v2 -= (v2 + v2)
    }
    if false {
        for fi18 := 0; fi18 < 6; fi18++ {
            switch ((("a" + "a") + "a") + "") {
            case "a":
                fmt.Println(fi18, v2)
            case "":
                var v19 uint = 370
                _ = v19
            }
            _ = fi18
        }
    } else {
        if (!false) {
            if true {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v2 = v2
                    _ = fi20
                }
            }
        }
        if ((128 > -283) || (!true)) {
            var v21 string = "foo bar"
            _ = v21
        } else {
            v2 -= (v2 % ((523) | 1))
            v2 = v2
            fmt.Println(v2)
        }
    }
    v2 = (v2 % (((v2 | v2)) | 1))
    fmt.Println(v2)
    switch -129 {
    case 128:
        var v22 uint64 = 793
        _ = v22
    case -66:
        v1 = append(v1, (v2 >> ((2) & 63)))
    case -1:
        var v23 int32 = 2147483647
        _ = v23
    }
    v3[1] = ((42 < 442) || (970 <= 1))
    v2 = (v2 - (v2 & 2147483648))
    var v24 int64 = 1
    switch "\n\t" {
    case "":
        delete(v3, (v2 / ((v2) | 1)))
    }
    v2 -= 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
