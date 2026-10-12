package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int32; F2 uint64 }
type GenS6 struct { F0 int16; F1 uint16; F2 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string }
func genFn0(p0 int16, p1 uint16, p2 uint64) int {
    var v7 uint32 = 966
    switch 42 {
    case 151:
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := []int{(fi8 >> ((uint(p1)) & 63)), int(v7), (fi8 - (fi8 ^ fi8)), (fi8 + 32767)}
            for fi10 := 0; fi10 < 8; fi10++ {
                switch ("\n\t" + "\n\t") {
                case "":
                    v11 := GenS1{F0: 1}
                    _ = v11
                case "\n\t":
                    var v12 bool = true
                    _ = v12
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    case -1:
        switch ("a" + ("foo bar" + (("\n\t" + "\n\t") + "foo bar"))) {
        case "":
            var v13 uint8 = 2
            _ = v13
        default:
            v14 := GenS1{F0: -2147483648}
            _ = v14
        }
    case -691:
        var v15 int64 = 42
        _ = v15
    default:
        if (162 <= 128) {
            v16 := GenS1{F0: 42}
            var v17 bool = false
            _ = v16
            _ = v17
        } else {
            v18 := GenS2{F0: v7}
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483647
    v3 := []int{42, 32767}
    v4 := []uint32{v2, v2}
    var v5 uint = 0
    v6 := []string{}
    v19 := GenS3{F0: -1, F1: 1}
    var v20 string = "foo bar"
    fmt.Println(v5, v5)
    if (!((v2 | v2) >= (v2 * v2))) {
        v5 -= v5
        var v21 uint8 = uint8(v2)
        if false {
            var v22 string = v20
            switch v22 {
            case "a":
                fmt.Println(v20, v2, v20)
            case "foo bar":
                var v23 uint32 = 0
                _ = v23
            case "":
                v21 = 128
            default:
                if (!(!((float32(0.0) == (1.5 / ((-2.25 * (1.5 / 3.14159)) + 1e10))) && true))) {
                    switch v22 {
                    case "a":
                        if true {
                            v5 = v5
                            v22 = (v20 + v20)
                            fmt.Println(v20, v2, v21)
                        }
                    }
                    v1[("hello" + v20)] = genFn0(32767, 2, 829)
                    v24 := GenS3{F0: 0, F1: (v5 % (((^(v5))) | 1))}
                    _ = v24
                }
            }
            fmt.Println(v22, v22)
            _ = v22
        } else {
            v6 = append(v6, v20)
        }
        _ = v21
    }
    v25 := GenS3{F0: 0, F1: (v5 % (((v5 + 509)) | 1))}
    v2 += 2
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 uint64 = 407
        v28 := GenS4{F0: (v5 ^ (v5 / ((v5) | 1)))}
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v29 := GenS5{F0: 947, F1: -1, F2: (v27 ^ (v27 & (v27 % (((v27 | v27)) | 1))))}
            _ = v29
        } else {
            v30 := GenS2{F0: 2147483648}
            fmt.Println(v20)
            _ = v30
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    v31 := GenS6{F0: 42, F1: 255, F2: "a"}
    var v32 string = (v20 + v20)
    fmt.Println(v5)
    switch 2147483647 {
    case 9223372036854775807:
        var v33 bool = (true || true)
        _ = v33
    case -1:
        var v34 uint32 = 325
        _ = v34
    case 32767:
        if (-1 > 0) {
            for fi35 := 0; fi35 < 2; fi35++ {
                _ = len(v4)
                _ = fi35
            }
        } else {
            for fi36 := 0; fi36 < 13; fi36++ {
                var v37 float32 = 0.0
                var v38 int16 = -32768
                if true {
                    var v39 bool = false
                    _ = v39
                }
                _ = fi36
                _ = v37
                _ = v38
            }
        }
    }
    if (680 != 255) {
        if false {
            var v40 int = 127
            if len(v6) > 0 {
                v6[0] = ((v20 + "hello") + v20)
            }
            _ = v40
        } else {
            v32 = v32
        }
        v41 := v1[v20]
        _ = v41
    } else {
        v20 = (v32 + (v32 + v20))
        fmt.Println(v32, v5)
        if (-228 != 32768) {
            for fi42 := 0; fi42 < 16; fi42++ {
                switch fi42 {
                case 32767:
                    if false {
                        for fi43 := 0; fi43 < 4; fi43++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi43
                        }
                    }
                case 32768:
                    var v44 string = (v20 + (v32 + v32))
                    _ = v44
                case 74:
                    if (fi42 > fi42) {
                        var v45 int32 = 546
                        _ = v45
                    }
                default:
                    var v46 uint16 = 32767
                    _ = v46
                }
                _ = fi42
            }
            v47 := GenS7{F0: 32767}
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        delete(v1, "\n\t")
        for fi49 := 0; fi49 < 14; fi49++ {
            switch v32 {
            case "\n\t":
                v20 = ((v20 + (("" + "hello") + (v20 + "\n\t"))) + "a")
            case "hello":
                v50 := GenS8{F0: ((v20 + (v20 + v32)) + v32)}
                _ = v50
            }
            _ = fi49
        }
        _ = fi48
    }
    if true {
        switch 128 {
        case 2147483647:
            var v51 float64 = 3.14159
            _ = v51
        case -818:
            for fi52 := 0; fi52 < 18; fi52++ {
                v20 = "\n\t"
                v32 = ""
                if (!(false && ((118 < 128) || (0 > 61)))) {
                    for fi53 := 0; fi53 < 1; fi53++ {
                        var v54 uint8 = 253
                        v55 := GenS8{F0: v20}
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    _ = len(v6)
                    var v56 int = (fi52 / ((33) | 1))
                    _ = v56
                }
                _ = fi52
            }
        case 1:
            v57 := GenS4{F0: v5}
            _ = v57
        default:
            var v58 int64 = 9223372036854775807
            _ = v58
        }
    }
    if ((!(true || true)) || true) {
        var v59 int32 = -472
        _ = v59
    }
    var v60 uint64 = 4294967295
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
