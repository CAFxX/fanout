package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 int64; F1 int; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint8) bool {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int32 = 470
        _ = fi6
        _ = v7
    }
    if false {
        var v8 int8 = -128
        fmt.Println(p1, p1, v8)
        _ = v8
    } else {
        if false {
            fmt.Println(p0)
            v9 := GenS1{F0: 1, F1: 255}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        p0 = (^(p0))
        if true {
            var v11 int64 = -24
            _ = v11
        }
        _ = fi10
    }
    var v12 int = -2147483648
    var v13 string = "a"
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return ((int(p1) >= (v12 ^ (v12 | int(p0)))) && true)
}

func main() {
    var v1 uint32 = 2
    v2 := []string{((("\n\t" + "\n\t") + "a") + "foo bar"), ("a" + (("\n\t" + "foo bar") + "\n\t"))}
    var v3 string = ("foo bar" + ("" + "hello"))
    v4 := make(map[int]bool)
    var v5 int32 = 2147483647
    v3 = "\n\t"
    if (true || false) {
        switch 32768 {
        case 1:
            v5 = -2147483648
        default:
            var v14 uint = 0
            _ = v14
        }
        switch ("foo bar" + ((("" + "\n\t") + (((v3 + v3) + "hello") + ((((v3 + v3) + v3) + "a") + v3))) + (v3 + v3))) {
        case "":
            for fi15 := 0; fi15 < 17; fi15++ {
                if false {
                    v16 := GenS2{F0: (-(v5))}
                    v2 = append(v2, "\n\t")
                    for fi17 := 0; fi17 < 6; fi17++ {
                        var v18 bool = (v1 > 980)
                        v18 = v18
                        v1 = v1
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                } else {
                    var v19 int = -399
                    _ = v19
                }
                _ = fi15
            }
        default:
            v5 = (v5 >> ((18446744073709551615) & 63))
        }
    } else {
        switch 32768 {
        case -1:
            v20 := GenS1{F0: 65535, F1: 18446744073709551615}
            _ = v20
        case 32768:
            v21 := GenS1{F0: 42, F1: 213}
            _ = v21
        }
        v3 = ("\n\t" + v3)
        var v22 int8 = 42
        _ = v22
    }
    var v23 uint8 = 29
    _ = len(v2)
    fmt.Println(v1, v3, v3)
    v5 = (v5 | -1)
    if (!true) {
        v24 := GenS1{F0: 0, F1: 0}
        v1 = (v1 | v1)
        if (!(1 < 0)) {
            var v25 uint8 = v23
            _ = v25
        } else {
            v5 = (v5 / ((422) | 1))
            v5 += 2147483647
            v26 := GenS3{F0: -157, F1: 0, F2: (("hello" + (v3 + ((v3 + ("foo bar" + v3)) + "\n\t"))) + v3)}
            _ = v26
        }
        _ = v24
    }
    v27 := GenS3{F0: 42, F1: (v23 * v23), F2: (v3 + (v3 + v3))}
    var v28 float64 = 3.14159
    for fi29 := 0; fi29 < 18; fi29++ {
        if (true && false) {
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 float32 = 1.5
                var v32 int32 = -1
                _ = fi30
                _ = v31
                _ = v32
            }
            delete(v4, 0)
            v1 -= v1
        } else {
            var v33 float32 = ((float32((-2.25 + v28)) + ((float32(v28) * (3.14159 * 0.1)) - (0.1 / 0.1))) * -2.25)
            _ = v33
        }
        v34 := GenS4{F0: 602, F1: 2147483647, F2: 0}
        _ = fi29
        _ = v34
    }
    v35 := GenS3{F0: -266, F1: v23, F2: v3}
    var v36 uint32 = v1
    if true {
        v37 := GenS5{F0: true}
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 int16 = -1
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    switch -2147483648 {
    case 127:
        for fi40 := 0; fi40 < 19; fi40++ {
            delete(v4, fi40)
            _ = fi40
        }
    case 840:
        v23 = (^(v23))
    default:
        v41 := GenS6{F0: false, F1: v5}
        _ = v41
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        v43 := GenS7{F0: 1}
        _ = fi42
        _ = v43
    }
    var v44 int64 = -403
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v5
}
