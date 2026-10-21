package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint8, p1 uint8, p2 uint16) int8 {
    v5 := []int{}
    if false {
        v5 = append(v5, -9223372036854775808)
        var v6 bool = (!(true || ((0.1 == -2.25) || (false && (0 <= 100)))))
        v6 = true
        _ = v6
    }
    v7 := GenS1{F0: 581, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 123
}
func genFn1(p0 string, p1 int8) uint16 {
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 = p0
        var v9 uint16 = 1
        _ = fi8
        _ = v9
    }
    p0 = (("" + "") + (("foo bar" + "") + "a"))
    v10 := GenS2{F0: 32767}
    _ = p0
    _ = p1
    _ = v10
    return 255
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 0
    v3 := make(map[int]uint32)
    var v4 float32 = 0.0
    var v11 float32 = v4
    v1 = (v1 << ((0) & 63))
    var v12 uint8 = 1
    var v13 float32 = v11
    var v14 string = (("\n\t" + "\n\t") + (("hello" + ("hello" + "foo bar")) + (("hello" + (("a" + (("\n\t" + (("\n\t" + "") + ("foo bar" + ""))) + "")) + "foo bar")) + ("foo bar" + "foo bar"))))
    if (!((v2 * (v2 & 0)) >= (v2 + v2))) {
        var v15 int = 4294967295
        _ = v15
    } else {
        switch (v14 + "") {
        case "a":
            if false {
                var v16 uint32 = (v2 & v2)
                _ = v16
            }
        default:
            fmt.Println(v1, v2, v14)
        }
        if (v14 >= "foo bar") {
            if (!false) {
                var v17 int = -2147483648
                v17 += (v17 >> ((4294967295) & 63))
                _ = v17
            }
            switch v14 {
            case "":
                if true {
                    v4 = v13
                }
            default:
                fmt.Println(v11, v2, v13)
            }
            v1 = genFn0((v12 + v12), (v12 & v12), 1)
        } else {
            v18 := GenS3{F0: (v1 + 78), F1: v2}
            _ = v18
        }
        var v19 bool = (v11 <= 1.5)
        _ = v19
    }
    v20 := v3[127]
    if false {
        v21 := GenS2{F0: -1}
        _ = v21
    } else {
        v2 -= (^(v2))
        var v22 uint32 = (v20 * v20)
        _ = v22
    }
    var v23 string = v14
    for fi24 := 0; fi24 < 16; fi24++ {
        v20 -= (v20 * (v2 & (v20 | (v20 / ((v2) | 1)))))
        _ = fi24
    }
    v25 := []int{4294967295, 503, -1}
    var v26 int32 = -195
    fmt.Println(v12)
    if (v2 != v20) {
        v27 := GenS2{F0: 0}
        v20 = 2147483648
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
