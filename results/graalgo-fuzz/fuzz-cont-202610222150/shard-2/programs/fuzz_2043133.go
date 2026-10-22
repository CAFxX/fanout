package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint }

func main() {
    v1 := []uint32{2147483647, 65535}
    var v2 uint16 = 65535
    v3 := []string{"\n\t"}
    var v4 bool = true
    v5 := []int{2147483647, 42, 4294967295, -2147483648}
    var v6 int64 = -1
    if (523 >= -32768) {
        v1 = append(v1, 0)
    } else {
        v7 := GenS1{F0: 475, F1: 693, F2: 4294967295}
        v2 += v2
        _ = v7
    }
    var v8 int32 = 2147483647
    var v9 string = "foo bar"
    var v10 int = 1
    fmt.Println(v4, v10, v6)
    switch "\n\t" {
    case "a":
        var v11 int64 = v6
        _ = v11
    }
    if v4 {
        var v12 float32 = ((1e10 * (0.0 - -2.25)) - 0.0)
        if (!v4) {
            v3 = append(v3, (v9 + v9))
            var v13 uint = 255
            var v14 int8 = -128
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 19; fi15++ {
                switch (("a" + v9) + ((v9 + v9) + "hello")) {
                case "foo bar":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        if false {
                            if (!(v6 == (v6 + 685))) {
                                if len(v3) > 0 {
                                    v3[0] = v9
                                }
                                v3 = append(v3, v9)
                                v9 = v9
                            }
                            v6 = (v6 ^ v6)
                        }
                        _ = fi16
                    }
                case "\n\t":
                    for fi17 := 0; fi17 < 14; fi17++ {
                        fmt.Println(v9)
                        _ = len(v5)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi17
                    }
                default:
                    if true {
                        var v18 float64 = -2.25
                        var v19 int8 = -128
                        _ = v18
                        _ = v19
                    }
                }
                var v20 uint32 = 42
                _ = fi15
                _ = v20
            }
            if v4 {
                for fi21 := 0; fi21 < 2; fi21++ {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v8 = 2147483647
                        fmt.Println(v12)
                        _ = fi22
                    }
                    fmt.Println(v4)
                    if ((v2 != (^(v2))) && v4) {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            var v24 float64 = 1e10
                            v25 := GenS1{F0: 905, F1: 0, F2: 1}
                            v4 = (902 > 0)
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = len(v3)
                    }
                    _ = fi21
                }
            } else {
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 uint8 = 0
                    for fi28 := 0; fi28 < 14; fi28++ {
                        fmt.Println(v10, v10)
                        var v29 bool = v4
                        v6 -= (v6 - 42)
                        v30 := GenS1{F0: 1, F1: 980, F2: 18446744073709551615}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
                v9 = v9
            }
            v31 := GenS1{F0: 18446744073709551615, F1: 65535, F2: 396}
            _ = v31
        }
        _ = v12
    } else {
        if v4 {
            if false {
                var v32 int = (v10 ^ -603)
                var v33 int64 = (v6 << ((1) & 63))
                _ = v32
                _ = v33
            } else {
                v34 := GenS1{F0: 0, F1: 0, F2: 4294967295}
                v35 := GenS1{F0: 621, F1: 4294967295, F2: 4294967295}
                _ = v34
                _ = v35
            }
        }
        v36 := GenS1{F0: 42, F1: 18446744073709551615, F2: 1}
        switch v9 {
        case "a":
            for fi37 := 0; fi37 < 14; fi37++ {
                if (v4 && v4) {
                    v10 -= (v10 % (((v10 >> ((402) & 63))) | 1))
                    fmt.Println(v10)
                } else {
                    if v4 {
                        v38 := GenS1{F0: 471, F1: 302, F2: 257}
                        _ = v38
                    } else {
                        var v39 float32 = 1e10
                        for fi40 := 0; fi40 < 5; fi40++ {
                            fmt.Println(v6, v39)
                            v4 = v4
                            if v4 {
                                fmt.Println(fi37, v9, v39)
                            } else {
                                fmt.Println(fi40)
                            }
                            _ = fi40
                        }
                        _ = v39
                    }
                    v10 = v10
                }
                v4 = ((v2 != (v2 | (^(v2)))) && v4)
                _ = fi37
            }
        case "foo bar":
            v41 := GenS2{F0: -1, F1: 44}
            _ = v41
        case "hello":
            if (true || v4) {
                v42 := GenS2{F0: v6, F1: 18446744073709551615}
                _ = v42
            }
        }
        _ = v36
    }
    fmt.Println(v2, v9)
    var v43 string = ("" + (("\n\t" + v9) + (("a" + v9) + v9)))
    v44 := GenS1{F0: 492, F1: 42, F2: 42}
    v4 = v4
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
