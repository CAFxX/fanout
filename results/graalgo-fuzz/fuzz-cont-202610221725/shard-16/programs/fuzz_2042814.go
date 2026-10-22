package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint32 }
func genFn0(p0 string, p1 uint, p2 string) bool {
    v4 := GenS1{F0: 9223372036854775807, F1: -9223372036854775808, F2: "a"}
    var v5 int64 = -1
    if (!true) {
        v6 := []int{}
        _ = len(v6)
        _ = v6
    }
    v7 := GenS1{F0: (v5 | (-(v5))), F1: v5, F2: p0}
    v8 := []int{-2147483648, 42, -741}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return (!false)
}
func genFn1(p0 bool, p1 int8, p2 int) bool {
    v9 := []int{(p2 - 4294967295)}
    fmt.Println(p0, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return ("foo bar" != "\n\t")
}

func main() {
    var v1 int8 = 86
    v2 := []uint32{2}
    v3 := make(map[int]string)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint64 = 372
        v12 := v3[(fi10 + 0)]
        _ = fi10
        _ = v11
        _ = v12
    }
    v1 += v1
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint32 = 2147483647
        if (1 <= uint16(v14)) {
            v15 := GenS1{F0: 1, F1: -1, F2: "a"}
            _ = v15
        }
        switch "a" {
        case "\n\t":
            var v16 bool = (-2147483648 <= 42)
            _ = v16
        case "":
            var v17 string = ""
            _ = v17
        default:
            v18 := GenS1{F0: 947, F1: 42, F2: "hello"}
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    v1 += v1
    var v19 int = -521
    v20 := GenS2{F0: v1, F1: 127}
    if true {
        if (!(("hello" == ("" + "foo bar")) || (v19 != v19))) {
            v21 := GenS1{F0: 159, F1: 9223372036854775807, F2: ("foo bar" + "a")}
            v2 = append(v2, 323)
            if true {
                _ = len(v3)
            }
            _ = v21
        }
        v19 += (v19 / ((v19) | 1))
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        v2 = append(v2, 2147483647)
        for fi23 := 0; fi23 < 12; fi23++ {
            if (!(((int64(fi22) > 9223372036854775807) || (42 <= 42)) || (false && (37 == 1)))) {
                v24 := v3[v19]
                fmt.Println(v1)
                _ = v24
            } else {
                v19 = fi23
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        _ = len(v3)
        v19 -= fi25
        _ = fi25
    }
    v26 := GenS2{F0: v1, F1: 255}
    if true {
        switch "a" {
        case "a":
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 float32 = 3.14159
                var v29 uint64 = 0
                v30 := GenS2{F0: v1, F1: 243}
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        case "foo bar":
            v19 -= 42
        default:
            v31 := GenS3{F0: 1, F1: 752}
            _ = v31
        }
    }
    var v32 uint = 42
    var v33 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v32
    _ = v33
}
