package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int64, p2 uint8) int32 {
    p1 = (p1 & (^(p1)))
    if ((false || (!(0 > 2147483647))) && false) {
        v7 := GenS1{F0: (p0 | p0)}
        if ((p2 + 127) != p0) {
            p2 = p2
        }
        var v8 int = 32767
        _ = v7
        _ = v8
    } else {
        switch "hello" {
        case "":
            v9 := GenS2{F0: -115}
            _ = v9
        case "foo bar":
            v10 := GenS3{F0: p2, F1: 1}
            _ = v10
        case "a":
            for fi11 := 0; fi11 < 12; fi11++ {
                p0 = p0
                p0 = (p2 + (^(p0)))
                _ = fi11
            }
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 bool = ((fi12 + 32768) <= (fi12 - -1))
                _ = fi12
                _ = v13
            }
        }
    }
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}
func genFn1(p0 uint16, p1 uint, p2 uint32) int8 {
    var v14 uint8 = 93
    var v15 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return v15
}

func main() {
    var v1 float64 = (3.14159 / 1.0)
    var v2 int64 = -9223372036854775808
    var v3 int32 = 417
    var v4 int8 = 42
    var v5 uint16 = 65535
    var v6 int8 = int8(v3)
    switch 240 {
    case 127:
        v1 = (1.5 * v1)
    }
    if (0 > 65535) {
        fmt.Println(v5)
        if (65535 != 0) {
            switch -262 {
            case 1:
                var v16 int64 = 448
                _ = v16
            case 128:
                v1 = v1
            case -9223372036854775808:
                for fi17 := 0; fi17 < 12; fi17++ {
                    if ("a" < "hello") {
                        var v18 int = 2147483647
                        v19 := GenS1{F0: 1}
                        v4 = v4
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                v21 := []int{2147483647, (fi20 - (-(fi20))), fi20, fi20}
                fmt.Println(v4, v2)
                _ = fi20
                _ = v21
            }
        } else {
            v2 = (-(v2))
            v22 := GenS2{F0: (v6 - (v6 << ((733) & 63)))}
            _ = v22
        }
        v23 := []int{9223372036854775807, 729, -668}
        _ = v23
    } else {
        switch 0 {
        case 865:
            v24 := []int{32767}
            _ = v24
        case -26:
            for fi25 := 0; fi25 < 15; fi25++ {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS2{F0: (v6 >> ((65535) & 63))}
                    _ = fi26
                    _ = v27
                }
                v6 += (v6 + (v4 - v6))
                var v28 int32 = (-(v3))
                _ = fi25
                _ = v28
            }
        case 127:
            switch 1 {
            case 128:
                switch 128 {
                case 127:
                    for fi29 := 0; fi29 < 20; fi29++ {
                        v30 := GenS4{F0: -161, F1: 0}
                        switch fi29 {
                        case -892:
                            var v31 bool = (false && ("\n\t" >= ("foo bar" + "\n\t")))
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                case 42:
                    var v32 float32 = -2.25
                    _ = v32
                case 4294967295:
                    var v33 int32 = (v3 + (-(v3)))
                    _ = v33
                }
            }
        default:
            var v34 int = int(v2)
            _ = v34
        }
        var v35 float32 = (0.0 * 0.0)
        _ = v35
    }
    var v36 int16 = 1000
    if ((v36 % ((0) | 1)) > v36) {
        v4 += (v6 >> ((42) & 63))
        for fi37 := 0; fi37 < 5; fi37++ {
            for fi38 := 0; fi38 < 16; fi38++ {
                v39 := GenS4{F0: (v36 >> ((463) & 63)), F1: 1}
                var v40 int32 = (v3 | v3)
                v41 := []int{fi38}
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 2; fi42++ {
                v36 -= (v36 ^ (v36 & v36))
                _ = fi42
            }
            var v43 int32 = (-(v3))
            _ = fi37
            _ = v43
        }
    }
    var v44 uint = 939
    v2 = (v2 * v2)
    v45 := []int{}
    v45 = append(v45, -2147483648)
    for fi46 := 0; fi46 < 6; fi46++ {
        if false {
            if true {
                v6 -= v4
                switch fi46 {
                case 42:
                    v47 := GenS5{F0: 0, F1: 4294967295, F2: 77}
                    _ = v47
                case 2147483647:
                    v3 = (v3 - (v3 % ((v3) | 1)))
                }
                for fi48 := 0; fi48 < 1; fi48++ {
                    v44 = v44
                    if len(v45) >= 2 {
                        v45 = v45[1:len(v45)]
                    }
                    v2 += v2
                    if len(v45) >= 2 {
                        v45 = v45[1:len(v45)]
                    }
                    v44 -= v44
                    _ = fi48
                }
            } else {
                var v49 uint32 = 2147483647
                _ = v49
            }
            var v50 float64 = v1
            _ = v50
        }
        _ = fi46
    }
    if (((v5 | v5) != (^(v5))) && ((^(v2)) < v2)) {
        switch "a" {
        case "foo bar":
            switch ("" + (("a" + (("a" + "") + (("" + "hello") + "foo bar"))) + (("" + ("a" + ("" + ""))) + ("" + "foo bar")))) {
            case "":
                v51 := GenS3{F0: 55, F1: 261}
                _ = v51
            case "\n\t":
                v3 = (v3 % ((895) | 1))
            case "a":
                var v52 int64 = (v2 / ((v2) | 1))
                _ = v52
            }
        default:
            var v53 bool = true
            _ = v53
        }
        for fi54 := 0; fi54 < 8; fi54++ {
            v55 := GenS6{F0: 860}
            var v56 int64 = v2
            _ = fi54
            _ = v55
            _ = v56
        }
    } else {
        fmt.Println(v36, v2, v44)
        v44 += v44
        var v57 int8 = v4
        _ = v57
    }
    v2 += v2
    fmt.Println(v5, v3, v36)
    for fi58 := 0; fi58 < 14; fi58++ {
        var v59 float64 = v1
        var v60 int16 = (v36 / (((^(v36))) | 1))
        for fi61 := 0; fi61 < 4; fi61++ {
            _ = len(v45)
            if len(v45) >= 2 {
                v45 = v45[1:len(v45)]
            }
            if len(v45) > 0 {
                v45[0] = (^(fi58))
            }
            _ = fi61
        }
        _ = fi58
        _ = v59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
