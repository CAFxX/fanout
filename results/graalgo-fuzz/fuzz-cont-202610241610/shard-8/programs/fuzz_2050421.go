package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16; F1 uint; F2 uint }
type GenS3 struct { F0 uint32; F1 int64; F2 uint }
type GenS4 struct { F0 uint; F1 int8; F2 int }
func genFn0(p0 uint8, p1 int16, p2 int32) int8 {
    var v7 uint16 = 65535
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint32 = 474
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := []int{fi8}
            switch ("" + "foo bar") {
            case "\n\t":
                var v12 uint32 = v9
                _ = v12
            case "foo bar":
                p1 -= p1
            case "":
                if (!false) {
                    var v13 uint = 255
                    fmt.Println(p2, v9)
                    v13 = v13
                    _ = v13
                }
            }
            var v14 float32 = -2.25
            _ = fi10
            _ = v11
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int8 = -1
    var v2 float64 = 1e10
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    v5 := []uint32{2147483647, 2147483647, 2147483647, 1}
    var v6 uint64 = 822
    var v15 uint32 = 42
    var v16 float64 = (v2 / 1e10)
    var v17 uint16 = 780
    delete(v4, "foo bar")
    switch int(v1) {
    case -771:
        var v18 uint32 = 2147483648
        _ = v18
    case -1:
        var v19 uint8 = 255
        _ = v19
    case -548:
        fmt.Println(v6, v16, v1)
    default:
        v16 = 3.14159
    }
    fmt.Println(v16)
    v20 := GenS1{F0: -686, F1: false}
    if ((v2 / float64(float32(v2))) > v2) {
        for fi21 := 0; fi21 < 6; fi21++ {
            v4["\n\t"] = v15
            switch 32767 {
            case -116:
                v6 = v6
            case 0:
                _ = len(v5)
            case 42:
                var v22 int8 = (v1 >> ((2) & 63))
                _ = v22
            default:
                var v23 float32 = 1e10
                _ = v23
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                fmt.Println(v15)
                switch (fi24 & -9223372036854775808) {
                case 32767:
                    var v25 int = 549
                    switch fi21 {
                    case -9223372036854775808:
                        v1 = -128
                    case 9223372036854775807:
                        if true {
                            v15 = v15
                        } else {
                            var v26 bool = (42 == (fi21 ^ (fi24 / (((fi24 & 246)) | 1))))
                            _ = v26
                        }
                    }
                    v5 = append(v5, v15)
                    v27 := GenS2{F0: -1, F1: 0, F2: 1}
                    delete(v3, "a")
                    v28 := GenS3{F0: 4294967295, F1: -46, F2: 196}
                    _ = len(v5)
                    switch fi24 {
                    case 9223372036854775807:
                        var v29 uint = 904
                        _ = v29
                    }
                    v6 -= v6
                    v30 := v3["\n\t"]
                    _ = v25
                    _ = v27
                    _ = v28
                    _ = v30
                default:
                    _ = len(v4)
                }
                _ = fi24
            }
            _ = fi21
        }
    } else {
        v2 -= (v2 * v16)
        v16 = (v2 * v16)
        v1 -= 1
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        switch "a" {
        case "\n\t":
            v32 := GenS4{F0: 255, F1: (-(v1)), F2: int(v15)}
            _ = v32
        case "hello":
            v2 += 3.14159
        default:
            if (!true) {
                v33 := GenS2{F0: 163, F1: 918, F2: 2}
                v34 := GenS2{F0: -912, F1: 742, F2: 42}
                var v35 uint32 = (^(v15))
                _ = v33
                _ = v34
                _ = v35
            }
        }
        _ = fi31
    }
    if (v15 != (v15 ^ v15)) {
        v15 = v15
        var v36 float64 = v16
        _ = v36
    } else {
        fmt.Println(v6)
    }
    v16 -= v16
    v15 = v15
    v37 := GenS1{F0: -2147483648, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
