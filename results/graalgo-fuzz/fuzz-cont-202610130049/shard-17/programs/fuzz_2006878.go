package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }
func genFn0(p0 int32) int16 {
    v4 := GenS1{F0: 4294967295, F1: 2}
    if true {
        var v5 int16 = 32767
        for fi6 := 0; fi6 < 10; fi6++ {
            p0 += (p0 >> ((18446744073709551615) & 63))
            v7 := GenS1{F0: 1, F1: 4294967295}
            for fi8 := 0; fi8 < 13; fi8++ {
                v5 += (v5 / ((v5) | 1))
                var v9 uint = 42
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: 760, F1: 0}
            _ = fi10
            _ = v11
        }
        var v12 uint8 = 127
        _ = v12
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 uint, p1 int) uint {
    if true {
        p1 = 32767
    } else {
        var v13 uint32 = 2147483648
        for fi14 := 0; fi14 < 2; fi14++ {
            switch (fi14 << (((p0 + p0)) & 63)) {
            case 871:
                var v15 float64 = -2.25
                _ = v15
            case 128:
                v16 := []int{(p1 * 127)}
                _ = v16
            default:
                v17 := []int{(p1 * (^(p1))), p1}
                _ = v17
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                if true {
                    v19 := GenS1{F0: 2, F1: (v13 | (^(v13)))}
                    var v20 uint32 = (^(v13))
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(fi18, fi14)
                    if (0.0 <= 1.5) {
                        v13 = v13
                        v13 -= v13
                        v21 := []int{fi18, p1}
                        var v22 int8 = 12
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        switch -821 {
                        case -1:
                            var v23 int32 = int32(fi18)
                            _ = v23
                        default:
                            _ = len(v21)
                        }
                        fmt.Println(fi18, v13, p1)
                        _ = v21
                        _ = v22
                    }
                }
                if ((100 > 0) || true) {
                    p1 = (fi18 << ((33) & 63))
                    v24 := []int{1, 543}
                    var v25 uint32 = v13
                    fmt.Println(fi18, fi18)
                    fmt.Println(v25, fi14, fi18)
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint16 = 255
                    var v27 uint8 = 101
                    var v28 int32 = 2147483647
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi18
            }
            switch "\n\t" {
            case "hello":
                for fi29 := 0; fi29 < 20; fi29++ {
                    v13 -= (v13 / (((v13 ^ 0)) | 1))
                    _ = fi29
                }
            case "":
                fmt.Println(v13, v13)
            case "a":
                for fi30 := 0; fi30 < 12; fi30++ {
                    switch (p1 / (((p1 << ((1) & 63))) | 1)) {
                    case 32768:
                        v31 := []int{fi14, fi30}
                        _ = v31
                    case 42:
                        var v32 uint64 = 0
                        _ = v32
                    }
                    v13 = 2
                    _ = fi30
                }
            }
            _ = fi14
        }
        _ = v13
    }
    switch p1 {
    case 0:
        p1 = 341
    case -180:
        var v33 float64 = (1e10 * (float64(3.14159) - -2.25))
        _ = v33
    case -169:
        var v34 int = p1
        _ = v34
    }
    _ = p0
    _ = p1
    return (p0 * (p0 ^ (p0 | p0)))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 uint32 = 2147483647
    if (32768 != 452) {
        _ = len(v1)
    } else {
        var v35 uint8 = 255
        _ = v35
    }
    v36 := []int{4294967295, 0, 273}
    _ = len(v1)
    var v37 bool = (("" + "a") != "\n\t")
    var v38 int16 = genFn0(0)
    var v39 uint64 = 0
    for fi40 := 0; fi40 < 15; fi40++ {
        v37 = (v37 && (1e10 < 1e10))
        _ = fi40
    }
    v36 = append(v36, 127)
    v41 := GenS2{F0: 32767, F1: 42, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v41
}
