package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 bool; F1 int64; F2 bool }
func genFn0() int16 {
    var v7 uint = 0
    v8 := GenS1{F0: 154, F1: 255, F2: (!(false || false))}
    var v9 uint8 = 127
    v10 := []int{}
    v7 = (v7 << (((v7 | (^(v7)))) & 63))
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 uint) int {
    v11 := GenS2{F0: 2147483647, F1: "hello", F2: 192}
    p1 += p1
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int32 = 691
        if false {
            v13 += 155
            p1 += (^(p1))
            for fi14 := 0; fi14 < 3; fi14++ {
                switch -1 {
                case 324:
                    v13 += -2147483648
                case 32767:
                    if ((fi14 != (fi12 << (((p1 * 18446744073709551615)) & 63))) || ((!false) || true)) {
                        if false {
                            v15 := []int{(fi14 % (((^(fi14))) | 1)), (fi14 - (fi14 << ((p1) & 63))), (-(fi12)), fi12}
                            v13 += (v13 % ((v13) | 1))
                            var v16 float32 = (0.1 + (3.14159 - (3.14159 - (0.0 / 1e10))))
                            var v17 string = "foo bar"
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        p0 += 42
                        var v18 float64 = 0.1
                        _ = v18
                    } else {
                        var v19 bool = (-9223372036854775808 > 0)
                        _ = v19
                    }
                case 0:
                    var v20 float64 = 1e10
                    _ = v20
                }
                var v21 uint8 = 228
                _ = fi14
                _ = v21
            }
        }
        _ = fi12
        _ = v13
    }
    var v22 float64 = (1e10 * (float64((3.14159 + 0.1)) + 1e10))
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := GenS1{F0: 0, F1: 255, F2: ((0 != 0) || ("hello" < "foo bar"))}
        p0 = p0
        _ = fi23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v22
    return 4294967295
}

func main() {
    var v1 float32 = (float32(((3.14159 + 0.0) + 0.0)) * 0.0)
    v2 := []string{"hello", "a", "hello"}
    var v3 int = 261
    v4 := make(map[int]bool)
    var v5 int32 = 1
    v6 := make(map[int]uint32)
    if false {
        fmt.Println(v3, v3, v3)
    } else {
        var v25 uint8 = 255
        var v26 uint64 = 0
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v29 float64 = 1e10
            fmt.Println(fi28, fi27, fi28)
            if (true || true) {
                var v30 uint16 = 42
                switch (-(fi27)) {
                case -9223372036854775808:
                    var v31 uint32 = 1
                    _ = v31
                case -215:
                    delete(v4, (fi28 & genFn1(4294967295, 2)))
                case 42:
                    var v32 uint32 = 65535
                    if true {
                        v3 -= fi27
                        v33 := GenS3{F0: v30, F1: 247, F2: (v32 * uint32(v30))}
                        _ = v33
                    } else {
                        var v34 bool = true
                        _ = v34
                    }
                    _ = v32
                default:
                    v1 -= (v1 + 1.5)
                }
                _ = v30
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v35 uint8 = 48
    v3 += (-(v3))
    if false {
        v5 += -886
    } else {
        v5 = v5
        if false {
            for fi36 := 0; fi36 < 17; fi36++ {
                v1 = v1
                if true {
                    v4[fi36] = false
                    delete(v6, fi36)
                } else {
                    if ((v5 << ((65535) & 63)) > 1) {
                        var v37 uint32 = 2147483648
                        _ = v37
                    } else {
                        if true {
                            _ = len(v6)
                        } else {
                            var v38 int = -1
                            if (!false) {
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                v3 = (fi36 + (v38 >> ((611) & 63)))
                                switch "\n\t" {
                                case "foo bar":
                                    v39 := v4[(v3 + (v3 % ((-567) | 1)))]
                                    _ = v39
                                }
                            } else {
                                fmt.Println(v5, v1, v35)
                                switch "\n\t" {
                                case "":
                                    v40 := v4[genFn1(41, 4294967295)]
                                    _ = v40
                                case "\n\t":
                                    for fi41 := 0; fi41 < 17; fi41++ {
                                        var v42 uint16 = 226
                                        _ = fi41
                                        _ = v42
                                    }
                                case "foo bar":
                                    if (1 < 1) {
                                        v43 := GenS1{F0: 1, F1: 981, F2: (2 < 0)}
                                        v35 = v35
                                        if (false || true) {
                                            var v44 uint8 = v35
                                            _ = v44
                                        } else {
                                            var v45 float32 = v1
                                            _ = v45
                                        }
                                        _ = v43
                                    } else {
                                        v5 -= -2147483648
                                        for fi46 := 0; fi46 < 10; fi46++ {
                                            var v47 float32 = float32(1.5)
                                            _ = fi46
                                            _ = v47
                                        }
                                    }
                                default:
                                    v3 -= (-(v3))
                                }
                            }
                            _ = v38
                        }
                        var v48 int = (fi36 - (fi36 % ((genFn1(410, 255)) | 1)))
                        _ = v48
                    }
                }
                _ = fi36
            }
        } else {
            v3 = genFn1(0, 255)
            for fi49 := 0; fi49 < 15; fi49++ {
                if true {
                    var v50 uint16 = 0
                    _ = v50
                }
                switch fi49 {
                case -774:
                    for fi51 := 0; fi51 < 4; fi51++ {
                        fmt.Println(v1, fi51, v3)
                        _ = fi51
                    }
                case 0:
                    v52 := GenS4{F0: -32768, F1: (v35 / ((42) | 1))}
                    _ = v52
                default:
                    var v53 float64 = float64(0.1)
                    _ = v53
                }
                v1 = 1.5
                _ = fi49
            }
            fmt.Println(v3, v3, v1)
        }
    }
    switch "hello" {
    case "":
        _ = len(v2)
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    default:
        v1 += (((v1 / v1) / v1) / (1e10 - 1.5))
    }
    if ((v5 % (((v5 << ((255) & 63))) | 1)) < v5) {
        v35 += 0
        v1 += (v1 / (v1 + float32(1.5)))
    }
    v4[v3] = ((933 <= 42) || true)
    fmt.Println(v3, v3)
    v54 := GenS4{F0: 32767, F1: (v35 - v35)}
    v35 += 82
    for fi55 := 0; fi55 < 4; fi55++ {
        var v56 int32 = -1
        v57 := v6[v3]
        v1 = ((v1 * (float32(3.14159) + v1)) * (0.1 - v1))
        _ = fi55
        _ = v56
        _ = v57
    }
    var v58 int = v3
    v59 := GenS2{F0: (-(v5)), F1: "a", F2: 647}
    v60 := GenS5{F0: (42 != -146), F1: 9223372036854775807, F2: (-1 != 0)}
    v1 = v1
    var v61 int = 365
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v54
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
