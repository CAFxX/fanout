package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int8; F1 uint16; F2 int32 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint64; F2 string }
type GenS6 struct { F0 int16; F1 int32; F2 int8 }
func genFn0() int8 {
    switch ("a" + ("a" + "\n\t")) {
    case "hello":
        v5 := []int{128, 926, -287}
        _ = v5
    default:
        var v6 uint32 = 65535
        _ = v6
    }
    var v7 uint = 491
    _ = v7
    return -123
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]uint32)
    var v3 int32 = -2147483648
    var v4 uint = 2
    var v8 int8 = -6
    var v9 int16 = 42
    if true {
        fmt.Println(v8)
        v1 += -1
    } else {
        var v10 int8 = (v8 * v8)
        if (-32768 != int16(v8)) {
            for fi11 := 0; fi11 < 19; fi11++ {
                v10 = v10
                v12 := GenS1{F0: 2147483648, F1: 558}
                _ = fi11
                _ = v12
            }
        } else {
            v13 := GenS1{F0: 42, F1: (v4 << ((v4) & 63))}
            var v14 float64 = (float64((3.14159 - 3.14159)) - 1.5)
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: 1, F1: 65535}
        _ = v10
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: 4294967295, F1: (v4 + v4)}
        _ = fi16
        _ = v17
    }
    var v18 float64 = 1e10
    v19 := GenS1{F0: 651, F1: 18446744073709551615}
    fmt.Println(v18)
    v20 := GenS1{F0: 42, F1: v4}
    v21 := GenS1{F0: 4294967295, F1: (v4 ^ v4)}
    v22 := GenS2{F0: (-(v1)), F1: (^(v3))}
    for fi23 := 0; fi23 < 18; fi23++ {
        if (!false) {
            if false {
                for fi24 := 0; fi24 < 5; fi24++ {
                    if false {
                        fmt.Println(fi23, fi23)
                    }
                    var v25 int64 = -9223372036854775808
                    switch 258 {
                    case -872:
                        if (((v25 % ((int64(v8)) | 1)) > 42) || (!(!(false && (0 < 2))))) {
                            v18 -= (0.1 * ((0.1 - v18) / v18))
                            v26 := GenS3{F0: (v8 << (((v4 / ((v4) | 1))) & 63)), F1: 905, F2: (v3 & v3)}
                            _ = v26
                        } else {
                            if false {
                                switch ("\n\t" + "a") {
                                case "\n\t":
                                    delete(v2, 42)
                                case "foo bar":
                                    if ((!(208 != v4)) && false) {
                                        v8 = (v8 / ((v8) | 1))
                                    } else {
                                        if ((v4 | (v4 - v4)) != v4) {
                                            switch (fi23 - fi24) {
                                            case -2147483648:
                                                v27 := GenS4{F0: 0, F1: 72}
                                                _ = v27
                                            }
                                            if true {
                                                if (("\n\t" + "a") == "a") {
                                                    var v28 int8 = int8(fi23)
                                                    v9 = 42
                                                    _ = v28
                                                } else {
                                                    v2[85] = 688
                                                    v1 += v25
                                                }
                                                v18 = v18
                                                var v29 float32 = 0.0
                                                _ = v29
                                            }
                                        }
                                    }
                                case "a":
                                    v30 := GenS5{F0: 2, F1: 0, F2: "a"}
                                    _ = v30
                                default:
                                    v9 -= v9
                                }
                                var v31 string = "\n\t"
                                _ = v31
                            }
                            v1 += v25
                        }
                    }
                    _ = fi24
                    _ = v25
                }
            }
        } else {
            v32 := GenS6{F0: 1, F1: v3, F2: 25}
            var v33 int8 = v8
            _ = v32
            _ = v33
        }
        var v34 int = fi23
        _ = fi23
        _ = v34
    }
    switch -622 {
    case 647:
        v4 = v4
    case 128:
        var v35 int8 = (^(v8))
        _ = v35
    default:
        var v36 float32 = (float32((v18 * float64(((1.5 - (-2.25 - 1.5)) * 0.0)))) * 1e10)
        _ = v36
    }
    var v37 string = "\n\t"
    for fi38 := 0; fi38 < 20; fi38++ {
        v1 = (v1 * (v1 + (-(v1))))
        if true {
            var v39 string = v37
            v40 := GenS1{F0: 807, F1: v4}
            _ = v39
            _ = v40
        }
        v9 -= int16(fi38)
        _ = fi38
    }
    fmt.Println(v8, v9, v8)
    var v41 string = v37
    v3 += v3
    if (v8 <= (v8 >> ((v4) & 63))) {
        switch 4294967295 {
        case 589:
            var v42 uint64 = 1
            _ = v42
        case 0:
            v37 = (v41 + v41)
        }
    } else {
        var v43 int64 = 0
        var v44 uint8 = 114
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v8
    _ = v9
}
