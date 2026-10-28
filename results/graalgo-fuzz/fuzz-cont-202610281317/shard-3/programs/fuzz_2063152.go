package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 uint16, p1 uint64) int64 {
    fmt.Println(p1, p0)
    v7 := []int{128, 2147483647}
    v8 := GenS1{F0: -13, F1: false}
    var v9 uint = 255
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int64) uint32 {
    p0 += p0
    for fi10 := 0; fi10 < 16; fi10++ {
        if ((fi10 * fi10) > -143) {
            for fi11 := 0; fi11 < 7; fi11++ {
                if (320 > 0) {
                    v12 := GenS2{F0: 128, F1: true}
                    fmt.Println(p0, fi10)
                    _ = v12
                }
                if true {
                    p0 += (p0 & 0)
                    fmt.Println(fi10, fi10, p0)
                    fmt.Println(fi11)
                }
                var v13 uint16 = 524
                v14 := []int{fi11, 448, (fi10 >> ((0) & 63)), fi10}
                p0 += (p0 - (p0 << ((4294967295) & 63)))
                var v15 int64 = p0
                _ = fi11
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 int16 = 0
            if ((p0 << ((554) & 63)) < p0) {
                v17 := []int{}
                _ = len(v17)
                _ = len(v17)
                _ = v17
            }
            _ = v16
        }
        p0 -= p0
        _ = fi10
    }
    _ = p0
    return 790
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 0
    var v3 int16 = -32768
    v4 := []uint32{}
    v5 := make(map[string]uint32)
    var v6 float64 = 1.5
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            if false {
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 uint = 18446744073709551615
                    _ = fi19
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                fmt.Println(v3, v2, v6)
                var v22 uint8 = 255
                _ = fi21
                _ = v22
            }
            _ = fi18
        }
        _ = len(v4)
        v23 := GenS2{F0: 128, F1: (!(float64(1.5) > 1e10))}
        _ = v23
    } else {
        var v24 uint64 = 429
        var v25 uint32 = 133
        switch 0 {
        case 2147483647:
            if false {
                var v26 int64 = 9223372036854775807
                v26 += v26
                _ = v26
            } else {
                v3 = v3
                v27 := GenS1{F0: (^(v2)), F1: (-1 > -2147483648)}
                v28 := GenS1{F0: -1, F1: false}
                _ = v27
                _ = v28
            }
        default:
            var v29 float32 = 0.0
            _ = v29
        }
        _ = v24
        _ = v25
    }
    v2 -= (-(v2))
    var v30 int16 = v3
    var v31 bool = (!false)
    var v32 float64 = v6
    fmt.Println(v2, v32, v31)
    var v33 float64 = (v6 - v6)
    v2 += (-(v2))
    v34 := GenS3{F0: -108, F1: -9223372036854775808, F2: -1}
    if (255 > 1) {
        var v35 uint32 = 65535
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v35
    }
    var v36 uint32 = 4294967295
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v37 uint = 194
    for fi38 := 0; fi38 < 17; fi38++ {
        v39 := GenS3{F0: 1, F1: genFn0(32767, 42), F2: 98}
        for fi40 := 0; fi40 < 5; fi40++ {
            v5["foo bar"] = (v36 * genFn1(-1))
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    if ("a" < "hello") {
        switch 214 {
        case -377:
            var v41 int64 = 9223372036854775807
            _ = v41
        }
    }
    if v31 {
        if v31 {
            fmt.Println(v3, v37, v31)
        }
        v42 := GenS4{F0: 233}
        if true {
            for fi43 := 0; fi43 < 9; fi43++ {
                v44 := GenS1{F0: (^(v2)), F1: true}
                var v45 int16 = v30
                v6 -= float64(((0.0 * 1.5) - ((0.0 * 0.0) - ((0.1 * 0.1) * 1e10))))
                _ = fi43
                _ = v44
                _ = v45
            }
            v36 -= genFn1(1)
            var v46 int = 1
            _ = v46
        }
        _ = v42
    } else {
        fmt.Println(v30, v2, v6)
    }
    v47 := GenS5{F0: 42, F1: false, F2: (^(v37))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
