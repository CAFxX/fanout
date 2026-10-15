package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
func genFn0(p0 int64, p1 uint32) bool {
    for fi7 := 0; fi7 < 8; fi7++ {
        p1 -= 132
        var v8 uint8 = 255
        _ = fi7
        _ = v8
    }
    var v9 float64 = float64(float32(((-2.25 - ((0.0 + float64((0.0 - 0.1))) / 0.1)) + (float64(-2.25) / ((1.5 / float64(float32(3.14159))) - (0.1 / 1.0))))))
    fmt.Println(v9, v9, v9)
    _ = p0
    _ = p1
    _ = v9
    return false
}
func genFn1() uint32 {
    var v10 int32 = 2147483647
    if (((465 == 2) && false) || false) {
        if ((!(!false)) && (((0.1 + 0.1) * 1e10) > float64(0.0))) {
            v10 = v10
            var v11 int16 = 32767
            _ = v11
        }
        v12 := GenS1{F0: 1, F1: 42, F2: 61}
        _ = v12
    } else {
        var v13 bool = ((((0.1 + 1.5) * ((1.5 - (1e10 - 1e10)) - 1e10)) < -2.25) && ((v10 | v10) >= (-(v10))))
        var v14 int8 = 1
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 1, F1: -834, F2: 235}
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            if (599 != 1) {
                var v18 uint64 = 42
                v19 := GenS1{F0: fi16, F1: -9223372036854775808, F2: 2}
                v18 += 2
                _ = v18
                _ = v19
            } else {
                v20 := []int{}
                var v21 float64 = 3.14159
                v20 = append(v20, 763)
                var v22 int64 = -253
                v10 = v10
                _ = v20
                _ = v21
                _ = v22
            }
            if (153 != 127) {
                v23 := GenS1{F0: (^(fi17)), F1: -749, F2: 127}
                if false {
                    var v24 uint16 = 32767
                    _ = v24
                } else {
                    var v25 bool = false
                    v25 = (968 != 2)
                    v26 := GenS1{F0: fi16, F1: 42, F2: 128}
                    v27 := []int{}
                    switch (fi16 | (fi17 / ((fi16) | 1))) {
                    case -1:
                        if (20 > 2) {
                            var v28 int32 = v10
                            v29 := GenS1{F0: fi16, F1: 42, F2: 19}
                            _ = v28
                            _ = v29
                        } else {
                            var v30 float64 = -2.25
                            var v31 uint = 347
                            if false {
                                v25 = ((9223372036854775807 != -252) && v25)
                            } else {
                                _ = len(v27)
                            }
                            _ = v30
                            _ = v31
                        }
                    case -9223372036854775808:
                        var v32 uint64 = 18446744073709551615
                        var v33 uint = 2
                        _ = v32
                        _ = v33
                    default:
                        _ = len(v27)
                    }
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v23
            }
            _ = fi17
        }
        v10 = (v10 | (v10 | v10))
        fmt.Println(fi16)
        _ = fi16
    }
    _ = v10
    _ = v15
    return 8
}

func main() {
    v1 := []string{("a" + "a")}
    var v2 float64 = ((1e10 - -2.25) / float64(0.1))
    v3 := make(map[int]string)
    v4 := []uint32{}
    v5 := []string{"hello"}
    v6 := make(map[int]string)
    v34 := v3[-2147483648]
    switch (((v34 + "") + (((v34 + v34) + "a") + v34)) + v34) {
    case "":
        fmt.Println(v34, v34)
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = (("" + v34) + "")
        }
    }
    var v35 uint = 4294967295
    var v36 uint8 = 207
    v6[42] = (v34 + (((((v34 + ("" + v34)) + v34) + "") + "\n\t") + v34))
    v36 -= v36
    fmt.Println(v35)
    v37 := GenS1{F0: 925, F1: 42, F2: (v36 ^ (^(v36)))}
    v34 = "\n\t"
    var v38 uint = (^(v35))
    v36 -= (v36 ^ (v36 << (((v35 << (((v38 | (v35 % ((v38) | 1)))) & 63))) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
