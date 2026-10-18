package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 uint8, p1 uint, p2 uint16) int8 {
    p2 -= p2
    v4 := GenS1{F0: 0}
    var v5 int = -597
    for fi6 := 0; fi6 < 15; fi6++ {
        switch "" {
        case "foo bar":
            v7 := []int{}
            _ = v7
        case "":
            v8 := []int{fi6, v5, (fi6 / ((v5) | 1))}
            _ = v8
        case "\n\t":
            if false {
                p0 -= (p0 - p0)
            } else {
                v9 := []int{(v5 - fi6), 0, (fi6 ^ (fi6 >> ((p1) & 63)))}
                var v10 uint = uint(p2)
                _ = len(v9)
                _ = v10
                _ = v9
            }
        default:
            p1 = p1
        }
        if true {
            fmt.Println(p1, p1)
            v11 := []int{(v5 | v5)}
            p1 += (p1 - p1)
            _ = v11
        }
        var v12 string = "\n\t"
        _ = fi6
        _ = v12
    }
    var v13 uint = (p1 >> (((^(p1))) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v4
    _ = v5
    return -7
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 uint16 = 32767
    var v3 uint16 = v2
    for fi14 := 0; fi14 < 8; fi14++ {
        v2 = v3
        for fi15 := 0; fi15 < 1; fi15++ {
            switch ("\n\t" + "") {
            case "foo bar":
                v16 := GenS2{F0: 65535, F1: -57}
                _ = v16
            case "hello":
                v3 += (v2 >> ((65535) & 63))
            case "\n\t":
                if (false && true) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v1 = append(v1, (^(fi14)))
                    switch "" {
                    case "a":
                        switch "foo bar" {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = fi14
                            }
                        case "\n\t":
                            v3 += v2
                        }
                    case "hello":
                        if false {
                            var v17 uint64 = 578
                            var v18 uint = 987
                            v19 := GenS1{F0: 161}
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    case "foo bar":
                        var v20 uint8 = 2
                        _ = v20
                    }
                } else {
                    var v21 int = -326
                    v22 := GenS2{F0: 4294967295, F1: -1}
                    _ = v21
                    _ = v22
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if ((!false) && true) {
                var v23 int64 = -9223372036854775808
                _ = v23
            } else {
                var v24 float32 = 1.5
                var v25 uint32 = 2147483648
                _ = v24
                _ = v25
            }
            _ = fi15
        }
        _ = fi14
    }
    v26 := GenS1{F0: 483}
    var v27 uint = 717
    v2 = v3
    if (2 == 128) {
        fmt.Println(v3, v3, v27)
        for fi28 := 0; fi28 < 3; fi28++ {
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 float64 = 1e10
                switch "hello" {
                case "\n\t":
                    v30 = (v30 + (1e10 + (((3.14159 * -2.25) / (-2.25 - (v30 * v30))) / v30)))
                default:
                    var v31 bool = (!false)
                    _ = v31
                }
                v27 = (v27 << (((v27 + v27)) & 63))
                _ = fi29
                _ = v30
            }
            if true {
                v3 = (v3 % ((v3) | 1))
                var v32 int = (fi28 >> ((346) & 63))
                v33 := GenS2{F0: 2, F1: 1}
                _ = v32
                _ = v33
            } else {
                var v34 uint64 = 42
                _ = v34
            }
            v27 += (^(v27))
            _ = fi28
        }
        v35 := GenS2{F0: (v27 >> (((v27 - v27)) & 63)), F1: 42}
        _ = v35
    }
    v27 = v27
    fmt.Println(v27)
    var v36 int64 = 9223372036854775807
    switch 0 {
    case -182:
        var v37 int64 = v36
        _ = v37
    }
    v2 = v2
    var v38 uint64 = 4294967295
    for fi39 := 0; fi39 < 18; fi39++ {
        v27 = (v27 | 1)
        var v40 int16 = 42
        _ = fi39
        _ = v40
    }
    fmt.Println(v36)
    var v41 uint8 = uint8(v3)
    fmt.Println(v41, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v38
    _ = v41
}
