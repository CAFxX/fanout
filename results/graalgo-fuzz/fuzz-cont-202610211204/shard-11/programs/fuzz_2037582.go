package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int64, p1 string, p2 uint64) uint8 {
    var v7 uint = 2
    v8 := []int{}
    p2 = p2
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int) uint8 {
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{fi9, p0, fi9}
        var v11 uint8 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS1{F0: 42, F1: -527}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            fmt.Println(fi14, fi14, fi14)
            _ = fi15
        }
        p0 = (p0 / (((fi14 % ((p0) | 1))) | 1))
        _ = fi14
    }
    _ = p0
    return 26
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 int8 = -1
    var v5 uint16 = 255
    var v6 uint = 1
    fmt.Println(v5, v6)
    for fi16 := 0; fi16 < 17; fi16++ {
        v4 -= 127
        _ = fi16
    }
    v2[703] = 2147483647
    if false {
        if (!(!false)) {
            switch 127 {
            case -1:
                var v17 int64 = -9223372036854775808
                _ = v17
            default:
                var v18 float32 = 0.0
                _ = v18
            }
            var v19 uint64 = uint64(v6)
            _ = v19
        } else {
            v6 = 1
        }
    } else {
        if (1 <= 18446744073709551615) {
            v4 = v4
        } else {
            v20 := GenS1{F0: 2147483647, F1: 200}
            _ = v20
        }
        switch "a" {
        case "foo bar":
            v5 += 556
        }
        v4 -= 100
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        fmt.Println(v4, v4, v4)
        _ = fi21
    }
    switch "\n\t" {
    case "":
        fmt.Println(v6)
    case "\n\t":
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 uint32 = 0
            v2[0] = v23
            if false {
                var v24 float64 = ((1.5 - ((((3.14159 + (((3.14159 * (1.5 / 1.0)) - (-2.25 + 3.14159)) / 0.1)) * (1.5 / float64(float32((0.1 + 1.5))))) - ((-2.25 / 3.14159) - 3.14159)) / 1.0)) - (float64((-2.25 - 0.0)) / 1.0))
                _ = v24
            } else {
                for fi25 := 0; fi25 < 11; fi25++ {
                    v26 := GenS2{F0: 961, F1: "a", F2: (v23 / (((^(v23))) | 1))}
                    switch -203 {
                    case 4294967295:
                        if true {
                            v27 := GenS1{F0: 0, F1: 0}
                            v23 = (v23 >> ((v6) & 63))
                            v4 -= v4
                            _ = v27
                        } else {
                            v28 := v1["foo bar"]
                            _ = v28
                        }
                    case -330:
                        var v29 int32 = 0
                        _ = v29
                    }
                    if (v6 > 367) {
                        v30 := GenS2{F0: 42, F1: "foo bar", F2: (v23 << ((v6) & 63))}
                        fmt.Println(fi22, v5)
                        _ = v30
                    }
                    if (!(false || false)) {
                        v31 := GenS1{F0: 42, F1: -270}
                        v32 := []int{(fi25 % (((fi22 | fi25)) | 1)), 2147483647, (fi25 / (((fi25 % (((fi22 / ((fi22) | 1))) | 1))) | 1))}
                        _ = v31
                        _ = v32
                    }
                    if (0 < 42) {
                        v23 -= (v23 & (v23 * v23))
                    } else {
                        switch (fi22 ^ fi25) {
                        case -9223372036854775808:
                            switch fi25 {
                            case 128:
                                v5 = (v5 / ((2) | 1))
                            case 42:
                                var v33 string = ""
                                _ = v33
                            case 2147483647:
                                v4 += 100
                                fmt.Println(v4, v23)
                                _ = len(v1)
                            }
                        case 128:
                            fmt.Println(fi25)
                        case 4294967295:
                            v34 := v2[(^(v23))]
                            _ = v34
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                var v35 int8 = 127
                for fi36 := 0; fi36 < 2; fi36++ {
                    v37 := []int{fi22, (fi22 + fi22), fi22}
                    var v38 uint16 = (v5 >> ((2) & 63))
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = v35
            }
            _ = fi22
            _ = v23
        }
    }
    v39 := []int{2147483647, 128, 2147483647, 641}
    fmt.Println(v4)
    v6 += v6
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := GenS3{F0: 42}
        _ = fi40
        _ = v41
    }
    var v42 bool = ((-118 < 0) || (!false))
    v43 := GenS1{F0: -575, F1: 1}
    v4 += (v4 % ((v4) | 1))
    if v42 {
        fmt.Println(v5, v5, v42)
    } else {
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
        var v44 int32 = 508
        fmt.Println(v4, v6)
        _ = v44
    }
    v39 = append(v39, 2147483647)
    if true {
        switch "foo bar" {
        case "":
            v45 := GenS4{F0: -678, F1: 1, F2: 42}
            _ = v45
        case "\n\t":
            var v46 uint32 = 42
            _ = v46
        }
        v47 := GenS5{F0: -941}
        _ = v47
    }
    var v48 uint = (^(v6))
    delete(v2, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v48
    _ = v5
    _ = v6
}
