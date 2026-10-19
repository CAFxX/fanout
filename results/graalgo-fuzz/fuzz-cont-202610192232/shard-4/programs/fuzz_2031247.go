package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int32; F1 int32 }
func genFn0(p0 uint64) int8 {
    v7 := []int{533, 42}
    p0 = p0
    p0 = (p0 / ((p0) | 1))
    _ = p0
    _ = v7
    return -46
}
func genFn1(p0 uint, p1 uint32, p2 uint16) int32 {
    p1 = 1
    switch -952 {
    case 32767:
        var v8 int = 32767
        _ = v8
    case -1:
        var v9 uint32 = 264
        _ = v9
    }
    switch 127 {
    case 32768:
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                p0 = (p0 | (p0 >> ((uint(p1)) & 63)))
                _ = fi11
            }
            var v12 uint16 = (^(p2))
            if false {
                v13 := []int{(fi10 & 42), 1, (fi10 ^ (fi10 % ((fi10) | 1))), fi10}
                _ = v13
            }
            _ = fi10
            _ = v12
        }
    case 56:
        var v14 uint32 = p1
        _ = v14
    case -1:
        v15 := GenS1{F0: 0}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 78
}

func main() {
    var v1 uint8 = 29
    var v2 string = ("a" + ("" + ("" + "hello")))
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 uint = 42
    var v6 int32 = 2147483647
    if (((float64(1e10) + 3.14159) * 0.1) > (0.0 - (-2.25 - -2.25))) {
        v5 += v5
        var v16 uint64 = uint64(v1)
        v16 -= (v16 - (v16 ^ v16))
        _ = v16
    } else {
        fmt.Println(v2, v5, v5)
    }
    v17 := []int{877, 2147483647, 810}
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 string = v2
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: -1, F1: 2}
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS2{F0: 32767, F1: v1}
        if true {
            v6 -= (v6 / ((v6) | 1))
        }
        var v23 uint = v5
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            v5 += 255
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi25
        }
        v2 = "a"
        v2 = ("foo bar" + v2)
        _ = fi24
    }
    if false {
        var v26 uint8 = (^(v1))
        _ = v26
    }
    var v27 float32 = (1e10 * (-2.25 - 0.0))
    var v28 float32 = 1.5
    if (false && (1000 <= 52)) {
        for fi29 := 0; fi29 < 16; fi29++ {
            if false {
                switch "a" {
                case "hello":
                    var v30 int = (^(fi29))
                    _ = v30
                case "foo bar":
                    fmt.Println(v1, v1)
                default:
                    if true {
                        var v31 int = fi29
                        var v32 int64 = 235
                        _ = v31
                        _ = v32
                    }
                }
                v17 = append(v17, 978)
            }
            _ = fi29
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            v2 = ""
            v34 := GenS3{F0: 42, F1: v2}
            _ = fi33
            _ = v34
        }
        if (!(false && true)) {
            v5 = (v5 ^ 18446744073709551615)
            v17 = append(v17, 42)
            v35 := GenS2{F0: 1, F1: (v1 & v1)}
            _ = v35
        }
    }
    v36 := GenS4{F0: v6, F1: (v6 | 2147483647)}
    v1 += (v1 ^ (v1 * v1))
    for fi37 := 0; fi37 < 13; fi37++ {
        for fi38 := 0; fi38 < 14; fi38++ {
            v39 := GenS2{F0: -321, F1: (^(v1))}
            if ((v28 >= v28) || true) {
                var v40 uint32 = 799
                _ = v40
            } else {
                if (!(!true)) {
                    if true {
                        var v41 string = v2
                        _ = v41
                    }
                    v28 = v28
                    switch "a" {
                    case "\n\t":
                        switch (v2 + ("foo bar" + "hello")) {
                        case "\n\t":
                            v42 := v4[fi37]
                            _ = v42
                        case "a":
                            fmt.Println(v28, v2)
                        case "":
                            if (((0.1 * 1e10) >= 3.14159) && false) {
                                fmt.Println(v28)
                            }
                        default:
                            var v43 float32 = (v27 + v28)
                            _ = v43
                        }
                    case "hello":
                        var v44 uint64 = 2
                        var v45 int16 = 723
                        if true {
                            v45 -= 0
                            v6 -= v6
                        } else {
                            if len(v17) > 0 {
                                v17[0] = fi38
                            }
                        }
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(fi37, v28)
                            switch ("\n\t" + "a") {
                            case "foo bar":
                                if (0 < 0) {
                                    v1 += (v1 << (((v5 / ((v5) | 1))) & 63))
                                    v6 = 2147483647
                                    v46 := GenS3{F0: 18446744073709551615, F1: v2}
                                    _ = v46
                                } else {
                                    switch fi38 {
                                    case -9223372036854775808:
                                        v6 -= (v6 & (v6 - (v6 >> (((^(v5))) & 63))))
                                    case 667:
                                        var v47 int16 = (v45 + (v45 % ((v45) | 1)))
                                        _ = v47
                                    case 128:
                                        v48 := v4[fi38]
                                        _ = v48
                                    default:
                                        var v49 int32 = (v6 % ((v6) | 1))
                                        _ = v49
                                    }
                                }
                            case "\n\t":
                                switch (fi37 - -190) {
                                case 42:
                                    switch -794 {
                                    case 9223372036854775807:
                                        v28 += (v27 / v27)
                                    }
                                }
                            }
                            var v50 float32 = -2.25
                            _ = v50
                        case "hello":
                            fmt.Println(v45)
                        }
                        _ = v44
                        _ = v45
                    }
                } else {
                    var v51 uint64 = 775
                    fmt.Println(v5)
                    v28 = 1.5
                    var v52 int8 = 0
                    _ = v51
                    _ = v52
                }
                if (1.5 >= -2.25) {
                    fmt.Println(fi38, fi38)
                    v53 := v3[2]
                    _ = v53
                } else {
                    fmt.Println(v2, v28, v5)
                }
            }
            if (v6 > v6) {
                v54 := GenS1{F0: -706}
                v55 := v4[fi38]
                _ = v54
                _ = v55
            }
            _ = fi38
            _ = v39
        }
        var v56 float64 = float64((float32(0.1) / ((v27 + (1e10 - (v28 + v28))) * v28)))
        _ = fi37
        _ = v56
    }
    v57 := GenS4{F0: v6, F1: -167}
    if true {
        var v58 uint64 = 0
        _ = v58
    }
    v5 = (^(v5))
    v28 = v27
    v27 = v28
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
