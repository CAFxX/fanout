package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8; F2 uint }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint8; F2 int32 }
func genFn0() bool {
    v6 := []int{}
    if true {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: (fi7 - (fi7 & 2147483647))}
            var v9 uint = 2
            fmt.Println(v9)
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    switch "" {
    case "":
        if ((0 <= 0) || (("" + ("foo bar" + "foo bar")) == "a")) {
            for fi10 := 0; fi10 < 10; fi10++ {
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 int8 = 1
                    v6 = append(v6, fi10)
                    var v13 int8 = (v12 | 0)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
        }
    case "hello":
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 string = "\n\t"
            _ = fi14
            _ = v15
        }
    default:
        var v16 float32 = (float32(-2.25) / 1.0)
        _ = v16
    }
    _ = v6
    return true
}

func main() {
    v1 := []uint32{0, 65535, 4294967295, 179}
    var v2 uint = 18446744073709551615
    v3 := []int{32768, 4294967295}
    v4 := []uint32{2147483647}
    var v5 uint = 231
    if (false || (false && true)) {
        v5 = v5
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if (1 <= 434) {
            if (-217 < 2147483647) {
                var v18 uint8 = 48
                _ = v18
            }
            var v19 uint32 = 1
            if false {
                var v20 uint8 = 0
                var v21 float32 = 3.14159
                _ = v20
                _ = v21
            }
            _ = v19
        } else {
            for fi22 := 0; fi22 < 18; fi22++ {
                v5 = (^(v2))
                var v23 uint16 = 0
                v23 -= 0
                v24 := GenS2{F0: "foo bar", F1: int8(v2), F2: 776}
                var v25 uint16 = (v23 + v23)
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            v3 = append(v3, fi17)
            fmt.Println(fi26, v5, v2)
            _ = fi26
        }
        var v27 float32 = 1e10
        _ = fi17
        _ = v27
    }
    fmt.Println(v5)
    v4 = append(v4, 42)
    switch "" {
    case "\n\t":
        if false {
            fmt.Println(v2)
            var v28 int64 = 0
            var v29 string = ("foo bar" + "")
            _ = v28
            _ = v29
        } else {
            var v30 int8 = 100
            _ = v30
        }
    case "hello":
        v2 += 780
    }
    v5 -= v2
    fmt.Println(v5, v2)
    for fi31 := 0; fi31 < 10; fi31++ {
        v32 := GenS2{F0: "\n\t", F1: 127, F2: v2}
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 int = fi33
            if false {
                v5 = 2
                var v35 string = "hello"
                _ = v35
            } else {
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
            var v36 uint16 = 32767
            _ = fi33
            _ = v34
            _ = v36
        }
        _ = fi31
        _ = v32
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        switch (-(fi37)) {
        case 128:
            if (0.1 >= 1.5) {
                v5 = (v5 / ((v5) | 1))
                v5 -= 255
                v38 := GenS3{F0: 255, F1: 898}
                _ = v38
            }
        case 831:
            fmt.Println(v2)
        case 0:
            v39 := GenS2{F0: (((("hello" + "foo bar") + "") + "\n\t") + "a"), F1: -71, F2: v2}
            _ = v39
        default:
            fmt.Println(v2, fi37, v5)
        }
        if true {
            var v40 int8 = 127
            _ = len(v4)
            v5 = 572
            _ = v40
        } else {
            var v41 string = "foo bar"
            _ = v41
        }
        switch fi37 {
        case -9223372036854775808:
            if true {
                v5 = v5
                var v42 uint16 = 65535
                _ = v42
            } else {
                for fi43 := 0; fi43 < 12; fi43++ {
                    var v44 bool = true
                    v5 = v2
                    _ = fi43
                    _ = v44
                }
            }
        default:
            v45 := GenS2{F0: ("foo bar" + ""), F1: 100, F2: (v5 + v2)}
            _ = v45
        }
        _ = fi37
    }
    var v46 uint = 0
    if (0.1 <= (1.5 - float64(0.1))) {
        if (997 < -2147483648) {
            for fi47 := 0; fi47 < 11; fi47++ {
                v48 := GenS4{F0: 42, F1: 0, F2: -2147483648}
                _ = fi47
                _ = v48
            }
            var v49 int8 = 127
            var v50 uint32 = uint32(v46)
            _ = v49
            _ = v50
        }
        if false {
            switch "" {
            case "hello":
                var v51 uint16 = 65535
                _ = v51
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
            }
            if true {
                var v52 int64 = -9223372036854775808
                _ = len(v3)
                fmt.Println(v46, v2)
                _ = v52
            }
            v53 := GenS3{F0: 433, F1: 18446744073709551615}
            _ = v53
        } else {
            for fi54 := 0; fi54 < 15; fi54++ {
                v46 += (v46 | 903)
                var v55 int16 = 1
                v46 = 4294967295
                _ = fi54
                _ = v55
            }
            v56 := GenS1{F0: 1}
            for fi57 := 0; fi57 < 8; fi57++ {
                var v58 int16 = 1000
                _ = fi57
                _ = v58
            }
            _ = v56
        }
    } else {
        fmt.Println(v2, v46, v2)
        v5 -= (v2 ^ v46)
        v5 = v46
    }
    v3 = append(v3, -2147483648)
    v5 = v5
    v46 -= 2
    v46 += (v2 ^ 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
