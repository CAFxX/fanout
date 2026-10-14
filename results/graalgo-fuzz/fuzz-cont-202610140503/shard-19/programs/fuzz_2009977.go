package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }

func main() {
    var v1 uint64 = 2
    var v2 int8 = 21
    var v3 bool = false
    v4 := make(map[int]uint32)
    v4[-9223372036854775808] = 65535
    v2 -= v2
    v1 += 2
    v5 := GenS1{F0: 80, F1: 4294967295}
    for fi6 := 0; fi6 < 10; fi6++ {
        switch (("foo bar" + "foo bar") + ("foo bar" + "a")) {
        case "hello":
            v7 := GenS1{F0: 0, F1: fi6}
            _ = v7
        case "\n\t":
            if v3 {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v3)
                case "\n\t":
                    var v8 int16 = 1
                    _ = v8
                default:
                    var v9 uint16 = 255
                    _ = v9
                }
                if v3 {
                    var v10 uint = 239
                    var v11 int = fi6
                    var v12 bool = true
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
        case "foo bar":
            if true {
                v13 := []int{(fi6 * 32767), (-(fi6)), 0, 4294967295}
                v14 := GenS2{F0: 1, F1: 32767, F2: 223}
                for fi15 := 0; fi15 < 10; fi15++ {
                    if v3 {
                        v16 := GenS3{F0: 0, F1: 2, F2: v3}
                        _ = v16
                    } else {
                        v17 := v4[1]
                        _ = v17
                    }
                    _ = fi15
                }
                _ = v13
                _ = v14
            } else {
                v1 = 2
                v2 = 1
                for fi18 := 0; fi18 < 10; fi18++ {
                    if (false || true) {
                        if (1.5 != -2.25) {
                            v19 := v4[(fi18 >> ((317) & 63))]
                            var v20 uint8 = 127
                            var v21 bool = (32767 < int(v2))
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            v22 := []int{}
                            var v23 float32 = -2.25
                            switch "hello" {
                            case "foo bar":
                                v2 += v2
                            case "":
                                var v24 uint = 797
                                _ = v24
                            }
                            _ = v22
                            _ = v23
                        }
                    }
                    v3 = (!v3)
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        fmt.Println(fi6)
                    case "":
                        v2 -= (v2 >> ((42) & 63))
                    }
                    v3 = (v3 || (42 <= 2147483647))
                    if false {
                        v25 := GenS2{F0: 65535, F1: 840, F2: 77}
                        switch (("" + ("hello" + "\n\t")) + "hello") {
                        case "hello":
                            v1 = 42
                        }
                        _ = v25
                    } else {
                        switch fi18 {
                        case 1:
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                v26 := GenS2{F0: 812, F1: 32767, F2: 166}
                                _ = v26
                            }
                        case 787:
                            v3 = ((0.0 + 0.1) == float64(-2.25))
                        }
                        if false {
                            switch fi6 {
                            case 2147483647:
                                var v27 string = "a"
                                _ = v27
                            default:
                                v3 = (-34 <= -9223372036854775808)
                            }
                            if ((!v3) && v3) {
                                v28 := []int{32767, fi6}
                                var v29 uint8 = 127
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                    _ = fi18
                }
            }
        }
        switch (fi6 ^ 42) {
        case 128:
            switch 584 {
            case 32767:
                v30 := GenS3{F0: -2147483648, F1: 65535, F2: v3}
                _ = v30
            case 32768:
                if ((0.0 * 1.5) >= 1.5) {
                    var v31 float64 = (0.1 - (float64((1.5 + 0.1)) * ((((1e10 - (((0.0 / 0.1) + float64((0.0 * 0.1))) - 3.14159)) - (3.14159 / 1e10)) / 1.5) - 0.1)))
                    _ = v31
                }
            }
        case 1:
            fmt.Println(v2, fi6)
        }
        _ = fi6
    }
    var v32 int8 = (v2 >> ((2) & 63))
    for fi33 := 0; fi33 < 15; fi33++ {
        if v3 {
            v4[int(v1)] = 2147483648
            v1 += v1
            switch "a" {
            case "foo bar":
                fmt.Println(v1, fi33, v1)
            case "\n\t":
                var v34 float32 = (0.1 / (1.5 / 1e10))
                _ = v34
            case "":
                fmt.Println(v3, v1)
            }
        }
        _ = len(v4)
        _ = fi33
    }
    v32 = v2
    v1 -= v1
    var v35 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
