package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint) int8 {
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint64 = 795
        if (!(65535 != 0)) {
            v7 := GenS1{F0: 1, F1: (v6 - 302)}
            var v8 uint = 255
            switch (("" + "foo bar") + "") {
            case "\n\t":
                switch "" {
                case "hello":
                    var v9 uint16 = 75
                    _ = v9
                }
            }
            _ = v7
            _ = v8
        } else {
            v6 += 0
        }
        var v10 int64 = 912
        _ = fi5
        _ = v10
        _ = v6
    }
    switch (("" + (("hello" + ("" + (("" + ("\n\t" + (((("foo bar" + "hello") + "foo bar") + "\n\t") + "\n\t"))) + "foo bar"))) + "hello")) + "") {
    case "a":
        p0 = 955
    }
    var v11 uint16 = 1
    var v12 int8 = 100
    v11 -= (^(v11))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 100
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (((((0.1 / ((0.1 - 0.0) + (1.5 * 3.14159))) / 1.0) - 0.0) - 1.5) + 0.1)
    v3 := make(map[int]string)
    var v4 int8 = 0
    v13 := GenS2{F0: 0, F1: 185}
    v2 = (v2 * ((v2 / v2) * ((v2 / v2) - v2)))
    for fi14 := 0; fi14 < 9; fi14++ {
        if (!false) {
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS1{F0: -1, F1: 4294967295}
                delete(v1, 2147483647)
                v17 := GenS1{F0: 239, F1: 27}
                _ = fi15
                _ = v16
                _ = v17
            }
            fmt.Println(fi14, v4, v2)
            fmt.Println(fi14, fi14)
        }
        _ = fi14
    }
    var v18 int32 = 771
    switch 1 {
    case 42:
        if (true || false) {
            fmt.Println(v4, v2)
            switch "" {
            case "hello":
                var v19 int32 = (v18 % ((v18) | 1))
                _ = v19
            case "foo bar":
                var v20 int8 = 42
                _ = v20
            case "a":
                switch ("hello" + "foo bar") {
                case "hello":
                    v21 := GenS3{F0: v4}
                    _ = v21
                case "":
                    if true {
                        v2 = v2
                    } else {
                        var v22 int8 = v4
                        v18 -= v18
                        _ = v22
                    }
                }
            default:
                v3[-608] = ("foo bar" + ("\n\t" + "foo bar"))
            }
            if false {
                fmt.Println(v4, v4)
            }
        } else {
            v23 := GenS1{F0: -1, F1: 42}
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 float32 = 1.5
                var v26 bool = false
                _ = fi24
                _ = v25
                _ = v26
            }
            v4 -= v4
            _ = v23
        }
    case -829:
        fmt.Println(v4, v2)
    case -2147483648:
        fmt.Println(v18, v18, v4)
    default:
        v4 -= int8(v18)
    }
    switch (("" + "") + "hello") {
    case "a":
        var v27 uint8 = 2
        _ = v27
    case "\n\t":
        if (774 >= -1) {
            for fi28 := 0; fi28 < 15; fi28++ {
                fmt.Println(v18, v18, v18)
                var v29 int16 = 0
                var v30 uint8 = 0
                _ = fi28
                _ = v29
                _ = v30
            }
            if true {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := []int{(fi31 / (((fi31 >> ((2) & 63))) | 1))}
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v18 += 0
                        v4 = v4
                        _ = len(v3)
                        v34 := GenS1{F0: 0, F1: 2}
                        _ = fi33
                        _ = v34
                    }
                    fmt.Println(v2, v4, v4)
                    _ = fi31
                    _ = v32
                }
                for fi35 := 0; fi35 < 3; fi35++ {
                    fmt.Println(v4, fi35, fi35)
                    _ = fi35
                }
                v4 -= (v4 + (v4 * v4))
            } else {
                v36 := GenS1{F0: 42, F1: 4294967295}
                v37 := GenS4{F0: 0}
                _ = v36
                _ = v37
            }
            fmt.Println(v2)
        } else {
            for fi38 := 0; fi38 < 11; fi38++ {
                v4 += -54
                for fi39 := 0; fi39 < 10; fi39++ {
                    var v40 uint = 1
                    _ = fi39
                    _ = v40
                }
                fmt.Println(v2, v18, v4)
                _ = fi38
            }
            var v41 bool = false
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 uint = 42
                var v44 uint64 = 731
                fmt.Println(v44)
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = v41
        }
    }
    v45 := GenS3{F0: (v4 & genFn0(810, 2))}
    fmt.Println(v18, v18)
    v46 := GenS5{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
