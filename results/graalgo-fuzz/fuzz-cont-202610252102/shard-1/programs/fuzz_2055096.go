package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"a", ("\n\t" + "a")}
    v3 := make(map[int]bool)
    var v4 float32 = 1.5
    v5 := []int{-995}
    v6 := []string{"foo bar", "foo bar", "\n\t"}
    switch "foo bar" {
    case "":
        v1[("\n\t" + ("hello" + "a"))] = false
    default:
        if false {
            var v7 bool = true
            _ = v7
        }
    }
    if true {
        if (9223372036854775807 >= -243) {
            var v8 uint64 = 2
            _ = v8
        }
        switch "\n\t" {
        case "\n\t":
            var v9 int = -2147483648
            _ = v9
        case "foo bar":
            fmt.Println(v4)
        case "a":
            v4 -= v4
        default:
            _ = len(v6)
        }
        v10 := GenS1{F0: 414, F1: "\n\t"}
        _ = v10
    } else {
        if ((1 < -9223372036854775808) || (-26 > 42)) {
            if true {
                if true {
                    fmt.Println(v4)
                    v4 += 0.0
                    v11 := GenS1{F0: 1, F1: "foo bar"}
                    _ = v11
                }
                v12 := GenS1{F0: 1, F1: "\n\t"}
                v13 := GenS2{F0: 736}
                _ = v12
                _ = v13
            } else {
                var v14 int8 = 1
                v14 -= v14
                if len(v6) > 0 {
                    v6[0] = "hello"
                }
                _ = v14
            }
            var v15 bool = false
            _ = v15
        } else {
            var v16 float64 = 0.0
            if true {
                if true {
                    var v17 int64 = -373
                    _ = v17
                }
                switch "foo bar" {
                case "a":
                    for fi18 := 0; fi18 < 20; fi18++ {
                        switch (fi18 << ((1) & 63)) {
                        case 4294967295:
                            var v19 bool = true
                            _ = v19
                        case -1:
                            var v20 bool = false
                            _ = v20
                        }
                        var v21 uint16 = 65535
                        var v22 int = -547
                        _ = fi18
                        _ = v21
                        _ = v22
                    }
                case "foo bar":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    v16 -= (v16 * v16)
                }
            }
            var v23 float32 = float32((v16 / v16))
            _ = v16
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v4 -= (v4 + (1.5 / (0.1 + v4)))
        _ = fi24
    }
    var v25 float64 = -2.25
    v25 -= ((((((float64(v4) + ((v25 * float64(1.5)) - v25)) + (0.0 / (float64(3.14159) + v25))) * 0.0) / v25) + float64((-2.25 / (-2.25 - float32(1e10))))) - (3.14159 + (0.1 + v25)))
    for fi26 := 0; fi26 < 9; fi26++ {
        v25 += v25
        _ = fi26
    }
    if (255 == 2) {
        var v27 uint = 0
        v25 -= float64(float32((v25 / v25)))
        v25 = (0.1 * (0.0 - v25))
        _ = v27
    }
    switch 465 {
    case 447:
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 int64 = -9223372036854775808
            _ = fi28
            _ = v29
        }
    case 84:
        var v30 int8 = -1
        _ = v30
    case 148:
        if true {
            if ((!(true && true)) && (float64(3.14159) != (v25 * (-2.25 + v25)))) {
                v31 := GenS3{F0: "", F1: -1}
                var v32 float32 = ((((v4 / (v4 - (-2.25 - (((1.5 - float32(v25)) + v4) * 0.0)))) + float32((float64(v4) / v25))) * (v4 * v4)) / (((1.5 * float32(((0.0 / (((v25 + v25) * 1e10) - v25)) - v25))) / (v4 + v4)) * v4))
                _ = v31
                _ = v32
            }
            var v33 uint16 = 1
            fmt.Println(v33, v33)
            _ = v33
        } else {
            v4 += 1.5
            for fi34 := 0; fi34 < 19; fi34++ {
                if (true && false) {
                    var v35 string = "foo bar"
                    var v36 uint16 = 909
                    for fi37 := 0; fi37 < 6; fi37++ {
                        fmt.Println(v4, v36)
                        var v38 uint = 65535
                        switch (fi34 >> (((v38 * v38)) & 63)) {
                        case 4294967295:
                            v5 = append(v5, (fi34 % ((fi37) | 1)))
                        }
                        _ = fi37
                        _ = v38
                    }
                    _ = v35
                    _ = v36
                } else {
                    v25 += v25
                    var v39 uint8 = 2
                    var v40 int64 = -28
                    _ = v39
                    _ = v40
                }
                switch "\n\t" {
                case "hello":
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v4 -= (-2.25 / (1e10 + v4))
                        v4 += v4
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi41
                    }
                case "":
                    v42 := GenS3{F0: "hello", F1: -1}
                    _ = v42
                default:
                    var v43 bool = true
                    _ = v43
                }
                _ = fi34
            }
            v25 = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
