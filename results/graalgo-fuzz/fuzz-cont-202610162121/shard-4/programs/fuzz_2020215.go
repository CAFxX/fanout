package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -1
    v3 := []uint32{4294967295, 0, 287}
    var v4 float64 = 1e10
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        switch fi5 {
        case 32767:
            fmt.Println(v2, fi5, fi5)
        }
        _ = fi5
    }
    if (false || false) {
        if (!false) {
            v2 = (v2 << ((65535) & 63))
            v4 -= v4
        }
        v2 -= v2
    } else {
        v2 -= v2
        var v6 int = -9223372036854775808
        _ = v6
    }
    var v7 string = ((("\n\t" + "") + "a") + "")
    switch 32768 {
    case 0:
        switch -819 {
        case 380:
            v8 := GenS1{F0: 424}
            _ = v8
        case 0:
            v2 += (v2 & -196)
        }
    case -972:
        if (!true) {
            var v9 int64 = (v2 ^ (v2 - (v2 << ((42) & 63))))
            v9 = (v2 ^ v2)
            if false {
                if (!false) {
                    v7 = "hello"
                    var v10 int = 32768
                    v11 := GenS2{F0: 1, F1: 42}
                    _ = v10
                    _ = v11
                } else {
                    var v12 int8 = -128
                    _ = v12
                }
            }
            _ = v9
        }
    default:
        if true {
            v3 = append(v3, 782)
            if false {
                delete(v1, 42)
                if true {
                    var v13 string = (("hello" + "hello") + (v7 + v7))
                    _ = v13
                }
            } else {
                v2 -= -555
                v14 := GenS2{F0: 814, F1: 888}
                _ = v14
            }
        }
    }
    fmt.Println(v4)
    v7 = ""
    var v15 bool = (true || (255 == 1))
    v2 -= v2
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v4, fi16, v2)
        v15 = (762 > 65535)
        v2 = 42
        _ = fi16
    }
    switch 4294967295 {
    case 32768:
        v15 = (true && (1 == 4294967295))
    case -936:
        if v15 {
            v7 = "hello"
            fmt.Println(v2, v2)
            v4 -= 0.1
        }
    case -423:
        var v17 uint32 = 65535
        _ = v17
    }
    var v18 uint8 = 127
    var v19 int8 = -1
    var v20 int32 = 0
    v2 += v2
    if true {
        var v21 uint64 = 0
        if true {
            _ = len(v3)
            var v22 int64 = (-(v2))
            _ = v22
        }
        switch (v7 + v7) {
        case "\n\t":
            v23 := GenS1{F0: (v21 % (((v21 >> ((0) & 63))) | 1))}
            _ = v23
        case "a":
            var v24 uint64 = v21
            _ = v24
        case "":
            if len(v3) > 0 {
                v3[0] = 65535
            }
        }
        _ = v21
    }
    v25 := v1[127]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
