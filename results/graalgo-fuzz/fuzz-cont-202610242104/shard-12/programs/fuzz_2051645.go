package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint32) int8 {
    var v6 uint32 = p0
    switch -195 {
    case 781:
        switch "a" {
        case "a":
            for fi7 := 0; fi7 < 3; fi7++ {
                v6 -= 454
                v8 := []int{(^(fi7)), (fi7 >> ((255) & 63)), (fi7 / ((652) | 1)), 9223372036854775807}
                var v9 string = ("hello" + (("hello" + "") + "\n\t"))
                _ = fi7
                _ = v8
                _ = v9
            }
        case "foo bar":
            for fi10 := 0; fi10 < 13; fi10++ {
                p0 -= v6
                fmt.Println(fi10, fi10)
                _ = fi10
            }
        default:
            if true {
                v6 = v6
                if (2 != 32767) {
                    var v11 string = "a"
                    if false {
                        switch -9223372036854775808 {
                        case 127:
                            for fi12 := 0; fi12 < 15; fi12++ {
                                var v13 string = v11
                                v14 := GenS1{F0: -1, F1: true, F2: 918}
                                var v15 int32 = 42
                                _ = fi12
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                        case 651:
                            for fi16 := 0; fi16 < 17; fi16++ {
                                var v17 int16 = 1
                                var v18 int8 = 96
                                v18 = (^(v18))
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        default:
                            v11 = "hello"
                        }
                        v6 = (v6 >> ((719) & 63))
                        fmt.Println(v11, p0, v11)
                    }
                    p0 = p0
                    _ = v11
                }
                fmt.Println(v6, p0)
            }
        }
    default:
        var v19 float32 = 0.0
        _ = v19
    }
    v20 := GenS1{F0: 32767, F1: (!false), F2: 1}
    fmt.Println(p0)
    _ = p0
    _ = v20
    _ = v6
    return -12
}
func genFn1(p0 uint8) uint32 {
    if false {
        for fi21 := 0; fi21 < 6; fi21++ {
            if ((42 < 1) && (!true)) {
                var v22 int32 = 1
                _ = v22
            }
            var v23 uint = 255
            v24 := []int{fi21}
            _ = fi21
            _ = v23
            _ = v24
        }
        v25 := []int{127, 9223372036854775807, 676, 412}
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 string = "hello"
        var v28 uint8 = p0
        switch v27 {
        case "a":
            v28 += 0
        case "\n\t":
            fmt.Println(v27, v27, v27)
        default:
            fmt.Println(fi26)
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = p0
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := []int{859, -937}
    var v4 string = ("foo bar" + (("" + ("foo bar" + "")) + ""))
    var v5 bool = true
    for fi29 := 0; fi29 < 5; fi29++ {
        fmt.Println(fi29, v4)
        _ = fi29
    }
    fmt.Println(v4)
    v5 = (42 < 128)
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := GenS1{F0: -32768, F1: v5, F2: 1}
            var v33 int8 = 42
            _ = fi31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 uint16 = 31
            _ = fi34
            _ = v35
        }
        _ = fi30
    }
    if (1 > 18446744073709551615) {
        v36 := GenS2{F0: 332}
        _ = len(v2)
        if (0 <= 341) {
            if len(v3) > 0 {
                v3[0] = -159
            }
        }
        _ = v36
    }
    switch -455 {
    case 32767:
        switch 484 {
        case 1:
            fmt.Println(v4, v5)
        }
    case 443:
        var v37 uint8 = 127
        _ = v37
    case 42:
        v38 := GenS1{F0: 103, F1: false, F2: 1}
        _ = v38
    default:
        var v39 string = v4
        _ = v39
    }
    var v40 uint = 1
    v40 -= v40
    fmt.Println(v5, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
