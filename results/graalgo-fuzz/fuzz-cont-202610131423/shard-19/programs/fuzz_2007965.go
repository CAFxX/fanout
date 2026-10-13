package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint16) int16 {
    var v4 int = 4294967295
    v5 := GenS1{F0: 42}
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 uint = 0
    var v2 int16 = 1000
    var v3 int8 = 1
    v1 = (v1 >> ((v1) & 63))
    v6 := GenS2{F0: 215, F1: false, F2: 0}
    for fi7 := 0; fi7 < 14; fi7++ {
        if ("a" >= "hello") {
            var v8 bool = true
            fmt.Println(v2)
            _ = v8
        }
        v1 = (v1 + 971)
        var v9 string = "hello"
        _ = fi7
        _ = v9
    }
    v10 := []int{}
    v2 -= (v2 / ((v2) | 1))
    v11 := GenS1{F0: 1}
    v12 := GenS1{F0: 29}
    if ((v1 << ((v1) & 63)) < (v1 % (((v1 << ((v1) & 63))) | 1))) {
        switch 522 {
        case 32768:
            var v13 int32 = -2147483648
            _ = v13
        default:
            if true {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v10 = append(v10, fi14)
                    _ = fi14
                }
                var v15 string = ("hello" + "\n\t")
                _ = v15
            }
        }
        v1 = 0
    }
    var v16 float64 = (float64(0.1) * 0.1)
    for fi17 := 0; fi17 < 15; fi17++ {
        if (!(!true)) {
            if false {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v18 uint = (v1 * 18446744073709551615)
                if true {
                    v19 := GenS3{F0: "foo bar", F1: (fi17 - (-(fi17)))}
                    v18 -= 255
                    v10 = append(v10, 1)
                    _ = v19
                }
                _ = v18
            } else {
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v2)
                    _ = fi20
                }
                if true {
                    v16 -= ((((0.1 / v16) - v16) / (float64(0.0) * v16)) / v16)
                    switch 1 {
                    case 2147483647:
                        v10 = append(v10, 9223372036854775807)
                    case 32767:
                        v21 := GenS4{F0: fi17, F1: 65535}
                        _ = v21
                    default:
                        v2 = (v2 | 1)
                    }
                    var v22 float32 = 0.1
                    _ = v22
                } else {
                    v2 = (-(v2))
                    switch int(v2) {
                    case 2147483647:
                        v16 += v16
                    case 42:
                        if ((v2 ^ v2) >= genFn0(136)) {
                            if false {
                                fmt.Println(fi17, fi17, v3)
                                var v23 bool = (false && true)
                                _ = v23
                            }
                            v1 += 522
                            v24 := GenS4{F0: fi17, F1: 32767}
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 3; fi25++ {
                                var v26 int32 = -2147483648
                                _ = fi25
                                _ = v26
                            }
                        }
                    }
                }
            }
            var v27 bool = true
            fmt.Println(v3, v27)
            _ = v27
        } else {
            v16 += -2.25
            if (float32(-2.25) != (-2.25 * 0.0)) {
                v16 += -2.25
                v2 = -32768
            } else {
                switch fi17 {
                case 2147483647:
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                case 945:
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if (uint16(v1) >= 529) {
                            var v29 float64 = ((v16 / (1.5 + 3.14159)) / 1.0)
                            v1 = 476
                            _ = v29
                        } else {
                            v3 = (-(v3))
                            v30 := GenS5{F0: 255, F1: 9223372036854775807}
                            _ = v30
                        }
                        _ = fi28
                    }
                }
                fmt.Println(v3, v2)
                v31 := GenS2{F0: 42, F1: true, F2: v3}
                _ = v31
            }
        }
        var v32 int32 = 374
        fmt.Println(fi17, v1, v16)
        _ = fi17
        _ = v32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 uint64 = 18446744073709551615
        v1 = (v1 & v1)
        v3 += v3
        _ = fi33
        _ = v34
    }
    var v35 uint64 = 18446744073709551615
    var v36 uint16 = 0
    var v37 uint64 = v35
    for fi38 := 0; fi38 < 5; fi38++ {
        var v39 uint16 = v36
        var v40 uint8 = 146
        var v41 uint16 = (v36 * 1)
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v6
}
