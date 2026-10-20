package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int64; F1 int8; F2 uint32 }
func genFn0(p0 int8, p1 uint32, p2 int64) uint {
    switch 0 {
    case 128:
        v4 := GenS1{F0: 2, F1: 0}
        _ = v4
    case 42:
        p1 += (^(p1))
    case 666:
        p0 += p0
    }
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int = (-(fi5))
        var v7 bool = false
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 0 {
    case 4294967295:
        v8 := []int{-279, 666}
        _ = v8
    }
    p1 -= (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1(p0 int8, p1 int64) bool {
    v9 := []int{32768, 170, 1}
    v10 := GenS1{F0: 0, F1: 32767}
    var v11 uint64 = 931
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return (1e10 >= (3.14159 - (1.5 / 1e10)))
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    var v3 float32 = float32(float64((1.5 + float32(0.0))))
    var v12 int32 = v2
    v12 = 2147483647
    var v13 int = 128
    v3 -= v3
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint32 = 1
        for fi16 := 0; fi16 < 8; fi16++ {
            if (0 == 42) {
                v17 := GenS1{F0: 0, F1: 0}
                v18 := []int{fi14, 4294967295, -9223372036854775808}
                _ = v17
                _ = v18
            } else {
                var v19 int64 = -9223372036854775808
                fmt.Println(v12, fi16)
                v20 := []int{}
                switch "" {
                case "a":
                    switch (fi16 | (fi14 ^ (fi16 / ((fi14) | 1)))) {
                    case 32768:
                        var v21 uint = genFn0(47, v15, -1)
                        _ = v21
                    case 36:
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    default:
                        switch fi16 {
                        case 128:
                            if true {
                                var v22 uint = genFn0(127, (v15 | (v15 / (((v15 * v15)) | 1))), (v19 | v19))
                                switch "" {
                                case "\n\t":
                                    v1[(v13 / ((fi14) | 1))] = true
                                case "foo bar":
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v23 uint64 = 325
                                        _ = v23
                                    default:
                                        var v24 uint32 = (v15 ^ (v15 * (v15 * v15)))
                                        _ = v24
                                    }
                                }
                                _ = v22
                            } else {
                                fmt.Println(v12)
                                fmt.Println(v19)
                            }
                        }
                    }
                }
                _ = v19
                _ = v20
            }
            _ = fi16
        }
        if true {
            v12 = v12
            switch (^(v13)) {
            case -123:
                var v25 uint = 4294967295
                _ = v25
            case 0:
                v26 := []int{(v13 * 9223372036854775807), fi14}
                _ = v26
            case 127:
                v15 = 4294967295
            }
        } else {
            fmt.Println(fi14, v3, v15)
        }
        _ = fi14
        _ = v15
    }
    var v27 uint16 = 2
    v3 = v3
    v12 = -822
    v28 := GenS2{F0: ("hello" + "\n\t"), F1: v13}
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            switch "a" {
            case "\n\t":
                var v31 int8 = 1
                _ = v31
            }
            v32 := GenS3{F0: -835, F1: -128, F2: 65535}
            _ = fi30
            _ = v32
        }
        _ = fi29
    }
    if (160 == 9223372036854775807) {
        var v33 int64 = 9223372036854775807
        var v34 int64 = (v33 + 0)
        v35 := v1[v13]
        _ = v33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        v37 := []int{1, v13, v13, (-(v13))}
        v3 = v3
        _ = fi36
        _ = v37
    }
    if ((false || false) || true) {
        v2 += (v12 | v12)
        v3 -= (v3 * v3)
    } else {
        v38 := GenS1{F0: 1, F1: v27}
        _ = v38
    }
    v12 = (v12 << ((4294967295) & 63))
    fmt.Println(v27, v3)
    var v39 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v39
}
