package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 int16, p1 uint8, p2 uint8) uint32 {
    v5 := GenS1{F0: 65535, F1: 2147483648, F2: 2}
    var v6 int8 = 1
    var v7 int32 = int32(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16) uint16 {
    switch 4294967295 {
    case -2147483648:
        if false {
            var v8 uint8 = 2
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint = 42
            _ = fi9
            _ = v10
        }
    }
    if (9223372036854775807 == -1) {
        v11 := GenS1{F0: 18446744073709551615, F1: 603, F2: 2147483647}
        v12 := GenS2{F0: 12}
        v13 := []int{1}
        _ = v11
        _ = v12
        _ = v13
    }
    p0 += p0
    var v14 int8 = -128
    _ = p0
    _ = v14
    return 587
}

func main() {
    v1 := []int{32767, 2147483647, 0}
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    v4 := []int{}
    if true {
        _ = len(v2)
        for fi15 := 0; fi15 < 12; fi15++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi15
        }
    } else {
        var v16 uint8 = 127
        _ = v16
    }
    _ = len(v4)
    var v17 int64 = 288
    var v18 float32 = -2.25
    v18 = v18
    var v19 string = "hello"
    _ = len(v1)
    switch "" {
    case "hello":
        var v20 uint32 = 2147483647
        _ = v20
    case "foo bar":
        if true {
            if true {
                _ = len(v2)
                fmt.Println(v18)
            }
        }
    case "":
        for fi21 := 0; fi21 < 4; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                v19 = (v19 + "")
                v19 = v19
                fmt.Println(fi21, v17)
                _ = fi22
            }
            v23 := GenS2{F0: 2}
            _ = fi21
            _ = v23
        }
    }
    if (false && (!true)) {
        var v24 uint = 725
        _ = v24
    } else {
        var v25 uint8 = 0
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 int64 = 0
            var v28 int32 = -294
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    var v29 uint = 18446744073709551615
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS3{F0: 0, F1: 462}
        var v32 uint = (v29 >> ((v29) & 63))
        var v33 int8 = 42
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    var v34 uint = (v29 + v29)
    if true {
        if false {
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 float64 = float64(v18)
                _ = fi35
                _ = v36
            }
            _ = len(v3)
            for fi37 := 0; fi37 < 8; fi37++ {
                v2[(-(fi37))] = uint32(v34)
                delete(v2, (-(fi37)))
                v34 += v29
                _ = fi37
            }
        } else {
            var v38 int64 = v17
            _ = v38
        }
        switch -2147483648 {
        case 647:
            v39 := GenS3{F0: -1, F1: -541}
            _ = v39
        case -2147483648:
            switch -9223372036854775808 {
            case -26:
                v40 := GenS1{F0: 4294967295, F1: 302, F2: 2}
                _ = v40
            default:
                fmt.Println(v34, v29, v29)
            }
        }
    }
    v41 := GenS1{F0: (v29 >> (((v29 | (v34 & (v34 + v34)))) & 63)), F1: 2147483648, F2: 2}
    var v42 int64 = v17
    var v43 int32 = 377
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
