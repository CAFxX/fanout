package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 bool; F2 int }
func genFn0(p0 uint64) int {
    var v6 uint64 = p0
    var v7 uint64 = (v6 % ((p0) | 1))
    v8 := []int{0, 9223372036854775807, -1}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -257
}

func main() {
    v1 := []string{"a"}
    var v2 int16 = -32768
    var v3 uint32 = 65535
    var v4 float64 = 0.0
    var v5 string = "foo bar"
    fmt.Println(v2, v4)
    if ((((((3.14159 + 0.1) * (((1e10 + 0.1) + 1e10) - 0.1)) - (1.5 / 3.14159)) - (1.5 + (-2.25 + 3.14159))) * (((0.1 / 0.1) / -2.25) + float32((3.14159 / v4)))) >= 1.5) {
        v2 -= -559
    } else {
        if true {
            var v9 uint8 = 0
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(v9)
                var v11 uint16 = 255
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(v2, v3, v5)
                if false {
                    v5 = v5
                    v2 = (v2 / ((1) | 1))
                }
                v4 -= v4
                _ = fi12
            }
            if (!(v3 >= (v3 ^ (^(v3))))) {
                v3 = 4294967295
                if true {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        var v14 string = ("\n\t" + "")
                        v3 = v3
                        v15 := GenS1{F0: -152}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 1; fi16++ {
                        _ = len(v1)
                        _ = fi16
                    }
                    if ((1e10 - (((((((1.5 * 1.5) / 1e10) * 0.0) * ((0.0 + 1.5) + -2.25)) * -2.25) * 1e10) / (3.14159 + ((-2.25 * -2.25) * (1.5 + (0.1 + 1.5)))))) < (1e10 * 1e10)) {
                        v4 = (((float64(-2.25) + v4) - ((v4 - 0.0) * v4)) + (v4 - float64(3.14159)))
                    } else {
                        v3 = (v3 * 2)
                        fmt.Println(v3, v5)
                    }
                }
            }
        }
    }
    v4 -= ((v4 - v4) + v4)
    v4 -= (v4 + v4)
    var v17 int8 = 60
    v4 -= v4
    if false {
        if (-2.25 != ((((1.5 * 3.14159) / 1e10) - 3.14159) - -2.25)) {
            for fi18 := 0; fi18 < 8; fi18++ {
                switch fi18 {
                case -444:
                    var v19 uint32 = (v3 & (v3 ^ v3))
                    _ = v19
                case 32767:
                    v20 := GenS2{F0: 220}
                    _ = v20
                case -2147483648:
                    if (false || ((fi18 % ((-986) | 1)) < (^(fi18)))) {
                        v4 = (3.14159 / -2.25)
                        v5 = (v5 + v5)
                    }
                }
                if (!true) {
                    switch 127 {
                    case 32768:
                        var v21 int = 689
                        _ = v21
                    case 0:
                        if false {
                            _ = len(v1)
                            switch (fi18 + (^(fi18))) {
                            case 23:
                                fmt.Println(fi18)
                            case 32768:
                                v3 -= (v3 | (^(v3)))
                            default:
                                v22 := GenS2{F0: 0}
                                _ = v22
                            }
                            var v23 int16 = 1000
                            _ = v23
                        }
                    }
                    v2 = v2
                    fmt.Println(fi18, v4)
                }
                var v24 uint64 = 686
                _ = fi18
                _ = v24
            }
            v3 = 4294967295
        }
    }
    fmt.Println(v5, v3, v4)
    if ((2147483647 == -1) && true) {
        v25 := GenS3{F0: true, F1: true, F2: -419}
        fmt.Println(v2, v17)
        switch 9223372036854775807 {
        case -781:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v25
    }
    var v26 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
