package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int64; F1 int32; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int32) int32 {
    if (false || (!(!false))) {
        var v6 int = 2147483647
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := []int{v6, fi7, fi7, v6}
            v9 := GenS1{F0: 2, F1: 271}
            var v10 int32 = p0
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        fmt.Println(p0)
        p0 -= p0
    }
    var v11 uint32 = 42
    v12 := GenS1{F0: 18446744073709551615, F1: 42}
    _ = p0
    _ = v11
    _ = v12
    return p0
}
func genFn1() int {
    if false {
        var v13 int = 970
        var v14 uint8 = 206
        var v15 float64 = (float64(3.14159) + (0.1 + (1e10 - 1.5)))
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 string = "\n\t"
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 167, F1: 2}
    var v19 int32 = genFn0(0)
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (!(2 == 128))
    v3 := make(map[string]string)
    var v4 uint8 = 2
    var v5 float32 = (-2.25 + (3.14159 + -2.25))
    _ = len(v3)
    var v20 int64 = -9223372036854775808
    v20 -= (v20 ^ v20)
    fmt.Println(v5, v2, v2)
    v5 -= v5
    var v21 uint64 = 1
    var v22 bool = (!(32767 > -32768))
    if false {
        if false {
            if (true && (-128 != -34)) {
                switch ("foo bar" + "hello") {
                case "\n\t":
                    var v23 uint32 = 296
                    _ = v23
                case "foo bar":
                    var v24 string = "\n\t"
                    _ = v24
                }
                v25 := GenS1{F0: (v21 / ((v21) | 1)), F1: v21}
                v1[4294967295] = 2147483647
                _ = v25
            } else {
                var v26 bool = v2
                _ = v26
            }
            var v27 int64 = (v20 - v20)
            _ = v27
        } else {
            if ((!(v22 && (1 < 2))) || v22) {
                v20 = (v20 << ((255) & 63))
                v28 := GenS2{F0: (-(v20)), F1: 0, F2: 304}
                for fi29 := 0; fi29 < 5; fi29++ {
                    var v30 float32 = float32(0.1)
                    switch (fi29 % (((fi29 >> ((2) & 63))) | 1)) {
                    case 2147483647:
                        if (310 >= 328) {
                            v31 := []int{}
                            for fi32 := 0; fi32 < 9; fi32++ {
                                v20 -= (^(v20))
                                var v33 int16 = 32767
                                var v34 float64 = (-2.25 - (1.5 + ((((-2.25 - ((1.5 * (3.14159 + ((1.5 * ((0.0 * 0.1) * (float64(0.1) + 0.0))) * ((1.5 - (0.0 - 3.14159)) * 0.1)))) - -2.25)) - 0.0) + (1e10 * 1.5)) * 0.0)))
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v20, v20, fi29)
                            _ = v31
                        } else {
                            v35 := []int{(^(fi29)), genFn1(), (fi29 / (((fi29 << ((0) & 63))) | 1))}
                            var v36 uint64 = v21
                            for fi37 := 0; fi37 < 10; fi37++ {
                                switch fi37 {
                                case 2147483647:
                                    var v38 bool = true
                                    _ = v38
                                case 1:
                                    _ = len(v35)
                                }
                                _ = fi37
                            }
                            _ = v35
                            _ = v36
                        }
                    case 1:
                        for fi39 := 0; fi39 < 17; fi39++ {
                            switch "hello" {
                            case "foo bar":
                                if (!v2) {
                                    v40 := v1[0]
                                    if (v2 || (((-1 < int64(fi39)) && (!(!true))) && false)) {
                                        v41 := GenS2{F0: v20, F1: -2147483648, F2: 0}
                                        _ = v41
                                    } else {
                                        v5 += (v5 + ((float32((1e10 - ((1.5 / (((0.1 + 1.5) / (3.14159 * 1e10)) - 1e10)) / 1.5))) - 3.14159) - (v5 + ((-2.25 * (float32(-2.25) / 3.14159)) + (v30 - 1e10)))))
                                        fmt.Println(fi29, v22, fi39)
                                        fmt.Println(fi39, v21, fi29)
                                    }
                                    _ = v40
                                } else {
                                    fmt.Println(fi29, v5, fi39)
                                }
                            case "hello":
                                var v42 float64 = (0.0 / 3.14159)
                                _ = v42
                            case "a":
                                v43 := GenS3{F0: uint64(v4), F1: 157}
                                _ = v43
                            }
                            fmt.Println(v30, v5)
                            var v44 string = "foo bar"
                            if false {
                                v45 := GenS3{F0: 42, F1: v4}
                                var v46 int8 = 42
                                _ = v45
                                _ = v46
                            } else {
                                v22 = false
                            }
                            _ = fi39
                            _ = v44
                        }
                    case -2147483648:
                        _ = len(v3)
                    }
                    v4 += v4
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
        }
        v47 := v3["\n\t"]
        _ = v47
    } else {
        var v48 string = "\n\t"
        if (0 < (v20 + (v20 >> ((65535) & 63)))) {
            if false {
                v22 = (1e10 >= 0.1)
                for fi49 := 0; fi49 < 16; fi49++ {
                    if true {
                        v4 = v4
                        v50 := []int{genFn1(), 4294967295, genFn1()}
                        v2 = (v2 || v2)
                        _ = v50
                    }
                    if true {
                        if (v2 && v22) {
                            v48 = "foo bar"
                        }
                        for fi51 := 0; fi51 < 19; fi51++ {
                            v21 += 42
                            var v52 float64 = 1.5
                            switch ("a" + v48) {
                            case "a":
                                var v53 int = fi49
                                _ = v53
                            case "\n\t":
                                fmt.Println(v5, v4, v22)
                            default:
                                var v54 int16 = -590
                                _ = v54
                            }
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        var v55 float32 = ((v5 / ((0.1 / 1.0) - v5)) / 1.0)
                        _ = v55
                    }
                    v20 -= v20
                    _ = fi49
                }
                var v56 int8 = -23
                _ = v56
            }
        } else {
            v57 := []int{4294967295, genFn1()}
            v3[v48] = "\n\t"
            _ = v57
        }
        _ = v48
    }
    fmt.Println(v21, v21, v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
