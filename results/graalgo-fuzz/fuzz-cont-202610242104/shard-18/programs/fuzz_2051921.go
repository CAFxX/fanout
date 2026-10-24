package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int8; F2 int16 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint16) string {
    v5 := GenS1{F0: 42}
    var v6 string = "foo bar"
    if true {
        if ((9223372036854775807 != 42) && (!((1 > -1) && false))) {
            if false {
                switch v6 {
                case "foo bar":
                    if true {
                        fmt.Println(p0)
                    } else {
                        for fi7 := 0; fi7 < 11; fi7++ {
                            v6 = v6
                            v8 := []int{(fi7 ^ (fi7 >> ((42) & 63))), 1, 32768}
                            v6 = v6
                            _ = fi7
                            _ = v8
                        }
                        v9 := GenS1{F0: 128}
                        switch "foo bar" {
                        case "":
                            v6 = "\n\t"
                        default:
                            v10 := GenS1{F0: 0}
                            _ = v10
                        }
                        _ = v9
                    }
                case "hello":
                    var v11 uint32 = 457
                    _ = v11
                default:
                    v12 := GenS2{F0: v6}
                    _ = v12
                }
                if true {
                    v13 := []int{9223372036854775807, -2147483648}
                    if len(v13) > 0 {
                        v13[0] = 9223372036854775807
                    }
                    if false {
                        v13 = append(v13, -9223372036854775808)
                    }
                    _ = v13
                }
            } else {
                v14 := GenS3{F0: -2147483648, F1: 0, F2: 32767}
                fmt.Println(v6, v6)
                if (1 != 898) {
                    var v15 int8 = 0
                    _ = v15
                }
                _ = v14
            }
        }
        v16 := []int{}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := []int{}
            p0 = (^(p0))
            switch fi17 {
            case 996:
                p0 -= (p0 * p0)
            }
            _ = fi17
            _ = v18
        }
        var v19 uint8 = 128
        var v20 int = 1
        _ = v19
        _ = v20
    }
    _ = p0
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]bool)
    var v3 uint32 = 42
    var v4 int64 = -9223372036854775808
    v21 := GenS3{F0: 2147483647, F1: 42, F2: 0}
    v1 = uint8(v3)
    var v22 float64 = 1.5
    v23 := GenS4{F0: v3, F1: 260}
    if (4294967295 == -2147483648) {
        if false {
            if true {
                v24 := GenS5{F0: false}
                var v25 float32 = 0.1
                _ = v24
                _ = v25
            }
        }
    }
    switch 42 {
    case 2147483647:
        if false {
            if false {
                v26 := GenS6{F0: -2147483648, F1: 1}
                _ = v26
            } else {
                v1 = v1
            }
            var v27 string = ("\n\t" + "")
            var v28 uint = 989
            _ = v27
            _ = v28
        }
    }
    v3 -= 443
    if ((2 <= 2) || true) {
        switch "\n\t" {
        case "a":
            var v29 uint32 = 65535
            _ = v29
        case "":
            if false {
                for fi30 := 0; fi30 < 9; fi30++ {
                    switch 42 {
                    case 9223372036854775807:
                        if (!(803 == -32768)) {
                            switch "a" {
                            case "":
                                v3 += v3
                            }
                            fmt.Println(v4)
                            if (!((v3 << ((1) & 63)) < v3)) {
                                var v31 float32 = float32(float64(0.0))
                                v3 += (^(v3))
                                _ = v31
                            }
                        } else {
                            v3 = (v3 >> ((2) & 63))
                            v32 := []int{(^(fi30)), -826, fi30, (fi30 ^ fi30)}
                            var v33 int = fi30
                            _ = v32
                            _ = v33
                        }
                    case 32767:
                        fmt.Println(v4)
                    }
                    fmt.Println(v1, fi30, v22)
                    _ = fi30
                }
                v34 := []int{-9223372036854775808}
                _ = v34
            }
        }
        var v35 uint8 = 127
        _ = v35
    } else {
        var v36 uint16 = 255
        _ = len(v2)
        var v37 int = 882
        _ = v36
        _ = v37
    }
    v38 := []int{-1, -2147483648, 32767, 32768}
    _ = len(v38)
    var v39 int32 = -2147483648
    var v40 uint8 = (v1 - v1)
    var v41 int = 127
    var v42 int64 = (v4 & int64(v3))
    v42 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
