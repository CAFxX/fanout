package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int64, p1 int16, p2 string) int64 {
    for fi5 := 0; fi5 < 12; fi5++ {
        switch -645 {
        case 9223372036854775807:
            var v6 uint16 = 839
            _ = v6
        case 42:
            for fi7 := 0; fi7 < 20; fi7++ {
                p0 = 42
                _ = fi7
            }
        case 335:
            var v8 float64 = 1.5
            _ = v8
        default:
            v9 := GenS1{F0: "", F1: p2, F2: 65535}
            _ = v9
        }
        v10 := []int{510, 127, fi5}
        p2 = p2
        _ = fi5
        _ = v10
    }
    p0 = (p0 ^ p0)
    _ = p0
    _ = p1
    _ = p2
    return (p0 | p0)
}
func genFn1() uint8 {
    switch 9223372036854775807 {
    case 1:
        v11 := GenS2{F0: (((3.14159 - 0.0) + (-2.25 - 3.14159)) > (1.5 * 1.5))}
        _ = v11
    default:
        var v12 int32 = 140
        _ = v12
    }
    var v13 uint16 = 1
    v13 += v13
    if (((((-2.25 / float64((0.1 + 0.1))) * -2.25) * 1e10) * 0.0) <= float64(((-2.25 + (1e10 / (float32(((1.5 - 1e10) - 0.1)) - 1.5))) + -2.25))) {
        v13 += 403
        for fi14 := 0; fi14 < 5; fi14++ {
            v13 = (v13 | v13)
            _ = fi14
        }
        var v15 int = -684
        _ = v15
    } else {
        if (1000 >= 32767) {
            v13 = v13
            fmt.Println(v13)
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 int16 = -707
                if true {
                    var v18 int64 = 1
                    _ = v18
                } else {
                    v17 = v17
                }
                v19 := []int{(^(fi16)), fi16, (fi16 * 4294967295)}
                _ = fi16
                _ = v17
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 bool = (-1 != 13)
            switch (("foo bar" + (("foo bar" + (("hello" + "a") + ("hello" + "\n\t"))) + "hello")) + "a") {
            case "foo bar":
                v21 = (-254 < 42)
            case "a":
                v22 := GenS1{F0: ("" + ("\n\t" + "")), F1: ("foo bar" + "\n\t"), F2: (v13 / ((1) | 1))}
                _ = v22
            default:
                v21 = (65535 == 373)
            }
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v13)
    _ = v13
    return 127
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 127
    v3 := make(map[uint32]string)
    var v4 bool = true
    fmt.Println(v4)
    if v4 {
        _ = len(v1)
        switch 2147483647 {
        case -636:
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(v4, fi23, v4)
                var v24 float64 = ((0.1 * float64((((1.5 + 1.5) / 1e10) - (float32(3.14159) - 3.14159)))) * 1.5)
                _ = fi23
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 13; fi25++ {
                v2 += -2147483648
                _ = fi25
            }
        }
        v2 += (v2 % (((-(v2))) | 1))
    } else {
        switch "" {
        case "foo bar":
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 string = "hello"
                v28 := GenS1{F0: "hello", F1: ("foo bar" + v27), F2: 0}
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 uint16 = 255
                var v31 int32 = 84
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        var v32 string = ("foo bar" + "foo bar")
        for fi33 := 0; fi33 < 3; fi33++ {
            if v4 {
                fmt.Println(fi33)
            }
            _ = fi33
        }
        _ = v32
    }
    v2 = (v2 + v2)
    v34 := []int{(v2 << ((42) & 63))}
    v35 := GenS2{F0: v4}
    if true {
        if (((((float64(-2.25) + (float64(1e10) + float64((-2.25 * float32((float64(-2.25) / -2.25)))))) == (1.5 / 1e10)) && ((!v4) || v4)) || (2147483647 <= -1)) || (false || ((109 <= -56) && v4))) {
            var v36 uint64 = 151
            _ = v36
        } else {
            var v37 int64 = 698
            var v38 float64 = 1e10
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 14; fi39++ {
            v40 := GenS3{F0: 0}
            var v41 uint64 = 42
            v2 = v2
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    var v42 float32 = 3.14159
    var v43 string = ("a" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v43
}
