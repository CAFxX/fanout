package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 string; F1 uint8 }

func main() {
    var v1 bool = true
    var v2 uint64 = 4294967295
    var v3 int = -2147483648
    var v4 int16 = 32767
    if v1 {
        v5 := []int{}
        _ = v5
    }
    if (v1 && (-2147483648 > -2147483648)) {
        v1 = (true && v1)
        fmt.Println(v2, v4, v2)
    } else {
        switch (v3 + 4294967295) {
        case 32768:
            if v1 {
                v6 := GenS1{F0: 0, F1: -1, F2: v2}
                _ = v6
            } else {
                var v7 int64 = -1
                v4 -= 1
                if v1 {
                    v7 -= v7
                    v8 := GenS2{F0: 61}
                    _ = v8
                }
                _ = v7
            }
        case -1:
            v3 -= 32768
        case 581:
            var v9 uint32 = 196
            _ = v9
        }
    }
    var v10 uint16 = 0
    for fi11 := 0; fi11 < 2; fi11++ {
        switch 127 {
        case 972:
            switch (v3 - (^(v3))) {
            case -161:
                if ("" != "") {
                    var v12 string = ("hello" + ("" + ""))
                    var v13 uint16 = v10
                    _ = v12
                    _ = v13
                } else {
                    var v14 float64 = (3.14159 / 3.14159)
                    if ((v4 - v4) == v4) {
                        v10 = v10
                        v14 += v14
                    }
                    v15 := []int{(fi11 ^ (v3 / ((v3) | 1))), int(v10), (v3 * 2147483647), v3}
                    _ = v14
                    _ = v15
                }
            case 1:
                switch fi11 {
                case 32768:
                    if true {
                        var v16 int16 = v4
                        _ = v16
                    }
                case 32767:
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v18 := []int{}
                        switch 127 {
                        case -1:
                            v3 += 0
                        case -2147483648:
                            var v19 uint32 = 30
                            _ = v19
                        }
                        var v20 uint32 = 0
                        if (v10 < (v10 % ((v10) | 1))) {
                            switch (fi11 | (v3 ^ (^(fi17)))) {
                            case 128:
                                var v21 uint16 = uint16(v2)
                                _ = v21
                            case 0:
                                v4 += v4
                            }
                        } else {
                            v20 -= (^(v20))
                            fmt.Println(v1)
                            if v1 {
                                var v22 uint16 = 32767
                                v10 -= v22
                                var v23 float64 = 1e10
                                _ = v22
                                _ = v23
                            }
                        }
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        _ = fi17
                        _ = v18
                        _ = v20
                    }
                default:
                    v24 := GenS3{F0: v4, F1: "a"}
                    _ = v24
                }
            case 32767:
                v25 := GenS4{F0: ("hello" + "foo bar"), F1: 2}
                _ = v25
            default:
                v26 := GenS1{F0: 42, F1: 0, F2: (v2 ^ 42)}
                _ = v26
            }
        case 42:
            v1 = (0 == (v10 / ((65535) | 1)))
        }
        _ = fi11
    }
    v27 := []int{32767, v3, (v3 * (v3 + (^(v3))))}
    if len(v27) > 0 {
        v27[0] = (v3 / (((-(v3))) | 1))
    }
    _ = len(v27)
    _ = len(v27)
    var v28 uint8 = 42
    v29 := GenS4{F0: (("" + (("a" + "") + "")) + "foo bar"), F1: 1}
    fmt.Println(v2, v10, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
