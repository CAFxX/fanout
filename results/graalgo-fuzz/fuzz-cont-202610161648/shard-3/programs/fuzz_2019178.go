package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 string, p1 uint64, p2 int8) int {
    var v7 uint8 = 42
    v8 := GenS1{F0: (v7 + (v7 / (((v7 * 255)) | 1))), F1: 18446744073709551615, F2: -32768}
    p2 -= (p2 << ((42) & 63))
    p1 = 1
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{}
        _ = len(v10)
        fmt.Println(p0)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 string) int {
    p0 = ((p0 + p0) + p0)
    var v11 int8 = -77
    _ = p0
    _ = v11
    return 32767
}

func main() {
    var v1 int32 = -431
    var v2 uint16 = 1
    var v3 float32 = 3.14159
    var v4 int16 = -32768
    var v5 uint = 349
    v6 := make(map[uint32]string)
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 bool = ((1 >= 127) || false)
            _ = fi13
            _ = v14
        }
        if (v1 > -1) {
            var v15 int64 = 9223372036854775807
            _ = v15
        } else {
            var v16 int = genFn0("hello", 42, -124)
            _ = v16
        }
        _ = fi12
    }
    v2 += 516
    v17 := GenS2{F0: "", F1: int64(v4)}
    fmt.Println(v4, v2)
    v1 += 650
    switch "" {
    case "":
        v3 -= v3
    case "hello":
        v2 -= 255
    case "\n\t":
        v18 := GenS1{F0: 153, F1: 18446744073709551615, F2: (v4 << ((0) & 63))}
        _ = v18
    default:
        var v19 string = "foo bar"
        _ = v19
    }
    if false {
        v2 = 42
        v20 := []int{516, int(v5)}
        _ = v20
    }
    fmt.Println(v1, v5, v5)
    var v21 uint64 = 216
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v3)
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS2{F0: ("hello" + "a"), F1: 624}
            delete(v6, 2)
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        v5 += 18446744073709551615
        _ = fi25
    }
    v5 = (^(v5))
    for fi26 := 0; fi26 < 15; fi26++ {
        if (("foo bar" != ("a" + "foo bar")) || true) {
            var v27 uint32 = 561
            v28 := []int{fi26, genFn0("\n\t", 200, 42), (fi26 + genFn1("a")), 127}
            switch ("" + ("\n\t" + "")) {
            case "a":
                var v29 int8 = int8(v2)
                _ = v29
            case "":
                fmt.Println(fi26)
            case "hello":
                v4 = (-(v4))
            default:
                var v30 int8 = 1
                _ = v30
            }
            _ = v27
            _ = v28
        } else {
            fmt.Println(v2)
            for fi31 := 0; fi31 < 5; fi31++ {
                fmt.Println(fi31, v5, v2)
                _ = fi31
            }
        }
        _ = fi26
    }
    var v32 int64 = 0
    var v33 int16 = v4
    v34 := []int{42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
