package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 string; F2 int8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint32; F2 bool }
func genFn0() uint8 {
    var v7 uint16 = 1
    if true {
        var v8 string = "\n\t"
        v9 := GenS1{F0: 0}
        _ = v8
        _ = v9
    }
    _ = v7
    return 120
}
func genFn1(p0 int32, p1 uint64, p2 int32) uint {
    if (!(535 > 432)) {
        v10 := []int{42, -45}
        switch "hello" {
        case "foo bar":
            var v11 uint64 = 2
            _ = v11
        case "":
            for fi12 := 0; fi12 < 17; fi12++ {
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi13
                }
                var v14 uint32 = 2
                _ = fi12
                _ = v14
            }
        case "\n\t":
            if true {
                switch ("\n\t" + "") {
                case "foo bar":
                    v10 = append(v10, 0)
                default:
                    v10 = append(v10, 9223372036854775807)
                }
                var v15 uint8 = genFn0()
                p0 = p2
                _ = v15
            }
        }
        if false {
            v10 = append(v10, 32768)
            if false {
                v16 := GenS1{F0: 18446744073709551615}
                _ = v16
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 int = 4294967295
                    _ = fi17
                    _ = v18
                }
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := GenS1{F0: (p1 * p1)}
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := GenS1{F0: (p1 >> ((19) & 63))}
                    _ = fi21
                    _ = v22
                }
            }
        } else {
            var v23 int16 = -269
            var v24 uint64 = (^(p1))
            _ = v23
            _ = v24
        }
        _ = v10
    }
    var v25 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return uint(p1)
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    v3 := []string{"foo bar"}
    var v4 float64 = 0.1
    v5 := []uint32{}
    v6 := []int{1, -9223372036854775808, 9223372036854775807, 32768}
    v1 = v1
    if true {
        _ = len(v6)
    } else {
        var v26 uint32 = 4294967295
        if (!false) {
            for fi27 := 0; fi27 < 17; fi27++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi27
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS1{F0: 1}
                fmt.Println(v1, v26)
                _ = fi28
                _ = v29
            }
            var v30 string = v1
            _ = v30
        } else {
            var v31 string = "hello"
            var v32 int32 = 205
            _ = v31
            _ = v32
        }
        _ = v26
    }
    if false {
        var v33 uint32 = 866
        fmt.Println(v33, v33)
        _ = v33
    } else {
        fmt.Println(v4, v4)
        fmt.Println(v1)
    }
    var v34 int8 = 8
    v35 := GenS1{F0: 2}
    var v36 string = (v1 + v1)
    var v37 uint8 = 18
    var v38 int8 = v34
    for fi39 := 0; fi39 < 6; fi39++ {
        for fi40 := 0; fi40 < 8; fi40++ {
            v41 := GenS2{F0: v1, F1: v1, F2: v38}
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            fmt.Println(fi42)
            if ((^(v38)) <= (v38 * v38)) {
                fmt.Println(v36, v34)
                var v43 int16 = -242
                _ = v43
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi42
        }
        _ = fi39
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        var v45 int64 = 1
        for fi46 := 0; fi46 < 8; fi46++ {
            v47 := v2[479]
            v37 += v37
            _ = fi46
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    var v48 int16 = -1
    var v49 uint = 255
    var v50 string = (((v1 + "a") + v1) + "\n\t")
    for fi51 := 0; fi51 < 17; fi51++ {
        var v52 uint64 = 1
        v50 = ""
        v53 := GenS3{F0: -1, F1: 1}
        _ = fi51
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 11; fi54++ {
        var v55 uint16 = 0
        var v56 uint8 = v37
        _ = fi54
        _ = v55
        _ = v56
    }
    var v57 bool = (1 > 42)
    for fi58 := 0; fi58 < 20; fi58++ {
        if v57 {
            v59 := v2[4294967295]
            if (0.1 <= (-2.25 * ((1e10 * 3.14159) + (float32(((1e10 + (0.1 / v4)) - v4)) * 0.0)))) {
                if true {
                    var v60 string = ((((("\n\t" + v50) + "") + v36) + v36) + (v1 + ((v50 + (v36 + v50)) + "foo bar")))
                    _ = v60
                }
                _ = len(v3)
                v61 := GenS1{F0: 4294967295}
                _ = v61
            }
            _ = v59
        } else {
            if v57 {
                switch fi58 {
                case -1:
                    fmt.Println(v37, v34, v1)
                case 32767:
                    for fi62 := 0; fi62 < 18; fi62++ {
                        var v63 float32 = 1e10
                        v38 = -29
                        _ = fi62
                        _ = v63
                    }
                case 1:
                    v64 := GenS4{F0: genFn1(-383, 558, 0), F1: 1, F2: (!((-(v38)) != -1))}
                    _ = v64
                }
            } else {
                for fi65 := 0; fi65 < 15; fi65++ {
                    v50 = (v1 + v1)
                    switch v50 {
                    case "foo bar":
                        delete(v2, 258)
                    case "\n\t":
                        var v66 uint32 = 2
                        _ = v66
                    case "a":
                        var v67 float32 = 3.14159
                        _ = v67
                    }
                    _ = fi65
                }
            }
            v57 = (("a" > ("\n\t" + "a")) && v57)
            v68 := v2[4294967295]
            _ = v68
        }
        var v69 bool = v57
        _ = fi58
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v57
    _ = v6
}
