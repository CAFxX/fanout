package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 string = "foo bar"
    var v2 int16 = -755
    var v3 float64 = float64(0.0)
    if true {
        v4 := []int{-9223372036854775808, 32768, -9223372036854775808, 4294967295}
        var v5 int8 = -34
        switch 0 {
        case 753:
            switch 42 {
            case 955:
                v2 = v2
            case 9223372036854775807:
                if true {
                    v4 = append(v4, -1)
                    if false {
                        fmt.Println(v2)
                        if false {
                            v5 = -128
                            v6 := GenS1{F0: 42, F1: ((("hello" + "\n\t") + "a") + "\n\t"), F2: 257}
                            for fi7 := 0; fi7 < 6; fi7++ {
                                v5 = v5
                                v4 = append(v4, int(v5))
                                var v8 int8 = (v5 + (v5 << ((11) & 63)))
                                _ = fi7
                                _ = v8
                            }
                            _ = v6
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if false {
                                if true {
                                    switch v1 {
                                    case "\n\t":
                                        v3 = v3
                                    }
                                }
                                switch v1 {
                                case "":
                                    if len(v4) > 0 {
                                        v4[0] = -127
                                    }
                                case "foo bar":
                                    v9 := GenS1{F0: 65535, F1: (v1 + v1), F2: 18446744073709551615}
                                    _ = v9
                                }
                                switch -1 {
                                case 32767:
                                    var v10 uint8 = 141
                                    _ = v10
                                case 2147483647:
                                    v2 -= (v2 + (^(v2)))
                                case 4294967295:
                                    v2 = v2
                                }
                            } else {
                                var v11 int = 9223372036854775807
                                for fi12 := 0; fi12 < 3; fi12++ {
                                    fmt.Println(v5, v2)
                                    _ = fi12
                                }
                                _ = v11
                            }
                            v3 -= ((v3 + 3.14159) + 0.1)
                        }
                    }
                } else {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        fmt.Println(fi13, v5, v5)
                        if len(v4) > 0 {
                            v4[0] = (fi13 % (((fi13 | 128)) | 1))
                        }
                        _ = fi13
                    }
                    for fi14 := 0; fi14 < 18; fi14++ {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 float64 = (v3 / 0.1)
                            _ = fi15
                            _ = v16
                        }
                        _ = fi14
                    }
                    if len(v4) > 0 {
                        v4[0] = 9223372036854775807
                    }
                }
            case 1:
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int64 = -9223372036854775808
                    switch ("\n\t" + v1) {
                    case "a":
                        if (1 == v5) {
                            if false {
                                v19 := GenS1{F0: 185, F1: v1, F2: 0}
                                _ = v19
                            } else {
                                v18 = (v18 * v18)
                            }
                            v20 := GenS1{F0: 1, F1: "", F2: 65535}
                            _ = v20
                        }
                    case "":
                        v1 = (v1 + "\n\t")
                    case "hello":
                        v4 = append(v4, fi17)
                    default:
                        for fi21 := 0; fi21 < 17; fi21++ {
                            v22 := GenS1{F0: 42, F1: "hello", F2: 0}
                            var v23 int = fi21
                            v4 = append(v4, fi17)
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    }
                    v24 := GenS1{F0: 1, F1: ((("hello" + v1) + v1) + v1), F2: 65535}
                    _ = fi17
                    _ = v18
                    _ = v24
                }
            default:
                var v25 int8 = v5
                _ = v25
            }
        }
        _ = v4
        _ = v5
    }
    var v26 bool = (!(42 <= 1))
    fmt.Println(v26)
    if (131 >= 113) {
        v27 := GenS1{F0: 255, F1: "hello", F2: 18446744073709551615}
        _ = v27
    } else {
        v28 := []int{340}
        _ = v28
    }
    v2 -= 42
    v29 := GenS2{F0: 57, F1: 18446744073709551615, F2: 4294967295}
    v30 := GenS3{F0: (v2 - v2)}
    var v31 int64 = 561
    v32 := GenS4{F0: v26}
    if (18446744073709551615 > 429) {
        for fi33 := 0; fi33 < 6; fi33++ {
            fmt.Println(v2)
            for fi34 := 0; fi34 < 17; fi34++ {
                v35 := []int{}
                _ = fi34
                _ = v35
            }
            var v36 uint16 = 1
            _ = fi33
            _ = v36
        }
        var v37 int16 = v2
        _ = v37
    } else {
        var v38 int16 = 32767
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
