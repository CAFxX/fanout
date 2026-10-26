package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int, p2 uint) int32 {
    p1 += (p1 - 32767)
    fmt.Println(p1, p0, p0)
    v7 := GenS1{F0: 628}
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 2}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int16 = 476
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
        switch "a" {
        case "":
            v13 := []int{0}
            _ = v13
        case "foo bar":
            fmt.Println(p1, p2, p2)
        default:
            fmt.Println(p0, p0, p0)
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1() int {
    switch "" {
    case "\n\t":
        if false {
            if false {
                v14 := []int{}
                _ = v14
            }
            switch 546 {
            case 1:
                v15 := GenS1{F0: 402}
                _ = v15
            }
            var v16 uint32 = 1
            _ = v16
        }
    case "hello":
        v17 := []int{4294967295}
        _ = v17
    default:
        var v18 int16 = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := []int{-1, fi19, (-(fi19))}
        _ = fi19
        _ = v20
    }
    var v21 int64 = -9223372036854775808
    _ = v21
    return 4294967295
}

func main() {
    var v1 int8 = -83
    v2 := []string{}
    v3 := []uint32{0, 2147483647}
    var v4 int32 = 679
    var v5 uint32 = 65535
    v6 := make(map[string]string)
    var v22 string = "hello"
    var v23 int16 = 1
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 uint8 = 234
        v4 -= (v4 >> ((473) & 63))
        v25 -= 2
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        if (((v5 << ((255) & 63)) >= 839) && true) {
            var v27 int = fi26
            fmt.Println(v22, v5, v22)
            var v28 int = -2147483648
            _ = v27
            _ = v28
        }
        _ = fi26
    }
    var v29 uint16 = uint16(v5)
    _ = len(v6)
    v22 = v22
    if false {
        v22 = v22
        v22 = "foo bar"
    } else {
        v30 := GenS1{F0: v5}
        v4 += v4
        _ = v30
    }
    fmt.Println(v23)
    var v31 bool = (v22 != v22)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
