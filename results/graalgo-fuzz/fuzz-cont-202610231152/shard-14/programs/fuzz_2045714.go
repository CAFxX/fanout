package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string; F2 int }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 bool = (683 != 771)
    var v5 float64 = ((1.5 + 1.5) * 1.5)
    var v6 string = "hello"
    switch 127 {
    case 0:
        var v7 uint64 = 42
        _ = v7
    case 32767:
        v8 := []int{-2147483648, 470}
        _ = v8
    }
    if v4 {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 42}
            delete(v1, "foo bar")
            _ = fi9
            _ = v10
        }
        var v11 float64 = v5
        _ = v11
    } else {
        fmt.Println(v5, v5)
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int32 = 42
            _ = fi12
            _ = v13
        }
        switch 32768 {
        case -212:
            v4 = v4
        case 4294967295:
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := []int{fi14, fi14}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                var v16 int64 = 9223372036854775807
                _ = fi14
                _ = v15
                _ = v16
            }
        case 314:
            var v17 bool = false
            _ = v17
        }
    }
    v18 := GenS1{F0: 18446744073709551615}
    for fi19 := 0; fi19 < 14; fi19++ {
        if false {
            if v4 {
                _ = len(v3)
                v4 = v4
                v5 += ((1e10 / -2.25) + (v5 * -2.25))
            } else {
                fmt.Println(fi19, v4, v4)
                for fi20 := 0; fi20 < 4; fi20++ {
                    v4 = v4
                    v4 = v4
                    delete(v1, v6)
                    _ = fi20
                }
            }
            v21 := GenS2{F0: 42, F1: (v6 + "\n\t"), F2: (fi19 - (-(fi19)))}
            if (v4 || v4) {
                fmt.Println(v4)
                var v22 int64 = -568
                v23 := []int{}
                _ = v22
                _ = v23
            }
            _ = v21
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            v25 := GenS1{F0: 494}
            if (586 != 862) {
                if true {
                    v26 := []int{}
                    if ((-2.25 * ((3.14159 - ((3.14159 / (-2.25 - 1.5)) + float32(-2.25))) * (0.0 / 1e10))) != 0.0) {
                        v6 = (((v6 + "") + ((v6 + "\n\t") + v6)) + v6)
                        if (682 <= 18446744073709551615) {
                            if (2 == 112) {
                                v4 = v4
                                v6 = ((((("" + "\n\t") + v6) + v6) + "foo bar") + "\n\t")
                            } else {
                                if (1 == 1) {
                                    var v27 bool = (4294967295 >= 2)
                                    _ = v27
                                }
                                v26 = append(v26, (fi24 & (fi19 * 32767)))
                                switch (^(fi24)) {
                                case -868:
                                    var v28 bool = (0.1 != (0.1 - 1e10))
                                    _ = v28
                                case 2147483647:
                                    v4 = (32767 == 42)
                                case -1:
                                    v4 = v4
                                default:
                                    if len(v26) >= 2 {
                                        v26 = v26[1:len(v26)]
                                    }
                                }
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                var v29 int32 = -1
                                switch "a" {
                                case "":
                                    v30 := GenS1{F0: 18446744073709551615}
                                    var v31 float64 = (0.0 * (v5 / 1.0))
                                    v29 = (-(v29))
                                    _ = v30
                                    _ = v31
                                default:
                                    v32 := GenS3{F0: 623, F1: -308}
                                    _ = v32
                                }
                                _ = v29
                            }
                        } else {
                            v6 = ("hello" + (((((v6 + v6) + v6) + "") + "\n\t") + ((((v6 + ((v6 + "a") + v6)) + ("a" + "")) + "\n\t") + v6)))
                        }
                    }
                    _ = v26
                } else {
                    var v33 int = 179
                    v34 := GenS3{F0: 0, F1: (fi19 | (fi19 / ((-206) | 1)))}
                    _ = v33
                    _ = v34
                }
                var v35 float64 = (v5 + (float64(((3.14159 - 1e10) / -2.25)) + 0.0))
                _ = v35
            }
            var v36 float64 = ((((v5 / (float64(1e10) * float64(-2.25))) - 0.1) - (1.5 / 0.1)) - v5)
            _ = fi24
            _ = v25
            _ = v36
        }
        _ = fi19
    }
    fmt.Println(v6)
    for fi37 := 0; fi37 < 5; fi37++ {
        v38 := GenS4{F0: 1000}
        var v39 uint32 = 0
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 float64 = v5
    fmt.Println(v40, v40, v40)
    var v41 uint16 = 480
    delete(v3, int(v41))
    v40 -= 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
