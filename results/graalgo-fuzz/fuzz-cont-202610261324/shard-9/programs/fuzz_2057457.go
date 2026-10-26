package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint = 163
    v2 := make(map[string]int)
    var v3 int8 = 1
    var v4 string = ("hello" + ("" + ("" + "\n\t")))
    var v5 float32 = 3.14159
    if false {
        v4 = v4
        fmt.Println(v5, v5, v5)
        v2[(v4 + "")] = 4294967295
    }
    v6 := GenS1{F0: "\n\t", F1: 2, F2: 4294967295}
    var v7 uint8 = 128
    switch v4 {
    case "a":
        fmt.Println(v7, v5)
    case "":
        v7 -= v7
    case "\n\t":
        switch -583 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 12; fi8++ {
                delete(v2, v4)
                _ = fi8
            }
        case 107:
            var v9 uint8 = v7
            _ = v9
        }
    default:
        v3 = v3
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        delete(v2, ((("foo bar" + v4) + v4) + ("" + (v4 + v4))))
        _ = fi10
    }
    v3 = v3
    if false {
        if false {
            fmt.Println(v7, v3, v1)
        }
    }
    v11 := []int{-527, 0}
    var v12 uint64 = 527
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int64 = -9223372036854775808
        switch fi13 {
        case -293:
            fmt.Println(v14)
        case 735:
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 uint16 = 1
                v17 := v2[(v4 + "\n\t")]
                if false {
                    v11 = append(v11, 32768)
                    v1 += v1
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        case 2147483647:
            fmt.Println(v14, v4, v4)
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            if (1 <= 32767) {
                v19 := GenS2{F0: (^(v12)), F1: 2}
                _ = v19
            }
            var v20 int32 = 42
            v3 = (-(v3))
            _ = fi18
            _ = v20
        }
        _ = fi13
        _ = v14
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS1{F0: "hello", F1: 42, F2: v1}
            v24 := GenS3{F0: 610, F1: 843, F2: (v3 / ((-110) | 1))}
            v5 += 1e10
            switch ("" + v4) {
            case "a":
                v25 := GenS1{F0: "a", F1: (v12 * v12), F2: (v1 ^ v1)}
                _ = v25
            case "foo bar":
                v4 = ""
            default:
                if ("hello" != v4) {
                    switch v4 {
                    case "a":
                        v5 += 1e10
                    case "foo bar":
                        v5 -= v5
                    }
                    v26 := GenS1{F0: ((v4 + v4) + ((v4 + ((((((v4 + (("foo bar" + "a") + (("foo bar" + v4) + v4))) + (v4 + v4)) + "hello") + v4) + v4) + "")) + "a")), F1: 4294967295, F2: (v1 ^ 65535)}
                    fmt.Println(fi22)
                    _ = v26
                } else {
                    v11 = append(v11, (fi21 % ((4294967295) | 1)))
                }
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    v4 = (v4 + v4)
    v5 = (v5 + (v5 - (1.5 + float32(1e10))))
    if (445 < 9223372036854775807) {
        for fi27 := 0; fi27 < 6; fi27++ {
            fmt.Println(v1, fi27, v5)
            _ = fi27
        }
        var v28 int32 = 1
        _ = v28
    } else {
        v29 := GenS4{F0: 2}
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            if (0.1 > 3.14159) {
                var v32 int8 = (v3 & v3)
                if false {
                    if (!false) {
                        var v33 uint16 = 897
                        v1 -= (^(v1))
                        v3 = v32
                        _ = v33
                    } else {
                        var v34 uint64 = (v12 | (v12 ^ (v12 - (^(v12)))))
                        _ = v34
                    }
                    var v35 float32 = (v5 - v5)
                    _ = v35
                } else {
                    var v36 int64 = -435
                    if true {
                        v32 += v32
                    } else {
                        var v37 uint64 = 4294967295
                        _ = v37
                    }
                    _ = v36
                }
                _ = v32
            } else {
                v7 += (^(v7))
                var v38 uint = uint(v7)
                _ = v38
            }
            _ = fi31
        }
        if false {
            v39 := GenS2{F0: v12, F1: 2147483647}
            if true {
                fmt.Println(v3, v5)
                v3 = (v3 | (v3 | (v3 + v3)))
            }
            _ = v39
        }
        _ = fi30
    }
    if false {
        switch (v4 + (v4 + (((((v4 + v4) + ("hello" + "foo bar")) + v4) + ("" + ("\n\t" + v4))) + v4))) {
        case "":
            if (440 >= 242) {
                fmt.Println(v4, v4, v4)
                for fi40 := 0; fi40 < 2; fi40++ {
                    v41 := GenS4{F0: 255}
                    for fi42 := 0; fi42 < 7; fi42++ {
                        v4 = ""
                        v11 = append(v11, -375)
                        fmt.Println(fi42, v5)
                        _ = fi42
                    }
                    _ = fi40
                    _ = v41
                }
                fmt.Println(v7, v1)
            } else {
                delete(v2, v4)
                v4 = "a"
                v4 = v4
            }
        }
        var v43 float64 = 1e10
        var v44 uint16 = 777
        _ = v43
        _ = v44
    } else {
        var v45 uint = uint(v7)
        for fi46 := 0; fi46 < 11; fi46++ {
            fmt.Println(v5, v3, v5)
            _ = fi46
        }
        v45 = v1
        _ = v45
    }
    var v47 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
