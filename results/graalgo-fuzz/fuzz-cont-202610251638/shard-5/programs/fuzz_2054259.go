package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int8; F2 int16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int8) uint64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: p0}
        switch "\n\t" {
        case "\n\t":
            v10 := GenS1{F0: 127}
            _ = v10
        case "hello":
            fmt.Println(fi8, fi8)
        case "foo bar":
            for fi11 := 0; fi11 < 18; fi11++ {
                p0 += 1
                p0 -= p0
                _ = fi11
            }
        }
        _ = fi8
        _ = v9
    }
    p0 = p0
    p0 = (p0 ^ p0)
    fmt.Println(p0)
    _ = p0
    return 2
}

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = (0.1 + ((1e10 - -2.25) / 3.14159))
    var v3 float32 = (float32(-2.25) - v2)
    v4 := make(map[string]int)
    var v5 uint64 = 4294967295
    v6 := make(map[uint32]string)
    v3 += (float32(((float64(1e10) / (1e10 * (1.5 * ((float64(float32(0.0)) - 0.0) - 3.14159)))) + (float64(1e10) * (1.5 - 0.0)))) * float32(1.5))
    var v12 float32 = 3.14159
    v3 -= ((v3 - 0.1) / 1.5)
    fmt.Println(v1, v12, v2)
    delete(v6, 384)
    v2 += v2
    if (122 != 42) {
        switch ("" + (((("foo bar" + "foo bar") + ("foo bar" + "hello")) + (("" + "\n\t") + (((("" + "hello") + "hello") + ("\n\t" + "foo bar")) + ("\n\t" + "\n\t")))) + "foo bar")) {
        case "hello":
            v3 = v2
        case "foo bar":
            fmt.Println(v2, v5)
        default:
            fmt.Println(v5, v3, v1)
        }
        if true {
            fmt.Println(v5)
            if false {
                v2 += 1e10
                v13 := GenS1{F0: 42}
                v14 := GenS1{F0: 112}
                _ = v13
                _ = v14
            } else {
                if ((-2147483648 > v1) || (-45 == -128)) {
                    v15 := GenS1{F0: -1}
                    fmt.Println(v12, v12, v5)
                    v16 := GenS1{F0: 102}
                    _ = v15
                    _ = v16
                } else {
                    v3 = (1.5 * v2)
                    fmt.Println(v12, v3, v2)
                }
            }
            v6[4294967295] = "a"
        }
        v2 = float32(0.1)
    } else {
        if (1e10 < v12) {
            var v17 int64 = 0
            v18 := GenS1{F0: -128}
            _ = v17
            _ = v18
        } else {
            fmt.Println(v1, v3)
            switch 128 {
            case -288:
                v19 := GenS2{F0: 255, F1: v5}
                _ = v19
            case 127:
                switch 1 {
                case 0:
                    var v20 int8 = 127
                    _ = v20
                case 32768:
                    v21 := []int{127, 9223372036854775807}
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 float64 = (((1e10 / float64(1.5)) * 1.5) + (float64((((v2 + v2) / ((0.0 + (v3 + v3)) - v2)) / (v12 - v12))) / 1.0))
                    fmt.Println(v3)
                    fmt.Println(v12)
                    _ = fi22
                    _ = v23
                }
            }
            var v24 int8 = 127
            _ = v24
        }
        var v25 int32 = 1
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 string = (("" + ("" + ("foo bar" + ("foo bar" + "\n\t")))) + "\n\t")
            for fi28 := 0; fi28 < 13; fi28++ {
                v25 -= (v1 / ((v1) | 1))
                v29 := GenS3{F0: (fi28 % ((fi28) | 1))}
                _ = fi28
                _ = v29
            }
            if true {
                fmt.Println(v5)
                v30 := GenS4{F0: v5, F1: 94, F2: -290}
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 string = "a"
                    v33 := GenS1{F0: -128}
                    v25 += (v25 / ((0) | 1))
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = v30
            }
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    var v34 uint = 255
    var v35 int64 = 894
    v36 := GenS5{F0: v5}
    var v37 int = -748
    var v38 int16 = -1
    v12 = v3
    for fi39 := 0; fi39 < 3; fi39++ {
        var v40 uint8 = 53
        v37 = v37
        _ = fi39
        _ = v40
    }
    v35 -= v35
    for fi41 := 0; fi41 < 3; fi41++ {
        v2 = 0.1
        _ = fi41
    }
    if (((v38 >> (((v34 * 719)) & 63)) == v38) || ((^(v5)) == (v5 * (v5 * (^(v5)))))) {
        if true {
            var v42 uint8 = 7
            for fi43 := 0; fi43 < 6; fi43++ {
                v12 = v2
                v37 += (v37 & (fi43 + 32767))
                v5 += 2
                _ = fi43
            }
            _ = v42
        } else {
            v44 := []int{v37, v37}
            fmt.Println(v12, v2)
            v35 = 9223372036854775807
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
