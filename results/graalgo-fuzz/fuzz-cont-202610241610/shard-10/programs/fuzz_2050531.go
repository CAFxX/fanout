package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int16) uint8 {
    var v4 int64 = 939
    if (!(!false)) {
        var v5 int64 = (v4 >> ((255) & 63))
        if true {
            if false {
                var v6 int = 127
                v5 += v4
                v6 += v6
                _ = v6
            } else {
                fmt.Println(v5, p0, p0)
            }
            v7 := GenS1{F0: "foo bar"}
            var v8 bool = true
            _ = v7
            _ = v8
        } else {
            p1 = (p1 & p1)
            p1 = p1
        }
        if ((v4 % (((v4 ^ 42)) | 1)) < v4) {
            p0 = 1
            v5 -= 1
            fmt.Println(v4)
        }
        _ = v5
    } else {
        if false {
            v9 := GenS1{F0: "a"}
            var v10 uint16 = 28
            fmt.Println(v4, p1, v4)
            _ = v10
            _ = v9
        }
        v4 += 1
        switch 42 {
        case 724:
            v11 := GenS2{F0: 42, F1: 993}
            _ = v11
        default:
            if (2147483647 > 674) {
                v4 = -9223372036854775808
                var v12 uint = 42
                v13 := []int{}
                _ = v12
                _ = v13
            }
        }
    }
    var v14 int = 128
    _ = p0
    _ = p1
    _ = v14
    _ = v4
    return 42
}
func genFn1(p0 int8) bool {
    var v15 int32 = 2147483647
    var v16 int16 = 1000
    v17 := []int{}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return ((v15 % ((v15) | 1)) <= v15)
}

func main() {
    var v1 bool = (true && (0 > 65535))
    var v2 int = 9223372036854775807
    var v3 int32 = -1
    if true {
        var v18 int64 = 1
        v19 := []int{(v2 << ((4294967295) & 63)), -656, 32768, (v2 << ((42) & 63))}
        _ = v18
        _ = v19
    } else {
        var v20 string = "foo bar"
        v21 := []int{-9223372036854775808, (v2 ^ (v2 ^ 0)), (v2 + v2)}
        v20 = "foo bar"
        _ = v20
        _ = v21
    }
    v2 = (v2 << ((230) & 63))
    v22 := GenS3{F0: v3, F1: 4294967295, F2: 313}
    switch "foo bar" {
    case "foo bar":
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := []int{(fi23 & (v2 * fi23))}
            _ = fi23
            _ = v24
        }
    case "":
        switch ("" + ("" + "a")) {
        case "":
            v25 := GenS2{F0: v3, F1: 619}
            _ = v25
        case "foo bar":
            fmt.Println(v1, v3, v2)
        case "hello":
            if (v1 && v1) {
                switch v2 {
                case 955:
                    v26 := []int{}
                    _ = v26
                case 9223372036854775807:
                    v27 := []int{32768, v2}
                    _ = v27
                }
            }
        default:
            var v28 bool = ((((-1 >= 42) && true) || (v1 && v1)) && v1)
            _ = v28
        }
    }
    switch (v2 ^ (v2 | 9223372036854775807)) {
    case 32768:
        if (true && ((-568 <= 1) || (!v1))) {
            var v29 uint16 = 65535
            v30 := []int{v2}
            v29 -= (^(v29))
            _ = v29
            _ = v30
        }
    }
    v31 := []int{(v2 / (((v2 * v2)) | 1)), (v2 ^ (v2 * v2)), (-(v2))}
    v32 := GenS1{F0: ("foo bar" + "a")}
    for fi33 := 0; fi33 < 2; fi33++ {
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
            fmt.Println(fi34, v2, fi34)
            if ((0 <= 4294967295) && (!(!(!(v1 || v1))))) {
                v35 := GenS4{F0: 2}
                v31 = append(v31, -9223372036854775808)
                switch (("a" + (("hello" + "hello") + "foo bar")) + ("foo bar" + "")) {
                case "a":
                    switch v2 {
                    case -575:
                        fmt.Println(v3, v3)
                    case 867:
                        if v1 {
                            var v36 int64 = -1
                            switch "" {
                            case "hello":
                                var v37 float32 = 0.0
                                _ = v37
                            case "foo bar":
                                var v38 int = v2
                                _ = v38
                            default:
                                var v39 bool = false
                                _ = v39
                            }
                            switch (-(fi34)) {
                            case 1:
                                var v40 uint8 = 0
                                _ = v40
                            }
                            _ = v36
                        }
                    case -536:
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                    default:
                        var v41 uint64 = 4294967295
                        _ = v41
                    }
                case "foo bar":
                    var v42 uint32 = 4294967295
                    _ = v42
                case "hello":
                    switch "hello" {
                    case "":
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                    }
                default:
                    var v43 int = -9223372036854775808
                    _ = v43
                }
                fmt.Println(fi33)
                _ = v35
            }
            _ = fi34
        }
        _ = fi33
    }
    if true {
        if len(v31) > 0 {
            v31[0] = (v2 - v2)
        }
        if len(v31) > 0 {
            v31[0] = (v2 % ((v2) | 1))
        }
    }
    v3 = v3
    v2 = 42
    v44 := GenS2{F0: v3, F1: 0}
    _ = len(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v44
}
