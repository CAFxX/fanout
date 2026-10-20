package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 int16; F2 uint }
type GenS4 struct { F0 uint; F1 uint64; F2 int64 }

func main() {
    v1 := []int{128, 267, 202}
    v2 := []string{"\n\t"}
    var v3 uint32 = 447
    var v4 int16 = 42
    fmt.Println(v4, v4)
    if len(v2) > 0 {
        v2[0] = ""
    }
    switch "" {
    case "a":
        switch ("" + "\n\t") {
        case "a":
            v5 := GenS1{F0: 127}
            _ = v5
        case "\n\t":
            v6 := GenS2{F0: -490, F1: "foo bar"}
            _ = v6
        }
    case "\n\t":
        switch 4294967295 {
        case 128:
            v3 = 4294967295
        case -2147483648:
            if (float64(((1.5 / -2.25) * 0.1)) <= 0.1) {
                _ = len(v1)
            }
        case 2147483647:
            var v7 uint16 = 32767
            _ = v7
        }
    case "foo bar":
        v3 += 0
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint = 0
        _ = fi8
        _ = v9
    }
    if false {
        v3 += (v3 ^ v3)
    } else {
        if (157 >= uint8(v3)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = v3
        } else {
            if false {
                fmt.Println(v3, v4)
                if false {
                    v4 += 661
                } else {
                    v10 := GenS3{F0: (^(v3)), F1: 42, F2: 1}
                    _ = v10
                }
            } else {
                v3 = v3
                for fi11 := 0; fi11 < 2; fi11++ {
                    v3 = (v3 % ((1) | 1))
                    fmt.Println(fi11)
                    v3 += (v3 + (v3 ^ (v3 << ((514) & 63))))
                    _ = fi11
                }
                v4 += v4
            }
            if false {
                v3 = 65535
                if false {
                    v3 += (v3 / (((^(v3))) | 1))
                    if (42 <= 42) {
                        fmt.Println(v3)
                    }
                    fmt.Println(v3, v4)
                }
            }
            var v12 float64 = 1.5
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            if false {
                if (!false) {
                    if (!false) {
                        if (0 < 0) {
                            fmt.Println(v4)
                            v4 = (v4 >> ((1) & 63))
                        }
                    }
                } else {
                    v1 = append(v1, fi13)
                }
            }
            _ = fi13
        }
        if (255 <= 0) {
            v4 -= 666
            v14 := GenS3{F0: v3, F1: 1000, F2: 65535}
            var v15 uint16 = 42
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: 42, F1: ("foo bar" + ((("" + "hello") + "foo bar") + (("foo bar" + "\n\t") + ("hello" + "a"))))}
            v4 = (v4 ^ (v4 - v4))
            _ = v16
        }
    }
    if true {
        for fi17 := 0; fi17 < 20; fi17++ {
            if (4294967295 >= 18446744073709551615) {
                var v18 uint32 = 1
                var v19 int64 = -9223372036854775808
                v1 = append(v1, -9223372036854775808)
                _ = v18
                _ = v19
            } else {
                v3 += v3
                v3 += (v3 / ((v3) | 1))
                v20 := GenS4{F0: 953, F1: 18446744073709551615, F2: 9223372036854775807}
                _ = v20
            }
            if len(v1) > 0 {
                v1[0] = (fi17 << ((65535) & 63))
            }
            var v21 uint16 = 65535
            _ = fi17
            _ = v21
        }
        if false {
            for fi22 := 0; fi22 < 4; fi22++ {
                v1 = append(v1, fi22)
                _ = fi22
            }
            var v23 uint8 = 13
            _ = v23
        }
    }
    v3 += 2147483648
    v3 = (v3 | 243)
    v3 -= v3
    fmt.Println(v3, v4)
    v3 = (v3 << ((34) & 63))
    if (true && (18446744073709551615 != 42)) {
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 string = "hello"
            fmt.Println(v25)
            var v26 float32 = 0.0
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            v4 += 1000
            _ = fi27
        }
    }
    if false {
        if false {
            switch ("" + "a") {
            case "":
                fmt.Println(v3)
            case "\n\t":
                var v28 int16 = 42
                _ = v28
            case "hello":
                v3 = 1
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                v30 := GenS3{F0: v3, F1: (v4 % ((1000) | 1)), F2: 550}
                _ = fi29
                _ = v30
            }
            var v31 float64 = 0.0
            _ = v31
        }
        v3 = 580
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("foo bar" + ("a" + ("a" + "hello"))) {
    case "hello":
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
    default:
        if ((true || true) && (-538 >= 881)) {
            v3 -= (v3 >> ((0) & 63))
            for fi32 := 0; fi32 < 10; fi32++ {
                if ((^(fi32)) > fi32) {
                    v4 += v4
                    if true {
                        v33 := GenS4{F0: 65535, F1: 770, F2: int64(v3)}
                        _ = v33
                    }
                    var v34 int8 = -117
                    _ = v34
                }
                _ = fi32
            }
        }
    }
    v35 := GenS3{F0: v3, F1: v4, F2: 255}
    var v36 uint8 = 128
    v37 := GenS1{F0: 100}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
