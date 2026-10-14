package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint32 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 int }
func genFn0() int64 {
    v7 := GenS1{F0: 42, F1: 1, F2: 42}
    var v8 int16 = 721
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32) string {
    var v9 bool = false
    v10 := []int{127, -788, -240, 0}
    v11 := GenS1{F0: 834, F1: 42, F2: p0}
    if true {
        p0 += (p0 * 275)
        if v9 {
            switch "a" {
            case "":
                v12 := GenS2{F0: 42, F1: 42, F2: 352}
                _ = v12
            case "foo bar":
                v9 = (float64(float32(((0.1 - (0.1 + (1e10 * 0.0))) + (0.1 * (float64(-2.25) + ((1e10 * 0.0) * 1.5)))))) < (0.1 - -2.25))
            default:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            v9 = (((("" + (("a" + ("hello" + "foo bar")) + "hello")) + "") + "hello") <= "a")
            p0 += (p0 - (p0 | 2))
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            v9 = (32767 == 1000)
            _ = fi13
        }
    }
    var v14 float32 = 1.5
    _ = p0
    _ = v10
    _ = v11
    _ = v14
    _ = v9
    return ("\n\t" + "")
}

func main() {
    var v1 int64 = -190
    v2 := []int{1}
    var v3 float32 = 3.14159
    var v4 int8 = 42
    var v5 uint = 65535
    var v6 float64 = (0.1 * ((0.1 + 0.1) - ((1e10 / (1.5 * 1.5)) + (-2.25 - ((1e10 - 1e10) / 3.14159)))))
    var v15 float32 = ((0.1 / (0.1 - v3)) - -2.25)
    var v16 int8 = (v4 + v4)
    v17 := GenS2{F0: -9223372036854775808, F1: 1, F2: 831}
    var v18 float64 = 1.5
    switch 0 {
    case -721:
        var v19 float32 = v3
        _ = v19
    default:
        var v20 int8 = 26
        _ = v20
    }
    if true {
        v18 = v18
        var v21 string = "foo bar"
        v22 := GenS3{F0: 42}
        _ = v21
        _ = v22
    }
    fmt.Println(v1, v4)
    if false {
        v5 = (v5 / ((v5) | 1))
        if true {
            _ = len(v2)
        } else {
            fmt.Println(v6)
            if (float32(float64(0.0)) <= 0.0) {
                var v23 int64 = v1
                v2 = append(v2, int(v4))
                v18 += ((1e10 * 0.0) + (v18 + ((3.14159 * v6) / -2.25)))
                _ = v23
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 int = -1
                    v5 = v5
                    switch "\n\t" {
                    case "foo bar":
                        var v26 bool = true
                        _ = v26
                    case "a":
                        for fi27 := 0; fi27 < 9; fi27++ {
                            v15 += v3
                            var v28 uint = v5
                            v2 = append(v2, -1)
                            _ = fi27
                            _ = v28
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                v16 = v4
                var v29 bool = false
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 1; fi30++ {
            v3 = 0.1
            _ = fi30
        }
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 3; fi32++ {
            if len(v2) > 0 {
                v2[0] = fi31
            }
            _ = fi32
        }
        v18 += (0.0 * (v18 - (v18 / float64(v3))))
        _ = fi31
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    for fi33 := 0; fi33 < 13; fi33++ {
        if (!true) {
            var v34 uint32 = 65535
            fmt.Println(v4, v16, v3)
            if true {
                fmt.Println(v3, v1)
                var v35 int64 = v1
                _ = v35
            }
            _ = v34
        } else {
            var v36 uint8 = 255
            switch "a" {
            case "\n\t":
                var v37 uint = 962
                _ = v37
            case "hello":
                switch "a" {
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        v38 := GenS2{F0: 127, F1: 42, F2: 1}
                        _ = v38
                    case "a":
                        v2 = append(v2, fi33)
                    }
                }
            default:
                for fi39 := 0; fi39 < 4; fi39++ {
                    if true {
                        if true {
                            var v40 int32 = 574
                            _ = v40
                        } else {
                            fmt.Println(v36, fi33, v18)
                        }
                        v18 += -2.25
                        var v41 uint16 = 1
                        fmt.Println(v16)
                        var v42 uint8 = (v36 & (v36 ^ (^(v36))))
                        var v43 int = fi39
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    fmt.Println(fi33, fi33)
                    _ = fi39
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v36
        }
        var v44 uint16 = 32767
        _ = fi33
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
