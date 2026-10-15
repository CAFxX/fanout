package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 uint32 }
type GenS3 struct { F0 int; F1 uint32; F2 int }
func genFn0(p0 uint16) int {
    v4 := GenS1{F0: 1, F1: -513}
    p0 -= p0
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{}
        for fi9 := 0; fi9 < 13; fi9++ {
            switch 1 {
            case 9223372036854775807:
                var v10 int16 = -1
                _ = v10
            case 32767:
                if ((!(!(4294967295 != 833))) && false) {
                    var v11 bool = true
                    _ = v11
                } else {
                    v8 = append(v8, fi9)
                }
            }
            fmt.Println(fi7, fi7)
            var v12 int64 = -123
            _ = fi9
            _ = v12
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi13
        }
        _ = fi7
        _ = v8
    }
    var v14 float32 = 0.1
    _ = p0
    _ = v14
    _ = v4
    return -596
}
func genFn1(p0 int8, p1 uint, p2 uint) uint {
    fmt.Println(p0)
    var v15 uint32 = 42
    fmt.Println(v15, p1)
    var v16 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"hello", "foo bar", ("foo bar" + ((("" + "") + ("hello" + "a")) + ""))}
    var v3 bool = true
    var v17 int8 = -1
    v3 = (!v3)
    v17 += (v17 % ((-61) | 1))
    v18 := GenS2{F0: 501, F1: 0, F2: 42}
    if v3 {
        switch "" {
        case "":
            v19 := GenS3{F0: 42, F1: 2, F2: 0}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 bool = (!(v3 && (1 <= 168)))
        _ = fi20
        _ = v21
    }
    switch -818 {
    case 765:
        fmt.Println(v17)
    }
    var v22 uint16 = 432
    v23 := GenS2{F0: 0, F1: 42, F2: 2147483648}
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
