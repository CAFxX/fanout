package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int8 }
func genFn0() int32 {
    v7 := GenS1{F0: ((1e10 - float32(((0.0 * (-2.25 * 0.1)) - (1.5 + 3.14159)))) == 1e10)}
    v8 := GenS1{F0: ((!(!(-7 >= 42))) && true)}
    var v9 int = 127
    var v10 string = ("a" + ("hello" + "foo bar"))
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint = 4294967295
    v2 := []string{}
    v3 := []uint32{}
    var v4 uint16 = 255
    var v5 uint32 = 498
    var v6 uint32 = uint32(v4)
    for fi11 := 0; fi11 < 2; fi11++ {
        if (uint(v4) == (^(v1))) {
            fmt.Println(v5, v5, v1)
        } else {
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS2{F0: -470, F1: -166}
                var v14 float32 = float32(float64(1e10))
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        v2 = append(v2, "hello")
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi11
    }
    if true {
        var v15 int64 = 42
        _ = v15
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 bool = false
        var v18 bool = v17
        var v19 int32 = genFn0()
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: (2147483647 <= 0)}
    fmt.Println(v1)
    fmt.Println(v6, v4)
    if (58 != 42) {
        v21 := GenS2{F0: 42, F1: -1}
        v2 = append(v2, "a")
        if false {
            if false {
                switch 128 {
                case 42:
                    fmt.Println(v4, v5)
                case -2147483648:
                    switch "hello" {
                    case "hello":
                        if len(v3) > 0 {
                            v3[0] = (^(v6))
                        }
                    case "foo bar":
                        if true {
                            for fi22 := 0; fi22 < 19; fi22++ {
                                v3 = append(v3, v6)
                                if ((v5 & v5) >= (v5 << (((v1 ^ v1)) & 63))) {
                                    v23 := GenS3{F0: (^(v4)), F1: 0}
                                    var v24 uint16 = v4
                                    _ = v23
                                    _ = v24
                                }
                                _ = fi22
                            }
                            _ = len(v3)
                            var v25 uint16 = v4
                            _ = v25
                        }
                    }
                default:
                    _ = len(v3)
                }
                fmt.Println(v5)
            } else {
                v5 = v5
                if ((v4 >> (((v1 | (v1 + 229))) & 63)) < (v4 / ((42) | 1))) {
                    v1 = v1
                } else {
                    if false {
                        var v26 string = (("hello" + "\n\t") + "hello")
                        for fi27 := 0; fi27 < 10; fi27++ {
                            v5 -= v5
                            v28 := GenS1{F0: false}
                            v29 := GenS4{F0: -128}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        fmt.Println(v4, v5)
                        _ = v26
                    } else {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            v6 = v6
                            _ = fi30
                        }
                    }
                }
                v6 += 0
            }
        } else {
            var v31 int = 2147483647
            v2 = append(v2, "\n\t")
            fmt.Println(v6, v5, v4)
            _ = v31
        }
        _ = v21
    } else {
        v5 -= v5
        v6 = v6
        v1 += v1
    }
    v1 = 4294967295
    if ((v6 >> (((v1 >> ((uint(v6)) & 63))) & 63)) < v5) {
        var v32 uint32 = (v6 & (v6 & (v6 ^ (v6 + v6))))
        _ = v32
    } else {
        switch "a" {
        case "\n\t":
            v5 = uint32(v4)
        case "foo bar":
            fmt.Println(v4)
        default:
            var v33 int16 = 119
            _ = v33
        }
        fmt.Println(v1)
    }
    var v34 int32 = -1
    if false {
        for fi35 := 0; fi35 < 14; fi35++ {
            switch (fi35 / ((-1) | 1)) {
            case 9223372036854775807:
                for fi36 := 0; fi36 < 17; fi36++ {
                    var v37 int = (fi36 / (((fi35 * fi35)) | 1))
                    v38 := GenS3{F0: v4, F1: 571}
                    v39 := GenS3{F0: v4, F1: 4294967295}
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v39
                }
            case -463:
                v34 = -624
            case 0:
                var v40 int64 = 1
                _ = v40
            }
            _ = fi35
        }
    } else {
        for fi41 := 0; fi41 < 11; fi41++ {
            for fi42 := 0; fi42 < 12; fi42++ {
                fmt.Println(fi41, v4, v6)
                _ = fi42
            }
            _ = fi41
        }
    }
    v43 := GenS1{F0: (128 < 251)}
    var v44 uint16 = (^(v4))
    v44 = uint16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
