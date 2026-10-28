package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint16; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 424
    var v3 uint64 = 42
    var v4 uint8 = 1
    for fi5 := 0; fi5 < 1; fi5++ {
        v3 = 978
        _ = fi5
    }
    delete(v1, 1)
    fmt.Println(v2)
    v6 := GenS1{F0: 883, F1: -1, F2: 42}
    var v7 int32 = v2
    v8 := GenS1{F0: 555, F1: 0, F2: (v7 + v7)}
    v9 := GenS2{F0: int16(v7), F1: 2}
    if true {
        switch "" {
        case "hello":
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 uint32 = 2
                _ = fi10
                _ = v11
            }
        case "\n\t":
            for fi12 := 0; fi12 < 8; fi12++ {
                if (("a" < "") && false) {
                    fmt.Println(v4)
                    var v13 int8 = -128
                    _ = v13
                }
                fmt.Println(fi12)
                _ = fi12
            }
        }
    }
    if true {
        var v14 uint = 255
        _ = v14
    }
    fmt.Println(v7, v4)
    var v15 int16 = 0
    switch "a" {
    case "foo bar":
        switch 4294967295 {
        case 9223372036854775807:
            if (!(false || (17 <= -86))) {
                var v16 int64 = 42
                v16 += v16
                _ = v16
            } else {
                switch "hello" {
                case "a":
                    var v17 string = "a"
                    _ = v17
                case "hello":
                    if false {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            fmt.Println(fi18, fi18)
                            _ = fi18
                        }
                        var v19 float64 = 1.5
                        _ = v19
                    }
                default:
                    v2 = (v2 >> ((42) & 63))
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 uint64 = 37
                    fmt.Println(v4, v7)
                    _ = fi20
                    _ = v21
                }
            }
        case 4294967295:
            for fi22 := 0; fi22 < 18; fi22++ {
                switch ("a" + "foo bar") {
                case "hello":
                    if true {
                        if (uint32(v4) <= 2147483647) {
                            fmt.Println(v4)
                            v7 = -2147483648
                        }
                    }
                case "":
                    v23 := GenS2{F0: v15, F1: 65535}
                    _ = v23
                case "a":
                    var v24 int = -9223372036854775808
                    _ = v24
                default:
                    v25 := GenS1{F0: 255, F1: v15, F2: v7}
                    _ = v25
                }
                switch ("\n\t" + "foo bar") {
                case "a":
                    switch (fi22 | (fi22 * (fi22 - fi22))) {
                    case 824:
                        var v26 int64 = -1
                        _ = v26
                    case -2147483648:
                        var v27 uint16 = 78
                        _ = v27
                    case 127:
                        v1[1] = 9223372036854775807
                    }
                case "\n\t":
                    v28 := GenS2{F0: v15, F1: 2}
                    _ = v28
                case "foo bar":
                    delete(v1, fi22)
                default:
                    switch (fi22 >> ((4294967295) & 63)) {
                    case 32767:
                        if ("foo bar" > "") {
                            fmt.Println(v3)
                            switch "hello" {
                            case "\n\t":
                                var v29 uint16 = 255
                                _ = v29
                            }
                        } else {
                            switch ("a" + "hello") {
                            case "hello":
                                var v30 int32 = v2
                                _ = v30
                            default:
                                for fi31 := 0; fi31 < 4; fi31++ {
                                    v2 = v2
                                    _ = fi31
                                }
                            }
                        }
                    case 42:
                        var v32 uint = 4294967295
                        _ = v32
                    }
                }
                _ = fi22
            }
        case 32767:
            v33 := []int{32767, -2147483648, 0, 4294967295}
            _ = v33
        }
    case "a":
        if (!true) {
            v7 = (v2 * v7)
            if true {
                v34 := GenS1{F0: 713, F1: -32768, F2: v7}
                v7 += (v7 | (v7 | (v7 - v7)))
                for fi35 := 0; fi35 < 14; fi35++ {
                    switch "" {
                    case "\n\t":
                        for fi36 := 0; fi36 < 15; fi36++ {
                            v37 := GenS3{F0: 65535, F1: 127, F2: 2}
                            var v38 bool = true
                            v38 = ("a" != "")
                            v39 := GenS1{F0: 2, F1: v15, F2: (^(v2))}
                            v15 += 771
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    case "a":
                        if false {
                            for fi40 := 0; fi40 < 19; fi40++ {
                                fmt.Println(v15)
                                _ = fi40
                            }
                        } else {
                            v4 = 127
                        }
                    case "foo bar":
                        v4 = (v4 >> ((18446744073709551615) & 63))
                    }
                    _ = fi35
                }
                _ = v34
            }
        }
    case "\n\t":
        var v41 float64 = 0.0
        _ = v41
    }
    v7 = v2
    v15 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
