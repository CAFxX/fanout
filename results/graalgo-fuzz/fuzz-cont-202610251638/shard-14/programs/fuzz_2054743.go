package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 uint; F1 string; F2 int16 }

func main() {
    var v1 string = "hello"
    var v2 uint = 916
    var v3 int = -315
    var v4 int64 = -1
    var v5 uint16 = 315
    var v6 uint64 = 42
    var v7 uint32 = 569
    if true {
        v8 := []int{127, (v3 >> (((v2 + uint(v7))) & 63))}
        _ = v8
    }
    if ((false || (v3 > v3)) || false) {
        var v9 int16 = 32767
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{-9223372036854775808, (^(fi10)), (v3 * 9223372036854775807)}
        for fi12 := 0; fi12 < 12; fi12++ {
            if true {
                if (true || false) {
                    switch int(v7) {
                    case -313:
                        var v13 int64 = (v4 ^ (v4 & 9223372036854775807))
                        _ = v13
                    }
                    switch (fi10 % ((fi12) | 1)) {
                    case -2147483648:
                        v14 := GenS1{F0: (^(v2))}
                        _ = v14
                    }
                }
                v15 := GenS1{F0: (^(v2))}
                fmt.Println(v1, v2)
                _ = v15
            } else {
                var v16 int = 32768
                var v17 uint8 = 192
                if ("" == (v1 + (v1 + v1))) {
                    v18 := GenS1{F0: 255}
                    var v19 uint16 = 65535
                    if ((!(((0.0 + float32(((1.5 + (float64(-2.25) * float64(1.5))) + -2.25))) >= 1e10) && false)) && false) {
                        switch 128 {
                        case -162:
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        case 124:
                            var v20 float64 = (0.0 + 3.14159)
                            _ = v20
                        default:
                            v21 := GenS1{F0: uint(v19)}
                            _ = v21
                        }
                        var v22 uint8 = (v17 | 0)
                        _ = v22
                    } else {
                        switch v16 {
                        case -9223372036854775808:
                            v23 := GenS1{F0: (^(v2))}
                            _ = v23
                        case 32767:
                            v5 -= (^(v19))
                        default:
                            v4 -= (v4 % ((v4) | 1))
                        }
                        v17 = (v17 | (^(v17)))
                    }
                    _ = v18
                    _ = v19
                } else {
                    switch (v1 + ((v1 + ("a" + "hello")) + v1)) {
                    case "":
                        v16 += v3
                    default:
                        var v24 uint32 = v7
                        _ = v24
                    }
                    fmt.Println(v7, fi10)
                    var v25 int16 = 1000
                    v11 = append(v11, 0)
                    _ = v25
                }
                _ = v16
                _ = v17
            }
            if ((^(v2)) < (v2 >> ((v2) & 63))) {
                if ((v4 <= int64(v2)) && (841 == v5)) {
                    v1 = ("" + (v1 + "hello"))
                    var v26 uint16 = (^(v5))
                    _ = v26
                } else {
                    v5 += (v5 / ((v5) | 1))
                    v27 := GenS2{F0: (v5 / ((352) | 1)), F1: (!false), F2: (!true)}
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v27
                }
            }
            v28 := GenS2{F0: (v5 / ((v5) | 1)), F1: (!(2 != uint64(v2))), F2: ((v5 << ((18446744073709551615) & 63)) == v5)}
            _ = fi12
            _ = v28
        }
        _ = fi10
        _ = v11
    }
    v29 := GenS1{F0: (v2 + v2)}
    for fi30 := 0; fi30 < 2; fi30++ {
        v31 := []int{2147483647, (fi30 & 1)}
        var v32 int8 = -95
        switch ("foo bar" + (v1 + ((v1 + "hello") + (v1 + v1)))) {
        case "":
            var v33 float64 = 3.14159
            _ = v33
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    var v34 uint = v2
    v35 := GenS2{F0: (v5 << ((2) & 63)), F1: true, F2: true}
    v36 := GenS3{F0: (^(v34)), F1: "foo bar", F2: 1000}
    v37 := []int{int(v5)}
    v38 := GenS3{F0: 896, F1: ("hello" + ""), F2: -132}
    var v39 int64 = int64(v5)
    var v40 float32 = ((((3.14159 / 1.5) / (1e10 + float32(1.5))) - ((1e10 + 0.1) + (3.14159 / 1e10))) - -2.25)
    var v41 float64 = 1e10
    switch v3 {
    case -658:
        v2 = 109
    }
    v40 -= ((v40 * 1.5) * v40)
    _ = len(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
