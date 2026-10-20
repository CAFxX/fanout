package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 float64 = (((1e10 * 0.1) - 0.1) + float64(1.5))
    var v2 bool = (!false)
    v3 := make(map[uint32]bool)
    var v4 float32 = ((0.1 - (0.0 / 0.1)) - (-2.25 - 0.0))
    v4 = v4
    for fi5 := 0; fi5 < 17; fi5++ {
        switch (fi5 >> ((65535) & 63)) {
        case 4294967295:
            var v6 int32 = 1
            _ = v6
        case 848:
            if true {
                v7 := GenS1{F0: -126}
                _ = v7
            } else {
                delete(v3, 1)
            }
        }
        switch (fi5 & 116) {
        case 667:
            var v8 int = 9223372036854775807
            _ = v8
        case 32767:
            v3[1] = v2
        case 432:
            switch ((("" + "\n\t") + "\n\t") + "") {
            case "foo bar":
                for fi9 := 0; fi9 < 10; fi9++ {
                    v10 := v3[2147483647]
                    v11 := GenS2{F0: 1}
                    v1 += v1
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case "":
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := v3[4294967295]
                    switch fi12 {
                    case 127:
                        delete(v3, 65535)
                    case 128:
                        v13 = (("a" < (((("a" + (("" + "") + (("a" + "hello") + "foo bar"))) + "a") + ("a" + ("foo bar" + "\n\t"))) + "hello")) || v13)
                    }
                    _ = fi12
                    _ = v13
                }
            }
        default:
            v14 := GenS3{F0: 42}
            _ = v14
        }
        var v15 int32 = 2147483647
        _ = fi5
        _ = v15
    }
    var v16 int16 = 1
    switch ("a" + ("hello" + (("a" + "foo bar") + ("a" + "")))) {
    case "":
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := []int{}
            _ = fi17
            _ = v18
        }
    case "\n\t":
        for fi19 := 0; fi19 < 4; fi19++ {
            if (-2147483648 < -2147483648) {
                if (!v2) {
                    if true {
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v21 := GenS4{F0: 43, F1: 42, F2: v16}
                            if (!false) {
                                v2 = false
                                switch "hello" {
                                case "foo bar":
                                    var v22 bool = false
                                    _ = v22
                                case "a":
                                    var v23 float64 = v1
                                    _ = v23
                                default:
                                    v16 = -1
                                    v24 := []int{(fi19 | fi20)}
                                    var v25 float32 = v4
                                    _ = v24
                                    _ = v25
                                }
                                v4 += v4
                            } else {
                                if false {
                                    if (23 <= 9223372036854775807) {
                                        var v26 bool = ((42 <= -76) && v2)
                                        v2 = (943 > -203)
                                        var v27 string = "a"
                                        if ((fi19 | (fi19 | fi20)) <= fi19) {
                                            var v28 uint32 = 194
                                            v29 := GenS5{F0: 1, F1: (v28 << ((255) & 63)), F2: 255}
                                            _ = v28
                                            _ = v29
                                        }
                                        v30 := []int{(fi19 / ((fi19) | 1)), 9223372036854775807, fi20, -969}
                                        _ = v26
                                        _ = v27
                                        _ = v30
                                    }
                                }
                                fmt.Println(v4)
                            }
                            v31 := []int{}
                            switch "foo bar" {
                            case "hello":
                                var v32 uint64 = 256
                                _ = v32
                            case "\n\t":
                                if v2 {
                                    var v33 uint64 = 42
                                    _ = v33
                                } else {
                                    var v34 float64 = v1
                                    _ = v34
                                }
                            }
                            var v35 float32 = v4
                            _ = fi20
                            _ = v21
                            _ = v31
                            _ = v35
                        }
                        v2 = true
                        v36 := GenS1{F0: 0}
                        _ = v36
                    } else {
                        var v37 uint8 = 241
                        v1 += v1
                        var v38 float32 = 3.14159
                        _ = v37
                        _ = v38
                    }
                    if (v2 && (((-(fi19)) <= (fi19 % (((fi19 ^ (^(fi19)))) | 1))) || (2 == 1))) {
                        v16 = v16
                        fmt.Println(fi19, fi19, v2)
                    } else {
                        if true {
                            for fi39 := 0; fi39 < 14; fi39++ {
                                v4 -= float32(v1)
                                v40 := []int{42, fi19}
                                _ = fi39
                                _ = v40
                            }
                            v41 := GenS2{F0: 2}
                            _ = v41
                        }
                        v2 = v2
                        if (0 > 65535) {
                            v16 -= v16
                            v1 = 3.14159
                            for fi42 := 0; fi42 < 9; fi42++ {
                                var v43 string = (("" + ((("hello" + "\n\t") + "\n\t") + ("foo bar" + (("" + ("foo bar" + "")) + ("" + "hello"))))) + ("foo bar" + "\n\t"))
                                _ = fi42
                                _ = v43
                            }
                        }
                    }
                }
                v44 := []int{(fi19 & (fi19 / ((-312) | 1))), -1, 719, fi19}
                _ = v44
            } else {
                fmt.Println(v4)
                v45 := []int{fi19, fi19, fi19}
                _ = v45
            }
            _ = fi19
        }
    default:
        if v2 {
            v46 := GenS6{F0: -1}
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        var v48 int32 = -1
        v49 := GenS6{F0: 869}
        fmt.Println(fi47)
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println(v1, v16, v2)
    v50 := GenS7{F0: 0}
    var v51 int32 = 1
    for fi52 := 0; fi52 < 19; fi52++ {
        v4 += 1e10
        if (!(fi52 > fi52)) {
            v51 -= (v51 ^ v51)
            var v53 int = fi52
            delete(v3, 119)
            _ = v53
        }
        for fi54 := 0; fi54 < 3; fi54++ {
            v55 := GenS7{F0: 42}
            if v2 {
                fmt.Println(v16, v51)
            } else {
                var v56 int32 = v51
                _ = v56
            }
            _ = fi54
            _ = v55
        }
        _ = fi52
    }
    v1 = v1
    var v57 float32 = float32(float64(1.5))
    if (969 > 65535) {
        switch 128 {
        case 32767:
            fmt.Println(v1)
        case -9223372036854775808:
            for fi58 := 0; fi58 < 1; fi58++ {
                if v2 {
                    var v59 float32 = v4
                    _ = v59
                }
                if ("\n\t" < "hello") {
                    var v60 string = ((("" + "") + "hello") + "")
                    var v61 bool = (586 <= 770)
                    v62 := []int{(fi58 * (fi58 + 223)), (-(fi58)), fi58}
                    _ = v60
                    _ = v61
                    _ = v62
                }
                _ = fi58
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v57
}
