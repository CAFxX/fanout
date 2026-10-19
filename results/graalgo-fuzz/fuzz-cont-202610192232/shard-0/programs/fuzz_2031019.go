package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int64; F2 bool }
func genFn0(p0 int32, p1 int64) uint32 {
    p1 += p1
    v4 := []int{-9223372036854775808, 2147483647, 4294967295, -154}
    fmt.Println(p1)
    if len(v4) > 0 {
        v4[0] = -1
    }
    if (255 <= 65535) {
        var v5 int = 9223372036854775807
        for fi6 := 0; fi6 < 7; fi6++ {
            if len(v4) > 0 {
                v4[0] = (v5 << ((454) & 63))
            }
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS1{F0: ((3.14159 <= float64(1.5)) || true)}
                var v9 int64 = (^(p1))
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int16 = 42
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch -616 {
                case 883:
                    v12 := GenS2{F0: 42}
                    _ = v12
                case 44:
                    v11 = v11
                case -1:
                    p1 = p1
                }
                _ = fi10
                _ = v11
            }
            _ = fi6
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 658
}

func main() {
    var v1 float64 = (0.0 * 0.0)
    v2 := []uint32{2, 2}
    v3 := make(map[uint32]bool)
    v1 += v1
    fmt.Println(v1)
    fmt.Println(v1)
    var v13 uint32 = 42
    v14 := GenS2{F0: 1}
    var v15 uint8 = uint8(v13)
    if (2 != 32767) {
        v1 += v1
        var v16 int32 = 798
        v17 := v3[v13]
        _ = v16
        _ = v17
    } else {
        var v18 int = 170
        var v19 string = ("foo bar" + (((("a" + "\n\t") + "a") + "foo bar") + ""))
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            if true {
                fmt.Println(v13, v13, v1)
                switch fi20 {
                case -9223372036854775808:
                    _ = len(v2)
                case -1:
                    v22 := GenS1{F0: true}
                    _ = v22
                }
            }
            v15 = v15
            _ = fi21
        }
        if ("\n\t" < "") {
            switch (^(fi20)) {
            case -547:
                switch (fi20 << ((1) & 63)) {
                case 2147483647:
                    switch "hello" {
                    case "hello":
                        delete(v3, (v13 / ((uint32(v15)) | 1)))
                    case "":
                        _ = len(v3)
                    }
                case 9223372036854775807:
                    for fi23 := 0; fi23 < 3; fi23++ {
                        var v24 int16 = 1
                        fmt.Println(v15, fi23)
                        v25 := GenS3{F0: 966, F1: -1, F2: false}
                        delete(v3, v13)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case -2147483648:
                    v26 := GenS2{F0: 955}
                    _ = v26
                default:
                    var v27 float64 = (v1 * ((-2.25 * (v1 * v1)) / 0.1))
                    _ = v27
                }
            case 0:
                var v28 float32 = 0.1
                _ = v28
            }
        }
        _ = len(v3)
        _ = fi20
    }
    v29 := v3[v13]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
}
