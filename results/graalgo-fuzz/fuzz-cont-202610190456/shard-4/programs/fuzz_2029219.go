package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 uint16; F1 bool; F2 int8 }
type GenS3 struct { F0 int; F1 string; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 uint) bool {
    var v7 uint32 = 65535
    if true {
        var v8 int8 = 42
        _ = v8
    } else {
        v9 := GenS1{F0: ("\n\t" + "hello"), F1: (false && false), F2: 4294967295}
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 13; fi10++ {
            switch ("hello" + "hello") {
            case "":
                p0 = (p0 & (^(p0)))
            case "hello":
                v7 = (v7 & v7)
            case "foo bar":
                v11 := GenS2{F0: 32767, F1: ((true && (!true)) && ((870 < 2) && (((1.5 + 1.5) - (0.1 / 0.1)) == -2.25))), F2: -128}
                _ = v11
            default:
                var v12 string = ((("\n\t" + "") + "\n\t") + "a")
                _ = v12
            }
            _ = fi10
        }
        if true {
            fmt.Println(v7, v7)
            v13 := []int{}
            fmt.Println(v7)
            _ = v13
        } else {
            var v14 int8 = 127
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := []int{fi16, (fi16 * (fi15 + 9223372036854775807)), (-(fi16))}
                var v18 float32 = 3.14159
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := []int{42, (fi15 * (fi15 * (-(fi15)))), (fi15 + (fi15 / ((fi15) | 1))), -159}
            _ = fi15
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 string = ""
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v7
    return true
}
func genFn1(p0 uint64) uint {
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := []int{fi22}
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        p0 = (p0 ^ 963)
        var v25 uint = 42
        _ = fi24
        _ = v25
    }
    p0 -= 42
    fmt.Println(p0)
    p0 -= 42
    _ = p0
    return 1
}

func main() {
    var v1 bool = (true || (1 > 4294967295))
    var v2 int8 = 42
    var v3 int64 = 393
    v4 := []int{9223372036854775807, -464}
    v5 := make(map[uint32]string)
    var v6 float64 = float64(3.14159)
    v3 = (-(v3))
    v26 := GenS3{F0: 127, F1: "a", F2: 0}
    fmt.Println(v3, v6)
    var v27 float32 = 1e10
    fmt.Println(v2, v27, v6)
    var v28 uint16 = 456
    fmt.Println(v6)
    v28 += (v28 ^ v28)
    for fi29 := 0; fi29 < 2; fi29++ {
        v5[uint32(v28)] = ("hello" + "\n\t")
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 int = fi29
            switch "\n\t" {
            case "a":
                switch (fi29 >> ((255) & 63)) {
                case -96:
                    var v32 int8 = 1
                    _ = v32
                default:
                    v33 := GenS4{F0: (v2 + (v2 % ((v2) | 1)))}
                    _ = v33
                }
            case "foo bar":
                fmt.Println(fi30, v6)
            case "hello":
                v34 := GenS5{F0: fi30, F1: 772, F2: -1}
                _ = v34
            }
            _ = fi30
            _ = v31
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi29
    }
    v28 -= 65535
    v4 = append(v4, -9223372036854775808)
    if v1 {
        fmt.Println(v28)
        var v35 uint64 = 2
        _ = v35
    } else {
        var v36 string = "hello"
        _ = v36
    }
    v27 = (3.14159 / v27)
    fmt.Println(v3, v28, v6)
    if false {
        for fi37 := 0; fi37 < 8; fi37++ {
            if v1 {
                if len(v4) > 0 {
                    v4[0] = fi37
                }
            } else {
                _ = len(v5)
            }
            v38 := v5[148]
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 9; fi39++ {
            v3 -= (v3 ^ v3)
            v3 -= v3
            switch 873 {
            case 1:
                v27 = v27
            }
            _ = fi39
        }
        var v40 int32 = -1
        _ = v40
    } else {
        v2 += v2
        switch (("hello" + "foo bar") + ((("foo bar" + ("hello" + (("foo bar" + "hello") + "a"))) + "hello") + "\n\t")) {
        case "":
            for fi41 := 0; fi41 < 17; fi41++ {
                for fi42 := 0; fi42 < 6; fi42++ {
                    v28 += (v28 & (v28 >> ((0) & 63)))
                    fmt.Println(v3, v2)
                    _ = fi42
                }
                _ = fi41
            }
        case "hello":
            var v43 float64 = v6
            _ = v43
        case "\n\t":
            for fi44 := 0; fi44 < 18; fi44++ {
                fmt.Println(v3)
                _ = fi44
            }
        default:
            var v45 int = -1
            _ = v45
        }
        var v46 int8 = (v2 ^ v2)
        _ = v46
    }
    delete(v5, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
