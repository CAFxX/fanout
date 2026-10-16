package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8) int {
    var v4 float32 = (3.14159 - (0.1 * -2.25))
    var v5 uint64 = 18446744073709551615
    _ = p0
    _ = v4
    _ = v5
    return -730
}
func genFn1(p0 int16, p1 uint32, p2 int8) string {
    p0 = 298
    v6 := []int{-782, 2147483647, -9223372036854775808, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}

func main() {
    var v1 int64 = 520
    var v2 string = "a"
    var v3 uint64 = 0
    var v7 uint64 = (^(v3))
    v8 := []int{}
    if false {
        v2 = v2
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: v2, F1: 663, F2: false}
            _ = fi9
            _ = v10
        }
        if false {
            var v11 bool = false
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            v8 = append(v8, (fi12 & (fi12 + (-(fi12)))))
            _ = fi12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            v2 = "foo bar"
            switch (fi13 << ((uint(v3)) & 63)) {
            case -527:
                if true {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v8 = append(v8, fi13)
                        _ = fi14
                    }
                }
            case -634:
                var v15 uint64 = (v3 << ((255) & 63))
                _ = v15
            case 1:
                _ = len(v8)
            default:
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := GenS2{F0: 42}
                    v3 -= v3
                    v3 -= 18446744073709551615
                    _ = fi16
                    _ = v17
                }
            }
            var v18 int32 = 42
            _ = fi13
            _ = v18
        }
    }
    if (("" + "\n\t") > (v2 + ((((v2 + v2) + "\n\t") + (v2 + v2)) + (v2 + v2)))) {
        for fi19 := 0; fi19 < 20; fi19++ {
            v7 = v7
            v1 = (v1 ^ v1)
            _ = fi19
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v20 := GenS2{F0: -2147483648}
    _ = len(v8)
    if (!false) {
        var v21 uint = 18446744073709551615
        v3 += (v7 >> ((42) & 63))
        _ = v21
    } else {
        fmt.Println(v7, v3, v2)
        switch v2 {
        case "foo bar":
            var v22 uint16 = 255
            _ = v22
        default:
            fmt.Println(v2)
        }
        switch 0 {
        case -9223372036854775808:
            var v23 int8 = 113
            _ = v23
        case 42:
            if len(v8) > 0 {
                v8[0] = -2147483648
            }
        }
    }
    var v24 uint64 = 0
    if ((-2147483648 <= 2147483647) && false) {
        for fi25 := 0; fi25 < 15; fi25++ {
            v7 = 2
            var v26 float32 = 0.0
            _ = fi25
            _ = v26
        }
        if true {
            v8 = append(v8, 0)
        } else {
            var v27 float32 = 0.0
            _ = v27
        }
    }
    var v28 int = 9223372036854775807
    v8 = append(v8, (v28 / ((1) | 1)))
    var v29 uint64 = (v3 - 42)
    v30 := GenS1{F0: v2, F1: -359, F2: ((-(v28)) > v28)}
    var v31 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v7
    _ = v8
}
