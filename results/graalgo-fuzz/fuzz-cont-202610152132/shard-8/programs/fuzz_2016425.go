package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 bool; F1 string; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint; F1 string; F2 uint16 }
type GenS7 struct { F0 int8 }
func genFn0() int8 {
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 11; fi7++ {
        switch ("hello" + "foo bar") {
        case "hello":
            v8 := GenS1{F0: 13}
            _ = v8
        case "":
            v9 := []int{}
            _ = v9
        default:
            v10 := GenS1{F0: 184}
            _ = v10
        }
        _ = fi7
    }
    if false {
        var v11 uint16 = v6
        _ = v11
    } else {
        var v12 string = "a"
        var v13 uint8 = 0
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 65535, F1: -333, F2: 2}
    _ = v14
    _ = v6
    return -128
}
func genFn1(p0 int8) int64 {
    for fi15 := 0; fi15 < 14; fi15++ {
        switch (fi15 * fi15) {
        case 1:
            v16 := GenS3{F0: (!((float32((1.5 * ((1.5 * 1.5) - (((3.14159 - (0.0 + -2.25)) - 1.5) * 3.14159)))) - 3.14159) >= (3.14159 * ((-2.25 / -2.25) / 1.0)))), F1: 180}
            _ = v16
        default:
            for fi17 := 0; fi17 < 11; fi17++ {
                if ((1 <= 876) || true) {
                    switch (fi15 >> ((255) & 63)) {
                    case 127:
                        p0 += p0
                    case 9223372036854775807:
                        var v18 uint16 = 32767
                        _ = v18
                    default:
                        p0 = (p0 * -1)
                    }
                } else {
                    var v19 uint8 = 127
                    _ = v19
                }
                if (true && (-1 < -140)) {
                    var v20 bool = false
                    if v20 {
                        var v21 uint32 = 65535
                        p0 = genFn0()
                        _ = v21
                    }
                    var v22 uint8 = 127
                    _ = v20
                    _ = v22
                } else {
                    if (2 < 895) {
                        p0 += (p0 * (p0 ^ genFn0()))
                        v23 := []int{fi17, fi17}
                        _ = v23
                    } else {
                        var v24 int = (fi17 % (((fi17 & -1)) | 1))
                        var v25 float32 = (-2.25 / 3.14159)
                        var v26 int16 = int16(fi15)
                        v27 := GenS1{F0: 1}
                        _ = v24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    p0 = (p0 / (((p0 / ((49) | 1))) | 1))
                }
                if (127 > 102) {
                    v28 := []int{(-(fi15)), -228, 265, 32767}
                    _ = v28
                }
                _ = fi17
            }
        }
        var v29 uint = 255
        _ = fi15
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 float32 = (3.14159 - (3.14159 + 1e10))
        _ = fi30
        _ = v31
    }
    var v32 int = -461
    v32 -= -782
    p0 = (p0 * p0)
    _ = p0
    _ = v32
    return -9223372036854775808
}

func main() {
    var v1 string = (("foo bar" + ("a" + "")) + "hello")
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 2
    var v4 int64 = -9223372036854775808
    v5 := make(map[uint32]int)
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 string = (v1 + v1)
        _ = fi33
        _ = v34
    }
    v35 := GenS4{F0: (128 == 127), F1: (v1 + ("" + ("hello" + v1))), F2: v3}
    if true {
        if false {
            delete(v5, 65535)
            for fi36 := 0; fi36 < 18; fi36++ {
                switch "foo bar" {
                case "\n\t":
                    switch v1 {
                    case "\n\t":
                        v1 = ((("" + (v1 + "hello")) + "foo bar") + v1)
                    }
                case "foo bar":
                    for fi37 := 0; fi37 < 18; fi37++ {
                        var v38 int = 2147483647
                        _ = fi37
                        _ = v38
                    }
                case "hello":
                    switch (v1 + v1) {
                    case "foo bar":
                        var v39 uint64 = 549
                        _ = v39
                    }
                }
                var v40 uint64 = 18446744073709551615
                _ = fi36
                _ = v40
            }
            if (((uint(v3) > 622) && true) || (1 < -1)) {
                for fi41 := 0; fi41 < 2; fi41++ {
                    for fi42 := 0; fi42 < 5; fi42++ {
                        switch 1 {
                        case 127:
                            v43 := []int{fi41, fi41, fi41, fi42}
                            _ = v43
                        }
                        fmt.Println(fi42, fi41, v1)
                        switch (fi41 % ((fi41) | 1)) {
                        case 127:
                            v3 -= v3
                        case -2147483648:
                            fmt.Println(v3, v1, v3)
                        case -1:
                            var v44 uint8 = 255
                            _ = v44
                        }
                        if false {
                            v1 = (v1 + "foo bar")
                        } else {
                            v45 := GenS5{F0: 679, F1: -764}
                            _ = v45
                        }
                        _ = fi42
                    }
                    var v46 uint = 587
                    _ = fi41
                    _ = v46
                }
            } else {
                var v47 uint64 = 4294967295
                switch "foo bar" {
                case "a":
                    v48 := GenS1{F0: 0}
                    _ = v48
                case "":
                    v49 := []int{9223372036854775807}
                    _ = v49
                case "\n\t":
                    for fi50 := 0; fi50 < 7; fi50++ {
                        for fi51 := 0; fi51 < 17; fi51++ {
                            v52 := []int{(^(fi51)), fi50}
                            var v53 int8 = -1
                            v47 -= (v47 | 361)
                            if true {
                                var v54 uint64 = 4294967295
                                _ = v54
                            }
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        _ = fi50
                    }
                default:
                    switch (v1 + (v1 + "foo bar")) {
                    case "\n\t":
                        if (v1 <= v1) {
                            delete(v2, 42)
                        }
                    }
                }
                var v55 int16 = -32768
                _ = v47
                _ = v55
            }
        }
        var v56 uint8 = 209
        for fi57 := 0; fi57 < 1; fi57++ {
            var v58 int8 = 24
            _ = fi57
            _ = v58
        }
        _ = v56
    } else {
        if false {
            v4 = (-(v4))
        } else {
            fmt.Println(v3, v1, v1)
        }
        v4 = v4
        var v59 int64 = v4
        _ = v59
    }
    v3 -= 65535
    v60 := GenS3{F0: false, F1: 2}
    delete(v5, 2)
    fmt.Println(v4, v4)
    v3 = (v3 & v3)
    v3 = (v3 << ((2) & 63))
    var v61 uint = 42
    v3 = 1
    v62 := GenS6{F0: 0, F1: "foo bar", F2: v3}
    for fi63 := 0; fi63 < 7; fi63++ {
        if (!true) {
            if (0 >= 1) {
                if true {
                    fmt.Println(v3)
                    v61 += 42
                } else {
                    var v64 int32 = -2147483648
                    delete(v2, uint32(v3))
                    var v65 int64 = v4
                    _ = v64
                    _ = v65
                }
            }
        }
        v66 := GenS7{F0: 100}
        _ = fi63
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v62
}
