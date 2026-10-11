package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint32) int {
    v7 := []int{-946}
    var v8 uint32 = (p1 + 0)
    v8 -= p1
    v9 := GenS1{F0: 186}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 949
}
func genFn1(p0 int8) uint {
    v10 := GenS1{F0: -1}
    var v11 int32 = -104
    var v12 bool = ((!(!true)) || (255 == 65535))
    v13 := GenS1{F0: 1}
    var v14 string = "foo bar"
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 2147483648
    var v3 int32 = -770
    v4 := []int{2147483647, 42}
    var v5 float32 = 0.1
    v6 := []int{42, -455, 1}
    for fi15 := 0; fi15 < 20; fi15++ {
        _ = len(v4)
        _ = fi15
    }
    fmt.Println(v1, v2, v5)
    fmt.Println(v2)
    var v16 uint16 = (v1 << ((4294967295) & 63))
    switch -363 {
    case 127:
        if (92 > 186) {
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 float32 = 1e10
                for fi19 := 0; fi19 < 4; fi19++ {
                    fmt.Println(v2, v3)
                    v16 = (v16 >> ((0) & 63))
                    _ = fi19
                }
                if (127 != 1) {
                    if true {
                        switch ("" + "\n\t") {
                        case "hello":
                            v20 := GenS1{F0: fi17}
                            _ = v20
                        case "foo bar":
                            if false {
                                for fi21 := 0; fi21 < 4; fi21++ {
                                    var v22 uint8 = 127
                                    _ = fi21
                                    _ = v22
                                }
                            } else {
                                var v23 int = (-(fi17))
                                if (false || (!false)) {
                                    v24 := GenS1{F0: (^(v23))}
                                    _ = v24
                                }
                                v23 -= (v23 ^ genFn0(74, (v2 / ((v2) | 1))))
                                _ = v23
                            }
                        case "\n\t":
                            var v25 uint32 = v2
                            _ = v25
                        }
                        var v26 uint32 = (v2 ^ (^(v2)))
                        _ = v26
                    }
                } else {
                    var v27 int8 = 114
                    v28 := GenS1{F0: -2147483648}
                    _ = v27
                    _ = v28
                }
                _ = fi17
                _ = v18
            }
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 float32 = v5
                var v31 int8 = -128
                var v32 int16 = -32768
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            v3 = 42
        } else {
            if false {
                fmt.Println(v2, v5)
                for fi33 := 0; fi33 < 14; fi33++ {
                    var v34 int64 = 42
                    _ = fi33
                    _ = v34
                }
                fmt.Println(v3, v2, v3)
            }
            var v35 uint64 = 4294967295
            _ = v35
        }
    case 32767:
        switch "" {
        case "hello":
            var v36 int8 = -128
            _ = v36
        default:
            for fi37 := 0; fi37 < 1; fi37++ {
                switch "\n\t" {
                case "foo bar":
                    var v38 bool = true
                    _ = v38
                case "":
                    var v39 uint8 = 53
                    _ = v39
                default:
                    if true {
                        if false {
                            var v40 int = genFn0(-1, 2147483647)
                            if false {
                                for fi41 := 0; fi41 < 10; fi41++ {
                                    if (float64((v5 - (((v5 / (v5 + float32(float64(v5)))) + (v5 + 0.1)) + v5))) > (1.5 - 0.1)) {
                                        var v42 uint64 = 911
                                        if ((v42 & (^(v42))) != v42) {
                                            v1 -= (v1 * 1)
                                        }
                                        v2 += (v2 * uint32(v16))
                                        _ = v42
                                    }
                                    var v43 bool = true
                                    var v44 int = (fi41 | 42)
                                    _ = fi41
                                    _ = v43
                                    _ = v44
                                }
                                v6 = append(v6, (^(v40)))
                                if (true && false) {
                                    v45 := GenS1{F0: genFn0(int8(v1), (v2 | v2))}
                                    fmt.Println(v2, v2)
                                    _ = v45
                                } else {
                                    v46 := GenS1{F0: (fi37 % (((fi37 | (fi37 | (v40 ^ genFn0(127, v2))))) | 1))}
                                    v40 = (v40 >> ((uint(v1)) & 63))
                                    _ = v46
                                }
                            }
                            if ("foo bar" < "foo bar") {
                                v4 = append(v4, (fi37 % ((genFn0(-128, (v2 ^ (v2 + v2)))) | 1)))
                                var v47 float32 = ((((v5 - v5) - v5) / 1e10) * 3.14159)
                                _ = v47
                            }
                            _ = v40
                        }
                        v6 = append(v6, (fi37 + (fi37 << ((18446744073709551615) & 63))))
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                }
                for fi48 := 0; fi48 < 3; fi48++ {
                    var v49 uint8 = 98
                    v50 := GenS2{F0: "foo bar", F1: -55, F2: (fi37 ^ genFn0(-1, uint32(v49)))}
                    var v51 uint = 42
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = fi37
            }
        }
    case -2147483648:
        for fi52 := 0; fi52 < 3; fi52++ {
            if (v5 == (float32(float64(0.0)) / (float32(float64(1.5)) - v5))) {
                v6 = append(v6, (fi52 << ((genFn1(100)) & 63)))
                switch "a" {
                case "\n\t":
                    fmt.Println(v2, v2, v1)
                case "":
                    _ = len(v6)
                }
                if (v5 != v5) {
                    var v53 int16 = 1000
                    v54 := GenS1{F0: -1}
                    if (4294967295 <= 18446744073709551615) {
                        v53 = v53
                        v6 = append(v6, (fi52 - (fi52 - (fi52 & 128))))
                        v2 -= v2
                    }
                    _ = v53
                    _ = v54
                } else {
                    v55 := GenS1{F0: (fi52 ^ int(v2))}
                    if true {
                        v56 := GenS3{F0: -1}
                        switch "\n\t" {
                        case "":
                            var v57 int32 = (^(v3))
                            _ = v57
                        case "hello":
                            v58 := GenS4{F0: 128}
                            _ = v58
                        default:
                            var v59 uint = 42
                            _ = v59
                        }
                        _ = v56
                    }
                    _ = v55
                }
            } else {
                var v60 uint32 = (v2 + (v2 + v2))
                var v61 bool = (128 != 106)
                for fi62 := 0; fi62 < 14; fi62++ {
                    var v63 int32 = v3
                    var v64 string = "\n\t"
                    _ = fi62
                    _ = v63
                    _ = v64
                }
                _ = v60
                _ = v61
            }
            fmt.Println(v2)
            var v65 uint8 = 127
            _ = fi52
            _ = v65
        }
    }
    var v66 bool = false
    v5 = v5
    v67 := GenS4{F0: 228}
    fmt.Println(v2, v2)
    var v68 int32 = (v3 << ((65535) & 63))
    v2 -= 2147483648
    var v69 int64 = -9223372036854775808
    var v70 int64 = (v69 ^ v69)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
