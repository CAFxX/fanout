package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 string) int {
    var v5 int16 = 32767
    v6 := []int{}
    _ = p0
    _ = v5
    _ = v6
    return 603
}
func genFn1(p0 uint32) bool {
    v7 := GenS1{F0: 255, F1: 0, F2: (p0 ^ p0)}
    p0 += p0
    var v8 uint8 = 35
    switch "a" {
    case "foo bar":
        v9 := GenS1{F0: v8, F1: 783, F2: (p0 ^ p0)}
        _ = v9
    case "":
        v10 := GenS2{F0: 128, F1: 1}
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return ((-77 > 1) && true)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    var v4 uint16 = 32767
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            v12 := GenS3{F0: 428}
            _ = v12
        }
        if true {
            v2 = "foo bar"
            v13 := []int{genFn0((v2 + (("\n\t" + "\n\t") + v2))), 32768, 9223372036854775807}
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 string = "foo bar"
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            fmt.Println(fi11)
            switch genFn0(("hello" + "\n\t")) {
            case 9223372036854775807:
                if ((!false) && (true && true)) {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v17 := GenS4{F0: 42, F1: v4}
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS3{F0: 9223372036854775807}
                    _ = v18
                } else {
                    v19 := v1[(fi11 * fi11)]
                    _ = v19
                }
            default:
                var v20 float32 = 3.14159
                _ = v20
            }
        }
        var v21 uint32 = 2147483647
        _ = fi11
        _ = v21
    }
    var v22 uint = 2
    v22 = v22
    v23 := GenS1{F0: 180, F1: v22, F2: 2147483647}
    if false {
        if true {
            fmt.Println(v2, v22, v2)
            var v24 string = v2
            v24 = (v2 + (("\n\t" + v2) + v24))
            _ = v24
        } else {
            v25 := GenS4{F0: 889, F1: 2}
            for fi26 := 0; fi26 < 8; fi26++ {
                v27 := GenS3{F0: 309}
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        v4 = 967
    }
    var v28 int8 = 109
    v22 = 788
    fmt.Println(v4, v22, v28)
    for fi29 := 0; fi29 < 9; fi29++ {
        v30 := v3[fi29]
        _ = fi29
        _ = v30
    }
    if true {
        if (-9223372036854775808 < -9223372036854775808) {
            v31 := GenS1{F0: 168, F1: (v22 - (^(v22))), F2: 389}
            v32 := []int{0}
            _ = v31
            _ = v32
        }
    } else {
        for fi33 := 0; fi33 < 8; fi33++ {
            if (((((("\n\t" + (v2 + "\n\t")) + v2) + "\n\t") + v2) + v2) < "\n\t") {
                var v34 uint16 = 706
                v22 = 255
                _ = v34
            } else {
                for fi35 := 0; fi35 < 4; fi35++ {
                    v36 := []int{837, fi33}
                    v22 += 1
                    _ = fi35
                    _ = v36
                }
            }
            var v37 uint = (v22 | v22)
            var v38 int = (fi33 << ((v22) & 63))
            _ = fi33
            _ = v37
            _ = v38
        }
    }
    v39 := []int{-229}
    for fi40 := 0; fi40 < 18; fi40++ {
        v4 = 824
        _ = fi40
    }
    if true {
        v4 += (v4 % ((v4) | 1))
        v22 += (v22 - v22)
        switch 32768 {
        case 9223372036854775807:
            delete(v1, 32768)
        default:
            for fi41 := 0; fi41 < 9; fi41++ {
                v4 = (v4 % (((v4 << (((v22 | v22)) & 63))) | 1))
                _ = fi41
            }
        }
    } else {
        v4 += v4
        var v42 string = "foo bar"
        _ = v42
    }
    v22 = (v22 + (v22 + v22))
    if (1 > 1) {
        v28 = -1
        v2 = "a"
    } else {
        v43 := GenS5{F0: v4}
        var v44 uint64 = 255
        _ = v43
        _ = v44
    }
    fmt.Println(v2)
    if ((float64(0.0) > -2.25) && true) {
        v28 -= v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v39
    _ = v4
}
