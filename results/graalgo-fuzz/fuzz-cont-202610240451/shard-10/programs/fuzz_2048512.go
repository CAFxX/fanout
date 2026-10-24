package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v4 int = 1
    v4 = (v4 | (v4 % ((513) | 1)))
    v1 -= (float64(((3.14159 + (-2.25 - 1e10)) - 0.1)) - 0.0)
    if true {
        v5 := []int{}
        var v6 uint32 = 42
        _ = v5
        _ = v6
    } else {
        fmt.Println(v4, v4)
        var v7 uint64 = 2
        _ = v7
    }
    v4 = v4
    var v8 bool = false
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v4)
        switch 32768 {
        case 752:
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 int64 = 9223372036854775807
                var v12 bool = (-498 != 42)
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            fmt.Println(v8)
        }
        var v13 int32 = -2147483648
        _ = fi9
        _ = v13
    }
    if (32767 >= v4) {
        switch "a" {
        case "hello":
            v4 -= (^(v4))
        case "a":
            if v8 {
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v8 = false
                        var v16 uint16 = 282
                        _ = fi15
                        _ = v16
                    }
                    v8 = (4294967295 < 42)
                    v17 := GenS1{F0: 255}
                    _ = fi14
                    _ = v17
                }
                var v18 string = ("a" + ("foo bar" + ""))
                _ = v18
            }
        }
        _ = len(v2)
    }
    v19 := GenS1{F0: 32767}
    if false {
        v1 -= v1
    } else {
        if false {
            if ((120 == 2) || true) {
                v20 := v3["foo bar"]
                _ = v20
            }
            v4 = v4
        } else {
            var v21 int32 = -2147483648
            v1 -= 3.14159
            _ = v21
        }
        var v22 uint = 42
        _ = v22
    }
    delete(v3, "foo bar")
    v1 -= v1
    _ = len(v2)
    switch "" {
    case "":
        v23 := GenS1{F0: 0}
        _ = v23
    case "foo bar":
        for fi24 := 0; fi24 < 16; fi24++ {
            if ((1 > 65535) || v8) {
                if v8 {
                    v4 = -695
                } else {
                    switch -847 {
                    case 0:
                        v25 := GenS2{F0: ((1.5 / -2.25) != v1), F1: -32768, F2: 2}
                        _ = v25
                    case 2147483647:
                        v26 := []int{v4, v4, 2147483647}
                        _ = v26
                    case 9223372036854775807:
                        v27 := []int{fi24}
                        _ = v27
                    }
                    var v28 float64 = (0.1 - float64((float32(v1) + 0.1)))
                    _ = v28
                }
            } else {
                var v29 int64 = -1
                _ = v29
            }
            v1 = v1
            var v30 int8 = 102
            _ = fi24
            _ = v30
        }
    default:
        v31 := GenS3{F0: 2, F1: -2147483648}
        _ = v31
    }
    if (0 > 1) {
        switch "\n\t" {
        case "a":
            if (15 < 905) {
                delete(v3, "a")
                switch (("hello" + "") + "a") {
                case "":
                    for fi32 := 0; fi32 < 19; fi32++ {
                        switch "" {
                        case "hello":
                            v33 := GenS4{F0: 981}
                            _ = v33
                        default:
                            v34 := GenS3{F0: 466, F1: fi32}
                            _ = v34
                        }
                        _ = fi32
                    }
                case "hello":
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v1 = v1
                        if (1000 <= 918) {
                            v36 := GenS3{F0: 18446744073709551615, F1: (fi35 * 2147483647)}
                            var v37 uint64 = 42
                            _ = v36
                            _ = v37
                        } else {
                            v8 = v8
                            v38 := GenS5{F0: -260}
                            _ = v38
                        }
                        _ = fi35
                    }
                }
            } else {
                var v39 uint64 = 0
                _ = v39
            }
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 uint32 = 0
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 14; fi42++ {
            v1 = (float64((0.0 * ((3.14159 / 1.5) / (-2.25 - 1.5)))) - (float64(0.0) / (float64(float32(v1)) / float64(-2.25))))
            v43 := GenS5{F0: 0}
            v8 = (0 >= 32767)
            _ = fi42
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
