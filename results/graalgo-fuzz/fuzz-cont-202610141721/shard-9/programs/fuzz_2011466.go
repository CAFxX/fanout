package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint }
func genFn0(p0 uint, p1 int64, p2 uint32) uint8 {
    var v4 int16 = -367
    if false {
        switch ("foo bar" + ("foo bar" + ("hello" + "a"))) {
        case "":
            var v5 uint8 = 1
            _ = v5
        case "a":
            p1 = (p1 >> (((p0 >> (((p0 * p0)) & 63))) & 63))
        case "\n\t":
            var v6 int = 641
            _ = v6
        default:
            var v7 uint8 = 42
            _ = v7
        }
    } else {
        var v8 int32 = -462
        p1 += (p1 + (p1 % (((p1 | -900)) | 1)))
        _ = v8
    }
    if true {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{fi9, (fi9 + 965), int(v4)}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: -957, F1: p2}
        _ = v11
    } else {
        v12 := GenS1{F0: 42, F1: (^(p2))}
        var v13 int = 0
        _ = v12
        _ = v13
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 46
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]uint32)
    var v3 int16 = 464
    var v14 uint8 = 0
    if (!false) {
        switch "" {
        case "a":
            if (true || false) {
                fmt.Println(v1, v3, v3)
                v3 = v3
                var v15 int32 = -1
                _ = v15
            }
        }
        if ((v1 + (^(v1))) != (v1 / ((v1) | 1))) {
            var v16 uint8 = (v14 / (((v14 * v14)) | 1))
            v17 := []int{452, 128, 32768, 42}
            fmt.Println(v3, v14, v1)
            _ = v16
            _ = v17
        } else {
            if true {
                if false {
                    v18 := []int{-959, 42}
                    _ = v18
                } else {
                    v19 := GenS2{F0: 840, F1: 127, F2: 112}
                    fmt.Println(v1, v3, v1)
                    v20 := []int{1}
                    _ = v19
                    _ = v20
                }
                v21 := GenS3{F0: 4294967295}
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (1e10 > 1e10) {
                        if true {
                            for fi23 := 0; fi23 < 8; fi23++ {
                                v1 -= (v1 / (((v1 & (v1 << ((42) & 63)))) | 1))
                                _ = fi23
                            }
                        } else {
                            fmt.Println(v1)
                        }
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v25 := GenS4{F0: 1, F1: 1, F2: v1}
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v26 := []int{fi22}
                            _ = v26
                        default:
                            v3 -= (v3 / ((v3) | 1))
                        }
                    }
                    v1 -= (^(v1))
                    v14 = v14
                    _ = fi22
                }
                _ = v21
            } else {
                for fi27 := 0; fi27 < 10; fi27++ {
                    switch "hello" {
                    case "a":
                        for fi28 := 0; fi28 < 17; fi28++ {
                            var v29 float64 = 3.14159
                            v1 = (v1 * v1)
                            _ = fi28
                            _ = v29
                        }
                    }
                    var v30 int = fi27
                    var v31 float32 = (1.5 + 0.1)
                    _ = fi27
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 15; fi32++ {
                    v14 = (^(v14))
                    _ = fi32
                }
            }
            v3 += (v3 | v3)
            if (!(true && false)) {
                if ((v14 - genFn0(2, 241, 2147483647)) == v14) {
                    for fi33 := 0; fi33 < 13; fi33++ {
                        v1 -= v1
                        fmt.Println(v3, v14, v1)
                        v34 := GenS1{F0: (^(fi33)), F1: 65535}
                        _ = fi33
                        _ = v34
                    }
                }
                v2[42] = 173
            }
        }
        fmt.Println(v3)
    }
    v3 -= 42
    v2[-1] = 42
    var v35 uint8 = (^(v14))
    v36 := []int{769}
    v37 := GenS1{F0: 32768, F1: 65535}
    var v38 uint = 0
    for fi39 := 0; fi39 < 7; fi39++ {
        v3 = (^(v3))
        _ = fi39
    }
    if false {
        var v40 bool = (false || true)
        fmt.Println(v40, v1, v1)
        _ = v40
    }
    if false {
        v3 += v3
        var v41 float64 = 1.5
        v42 := GenS5{F0: v35, F1: (v38 ^ v38), F2: (v38 + (v38 / (((v38 / ((v38) | 1))) | 1)))}
        _ = v41
        _ = v42
    } else {
        var v43 uint8 = v35
        fmt.Println(v35)
        v43 = (v43 - v35)
        _ = v43
    }
    switch "" {
    case "\n\t":
        for fi44 := 0; fi44 < 17; fi44++ {
            v14 = genFn0((v38 | v38), 54, 0)
            _ = fi44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
