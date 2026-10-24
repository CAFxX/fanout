package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS4 struct { F0 int; F1 uint64; F2 uint16 }

func main() {
    var v1 int8 = -1
    var v2 float64 = float64(-2.25)
    var v3 int8 = v1
    var v4 bool = (!false)
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int8 = (v1 + v1)
        if (647 < 967) {
            switch "hello" {
            case "foo bar":
                v7 := []int{(fi5 ^ (fi5 % ((-1) | 1))), 211, (^(fi5)), fi5}
                _ = v7
            case "a":
                v2 += v2
            case "":
                if v4 {
                    fmt.Println(v6)
                } else {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v6 = v3
                        _ = fi8
                    }
                }
            default:
                switch (fi5 - fi5) {
                case -129:
                    v9 := GenS1{F0: 18446744073709551615, F1: -128, F2: (false || false)}
                    _ = v9
                case 804:
                    if (!v4) {
                        v2 += v2
                        v6 = (v6 | (v3 % ((42) | 1)))
                    } else {
                        v10 := GenS1{F0: 0, F1: 127, F2: true}
                        if false {
                            v1 = (-(v6))
                        }
                        _ = v10
                    }
                case -717:
                    for fi11 := 0; fi11 < 18; fi11++ {
                        v12 := GenS1{F0: 637, F1: (v3 ^ 0), F2: v4}
                        _ = fi11
                        _ = v12
                    }
                default:
                    switch fi5 {
                    case -1:
                        if v4 {
                            var v13 float32 = -2.25
                            v6 -= -128
                            var v14 uint8 = 2
                            _ = v13
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 11; fi15++ {
                                var v16 int32 = 0
                                v6 = (v1 ^ v1)
                                _ = fi15
                                _ = v16
                            }
                            v17 := GenS2{F0: 0}
                            var v18 float64 = 0.0
                            _ = v17
                            _ = v18
                        }
                    }
                }
            }
        } else {
            var v19 float64 = v2
            _ = v19
        }
        var v20 float32 = 3.14159
        _ = fi5
        _ = v20
        _ = v6
    }
    v4 = false
    fmt.Println(v1, v1)
    v21 := GenS3{F0: 32767, F1: -1, F2: 0}
    if (1 >= 4294967295) {
        for fi22 := 0; fi22 < 16; fi22++ {
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 int64 = -1
                v25 := []int{fi22, (^(fi22)), 0}
                v26 := GenS2{F0: (v24 - -101)}
                v4 = ((v24 >= (v24 ^ v24)) || v4)
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi22
        }
    } else {
        v27 := GenS1{F0: 202, F1: 109, F2: v4}
        v28 := []int{}
        var v29 uint16 = 1
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS4{F0: 1, F1: 0, F2: 127}
    switch -444 {
    case 0:
        v31 := GenS3{F0: 2, F1: 2147483647, F2: 255}
        _ = v31
    default:
        for fi32 := 0; fi32 < 9; fi32++ {
            v33 := GenS1{F0: 1, F1: 0, F2: v4}
            _ = fi32
            _ = v33
        }
    }
    fmt.Println(v1, v2)
    v3 = v1
    var v34 int16 = 0
    v35 := GenS1{F0: 18446744073709551615, F1: (v1 + 100), F2: (!(!v4))}
    fmt.Println(v2, v1, v4)
    v36 := GenS1{F0: 4294967295, F1: (v1 ^ (v1 | (v1 | (v3 * (v1 - v3))))), F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
