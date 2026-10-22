package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint16 }
func genFn0() int64 {
    var v7 uint32 = 2
    if (3.14159 >= ((((-2.25 - -2.25) + 0.1) + 0.0) + -2.25)) {
        v7 += 5
        v8 := GenS1{F0: 127}
        _ = v8
    } else {
        v9 := []int{9223372036854775807}
        v7 = 368
        _ = v9
    }
    if true {
        if true {
            var v10 uint16 = 1
            for fi11 := 0; fi11 < 16; fi11++ {
                v7 = v7
                _ = fi11
            }
            v7 = v7
            _ = v10
        }
    }
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint32, p1 int8) uint64 {
    var v12 string = (("\n\t" + "foo bar") + "hello")
    if true {
        if false {
            if false {
                var v13 uint64 = 18446744073709551615
                var v14 bool = (!false)
                v15 := GenS2{F0: int8(v13), F1: uint32(v13)}
                _ = v13
                _ = v14
                _ = v15
            } else {
                fmt.Println(v12)
            }
        } else {
            v16 := GenS1{F0: 1}
            _ = v16
        }
        fmt.Println(v12, p1, v12)
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int8 = p1
            var v19 int16 = -1
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    var v20 int32 = -591
    for fi21 := 0; fi21 < 11; fi21++ {
        if (-2.25 >= ((0.0 - -2.25) * 0.0)) {
            v20 += v20
            if ((1.5 != 1.5) && true) {
                v22 := GenS1{F0: (fi21 + fi21)}
                var v23 int16 = 1000
                _ = v22
                _ = v23
            } else {
                p0 -= p0
                fmt.Println(p0, p0)
                for fi24 := 0; fi24 < 15; fi24++ {
                    p1 = (p1 | (p1 + p1))
                    fmt.Println(fi21, v20, p1)
                    v25 := GenS1{F0: (fi24 >> ((65535) & 63))}
                    var v26 uint32 = (p0 | (p0 << ((479) & 63)))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            if true {
                v12 = (v12 + v12)
                v27 := []int{347, fi21, fi21, fi21}
                var v28 uint8 = 2
                _ = v27
                _ = v28
            } else {
                if true {
                    if (!false) {
                        switch fi21 {
                        case 4294967295:
                            v29 := []int{-982, fi21, fi21}
                            _ = v29
                        case 1:
                            var v30 int64 = 9223372036854775807
                            _ = v30
                        }
                    } else {
                        v31 := GenS3{F0: (v12 + v12), F1: 216}
                        _ = v31
                    }
                } else {
                    switch (fi21 + -2147483648) {
                    case 4294967295:
                        fmt.Println(fi21, p0)
                    }
                    switch "hello" {
                    case "":
                        for fi32 := 0; fi32 < 19; fi32++ {
                            var v33 bool = false
                            v34 := GenS3{F0: "foo bar", F1: 119}
                            var v35 uint16 = 0
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    case "\n\t":
                        v36 := []int{(fi21 - fi21), fi21}
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v38 := []int{0, fi37}
                        _ = fi37
                        _ = v38
                    }
                }
            }
        } else {
            v39 := []int{(fi21 << ((255) & 63)), (fi21 / ((4294967295) | 1))}
            var v40 float32 = -2.25
            if ((true || true) || ((p0 >> ((65535) & 63)) > p0)) {
                fmt.Println(fi21)
                switch 4294967295 {
                case -1:
                    var v41 uint64 = 584
                    _ = v41
                case 127:
                    fmt.Println(v12, p1)
                default:
                    for fi42 := 0; fi42 < 1; fi42++ {
                        var v43 int8 = 127
                        if true {
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                            p0 += p0
                        } else {
                            var v44 float64 = 1.5
                            var v45 bool = false
                            _ = v44
                            _ = v45
                        }
                        _ = fi42
                        _ = v43
                    }
                }
            }
            _ = v39
            _ = v40
        }
        var v46 int = (-(fi21))
        var v47 int16 = 251
        _ = fi21
        _ = v46
        _ = v47
    }
    var v48 int8 = 4
    _ = p0
    _ = p1
    _ = v12
    _ = v20
    _ = v48
    return 84
}

func main() {
    v1 := []int{128, 4294967295, -638, 127}
    var v2 uint = 255
    var v3 float64 = 0.1
    var v4 uint64 = 742
    var v5 uint32 = 2
    v6 := make(map[string]bool)
    v5 = (v5 ^ 1)
    for fi49 := 0; fi49 < 12; fi49++ {
        var v50 int16 = -32768
        v51 := GenS4{F0: (v2 | 1), F1: (!false)}
        _ = fi49
        _ = v50
        _ = v51
    }
    if (false || false) {
        v4 += (v4 % (((v4 / ((genFn1(v5, 85)) | 1))) | 1))
    } else {
        for fi52 := 0; fi52 < 8; fi52++ {
            var v53 uint = 837
            v5 = v5
            v4 = (^(v4))
            _ = fi52
            _ = v53
        }
        if false {
            v4 -= 0
            v5 -= 912
            switch "hello" {
            case "foo bar":
                v2 -= (^(v2))
            case "hello":
                var v54 uint64 = (v4 - v4)
                _ = v54
            default:
                _ = len(v6)
            }
        }
    }
    for fi55 := 0; fi55 < 3; fi55++ {
        v56 := GenS2{F0: 96, F1: 115}
        var v57 uint8 = 209
        _ = fi55
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        delete(v6, "\n\t")
        var v59 string = (("\n\t" + "foo bar") + ("foo bar" + "foo bar"))
        _ = fi58
        _ = v59
    }
    if true {
        v3 = ((1e10 - (v3 - 1e10)) * (float64(-2.25) / float64(float32(0.1))))
        v5 = v5
    }
    v4 -= v4
    var v60 float32 = 1.5
    switch (("hello" + "\n\t") + ("\n\t" + ("foo bar" + "\n\t"))) {
    case "hello":
        for fi61 := 0; fi61 < 12; fi61++ {
            var v62 float32 = v60
            _ = fi61
            _ = v62
        }
    case "foo bar":
        if (true && true) {
            for fi63 := 0; fi63 < 10; fi63++ {
                var v64 int64 = -800
                var v65 int16 = int16(v5)
                var v66 uint32 = v5
                _ = fi63
                _ = v64
                _ = v65
                _ = v66
            }
        }
    case "":
        if false {
            var v67 string = ""
            _ = v67
        } else {
            var v68 int64 = 1
            _ = v68
        }
    }
    var v69 int32 = 2147483647
    switch 32767 {
    case -648:
        var v70 uint = (v2 << (((v2 >> ((v2) & 63))) & 63))
        _ = v70
    case 32767:
        switch "\n\t" {
        case "":
            switch ((((("a" + (("foo bar" + ("\n\t" + "a")) + ("foo bar" + "foo bar"))) + ("hello" + (("hello" + "a") + "foo bar"))) + "\n\t") + ("\n\t" + "hello")) + "\n\t") {
            case "":
                v60 = v60
            }
        }
    }
    var v71 uint32 = v5
    _ = len(v1)
    fmt.Println(v71)
    if true {
        if ((1.5 - (v60 + (((-2.25 + v60) + v60) * 3.14159))) < v60) {
            for fi72 := 0; fi72 < 11; fi72++ {
                var v73 uint16 = 2
                _ = fi72
                _ = v73
            }
            switch (("foo bar" + ("a" + ("a" + "a"))) + "") {
            case "\n\t":
                v74 := GenS5{F0: uint16(v4)}
                _ = v74
            case "a":
                v3 -= v3
            case "foo bar":
                var v75 float64 = 0.1
                _ = v75
            }
            v3 = (v3 / 1.5)
        } else {
            v2 = (v2 % (((v2 * (v2 - v2))) | 1))
        }
    }
    var v76 uint32 = v5
    v76 = v71
    var v77 string = (("a" + "foo bar") + (("a" + "") + "a"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v69
    _ = v71
    _ = v76
    _ = v77
}
