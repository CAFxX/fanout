package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int16; F1 string; F2 uint8 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 string }
func genFn0(p0 uint8, p1 int8, p2 int16) int64 {
    if true {
        fmt.Println(p1)
        var v7 uint16 = uint16(p0)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 bool = (32767 != 0)
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        p2 = p2
        if (42 != 0) {
            if true {
                v11 := []int{}
                switch -916 {
                case 127:
                    var v12 string = ""
                    _ = v12
                }
                _ = v11
            }
            v13 := []int{}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v13
        } else {
            p0 = (p0 * 127)
            v14 := GenS1{F0: 144}
            _ = v14
        }
        _ = fi10
    }
    var v15 uint64 = 835
    switch -478 {
    case 4294967295:
        p2 = p2
    default:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 0
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    v5 := []uint32{2147483647, 4294967295}
    var v6 int64 = -9223372036854775808
    v16 := GenS2{F0: 851}
    v17 := GenS3{F0: false}
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint = 255
        if len(v5) > 0 {
            v5[0] = 2
        }
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 0
    v20 = v20
    var v21 string = ("\n\t" + "foo bar")
    v1 = 1.5
    var v22 uint = 0
    v1 += float64(float32(1e10))
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            if true {
                v25 := GenS4{F0: 127, F1: (!false)}
                switch fi23 {
                case 32767:
                    if len(v5) > 0 {
                        v5[0] = 2147483648
                    }
                case -2147483648:
                    if true {
                        v3[2147483647] = ((^(v22)) >= v22)
                        var v26 uint = (v22 / ((v22) | 1))
                        _ = v26
                    }
                case 127:
                    fmt.Println(v21, v1)
                }
                if (true || (-1 == 12)) {
                    v27 := GenS1{F0: (v22 >> (((v22 ^ v22)) & 63))}
                    switch 705 {
                    case 127:
                        v28 := GenS5{F0: 0, F1: ((v21 + "a") + v21), F2: 173}
                        _ = v28
                    }
                    var v29 int8 = 42
                    _ = v27
                    _ = v29
                }
                _ = v25
            } else {
                var v30 bool = (!true)
                var v31 uint16 = 255
                v32 := GenS6{F0: 32767, F1: (811 != 813)}
                v4[(fi24 ^ fi24)] = (fi23 / (((fi23 - (fi23 & 128))) | 1))
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi24
        }
        var v33 string = v21
        _ = fi23
        _ = v33
    }
    v34 := GenS7{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
