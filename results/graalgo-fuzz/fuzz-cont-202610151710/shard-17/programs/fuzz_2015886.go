package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
func genFn0() int32 {
    var v6 uint = 42
    v7 := GenS1{F0: 9223372036854775807, F1: 127}
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint32, p1 uint16, p2 int) uint64 {
    p2 = p2
    p1 = (p1 + (p1 + p1))
    v8 := GenS1{F0: 42, F1: -104}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}

func main() {
    var v1 float32 = ((0.1 / (1e10 + float32(1e10))) * ((((0.1 - (float32((3.14159 * -2.25)) + (0.0 - ((0.1 * 0.1) - 0.0)))) * 3.14159) - (-2.25 * 3.14159)) / 1.5))
    v2 := []uint32{2147483647, 316}
    v3 := make(map[string]string)
    var v4 uint16 = 631
    v5 := make(map[string]bool)
    var v9 int32 = -538
    var v10 bool = (!false)
    switch -2147483648 {
    case 2147483647:
        fmt.Println(v10)
    case 1:
        v4 += (v4 | v4)
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch (("a" + "hello") + "\n\t") {
        case "a":
            v12 := GenS1{F0: 9223372036854775807, F1: 110}
            _ = v12
        case "hello":
            v13 := v5[""]
            _ = v13
        default:
            var v14 uint64 = 42
            _ = v14
        }
        _ = fi11
    }
    switch "a" {
    case "":
        for fi15 := 0; fi15 < 18; fi15++ {
            switch (("foo bar" + "\n\t") + "foo bar") {
            case "foo bar":
                var v16 int32 = v9
                _ = v16
            default:
                v4 = (^(v4))
            }
            var v17 int8 = -1
            for fi18 := 0; fi18 < 6; fi18++ {
                if v10 {
                    if ((0.1 <= -2.25) && (-9223372036854775808 != 9223372036854775807)) {
                        v10 = ((^(v17)) == v17)
                    }
                    if (4294967295 <= 0) {
                        v9 = genFn0()
                    } else {
                        if true {
                            _ = len(v5)
                            var v19 uint8 = 69
                            v20 := GenS2{F0: v9, F1: "hello", F2: 4294967295}
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v10, v10, v10)
                        }
                    }
                    fmt.Println(fi15)
                } else {
                    v17 = v17
                }
                _ = len(v5)
                _ = fi18
            }
            _ = fi15
            _ = v17
        }
    case "a":
        if ((v4 >> ((726) & 63)) < (v4 / ((v4) | 1))) {
            v1 += v1
            v10 = true
        }
    default:
        var v21 int16 = 1
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 float32 = ((3.14159 * (v1 - v1)) / (((v1 - v1) / (float32((3.14159 / (0.1 / float64(v1)))) - float32(0.0))) + float32(1e10)))
        _ = fi22
        _ = v23
    }
    v1 = v1
    v4 = (^(v4))
    var v24 float32 = v1
    v25 := GenS2{F0: (v9 & (v9 & (v9 - genFn0()))), F1: "a", F2: 0}
    v4 -= v4
    fmt.Println(v9, v10)
    v4 += (^(v4))
    v9 = (v9 & (^(v9)))
    if len(v2) > 0 {
        v2[0] = 2
    }
    v2 = append(v2, 523)
    if true {
        for fi26 := 0; fi26 < 12; fi26++ {
            v5["hello"] = true
            _ = fi26
        }
        delete(v3, "")
        v27 := v3["a"]
        _ = v27
    } else {
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "\n\t":
            var v28 int = 4294967295
            _ = v28
        case "hello":
            if (!(v10 || (-355 != -109))) {
                var v29 uint32 = 42
                v4 = (v4 - (^(v4)))
                switch 629 {
                case 32767:
                    v9 += v9
                }
                _ = v29
            } else {
                v9 = (-(v9))
            }
        }
        v30 := v3[""]
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
