package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int32 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int32; F1 string; F2 int8 }

func main() {
    v1 := []uint32{546, 771}
    v2 := []string{"foo bar"}
    var v3 float32 = float32(3.14159)
    v4 := []int{}
    switch ("\n\t" + "\n\t") {
    case "":
        v3 = float32(-2.25)
    default:
        var v5 uint8 = 12
        _ = v5
    }
    var v6 int16 = 1
    fmt.Println(v3)
    var v7 int16 = (v6 ^ v6)
    v8 := GenS1{F0: -2147483648, F1: 65535}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v9 int64 = 0
    v9 -= v9
    if ("hello" <= "foo bar") {
        var v10 uint = 662
        _ = v10
    } else {
        fmt.Println(v6)
        if false {
            v6 = v6
            if false {
                var v11 uint = 956
                switch "\n\t" {
                case "foo bar":
                    v3 += v3
                case "\n\t":
                    if (!(1 != 867)) {
                        for fi12 := 0; fi12 < 1; fi12++ {
                            for fi13 := 0; fi13 < 20; fi13++ {
                                v14 := GenS1{F0: 95, F1: 42}
                                _ = fi13
                                _ = v14
                            }
                            _ = fi12
                        }
                        var v15 int32 = 2147483647
                        _ = v15
                    }
                case "a":
                    v3 = v3
                default:
                    _ = len(v1)
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    v3 = float32(0.1)
                    var v17 int8 = 42
                    v11 -= v11
                    _ = fi16
                    _ = v17
                }
                _ = v11
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                _ = fi18
            }
        } else {
            var v19 int16 = (v6 / (((-(v6))) | 1))
            var v20 int64 = -1
            _ = v19
            _ = v20
        }
        _ = len(v4)
    }
    v21 := GenS2{F0: 0, F1: 4294967295, F2: 42}
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS2{F0: 1, F1: 40, F2: 468}
        if false {
            if false {
                var v24 int8 = -61
                var v25 uint16 = 32767
                v24 = v24
                _ = v24
                _ = v25
            }
            var v26 string = ("a" + "\n\t")
            _ = len(v2)
            _ = v26
        } else {
            for fi27 := 0; fi27 < 8; fi27++ {
                switch 826 {
                case 4294967295:
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v28 uint32 = 42
                        _ = v28
                    } else {
                        v6 += v7
                        v29 := GenS3{F0: 42, F1: false}
                        _ = v29
                    }
                    v7 += v6
                case 32767:
                    var v30 uint32 = 65535
                    _ = v30
                case 32768:
                    v31 := GenS4{F0: 0, F1: "foo bar", F2: int8(v7)}
                    _ = v31
                }
                _ = fi27
            }
            var v32 uint16 = 1
            _ = v32
        }
        var v33 int16 = v6
        _ = fi22
        _ = v23
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 float32 = v3
        var v36 float32 = (0.0 + (0.0 + ((v35 / v35) + (-2.25 + (v3 * 0.0)))))
        _ = fi34
        _ = v35
        _ = v36
    }
    v7 -= 764
    if (80 < 183) {
        v6 = (v7 * v6)
    } else {
        if (!((v9 ^ v9) < -544)) {
            v6 -= (v7 / ((v6) | 1))
        } else {
            switch "foo bar" {
            case "foo bar":
                v9 = v9
            default:
                v9 -= (v9 ^ v9)
            }
        }
    }
    fmt.Println(v6, v6)
    v7 = (v7 * v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v9 += (v9 - (v9 / ((v9) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
