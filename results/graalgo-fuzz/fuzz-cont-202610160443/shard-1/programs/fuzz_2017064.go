package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int; F1 uint; F2 uint32 }
type GenS5 struct { F0 bool; F1 bool; F2 string }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint8) int {
    p0 = p0
    var v5 int8 = -128
    var v6 int64 = 9223372036854775807
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 674
    v3 := make(map[int]string)
    var v4 uint16 = 1
    v4 = (v2 >> ((2) & 63))
    for fi7 := 0; fi7 < 16; fi7++ {
        v4 = (v4 & (v2 | (v2 ^ 0)))
        _ = fi7
    }
    v2 += 255
    v1 += (v1 / ((v1) | 1))
    v8 := GenS1{F0: (v1 | 1), F1: 1}
    v9 := GenS2{F0: 0, F1: 2}
    v10 := []int{1, 0}
    var v11 int32 = 729
    if (v4 > v2) {
        if true {
            if true {
                fmt.Println(v4, v4)
                if false {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        if (128 <= fi12) {
                            var v13 int = 163
                            switch ("\n\t" + "a") {
                            case "\n\t":
                                for fi14 := 0; fi14 < 1; fi14++ {
                                    if (0 < 2) {
                                        var v15 float64 = (3.14159 + 3.14159)
                                        v16 := GenS3{F0: (^(v1)), F1: (v13 - -917)}
                                        switch ("hello" + "") {
                                        case "\n\t":
                                            var v17 string = ("\n\t" + ("hello" + "a"))
                                            _ = v17
                                        case "hello":
                                            v13 -= -641
                                        }
                                        _ = v15
                                        _ = v16
                                    }
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    v18 := GenS4{F0: fi12, F1: 490, F2: uint32(v2)}
                                    _ = fi14
                                    _ = v18
                                }
                            }
                            _ = v13
                        } else {
                            var v19 int64 = 1
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 13; fi20++ {
                            v4 = (^(v2))
                            v2 -= v4
                            v21 := GenS5{F0: true, F1: (1 > uint32(v2)), F2: "foo bar"}
                            v4 = v4
                            _ = fi20
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 18; fi22++ {
                            v11 = (v11 + -2147483648)
                            _ = fi22
                        }
                        _ = fi12
                    }
                }
                v23 := GenS1{F0: (v1 | v1), F1: 295}
                _ = v23
            } else {
                v24 := GenS1{F0: 1000, F1: 394}
                _ = len(v10)
                _ = v24
            }
            v10 = append(v10, 128)
            var v25 int64 = -460
            _ = v25
        } else {
            v11 = (v11 | (v11 / (((v11 | (v11 * int32(v1)))) | 1)))
            v11 -= (v11 - v11)
        }
        v26 := GenS4{F0: 625, F1: 1, F2: 42}
        switch -9223372036854775808 {
        case 558:
            fmt.Println(v11, v1, v4)
        case -2147483648:
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 uint8 = 1
                v3[(fi27 - fi27)] = "\n\t"
                _ = fi27
                _ = v28
            }
        case 4294967295:
            var v29 bool = (float32(1.5) == 1e10)
            _ = v29
        default:
            if false {
                var v30 uint32 = 997
                for fi31 := 0; fi31 < 6; fi31++ {
                    v32 := GenS3{F0: (v1 ^ 0), F1: fi31}
                    _ = fi31
                    _ = v32
                }
                if true {
                    var v33 uint = 582
                    v34 := GenS5{F0: (false && (1 < 2)), F1: ((v2 + v4) <= v4), F2: ("a" + ("a" + "a"))}
                    var v35 uint = v33
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = v30
            }
        }
        _ = v26
    } else {
        v11 = (v11 & (v11 / ((v11) | 1)))
        v36 := GenS6{F0: 619, F1: -882}
        v37 := GenS5{F0: false, F1: false, F2: "hello"}
        _ = v36
        _ = v37
    }
    v38 := GenS7{F0: 2}
    var v39 int32 = (v11 | (^(v11)))
    v40 := GenS8{F0: 1, F1: v39}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v8
    _ = v9
}
