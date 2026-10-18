package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8; F1 int; F2 int8 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0(p0 int8) int {
    fmt.Println(p0)
    switch -9223372036854775808 {
    case 128:
        if (-32768 >= 32767) {
            fmt.Println(p0)
            p0 += p0
        }
    case 127:
        v6 := GenS1{F0: p0, F1: -50}
        _ = v6
    case -1:
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 0, F1: 1}
    _ = p0
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{0, 1}
    var v2 string = (("hello" + "\n\t") + "foo bar")
    var v3 uint16 = 746
    var v4 int8 = -18
    v5 := []int{}
    var v10 int8 = 100
    if true {
        var v11 string = v2
        for fi12 := 0; fi12 < 18; fi12++ {
            switch (v11 + ("foo bar" + (v2 + ("a" + (v11 + "foo bar"))))) {
            case "foo bar":
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := GenS2{F0: (v4 << ((537) & 63)), F1: genFn0(v10), F2: -128}
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                v2 = "foo bar"
            default:
                if false {
                    v15 := GenS2{F0: (v10 ^ -43), F1: genFn0((v10 - (v4 - (v10 % ((-1) | 1))))), F2: v10}
                    v11 = v2
                    if true {
                        switch (fi12 - genFn0(v10)) {
                        case 128:
                            v16 := GenS3{F0: (!false), F1: (true && true), F2: 92}
                            _ = v16
                        case 4294967295:
                            if false {
                                if false {
                                    v3 = 255
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                } else {
                                    switch v11 {
                                    case "foo bar":
                                        v11 = (("\n\t" + v2) + (v11 + "a"))
                                    }
                                    v17 := GenS3{F0: true, F1: ((-2.25 * -2.25) == (((float64(0.1) + 0.1) + 1.5) * (0.0 * (-2.25 + -2.25)))), F2: 134}
                                    fmt.Println(v11, v3)
                                    _ = v17
                                }
                                var v18 int64 = 323
                                _ = v18
                            }
                        default:
                            var v19 uint64 = 42
                            _ = v19
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 934
                        }
                        var v20 float32 = 0.1
                        if (!false) {
                            for fi21 := 0; fi21 < 8; fi21++ {
                                var v22 int32 = -1
                                _ = fi21
                                _ = v22
                            }
                        }
                        _ = v20
                    }
                    _ = v15
                }
            }
            v3 -= v3
            _ = fi12
        }
        _ = v11
    }
    fmt.Println(v4)
    var v23 string = v2
    v24 := GenS1{F0: -128, F1: 1}
    v4 = v4
    var v25 int64 = 42
    var v26 uint16 = 902
    v3 -= (v3 * (v3 % ((v26) | 1)))
    v4 = (v4 & v4)
    var v27 int = 9223372036854775807
    v28 := GenS4{F0: 42, F1: 128}
    fmt.Println(v25, v2, v2)
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := GenS3{F0: ((!(!false)) && (-452 > (^(v25)))), F1: (!true), F2: 127}
        v10 += v4
        for fi31 := 0; fi31 < 18; fi31++ {
            v26 = 65535
            v25 -= v25
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    v1 = append(v1, 953)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
