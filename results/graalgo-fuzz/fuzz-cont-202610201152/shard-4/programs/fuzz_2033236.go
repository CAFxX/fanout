package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 uint8 = 2
    v5 := make(map[string]int)
    v4 += 128
    switch 2147483647 {
    case 32768:
        switch -1 {
        case 128:
            switch (("" + (("foo bar" + "\n\t") + "")) + ("a" + (("a" + "a") + ("a" + "foo bar")))) {
            case "hello":
                v6 := GenS1{F0: 127, F1: true, F2: 0}
                _ = v6
            default:
                for fi7 := 0; fi7 < 7; fi7++ {
                    var v8 int32 = -365
                    _ = fi7
                    _ = v8
                }
            }
        case 127:
            for fi9 := 0; fi9 < 4; fi9++ {
                if true {
                    v10 := GenS2{F0: 438, F1: int64(v4), F2: 18446744073709551615}
                    v11 := v5["\n\t"]
                    fmt.Println(v11, v4, fi9)
                    _ = v10
                    _ = v11
                } else {
                    if true {
                        v12 := GenS1{F0: 127, F1: false, F2: 42}
                        if ((100 >= 0) && (255 <= 42)) {
                            if (true || false) {
                                switch (-(fi9)) {
                                case 32768:
                                    v13 := GenS2{F0: 18446744073709551615, F1: 0, F2: 119}
                                    _ = v13
                                }
                                v4 = v4
                            } else {
                                v4 -= v4
                            }
                            if true {
                                var v14 int8 = 42
                                _ = v14
                            }
                            v15 := GenS2{F0: 713, F1: -1, F2: 2}
                            _ = v15
                        }
                        _ = v12
                    }
                }
                var v16 uint16 = 65535
                _ = fi9
                _ = v16
            }
        default:
            v17 := GenS3{F0: ((3.14159 / (((0.0 + (((-2.25 * ((((3.14159 * 3.14159) + (0.0 - -2.25)) * 1.5) * 0.0)) - 1e10) - 0.0)) * (1.5 - 0.1)) - -2.25)) >= -2.25), F1: v4}
            _ = v17
        }
    case 1:
        var v18 bool = (!((1 < 32767) || ((0.1 / -2.25) != float32((0.0 * (-2.25 / (1.5 - -2.25)))))))
        _ = v18
    default:
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS4{F0: (v4 >> ((65535) & 63)), F1: 42}
            for fi21 := 0; fi21 < 2; fi21++ {
                switch "" {
                case "a":
                    v22 := GenS5{F0: 255, F1: (((false || (1 > 0)) && (255 > (v4 << ((737) & 63)))) && (false || false)), F2: 65535}
                    _ = v22
                case "\n\t":
                    var v23 int64 = 1
                    _ = v23
                case "foo bar":
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            v24 := GenS5{F0: 326, F1: (!(true && (42 != 367))), F2: 824}
                            v4 = v4
                            _ = v24
                        }
                    case "hello":
                        switch "foo bar" {
                        case "a":
                            var v25 uint16 = 255
                            _ = v25
                        case "":
                            v26 := GenS1{F0: 94, F1: true, F2: 1}
                            _ = v26
                        case "hello":
                            v4 -= (v4 * v4)
                        }
                    default:
                        fmt.Println(v4, v4, fi21)
                    }
                default:
                    fmt.Println(fi21)
                }
                _ = fi21
            }
            fmt.Println(fi19)
            _ = fi19
            _ = v20
        }
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 float64 = 0.0
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v30 string = ("" + "\n\t")
    v31 := GenS4{F0: 16, F1: 62}
    var v32 string = "foo bar"
    v30 = "foo bar"
    fmt.Println(v32, v4, v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
