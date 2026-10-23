package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 string; F2 int8 }
type GenS8 struct { F0 bool; F1 int64; F2 uint64 }
func genFn0() int64 {
    var v4 uint8 = 255
    var v5 float32 = 1e10
    if (!(!true)) {
        var v6 uint64 = 540
        switch (("a" + "\n\t") + "") {
        case "a":
            var v7 uint64 = (v6 >> ((255) & 63))
            _ = v7
        case "\n\t":
            v8 := []int{}
            _ = v8
        default:
            var v9 uint64 = v6
            _ = v9
        }
        var v10 int8 = -7
        _ = v10
        _ = v6
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            if (32767 != int16(fi12)) {
                if true {
                    v13 := []int{fi12}
                    v5 = v5
                    _ = v13
                }
            } else {
                switch (((("a" + "foo bar") + "a") + "\n\t") + "a") {
                case "foo bar":
                    if (1 < 267) {
                        switch fi12 {
                        case 2147483647:
                            fmt.Println(v5, v4, v5)
                        case 42:
                            if false {
                                v5 -= v5
                            } else {
                                fmt.Println(v4, v5, fi11)
                            }
                        default:
                            v5 = float32((3.14159 - 1e10))
                        }
                        if (1.5 != 1e10) {
                            var v14 uint8 = (v4 >> ((18446744073709551615) & 63))
                            if ((false && true) || (!true)) {
                                if false {
                                    var v15 string = ("a" + "")
                                    _ = v15
                                } else {
                                    v16 := GenS1{F0: -1, F1: 247}
                                    _ = v16
                                }
                            } else {
                                if (42 >= 500) {
                                    fmt.Println(v5, v5)
                                    var v17 float64 = (0.0 / float64(-2.25))
                                    v14 += (v14 / ((2) | 1))
                                    v4 = v4
                                    var v18 uint8 = v14
                                    _ = v17
                                    _ = v18
                                }
                                v19 := GenS2{F0: 405}
                                _ = v19
                            }
                            var v20 int = (-(fi11))
                            _ = v14
                            _ = v20
                        } else {
                            if (!(-110 <= -280)) {
                                v21 := []int{fi12, (fi12 << ((18446744073709551615) & 63))}
                                switch fi12 {
                                case 32768:
                                    v4 += (v4 - v4)
                                default:
                                    var v22 float64 = (float64((v5 + ((1e10 - ((float32(0.0) * float32((1e10 + 1e10))) * 3.14159)) * (3.14159 - v5)))) + 3.14159)
                                    var v23 int8 = -1
                                    switch "a" {
                                    case "a":
                                        _ = len(v21)
                                    default:
                                        v5 += (float32(-2.25) * 1e10)
                                    }
                                    _ = v22
                                    _ = v23
                                }
                                _ = v21
                            } else {
                                var v24 int16 = 0
                                _ = v24
                            }
                            if true {
                                fmt.Println(v4, fi12, v4)
                                v25 := []int{fi11, fi11, 1}
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    var v26 string = ("hello" + "hello")
                                    _ = v26
                                case "a":
                                    switch (fi12 % ((fi12) | 1)) {
                                    case -305:
                                        v5 -= ((v5 / v5) + -2.25)
                                    case -234:
                                        var v27 string = "foo bar"
                                        _ = v27
                                    default:
                                        v28 := GenS3{F0: (v4 ^ v4)}
                                        _ = v28
                                    }
                                default:
                                    v29 := GenS4{F0: 42}
                                    _ = v29
                                }
                                _ = v25
                            }
                        }
                        var v30 float64 = 1e10
                        _ = v30
                    }
                case "hello":
                    switch 1 {
                    case 1:
                        v5 += float32(3.14159)
                        if false {
                            fmt.Println(v5, fi12, fi12)
                        } else {
                            v31 := []int{(fi12 & 0), (-(fi11)), (-(fi11)), fi12}
                            _ = v31
                        }
                        v5 = (0.0 / v5)
                        v32 := GenS2{F0: 32767}
                        _ = v32
                    default:
                        v33 := GenS5{F0: v4, F1: 382}
                        _ = v33
                    }
                default:
                    if ((true && true) || false) {
                        v4 -= (^(v4))
                    } else {
                        v34 := GenS6{F0: 42}
                        v35 := []int{fi11}
                        var v36 uint64 = 18446744073709551615
                        var v37 uint16 = 255
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
                v4 -= 128
                fmt.Println(fi12)
            }
            v38 := GenS7{F0: 18446744073709551615, F1: "", F2: 42}
            _ = fi12
            _ = v38
        }
        for fi39 := 0; fi39 < 6; fi39++ {
            v40 := GenS2{F0: 267}
            v5 += 3.14159
            _ = fi39
            _ = v40
        }
        if false {
            for fi41 := 0; fi41 < 7; fi41++ {
                v42 := GenS2{F0: 255}
                var v43 uint16 = 126
                _ = fi41
                _ = v42
                _ = v43
            }
            for fi44 := 0; fi44 < 15; fi44++ {
                var v45 string = ("foo bar" + "a")
                v46 := GenS1{F0: 1, F1: 2}
                v47 := []int{fi11, 880, (fi11 * fi11)}
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
            }
        } else {
            v48 := GenS2{F0: 32767}
            var v49 bool = true
            _ = v48
            _ = v49
        }
        _ = fi11
    }
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 uint, p1 int) uint16 {
    switch (p1 % (((p1 * p1)) | 1)) {
    case 2147483647:
        p1 = (-(p1))
    default:
        v50 := GenS6{F0: -483}
        _ = v50
    }
    var v51 int32 = 448
    fmt.Println(p0, p1, p0)
    switch (p1 >> (((p0 << ((p0) & 63))) & 63)) {
    case 32767:
        for fi52 := 0; fi52 < 6; fi52++ {
            v53 := []int{fi52, (fi52 * (p1 ^ (fi52 % ((fi52) | 1)))), (p1 - -2), 127}
            _ = fi52
            _ = v53
        }
    case 4294967295:
        v54 := []int{p1}
        _ = v54
    }
    _ = p0
    _ = p1
    _ = v51
    return 196
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    fmt.Println(v1)
    fmt.Println(v1)
    v55 := GenS6{F0: 42}
    switch 710 {
    case 42:
        switch "a" {
        case "":
            switch ("" + "foo bar") {
            case "\n\t":
                for fi56 := 0; fi56 < 13; fi56++ {
                    if false {
                        for fi57 := 0; fi57 < 2; fi57++ {
                            v1 -= v1
                            v58 := GenS8{F0: (!(505 >= 1)), F1: 42, F2: 1}
                            var v59 uint16 = 255
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                    }
                    switch "" {
                    case "\n\t":
                        fmt.Println(fi56, v1)
                    case "a":
                        v60 := GenS5{F0: 0, F1: 0}
                        _ = v60
                    case "hello":
                        var v61 int16 = 401
                        _ = v61
                    }
                    _ = fi56
                }
            case "foo bar":
                v2[""] = ("" + ("hello" + "foo bar"))
            default:
                if false {
                    v62 := []int{416, 0, -679, 2147483647}
                    _ = v62
                }
            }
        case "\n\t":
            v1 += ((v1 - ((1.5 * ((float64(((1e10 - (((0.0 - 1e10) / -2.25) / 0.1)) / 0.1)) * v1) / (((1.5 - 0.0) - 0.0) / v1))) / 1.0)) * -2.25)
        default:
            var v63 float32 = 1e10
            _ = v63
        }
    }
    var v64 uint = 4294967295
    fmt.Println(v64, v64)
    for fi65 := 0; fi65 < 10; fi65++ {
        v64 = (v64 | (v64 % ((v64) | 1)))
        _ = fi65
    }
    var v66 int16 = -32768
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v64
    _ = v66
}
