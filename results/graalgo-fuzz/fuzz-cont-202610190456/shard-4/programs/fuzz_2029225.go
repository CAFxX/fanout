package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int8, p1 uint, p2 uint64) int16 {
    v6 := GenS1{F0: -232}
    var v7 uint16 = 1
    fmt.Println(p1, p1)
    for fi8 := 0; fi8 < 7; fi8++ {
        if (p2 != (p2 >> ((307) & 63))) {
            p1 = p1
            v9 := GenS2{F0: 2, F1: 4294967295, F2: (("" + "a") + "a")}
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS2{F0: 465, F1: 9223372036854775807, F2: "a"}
                var v12 int16 = 895
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := []int{fi13}
                var v15 int8 = -48
                var v16 uint16 = v7
                v7 = v16
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        if (32767 == 9223372036854775807) {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint32 = 42
                if true {
                    var v19 float32 = 1e10
                    var v20 int64 = -9223372036854775808
                    if true {
                        var v21 int = (-(fi17))
                        var v22 uint64 = 484
                        var v23 int8 = p0
                        fmt.Println(v22, v18, p1)
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        var v24 uint64 = p2
                        _ = v24
                    }
                    _ = v19
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
            var v25 float32 = 1.5
            _ = v25
        } else {
            switch "foo bar" {
            case "a":
                var v26 uint64 = (p2 % ((p2) | 1))
                _ = v26
            case "\n\t":
                v27 := GenS3{F0: 348, F1: p1, F2: -334}
                _ = v27
            default:
                v28 := []int{2147483647, int(p1), (-(fi8)), fi8}
                _ = v28
            }
        }
        _ = fi8
    }
    var v29 string = (("hello" + (("foo bar" + "hello") + (("" + ((("a" + ("hello" + ("hello" + ""))) + "a") + ("\n\t" + "hello"))) + ""))) + "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 0
    var v2 uint64 = 64
    var v3 float32 = 1e10
    v4 := []string{"foo bar", ("" + "\n\t")}
    v5 := make(map[int]uint32)
    v30 := GenS2{F0: 915, F1: 367, F2: "hello"}
    fmt.Println(v1, v2, v1)
    for fi31 := 0; fi31 < 19; fi31++ {
        v2 -= (v1 - v1)
        _ = fi31
    }
    v4 = append(v4, ("hello" + ((("hello" + "hello") + "foo bar") + (("" + "\n\t") + "\n\t"))))
    v3 += v3
    v32 := GenS1{F0: 42}
    v1 = (v2 & v2)
    for fi33 := 0; fi33 < 6; fi33++ {
        if (v2 >= (v1 % ((v2) | 1))) {
            v1 += (v2 + (v2 * (v1 | 2)))
            var v34 int = (fi33 << ((2) & 63))
            var v35 int8 = -1
            _ = v34
            _ = v35
        }
        if false {
            var v36 int32 = -1
            var v37 uint = uint(v2)
            _ = v36
            _ = v37
        } else {
            v1 = (^(v2))
            if false {
                v1 += (v2 ^ 513)
            } else {
                var v38 bool = (v3 != 1e10)
                _ = v38
            }
        }
        if false {
            _ = len(v4)
            switch "" {
            case "hello":
                switch fi33 {
                case 148:
                    var v39 int8 = 0
                    _ = v39
                case 2147483647:
                    var v40 int = -975
                    _ = v40
                }
            case "\n\t":
                var v41 int = 4294967295
                _ = v41
            default:
                fmt.Println(v1, v3, fi33)
            }
            var v42 uint16 = 2
            _ = v42
        }
        _ = fi33
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        v3 -= 3.14159
        _ = fi43
    }
    fmt.Println(v1)
    var v44 int32 = -800
    if ((!false) || (false && false)) {
        var v45 float32 = v3
        fmt.Println(v1, v3, v1)
        _ = v45
    } else {
        switch "foo bar" {
        case "a":
            var v46 uint64 = (v2 | v1)
            _ = v46
        case "":
            v47 := GenS4{F0: -9223372036854775808, F1: (((1e10 + (-2.25 / ((1e10 * 0.1) * 3.14159))) + -2.25) >= 0.1)}
            _ = v47
        }
        switch -9223372036854775808 {
        case 832:
            if false {
                v48 := GenS1{F0: 220}
                fmt.Println(v44)
                _ = v48
            }
        }
    }
    if ((0 < (-(v44))) || false) {
        v49 := GenS4{F0: 836, F1: (!(((47 > 244) && (false && (980 != 65535))) && ((0.1 / ((0.0 / float64(v3)) + 3.14159)) != 0.1)))}
        if len(v4) > 0 {
            v4[0] = "a"
        }
        _ = v49
    }
    v50 := v5[-916]
    _ = len(v4)
    v1 -= v1
    fmt.Println(v50, v50)
    v44 = (v44 << ((218) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v44
    _ = v5
    _ = v50
}
