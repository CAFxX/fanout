package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int64 }
func genFn0() int64 {
    if (42 >= 2) {
        if ((3.14159 > 3.14159) || (true || true)) {
            var v5 uint8 = 2
            v6 := GenS1{F0: 125}
            switch ("a" + "a") {
            case "hello":
                for fi7 := 0; fi7 < 15; fi7++ {
                    v8 := []int{}
                    v8 = append(v8, fi7)
                    if true {
                        switch fi7 {
                        case 0:
                            switch "a" {
                            case "a":
                                var v9 uint16 = 65535
                                _ = v9
                            }
                        case 9223372036854775807:
                            v5 -= 255
                        case -695:
                            var v10 uint = 410
                            _ = v10
                        default:
                            if len(v8) > 0 {
                                v8[0] = (^(fi7))
                            }
                        }
                        for fi11 := 0; fi11 < 15; fi11++ {
                            v12 := GenS1{F0: 47}
                            if (true && (!(0 <= 4294967295))) {
                                fmt.Println(v5, fi11)
                                v5 = v5
                            } else {
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                            _ = fi11
                            _ = v12
                        }
                        if (!(0 > 255)) {
                            for fi13 := 0; fi13 < 16; fi13++ {
                                if (!((-524 >= 0) || (v5 == (^(v5))))) {
                                    v5 = (^(v5))
                                    if len(v8) > 0 {
                                        v8[0] = fi7
                                    }
                                    var v14 string = ("hello" + ((("" + (("\n\t" + "") + (("\n\t" + "hello") + ("foo bar" + "")))) + ("" + "hello")) + ((("foo bar" + ("" + "")) + (("a" + "hello") + "\n\t")) + ((((("a" + ("" + "hello")) + "foo bar") + (("hello" + "") + "\n\t")) + ("hello" + "\n\t")) + "a"))))
                                    _ = v14
                                }
                                switch fi13 {
                                case 42:
                                    var v15 uint32 = 2147483647
                                    _ = v15
                                case 127:
                                    var v16 uint16 = 65535
                                    _ = v16
                                case 922:
                                    v5 -= 255
                                default:
                                    switch "foo bar" {
                                    case "":
                                        v17 := GenS1{F0: 779}
                                        _ = v17
                                    case "foo bar":
                                        _ = len(v8)
                                    case "hello":
                                        var v18 uint32 = 2147483647
                                        _ = v18
                                    default:
                                        v5 += v5
                                    }
                                }
                                v5 += 127
                                _ = fi13
                            }
                        } else {
                            var v19 int32 = 994
                            if (((0 > 2) || (!true)) || false) {
                                v20 := GenS2{F0: 1, F1: 65535, F2: 1}
                                v19 += (v19 + 491)
                                _ = v20
                            } else {
                                fmt.Println(fi7)
                                var v21 float64 = (1e10 * (1e10 / 3.14159))
                                _ = v21
                            }
                            _ = v19
                        }
                    } else {
                        v5 += (v5 | v5)
                    }
                    _ = fi7
                    _ = v8
                }
            case "":
                v5 = v5
            }
            _ = v5
            _ = v6
        } else {
            v22 := GenS2{F0: 2, F1: 255, F2: 0}
            v23 := GenS1{F0: 883}
            _ = v22
            _ = v23
        }
        var v24 float64 = 0.0
        _ = v24
    } else {
        var v25 uint64 = 182
        _ = v25
    }
    var v26 uint8 = 2
    v27 := []int{-2147483648}
    v26 = v26
    _ = v26
    _ = v27
    return 519
}

func main() {
    var v1 uint32 = 79
    var v2 float64 = (0.0 / ((1.5 * 1e10) - (((0.0 - ((0.1 / float64(-2.25)) * 1.5)) * ((-2.25 * (1.5 * float64(0.0))) * float64(0.0))) - -2.25)))
    var v3 uint = 741
    var v4 uint32 = (v1 % (((^(v1))) | 1))
    switch ("\n\t" + "hello") {
    case "foo bar":
        v28 := GenS3{F0: 255, F1: 2, F2: 0}
        _ = v28
    }
    switch ("" + "foo bar") {
    case "hello":
        v29 := GenS4{F0: -2147483648, F1: 128}
        _ = v29
    case "a":
        v30 := GenS5{F0: (v4 >> ((797) & 63)), F1: 9223372036854775807}
        _ = v30
    case "foo bar":
        if ((v1 % ((2147483648) | 1)) == (^(v1))) {
            fmt.Println(v1, v4, v3)
        }
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 string = ((("a" + "a") + "") + ("a" + "\n\t"))
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        var v34 int = 4294967295
        _ = fi33
        _ = v34
    }
    fmt.Println(v2)
    v1 += 2
    v35 := GenS3{F0: 531, F1: 0, F2: -2147483648}
    var v36 int64 = genFn0()
    switch -830 {
    case -2147483648:
        v1 += (v4 >> (((v3 * v3)) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
