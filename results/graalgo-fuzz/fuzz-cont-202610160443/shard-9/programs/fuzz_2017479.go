package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 bool; F1 uint64 }
func genFn0() uint8 {
    var v7 bool = true
    v7 = (-1 < 0)
    if v7 {
        var v8 int16 = 32767
        var v9 bool = true
        _ = v8
        _ = v9
    }
    _ = v7
    return 127
}
func genFn1(p0 uint64, p1 uint, p2 int) uint16 {
    if true {
        fmt.Println(p2, p1)
        var v10 uint64 = (p0 - (p0 ^ (p0 << ((p1) & 63))))
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        p0 = 18446744073709551615
        switch "foo bar" {
        case "hello":
            switch p2 {
            case -1:
                p1 = p1
            default:
                var v12 int32 = -990
                _ = v12
            }
        }
        v13 := []int{(fi11 % ((p2) | 1)), p2}
        _ = fi11
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]uint32)
    var v3 float64 = (3.14159 / float64(1e10))
    v4 := []uint32{2147483647, 4294967295}
    v5 := make(map[string]bool)
    v6 := make(map[uint32]string)
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int = 885
            v1 = 42
            _ = fi15
            _ = v16
        }
        var v17 uint16 = 1
        if false {
            delete(v5, "")
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 int16 = 32767
                var v20 string = "foo bar"
                var v21 int8 = (v1 >> ((2) & 63))
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        _ = fi14
        _ = v17
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int8 = (v1 & (v1 << ((255) & 63)))
        v23 = (v1 - v23)
        v24 := GenS1{F0: -54}
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 int = fi25
        if ((float32((float64(0.1) - (v3 + (v3 * v3)))) < -2.25) && true) {
            v1 -= (v1 % ((-1) | 1))
            v4 = append(v4, 757)
            _ = len(v4)
        }
        if len(v4) > 0 {
            v4[0] = 42
        }
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: 4294967295, F1: 226}
    if false {
        if true {
            switch -851 {
            case 1:
                v4 = append(v4, 2147483648)
            case -2147483648:
                var v28 float32 = -2.25
                _ = v28
            case 394:
                _ = len(v6)
            }
            if (1 >= 182) {
                v1 -= -117
                for fi29 := 0; fi29 < 10; fi29++ {
                    fmt.Println(v1, fi29, fi29)
                    v5[""] = ("hello" >= ((("\n\t" + "hello") + "hello") + ((("a" + "a") + "a") + ((("\n\t" + "\n\t") + "hello") + "a"))))
                    v30 := GenS3{F0: true, F1: genFn1(4294967295, 18446744073709551615, (fi29 * (-(fi29)))), F2: 2}
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v1, v3)
            }
            v31 := GenS2{F0: 42, F1: 1}
            _ = v31
        }
        v32 := GenS4{F0: 1, F1: ("" + "a")}
        _ = v32
    }
    var v33 int64 = 1
    v2[32767] = 875
    switch 4294967295 {
    case -303:
        for fi34 := 0; fi34 < 6; fi34++ {
            v1 = (-(v1))
            var v35 int16 = -61
            var v36 int = 32768
            _ = fi34
            _ = v35
            _ = v36
        }
    case 390:
        v33 += (v33 + -192)
    case 32:
        var v37 uint64 = 971
        _ = v37
    default:
        switch "" {
        case "hello":
            for fi38 := 0; fi38 < 11; fi38++ {
                switch ("" + "foo bar") {
                case "foo bar":
                    var v39 int16 = 32767
                    _ = v39
                case "a":
                    v3 = v3
                default:
                    v40 := GenS5{F0: (!false), F1: 356}
                    _ = v40
                }
                v4 = append(v4, 4294967295)
                for fi41 := 0; fi41 < 3; fi41++ {
                    if ((v3 * (1e10 * float64(0.0))) != -2.25) {
                        fmt.Println(fi38, fi41)
                        var v42 float32 = 0.1
                        _ = v42
                    } else {
                        var v43 uint32 = 824
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v43
                    }
                    v33 = int64(v1)
                    _ = fi41
                }
                _ = fi38
            }
        case "foo bar":
            fmt.Println(v33, v3, v3)
        default:
            var v44 uint64 = 726
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
