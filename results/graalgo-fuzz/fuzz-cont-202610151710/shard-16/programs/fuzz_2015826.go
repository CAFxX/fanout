package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
func genFn0(p0 int64, p1 int16) uint64 {
    v4 := GenS1{F0: 0, F1: 1, F2: 0}
    v5 := []int{42}
    for fi6 := 0; fi6 < 2; fi6++ {
        switch "foo bar" {
        case "\n\t":
            var v7 int8 = -85
            _ = v7
        case "hello":
            fmt.Println(p0, p0)
        case "":
            var v8 uint = 1
            _ = v8
        }
        _ = fi6
    }
    var v9 int64 = int64(p1)
    if true {
        var v10 int32 = -291
        v11 := GenS1{F0: 101, F1: 42, F2: 0}
        _ = v10
        _ = v11
    } else {
        switch 153 {
        case 32767:
            p1 = p1
        case 212:
            var v12 float32 = 1e10
            _ = v12
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            v5 = append(v5, 9223372036854775807)
            v14 := GenS2{F0: 0}
            for fi15 := 0; fi15 < 9; fi15++ {
                if len(v5) > 0 {
                    v5[0] = 9223372036854775807
                }
                if true {
                    var v16 int8 = -49
                    if false {
                        if (0 > -1) {
                            var v17 int32 = 0
                            if false {
                                v18 := GenS2{F0: 65535}
                                _ = v18
                            }
                            _ = v17
                        } else {
                            v5 = append(v5, (fi13 & -717))
                            switch fi13 {
                            case 2147483647:
                                p1 -= (p1 * 0)
                            default:
                                if (4294967295 < 4294967295) {
                                    var v19 int16 = p1
                                    _ = v19
                                }
                            }
                            fmt.Println(p0)
                        }
                        var v20 uint8 = 96
                        var v21 uint16 = 401
                        fmt.Println(v20, v21, v21)
                        var v22 uint = 255
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = v16
                } else {
                    p0 += (p0 + (v9 >> ((65535) & 63)))
                }
                fmt.Println(fi15, p0, p1)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v9
    return 810
}
func genFn1(p0 int16) int32 {
    for fi23 := 0; fi23 < 10; fi23++ {
        p0 += (p0 << ((303) & 63))
        var v24 uint64 = 42
        _ = fi23
        _ = v24
    }
    var v25 int8 = -128
    var v26 int32 = -2147483648
    _ = p0
    _ = v25
    _ = v26
    return v26
}

func main() {
    var v1 bool = false
    var v2 float32 = (1e10 * ((3.14159 * (1e10 / float32((1.5 / float64(3.14159))))) / float32((((-2.25 - 0.0) - 0.0) / float64((1e10 + 1e10))))))
    var v3 string = "a"
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := []int{}
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int16 = 32767
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := []int{fi29, (fi31 / ((fi31) | 1))}
            fmt.Println(fi31, v2)
            fmt.Println(fi29, fi29, v30)
            v30 = (v30 + (v30 ^ (v30 >> ((2) & 63))))
            v33 := GenS1{F0: 0, F1: 42, F2: 127}
            fmt.Println(fi31)
            if v1 {
                v3 = v3
                var v34 int64 = -9223372036854775808
                _ = v34
            }
            _ = len(v32)
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    var v35 float32 = v2
    fmt.Println(v35)
    v1 = true
    v36 := GenS3{F0: 125, F1: 1, F2: -995}
    var v37 uint = 209
    for fi38 := 0; fi38 < 19; fi38++ {
        switch fi38 {
        case -9223372036854775808:
            fmt.Println(fi38, v3)
        case 4294967295:
            v3 = v3
        case -2147483648:
            for fi39 := 0; fi39 < 15; fi39++ {
                var v40 int8 = 127
                if true {
                    v37 -= v37
                    fmt.Println(v1)
                    var v41 string = v3
                    _ = v41
                }
                _ = fi39
                _ = v40
            }
        }
        _ = fi38
    }
    var v42 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v42
}
