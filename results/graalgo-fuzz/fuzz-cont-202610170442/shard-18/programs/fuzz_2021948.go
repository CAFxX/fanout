package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 string, p1 uint64, p2 uint16) uint {
    fmt.Println(p0)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int8, p1 int32, p2 uint32) uint32 {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: ("" + ""), F1: p0}
        p1 = 42
        v7 := []int{fi5, 2147483647}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "hello"
    v3 := []string{v2, v2, v2}
    var v4 uint32 = 879
    for fi8 := 0; fi8 < 6; fi8++ {
        v4 -= v4
        fmt.Println(fi8, v2)
        if true {
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
        _ = fi8
    }
    v10 := GenS2{F0: 42}
    v4 = (v4 + 42)
    var v11 string = "foo bar"
    for fi12 := 0; fi12 < 19; fi12++ {
        v3 = append(v3, "a")
        switch v2 {
        case "foo bar":
            var v13 uint32 = genFn1(-128, 2147483647, (v4 ^ v4))
            _ = v13
        default:
            var v14 string = ""
            _ = v14
        }
        v4 += v4
        _ = fi12
    }
    switch ("\n\t" + (v2 + v2)) {
    case "foo bar":
        v3 = append(v3, (v2 + v11))
    case "hello":
        v4 += v4
    case "\n\t":
        v4 -= 1
    }
    fmt.Println(v2, v11)
    v11 = v2
    var v15 uint64 = 491
    if false {
        var v16 int = 1
        if false {
            v11 = (v11 + "\n\t")
            for fi17 := 0; fi17 < 1; fi17++ {
                v1[fi17] = -362
                var v18 uint64 = v15
                _ = fi17
                _ = v18
            }
            delete(v1, 108)
        }
        _ = v16
    } else {
        switch v2 {
        case "\n\t":
            switch 369 {
            case 32768:
                v19 := GenS3{F0: 144}
                _ = v19
            case 0:
                v15 = 42
            }
        case "":
            v11 = (v2 + v11)
        case "foo bar":
            var v20 uint8 = 1
            _ = v20
        default:
            delete(v1, 32767)
        }
    }
    switch 2147483647 {
    case -243:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    default:
        _ = len(v3)
    }
    var v21 uint = 4294967295
    _ = len(v3)
    v22 := v1[364]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
