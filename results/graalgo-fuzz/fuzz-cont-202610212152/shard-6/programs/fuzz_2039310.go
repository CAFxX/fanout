package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0() bool {
    v4 := []int{2147483647}
    v5 := GenS1{F0: 9223372036854775807}
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 bool) int {
    p0 = true
    var v6 int32 = 2
    _ = p0
    _ = v6
    return 88
}

func main() {
    v1 := []int{}
    var v2 uint64 = 331
    var v3 int64 = 9223372036854775807
    v2 = (v2 * (v2 & (v2 >> ((406) & 63))))
    v3 = (v3 ^ (v3 | -9223372036854775808))
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 = (v2 - (v2 / (((v2 / ((v2) | 1))) | 1)))
        var v8 int = (fi7 | genFn1((23 >= 1)))
        for fi9 := 0; fi9 < 19; fi9++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int32 = 827
        if ((668 <= 938) && (127 == 255)) {
            for fi12 := 0; fi12 < 8; fi12++ {
                if len(v1) > 0 {
                    v1[0] = 32767
                }
                _ = fi12
            }
            switch "" {
            case "hello":
                v1 = append(v1, fi10)
            case "\n\t":
                fmt.Println(fi10, fi10)
            default:
                v3 += -1
            }
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(fi13)
                _ = fi13
            }
        } else {
            if false {
                v14 := GenS2{F0: v11, F1: int64(v11)}
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS2{F0: v11, F1: (^(v3))}
                    var v17 string = (("\n\t" + "") + ((("a" + "\n\t") + "a") + ""))
                    if false {
                        var v18 int64 = (-(v3))
                        if true {
                            v3 = (v3 << ((42) & 63))
                        }
                        if ((false || true) && false) {
                            v17 = (v17 + v17)
                            v17 = "hello"
                            var v19 string = ""
                            _ = v19
                        }
                        _ = v18
                    } else {
                        var v20 int = 1
                        _ = v20
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v14
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    fmt.Println(v11)
                    v3 = (v3 % (((v3 | 770)) | 1))
                    v22 := GenS1{F0: (v3 & v3)}
                    _ = fi21
                    _ = v22
                }
                v11 -= v11
                fmt.Println(v3, v2, v2)
            }
            v11 += (-(v11))
            if len(v1) > 0 {
                v1[0] = (-(fi10))
            }
        }
        _ = fi10
        _ = v11
    }
    if (!false) {
        v1 = append(v1, -934)
        v23 := GenS2{F0: 0, F1: v3}
        for fi24 := 0; fi24 < 2; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int16 = 42
                fmt.Println(v26, fi25, fi24)
                var v27 uint8 = 2
                _ = fi25
                _ = v26
                _ = v27
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi24
        }
        _ = v23
    }
    v2 += 102
    v28 := GenS2{F0: -2147483648, F1: (v3 >> ((1) & 63))}
    fmt.Println(v3, v2)
    for fi29 := 0; fi29 < 13; fi29++ {
        v1 = append(v1, 491)
        for fi30 := 0; fi30 < 1; fi30++ {
            _ = len(v1)
            _ = fi30
        }
        v2 = (v2 | v2)
        _ = fi29
    }
    var v31 int32 = -1
    var v32 bool = true
    _ = len(v1)
    switch -657 {
    case 2147483647:
        v1 = append(v1, genFn1((((^(v2)) >= 87) || ((!v32) || v32))))
    default:
        switch -588 {
        case 2147483647:
            var v33 string = ("" + "hello")
            _ = v33
        default:
            var v34 float64 = 1e10
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 int64 = v3
        v32 = v32
        fmt.Println(fi35)
        _ = fi35
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
}
