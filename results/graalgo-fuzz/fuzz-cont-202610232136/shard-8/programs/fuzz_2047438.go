package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int16; F1 uint; F2 int32 }
func genFn0(p0 int32, p1 int8) bool {
    p1 = (p1 >> ((246) & 63))
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: "a"}
        var v9 uint32 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 1}
    if false {
        p1 = p1
        var v11 int32 = (p0 / (((p0 ^ (p0 * (p0 << ((993) & 63))))) | 1))
        v11 += 42
        _ = v11
    }
    p0 -= -1
    _ = p0
    _ = p1
    _ = v10
    return true
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    var v4 int16 = -540
    v5 := []string{"", ""}
    v6 := make(map[int]bool)
    v12 := GenS3{F0: 218, F1: ("a" + (("\n\t" + "foo bar") + (((("foo bar" + "hello") + "a") + "\n\t") + "\n\t")))}
    fmt.Println(v4)
    fmt.Println(v4)
    v4 = (v4 % (((v4 & 0)) | 1))
    fmt.Println(v4)
    v4 += (-(v4))
    if (0 == 2) {
        switch "\n\t" {
        case "":
            var v13 uint = 18446744073709551615
            _ = v13
        }
        var v14 int = 128
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        if true {
            var v16 uint16 = 396
            v17 := GenS1{F0: "foo bar"}
            v18 := v1[575]
            _ = v16
            _ = v17
            _ = v18
        } else {
            var v19 uint = 398
            _ = v19
        }
        v20 := GenS4{F0: -1, F1: 0, F2: 2147483647}
        _ = fi15
        _ = v20
    }
    var v21 string = "foo bar"
    switch 9223372036854775807 {
    case 32767:
        for fi22 := 0; fi22 < 6; fi22++ {
            v21 = (v21 + "a")
            _ = fi22
        }
    default:
        for fi23 := 0; fi23 < 4; fi23++ {
            v4 += int16(fi23)
            _ = fi23
        }
    }
    v6[9223372036854775807] = true
    v21 = v21
    _ = len(v2)
    var v24 uint16 = 0
    var v25 int16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
