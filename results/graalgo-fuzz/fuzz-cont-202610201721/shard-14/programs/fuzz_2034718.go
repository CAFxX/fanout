package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
func genFn0(p0 int16, p1 int32, p2 string) string {
    var v6 uint = 255
    v6 += (v6 >> ((v6) & 63))
    p2 = "foo bar"
    v7 := []int{}
    switch 687 {
    case 32767:
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int64 = 43
            var v10 uint = v6
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        p1 = 0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ("a" + p2)
}
func genFn1() uint32 {
    var v11 int = -92
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: 0, F1: 0}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if ((-2.25 + 0.1) <= (float64(1e10) + ((0.0 + 1.5) / 3.14159))) {
            var v15 string = ""
            for fi16 := 0; fi16 < 11; fi16++ {
                if (255 == 116) {
                    if ((0.0 / 3.14159) != (1e10 - 1e10)) {
                        if (326 > -32768) {
                            v11 = fi16
                        }
                        v17 := GenS1{F0: 32767, F1: 1}
                        v18 := []int{(fi16 % ((-2147483648) | 1)), (fi14 ^ fi14), v11, (fi14 | (fi14 ^ fi14))}
                        _ = v17
                        _ = v18
                    }
                } else {
                    var v19 bool = (2 <= 18446744073709551615)
                    var v20 uint16 = 32767
                    _ = v19
                    _ = v20
                }
                var v21 uint8 = 2
                _ = fi16
                _ = v21
            }
            v15 = ("" + v15)
            _ = v15
        } else {
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := []int{}
                var v24 uint16 = 424
                if true {
                    v24 = (v24 << ((0) & 63))
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        _ = fi14
    }
    v25 := []int{-644, (-(v11))}
    var v26 bool = true
    _ = v11
    _ = v25
    _ = v26
    return 541
}

func main() {
    var v1 int32 = 42
    var v2 float64 = float64(-2.25)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 34
    v5 := make(map[uint32]string)
    switch "foo bar" {
    case "hello":
        v27 := []int{}
        _ = v27
    case "":
        var v28 uint32 = v4
        _ = v28
    case "foo bar":
        var v29 uint64 = 149
        _ = v29
    default:
        fmt.Println(v1, v1, v1)
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            v1 -= v1
            _ = fi31
        }
        switch 32767 {
        case 181:
            v32 := GenS1{F0: uint16(v4), F1: 42}
            _ = v32
        case 64:
            var v33 int32 = (v1 & 0)
            _ = v33
        case 42:
            v4 += v4
        }
        v34 := []int{0, fi30, (fi30 & 257), fi30}
        _ = fi30
        _ = v34
    }
    v35 := GenS1{F0: 1, F1: -2147483648}
    var v36 int = -6
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 int64 = 605
            v40 := []int{(fi38 >> ((214) & 63)), fi37, fi38}
            switch "foo bar" {
            case "\n\t":
                v1 = (v1 - -1)
                v1 -= (v1 * v1)
                v36 -= (fi37 - fi38)
                var v41 uint16 = 121
                _ = v41
            case "a":
                fmt.Println(v4, v1, fi38)
            default:
                var v42 uint32 = (v4 ^ genFn1())
                _ = v42
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        var v43 int8 = 127
        _ = fi37
        _ = v43
    }
    switch v36 {
    case 128:
        if false {
            var v44 int64 = -1
            _ = v44
        } else {
            v45 := GenS2{F0: (^(v4)), F1: genFn1(), F2: ((v1 > 920) && ((v36 / (((v36 / (((v36 ^ (v36 >> ((18446744073709551615) & 63)))) | 1))) | 1)) >= -2147483648))}
            _ = v45
        }
    case 32768:
        if (1 < -405) {
            v1 = (v1 & 42)
            var v46 float32 = 3.14159
            v36 += v36
            _ = v46
        }
    }
    var v47 uint8 = 255
    fmt.Println(v1)
    v2 = ((0.0 - float64(1e10)) / (v2 * (v2 - (v2 * v2))))
    for fi48 := 0; fi48 < 7; fi48++ {
        fmt.Println(v47, v4)
        fmt.Println(v2, v2, fi48)
        _ = fi48
    }
    if false {
        var v49 int32 = int32(v4)
        var v50 uint64 = 66
        var v51 int64 = -309
        _ = v49
        _ = v50
        _ = v51
    }
    switch (^(v36)) {
    case 302:
        fmt.Println(v1, v47)
    case -241:
        var v52 uint8 = (v47 + v47)
        _ = v52
    }
    var v53 uint32 = genFn1()
    fmt.Println(v2, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v47
    _ = v5
    _ = v53
}
