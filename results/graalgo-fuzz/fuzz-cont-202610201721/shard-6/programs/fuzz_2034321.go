package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int8; F1 uint; F2 int }
func genFn0() int16 {
    var v5 int = 56
    v5 -= v5
    var v6 float64 = (3.14159 * 1e10)
    v7 := []int{4294967295}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            switch (fi8 / ((0) | 1)) {
            case 32767:
                var v10 string = ""
                _ = v10
            case 543:
                v11 := GenS1{F0: 1000, F1: 42}
                _ = v11
            default:
                switch "a" {
                case "a":
                    var v12 uint64 = 42
                    _ = v12
                case "\n\t":
                    v13 := GenS2{F0: 221}
                    v14 := GenS2{F0: 128}
                    fmt.Println(fi8, fi9)
                    var v15 float64 = v6
                    switch "hello" {
                    case "a":
                        switch fi9 {
                        case 281:
                            var v16 uint = 0
                            _ = v16
                        }
                        _ = len(v7)
                        var v17 uint = 0
                        _ = v17
                    case "foo bar":
                        v5 = v5
                        fmt.Println(v15, v15, v15)
                        v5 = -280
                    case "":
                        if (!true) {
                            v6 += 1e10
                            switch "foo bar" {
                            case "hello":
                                v7 = append(v7, 32768)
                            case "foo bar":
                                v18 := GenS3{F0: "hello"}
                                _ = v18
                            default:
                                var v19 int64 = 0
                                _ = v19
                            }
                            v7 = append(v7, fi9)
                            v6 = v6
                            var v20 int64 = 566
                            _ = v20
                        }
                    }
                    switch "" {
                    case "hello":
                        if len(v7) > 0 {
                            v7[0] = (-(fi8))
                        }
                        v15 = 0.1
                        v21 := GenS2{F0: 255}
                        _ = v21
                    case "":
                        fmt.Println(fi8, fi9)
                    }
                    _ = v13
                    _ = v14
                    _ = v15
                case "":
                    if len(v7) > 0 {
                        v7[0] = v5
                    }
                }
            }
            v5 -= v5
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi9
        }
        var v22 int16 = 1000
        var v23 uint16 = 32767
        _ = fi8
        _ = v22
        _ = v23
    }
    _ = v5
    _ = v6
    _ = v7
    return -63
}
func genFn1(p0 bool) uint32 {
    if (false && p0) {
        p0 = true
    } else {
        v24 := GenS1{F0: -32768, F1: 928}
        _ = v24
    }
    p0 = p0
    for fi25 := 0; fi25 < 1; fi25++ {
        p0 = p0
        var v26 uint8 = 0
        _ = fi25
        _ = v26
    }
    p0 = (1 == -291)
    _ = p0
    return 42
}

func main() {
    var v1 int = -753
    v2 := []int{(-(v1)), v1, (^(v1)), (^(v1))}
    v3 := make(map[string]bool)
    v4 := []uint32{13}
    v1 = (v1 << ((925) & 63))
    v1 = (v1 & v1)
    v1 -= v1
    v27 := GenS2{F0: 1}
    var v28 uint8 = 2
    v29 := v3["foo bar"]
    if (!false) {
        _ = len(v4)
    }
    if (1 >= -146) {
        v28 -= 128
        v29 = (((-2.25 / (3.14159 - (float32((1e10 * 0.0)) / 0.1))) >= float32(0.1)) || (!(586 >= 255)))
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 uint = 42
            _ = fi30
            _ = v31
        }
    }
    if v29 {
        v32 := GenS3{F0: ""}
        v28 -= (^(v28))
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        v29 = (146 > 2)
        fmt.Println(v1, v29, v1)
        if (21 < fi33) {
            v34 := GenS4{F0: 1, F1: v28}
            if v29 {
                for fi35 := 0; fi35 < 17; fi35++ {
                    var v36 uint64 = 4294967295
                    v37 := GenS3{F0: "foo bar"}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                v38 := GenS2{F0: (v28 | v28)}
                var v39 int = fi33
                _ = v38
                _ = v39
            } else {
                v2 = append(v2, -1)
                fmt.Println(v1)
            }
            _ = v34
        } else {
            v28 = (v28 % ((v28) | 1))
            if v29 {
                v2 = append(v2, v1)
            }
        }
        _ = fi33
    }
    if (32767 < int16(v28)) {
        var v40 int16 = -1
        _ = v40
    } else {
        v41 := GenS5{F0: 42, F1: 396, F2: (v1 % ((v1) | 1))}
        v28 = 128
        var v42 uint8 = (^(v28))
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        fmt.Println(v29, v29, v1)
        _ = fi43
    }
    switch "a" {
    case "\n\t":
        v44 := GenS1{F0: -920, F1: 4294967295}
        _ = v44
    }
    fmt.Println(v1, v28)
    for fi45 := 0; fi45 < 9; fi45++ {
        fmt.Println(v29)
        _ = fi45
    }
    var v46 string = ("\n\t" + ("foo bar" + ""))
    if ((-2.25 - 1.5) != 0.0) {
        _ = len(v3)
    } else {
        v29 = (42 != 640)
    }
    delete(v3, (("foo bar" + v46) + v46))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v46
}
