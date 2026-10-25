package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint16; F1 uint; F2 uint64 }
func genFn0() bool {
    var v4 uint16 = 17
    v5 := []int{608, 128}
    v5 = append(v5, 32767)
    var v6 float32 = float32(1.5)
    for fi7 := 0; fi7 < 6; fi7++ {
        if (!((637 >= 4294967295) || (0 == 0))) {
            if false {
                v4 = v4
            } else {
                v4 -= 255
                var v8 uint32 = uint32(v4)
                _ = v8
            }
            if (v4 >= v4) {
                if ((888 >= 2147483648) && ((fi7 % ((fi7) | 1)) <= fi7)) {
                    v9 := GenS1{F0: 840, F1: v4}
                    var v10 uint64 = 42
                    if len(v5) > 0 {
                        v5[0] = -2147483648
                    }
                    _ = v10
                    _ = v9
                } else {
                    fmt.Println(v6, v6, v4)
                }
                if (!((float32((-2.25 + ((-2.25 / (float64(1.5) + 3.14159)) * (0.1 + 1e10)))) - v6) < v6)) {
                    if len(v5) > 0 {
                        v5[0] = fi7
                    }
                }
                v11 := GenS1{F0: 2147483647, F1: v4}
                _ = v11
            } else {
                if true {
                    switch "\n\t" {
                    case "hello":
                        switch "foo bar" {
                        case "hello":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "\n\t":
                            _ = len(v5)
                        case "":
                            var v12 int = fi7
                            _ = v12
                        }
                    }
                    var v13 float64 = float64((v6 + float32(-2.25)))
                    switch "" {
                    case "foo bar":
                        var v14 int32 = 42
                        _ = v14
                    }
                    _ = v13
                }
                v5 = append(v5, (fi7 - (fi7 * (fi7 << ((0) & 63)))))
                for fi15 := 0; fi15 < 10; fi15++ {
                    switch fi7 {
                    case 2147483647:
                        v16 := GenS1{F0: 1, F1: (^(v4))}
                        _ = v16
                    case 4294967295:
                        if false {
                            v6 -= ((1e10 - 0.0) * ((float32(1e10) * v6) / -2.25))
                            var v17 int64 = 175
                            var v18 uint32 = 42
                            _ = v17
                            _ = v18
                        }
                    case 51:
                        var v19 int32 = -2147483648
                        _ = v19
                    }
                    _ = fi15
                }
            }
            if (3.14159 != float64((0.0 - ((v6 + float32(1.5)) - ((v6 * float32(1.5)) - v6))))) {
                v5 = append(v5, -720)
                for fi20 := 0; fi20 < 7; fi20++ {
                    v5 = append(v5, 2147483647)
                    _ = fi20
                }
            } else {
                if (!(true || (!true))) {
                    if len(v5) > 0 {
                        v5[0] = (fi7 << ((1) & 63))
                    }
                    fmt.Println(fi7, v4, v6)
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v6 -= v6
                        v22 := GenS2{F0: 9223372036854775807, F1: ("hello" + "\n\t")}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    v23 := GenS2{F0: 0, F1: "hello"}
                    v24 := GenS2{F0: 217, F1: ("a" + "foo bar")}
                    v4 = v4
                    _ = v23
                    _ = v24
                }
                var v25 int8 = -65
                _ = v25
            }
        } else {
            var v26 uint16 = v4
            fmt.Println(v26)
            if len(v5) > 0 {
                v5[0] = (fi7 / ((fi7) | 1))
            }
            _ = v26
        }
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 2147483647
    var v3 bool = (true || (0 > 127))
    switch v2 {
    case 42:
        v2 -= -2147483648
    case -596:
        v2 = -116
    case 128:
        var v27 float64 = (0.0 - 0.1)
        _ = v27
    }
    var v28 int8 = -38
    fmt.Println(v3, v2, v3)
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS2{F0: 9223372036854775807, F1: ("\n\t" + "\n\t")}
        v31 := GenS3{F0: 1, F1: 815, F2: 933}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int8 = 42
    v33 := []int{}
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 string = "foo bar"
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        _ = fi34
        _ = v35
    }
    v32 -= (^(v32))
    var v36 int64 = 1
    var v37 uint8 = 14
    for fi38 := 0; fi38 < 17; fi38++ {
        v39 := GenS2{F0: v36, F1: "\n\t"}
        fmt.Println(v37)
        _ = fi38
        _ = v39
    }
    v3 = v3
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    if v3 {
        v32 = (-(v28))
        for fi40 := 0; fi40 < 9; fi40++ {
            var v41 uint8 = (v37 % ((v37) | 1))
            _ = fi40
            _ = v41
        }
    }
    v3 = v3
    for fi42 := 0; fi42 < 3; fi42++ {
        var v43 float64 = (3.14159 + (1e10 / float64((((1e10 * float32((float64(((((float32(0.0) / 3.14159) / 0.1) + 1.5) / 1.0)) / 1.0))) + 0.1) + 0.0))))
        _ = fi42
        _ = v43
    }
    v36 += v36
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
}
