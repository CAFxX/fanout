package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 uint }
func genFn0(p0 int8, p1 uint8, p2 int8) uint32 {
    p0 = (p2 / (((p0 - -1)) | 1))
    fmt.Println(p1, p0, p0)
    p0 += p0
    fmt.Println(p1, p2, p0)
    var v7 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1() int64 {
    if false {
        v8 := GenS1{F0: (1 >= 1), F1: 4294967295, F2: 42}
        switch "a" {
        case "\n\t":
            var v9 bool = true
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint = 1
                _ = fi10
                _ = v11
            }
        case "a":
            for fi12 := 0; fi12 < 2; fi12++ {
                switch ("foo bar" + (("a" + "a") + "a")) {
                case "foo bar":
                    var v13 uint8 = 42
                    _ = v13
                case "":
                    if true {
                        var v14 int8 = 100
                        v14 -= (v14 + (v14 << ((681) & 63)))
                        _ = v14
                    } else {
                        v15 := GenS1{F0: (false && (!(!false))), F1: 1, F2: 385}
                        v16 := GenS1{F0: (((((3.14159 / (float64(-2.25) - 1e10)) - (1.5 / ((3.14159 * (float64(1.5) * 0.1)) - (1.5 + 3.14159)))) / (1.5 * -2.25)) / (3.14159 * (float64((float32(1e10) + ((1.5 + (float32(((-2.25 * (1e10 * 1.5)) * 3.14159)) * (0.0 + 0.0))) * ((((1.5 * 3.14159) * 1e10) * ((((1.5 + ((3.14159 - 1.5) * 3.14159)) * (1e10 + 1e10)) / 0.1) * 0.1)) / (0.1 + (1.5 + (1.5 + 1e10))))))) / ((1.5 - (0.0 * (0.0 * ((0.0 - -2.25) + float64(1.5))))) * -2.25)))) == 3.14159), F1: 354, F2: -862}
                        var v17 uint8 = 255
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                case "\n\t":
                    var v18 uint8 = 170
                    _ = v18
                default:
                    if (fi12 >= fi12) {
                        var v19 uint64 = 50
                        var v20 int8 = 127
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint8 = 1
                        v21 += (^(v21))
                        _ = v21
                    }
                }
                _ = fi12
            }
        }
        var v22 int64 = 1
        _ = v22
        _ = v8
    }
    var v23 float64 = 1e10
    var v24 int8 = -19
    _ = v23
    _ = v24
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = v1
    v3 := []string{"hello"}
    var v4 int = 1
    var v5 int8 = 116
    var v6 int8 = 88
    v5 = 127
    v4 = -863
    var v25 float32 = 3.14159
    for fi26 := 0; fi26 < 1; fi26++ {
        v25 += -2.25
        _ = fi26
    }
    if (true || false) {
        var v27 uint32 = 2147483648
        _ = v27
    }
    if true {
        var v28 int = (v4 * v4)
        v29 := GenS1{F0: true, F1: 42, F2: 0}
        v30 := GenS2{F0: genFn0((v6 | (v6 / ((v5) | 1))), 1, (v5 | v5))}
        _ = v28
        _ = v29
        _ = v30
    } else {
        var v31 string = "foo bar"
        _ = v31
    }
    if ((v4 + 32767) <= (v4 << ((296) & 63))) {
        if len(v3) > 0 {
            v3[0] = (("foo bar" + "foo bar") + "hello")
        }
        var v32 bool = (!true)
        _ = v32
    } else {
        v2 -= genFn1()
    }
    var v33 bool = true
    v5 += v5
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS3{F0: 18446744073709551615, F1: v33, F2: 159}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch (fi34 * v4) {
        case 386:
            v4 -= v4
        case -990:
            v3 = append(v3, "hello")
        default:
            v25 -= ((v25 * (v25 - ((v25 / v25) + v25))) - (-2.25 / (0.0 + float32((1.5 - 0.1)))))
        }
        _ = fi34
        _ = v35
    }
    switch "\n\t" {
    case "":
        _ = len(v3)
    case "hello":
        var v36 int8 = v5
        _ = v36
    case "\n\t":
        if v33 {
            if (0 >= 32767) {
                var v37 bool = true
                _ = v37
            }
        }
    }
    v25 += v25
    if v33 {
        fmt.Println(v4, v6, v6)
    } else {
        var v38 int16 = 0
        v38 = v38
        _ = v38
    }
    _ = len(v3)
    if v33 {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    } else {
        fmt.Println(v6, v4, v6)
        for fi39 := 0; fi39 < 13; fi39++ {
            for fi40 := 0; fi40 < 17; fi40++ {
                v41 := GenS2{F0: genFn0((v5 % ((v5) | 1)), 2, (^(v6)))}
                _ = fi40
                _ = v41
            }
            v2 += (^(v2))
            var v42 uint8 = 192
            _ = fi39
            _ = v42
        }
        var v43 int64 = genFn1()
        _ = v43
    }
    v25 -= (-2.25 * (-2.25 / 1.0))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
