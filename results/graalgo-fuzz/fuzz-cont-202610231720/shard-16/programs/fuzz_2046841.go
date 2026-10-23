package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 string }
func genFn0(p0 uint64, p1 string, p2 uint8) int8 {
    var v5 int8 = 0
    if (true && ((65535 < 194) && false)) {
        v6 := []int{}
        v7 := GenS1{F0: (p0 / ((p0) | 1)), F1: (1.5 != 0.0), F2: 799}
        if false {
            v8 := GenS1{F0: (p0 / (((p0 + p0)) | 1)), F1: (p1 < p1), F2: uint(p0)}
            _ = v8
        }
        _ = v6
        _ = v7
    }
    var v9 int = 238
    v10 := GenS1{F0: p0, F1: false, F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return int8(p2)
}
func genFn1(p0 uint64, p1 int32) int {
    if (4294967295 == 241) {
        v11 := GenS2{F0: genFn0((^(p0)), "foo bar", 1)}
        var v12 uint8 = 142
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            p0 = p0
            var v14 int = fi13
            v14 = 393
            _ = fi13
            _ = v14
        }
        p1 += p1
    }
    p0 -= 4294967295
    var v15 uint32 = 2147483647
    var v16 float64 = -2.25
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return int(p0)
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{4294967295}
    var v3 uint64 = 1
    v4 := make(map[int]uint32)
    var v17 int64 = 0
    v17 = 926
    v3 = (^(v3))
    var v18 uint16 = 779
    var v19 bool = (255 <= 0)
    v3 = (v3 << ((1) & 63))
    switch (("a" + "a") + "foo bar") {
    case "a":
        var v20 int16 = 0
        _ = v20
    default:
        v21 := GenS1{F0: 42, F1: false, F2: 395}
        _ = v21
    }
    switch -2147483648 {
    case 32768:
        v22 := GenS1{F0: (v3 | 211), F1: (!v19), F2: 1}
        _ = v22
    case 4294967295:
        if (0.1 > 0.0) {
            var v23 int64 = v17
            _ = v23
        } else {
            if (v17 > (v17 | (v17 >> ((2) & 63)))) {
                v24 := GenS1{F0: v3, F1: (v3 < 42), F2: 551}
                _ = v24
            } else {
                v18 -= (v18 / (((v18 + (v18 * v18))) | 1))
                fmt.Println(v17, v3)
            }
            v2 = append(v2, 42)
            v3 = (v3 - v3)
        }
    case 42:
        var v25 uint16 = (v18 << ((65535) & 63))
        _ = v25
    default:
        fmt.Println(v19, v19)
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v27 float32 = 0.1
        _ = fi26
        _ = v27
    }
    v18 -= 7
    var v28 bool = (730 < 0)
    var v29 float64 = 1.5
    v29 = 1.5
    v19 = true
    switch 127 {
    case 32767:
        for fi30 := 0; fi30 < 3; fi30++ {
            v31 := GenS3{F0: -297, F1: v28, F2: "a"}
            var v32 uint8 = 128
            v19 = ((fi30 - -894) <= fi30)
            _ = fi30
            _ = v31
            _ = v32
        }
    case -642:
        switch "" {
        case "hello":
            if v28 {
                var v33 float32 = 0.0
                _ = v33
            }
        case "":
            for fi34 := 0; fi34 < 2; fi34++ {
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(fi34, v17, v17)
                }
                fmt.Println(v28, v17)
                _ = fi34
            }
        case "\n\t":
            for fi35 := 0; fi35 < 4; fi35++ {
                for fi36 := 0; fi36 < 5; fi36++ {
                    var v37 string = (("a" + "a") + "")
                    if v19 {
                        fmt.Println(v19, v29, v17)
                        v17 = (v17 + v17)
                        var v38 uint = 42
                        _ = v38
                    }
                    var v39 uint = 394
                    _ = fi36
                    _ = v37
                    _ = v39
                }
                _ = fi35
            }
        }
    case 181:
        var v40 uint = 2
        _ = v40
    }
    v41 := GenS1{F0: 4294967295, F1: (v29 <= ((0.1 * (v29 / float64(3.14159))) - 3.14159)), F2: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
}
