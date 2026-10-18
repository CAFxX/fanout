package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint8; F1 int8; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 bool }
func genFn0(p0 int16, p1 bool) uint16 {
    v7 := GenS1{F0: 2}
    for fi8 := 0; fi8 < 19; fi8++ {
        if p1 {
            p0 -= (-(p0))
            var v9 int = fi8
            var v10 uint64 = 0
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                p0 += p0
                v12 := []int{fi8, fi11}
                if (-29 == -1) {
                    p1 = (0.1 >= 0.0)
                    if (p1 || p1) {
                        if p1 {
                            if (!(!true)) {
                                if (true || (!(((902 < 0) || false) && p1))) {
                                    if p1 {
                                        if false {
                                            switch ("hello" + "\n\t") {
                                            case "\n\t":
                                                v13 := GenS1{F0: 300}
                                                _ = v13
                                            case "hello":
                                                var v14 uint8 = 110
                                                _ = v14
                                            default:
                                                v15 := GenS1{F0: 232}
                                                _ = v15
                                            }
                                            v16 := GenS2{F0: 184, F1: p0}
                                            var v17 float32 = 1.5
                                            _ = v16
                                            _ = v17
                                        }
                                    }
                                    v12 = append(v12, fi8)
                                    var v18 int16 = p0
                                    _ = v18
                                }
                                p0 = (^(p0))
                                p1 = true
                            }
                            p0 = p0
                            var v19 uint8 = 255
                            var v20 float32 = 1e10
                            if (432 <= 0) {
                                v12 = append(v12, fi8)
                            } else {
                                if false {
                                    fmt.Println(p0, v19, p0)
                                    switch 127 {
                                    case -1:
                                        switch fi8 {
                                        case 0:
                                            var v21 float32 = v20
                                            _ = v21
                                        case -408:
                                            v22 := GenS1{F0: 720}
                                            _ = v22
                                        }
                                    }
                                    v23 := GenS1{F0: 42}
                                    _ = v23
                                } else {
                                    if (0 > 107) {
                                        if true {
                                            v20 += v20
                                            v24 := GenS1{F0: 0}
                                            v12 = append(v12, (-(fi8)))
                                            v25 := GenS3{F0: 128, F1: 1, F2: (^(p0))}
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 bool = ((0 >= 42) && ((1.5 * float64((v20 / v20))) != 1e10))
                                        fmt.Println(p0, p0, fi8)
                                        _ = v26
                                    }
                                    v27 := GenS2{F0: 42, F1: (p0 / ((p0) | 1))}
                                    v28 := GenS1{F0: 4294967295}
                                    _ = v27
                                    _ = v28
                                }
                                if true {
                                    if true {
                                        p1 = p1
                                        p1 = p1
                                        var v29 uint16 = uint16(v19)
                                        _ = v29
                                    } else {
                                        p0 -= p0
                                        var v30 uint8 = (v19 / ((2) | 1))
                                        fmt.Println(p1, p0, fi11)
                                        _ = v30
                                    }
                                    v31 := GenS1{F0: 238}
                                    if p1 {
                                        v12 = append(v12, 1)
                                        fmt.Println(v20)
                                        v19 = (v19 >> ((18446744073709551615) & 63))
                                    } else {
                                        v32 := GenS4{F0: 380, F1: 18446744073709551615}
                                        var v33 int32 = -1
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v31
                                } else {
                                    var v34 uint32 = 727
                                    _ = v34
                                }
                            }
                            _ = v19
                            _ = v20
                        }
                        if p1 {
                            fmt.Println(fi11, p0)
                            p0 += p0
                        }
                        fmt.Println(fi11, fi8, fi8)
                    }
                }
                _ = fi11
                _ = v12
            }
            p1 = (305 != 2)
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1() string {
    if (!true) {
        v35 := []int{93, 127, 42, -2147483648}
        _ = v35
    }
    var v36 uint = 42
    if true {
        v36 -= v36
        for fi37 := 0; fi37 < 7; fi37++ {
            for fi38 := 0; fi38 < 9; fi38++ {
                var v39 int64 = -1000
                v40 := []int{int(v36), (fi38 - fi37), (fi37 & 42)}
                v39 -= -9223372036854775808
                var v41 bool = false
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi37
        }
    } else {
        v42 := GenS5{F0: 106}
        _ = v42
    }
    v36 -= (v36 / ((935) | 1))
    switch "\n\t" {
    case "foo bar":
        v36 = (v36 + (v36 | 18446744073709551615))
    case "":
        v36 += v36
    case "hello":
        v36 = 1
    }
    _ = v36
    return ("\n\t" + "")
}

func main() {
    v1 := []uint32{349, 898}
    v2 := make(map[int]uint32)
    var v3 uint = 65535
    var v4 string = ("\n\t" + "foo bar")
    var v5 int = 42
    v6 := make(map[string]int)
    if false {
        var v43 int8 = -1
        switch v4 {
        case "foo bar":
            v3 = v3
        case "a":
            v5 += int(v3)
        case "\n\t":
            v3 = 577
        }
        _ = v43
    } else {
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 uint64 = 1
            if false {
                fmt.Println(fi44, fi44, v5)
                var v46 uint8 = 42
                v3 -= 4294967295
                _ = v46
            }
            _ = fi44
            _ = v45
        }
        _ = len(v1)
        var v47 float32 = 1e10
        _ = v47
    }
    switch v4 {
    case "\n\t":
        v4 = "a"
    }
    v3 = (v3 << (((v3 / (((v3 % (((v3 % ((v3) | 1))) | 1))) | 1))) & 63))
    v5 = 42
    switch (v5 << (((v3 & v3)) & 63)) {
    case -54:
        v48 := GenS1{F0: 2}
        _ = v48
    case 127:
        var v49 bool = (v3 == 65535)
        _ = v49
    }
    if ((((v4 + (v4 + "\n\t")) + (v4 + v4)) + v4) != (v4 + (v4 + v4))) {
        var v50 uint16 = 65535
        _ = v50
    } else {
        v51 := GenS4{F0: 4294967295, F1: (v3 >> ((255) & 63))}
        fmt.Println(v4)
        _ = v51
    }
    switch v5 {
    case 0:
        fmt.Println(v3, v3)
    case 1:
        var v52 int = (v5 * (^(v5)))
        _ = v52
    case -9223372036854775808:
        if (v5 < v5) {
            if true {
                v53 := GenS6{F0: 65535, F1: (float32(-2.25) >= float32(1.5))}
                _ = v53
            }
            var v54 uint = (v3 | 4294967295)
            _ = v54
        } else {
            if false {
                v4 = (v4 + ("hello" + v4))
                for fi55 := 0; fi55 < 11; fi55++ {
                    for fi56 := 0; fi56 < 5; fi56++ {
                        v5 -= (^(fi56))
                        v57 := GenS5{F0: -1}
                        _ = fi56
                        _ = v57
                    }
                    var v58 int32 = 1
                    var v59 string = v4
                    _ = fi55
                    _ = v58
                    _ = v59
                }
            }
            v60 := GenS6{F0: 890, F1: (((3.14159 - (((0.0 * float32(0.1)) * (1e10 / ((((1.5 * -2.25) * float32(3.14159)) + (0.0 * 0.0)) * 3.14159))) + float32(float64(((0.0 - float32((0.0 / 1.5))) - 1.5))))) > 1e10) && true)}
            v61 := v6[(v4 + (v4 + (v4 + ("a" + "foo bar"))))]
            _ = v60
            _ = v61
        }
    }
    var v62 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
}
