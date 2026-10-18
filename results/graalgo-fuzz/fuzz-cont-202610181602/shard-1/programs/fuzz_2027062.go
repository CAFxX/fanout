package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint; F1 int32; F2 int16 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 uint64; F2 int8 }
func genFn0() int64 {
    var v4 int = -1
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 0}
        _ = fi5
        _ = v6
    }
    var v7 int = -9223372036854775808
    _ = v4
    _ = v7
    return -1
}
func genFn1(p0 bool, p1 bool) string {
    switch "" {
    case "foo bar":
        v8 := []int{-22, 32767}
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS2{F0: 18446744073709551615, F1: 2147483648, F2: 0}
            _ = fi9
            _ = v10
        }
    case "hello":
        p0 = p1
    }
    p0 = p0
    _ = p0
    _ = p1
    return (("\n\t" + "a") + "")
}

func main() {
    var v1 int32 = 634
    v2 := make(map[string]uint32)
    var v3 uint64 = 311
    if (false || (32767 >= 399)) {
        fmt.Println(v3)
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := []int{(fi11 - -1), (fi11 / ((59) | 1)), -9223372036854775808, fi11}
            var v13 uint16 = 65535
            v14 := GenS1{F0: 1}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            v3 -= 873
            v3 = v3
            for fi16 := 0; fi16 < 9; fi16++ {
                v3 += v3
                var v17 int = fi15
                var v18 uint32 = 644
                v19 := GenS3{F0: 255, F1: 84}
                fmt.Println(v18, v3, v1)
                v3 -= (v3 | v3)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
    } else {
        v20 := GenS3{F0: 2, F1: 0}
        if true {
            _ = len(v2)
            switch 32767 {
            case 1:
                fmt.Println(v1)
            default:
                if true {
                    v21 := GenS4{F0: 42, F1: v1, F2: 259}
                    v3 = v3
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        if true {
                            switch "\n\t" {
                            case "foo bar":
                                var v23 int32 = v1
                                _ = v23
                            default:
                                if (fi22 < (fi22 - 127)) {
                                    for fi24 := 0; fi24 < 11; fi24++ {
                                        fmt.Println(fi24)
                                        _ = fi24
                                    }
                                } else {
                                    var v25 float64 = 1e10
                                    if true {
                                        v26 := []int{fi22, 1, fi22, 1}
                                        v25 = (v25 - v25)
                                        _ = v26
                                    }
                                    switch fi22 {
                                    case 32767:
                                        fmt.Println(v3)
                                    case -9223372036854775808:
                                        v1 = 879
                                    case 4294967295:
                                        var v27 int = 128
                                        _ = v27
                                    }
                                    _ = v25
                                }
                            }
                            for fi28 := 0; fi28 < 16; fi28++ {
                                if ((3.14159 + (1.5 / (0.1 + 0.0))) >= 3.14159) {
                                    v29 := GenS3{F0: 226, F1: 42}
                                    var v30 float64 = 1e10
                                    _ = v29
                                    _ = v30
                                }
                                v3 = v3
                                _ = fi28
                            }
                            fmt.Println(v3, fi22, fi22)
                        }
                        v31 := GenS5{F0: 1, F1: 42}
                        _ = fi22
                        _ = v31
                    }
                    var v32 uint = 621
                    _ = v32
                }
            }
            switch ("foo bar" + "foo bar") {
            case "":
                v33 := GenS6{F0: 0}
                _ = v33
            default:
                var v34 int8 = -128
                _ = v34
            }
        }
        v3 = 4294967295
        _ = v20
    }
    var v35 int16 = 0
    v36 := GenS7{F0: 0, F1: v3, F2: 100}
    var v37 bool = (false && false)
    v2["foo bar"] = 1
    for fi38 := 0; fi38 < 15; fi38++ {
        v2["\n\t"] = 605
        _ = fi38
    }
    var v39 float32 = (0.0 * (0.1 - 0.1))
    v35 += v35
    var v40 uint32 = 42
    var v41 uint8 = 255
    fmt.Println(v41, v41)
    fmt.Println(v37, v40)
    var v42 float64 = 0.0
    for fi43 := 0; fi43 < 7; fi43++ {
        delete(v2, ("hello" + "a"))
        var v44 uint64 = (^(v3))
        v45 := v2["a"]
        _ = fi43
        _ = v44
        _ = v45
    }
    var v46 uint = 891
    for fi47 := 0; fi47 < 10; fi47++ {
        v3 = (v3 ^ (v3 - v3))
        var v48 int = fi47
        _ = fi47
        _ = v48
    }
    fmt.Println(v40, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v46
}
