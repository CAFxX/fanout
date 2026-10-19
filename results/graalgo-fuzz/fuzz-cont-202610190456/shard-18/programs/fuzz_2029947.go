package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
func genFn0(p0 uint8, p1 uint8, p2 int32) string {
    p1 -= (p1 & (^(p1)))
    fmt.Println(p0)
    fmt.Println(p1, p0)
    var v7 float64 = (-2.25 / float64(float32(1.5)))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("foo bar" + "\n\t")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -605
    var v3 uint64 = 2
    var v4 int16 = 1
    v5 := make(map[int]string)
    var v6 int = 32768
    fmt.Println(v2, v3)
    if (!false) {
        if (!false) {
            fmt.Println(v6, v4)
            var v8 int32 = 2147483647
            _ = v8
        } else {
            var v9 uint16 = 255
            var v10 int16 = (v4 / ((v4) | 1))
            var v11 int = (v6 ^ v6)
            _ = v10
            _ = v11
            _ = v9
        }
    } else {
        v4 += -1
        var v12 int32 = 0
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 3; fi13++ {
            if ((2147483647 < 42) && false) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS1{F0: 0, F1: (^(v3))}
                    var v16 int16 = v4
                    v2 = (v2 + (v2 * v2))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 uint = 253
                _ = v17
            } else {
                v18 := []int{(v6 + -916), (-(v6))}
                var v19 int32 = -86
                if (v4 < (v4 * v4)) {
                    v19 = (-(v19))
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v21 := GenS2{F0: v19, F1: 128}
                        v3 -= 0
                        _ = fi20
                        _ = v21
                    }
                }
                _ = v18
                _ = v19
            }
            var v22 int = (fi13 % ((fi13) | 1))
            var v23 uint16 = 164
            _ = fi13
            _ = v22
            _ = v23
        }
    } else {
        var v24 uint = 0
        _ = v24
    }
    if false {
        v25 := GenS1{F0: 854, F1: (v3 ^ v3)}
        if (4294967295 != 622) {
            for fi26 := 0; fi26 < 18; fi26++ {
                switch -315 {
                case 127:
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 uint16 = 819
                        var v29 int8 = -15
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                case -669:
                    for fi30 := 0; fi30 < 15; fi30++ {
                        switch "a" {
                        case "":
                            v2 = 1
                        case "hello":
                            v4 = (v4 % ((v4) | 1))
                        case "\n\t":
                            v3 -= v3
                        default:
                            v31 := v5[(fi26 & (fi30 / ((fi30) | 1)))]
                            _ = v31
                        }
                        var v32 bool = (!((!(!true)) && false))
                        var v33 bool = ((1.5 >= 0.1) && v32)
                        _ = fi30
                        _ = v32
                        _ = v33
                    }
                case -9223372036854775808:
                    var v34 string = ""
                    _ = v34
                }
                _ = fi26
            }
            v1[0] = (175 >= 65535)
        }
        _ = v25
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        if (!true) {
            var v36 bool = true
            if v36 {
                fmt.Println(v6, v2)
                switch "a" {
                case "foo bar":
                    var v37 bool = ((v2 >> ((1) & 63)) <= v2)
                    _ = v37
                case "a":
                    for fi38 := 0; fi38 < 6; fi38++ {
                        v39 := v5[fi38]
                        _ = fi38
                        _ = v39
                    }
                case "hello":
                    fmt.Println(v36)
                }
                var v40 float32 = 0.1
                _ = v40
            }
            var v41 int8 = 42
            _ = v36
            _ = v41
        } else {
            var v42 int8 = 127
            var v43 uint16 = 410
            switch (fi35 | v6) {
            case 0:
                for fi44 := 0; fi44 < 5; fi44++ {
                    fmt.Println(v42)
                    v45 := v1[2147483648]
                    _ = fi44
                    _ = v45
                }
            case -2147483648:
                switch (("foo bar" + ("a" + ("hello" + "a"))) + "foo bar") {
                case "a":
                    v46 := GenS2{F0: 0, F1: 182}
                    _ = v46
                default:
                    var v47 string = (("\n\t" + "hello") + ((((("" + (("hello" + ("a" + ("\n\t" + ("" + "a")))) + (("" + ("foo bar" + ("\n\t" + "a"))) + (("hello" + "hello") + ((("a" + "hello") + "") + ""))))) + ("foo bar" + ("" + "foo bar"))) + ("foo bar" + ("foo bar" + "hello"))) + ("foo bar" + (("" + ("hello" + "\n\t")) + "\n\t"))) + "foo bar"))
                    _ = v47
                }
            default:
                var v48 uint16 = (v43 ^ 551)
                _ = v48
            }
            _ = v42
            _ = v43
        }
        if (false || (!false)) {
            fmt.Println(fi35)
            var v49 int64 = v2
            _ = v49
        }
        _ = fi35
    }
    for fi50 := 0; fi50 < 12; fi50++ {
        v51 := []int{fi50, -1, fi50}
        switch 128 {
        case 4294967295:
            var v52 float32 = ((float32(0.0) / 1e10) * 3.14159)
            _ = v52
        case 32767:
            v53 := GenS2{F0: 623, F1: 224}
            _ = v53
        }
        v3 = 277
        _ = fi50
        _ = v51
    }
    for fi54 := 0; fi54 < 6; fi54++ {
        var v55 uint64 = v3
        v56 := []int{}
        _ = fi54
        _ = v55
        _ = v56
    }
    v57 := []int{4294967295, 9223372036854775807, v6}
    if true {
        if true {
            if true {
                v2 = 0
                switch "hello" {
                case "\n\t":
                    _ = len(v57)
                case "":
                    v58 := GenS2{F0: 2147483647, F1: 42}
                    _ = v58
                default:
                    v59 := GenS3{F0: 164, F1: "", F2: -1}
                    _ = v59
                }
                v1[65535] = ("a" > "hello")
            }
            v60 := GenS3{F0: 615, F1: "", F2: 1000}
            if false {
                v61 := GenS1{F0: 1, F1: 0}
                _ = v61
            }
            _ = v60
        }
    } else {
        for fi62 := 0; fi62 < 2; fi62++ {
            switch "hello" {
            case "hello":
                v4 = (v4 ^ v4)
            case "":
                v3 -= 1
            default:
                v63 := GenS2{F0: -1, F1: 1}
                _ = v63
            }
            var v64 string = ("hello" + "foo bar")
            fmt.Println(v64)
            _ = fi62
            _ = v64
        }
    }
    switch -524 {
    case 0:
        v4 -= v4
    case 636:
        for fi65 := 0; fi65 < 13; fi65++ {
            if len(v57) > 0 {
                v57[0] = (-(v6))
            }
            _ = fi65
        }
    }
    v3 += v3
    if ((v6 >> ((273) & 63)) == (^(v6))) {
        fmt.Println(v4, v3)
        v66 := GenS3{F0: 2147483648, F1: "hello", F2: v4}
        _ = v66
    }
    var v67 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v67
}
