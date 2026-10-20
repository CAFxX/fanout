package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
func genFn0() bool {
    for fi6 := 0; fi6 < 12; fi6++ {
        if false {
            var v7 uint64 = 0
            switch (fi6 ^ fi6) {
            case -689:
                for fi8 := 0; fi8 < 9; fi8++ {
                    switch fi8 {
                    case 128:
                        v7 = v7
                    default:
                        if ((1.5 == 1e10) || true) {
                            v9 := []int{}
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            if true {
                                v7 -= v7
                            } else {
                                var v10 uint64 = 801
                                var v11 uint64 = v7
                                var v12 uint16 = 255
                                _ = v10
                                _ = v11
                                _ = v12
                            }
                            _ = v9
                        }
                    }
                    _ = fi8
                }
            default:
                for fi13 := 0; fi13 < 11; fi13++ {
                    v7 = v7
                    _ = fi13
                }
            }
            v14 := []int{fi6, fi6, 127}
            _ = v14
            _ = v7
        }
        var v15 uint16 = 1
        v16 := []int{-1, fi6, fi6}
        _ = fi6
        _ = v15
        _ = v16
    }
    var v17 uint8 = 46
    _ = v17
    return true
}
func genFn1(p0 int, p1 int64, p2 int64) bool {
    var v18 string = ("\n\t" + "\n\t")
    p0 += 127
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (!(0 >= 127))
}

func main() {
    var v1 int64 = 50
    var v2 float32 = 3.14159
    var v3 float64 = 1e10
    var v4 int16 = 32767
    var v5 string = "foo bar"
    v19 := []int{499, 127, 32767, -717}
    v19 = append(v19, 203)
    var v20 uint32 = 0
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    _ = len(v19)
    v2 = v2
    v5 = v5
    v21 := GenS1{F0: 396, F1: 50, F2: 248}
    v5 = "foo bar"
    v20 += (^(v20))
    for fi22 := 0; fi22 < 15; fi22++ {
        if false {
            if (-100 >= -1) {
                v23 := GenS1{F0: -1, F1: (^(v1)), F2: 1}
                fmt.Println(v4, v4, v20)
                _ = v23
            }
            v1 += 9223372036854775807
        } else {
            switch (-(fi22)) {
            case 886:
                var v24 int16 = 1000
                _ = v24
            case -9223372036854775808:
                v25 := GenS1{F0: -695, F1: 9223372036854775807, F2: 983}
                _ = v25
            default:
                var v26 int32 = -2147483648
                _ = v26
            }
            if (false && (3.14159 >= 1.5)) {
                switch (fi22 + (fi22 << ((18446744073709551615) & 63))) {
                case -822:
                    var v27 float64 = float64((((v2 * -2.25) / (v2 - (float32((((v3 + (v3 - (v3 - 1.5))) - 0.0) - float64(v2))) - (v2 * (3.14159 * v2))))) * 3.14159))
                    _ = v27
                default:
                    var v28 uint8 = 254
                    _ = v28
                }
                v19 = append(v19, 2147483647)
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 uint64 = uint64(v20)
                    _ = fi29
                    _ = v30
                }
            }
        }
        switch 351 {
        case -2147483648:
            switch ("" + v5) {
            case "foo bar":
                for fi31 := 0; fi31 < 10; fi31++ {
                    v5 = "hello"
                    switch v5 {
                    case "\n\t":
                        v2 += 1e10
                    default:
                        v19 = append(v19, (fi31 + (-(fi22))))
                    }
                    v2 = (v2 * 1e10)
                    if (v3 <= float64(v2)) {
                        var v32 uint64 = 42
                        _ = v32
                    } else {
                        var v33 float64 = 3.14159
                        var v34 uint16 = 0
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                }
            case "a":
                var v35 int8 = 83
                _ = v35
            default:
                if len(v19) > 0 {
                    v19[0] = 32767
                }
            }
        }
        if len(v19) > 0 {
            v19[0] = -1
        }
        _ = fi22
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v36 float32 = 0.0
    v36 += 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
