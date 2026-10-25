package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16, p1 uint, p2 int8) uint64 {
    p2 = 42
    var v4 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}
func genFn1(p0 uint32, p1 uint8, p2 uint) uint {
    v5 := []int{2147483647, 510}
    if (!true) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch "a" {
        case "a":
            var v6 bool = (!(42 != int64(p1)))
            _ = v6
        case "\n\t":
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 bool = ((^(fi7)) == -844)
                var v9 string = "hello"
                _ = fi7
                _ = v8
                _ = v9
            }
        case "hello":
            var v10 uint = 934
            _ = v10
        }
        var v11 int8 = -128
        _ = v11
    }
    p2 += (p2 | p2)
    p2 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (3.14159 - (3.14159 / 1e10))
    v3 := []int{-9223372036854775808, -2147483648}
    fmt.Println(v2)
    switch 32767 {
    case 128:
        if ((452 <= 0) && ((("" + "\n\t") + ("\n\t" + ("" + ("hello" + "")))) <= "a")) {
            if (!(v2 == (1.5 / v2))) {
                var v12 string = "\n\t"
                var v13 uint8 = 0
                _ = v12
                _ = v13
            }
            var v14 bool = true
            v14 = (735 == 0)
            _ = v14
        }
    default:
        v1[2147483648] = 4294967295
    }
    v2 = (v2 + v2)
    var v15 uint32 = 2
    v15 -= v15
    fmt.Println(v2, v15)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    v1[(v15 | 1)] = v15
    v2 += float64(0.1)
    var v16 uint = 0
    fmt.Println(v15, v2)
    v17 := GenS1{F0: 1}
    switch 128 {
    case 32767:
        var v18 float64 = 0.1
        _ = v18
    case -1:
        var v19 string = "foo bar"
        _ = v19
    case 2147483647:
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint8 = 128
            _ = fi20
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
