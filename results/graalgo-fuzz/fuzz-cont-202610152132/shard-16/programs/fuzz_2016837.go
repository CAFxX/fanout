package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint64 = 905
    var v3 uint16 = 1
    if true {
        for fi4 := 0; fi4 < 14; fi4++ {
            fmt.Println(fi4, fi4)
            switch "a" {
            case "":
                v5 := []int{402, fi4, (fi4 | (fi4 / ((1) | 1))), (fi4 | fi4)}
                _ = v5
            case "hello":
                var v6 int = fi4
                _ = v6
            case "\n\t":
                v7 := []int{4294967295, -2147483648, fi4, 935}
                _ = v7
            default:
                if true {
                    v3 = v3
                } else {
                    var v8 int16 = -325
                    _ = v8
                }
            }
            var v9 int16 = -1
            _ = fi4
            _ = v9
        }
        switch "hello" {
        case "a":
            for fi10 := 0; fi10 < 12; fi10++ {
                if (!(true && ((("a" > "a") || false) && true))) {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        fmt.Println(fi11, v3)
                        v2 += v2
                        v12 := GenS1{F0: -40}
                        v3 -= (v3 - v3)
                        var v13 uint = 566
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                } else {
                    if (!(-552 < 42)) {
                        switch -9223372036854775808 {
                        case 188:
                            v1 += v1
                        case -203:
                            if (false && (float64((((1.5 / 1.0) - (1.5 * (((float32(1e10) + 0.1) / 0.1) * 0.0))) * -2.25)) >= (-2.25 - (0.1 * 0.0)))) {
                                var v14 uint = 255
                                _ = v14
                            } else {
                                if true {
                                    fmt.Println(v3, fi10, v1)
                                    v1 = 9223372036854775807
                                } else {
                                    if false {
                                        var v15 uint64 = (v2 << ((943) & 63))
                                        fmt.Println(v15, v3)
                                        if ((!true) || false) {
                                            if false {
                                                v3 += (^(v3))
                                                var v16 bool = false
                                                _ = v16
                                            } else {
                                                if false {
                                                    fmt.Println(fi10, v15, v1)
                                                    fmt.Println(v2, v3, fi10)
                                                }
                                            }
                                            v1 += v1
                                        }
                                        _ = v15
                                    }
                                    for fi17 := 0; fi17 < 5; fi17++ {
                                        var v18 int16 = 1000
                                        fmt.Println(v1, v2, v1)
                                        _ = fi17
                                        _ = v18
                                    }
                                    v1 += v1
                                }
                            }
                        case 128:
                            v3 += 32767
                        }
                        v3 = v3
                        v2 = v2
                    }
                }
                var v19 int64 = (v1 % ((v1) | 1))
                v20 := GenS2{F0: 42, F1: 2147483647, F2: 799}
                _ = fi10
                _ = v19
                _ = v20
            }
        case "hello":
            var v21 string = ("foo bar" + "foo bar")
            _ = v21
        default:
            if (true || (v3 <= v3)) {
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 string = (("\n\t" + "foo bar") + ("foo bar" + "a"))
                    _ = fi22
                    _ = v23
                }
                var v24 uint = 0
                _ = v24
            } else {
                var v25 float32 = 1e10
                if (!(!false)) {
                    fmt.Println(v3, v2, v3)
                    var v26 int16 = 0
                    if true {
                        v2 = v2
                    } else {
                        v27 := GenS2{F0: v1, F1: 2147483647, F2: 0}
                        switch 1 {
                        case 42:
                            v2 -= v2
                        default:
                            v28 := []int{32767}
                            _ = v28
                        }
                        _ = v27
                    }
                    _ = v26
                }
                _ = v25
            }
        }
    } else {
        for fi29 := 0; fi29 < 5; fi29++ {
            if false {
                switch (fi29 | (fi29 | fi29)) {
                case -9223372036854775808:
                    v1 -= 0
                case 1:
                    fmt.Println(v3, fi29, v2)
                case 128:
                    if false {
                        fmt.Println(v2, v2, v2)
                        var v30 uint16 = v3
                        for fi31 := 0; fi31 < 12; fi31++ {
                            var v32 uint64 = (v2 / ((2) | 1))
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                default:
                    var v33 string = "\n\t"
                    _ = v33
                }
            } else {
                var v34 uint = 2
                fmt.Println(v2, v2, v34)
                _ = v34
            }
            _ = fi29
        }
        for fi35 := 0; fi35 < 17; fi35++ {
            v36 := []int{32767}
            var v37 int64 = v1
            v37 -= (v1 << ((268) & 63))
            _ = fi35
            _ = v36
            _ = v37
        }
        var v38 int64 = 42
        _ = v38
    }
    fmt.Println(v2)
    v2 -= (v2 ^ 2)
    var v39 int = 32768
    var v40 uint = 4294967295
    v39 += (-(v39))
    fmt.Println(v3, v40)
    v3 = v3
    v39 += -463
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}
