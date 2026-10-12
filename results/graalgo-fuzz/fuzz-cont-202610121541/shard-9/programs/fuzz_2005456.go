package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 bool = (2147483648 > 764)
    var v2 uint8 = 0
    var v3 bool = v1
    var v4 uint16 = 985
    var v5 int32 = -302
    v5 -= (v5 / (((v5 + 0)) | 1))
    v6 := GenS1{F0: 42, F1: -2147483648, F2: v2}
    fmt.Println(v1)
    v7 := GenS1{F0: 42, F1: 0, F2: v2}
    v4 = v4
    switch ("hello" + "hello") {
    case "a":
        var v8 uint32 = 4294967295
        _ = v8
    case "\n\t":
        v9 := GenS1{F0: 1, F1: -906, F2: 20}
        _ = v9
    case "hello":
        var v10 int = -157
        _ = v10
    }
    switch "" {
    case "\n\t":
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 12; fi12++ {
                if ((v4 >= v4) || v3) {
                    var v13 float64 = (0.0 + (1.5 - float64(float32(0.1))))
                    _ = v13
                } else {
                    if (!(0 < 967)) {
                        var v14 int32 = v5
                        if v3 {
                            v3 = ((v14 * (v5 % (((v5 / ((-802) | 1))) | 1))) <= v5)
                            v3 = false
                        }
                        _ = v14
                    }
                }
                v1 = (1 < -1)
                _ = fi12
            }
            var v15 float64 = float64(float32(3.14159))
            _ = fi11
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 15; fi16++ {
            if ((v4 >> ((255) & 63)) >= v4) {
                v5 = (v5 << ((1) & 63))
                switch fi16 {
                case 562:
                    for fi17 := 0; fi17 < 20; fi17++ {
                        var v18 float32 = float32(1.5)
                        _ = fi17
                        _ = v18
                    }
                }
            } else {
                var v19 uint32 = 42
                var v20 int64 = 0
                fmt.Println(v4)
                _ = v19
                _ = v20
            }
            var v21 string = ""
            _ = fi16
            _ = v21
        }
    }
    v22 := []int{}
    v23 := GenS1{F0: 42, F1: 252, F2: v2}
    v22 = append(v22, 820)
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            v22 = append(v22, (fi25 ^ (fi24 & (fi24 / (((-(fi24))) | 1)))))
            if true {
                if (!v3) {
                    v1 = (!false)
                }
                var v26 int = (fi25 % (((fi25 & (-(fi25)))) | 1))
                var v27 bool = false
                v4 = (v4 << ((42) & 63))
                v28 := GenS2{F0: 18446744073709551615}
                _ = v26
                _ = v27
                _ = v28
            } else {
                if v1 {
                    v1 = (((0.1 - 3.14159) >= -2.25) && (2147483647 == fi25))
                } else {
                    v29 := GenS3{F0: -203}
                    fmt.Println(v3, fi24, v3)
                    v2 += v2
                    v4 = 971
                    _ = v29
                }
                if len(v22) > 0 {
                    v22[0] = (-(fi24))
                }
                var v30 int16 = 42
                _ = v30
            }
            v31 := GenS4{F0: 9223372036854775807, F1: 64}
            v32 := GenS2{F0: 0}
            _ = fi25
            _ = v31
            _ = v32
        }
        _ = fi24
    }
    v33 := GenS5{F0: (v2 % ((255) | 1))}
    var v34 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
