package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0() int32 {
    var v7 bool = ((true || false) || (false && (42 == 57)))
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: ((fi8 & (fi8 / (((fi8 / ((fi8) | 1))) | 1))) <= 125), F1: "\n\t"}
        _ = fi8
        _ = v9
    }
    v7 = ((0 <= -2147483648) && (20 <= 32767))
    v10 := []int{127}
    if len(v10) > 0 {
        v10[0] = 0
    }
    _ = v10
    _ = v7
    return 342
}
func genFn1(p0 uint8, p1 uint32) int32 {
    var v11 float32 = (-2.25 / float32(0.1))
    if ((p0 >= (p0 << ((42) & 63))) || true) {
        if true {
            v12 := GenS1{F0: (-494 < 100), F1: ("foo bar" + "\n\t")}
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := []int{(fi13 | (fi13 & (fi13 / ((fi13) | 1)))), fi13}
                var v15 float32 = (0.1 * float32(-2.25))
                if (false && ((((p0 & p0) >= 128) && (930 >= p1)) || (true || false))) {
                    if ((false || ((("hello" + "") + ("\n\t" + "")) > ("" + (("hello" + ("a" + "\n\t")) + "\n\t")))) || false) {
                        var v16 uint64 = 669
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v14 = append(v14, fi13)
                            _ = fi17
                        }
                        v14 = append(v14, fi13)
                    }
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: (-77 > 42), F1: "\n\t"}
        v20 := GenS2{F0: p0}
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v11
    return 42
}

func main() {
    var v1 int64 = 0
    var v2 bool = (3.14159 <= 0.1)
    v3 := make(map[int]int)
    var v4 uint32 = 783
    var v5 float32 = 1.5
    var v6 string = (("" + "hello") + "\n\t")
    v21 := GenS2{F0: 0}
    if v2 {
        if v2 {
            if (-1 <= -55) {
                var v22 uint8 = 42
                _ = v22
            } else {
                var v23 uint8 = 205
                for fi24 := 0; fi24 < 4; fi24++ {
                    v25 := []int{}
                    _ = fi24
                    _ = v25
                }
                _ = v23
            }
            var v26 uint16 = 903
            _ = v26
        }
        if v2 {
            var v27 uint = 621
            v28 := []int{}
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 2; fi29++ {
                v6 = "foo bar"
                fmt.Println(fi29, v5)
                fmt.Println(v4)
                _ = fi29
            }
            if (541 != 0) {
                switch 0 {
                case 127:
                    v30 := GenS3{F0: 128, F1: (!v2)}
                    _ = v30
                default:
                    if ((-32768 >= -925) || (!v2)) {
                        v6 = v6
                        var v31 uint64 = 0
                        for fi32 := 0; fi32 < 13; fi32++ {
                            for fi33 := 0; fi33 < 3; fi33++ {
                                var v34 int64 = (v1 - -1)
                                _ = fi33
                                _ = v34
                            }
                            v35 := []int{fi32, int(v31), fi32}
                            if true {
                                v5 -= 3.14159
                                v36 := GenS2{F0: 116}
                                var v37 int = 42
                                _ = v36
                                _ = v37
                            } else {
                                v2 = true
                            }
                            _ = fi32
                            _ = v35
                        }
                        _ = v31
                    } else {
                        var v38 uint16 = 370
                        if false {
                            v39 := GenS4{F0: 127}
                            if v2 {
                                var v40 float32 = float32(-2.25)
                                fmt.Println(v1, v38)
                                _ = v40
                            }
                            fmt.Println(v2, v4)
                            _ = v39
                        }
                        _ = v38
                    }
                }
            } else {
                v41 := GenS4{F0: 32768}
                v42 := GenS2{F0: 197}
                v5 = v5
                _ = v41
                _ = v42
            }
            var v43 uint8 = 2
            _ = v43
        }
    } else {
        v4 -= 2
        v3[-2147483648] = 128
        if (!v2) {
            for fi44 := 0; fi44 < 4; fi44++ {
                switch (v6 + v6) {
                case "":
                    var v45 uint = 0
                    _ = v45
                default:
                    v1 = (^(v1))
                }
                _ = fi44
            }
            delete(v3, 128)
        } else {
            for fi46 := 0; fi46 < 15; fi46++ {
                var v47 bool = false
                switch "a" {
                case "a":
                    v5 = 0.1
                case "hello":
                    switch fi46 {
                    case -2147483648:
                        _ = len(v3)
                    case -1:
                        v48 := GenS4{F0: fi46}
                        _ = v48
                    case -348:
                        var v49 bool = ((2 > 127) || (false || true))
                        _ = v49
                    }
                }
                _ = fi46
                _ = v47
            }
        }
    }
    v6 = "a"
    var v50 string = ((("" + v6) + ("hello" + v6)) + "")
    var v51 int64 = (v1 / ((604) | 1))
    var v52 float64 = ((1e10 + -2.25) + (float64((3.14159 - v5)) * 0.0))
    v5 += 0.1
    for fi53 := 0; fi53 < 3; fi53++ {
        for fi54 := 0; fi54 < 15; fi54++ {
            v55 := GenS2{F0: 45}
            switch fi53 {
            case 790:
                var v56 int16 = 1000
                _ = v56
            case -201:
                if v2 {
                    v51 += -390
                    switch 522 {
                    case 32768:
                        var v57 bool = (v2 || v2)
                        _ = v57
                    case 498:
                        v1 += -1
                    }
                    if (127 < uint8(v4)) {
                        v1 = 174
                        v6 = (v6 + "")
                        v58 := GenS4{F0: fi54}
                        _ = v58
                    }
                }
            case 406:
                v59 := GenS1{F0: v2, F1: ("" + "hello")}
                _ = v59
            default:
                switch fi54 {
                case 128:
                    v50 = v6
                }
            }
            var v60 int64 = -1
            switch (v6 + "") {
            case "hello":
                var v61 uint64 = 42
                _ = v61
            case "\n\t":
                v62 := []int{}
                _ = v62
            case "foo bar":
                var v63 int64 = v51
                fmt.Println(v6, v4, v5)
                var v64 int = 127
                _ = v63
                _ = v64
            }
            v51 -= v51
            fmt.Println(v2)
            _ = fi54
            _ = v55
            _ = v60
        }
        _ = fi53
    }
    fmt.Println(v5)
    if (-75 != 100) {
        for fi65 := 0; fi65 < 16; fi65++ {
            var v66 uint32 = 0
            v67 := GenS5{F0: ((v2 || true) || true), F1: 42, F2: 65535}
            v68 := v3[fi65]
            _ = fi65
            _ = v66
            _ = v67
            _ = v68
        }
    }
    v51 -= (-(v51))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
