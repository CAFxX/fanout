package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() int32 {
    var v4 int = 527
    if true {
        v5 := GenS1{F0: 42}
        switch v4 {
        case 1:
            v6 := []int{v4}
            _ = v6
        }
        fmt.Println(v4)
        _ = v5
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            if (1000 <= 1000) {
                if false {
                    fmt.Println(v4, v4)
                } else {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        v4 = 301
                        var v9 uint64 = 0
                        v10 := GenS1{F0: 32767}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    v4 += (v4 << ((0) & 63))
                    v11 := GenS1{F0: -1}
                    _ = v11
                }
                if (32 != 2) {
                    var v12 bool = (!(!false))
                    if (!(false && (!(v12 && (-1 == -1))))) {
                        fmt.Println(fi7)
                        switch (-(fi7)) {
                        case -9223372036854775808:
                            var v13 int16 = -32768
                            _ = v13
                        case -255:
                            v4 += v4
                        default:
                            v4 += v4
                        }
                    } else {
                        v14 := []int{128}
                        v4 += (fi7 - (v4 + fi7))
                        _ = v14
                    }
                    v12 = v12
                    _ = v12
                }
            }
            v4 = (v4 | (fi7 - 1))
            v4 -= (v4 - v4)
            _ = fi7
        }
    }
    var v15 int8 = 24
    fmt.Println(v4, v4)
    _ = v15
    _ = v4
    return -305
}
func genFn1(p0 uint16, p1 int8) uint64 {
    fmt.Println(p0, p0)
    var v16 uint = 4294967295
    p1 -= (p1 & p1)
    v16 += (v16 << ((199) & 63))
    switch (("foo bar" + "hello") + "foo bar") {
    case "hello":
        v17 := []int{}
        _ = v17
    case "":
        var v18 int8 = 0
        _ = v18
    case "\n\t":
        v19 := GenS1{F0: 1}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v16
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 42
    var v3 uint16 = 255
    if true {
        for fi20 := 0; fi20 < 1; fi20++ {
            v2 += v2
            switch fi20 {
            case 127:
                v21 := GenS2{F0: 290, F1: 458}
                _ = v21
            default:
                for fi22 := 0; fi22 < 1; fi22++ {
                    fmt.Println(fi22)
                    fmt.Println(v3, v2)
                    _ = fi22
                }
            }
            _ = fi20
        }
        v3 += (v3 << ((4294967295) & 63))
    }
    if false {
        v2 = v2
    } else {
        var v23 int16 = v2
        var v24 uint = 4294967295
        var v25 int = 32767
        _ = v23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := GenS3{F0: true}
        v3 = (v3 % ((v3) | 1))
        v28 := v1["foo bar"]
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v3, v2)
    var v29 float64 = 3.14159
    for fi30 := 0; fi30 < 16; fi30++ {
        if false {
            v2 -= -433
            var v31 uint32 = 2147483647
            v32 := GenS1{F0: v2}
            _ = v31
            _ = v32
        }
        fmt.Println(v29, fi30)
        v29 = v29
        _ = fi30
    }
    v3 -= v3
    v33 := GenS2{F0: genFn0(), F1: 340}
    switch "hello" {
    case "foo bar":
        v34 := []int{924}
        _ = v34
    case "hello":
        v35 := GenS4{F0: "a", F1: 0}
        _ = v35
    case "\n\t":
        v3 += 743
    default:
        v2 += v2
    }
    if (v3 <= (^(v3))) {
        fmt.Println(v29, v3, v3)
        fmt.Println(v29, v3)
    } else {
        var v36 bool = (!false)
        _ = v36
    }
    v29 = v29
    for fi37 := 0; fi37 < 4; fi37++ {
        v29 = v29
        fmt.Println(v29, v3, v29)
        if (0 > 42) {
            fmt.Println(v29)
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
}
