package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int; F1 uint; F2 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32; F1 int; F2 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8, p1 uint8) bool {
    var v7 string = ("a" + ("\n\t" + "hello"))
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: p1}
        v10 := GenS2{F0: p1, F1: 255}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 float32 = float32(0.0)
        var v13 float32 = ((0.0 / ((((v12 * (v12 - 3.14159)) - v12) - -2.25) - v12)) - (v12 + v12))
        var v14 int64 = 794
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := []int{-31, fi15}
        var v17 float32 = 0.1
        var v18 uint32 = 2
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v7
    return (v7 >= v7)
}
func genFn1(p0 string, p1 uint16, p2 int) uint32 {
    var v19 int16 = -32768
    switch (p2 / ((-1) | 1)) {
    case 42:
        fmt.Println(p1, p2, p0)
    }
    switch p2 {
    case 1:
        var v20 int32 = 2147483647
        _ = v20
    case 868:
        for fi21 := 0; fi21 < 4; fi21++ {
            p1 = (^(p1))
            _ = fi21
        }
    default:
        v22 := GenS2{F0: 42, F1: 255}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 1
}

func main() {
    var v1 int = 495
    v2 := []int{v1, -2147483648}
    var v3 int16 = 32767
    var v4 float32 = 1e10
    var v5 uint64 = 4294967295
    var v6 int64 = 42
    v4 = (3.14159 + 0.0)
    fmt.Println(v5, v5)
    var v23 uint16 = 2
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := GenS3{F0: uint32(v23)}
            var v27 int64 = 42
            var v28 float64 = 0.1
            if len(v2) > 0 {
                v2[0] = (v1 - (v1 % ((32768) | 1)))
            }
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    var v29 int8 = 127
    v4 -= 1e10
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 uint8 = 127
        _ = fi30
        _ = v31
    }
    if true {
        v29 -= -128
        var v32 int8 = int8(v1)
        _ = v32
    } else {
        fmt.Println(v23, v23, v29)
    }
    switch "a" {
    case "a":
        switch (("a" + ("hello" + ("" + ("a" + ("hello" + (("\n\t" + ("" + (("" + "") + "hello"))) + "a")))))) + ("hello" + "a")) {
        case "foo bar":
            v33 := GenS4{F0: ""}
            _ = v33
        case "hello":
            v6 = v6
        case "a":
            for fi34 := 0; fi34 < 10; fi34++ {
                v35 := GenS5{F0: 184, F1: (^(v5))}
                _ = fi34
                _ = v35
            }
        default:
            v36 := GenS6{F0: v1, F1: v5}
            _ = v36
        }
    case "\n\t":
        switch "foo bar" {
        case "foo bar":
            switch (v1 % ((-2147483648) | 1)) {
            case 32767:
                v37 := GenS7{F0: 128, F1: 255, F2: ("a" + "")}
                _ = v37
            case -2147483648:
                v2 = append(v2, 32767)
            default:
                var v38 uint32 = uint32(v23)
                _ = v38
            }
        case "":
            switch ((("foo bar" + "foo bar") + "a") + ("" + "hello")) {
            case "a":
                if true {
                    v39 := GenS3{F0: 2147483647}
                    v40 := GenS8{F0: 2}
                    _ = v39
                    _ = v40
                }
            case "foo bar":
                var v41 int64 = (v6 / ((v6) | 1))
                _ = v41
            case "\n\t":
                switch "" {
                case "\n\t":
                    v6 = v6
                }
            default:
                for fi42 := 0; fi42 < 14; fi42++ {
                    for fi43 := 0; fi43 < 11; fi43++ {
                        switch fi42 {
                        case 128:
                            v4 = ((v4 - v4) / (0.1 - (v4 * v4)))
                        case 127:
                            var v44 uint16 = (v23 / ((v23) | 1))
                            _ = v44
                        default:
                            var v45 int16 = -114
                            switch int(v6) {
                            case 128:
                                var v46 string = ("a" + "\n\t")
                                _ = v46
                            default:
                                var v47 int32 = -1
                                _ = v47
                            }
                            var v48 bool = false
                            _ = v45
                            _ = v48
                        }
                        var v49 uint32 = 42
                        _ = fi43
                        _ = v49
                    }
                    v50 := GenS9{F0: 2147483647}
                    if false {
                        v51 := GenS10{F0: 42, F1: (v1 % (((v1 | fi42)) | 1)), F2: 18446744073709551615}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v51
                    } else {
                        var v52 float64 = 0.1
                        if ((true && true) && ((v5 << ((0) & 63)) <= v5)) {
                            var v53 string = "\n\t"
                            var v54 uint32 = 2
                            _ = v53
                            _ = v54
                        }
                        _ = v52
                    }
                    _ = fi42
                    _ = v50
                }
            }
        case "a":
            v4 = (3.14159 * (-2.25 * (float32(1e10) + v4)))
        default:
            v23 += (v23 - v23)
        }
    default:
        if true {
            fmt.Println(v1)
            fmt.Println(v29, v1)
            for fi55 := 0; fi55 < 7; fi55++ {
                v1 = 32767
                _ = fi55
            }
        }
    }
    v56 := GenS11{F0: v6}
    v57 := GenS12{F0: 497, F1: (v6 ^ (^(v6)))}
    v5 += (v5 | 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
