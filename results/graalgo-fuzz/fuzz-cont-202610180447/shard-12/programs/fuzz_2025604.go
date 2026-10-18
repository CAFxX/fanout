package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int8 }
func genFn0() uint16 {
    var v7 int64 = 866
    if (true && true) {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{(fi8 & -881), (fi8 / ((4294967295) | 1)), 32767}
            var v10 uint16 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint64 = 4294967295
            var v13 int8 = -112
            switch 9223372036854775807 {
            case 9223372036854775807:
                if true {
                    var v14 uint32 = 568
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := []int{}
                        _ = fi15
                        _ = v16
                    }
                    fmt.Println(v7)
                    _ = v14
                } else {
                    v17 := GenS1{F0: 117, F1: 2}
                    _ = v17
                }
            case 42:
                v18 := []int{}
                _ = v18
            case -1:
                v19 := []int{(fi11 % ((685) | 1))}
                _ = v19
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        v7 -= v7
    }
    var v20 float64 = (0.0 - 1e10)
    if true {
        var v21 float32 = 1.5
        _ = v21
    } else {
        switch 717 {
        case 32767:
            v20 = 3.14159
        case 807:
            var v22 int32 = 42
            _ = v22
        default:
            v23 := []int{629, 9223372036854775807, -587}
            _ = v23
        }
    }
    _ = v20
    _ = v7
    return 539
}
func genFn1() int {
    switch "\n\t" {
    case "hello":
        v24 := GenS1{F0: 0, F1: 42}
        _ = v24
    case "":
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(fi25)
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint8 = 0
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    var v28 bool = ((255 == 0) || false)
    var v29 uint = 928
    _ = v28
    _ = v29
    return -763
}

func main() {
    var v1 int16 = -1
    var v2 uint64 = 0
    v3 := make(map[uint32]uint32)
    var v4 uint = 255
    v5 := make(map[int]uint32)
    var v6 uint64 = 453
    if (2 < (v6 % ((0) | 1))) {
        var v30 string = ((("a" + "foo bar") + "") + "hello")
        var v31 int32 = -392
        _ = v30
        _ = v31
    } else {
        fmt.Println(v4, v6, v1)
    }
    v1 = v1
    if (!true) {
        var v32 float64 = (1.5 / ((1.5 - float64(0.0)) - 0.1))
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 int32 = 42
            v32 += (float64(1.5) - v32)
            _ = fi33
            _ = v34
        }
        var v35 int16 = 0
        _ = v32
        _ = v35
    }
    var v36 uint = 280
    v37 := GenS2{F0: 1, F1: -171, F2: 745}
    v36 = v4
    switch -680 {
    case 773:
        var v38 float64 = (0.0 / ((1.5 / 1.0) - 0.1))
        _ = v38
    case -890:
        v4 += 18446744073709551615
    case 4294967295:
        var v39 bool = false
        _ = v39
    default:
        delete(v3, 539)
    }
    fmt.Println(v36, v1)
    switch "foo bar" {
    case "":
        switch 32767 {
        case 9223372036854775807:
            if (int8(v6) < 42) {
                if ("" > ("a" + "")) {
                    v3[4294967295] = 0
                    var v40 float32 = float32(0.0)
                    _ = v40
                }
                v41 := GenS3{F0: 0, F1: ""}
                for fi42 := 0; fi42 < 6; fi42++ {
                    if true {
                        var v43 uint8 = 127
                        v44 := GenS4{F0: v1}
                        v2 = 42
                        _ = v43
                        _ = v44
                    } else {
                        delete(v3, 2147483647)
                        v4 += 644
                        v6 = 814
                    }
                    _ = fi42
                }
                _ = v41
            } else {
                v2 = v6
            }
        case 0:
            var v45 int32 = 0
            _ = v45
        case 214:
            v46 := []int{genFn1()}
            _ = v46
        default:
            if true {
                v47 := GenS5{F0: 42, F1: 675, F2: 1}
                _ = v47
            }
        }
    case "foo bar":
        var v48 uint64 = v6
        _ = v48
    default:
        var v49 uint = (v4 | (v36 << ((v4) & 63)))
        _ = v49
    }
    if (!(383 != -759)) {
        if (254 > 1) {
            v50 := []int{}
            switch ("" + ("foo bar" + ("a" + ("hello" + (("foo bar" + ((("hello" + "a") + "foo bar") + ("a" + "foo bar"))) + "hello"))))) {
            case "":
                if (32767 <= 2) {
                    var v51 bool = (!(200 <= uint8(v2)))
                    v4 = 42
                    v52 := GenS1{F0: 128, F1: 255}
                    _ = v51
                    _ = v52
                }
            case "\n\t":
                v53 := GenS2{F0: 0, F1: -1, F2: 656}
                _ = v53
            }
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 uint8 = 127
                _ = fi54
                _ = v55
            }
            _ = v50
        } else {
            v56 := v3[2147483648]
            for fi57 := 0; fi57 < 5; fi57++ {
                v2 = v6
                var v58 int64 = -811
                if (((0.1 / -2.25) * ((1.5 - -2.25) * 0.0)) == 3.14159) {
                    var v59 int8 = int8(v6)
                    if true {
                        var v60 int64 = -974
                        _ = v60
                    }
                    v4 += (v4 / (((v4 ^ (v4 >> (((v36 - (v36 << ((v4) & 63)))) & 63)))) | 1))
                    _ = v59
                }
                _ = fi57
                _ = v58
            }
            _ = v56
        }
        delete(v3, 2147483648)
        v61 := []int{127}
        _ = v61
    } else {
        v62 := []int{32768, -770}
        v63 := GenS6{F0: 255, F1: 1, F2: 0}
        _ = v62
        _ = v63
    }
    v3[2147483647] = 0
    switch genFn1() {
    case 127:
        v1 += 1
    case 0:
        var v64 int64 = 9223372036854775807
        _ = v64
    case -102:
        if (uint16(v2) <= 1) {
            v2 -= v2
            v36 -= (v36 - 311)
        }
    default:
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
