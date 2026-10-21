package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int; F1 bool; F2 bool }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint, p2 uint16) int16 {
    var v7 string = "foo bar"
    v8 := GenS1{F0: (p0 | (p0 - (p0 >> ((355) & 63)))), F1: 2, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32, p1 uint) int {
    var v9 string = "\n\t"
    v10 := GenS1{F0: 21, F1: 65535, F2: 381}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -257
}

func main() {
    v1 := []int{953}
    v2 := make(map[string]int)
    v3 := []int{-872, 586, 4294967295}
    var v4 string = (("hello" + "") + ("foo bar" + ("foo bar" + ("a" + "foo bar"))))
    v5 := []int{-2147483648}
    var v6 int32 = 42
    if (true && (999 <= 2147483647)) {
        if (2 == 0) {
            var v11 uint8 = 2
            fmt.Println(v11, v4)
            switch genFn1(2, 1) {
            case 549:
                fmt.Println(v6)
            case 32768:
                v12 := v2[(((((v4 + ("" + (("hello" + "hello") + ""))) + v4) + "a") + v4) + v4)]
                _ = v12
            default:
                fmt.Println(v11)
            }
            _ = v11
        } else {
            v13 := GenS2{F0: (v4 + v4), F1: 255, F2: 42}
            fmt.Println(v4)
            var v14 int16 = 1
            _ = v13
            _ = v14
        }
        v6 -= (v6 | (-(v6)))
    } else {
        var v15 float32 = (-2.25 - (1.5 - (0.1 / 3.14159)))
        switch v4 {
        case "a":
            fmt.Println(v6)
        case "":
            v16 := GenS2{F0: "foo bar", F1: 2, F2: 268}
            _ = v16
        }
        _ = v15
    }
    v4 = ("a" + (v4 + v4))
    for fi17 := 0; fi17 < 7; fi17++ {
        if false {
            for fi18 := 0; fi18 < 6; fi18++ {
                v4 = v4
                switch fi18 {
                case -1:
                    switch (("\n\t" + "foo bar") + "") {
                    case "foo bar":
                        if len(v5) > 0 {
                            v5[0] = (^(fi18))
                        }
                    case "a":
                        v19 := GenS3{F0: 377, F1: 119}
                        _ = v19
                    case "":
                        var v20 uint = 670
                        _ = v20
                    default:
                        var v21 bool = (0 != 1)
                        _ = v21
                    }
                default:
                    v6 -= -1
                }
                if (false && (true && (true && (v6 != -2147483648)))) {
                    v4 = v4
                }
                _ = fi18
            }
            fmt.Println(v6, v4)
        } else {
            delete(v2, v4)
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            _ = len(v2)
            if false {
                var v23 float32 = 0.1
                var v24 int16 = 0
                v23 += (3.14159 / v23)
                _ = v23
                _ = v24
            } else {
                fmt.Println(v4, v4)
            }
            var v25 uint8 = 127
            _ = fi22
            _ = v25
        }
        _ = fi17
    }
    var v26 int16 = genFn0(0, 0, 2)
    if false {
        var v27 float32 = (1.5 - (0.0 + -2.25))
        var v28 bool = true
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := v2["a"]
        v31 := GenS3{F0: 176, F1: 83}
        var v32 uint64 = 4294967295
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    v33 := GenS4{F0: 520, F1: (213 <= 65535), F2: true}
    v6 = v6
    v34 := v2[v4]
    for fi35 := 0; fi35 < 14; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            switch v4 {
            case "":
                if (true || true) {
                    fmt.Println(v26)
                    if ((v6 > 42) && false) {
                        delete(v2, "\n\t")
                        var v37 string = ""
                        _ = v37
                    } else {
                        var v38 float64 = 0.1
                        v26 += 42
                        _ = v38
                    }
                }
            case "\n\t":
                v39 := GenS3{F0: 0, F1: 1}
                _ = v39
            default:
                switch "a" {
                case "a":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "":
                    var v40 int8 = 55
                    _ = v40
                case "\n\t":
                    if false {
                        v6 = -96
                        fmt.Println(v34, v26, v4)
                    }
                default:
                    v26 = -1
                }
            }
            v41 := GenS5{F0: -387, F1: 42}
            var v42 uint8 = 217
            fmt.Println(v6, fi36, v4)
            _ = fi36
            _ = v41
            _ = v42
        }
        _ = fi35
    }
    if (!(0 <= 32767)) {
        var v43 int = v34
        _ = v43
    }
    v44 := GenS6{F0: 1}
    var v45 int32 = (v6 & v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
