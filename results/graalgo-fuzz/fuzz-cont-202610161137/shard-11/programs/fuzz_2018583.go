package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 uint64) uint {
    switch "\n\t" {
    case "foo bar":
        p0 = p0
    case "":
        v6 := []int{51}
        _ = v6
    default:
        fmt.Println(p0)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 string = ("\n\t" + "\n\t")
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(p0, fi7, fi9)
            var v10 string = (v8 + (v8 + "hello"))
            v11 := []int{(^(fi9)), (fi9 & fi7), 127, fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS1{F0: ("\n\t" + "hello"), F1: p0}
    var v13 float32 = 3.14159
    p0 = p0
    _ = p0
    _ = v12
    _ = v13
    return 682
}
func genFn1(p0 uint32) int32 {
    v14 := GenS2{F0: 21, F1: 1}
    var v15 uint8 = 127
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int8 = -1
        _ = fi16
        _ = v17
    }
    p0 += 2
    _ = p0
    _ = v14
    _ = v15
    return -497
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{2147483647, 42, 32767}
    v3 := make(map[int]bool)
    var v4 uint = 4294967295
    var v5 int16 = 506
    v18 := GenS1{F0: "foo bar", F1: 9}
    if false {
        fmt.Println(v5, v5)
        switch "" {
        case "a":
            v19 := GenS3{F0: 0, F1: "\n\t"}
            _ = v19
        case "":
            _ = len(v3)
        case "hello":
            var v20 string = ("foo bar" + "")
            _ = v20
        }
    } else {
        v4 = v4
    }
    var v21 uint = (^(v4))
    var v22 string = "a"
    var v23 uint16 = 2
    if ((v5 << ((0) & 63)) == v5) {
        v23 -= v23
        var v24 uint32 = 65535
        _ = v24
    } else {
        var v25 int = 11
        for fi26 := 0; fi26 < 4; fi26++ {
            _ = len(v3)
            _ = fi26
        }
        var v27 uint32 = uint32(v23)
        _ = v25
        _ = v27
    }
    v4 = (v21 | 0)
    delete(v1, (((("foo bar" + v22) + ((v22 + "foo bar") + v22)) + "hello") + "foo bar"))
    v22 = (v22 + v22)
    delete(v3, 0)
    v5 = (v5 & (v5 | (v5 ^ (v5 % ((-1) | 1)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
