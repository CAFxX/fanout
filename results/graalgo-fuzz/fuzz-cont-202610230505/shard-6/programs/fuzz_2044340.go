package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint }
func genFn0() uint {
    if (48 >= 42) {
        var v7 int16 = 217
        _ = v7
    } else {
        var v8 uint16 = 602
        var v9 uint8 = 254
        _ = v8
        _ = v9
    }
    v10 := []int{1, 862}
    switch 32767 {
    case 32767:
        v10 = append(v10, 32767)
    case -265:
        var v11 uint8 = 128
        _ = v11
    default:
        var v12 uint16 = 65535
        _ = v12
    }
    if (784 >= 2) {
        switch -629 {
        case 128:
            var v13 int = 128
            _ = v13
        case 4294967295:
            if false {
                var v14 uint64 = 2
                var v15 uint64 = (v14 - 389)
                v15 = v14
                _ = v14
                _ = v15
            } else {
                _ = len(v10)
                var v16 int32 = -612
                _ = v16
            }
        }
        switch -2147483648 {
        case 128:
            var v17 int16 = 0
            _ = v17
        case -2147483648:
            switch ("\n\t" + "\n\t") {
            case "":
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 uint64 = 351
                    if (128 > 255) {
                        if (!(!(v19 > (v19 >> ((2) & 63))))) {
                            if false {
                                v19 -= (^(v19))
                            }
                        } else {
                            v20 := GenS1{F0: 42}
                            var v21 float32 = 0.0
                            _ = v20
                            _ = v21
                        }
                    } else {
                        switch "" {
                        case "\n\t":
                            for fi22 := 0; fi22 < 11; fi22++ {
                                v23 := GenS1{F0: -689}
                                var v24 uint32 = uint32(v19)
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        default:
                            _ = len(v10)
                        }
                        var v25 bool = false
                        _ = v25
                    }
                    _ = fi18
                    _ = v19
                }
            case "a":
                var v26 string = "\n\t"
                _ = v26
            }
        case 32768:
            var v27 uint64 = 42
            _ = v27
        default:
            _ = len(v10)
        }
    } else {
        var v28 int = -9223372036854775808
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v28
    }
    var v29 int64 = 1
    _ = v10
    _ = v29
    return 0
}
func genFn1(p0 int32, p1 string, p2 int64) string {
    v30 := GenS1{F0: 0}
    switch p1 {
    case "foo bar":
        v31 := []int{249, 128, 42, 1}
        _ = v31
    case "a":
        if false {
            v32 := []int{4294967295, 0, 4294967295, -1}
            _ = v32
        }
    case "":
        p2 += (p2 % ((0) | 1))
    default:
        for fi33 := 0; fi33 < 9; fi33++ {
            var v34 int16 = -1
            p2 = (p2 ^ 1)
            _ = fi33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        p0 = (p0 << ((378) & 63))
        var v36 float64 = ((1.5 - (float64(1e10) - 0.0)) / (((1e10 + (0.0 + (1.5 - (1.5 / float64(3.14159))))) / -2.25) + 1e10))
        _ = fi35
        _ = v36
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return ((p1 + ((p1 + "a") + (p1 + p1))) + (("hello" + "") + (p1 + p1)))
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]bool)
    var v3 uint = 65535
    v4 := make(map[int]bool)
    var v5 uint = v3
    v6 := make(map[uint32]int)
    v37 := GenS1{F0: 1}
    var v38 int64 = 1
    v38 = v38
    if false {
        v39 := []int{32768, 401, -252}
        _ = v39
    }
    if false {
        v38 = v38
    }
    var v40 int = 813
    v1 = (v1 | v1)
    switch "\n\t" {
    case "foo bar":
        v41 := GenS2{F0: "a", F1: v5}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
