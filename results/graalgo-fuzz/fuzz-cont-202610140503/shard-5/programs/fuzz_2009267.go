package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    var v1 float64 = ((float64(3.14159) * 3.14159) + -2.25)
    var v2 uint16 = 867
    var v3 float32 = 0.1
    var v4 float32 = (-2.25 + 0.1)
    v5 := GenS1{F0: v2, F1: true, F2: 0}
    v6 := []int{4294967295}
    v1 = (((-2.25 / v1) * (1.5 - 0.1)) * float64(1e10))
    v4 = float32((v1 / float64(3.14159)))
    switch "" {
    case "foo bar":
        if (true && true) {
            v2 = 42
        }
    case "":
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int8 = 1
                var v10 int = fi8
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    case "hello":
        var v11 float32 = v4
        _ = v11
    default:
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint8 = 196
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v2)
    var v14 uint16 = v2
    switch "" {
    case "\n\t":
        switch 128 {
        case 4294967295:
            v15 := GenS2{F0: 0, F1: 1}
            _ = v15
        case -980:
            fmt.Println(v1, v3)
        default:
            for fi16 := 0; fi16 < 17; fi16++ {
                v6 = append(v6, (fi16 / ((fi16) | 1)))
                v2 -= v2
                _ = fi16
            }
        }
    case "a":
        v2 += (v2 + 65535)
    }
    var v17 int8 = -128
    fmt.Println(v4)
    if true {
        v18 := GenS2{F0: (v17 * v17), F1: -1}
        switch "\n\t" {
        case "hello":
            v14 = v2
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                for fi19 := 0; fi19 < 19; fi19++ {
                    switch (fi19 >> ((4294967295) & 63)) {
                    case -608:
                        switch ((("\n\t" + "hello") + "\n\t") + "") {
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "a":
                            v6 = append(v6, (fi19 - fi19))
                        }
                    }
                    _ = len(v6)
                    var v20 float64 = 1e10
                    _ = fi19
                    _ = v20
                }
            }
        case "":
            if ((float32(v1) + (v4 - 1e10)) == 0.1) {
                switch ("a" + "foo bar") {
                case "foo bar":
                    v14 += (^(v14))
                case "hello":
                    var v21 int8 = v17
                    _ = v21
                case "a":
                    v14 = 2
                default:
                    var v22 uint32 = 433
                    _ = v22
                }
                var v23 int32 = 0
                _ = v23
            }
        default:
            var v24 int64 = 0
            _ = v24
        }
        v25 := GenS2{F0: (v17 >> ((0) & 63)), F1: v17}
        _ = v18
        _ = v25
    }
    var v26 int16 = int16(v17)
    switch "hello" {
    case "\n\t":
        for fi27 := 0; fi27 < 3; fi27++ {
            v28 := GenS2{F0: (v17 << ((65535) & 63)), F1: 9}
            v17 = v17
            for fi29 := 0; fi29 < 4; fi29++ {
                v1 += 0.1
                _ = fi29
            }
            _ = fi27
            _ = v28
        }
    case "a":
        switch "hello" {
        case "foo bar":
            var v30 int8 = v17
            _ = v30
        case "a":
            for fi31 := 0; fi31 < 9; fi31++ {
                switch "a" {
                case "a":
                    var v32 bool = false
                    _ = v32
                case "hello":
                    v17 = (v17 % ((v17) | 1))
                default:
                    v3 = (v4 * v4)
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi31
            }
        default:
            if (("hello" + "") == ("\n\t" + ("foo bar" + "hello"))) {
                switch ("\n\t" + ("foo bar" + "")) {
                case "hello":
                    var v33 uint = 42
                    _ = v33
                default:
                    v17 = 1
                }
            } else {
                fmt.Println(v17, v14, v14)
                for fi34 := 0; fi34 < 10; fi34++ {
                    switch (fi34 | (fi34 + (fi34 ^ 32767))) {
                    case -2147483648:
                        v26 = v26
                    case 2147483647:
                        if len(v6) > 0 {
                            v6[0] = fi34
                        }
                    }
                    _ = fi34
                }
                var v35 float64 = v1
                _ = v35
            }
        }
    default:
        v14 += 689
    }
    v36 := GenS3{F0: ((2147483647 < 2147483647) || false), F1: 1}
    switch 1 {
    case 0:
        if (519 > 96) {
            switch (("hello" + "a") + "foo bar") {
            case "\n\t":
                fmt.Println(v4)
            case "a":
                var v37 float64 = v1
                _ = v37
            default:
                _ = len(v6)
            }
            for fi38 := 0; fi38 < 10; fi38++ {
                v3 = v4
                _ = fi38
            }
        }
    }
    switch "hello" {
    case "\n\t":
        if (-9223372036854775808 > 0) {
            for fi39 := 0; fi39 < 8; fi39++ {
                for fi40 := 0; fi40 < 18; fi40++ {
                    var v41 float32 = v3
                    var v42 uint = 1
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                var v43 uint = 18446744073709551615
                var v44 int16 = 1000
                _ = fi39
                _ = v43
                _ = v44
            }
        }
    default:
        v3 -= v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
