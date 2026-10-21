package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 bool; F1 int64 }
func genFn0(p0 uint16) string {
    v6 := GenS1{F0: (p0 * (p0 >> ((65535) & 63)))}
    fmt.Println(p0)
    var v7 string = ("foo bar" + "")
    v8 := []int{-9223372036854775808, -9223372036854775808}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 uint16, p1 int8, p2 bool) uint32 {
    var v9 uint16 = (p0 | p0)
    for fi10 := 0; fi10 < 1; fi10++ {
        if ((((42 > 45) && p2) || p2) && (((fi10 & -9223372036854775808) >= (fi10 / ((fi10) | 1))) && p2)) {
            v11 := GenS1{F0: p0}
            _ = v11
        } else {
            var v12 float64 = 1e10
            switch (fi10 - fi10) {
            case -2147483648:
                for fi13 := 0; fi13 < 3; fi13++ {
                    switch (fi10 + fi10) {
                    case 0:
                        if p2 {
                            p0 = (^(p0))
                        }
                    case -2147483648:
                        if true {
                            p0 = 65535
                        }
                    case -1:
                        fmt.Println(p1, v12, p1)
                    }
                    v12 = v12
                    _ = fi13
                }
            }
            _ = v12
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 float64 = 0.0
            _ = fi14
            _ = v15
        }
        var v16 bool = (!(-69 != 9223372036854775807))
        _ = fi10
        _ = v16
    }
    var v17 uint16 = (^(p0))
    switch 2147483647 {
    case 55:
        v18 := GenS1{F0: v17}
        _ = v18
    default:
        var v19 float32 = float32(1.5)
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v9
    return 4294967295
}

func main() {
    var v1 int = 495
    v2 := make(map[uint32]uint32)
    v3 := []string{"\n\t"}
    var v4 uint8 = 128
    var v5 int16 = 1000
    for fi20 := 0; fi20 < 17; fi20++ {
        v1 += (fi20 + (v1 % ((v1) | 1)))
        _ = fi20
    }
    var v21 uint8 = 167
    v21 = (^(v21))
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 uint32 = genFn1(1, -64, (227 > 2))
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = fi23
            _ = v24
        }
        if (!(!(!true))) {
            if true {
                if ((v1 <= (fi22 + (v1 ^ fi22))) || true) {
                    delete(v2, 1)
                    var v25 float32 = 0.0
                    _ = v25
                } else {
                    if (2 < 42) {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            v21 -= 127
                            v1 = fi26
                            _ = fi26
                        }
                        var v27 uint32 = 1
                        _ = v27
                    } else {
                        v28 := GenS2{F0: 1, F1: 0, F2: 42}
                        for fi29 := 0; fi29 < 5; fi29++ {
                            var v30 int64 = -9223372036854775808
                            fmt.Println(v30)
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    }
                    var v31 int16 = 0
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    _ = v31
                }
                _ = len(v2)
                v32 := GenS2{F0: 57, F1: 2, F2: 861}
                _ = v32
            } else {
                var v33 int64 = 0
                v21 -= v4
                _ = v33
            }
            if true {
                v5 = 42
            } else {
                fmt.Println(v1)
                v5 = (v5 * v5)
            }
        } else {
            var v34 int = (v1 | -2147483648)
            var v35 bool = (!("\n\t" > "hello"))
            _ = v34
            _ = v35
        }
        v5 = 42
        _ = fi22
    }
    v36 := GenS3{F0: "", F1: uint64(v21), F2: 4294967295}
    v1 += v1
    if (4294967295 <= 2) {
        v37 := GenS4{F0: 139, F1: (v5 * v5)}
        fmt.Println(v1, v4)
        _ = v37
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        fmt.Println(v4)
        switch "\n\t" {
        case "foo bar":
            if true {
                var v39 int = v1
                for fi40 := 0; fi40 < 7; fi40++ {
                    v39 -= fi40
                    if (0.1 == float32((3.14159 - (1e10 / 1.0)))) {
                        var v41 int = (v1 - fi40)
                        switch "" {
                        case "foo bar":
                            var v42 uint16 = 999
                            _ = v42
                        case "\n\t":
                            var v43 float32 = (0.1 + (((3.14159 - 3.14159) + -2.25) - (1e10 * (3.14159 * float32(1e10)))))
                            _ = v43
                        case "hello":
                            var v44 uint16 = 1
                            _ = v44
                        }
                        var v45 uint8 = 0
                        _ = v41
                        _ = v45
                    }
                    v39 += 127
                    _ = fi40
                }
                _ = v39
            }
        case "\n\t":
            _ = len(v3)
        case "":
            if true {
                v4 += 168
                v5 += 42
            }
        }
        v5 -= (^(v5))
        _ = fi38
    }
    v5 -= v5
    if true {
        var v46 float64 = -2.25
        v46 -= v46
        switch (v1 | (^(v1))) {
        case 936:
            for fi47 := 0; fi47 < 8; fi47++ {
                if false {
                    v48 := GenS4{F0: 18446744073709551615, F1: 55}
                    if false {
                        var v49 int16 = (v5 - v5)
                        _ = v49
                    }
                    v4 += 255
                    _ = v48
                }
                _ = fi47
            }
        case 42:
            _ = len(v3)
        default:
            var v50 string = ""
            _ = v50
        }
        _ = v46
    }
    var v51 int = v1
    var v52 float32 = 1.5
    for fi53 := 0; fi53 < 14; fi53++ {
        if (!(false || true)) {
            var v54 float64 = 1.5
            v2[genFn1(42, 19, ((true && (27 <= 2147483648)) && (2147483647 >= (fi53 - (fi53 | v1)))))] = genFn1(0, -122, ((false && true) || true))
            for fi55 := 0; fi55 < 1; fi55++ {
                var v56 int16 = (v5 % ((654) | 1))
                var v57 uint64 = 18446744073709551615
                _ = fi55
                _ = v56
                _ = v57
            }
            _ = v54
        }
        _ = fi53
    }
    var v58 int16 = (-(v5))
    var v59 int32 = 42
    if true {
        for fi60 := 0; fi60 < 15; fi60++ {
            if (!((1 == (v4 % ((v21) | 1))) || (!(-771 >= 1)))) {
                v61 := GenS2{F0: 127, F1: 489, F2: 65535}
                var v62 float32 = (float32(3.14159) - ((v52 + v52) + 0.1))
                var v63 float32 = float32(0.0)
                _ = v61
                _ = v62
                _ = v63
            }
            v64 := GenS5{F0: false, F1: 0}
            var v65 float64 = (0.1 * 1.5)
            _ = fi60
            _ = v64
            _ = v65
        }
    }
    var v66 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v58
    _ = v59
    _ = v66
}
