package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 bool }
func genFn0(p0 int16, p1 uint64) int64 {
    p1 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{fi7, 0}
        p1 += (p1 | 131)
        var v9 int16 = p0
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint = 483
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := []int{fi11, fi11}
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: (p1 * (p1 >> (((v10 + (v10 * v10))) & 63)))}
        _ = fi11
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    var v1 int16 = 113
    v2 := []string{}
    var v3 int16 = 1000
    v4 := make(map[uint32]string)
    var v5 int64 = 739
    v6 := []uint32{2, 0}
    if false {
        if (!false) {
            fmt.Println(v5, v3, v5)
            var v15 bool = false
            if (!(393 > 42)) {
                v16 := GenS2{F0: -521}
                for fi17 := 0; fi17 < 20; fi17++ {
                    v18 := GenS2{F0: fi17}
                    fmt.Println(v1, v1)
                    var v19 uint16 = 2
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := v4[244]
                _ = v16
                _ = v20
            }
            _ = v15
        }
    } else {
        if (!((true && true) && false)) {
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS1{F0: 4294967295}
                for fi23 := 0; fi23 < 6; fi23++ {
                    var v24 int8 = 1
                    fmt.Println(fi21, fi23)
                    if (v24 >= 42) {
                        var v25 uint64 = 2
                        if true {
                            if (2 != 2) {
                                switch (fi23 * (fi21 + -2147483648)) {
                                case 1:
                                    v26 := GenS2{F0: (fi23 % ((fi23) | 1))}
                                    _ = v26
                                case 42:
                                    var v27 float64 = ((float64((0.1 - 3.14159)) * 1e10) * (0.0 + float64(((1e10 * -2.25) - (1e10 + 0.1)))))
                                    _ = v27
                                case 2147483647:
                                    v25 = (v25 | v25)
                                default:
                                    v24 += (v24 % ((v24) | 1))
                                }
                                v28 := GenS3{F0: fi21, F1: 65535, F2: 65}
                                _ = v28
                            } else {
                                var v29 float32 = (-2.25 + 0.1)
                                v30 := GenS1{F0: v25}
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v1, v5)
                                case "":
                                    fmt.Println(v25, v24)
                                default:
                                    var v31 bool = (((v25 / ((0) | 1)) > v25) || false)
                                    _ = v31
                                }
                                if len(v6) > 0 {
                                    v6[0] = 42
                                }
                                v1 += v3
                                _ = v29
                                _ = v30
                            }
                            v32 := v4[939]
                            _ = v32
                        }
                        _ = v25
                    }
                    fmt.Println(v3)
                    switch (fi21 / (((fi23 - fi21)) | 1)) {
                    case 42:
                        switch -295 {
                        case 1:
                            v33 := v4[42]
                            _ = v33
                        case -1:
                            v5 += (v5 ^ (v5 + v5))
                        case -737:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
            for fi34 := 0; fi34 < 13; fi34++ {
                v5 += (v5 - int64(fi34))
                v1 -= (v1 ^ 42)
                _ = len(v2)
                _ = fi34
            }
        }
        v5 = int64(v1)
        v3 = 1000
    }
    if (false && true) {
        var v35 bool = true
        _ = v35
    } else {
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 bool = (false || (861 < 2147483647))
            _ = fi36
            _ = v37
        }
        switch 32768 {
        case 4294967295:
            v3 = (v3 << ((255) & 63))
        default:
            v1 += (v3 & v1)
        }
    }
    v38 := GenS4{F0: (v5 << ((905) & 63)), F1: -481}
    v1 -= v3
    v3 = (v1 | (v1 - -32768))
    var v39 uint = 18446744073709551615
    if true {
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 uint8 = 2
            fmt.Println(v41, v5, fi40)
            var v42 int32 = -836
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS5{F0: true}
        _ = v43
    }
    v3 += (-(v1))
    if (0 <= 42) {
        for fi44 := 0; fi44 < 3; fi44++ {
            v1 -= v1
            var v45 int = 42
            _ = fi44
            _ = v45
        }
        for fi46 := 0; fi46 < 18; fi46++ {
            v1 -= -333
            v3 = (v3 >> ((44) & 63))
            v1 -= (v1 - v3)
            _ = fi46
        }
        switch 0 {
        case 32767:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 0:
            v47 := GenS2{F0: -744}
            _ = v47
        }
    } else {
        v4[1] = "a"
    }
    var v48 int16 = v1
    v5 += v5
    v39 += v39
    for fi49 := 0; fi49 < 12; fi49++ {
        _ = len(v6)
        _ = fi49
    }
    v48 = -1
    fmt.Println(v48, v3, v39)
    var v50 int32 = -546
    v5 = (v5 ^ v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v48
    _ = v5
    _ = v50
    _ = v6
}
