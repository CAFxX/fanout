package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    var v3 float64 = v1
    var v4 int8 = 12
    var v5 float32 = -2.25
    switch ("a" + "a") {
    case "":
        v6 := GenS1{F0: 134, F1: 892}
        _ = v6
    case "hello":
        var v7 uint64 = 665
        _ = v7
    case "a":
        v1 = -2.25
    }
    delete(v2, "a")
    var v8 uint8 = 255
    switch "hello" {
    case "":
        v9 := []int{904, -2147483648}
        _ = v9
    case "\n\t":
        v10 := GenS2{F0: -1}
        _ = v10
    case "foo bar":
        switch 4294967295 {
        case -102:
            if false {
                if true {
                    var v11 float32 = v5
                    v12 := GenS3{F0: v4}
                    v3 = v3
                    _ = v11
                    _ = v12
                } else {
                    if ((1 >= -32768) && true) {
                        fmt.Println(v4, v3)
                        v1 -= v3
                        var v13 int32 = -2147483648
                        _ = v13
                    }
                    var v14 uint16 = 255
                    v14 = 828
                    _ = v14
                }
            } else {
                fmt.Println(v5, v3, v5)
                var v15 uint16 = 0
                v3 -= 1.5
                _ = v15
            }
        }
    }
    switch (("" + "hello") + ("foo bar" + "a")) {
    case "hello":
        var v16 float64 = v1
        _ = v16
    case "":
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v4, v8)
        case "\n\t":
            var v17 float64 = v1
            _ = v17
        case "":
            var v18 bool = true
            _ = v18
        default:
            v2[((("\n\t" + "") + (("hello" + (("a" + ("" + "foo bar")) + "a")) + "\n\t")) + (("\n\t" + "foo bar") + (("foo bar" + ("" + "hello")) + "a")))] = 42
        }
    case "foo bar":
        fmt.Println(v5, v8, v1)
    default:
        v19 := GenS4{F0: -1}
        _ = v19
    }
    var v20 uint32 = 102
    var v21 uint = 255
    if false {
        for fi22 := 0; fi22 < 18; fi22++ {
            if (((((((v1 - ((0.0 * v3) + (v3 + (v1 * v3)))) / (v1 / ((-2.25 + v1) - v1))) + v3) - 1e10) * (((float64(v5) * v3) + 0.0) - ((((((3.14159 - v1) + 1e10) - v3) + 1.5) * ((float64(v5) / ((-2.25 + (1e10 + 0.0)) - v1)) / ((v3 / ((((3.14159 / v1) - ((3.14159 + v3) * (float64(v5) / v3))) + 0.0) + ((float64(((float32(((0.1 * float64(v5)) * 1e10)) + 1e10) * v5)) / v1) / v1))) - float64(-2.25)))) + ((v1 / -2.25) + 3.14159)))) * 3.14159) > (v1 + v3)) {
                v20 = (v20 - (^(v20)))
                v23 := GenS4{F0: 181}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 16; fi24++ {
                    v25 := []int{(fi22 ^ (fi24 << (((v21 ^ (v21 << (((v21 - v21)) & 63)))) & 63))), fi22, -1, fi24}
                    var v26 float64 = (v1 - 1.5)
                    var v27 uint = v21
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            _ = fi22
        }
        v3 = (1.5 - 3.14159)
        switch "" {
        case "hello":
            v28 := v2["foo bar"]
            _ = v28
        default:
            v29 := []int{127}
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        if (2 >= 255) {
            v1 -= (((1e10 + v3) / 1e10) + (3.14159 + 0.1))
            v21 = v21
            v3 = v3
        }
        fmt.Println(v8)
        if (fi30 != -501) {
            for fi31 := 0; fi31 < 17; fi31++ {
                delete(v2, "")
                v20 = 984
                if false {
                    v5 -= v5
                    if true {
                        v32 := GenS3{F0: (v4 ^ (v4 / (((v4 << ((37) & 63))) | 1)))}
                        fmt.Println(v21)
                        v33 := GenS5{F0: v8, F1: fi31}
                        _ = v32
                        _ = v33
                    }
                    switch (fi31 % ((fi30) | 1)) {
                    case 885:
                        var v34 int16 = 42
                        _ = v34
                    case -968:
                        var v35 int64 = 42
                        _ = v35
                    }
                } else {
                    if false {
                        v5 = 1e10
                        if true {
                            switch ("foo bar" + "") {
                            case "a":
                                v36 := []int{}
                                _ = v36
                            case "\n\t":
                                if (!("" != "foo bar")) {
                                    v5 -= 1e10
                                }
                            default:
                                fmt.Println(v3)
                            }
                        }
                        var v37 int16 = 261
                        _ = v37
                    } else {
                        v38 := GenS6{F0: -255}
                        var v39 int8 = v4
                        fmt.Println(v5, v21)
                        _ = v38
                        _ = v39
                    }
                    switch (fi31 * (fi30 * (fi30 | 79))) {
                    case -503:
                        v40 := GenS6{F0: 1}
                        _ = v40
                    case 32768:
                        var v41 int = (fi30 & (-(fi31)))
                        _ = v41
                    case -354:
                        var v42 uint32 = (v20 ^ (v20 % (((v20 + (v20 << (((^(v21))) & 63)))) | 1)))
                        _ = v42
                    default:
                        var v43 int8 = (v4 ^ -1)
                        _ = v43
                    }
                }
                v44 := v2["foo bar"]
                _ = fi31
                _ = v44
            }
        } else {
            v45 := GenS7{F0: 42, F1: -1}
            _ = v45
        }
        _ = fi30
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 bool = (1000 >= -1)
        for fi48 := 0; fi48 < 10; fi48++ {
            fmt.Println(v8, v5, v5)
            switch 2147483647 {
            case 127:
                var v49 int = fi48
                _ = v49
            }
            _ = fi48
        }
        v50 := GenS8{F0: (^(v8))}
        _ = fi46
        _ = v47
        _ = v50
    }
    var v51 bool = (true && (false && true))
    v20 -= v20
    var v52 int8 = 1
    v3 = 0.1
    v51 = (((1 > -1) || (int16(v52) != 705)) || v51)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v8
}
