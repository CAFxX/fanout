package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int16; F2 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint, p1 int64) int {
    if (!false) {
        fmt.Println(p1, p1)
        v4 := []int{127, 359, -1, 702}
        _ = v4
    }
    if false {
        v5 := GenS1{F0: 255, F1: true, F2: 2147483647}
        switch "foo bar" {
        case "hello":
            var v6 uint8 = 2
            _ = v6
        case "a":
            p0 += (p0 | p0)
        case "":
            for fi7 := 0; fi7 < 14; fi7++ {
                for fi8 := 0; fi8 < 5; fi8++ {
                    p0 = (p0 * (^(p0)))
                    _ = fi8
                }
                _ = fi7
            }
        }
        _ = v5
    }
    fmt.Println(p0, p1)
    var v9 uint64 = 498
    var v10 uint8 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v11 float64 = 0.0
    switch (v1 + v1) {
    case "\n\t":
        var v12 bool = ((!(!(42 <= -2147483648))) && (!false))
        _ = v12
    }
    var v13 uint8 = 2
    if (true || (!(0.0 < 1.5))) {
        var v14 uint64 = 781
        var v15 float32 = 0.1
        v16 := GenS2{F0: 2147483647, F1: 1}
        _ = v14
        _ = v15
        _ = v16
    } else {
        v17 := GenS3{F0: -2147483648}
        v18 := []int{2147483647, 626}
        var v19 float64 = 1e10
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 bool = false
    var v21 int8 = 1
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 bool = false
        v11 -= (1e10 * (((0.0 / float64((3.14159 * 3.14159))) * v11) - v11))
        _ = fi22
        _ = v23
    }
    v24 := GenS4{F0: 405, F1: 1000, F2: v20}
    fmt.Println(v11)
    switch -2147483648 {
    case 180:
        var v25 uint64 = 0
        _ = v25
    case 2147483647:
        var v26 float64 = (v11 + v11)
        _ = v26
    case -749:
        for fi27 := 0; fi27 < 7; fi27++ {
            if true {
                switch "foo bar" {
                case "":
                    for fi28 := 0; fi28 < 7; fi28++ {
                        v29 := GenS5{F0: v20, F1: 0}
                        switch v1 {
                        case "":
                            var v30 uint16 = 1
                            _ = v30
                        case "\n\t":
                            fmt.Println(v13, v11, fi27)
                            switch fi28 {
                            case -202:
                                v11 = ((v11 + 0.0) + v11)
                            case 128:
                                var v31 string = (((v1 + (v1 + "\n\t")) + v1) + ("foo bar" + v1))
                                _ = v31
                            case 127:
                                if v20 {
                                    v32 := v3[65535]
                                    fmt.Println(v20, fi28)
                                    v33 := GenS5{F0: false, F1: int16(fi27)}
                                    v1 = "hello"
                                    if (uint32(v13) < 1) {
                                        var v34 float32 = 3.14159
                                        _ = v34
                                    }
                                    _ = v32
                                    _ = v33
                                }
                            default:
                                switch genFn0(18446744073709551615, 53) {
                                case -1:
                                    var v35 int64 = -599
                                    _ = v35
                                case -694:
                                    fmt.Println(v1, v20, fi27)
                                case 32768:
                                    fmt.Println(v21)
                                }
                            }
                            switch v1 {
                            case "a":
                                v36 := []int{(-(fi27)), 1}
                                _ = v36
                            case "foo bar":
                                v37 := GenS1{F0: 0, F1: (v20 && (((v11 - v11) > (v11 * 1e10)) && true)), F2: 321}
                                _ = v37
                            case "\n\t":
                                v38 := []int{}
                                _ = v38
                            default:
                                switch genFn0(5, -9223372036854775808) {
                                case 1:
                                    v21 -= v21
                                case 28:
                                    v1 = "\n\t"
                                }
                            }
                            v20 = v20
                        default:
                            v39 := GenS6{F0: (v21 + 98)}
                            _ = v39
                        }
                        _ = fi28
                        _ = v29
                    }
                case "hello":
                    switch genFn0(4294967295, -420) {
                    case 128:
                        var v40 uint32 = 65535
                        _ = v40
                    default:
                        fmt.Println(v21, v1, v21)
                    }
                }
                switch 127 {
                case -1:
                    v20 = v20
                case -994:
                    v13 = 128
                case 127:
                    var v41 uint32 = 2147483647
                    _ = v41
                }
            } else {
                v42 := []int{(^(fi27)), genFn0(519, -9223372036854775808)}
                v1 = ((v1 + "hello") + v1)
                _ = v42
            }
            v43 := []int{805, 4294967295, (fi27 & fi27)}
            _ = fi27
            _ = v43
        }
    default:
        v13 -= 34
    }
    v11 += v11
    var v44 float32 = float32((v11 - v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v44
}
