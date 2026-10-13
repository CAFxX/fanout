package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 string) int8 {
    if true {
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
    }
    var v5 float64 = 0.0
    if (-789 == -879) {
        var v6 int64 = -1
        for fi7 := 0; fi7 < 14; fi7++ {
            p0 = "foo bar"
            v8 := []int{(-(fi7)), (fi7 % ((fi7) | 1)), 2147483647}
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 int64 = v6
                var v11 uint32 = 850
                v12 := GenS1{F0: true, F1: 65535, F2: true}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        var v13 uint64 = 0
        _ = v13
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 int8, p1 int64, p2 uint) uint8 {
    v14 := []int{-9223372036854775808, 127, 0}
    v15 := GenS1{F0: (!(!(true && false))), F1: p2, F2: (uint16(p2) == 222)}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 127
}

func main() {
    var v1 uint16 = 193
    var v2 uint32 = 410
    var v3 string = ("hello" + ("a" + "\n\t"))
    v4 := make(map[string]bool)
    v16 := []int{-9223372036854775808, 42}
    if true {
        v17 := GenS1{F0: false, F1: 1, F2: true}
        switch 127 {
        case 42:
            var v18 uint16 = (v1 * v1)
            _ = v18
        case -613:
            if false {
                for fi19 := 0; fi19 < 2; fi19++ {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        v1 += 16
                        if false {
                            v3 = (("foo bar" + v3) + v3)
                            switch "foo bar" {
                            case "a":
                                switch 42 {
                                case 127:
                                    fmt.Println(v2)
                                }
                                fmt.Println(v3, v1, v2)
                            }
                            fmt.Println(v3, v3)
                            var v21 uint16 = 0
                            switch fi19 {
                            case -2147483648:
                                var v22 int64 = 288
                                _ = v22
                            case 46:
                                v21 -= (v21 << ((65535) & 63))
                            }
                            _ = v21
                        }
                        _ = fi20
                    }
                    _ = fi19
                }
                v1 = 2
                var v23 uint64 = 42
                _ = v23
            }
        }
        _ = v17
    }
    fmt.Println(v2, v3, v2)
    if (-32768 >= -32768) {
        for fi24 := 0; fi24 < 15; fi24++ {
            if false {
                v2 += (v2 % (((v2 | (v2 & 42))) | 1))
                v2 += (v2 / ((v2) | 1))
                v25 := GenS2{F0: genFn0(((v3 + v3) + "\n\t"))}
                _ = v25
            } else {
                fmt.Println(v1, v3, v1)
            }
            _ = fi24
        }
    }
    v1 = (v1 % ((v1) | 1))
    var v26 float64 = (3.14159 * -2.25)
    var v27 string = "foo bar"
    v3 = "foo bar"
    var v28 int16 = 116
    if (v1 != (v1 - 42)) {
        for fi29 := 0; fi29 < 15; fi29++ {
            v3 = ""
            for fi30 := 0; fi30 < 10; fi30++ {
                if true {
                    v16 = append(v16, fi30)
                    if ((v26 + float64((0.1 * 3.14159))) > (float64(-2.25) / ((float64(((1e10 + 1.5) * 1.5)) - v26) + (float64((1e10 + (-2.25 * 1e10))) - v26)))) {
                        var v31 int16 = (v28 / ((-543) | 1))
                        if (!true) {
                            if false {
                                var v32 uint8 = 42
                                v3 = v3
                                _ = v32
                            }
                            var v33 int16 = v31
                            _ = v33
                        } else {
                            v2 -= 0
                        }
                        _ = v31
                    }
                    var v34 uint32 = v2
                    _ = v34
                }
                _ = fi30
            }
            _ = fi29
        }
    }
    v35 := GenS1{F0: (9223372036854775807 != 32767), F1: 752, F2: false}
    v36 := GenS3{F0: 69, F1: v28}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
