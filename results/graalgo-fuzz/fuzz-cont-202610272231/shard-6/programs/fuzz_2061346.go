package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 string }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 int32, p2 int32) uint64 {
    switch p0 {
    case "a":
        p0 = "foo bar"
    case "foo bar":
        p0 = "\n\t"
    case "":
        for fi5 := 0; fi5 < 14; fi5++ {
            if false {
                v6 := []int{572}
                p1 -= (p2 / ((p1) | 1))
                _ = v6
            } else {
                switch (p0 + p0) {
                case "hello":
                    p0 = (p0 + (p0 + p0))
                case "\n\t":
                    v7 := GenS1{F0: 9223372036854775807, F1: 1, F2: (p0 + p0)}
                    _ = v7
                case "":
                    v8 := GenS1{F0: 1, F1: 1, F2: (p0 + p0)}
                    _ = v8
                }
                v9 := GenS1{F0: 568, F1: -1, F2: "hello"}
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS1{F0: 1, F1: 21, F2: (((((p0 + ("" + p0)) + p0) + p0) + "hello") + p0)}
                    _ = fi10
                    _ = v11
                }
                _ = v9
            }
            _ = fi5
        }
    }
    v12 := []int{-9223372036854775808, -419, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 263
}
func genFn1(p0 int32) int16 {
    v13 := GenS1{F0: -125, F1: -109, F2: ""}
    for fi14 := 0; fi14 < 7; fi14++ {
        p0 -= p0
        _ = fi14
    }
    if (1000 != 32767) {
        var v15 float32 = ((1e10 * 0.0) * 0.0)
        var v16 bool = true
        _ = v15
        _ = v16
    } else {
        if false {
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := []int{(fi17 - 32767)}
                var v19 uint8 = 11
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS1{F0: -9223372036854775808, F1: -27, F2: ((("foo bar" + "") + "\n\t") + "")}
            _ = v20
        } else {
            p0 = (p0 / (((p0 / ((p0) | 1))) | 1))
        }
    }
    v21 := GenS2{F0: -132, F1: 186, F2: 127}
    v22 := GenS3{F0: 42}
    _ = p0
    _ = v13
    _ = v21
    _ = v22
    return -495
}

func main() {
    var v1 int16 = 148
    v2 := []string{"\n\t", ("\n\t" + "hello")}
    var v3 uint16 = 922
    var v4 int64 = 790
    var v23 string = "\n\t"
    var v24 uint16 = v3
    v2 = append(v2, v23)
    var v25 bool = true
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS2{F0: -1, F1: 2, F2: 127}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int64 = -1
        _ = fi28
        _ = v29
    }
    fmt.Println(v25, v1, v25)
    for fi30 := 0; fi30 < 6; fi30++ {
        v24 -= 42
        _ = fi30
    }
    v1 = (v1 / ((v1) | 1))
    if (((((v23 + (v23 + v23)) + ("\n\t" + "foo bar")) + (v23 + v23)) + ((("hello" + "foo bar") + v23) + v23)) != v23) {
        v4 += (^(v4))
        for fi31 := 0; fi31 < 1; fi31++ {
            if (!((fi31 << ((1) & 63)) >= fi31)) {
                switch (v23 + (v23 + "")) {
                case "\n\t":
                    var v32 float64 = 0.1
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v34 := GenS1{F0: (^(v4)), F1: 127, F2: v23}
                        if (!(!v25)) {
                            v1 -= 1000
                            switch fi33 {
                            case 0:
                                v35 := GenS4{F0: genFn1(-712)}
                                _ = v35
                            }
                            v4 = 0
                        }
                        var v36 float32 = 1.5
                        _ = fi33
                        _ = v34
                        _ = v36
                    }
                }
                for fi37 := 0; fi37 < 10; fi37++ {
                    var v38 int64 = (v4 / (((v4 + (v4 / ((v4) | 1)))) | 1))
                    _ = fi37
                    _ = v38
                }
                for fi39 := 0; fi39 < 12; fi39++ {
                    v25 = (((v24 ^ (v24 ^ 42)) < v3) || ((v3 << ((18446744073709551615) & 63)) >= (^(v24))))
                    v24 = v3
                    _ = fi39
                }
            }
            var v40 int16 = v1
            _ = fi31
            _ = v40
        }
        var v41 int8 = 114
        _ = v41
    }
    v1 += v1
    for fi42 := 0; fi42 < 3; fi42++ {
        _ = len(v2)
        fmt.Println(v24, fi42)
        fmt.Println(fi42, v23)
        _ = fi42
    }
    v43 := GenS3{F0: 255}
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "a")
    }
    v25 = (1 <= 458)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v43
}
