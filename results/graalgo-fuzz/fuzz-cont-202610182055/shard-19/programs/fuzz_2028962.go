package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string, p1 int64, p2 uint8) uint16 {
    p2 += p2
    fmt.Println(p2, p0)
    v7 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 int8, p1 int16) uint {
    switch 32768 {
    case -232:
        var v8 uint32 = 1
        _ = v8
    case 127:
        v9 := []int{}
        _ = v9
    default:
        switch ("hello" + "") {
        case "\n\t":
            v10 := []int{2147483647}
            _ = v10
        }
    }
    var v11 bool = (!true)
    _ = p0
    _ = p1
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := []uint32{65535, 65535, 774, 0}
    var v2 float64 = 0.0
    v3 := []uint32{2147483648, 865, 2147483648}
    v4 := make(map[int]int)
    var v5 uint = 65535
    v6 := []int{-1, 32768}
    v5 = (^(v5))
    fmt.Println(v2, v5)
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int8 = 1
        _ = fi12
        _ = v13
    }
    if (42 > -3) {
        var v14 uint8 = 110
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v15 := GenS1{F0: 106}
    var v16 uint = (v5 & genFn1(84, -11))
    switch 4294967295 {
    case 32767:
        var v17 int64 = -1
        _ = v17
    default:
        if (!(true && (637 <= 992))) {
            var v18 float32 = (3.14159 / float32(v2))
            v5 = (v16 - (^(v5)))
            v19 := v4[127]
            _ = v18
            _ = v19
        } else {
            var v20 int16 = -1
            v16 += 0
            if (!(677 != 2147483647)) {
                v20 = v20
            } else {
                var v21 float32 = 0.0
                _ = v21
            }
            _ = v20
        }
    }
    var v22 int16 = 829
    switch "foo bar" {
    case "\n\t":
        if (((v5 >> ((255) & 63)) <= 883) && (((0 != -9223372036854775808) && (255 >= 0)) || (1 != 65535))) {
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 string = "\n\t"
                var v25 uint32 = 2
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 bool = ((true || true) && (0 != 2147483647))
            _ = v26
        } else {
            v22 = v22
            var v27 int32 = 0
            _ = v27
        }
    case "":
        v2 = v2
    case "foo bar":
        if (true || ((((v22 % ((v22) | 1)) != (v22 * 1)) && false) && true)) {
            _ = len(v4)
            var v28 uint8 = 0
            _ = v28
        }
    default:
        v2 = (3.14159 + float64(0.0))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
