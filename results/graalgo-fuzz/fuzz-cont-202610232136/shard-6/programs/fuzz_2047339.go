package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int64; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int }
func genFn0(p0 int64) int16 {
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            if false {
                var v6 float32 = 0.1
                switch "\n\t" {
                case "a":
                    v7 := GenS1{F0: 119, F1: (-(fi5))}
                    _ = v7
                case "hello":
                    v8 := []int{fi4}
                    if ((!false) && (2 > 42)) {
                        fmt.Println(fi5)
                        v9 := GenS1{F0: 0, F1: 128}
                        v10 := GenS1{F0: 128, F1: -558}
                        _ = v10
                        _ = v9
                    }
                    p0 = (p0 << ((42) & 63))
                    _ = v8
                default:
                    fmt.Println(fi5, fi5)
                }
                if false {
                    v11 := GenS2{F0: ((0.1 + 0.1) > 0.0)}
                    v12 := []int{-2147483648, (^(fi5)), fi4}
                    _ = v11
                    _ = v12
                }
                switch (fi5 | fi5) {
                case 1:
                    v13 := []int{}
                    _ = v13
                default:
                    v14 := GenS1{F0: 102, F1: (fi5 ^ fi4)}
                    _ = v14
                }
                var v15 int32 = 963
                fmt.Println(p0)
                var v16 float64 = 0.0
                _ = v15
                _ = v16
                _ = v6
            }
            _ = fi5
        }
        p0 += 9223372036854775807
        var v17 uint = 2
        _ = fi4
        _ = v17
    }
    p0 -= (p0 % ((p0) | 1))
    switch "" {
    case "foo bar":
        for fi18 := 0; fi18 < 16; fi18++ {
            switch (fi18 % ((fi18) | 1)) {
            case 32767:
                switch (-(fi18)) {
                case 127:
                    v19 := GenS3{F0: 959, F1: p0, F2: 867}
                    _ = v19
                case -2147483648:
                    switch "hello" {
                    case "hello":
                        v20 := []int{}
                        _ = v20
                    case "a":
                        if true {
                            switch ("" + "") {
                            case "\n\t":
                                switch fi18 {
                                case 32767:
                                    p0 -= p0
                                case -9223372036854775808:
                                    var v21 int16 = -1
                                    _ = v21
                                case 1:
                                    var v22 int64 = (p0 << ((2) & 63))
                                    _ = v22
                                }
                            }
                            v23 := GenS4{F0: 65535}
                            var v24 float64 = (0.0 * 3.14159)
                            _ = v23
                            _ = v24
                        }
                    case "\n\t":
                        p0 = p0
                    }
                }
            default:
                fmt.Println(p0, fi18)
            }
            v25 := []int{fi18, 32768}
            _ = fi18
            _ = v25
        }
    case "":
        v26 := []int{-171, 127, 4294967295, -2147483648}
        _ = v26
    }
    _ = p0
    return 375
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1000
    var v27 int16 = genFn0(944)
    switch 0 {
    case -1:
        var v28 bool = true
        _ = v28
    case 32767:
        var v29 float32 = (-2.25 / float32(1e10))
        _ = v29
    case -450:
        v30 := GenS5{F0: 975, F1: 42}
        _ = v30
    }
    var v31 float64 = 0.1
    v32 := GenS6{F0: 0, F1: 65535, F2: 768}
    for fi33 := 0; fi33 < 19; fi33++ {
        if true {
            for fi34 := 0; fi34 < 9; fi34++ {
                v35 := GenS3{F0: 2147483647, F1: 0, F2: 0}
                _ = fi34
                _ = v35
            }
            var v36 uint16 = 32767
            switch "hello" {
            case "":
                fmt.Println(fi33, v31, fi33)
            }
            _ = v36
        } else {
            switch "foo bar" {
            case "a":
                delete(v2, 2147483648)
            }
        }
        v37 := GenS7{F0: v3, F1: 784}
        for fi38 := 0; fi38 < 15; fi38++ {
            v39 := []int{fi33, -9223372036854775808}
            v40 := GenS2{F0: true}
            v27 -= v3
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi33
        _ = v37
    }
    v31 += -2.25
    fmt.Println(v31)
    v41 := GenS8{F0: 1}
    var v42 uint32 = 0
    v27 = genFn0(1)
    v43 := GenS9{F0: 42}
    var v44 int = 587
    for fi45 := 0; fi45 < 16; fi45++ {
        v31 = v31
        v3 = v3
        var v46 float32 = ((1.5 / 1.0) * ((float32((v31 - 1e10)) - (float32(v31) / -2.25)) + (1e10 / 0.1)))
        _ = fi45
        _ = v46
    }
    v27 = v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
