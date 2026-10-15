package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 uint = 880
    var v2 int64 = -823
    var v3 string = "a"
    var v4 bool = (42 <= 55)
    v5 := []uint32{4294967295, 443, 556}
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int64 = v2
        if true {
            v9 := GenS1{F0: 504, F1: 301}
            _ = v9
        } else {
            v2 = 42
            switch v3 {
            case "a":
                if v4 {
                    if false {
                        if (!false) {
                            fmt.Println(v2)
                            var v10 uint64 = 136
                            switch (fi7 | (fi7 + (fi7 - fi7))) {
                            case -1:
                                for fi11 := 0; fi11 < 4; fi11++ {
                                    if len(v5) > 0 {
                                        v5[0] = 454
                                    }
                                    _ = fi11
                                }
                            }
                            _ = v10
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 0
                            }
                        }
                        fmt.Println(v1, fi7)
                        for fi12 := 0; fi12 < 8; fi12++ {
                            if v4 {
                                if len(v5) > 0 {
                                    v5[0] = 42
                                }
                                if false {
                                    v3 = v3
                                } else {
                                    v13 := GenS1{F0: 816, F1: 1}
                                    _ = v13
                                }
                            } else {
                                switch -2147483648 {
                                case -9223372036854775808:
                                    fmt.Println(v1)
                                }
                                var v14 uint16 = 497
                                _ = v14
                            }
                            v3 = (v3 + "\n\t")
                            v2 = (v2 - (-(v8)))
                            _ = fi12
                        }
                    }
                    if (false || ((((^(fi7)) >= 4294967295) || v4) && (!(!v4)))) {
                        var v15 int16 = 32767
                        _ = v15
                    }
                }
            case "foo bar":
                for fi16 := 0; fi16 < 10; fi16++ {
                    if ((v8 & (v2 ^ 42)) < v8) {
                        v1 = 0
                    } else {
                        v1 = 42
                        v17 := GenS1{F0: 18446744073709551615, F1: 255}
                        if (false || (-417 > 1)) {
                            fmt.Println(v4)
                            v6 += v6
                        } else {
                            if (false && v4) {
                                v6 += (0.1 - 0.0)
                            } else {
                                v4 = (v8 <= (v2 % ((v8) | 1)))
                                fmt.Println(v6)
                            }
                        }
                        _ = v17
                    }
                    var v18 int = (fi7 * (fi16 >> (((v1 % ((v1) | 1))) & 63)))
                    _ = fi16
                    _ = v18
                }
            }
            switch v3 {
            case "hello":
                for fi19 := 0; fi19 < 15; fi19++ {
                    if (!(32767 > 169)) {
                        var v20 string = v3
                        switch fi19 {
                        case 488:
                            v21 := GenS1{F0: 42, F1: 2}
                            _ = v21
                        case -946:
                            if false {
                                v22 := GenS1{F0: 18446744073709551615, F1: 0}
                                _ = v22
                            }
                        }
                        v23 := GenS1{F0: 367, F1: 65535}
                        var v24 uint64 = 1
                        if v4 {
                            var v25 string = "foo bar"
                            v26 := GenS1{F0: v24, F1: 249}
                            var v27 float64 = v6
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS1{F0: 961, F1: 555}
                        var v29 uint32 = 4294967295
                        _ = v20
                        _ = v23
                        _ = v24
                        _ = v28
                        _ = v29
                    } else {
                        if v4 {
                            var v30 uint8 = 200
                            _ = v30
                        }
                        var v31 uint32 = 0
                        v5 = append(v5, (v31 >> ((v1) & 63)))
                        var v32 uint32 = 335
                        _ = v31
                        _ = v32
                    }
                    v2 -= v2
                    switch fi7 {
                    case 4294967295:
                        v5 = append(v5, 2147483647)
                    default:
                        var v33 uint64 = 18446744073709551615
                        _ = v33
                    }
                    _ = fi19
                }
            default:
                v2 -= (v2 - (-(v8)))
            }
        }
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 int64 = (v2 | v8)
            _ = fi34
            _ = v35
        }
        _ = fi7
        _ = v8
    }
    var v36 int64 = (^(v2))
    v1 = 4294967295
    var v37 int32 = -1
    v36 -= 1
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v38 int8 = -90
        _ = v38
    case "a":
        v39 := GenS1{F0: 776, F1: 255}
        _ = v39
    default:
        fmt.Println(v6, v6)
    }
    fmt.Println(v37)
    var v40 string = ("" + "a")
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
