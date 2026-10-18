package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int8 = 111
    var v2 int = -2147483648
    var v3 uint64 = 1
    var v4 float64 = (3.14159 - ((((((float64(1e10) * ((0.1 / (1.5 / 1e10)) / 0.1)) + 3.14159) - 3.14159) - -2.25) / 1.0) + float64((0.0 + 3.14159))))
    var v5 uint16 = 594
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v2)
        var v8 int32 = -1
        _ = fi7
        _ = v8
    }
    var v9 int32 = int32(v3)
    if (("a" + (("foo bar" + "a") + (("\n\t" + ("a" + "\n\t")) + "\n\t"))) >= "foo bar") {
        v10 := []int{-656, v2, v2, v2}
        var v11 uint16 = (v5 % ((v5) | 1))
        _ = v10
        _ = v11
    } else {
        v3 -= (v3 | v3)
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint16 = v5
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v1, v3, v4)
        for fi15 := 0; fi15 < 17; fi15++ {
            switch ((("a" + "foo bar") + ("a" + "foo bar")) + ("hello" + ("foo bar" + "hello"))) {
            case "":
                v16 := []int{}
                var v17 uint8 = 1
                if (42 >= 1) {
                    v16 = append(v16, (fi15 ^ v2))
                }
                v18 := GenS1{F0: 9223372036854775807}
                var v19 uint16 = v5
                var v20 float32 = 0.1
                fmt.Println(fi15, v19, fi14)
                v19 = v5
                _ = v16
                _ = v17
                _ = v18
                _ = v19
                _ = v20
            default:
                fmt.Println(v9)
            }
            _ = fi15
        }
        switch "hello" {
        case "foo bar":
            var v21 uint32 = 0
            _ = v21
        case "":
            switch "hello" {
            case "a":
                switch v2 {
                case 42:
                    v3 = v3
                case -564:
                    v22 := GenS2{F0: -32768, F1: -1, F2: 2147483648}
                    _ = v22
                case -1:
                    var v23 float64 = (0.0 / (v4 + v4))
                    _ = v23
                default:
                    var v24 uint16 = 0
                    _ = v24
                }
            default:
                if (!false) {
                    switch fi14 {
                    case 0:
                        var v25 float64 = v4
                        _ = v25
                    case 128:
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS1{F0: -1}
                            v28 := []int{v2}
                            v28 = append(v28, int(v9))
                            fmt.Println(v2, v2, v3)
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    case -626:
                        for fi29 := 0; fi29 < 5; fi29++ {
                            v5 = v5
                            _ = fi29
                        }
                    }
                    var v30 uint8 = 128
                    _ = v30
                }
            }
        case "hello":
            for fi31 := 0; fi31 < 14; fi31++ {
                if false {
                    v9 -= (-(v9))
                    var v32 int16 = 393
                    _ = v32
                } else {
                    fmt.Println(v3)
                }
                var v33 int8 = (v1 * (-(v1)))
                v34 := GenS2{F0: 32767, F1: 1000, F2: 2147483648}
                _ = fi31
                _ = v33
                _ = v34
            }
        }
        _ = fi14
    }
    if true {
        v9 = v9
        switch "a" {
        case "\n\t":
            var v35 float64 = (0.0 + v4)
            _ = v35
        case "":
            for fi36 := 0; fi36 < 16; fi36++ {
                fmt.Println(v9, v2)
                for fi37 := 0; fi37 < 15; fi37++ {
                    switch ("" + ("foo bar" + "foo bar")) {
                    case "":
                        var v38 string = "\n\t"
                        _ = v38
                    }
                    _ = fi37
                }
                _ = fi36
            }
        case "foo bar":
            if false {
                v5 = v5
            }
        }
    }
    v39 := GenS3{F0: 1}
    v40 := GenS2{F0: 1, F1: 0, F2: 182}
    v9 -= (-(v9))
    switch 0 {
    case 32767:
        var v41 uint32 = 65535
        _ = v41
    case -1:
        v4 = v4
    default:
        switch ("hello" + "\n\t") {
        case "foo bar":
            v42 := GenS3{F0: 42}
            _ = v42
        default:
            v43 := []int{int(v3), 9223372036854775807, v2, (v2 - (^(v2)))}
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
