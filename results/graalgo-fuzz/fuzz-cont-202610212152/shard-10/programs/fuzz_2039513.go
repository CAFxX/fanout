package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint8, p1 int64) uint8 {
    p0 = p0
    fmt.Println(p0, p1)
    fmt.Println(p1, p0)
    p0 += 0
    var v5 int32 = 1
    _ = p0
    _ = p1
    _ = v5
    return 255
}

func main() {
    var v1 int32 = -479
    var v2 bool = true
    v3 := make(map[uint32]string)
    var v4 int32 = -2147483648
    v1 -= v1
    var v6 string = ("a" + (("foo bar" + "") + ("hello" + ("foo bar" + "\n\t"))))
    if (((0 > 617) && ((0 >= 100) || (v1 >= v4))) || v2) {
        v7 := []int{697}
        if (1 > 4294967295) {
            v6 = (v6 + "hello")
            if v2 {
                var v8 string = v6
                _ = v8
            }
        }
        _ = v7
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 bool = v2
            var v11 bool = v10
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(v6)
        switch 1 {
        case 629:
            v2 = (!(v2 && v2))
        case 9223372036854775807:
            v12 := GenS1{F0: 497, F1: (v1 ^ v1)}
            _ = v12
        }
    }
    v6 = v6
    var v13 float64 = 1e10
    var v14 int = -259
    var v15 float32 = 1e10
    if v2 {
        fmt.Println(v15, v1)
    } else {
        if (30 != -365) {
            if false {
                var v16 int = v14
                _ = v16
            }
            switch "" {
            case "foo bar":
                var v17 uint = 65535
                _ = v17
            case "a":
                var v18 int32 = v1
                _ = v18
            case "\n\t":
                v14 = v14
            }
        }
        var v19 uint = 2
        _ = v19
    }
    var v20 uint16 = 65
    if false {
        var v21 int8 = 101
        _ = v21
    } else {
        v22 := []int{v14, 0}
        v15 = 0.0
        _ = v22
    }
    v1 -= -1
    fmt.Println(v2, v20)
    v23 := []int{v14, (v14 >> ((14) & 63))}
    var v24 int32 = (-(v4))
    for fi25 := 0; fi25 < 6; fi25++ {
        v15 = ((v15 / (float32(1.5) + 3.14159)) - (0.1 * v15))
        v26 := GenS2{F0: 65535, F1: 1}
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v6
}
