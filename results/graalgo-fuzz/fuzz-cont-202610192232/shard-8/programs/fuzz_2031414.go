package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 bool }
func genFn0(p0 int64) uint32 {
    v5 := []int{4294967295, 127, 9223372036854775807}
    p0 = p0
    v6 := GenS1{F0: 157, F1: 128}
    _ = p0
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int8) string {
    fmt.Println(p0)
    if (!false) {
        fmt.Println(p0)
        p0 = 1
    }
    fmt.Println(p0)
    p0 = p0
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            p0 = (-(p0))
            var v8 uint8 = 1
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 int16 = -422
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            if (!false) {
                if (!(!true)) {
                    p0 -= 0
                } else {
                    p0 = 42
                    var v11 bool = true
                    var v12 int16 = 352
                    _ = v11
                    _ = v12
                }
                var v13 uint = 948
                _ = v13
            } else {
                fmt.Println(p0)
                var v14 int8 = p0
                var v15 float32 = (0.0 / 0.1)
                _ = v14
                _ = v15
            }
        }
        _ = fi7
    }
    _ = p0
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    var v3 int64 = 42
    v4 := []int{2147483647, 4294967295, 515}
    var v16 int32 = 1
    var v17 uint16 = 588
    if (-1 <= 42) {
        v3 -= (v3 / (((v3 << ((65535) & 63))) | 1))
        v18 := GenS1{F0: (v17 | (v2 + 42)), F1: 127}
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int64 = (v3 << ((65535) & 63))
        v21 := v1[4294967295]
        switch (fi19 % (((v21 + (fi19 >> ((4294967295) & 63)))) | 1)) {
        case 32768:
            v22 := GenS1{F0: 2, F1: 2}
            _ = v22
        case -1:
            fmt.Println(v17, v3, v2)
        default:
            v17 += (v17 | 65535)
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v23 int64 = (v3 | v3)
    var v24 int64 = 87
    if true {
        fmt.Println(v3, v17, v24)
        for fi25 := 0; fi25 < 8; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 bool = true
                _ = fi26
                _ = v27
            }
            switch (fi25 | (fi25 & fi25)) {
            case -9223372036854775808:
                v3 -= -9223372036854775808
            default:
                switch "foo bar" {
                case "a":
                    fmt.Println(v17)
                case "foo bar":
                    if (!(!(0 > 9223372036854775807))) {
                        fmt.Println(fi25, v17)
                        v28 := GenS1{F0: v17, F1: 203}
                        _ = v28
                    } else {
                        fmt.Println(v16, v2, v23)
                    }
                case "hello":
                    for fi29 := 0; fi29 < 13; fi29++ {
                        var v30 int32 = v16
                        switch fi25 {
                        case 0:
                            v2 += (^(v17))
                        case 221:
                            var v31 int64 = v3
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                }
            }
            _ = fi25
        }
    } else {
        var v32 float64 = float64((3.14159 - (0.0 - 1e10)))
        v33 := GenS2{F0: 0}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        fmt.Println(v3, fi34, v24)
        var v35 string = ""
        v16 += -1
        _ = fi34
        _ = v35
    }
    v36 := GenS3{F0: 65535, F1: (true && (false && false))}
    delete(v1, genFn0(v24))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v4
}
