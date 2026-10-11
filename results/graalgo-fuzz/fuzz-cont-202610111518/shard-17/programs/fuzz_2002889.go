package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 int }

func main() {
    var v1 int = -932
    var v2 int64 = 1
    v3 := make(map[uint32]uint32)
    var v4 int64 = v2
    var v5 uint8 = 42
    v6 := []uint32{680, 583, 2}
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            if false {
                fmt.Println(v2, v5)
                var v9 int32 = 400
                v9 = v9
                _ = v9
            } else {
                v1 -= fi7
                var v10 uint = 1
                v5 = v5
                _ = v10
            }
            _ = fi8
        }
        _ = fi7
    }
    var v11 int16 = -355
    var v12 int = v1
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v2, fi13, v1)
            v6 = append(v6, 1)
            _ = fi14
        }
        v11 -= (v11 | v11)
        _ = fi13
    }
    var v15 int16 = -43
    if (-1 > -128) {
        switch v12 {
        case 4294967295:
            var v16 int8 = 122
            _ = v16
        case 238:
            if ((1 != int64(v11)) || true) {
                switch v12 {
                case 4294967295:
                    var v17 uint16 = 42
                    _ = v17
                default:
                    var v18 int64 = v2
                    _ = v18
                }
                v12 = (v1 & 1)
            }
        default:
            v1 -= (v12 >> ((65535) & 63))
        }
        fmt.Println(v5, v2)
        var v19 int8 = 100
        _ = v19
    } else {
        v11 -= 159
        if true {
            v3[65535] = 1
        } else {
            fmt.Println(v4)
            v15 -= 1
        }
        if (1 == 470) {
            v20 := GenS1{F0: false}
            _ = v20
        } else {
            v15 += -1
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        fmt.Println(v15, v1, v15)
        _ = fi21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS1{F0: (!true)}
            fmt.Println(v15)
            _ = fi23
            _ = v24
        }
        v25 := GenS2{F0: "", F1: (!((0 == 4294967295) || (0.1 != 1.5)))}
        _ = fi22
        _ = v25
    }
    v26 := GenS2{F0: (("\n\t" + "hello") + "foo bar"), F1: true}
    var v27 uint = 255
    v3[746] = 2
    var v28 int16 = v15
    if (int(v5) < (v1 << ((v27) & 63))) {
        v29 := GenS3{F0: 264, F1: 107, F2: (v12 << (((v27 | (v27 | v27))) & 63))}
        for fi30 := 0; fi30 < 6; fi30++ {
            v31 := GenS3{F0: 979, F1: uint32(v27), F2: (v12 << (((v27 | (v27 >> ((v27) & 63)))) & 63))}
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 uint16 = 42
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 string = ("" + "")
                var v36 int16 = v15
                _ = fi34
                _ = v35
                _ = v36
            }
            v37 := GenS1{F0: ("foo bar" == (("foo bar" + "") + "hello"))}
            _ = fi32
            _ = v33
            _ = v37
        }
        _ = v29
    } else {
        v38 := GenS3{F0: 659, F1: 1, F2: (v12 | 2147483647)}
        var v39 int32 = -453
        _ = v38
        _ = v39
    }
    v4 = 247
    v40 := GenS2{F0: "\n\t", F1: (127 >= 42)}
    var v41 int8 = 87
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
