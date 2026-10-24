package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int64; F2 uint }

func main() {
    v1 := []int{}
    var v2 bool = (1 < -2147483648)
    var v3 bool = v2
    var v4 int32 = 2147483647
    var v5 float32 = float32(float64(-2.25))
    v6 := GenS1{F0: 42, F1: 0, F2: 32768}
    switch 4294967295 {
    case -856:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch "a" {
    case "\n\t":
        if (((v5 * (v5 + v5)) - 1.5) > v5) {
            v7 := GenS2{F0: 0}
            _ = len(v1)
            if true {
                var v8 uint16 = 65535
                v2 = v2
                var v9 float32 = ((float32(0.0) * (1e10 * v5)) * (1e10 + (0.1 / 0.1)))
                _ = v8
                _ = v9
            } else {
                switch 32767 {
                case -504:
                    _ = len(v1)
                case 4294967295:
                    var v10 float32 = (((v5 + (v5 + v5)) - float32((1e10 + 0.1))) + 0.1)
                    _ = v10
                case 42:
                    v4 = 42
                }
                v4 = v4
                v11 := GenS2{F0: 4294967295}
                _ = v11
            }
            _ = v7
        } else {
            switch int(v4) {
            case 32767:
                fmt.Println(v5, v4, v4)
            case 496:
                if (127 < 42) {
                    if len(v1) > 0 {
                        v1[0] = -9223372036854775808
                    }
                    v4 += v4
                    switch 128 {
                    case -533:
                        v4 = v4
                    case 1:
                        var v12 int64 = -36
                        _ = v12
                    }
                } else {
                    v4 = 0
                    if v2 {
                        for fi13 := 0; fi13 < 18; fi13++ {
                            var v14 int = fi13
                            _ = fi13
                            _ = v14
                        }
                        v15 := GenS3{F0: ("hello" + "foo bar"), F1: 0, F2: 65535}
                        _ = v15
                    }
                }
            case 42:
                v16 := GenS1{F0: 60, F1: 97, F2: 32767}
                _ = v16
            default:
                var v17 int32 = v4
                _ = v17
            }
            v2 = false
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 int64 = -673
                fmt.Println(v3, v2)
                fmt.Println(v19)
                _ = fi18
                _ = v19
            }
        }
    case "foo bar":
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 float32 = v5
            switch "foo bar" {
            case "a":
                for fi22 := 0; fi22 < 10; fi22++ {
                    v5 = 3.14159
                    if false {
                        if true {
                            v23 := GenS1{F0: 4294967295, F1: 562, F2: fi22}
                            if (fi20 > (fi20 >> ((42) & 63))) {
                                switch "foo bar" {
                                case "a":
                                    v1 = append(v1, fi20)
                                default:
                                    v2 = v2
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v4 -= (^(v4))
                                v4 += v4
                                var v24 int64 = 411
                                _ = v24
                            }
                            _ = v23
                        }
                        fmt.Println(fi22, fi22)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi22
                }
            case "foo bar":
                v4 += (v4 >> ((42) & 63))
            case "\n\t":
                if (v4 <= v4) {
                    v5 += 0.1
                    v25 := GenS3{F0: "\n\t", F1: 9223372036854775807, F2: 255}
                    _ = v25
                }
            default:
                v26 := GenS1{F0: 42, F1: (fi20 >> ((42) & 63)), F2: fi20}
                _ = v26
            }
            fmt.Println(fi20, v21, fi20)
            _ = fi20
            _ = v21
        }
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    v5 -= v5
    switch -771 {
    case 32767:
        for fi27 := 0; fi27 < 11; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                fmt.Println(v2)
                var v29 uint64 = 614
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
    case 4294967295:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v30 uint32 = 276
        _ = v30
    }
    fmt.Println(v4, v5, v3)
    var v31 uint32 = 2
    v32 := GenS2{F0: 2}
    var v33 string = ("a" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
