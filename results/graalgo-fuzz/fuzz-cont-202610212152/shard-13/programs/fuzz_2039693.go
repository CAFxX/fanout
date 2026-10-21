package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint16; F2 string }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int32; F1 int8; F2 uint16 }
func genFn0(p0 uint64, p1 uint32) int32 {
    p0 = (p0 + 916)
    for fi6 := 0; fi6 < 19; fi6++ {
        if false {
            v7 := GenS1{F0: 0}
            fmt.Println(fi6, p1, p0)
            var v8 uint16 = 2
            _ = v7
            _ = v8
        }
        var v9 float64 = ((3.14159 + -2.25) * (0.0 * float64(0.0)))
        _ = fi6
        _ = v9
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]int)
    var v3 bool = false
    var v4 int = 32767
    v5 := make(map[int]bool)
    v10 := []int{(v4 ^ 0), v4}
    var v11 int32 = -1
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v10 = append(v10, 1)
            _ = fi13
        }
        _ = fi12
    }
    if (!v3) {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 string = "hello"
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: 888}
        if (0 < 4294967295) {
            v17 := GenS2{F0: -9223372036854775808, F1: 562}
            _ = v17
        } else {
            var v18 string = ("\n\t" + "foo bar")
            _ = v18
        }
        _ = v16
    } else {
        v19 := GenS2{F0: 1, F1: -1}
        fmt.Println(v1, v11, v3)
        _ = v19
    }
    _ = len(v10)
    if v3 {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    if true {
        for fi20 := 0; fi20 < 4; fi20++ {
            if (42 > 522) {
                for fi21 := 0; fi21 < 6; fi21++ {
                    _ = len(v2)
                    switch "" {
                    case "foo bar":
                        var v22 int16 = -638
                        _ = v22
                    case "hello":
                        _ = len(v10)
                        if ((v1 % ((v1) | 1)) == 1000) {
                            v11 += v11
                        } else {
                            fmt.Println(fi20, v3)
                        }
                    default:
                        v4 = fi20
                    }
                    _ = len(v2)
                    _ = fi21
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v3 = v3
            } else {
                switch int(v11) {
                case -552:
                    v23 := GenS3{F0: 1}
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v3, v1, v4)
                default:
                    v11 = (v11 << ((255) & 63))
                }
                var v25 float32 = -2.25
                var v26 uint = 335
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi20
        }
        var v27 int8 = -1
        _ = v27
    } else {
        var v28 uint16 = 1
        v29 := GenS4{F0: 42, F1: (v28 * (^(v28))), F2: ("hello" + ("" + ""))}
        fmt.Println(v3, v4, v1)
        _ = v28
        _ = v29
    }
    v30 := GenS5{F0: v4, F1: 1}
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(v11, v3)
        if v3 {
            var v32 bool = v3
            v4 += (v4 * 128)
            switch (v4 ^ -2147483648) {
            case -2147483648:
                if v3 {
                    v10 = append(v10, v4)
                    var v33 uint8 = 195
                    var v34 uint64 = 2
                    _ = v33
                    _ = v34
                } else {
                    v1 = v1
                    fmt.Println(fi31, v3)
                }
            case 32768:
                switch (("\n\t" + "a") + "foo bar") {
                case "hello":
                    var v35 int64 = 447
                    _ = v35
                default:
                    _ = len(v2)
                }
            }
            _ = v32
        }
        if v3 {
            var v36 float64 = (0.1 * ((0.0 + 0.0) / 1.0))
            v37 := GenS6{F0: v11, F1: 0, F2: 32767}
            fmt.Println(v1, v1, v11)
            _ = v36
            _ = v37
        } else {
            var v38 int32 = v11
            _ = v38
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
