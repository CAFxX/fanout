package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0() uint8 {
    v6 := GenS1{F0: 1, F1: 15}
    var v7 bool = (0 < -9223372036854775808)
    if ((2 == 758) || v7) {
        var v8 int = 503
        _ = v8
    }
    v7 = (-2147483648 == -704)
    if false {
        if false {
            var v9 uint16 = 1
            v9 = 32767
            v10 := []int{663, 127}
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := []int{}
                v13 := GenS1{F0: -1, F1: 255}
                v7 = (1 != 255)
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    } else {
        if (false || v7) {
            var v14 uint = 318
            v15 := []int{-596}
            _ = v14
            _ = v15
        }
    }
    _ = v6
    _ = v7
    return 255
}
func genFn1() int64 {
    var v16 uint32 = 4294967295
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int16 = 236
        v16 -= (v16 | v16)
        _ = fi17
        _ = v18
    }
    v19 := []int{1, -9223372036854775808}
    _ = v16
    _ = v19
    return 1
}

func main() {
    v1 := []string{}
    var v2 int64 = 497
    var v3 float64 = 3.14159
    v4 := []string{("hello" + "\n\t")}
    var v5 uint64 = 42
    for fi20 := 0; fi20 < 4; fi20++ {
        v3 -= (float64(float32((v3 * v3))) / 1.0)
        var v21 int64 = v2
        _ = fi20
        _ = v21
    }
    var v22 int32 = -2147483648
    for fi23 := 0; fi23 < 20; fi23++ {
        v3 -= 3.14159
        v22 = (v22 - v22)
        if (255 == 1) {
            for fi24 := 0; fi24 < 11; fi24++ {
                if (v3 < ((0.1 * v3) / v3)) {
                    fmt.Println(fi23, v5, fi23)
                    if true {
                        var v25 int16 = 278
                        v5 += (v5 << ((42) & 63))
                        _ = v25
                    }
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + "\n\t")
                    }
                }
                _ = fi24
            }
            var v26 bool = true
            v5 = (v5 / (((v5 ^ 4294967295)) | 1))
            _ = v26
        }
        _ = fi23
    }
    var v27 uint = 997
    v22 = v22
    var v28 int16 = 1000
    if (false || true) {
        v27 = (v27 + v27)
    }
    if true {
        v3 -= -2.25
        v27 = (^(v27))
    } else {
        v3 += v3
        v27 = v27
        v5 -= v5
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    switch (("a" + "foo bar") + ("\n\t" + "\n\t")) {
    case "a":
        if (4294967295 == 1) {
            fmt.Println(v5, v5, v22)
            if (255 <= 2) {
                var v29 uint8 = 1
                if (!true) {
                    for fi30 := 0; fi30 < 6; fi30++ {
                        for fi31 := 0; fi31 < 6; fi31++ {
                            if (true && (true && (((^(v5)) == (^(v5))) && (v27 <= v27)))) {
                                if (0 == v29) {
                                    fmt.Println(fi30, v27)
                                    var v32 uint16 = 65535
                                    _ = v32
                                } else {
                                    v1 = append(v1, "hello")
                                    fmt.Println(fi30, v2)
                                    if false {
                                        v5 = v5
                                    } else {
                                        v33 := GenS1{F0: -65, F1: 255}
                                        fmt.Println(v29, v3, v29)
                                        if (!true) {
                                            var v34 int8 = 42
                                            var v35 int8 = (^(v34))
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = v33
                                    }
                                }
                                v36 := GenS2{F0: (!true)}
                                _ = v36
                            }
                            _ = fi31
                        }
                        _ = fi30
                    }
                } else {
                    v3 -= v3
                }
                _ = v29
            } else {
                var v37 int8 = 27
                v38 := GenS3{F0: v28}
                if ((-2.25 + 0.1) < (1e10 * (v3 + (v3 - (0.0 - v3))))) {
                    fmt.Println(v27)
                    _ = len(v1)
                    v39 := GenS2{F0: (false && false)}
                    _ = v39
                } else {
                    v3 += 1e10
                    v3 += v3
                }
                _ = v37
                _ = v38
            }
        }
    }
    v27 = (v27 + v27)
    v40 := GenS2{F0: (2147483647 > 4294967295)}
    v41 := GenS3{F0: (v28 / (((^(v28))) | 1))}
    if false {
        var v42 string = (("a" + "a") + "a")
        _ = v42
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
