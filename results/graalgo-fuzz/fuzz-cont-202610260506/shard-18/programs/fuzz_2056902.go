package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int8; F1 int; F2 uint8 }
type GenS4 struct { F0 int; F1 string; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint32; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 uint32; F2 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 int; F2 int64 }
func genFn0(p0 int, p1 bool, p2 string) uint8 {
    switch p2 {
    case "a":
        p1 = (-32768 != 32767)
    default:
        fmt.Println(p0, p2, p2)
    }
    v4 := []int{p0, p0, (p0 ^ p0), (p0 / ((-2147483648) | 1))}
    for fi5 := 0; fi5 < 17; fi5++ {
        switch ((("\n\t" + "hello") + p2) + (p2 + "a")) {
        case "\n\t":
            for fi6 := 0; fi6 < 4; fi6++ {
                v7 := GenS1{F0: 42, F1: 243}
                _ = fi6
                _ = v7
            }
        case "":
            for fi8 := 0; fi8 < 8; fi8++ {
                if p1 {
                    switch "foo bar" {
                    case "foo bar":
                        if p1 {
                            var v9 bool = p1
                            var v10 float64 = float64(0.1)
                            v11 := GenS1{F0: 1000, F1: 196}
                            _ = v10
                            _ = v11
                            _ = v9
                        } else {
                            var v12 int16 = -32768
                            var v13 int8 = 127
                            _ = v12
                            _ = v13
                        }
                    default:
                        if p1 {
                            fmt.Println(fi5, fi8, fi8)
                            fmt.Println(fi5, p2)
                            v14 := GenS2{F0: 42, F1: 100}
                            _ = v14
                        } else {
                            var v15 uint = 758
                            p0 = fi8
                            v16 := GenS3{F0: -17, F1: 1, F2: 255}
                            _ = v15
                            _ = v16
                        }
                    }
                    p1 = (1 < 0)
                    if true {
                        if p1 {
                            fmt.Println(p2, p2)
                            switch p2 {
                            case "a":
                                fmt.Println(p2, fi5, p1)
                            case "foo bar":
                                p0 = (p0 >> ((42) & 63))
                            default:
                                if p1 {
                                    _ = len(v4)
                                    var v17 int64 = 1
                                    var v18 float64 = float64((0.1 / ((1e10 * 0.1) / (((float32((((((1e10 + 3.14159) + ((3.14159 * 0.0) * 0.1)) * float64(((1e10 * (0.1 * 0.1)) * 1e10))) + 0.0) - float64(((-2.25 * (((float32(0.1) / 1.0) + float32(1e10)) - 1.5)) + -2.25)))) + 0.1) * (0.1 / 1e10)) + 3.14159))))
                                    _ = v17
                                    _ = v18
                                }
                            }
                            v19 := GenS4{F0: -2147483648, F1: ("a" + ((p2 + ((p2 + p2) + ((p2 + "\n\t") + p2))) + p2)), F2: 50}
                            _ = v19
                        }
                        v4 = append(v4, p0)
                        _ = len(v4)
                    }
                } else {
                    fmt.Println(p2)
                    if (p1 || (!p1)) {
                        if (!p1) {
                            fmt.Println(p1, fi8)
                        } else {
                            var v20 uint8 = 1
                            v21 := GenS5{F0: 4294967295, F1: 356, F2: p0}
                            v22 := GenS2{F0: 2147483647, F1: 100}
                            var v23 bool = p1
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                    var v24 uint = 42
                    fmt.Println(fi5, fi8, fi8)
                    v25 := GenS6{F0: -285}
                    _ = v24
                    _ = v25
                }
                _ = fi8
            }
        }
        var v26 float32 = 1e10
        _ = fi5
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 int64 = 475
    var v2 bool = false
    var v3 uint = 760
    switch 42 {
    case -9223372036854775808:
        v27 := GenS7{F0: -1, F1: 290, F2: (1 > -2147483648)}
        _ = v27
    case 42:
        v2 = false
    case 127:
        v2 = true
    }
    v3 = v3
    var v28 int64 = v1
    switch "\n\t" {
    case "":
        var v29 int16 = 32767
        _ = v29
    default:
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 bool = (!(!false))
            for fi32 := 0; fi32 < 9; fi32++ {
                v2 = (-2.25 <= 1e10)
                var v33 int32 = -621
                fmt.Println(fi32, v31)
                if v31 {
                    v34 := []int{fi32}
                    var v35 int64 = v28
                    _ = v34
                    _ = v35
                }
                _ = fi32
                _ = v33
            }
            _ = fi30
            _ = v31
        }
    }
    v36 := GenS8{F0: 196}
    var v37 float32 = float32(1.5)
    v38 := GenS9{F0: -2147483648, F1: -201, F2: v1}
    for fi39 := 0; fi39 < 18; fi39++ {
        v28 = v28
        v28 = v28
        _ = fi39
    }
    var v40 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v40
}
