package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 int, p2 int16) int32 {
    var v7 int32 = 2147483647
    if false {
        var v8 uint8 = (p0 << ((873) & 63))
        for fi9 := 0; fi9 < 12; fi9++ {
            v8 -= (p0 >> ((207) & 63))
            p1 -= (fi9 / (((p1 % ((128) | 1))) | 1))
            fmt.Println(p2, fi9, p2)
            _ = fi9
        }
        _ = v8
    }
    if (100 != 0) {
        fmt.Println(p2, p1, v7)
    }
    switch p1 {
    case -9223372036854775808:
        if ((("\n\t" + "\n\t") + "") > ("" + "hello")) {
            switch p1 {
            case 767:
                p1 += (p1 >> ((1) & 63))
            }
        } else {
            switch 42 {
            case -736:
                fmt.Println(v7, p1)
            }
        }
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return int32(p2)
}
func genFn1(p0 string, p1 bool) int64 {
    p0 = (("hello" + "hello") + ("hello" + ("hello" + "hello")))
    if (p1 || (0 > 569)) {
        var v10 int64 = 42
        _ = v10
    } else {
        p1 = (!(!true))
        v11 := GenS1{F0: -593}
        _ = v11
    }
    p0 = "hello"
    v12 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v12
    return 1
}

func main() {
    var v1 uint32 = 928
    var v2 int64 = -775
    v3 := []uint32{}
    var v4 float64 = 0.0
    v5 := make(map[uint32]int)
    var v6 int64 = (v2 / (((^(v2))) | 1))
    for fi13 := 0; fi13 < 1; fi13++ {
        delete(v5, (v1 * (^(v1))))
        fmt.Println(v1)
        _ = fi13
    }
    v4 = -2.25
    v6 = 0
    switch -243 {
    case -1:
        v2 = (v6 + 42)
    }
    var v14 int64 = v6
    if true {
        switch -9223372036854775808 {
        case 535:
            v3 = append(v3, (v1 / (((v1 >> ((145) & 63))) | 1)))
        case 553:
            v15 := v5[(v1 % (((v1 & v1)) | 1))]
            _ = v15
        case 32767:
            var v16 int8 = -83
            _ = v16
        }
        v14 += v14
    }
    var v17 uint32 = (v1 | v1)
    if len(v3) > 0 {
        v3[0] = (v17 << ((171) & 63))
    }
    fmt.Println(v2)
    v2 = v14
    v18 := v5[0]
    var v19 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
