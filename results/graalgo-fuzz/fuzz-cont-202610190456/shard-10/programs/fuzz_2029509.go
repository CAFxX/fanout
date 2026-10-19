package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint = 0
    var v2 int32 = -245
    var v3 bool = false
    var v4 bool = (true || v3)
    v5 := make(map[string]int)
    v3 = (false && (-1 == 42))
    var v6 uint16 = uint16(v1)
    var v7 uint = (v1 / ((v1) | 1))
    if false {
        var v8 int32 = (v2 / ((v2) | 1))
        var v9 int64 = -9223372036854775808
        _ = v8
        _ = v9
    }
    switch 496 {
    case 738:
        v2 = -916
    case 4294967295:
        var v10 int16 = 1000
        _ = v10
    }
    if v3 {
        if (((v1 & 655) > (v1 % (((^(v1))) | 1))) && false) {
            v2 += (v2 % ((840) | 1))
            v3 = (!v4)
            var v11 int16 = -112
            _ = v11
        } else {
            var v12 uint16 = v6
            v7 += v1
            _ = v12
        }
        fmt.Println(v3, v6)
    } else {
        var v13 int64 = -1
        _ = v13
    }
    fmt.Println(v2, v1, v1)
    v14 := GenS1{F0: 0, F1: 42, F2: 2}
    v2 -= v2
    v15 := GenS1{F0: 0, F1: 2, F2: 4294967295}
    if ((!(255 >= 128)) || (1 == -1)) {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int8 = -31
            var v18 float64 = -2.25
            v17 = (^(v17))
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(v2, v4)
            _ = fi19
        }
        if (false || (!(2 < 4294967295))) {
            v20 := GenS2{F0: v6}
            var v21 uint = 1
            for fi22 := 0; fi22 < 7; fi22++ {
                for fi23 := 0; fi23 < 17; fi23++ {
                    if true {
                        switch "hello" {
                        case "\n\t":
                            delete(v5, "foo bar")
                            v24 := GenS3{F0: (-(v2)), F1: 32767}
                            _ = v24
                        case "hello":
                            if (!((1e10 * 3.14159) < (3.14159 / (3.14159 * -2.25)))) {
                                var v25 uint32 = 682
                                v25 = v25
                                switch (fi23 << (((v21 * (v7 >> (((v1 / ((42) | 1))) & 63)))) & 63)) {
                                case -94:
                                    v2 -= v2
                                case 32767:
                                    if (!(!(v4 && (v6 == v6)))) {
                                        switch ("" + (("" + "foo bar") + "foo bar")) {
                                        case "\n\t":
                                            if v4 {
                                                v5["foo bar"] = 4294967295
                                            }
                                            v25 -= 684
                                            var v26 string = "hello"
                                            v5[(v26 + v26)] = (fi22 * (fi23 ^ fi23))
                                            var v27 float32 = -2.25
                                            v28 := []int{fi22, (fi23 << ((2) & 63)), fi22, (^(fi22))}
                                            var v29 int = fi23
                                            v30 := GenS2{F0: (v6 + v6)}
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                    } else {
                                        var v31 float32 = (((0.0 / 0.1) / (((-2.25 / -2.25) * 1e10) + ((0.1 / 0.1) - (1e10 - (-2.25 * 1e10))))) / (float32(-2.25) - float32((((-2.25 + 1e10) * -2.25) + -2.25))))
                                        if (false && ((0 != (v25 / ((v25) | 1))) && v4)) {
                                            fmt.Println(v4, v21)
                                            v7 = 255
                                        }
                                        _ = v31
                                    }
                                }
                                _ = v25
                            }
                        }
                    }
                    _ = fi23
                }
                for fi32 := 0; fi32 < 19; fi32++ {
                    v2 += v2
                    fmt.Println(v4)
                    var v33 int64 = 9223372036854775807
                    _ = fi32
                    _ = v33
                }
                var v34 uint = v7
                _ = fi22
                _ = v34
            }
            _ = v20
            _ = v21
        }
    } else {
        fmt.Println(v4, v3, v2)
    }
    fmt.Println(v2, v1, v1)
    var v35 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
