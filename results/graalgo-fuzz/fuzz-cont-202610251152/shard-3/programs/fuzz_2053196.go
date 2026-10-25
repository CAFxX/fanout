package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := []int{-1, 32768, 4294967295, 32768}
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    var v4 int32 = 1
    for fi5 := 0; fi5 < 20; fi5++ {
        v4 = (v4 / ((0) | 1))
        var v6 float64 = (((((((-2.25 - float64(0.0)) - 3.14159) / float64((0.0 - 3.14159))) + ((float64(1e10) - 1.5) - 0.1)) - 0.1) * (-2.25 - (float64(float32(3.14159)) + (0.1 / 1.5)))) + float64((1.5 - (0.0 + 0.1))))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v4 = (v4 % (((v4 % (((v4 & v4)) | 1))) | 1))
        delete(v2, 2147483647)
        v4 += (^(v4))
        _ = fi7
    }
    fmt.Println(v4)
    var v8 int64 = -878
    if true {
        var v9 bool = false
        v2[0] = 660
        _ = v9
    } else {
        if ((true || (-57 >= 0)) || true) {
            _ = len(v2)
        } else {
            v10 := GenS1{F0: (v8 | v8), F1: 1}
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint32 = 211
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        v4 = 42
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 -= 2147483647
        _ = fi13
    }
    var v14 uint64 = 413
    if (2 <= 4294967295) {
        v15 := v2[65535]
        for fi16 := 0; fi16 < 3; fi16++ {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            var v17 uint8 = 2
            switch ("" + "hello") {
            case "\n\t":
                for fi18 := 0; fi18 < 2; fi18++ {
                    v19 := GenS1{F0: 42, F1: fi16}
                    if false {
                        v20 := GenS1{F0: -9223372036854775808, F1: fi18}
                        if (!((true && false) && (v8 > (-(v8))))) {
                            if len(v1) > 0 {
                                v1[0] = (^(fi18))
                            }
                        }
                        v21 := v2[v15]
                        _ = v20
                        _ = v21
                    } else {
                        v4 -= v4
                    }
                    v22 := GenS1{F0: (v8 | v8), F1: 9223372036854775807}
                    v23 := GenS1{F0: (v8 % ((v8) | 1)), F1: (fi16 << ((474) & 63))}
                    _ = fi18
                    _ = v19
                    _ = v22
                    _ = v23
                }
            case "":
                if (true && (-1 == 127)) {
                    for fi24 := 0; fi24 < 9; fi24++ {
                        var v25 int = 127
                        _ = fi24
                        _ = v25
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v26 uint = 344
                    if false {
                        var v27 uint8 = (v17 - (v17 + v17))
                        var v28 uint64 = (v14 / (((v14 ^ v14)) | 1))
                        v29 := GenS1{F0: (^(v8)), F1: 663}
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        v14 -= (v14 % (((v14 ^ v14)) | 1))
                        var v30 int64 = (^(v8))
                        v17 -= (^(v17))
                        _ = v30
                    }
                    _ = v26
                }
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    v8 = (v8 / (((^(v8))) | 1))
    var v31 int64 = 0
    v3 = append(v3, "\n\t")
    fmt.Println(v31, v14, v14)
    switch 0 {
    case -646:
        if false {
            delete(v2, 174)
            for fi32 := 0; fi32 < 19; fi32++ {
                v4 += v4
                v14 = (v14 * 0)
                _ = fi32
            }
            var v33 int16 = 42
            _ = v33
        } else {
            v14 -= 576
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 uint32 = 2147483647
                v8 -= -432
                _ = fi34
                _ = v35
            }
        }
    default:
        var v36 uint16 = 227
        _ = v36
    }
    v8 += (^(v8))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v8
}
