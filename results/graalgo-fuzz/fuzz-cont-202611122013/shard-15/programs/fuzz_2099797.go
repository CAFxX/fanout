package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 bool }
type GenS4 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int = 32767
    var v2 bool = false
    var v3 uint16 = 812
    var v4 float32 = ((1.5 / 3.14159) - 3.14159)
    fmt.Println(v3, v3, v2)
    if v2 {
        if true {
            if (-162 == 42) {
                for fi5 := 0; fi5 < 17; fi5++ {
                    for fi6 := 0; fi6 < 2; fi6++ {
                        switch 127 {
                        case 0:
                            var v7 string = ("foo bar" + "")
                            _ = v7
                        default:
                            v8 := GenS1{F0: v2}
                            _ = v8
                        }
                        var v9 uint16 = (v3 - (v3 << ((18446744073709551615) & 63)))
                        _ = fi6
                        _ = v9
                    }
                    _ = fi5
                }
                v10 := []int{v1, -804, (v1 / (((^(v1))) | 1))}
                _ = v10
            } else {
                v4 -= 3.14159
                fmt.Println(v4, v4, v2)
            }
            for fi11 := 0; fi11 < 17; fi11++ {
                if (0 < -817) {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v13 := GenS2{F0: 1}
                        _ = fi12
                        _ = v13
                    }
                }
                _ = fi11
            }
        } else {
            if ((!v2) && (v2 || (!(2147483647 > 2147483648)))) {
                v1 -= v1
                v4 = ((1e10 * v4) + 0.1)
            } else {
                fmt.Println(v3, v3)
                v14 := []int{(v1 / ((v1) | 1)), v1}
                _ = v14
            }
            var v15 float32 = v4
            _ = v15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := []int{}
            v2 = (((v3 ^ (v3 / (((v3 + v3)) | 1))) >= 42) && v2)
            if len(v17) > 0 {
                v17[0] = -604
            }
            _ = fi16
            _ = v17
        }
        v1 = v1
    }
    v18 := GenS3{F0: 0, F1: (^(v3)), F2: true}
    var v19 bool = v2
    var v20 uint64 = 42
    switch v1 {
    case 2147483647:
        switch -2147483648 {
        case -374:
            var v21 int64 = 42
            _ = v21
        case 850:
            v22 := []int{128}
            _ = v22
        default:
            fmt.Println(v2)
        }
    case -276:
        fmt.Println(v3, v1)
    case 983:
        v3 = v3
    }
    v23 := GenS1{F0: v2}
    fmt.Println(v20)
    var v24 int16 = 0
    var v25 string = ("" + "foo bar")
    v1 = int(v24)
    if true {
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := []int{-102, (^(v1))}
            _ = fi26
            _ = v27
        }
        v28 := []int{0, int(v3), 9223372036854775807}
        if v2 {
            v29 := GenS1{F0: (-1 != 42)}
            _ = len(v28)
            _ = v29
        } else {
            var v30 float64 = 1e10
            fmt.Println(v19, v4, v4)
            _ = v30
        }
        _ = v28
    }
    v31 := GenS1{F0: ((v24 | 1000) > v24)}
    var v32 int16 = v24
    v33 := GenS3{F0: (v20 * v20), F1: 219, F2: v2}
    v34 := []int{(-(v1))}
    v35 := GenS4{F0: 34, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
