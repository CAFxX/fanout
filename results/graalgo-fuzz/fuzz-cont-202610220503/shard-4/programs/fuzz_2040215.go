package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 bool }
func genFn0(p0 uint, p1 uint32, p2 uint16) uint {
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{32767}
        p0 = p0
        fmt.Println(p0, p1)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 + (^(p0)))
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{}
    v3 := make(map[int]bool)
    var v4 uint16 = 255
    switch v1 {
    case "\n\t":
        var v8 int32 = -2147483648
        _ = v8
    case "hello":
        if len(v2) > 0 {
            v2[0] = 128
        }
    default:
        fmt.Println(v1, v4)
    }
    v9 := GenS1{F0: 32767}
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS2{F0: 0, F1: 879, F2: true}
        fmt.Println(v4, v4, v4)
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            fmt.Println(v4)
            switch -1 {
            case 32768:
                switch -1 {
                case 2147483647:
                    if (2 == 179) {
                        v1 = (((v1 + ((v1 + v1) + "foo bar")) + (((("" + "") + ((("hello" + v1) + v1) + (v1 + v1))) + v1) + "")) + (("\n\t" + "hello") + ("\n\t" + "")))
                        var v12 int8 = -25
                        _ = v12
                    }
                }
            default:
                var v13 int8 = 100
                _ = v13
            }
        }
    case "hello":
        switch 42 {
        case 42:
            if (2147483648 == 2) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi14
                }
                v4 -= (v4 + v4)
            } else {
                if true {
                    switch -208 {
                    case 32768:
                        var v15 string = ""
                        _ = v15
                    }
                } else {
                    var v16 uint16 = v4
                    _ = v16
                }
                fmt.Println(v1, v4)
            }
        case 32768:
            if true {
                var v17 uint = 327
                _ = v17
            }
        case -9223372036854775808:
            v18 := GenS3{F0: true, F1: 1}
            _ = v18
        default:
            for fi19 := 0; fi19 < 20; fi19++ {
                v20 := GenS1{F0: 32767}
                _ = fi19
                _ = v20
            }
        }
    case "a":
        v21 := GenS1{F0: 32767}
        _ = v21
    }
    var v22 uint64 = 0
    var v23 int64 = 1
    fmt.Println(v23)
    var v24 string = ""
    v25 := GenS4{F0: 2}
    switch "" {
    case "\n\t":
        v26 := GenS5{F0: v4, F1: (v1 + v24)}
        _ = v26
    case "hello":
        var v27 int8 = 40
        _ = v27
    case "":
        v24 = "hello"
    default:
        switch v1 {
        case "hello":
            v23 = -9223372036854775808
        case "":
            var v28 uint32 = 786
            _ = v28
        default:
            for fi29 := 0; fi29 < 10; fi29++ {
                switch 9223372036854775807 {
                case 136:
                    var v30 float32 = 1.5
                    _ = v30
                default:
                    v31 := GenS4{F0: 1}
                    _ = v31
                }
                _ = fi29
            }
        }
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        for fi33 := 0; fi33 < 8; fi33++ {
            if false {
                switch fi33 {
                case -768:
                    switch (fi32 & fi33) {
                    case -984:
                        var v34 int64 = (v23 * v23)
                        _ = v34
                    case 32767:
                        v23 += v23
                    }
                }
                v35 := GenS6{F0: false}
                v36 := GenS2{F0: (^(v22)), F1: 255, F2: (((0.0 / (1.5 * (float64(0.1) + ((-2.25 + 0.1) - 1.5)))) >= float64((1.5 + 0.1))) || (false && false))}
                _ = v35
                _ = v36
            }
            _ = fi33
        }
        switch v24 {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = (fi32 | fi32)
            }
        case "":
            v2 = append(v2, (fi32 ^ (-(fi32))))
        case "hello":
            for fi37 := 0; fi37 < 3; fi37++ {
                v24 = v24
                _ = fi37
            }
        default:
            v1 = v24
        }
        if (true || (false || (128 != 1))) {
            var v38 string = v1
            _ = v38
        } else {
            var v39 uint32 = 65535
            if (!true) {
                v40 := GenS6{F0: (!((v22 != 1) && (int(v39) == fi32)))}
                v1 = v1
                _ = v40
            }
            _ = v39
        }
        _ = fi32
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v42 := GenS5{F0: (v4 / ((703) | 1)), F1: "\n\t"}
        switch fi41 {
        case 32768:
            if len(v2) > 0 {
                v2[0] = fi41
            }
        case 402:
            v43 := GenS4{F0: 742}
            _ = v43
        default:
            if (((1e10 / (1.5 / 3.14159)) + 1e10) == (3.14159 - 0.1)) {
                var v44 float64 = (1.5 / float64(-2.25))
                v22 -= v22
                _ = v44
            }
        }
        _ = fi41
        _ = v42
    }
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
