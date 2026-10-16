package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }
func genFn0(p0 int64) uint8 {
    var v6 int = 42
    p0 = p0
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int16 = -1
        var v9 float64 = float64(0.0)
        fmt.Println(p0, v6)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return 201
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = false
    v3 := []int{1, 128}
    v4 := []uint32{}
    var v5 uint16 = 1
    v10 := GenS1{F0: 2147483647, F1: -1, F2: 18446744073709551615}
    if false {
        v11 := GenS1{F0: -2147483648, F1: 1, F2: 1}
        _ = v11
    }
    if (1 != -278) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if v2 {
            var v12 uint8 = 105
            _ = v12
        } else {
            v5 += 51
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                v5 -= v5
                _ = fi13
            }
        }
    } else {
        if (!((-2147483648 <= 0) && (v1 <= (v1 % (((v1 >> ((65535) & 63))) | 1))))) {
            var v14 uint16 = v5
            _ = len(v4)
            _ = v14
        }
        fmt.Println(v2, v5, v1)
    }
    if ((v1 != v1) && true) {
        switch 32767 {
        case 127:
            for fi15 := 0; fi15 < 5; fi15++ {
                v5 = v5
                var v16 uint32 = 145
                if (((v1 << ((30) & 63)) < (v1 ^ v1)) && (!(4294967295 >= 18446744073709551615))) {
                    v17 := GenS2{F0: v2, F1: ((v5 / (((v5 & (v5 & 65535))) | 1)) >= 2), F2: 663}
                    if true {
                        if v2 {
                            var v18 uint32 = 646
                            _ = v18
                        }
                        var v19 uint32 = 737
                        _ = v19
                    }
                    var v20 uint = 0
                    _ = v17
                    _ = v20
                } else {
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if ((float32((0.1 / 0.1)) * 0.0) <= float32(0.1)) {
                            if (((-2.25 - 0.1) / 1.0) <= float64(-2.25)) {
                                if v2 {
                                    var v21 string = "\n\t"
                                    fmt.Println(v21, v21)
                                    _ = v21
                                }
                            } else {
                                v1 += (v1 - v1)
                            }
                            _ = len(v3)
                        }
                    } else {
                        if (true || v2) {
                            var v22 float64 = (float64(float32((float64(float32(1e10)) / 1.0))) / 3.14159)
                            _ = v22
                        }
                    }
                    v23 := GenS2{F0: (-315 <= (^(fi15))), F1: (18446744073709551615 != 255), F2: 2147483647}
                    _ = v23
                }
                _ = fi15
                _ = v16
            }
        }
        for fi24 := 0; fi24 < 20; fi24++ {
            v3 = append(v3, (fi24 & fi24))
            v5 += (v5 & v5)
            switch "a" {
            case "hello":
                var v25 bool = (("\n\t" + "foo bar") <= ("\n\t" + ("" + "a")))
                _ = v25
            }
            _ = fi24
        }
        switch 42 {
        case 128:
            var v26 bool = (181 >= genFn0((^(v1))))
            _ = v26
        }
    } else {
        var v27 uint64 = 991
        _ = v27
    }
    _ = len(v4)
    _ = len(v4)
    var v28 int32 = 2147483647
    fmt.Println(v2, v2, v2)
    _ = len(v3)
    v28 -= v28
    v1 = 929
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
