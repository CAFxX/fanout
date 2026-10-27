package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint16 = 1
    v2 := []int{}
    var v3 int = -1
    var v4 uint = 818
    v5 := []uint32{42, 0, 2147483648}
    if (!false) {
        switch "foo bar" {
        case "a":
            v6 := GenS1{F0: 950}
            _ = v6
        case "":
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(v3)
                if true {
                    v4 = (^(v4))
                } else {
                    fmt.Println(v1, v3, v4)
                }
                fmt.Println(fi7, v3)
                _ = fi7
            }
        case "\n\t":
            fmt.Println(v4, v3, v4)
        }
    } else {
        v1 = v1
    }
    var v8 uint16 = (v1 | v1)
    for fi9 := 0; fi9 < 10; fi9++ {
        v8 = 255
        var v10 uint32 = 2
        v2 = append(v2, 1)
        _ = fi9
        _ = v10
    }
    if (-9223372036854775808 <= (^(v3))) {
        switch 9223372036854775807 {
        case 893:
            v3 = (v3 + 791)
        case 311:
            for fi11 := 0; fi11 < 6; fi11++ {
                switch "\n\t" {
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "":
                    v12 := GenS2{F0: ("" + "a"), F1: 729, F2: true}
                    _ = v12
                case "hello":
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                }
                _ = fi11
            }
        default:
            var v13 int = 128
            _ = v13
        }
        v14 := GenS3{F0: 127, F1: "\n\t"}
        fmt.Println(v8)
        _ = v14
    } else {
        if (127 == -46) {
            if false {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v4 += v4
                    var v16 int16 = 1000
                    v17 := GenS2{F0: ("\n\t" + "foo bar"), F1: (v16 & 42), F2: (true || false)}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 float32 = -2.25
                    v20 := GenS4{F0: v8, F1: 251}
                    for fi21 := 0; fi21 < 13; fi21++ {
                        if true {
                            if ("a" >= "a") {
                                var v22 int16 = 32767
                                v2 = append(v2, (^(fi21)))
                                _ = v22
                            }
                            var v23 float64 = 0.1
                            _ = v23
                        }
                        if true {
                            if (1e10 < (0.1 / 1.5)) {
                                v3 += (v3 + 32767)
                                var v24 uint8 = 255
                                _ = v24
                            } else {
                                fmt.Println(fi18)
                                fmt.Println(fi18, v8)
                            }
                        }
                        _ = fi21
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            } else {
                var v25 string = ("a" + ("hello" + ""))
                fmt.Println(v8, v8)
                _ = v25
            }
            fmt.Println(v8, v3)
        } else {
            var v26 float32 = 0.0
            switch (v3 + (v3 >> ((v4) & 63))) {
            case -9223372036854775808:
                switch (("hello" + ("foo bar" + "")) + "\n\t") {
                case "a":
                    v4 -= v4
                default:
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v28 := GenS5{F0: v8, F1: 65535}
                        _ = fi27
                        _ = v28
                    }
                }
            case -1:
                if false {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        for fi30 := 0; fi30 < 13; fi30++ {
                            var v31 uint = 4294967295
                            v3 += (fi30 % ((v3) | 1))
                            _ = fi30
                            _ = v31
                        }
                        var v32 bool = (("foo bar" + "foo bar") <= ("\n\t" + "foo bar"))
                        if (95 != 128) {
                            var v33 float32 = (1e10 * float32((0.1 / 0.1)))
                            var v34 int32 = 0
                            _ = v33
                            _ = v34
                        }
                        _ = fi29
                        _ = v32
                    }
                } else {
                    var v35 float32 = 1.5
                    v26 = v26
                    var v36 uint8 = 0
                    _ = v35
                    _ = v36
                }
            case 1:
                v37 := GenS6{F0: v1, F1: 824}
                _ = v37
            default:
                v38 := GenS4{F0: 828, F1: 18446744073709551615}
                _ = v38
            }
            _ = v26
        }
        v1 = v1
    }
    v3 = (v3 << ((v4) & 63))
    fmt.Println(v8, v1, v3)
    var v39 uint = v4
    var v40 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v8
}
