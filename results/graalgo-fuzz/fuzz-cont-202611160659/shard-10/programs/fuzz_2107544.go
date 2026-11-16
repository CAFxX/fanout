package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
type GenS4 struct { F0 uint; F1 int32; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
func genFn0(p0 bool, p1 int8, p2 uint32) int8 {
    switch "foo bar" {
    case "\n\t":
        for fi4 := 0; fi4 < 3; fi4++ {
            p0 = (!(!(!p0)))
            p1 = p1
            _ = fi4
        }
    case "a":
        v5 := GenS1{F0: (p2 ^ p2), F1: 0}
        _ = v5
    case "":
        fmt.Println(p1, p1)
    default:
        fmt.Println(p1, p0, p2)
    }
    if p0 {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: p2, F1: 154}
            v8 := []int{fi6, (fi6 & 1)}
            fmt.Println(p2, fi6)
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS2{F0: 255, F1: 954, F2: 0}
        for fi10 := 0; fi10 < 13; fi10++ {
            p1 = int8(p2)
            p1 = p1
            _ = fi10
        }
        _ = v9
    } else {
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    var v1 uint32 = 865
    var v2 float32 = 0.0
    var v3 bool = ((!(-62 > 1)) || false)
    var v11 uint8 = 255
    var v12 uint64 = 402
    var v13 uint8 = 255
    v11 = (^(v11))
    var v14 string = ("foo bar" + "hello")
    v15 := []int{-1, -9223372036854775808, -814}
    if len(v15) > 0 {
        v15[0] = -580
    }
    var v16 int64 = -1
    v17 := GenS3{F0: 812, F1: 2, F2: 42}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v18 := GenS4{F0: 242, F1: 833, F2: -1}
    if (v3 && false) {
        switch "a" {
        case "a":
            v19 := GenS5{F0: 2147483647}
            _ = v19
        case "":
            var v20 string = ""
            _ = v20
        case "\n\t":
            v16 = (v16 | (v16 << ((1) & 63)))
        }
        var v21 int32 = 1
        _ = v21
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            v1 -= v1
            _ = fi22
        }
        var v23 uint = 255
        _ = v23
    }
    v24 := GenS3{F0: (v12 >> ((279) & 63)), F1: 936, F2: (v16 ^ -739)}
    fmt.Println(v1, v3)
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(v14, fi25)
        _ = fi25
    }
    v26 := GenS6{F0: ((-9223372036854775808 < -1) && (v16 <= v16))}
    v1 = v1
    var v27 uint32 = uint32(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
}
