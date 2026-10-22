package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 uint16, p1 int32, p2 bool) uint64 {
    p1 += 725
    v6 := GenS1{F0: -9223372036854775808}
    var v7 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() int8 {
    v8 := GenS2{F0: -590}
    var v9 uint32 = 1
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int64 = -79
        var v12 uint = 65535
        v12 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v8
    _ = v9
    return 86
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 263
    var v3 uint16 = 1
    var v4 uint = v2
    v5 := make(map[string]int)
    fmt.Println(v3)
    var v13 bool = (32767 < -1)
    v14 := GenS1{F0: 630}
    fmt.Println(v13)
    fmt.Println(v4, v4)
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            v3 -= v3
            switch "" {
            case "\n\t":
                var v17 uint = (v2 | v4)
                _ = v17
            case "hello":
                v5["a"] = -1
            }
            _ = fi16
        }
        if (!((v2 & v4) >= 542)) {
            fmt.Println(fi15, v13)
            v18 := []int{(fi15 % (((fi15 % ((-1) | 1))) | 1)), fi15}
            v18 = append(v18, 128)
            _ = v18
        } else {
            v19 := []int{fi15, (fi15 - (fi15 % ((fi15) | 1))), 187}
            v20 := v5[(("" + "") + (("foo bar" + ("a" + "\n\t")) + "hello"))]
            var v21 uint8 = 42
            _ = v19
            _ = v20
            _ = v21
        }
        switch 0 {
        case -1:
            var v22 uint64 = 702
            _ = v22
        }
        _ = fi15
    }
    switch 128 {
    case 32767:
        v23 := []int{}
        _ = v23
    }
    _ = len(v5)
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS3{F0: 250}
        _ = fi24
        _ = v25
    }
    switch 9223372036854775807 {
    case 127:
        fmt.Println(v13, v13)
    case -1:
        v2 = (^(v2))
    default:
        if false {
            v26 := GenS2{F0: -1}
            if v13 {
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 float32 = float32(float64(0.0))
                    var v29 int32 = 1
                    if (18446744073709551615 >= 556) {
                        fmt.Println(v2, v4)
                        switch (((("a" + "\n\t") + ("" + ("a" + ("foo bar" + ("" + "\n\t"))))) + "a") + "foo bar") {
                        case "\n\t":
                            if (v28 > (((0.1 / v28) * ((v28 * (v28 / 1.5)) + 0.1)) / v28)) {
                                if (genFn0(782, 1, (0 <= 1)) > genFn0(32767, -2147483648, (v2 != 18446744073709551615))) {
                                    for fi30 := 0; fi30 < 12; fi30++ {
                                        v31 := GenS4{F0: -429, F1: genFn0((v3 - 0), int32(fi30), v13)}
                                        _ = fi30
                                        _ = v31
                                    }
                                } else {
                                    fmt.Println(v13)
                                    if v13 {
                                        var v32 int8 = -128
                                        _ = v32
                                    }
                                }
                                v4 = (v4 / ((v2) | 1))
                                v3 = 42
                            }
                        case "hello":
                            var v33 string = ((("a" + "foo bar") + ("\n\t" + "\n\t")) + "hello")
                            _ = v33
                        default:
                            for fi34 := 0; fi34 < 4; fi34++ {
                                v28 += 0.1
                                _ = fi34
                            }
                        }
                    } else {
                        if v13 {
                            v1["\n\t"] = "a"
                            v28 += 1.5
                            v4 -= v2
                        } else {
                            switch ("a" + "foo bar") {
                            case "hello":
                                v35 := []int{fi27, fi27, fi27, 4294967295}
                                _ = v35
                            }
                            v2 = (v4 - v2)
                            if false {
                                v4 = (^(v4))
                                if v13 {
                                    v3 += (v3 % ((42) | 1))
                                }
                            }
                        }
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v13 = (!(1 > 42))
            }
            _ = v26
        } else {
            var v36 float64 = (3.14159 / 1.0)
            _ = v36
        }
    }
    var v37 float32 = 0.1
    var v38 int16 = 1000
    for fi39 := 0; fi39 < 19; fi39++ {
        var v40 int32 = 815
        v4 = (v2 + (v2 ^ v2))
        _ = fi39
        _ = v40
    }
    v41 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
