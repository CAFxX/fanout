package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 int16; F1 uint32; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint64 }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 int64 = 0
    var v5 int64 = v4
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            if len(v2) > 0 {
                v2[0] = 0
            }
            var v8 int64 = v5
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            _ = fi7
            _ = v8
        }
        if ((0.1 != 0.1) || false) {
            v9 := GenS1{F0: v4, F1: -128, F2: 448}
            v10 := GenS2{F0: 794, F1: fi6, F2: "hello"}
            var v11 uint32 = 2147483648
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := v3[(^(fi12))]
        var v14 int64 = v5
        switch v13 {
        case "foo bar":
            fmt.Println(v4)
        default:
            v15 := v3[(fi12 & fi12)]
            _ = v15
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(v5, v1)
        _ = fi16
    }
    fmt.Println(v4, v5)
    var v17 uint16 = 0
    v1 = 1.5
    _ = len(v2)
    if (18446744073709551615 <= 4294967295) {
        fmt.Println(v17, v5, v1)
        switch 127 {
        case 1:
            if len(v2) > 0 {
                v2[0] = 0
            }
        case 2147483647:
            v18 := GenS3{F0: 42, F1: 2147483648, F2: ""}
            _ = v18
        case -2147483648:
            for fi19 := 0; fi19 < 15; fi19++ {
                if (32767 != -1) {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 int32 = -1
                        var v22 uint = 65535
                        v3[(fi19 % ((fi19) | 1))] = "a"
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS3{F0: -32768, F1: 298, F2: (("a" + "foo bar") + "\n\t")}
                    if false {
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v1 = 0.0
                            var v25 int16 = 0
                            var v26 bool = false
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        var v27 int = fi19
                        v28 := GenS3{F0: 1, F1: 2147483647, F2: ("foo bar" + "hello")}
                        _ = v27
                        _ = v28
                    }
                    _ = v23
                } else {
                    switch (("hello" + (("\n\t" + "foo bar") + ("" + ""))) + ("" + "a")) {
                    case "foo bar":
                        for fi29 := 0; fi29 < 12; fi29++ {
                            var v30 float32 = 0.0
                            switch ((("a" + "foo bar") + "a") + "a") {
                            case "a":
                                switch ("foo bar" + "\n\t") {
                                case "a":
                                    var v31 bool = ("a" == "foo bar")
                                    _ = v31
                                case "\n\t":
                                    v30 -= ((v30 * v30) + float32(3.14159))
                                default:
                                    v4 -= (v4 / (((v4 << ((414) & 63))) | 1))
                                }
                            case "":
                                var v32 uint = 2
                                _ = v32
                            case "hello":
                                fmt.Println(v30, fi29)
                            }
                            _ = fi29
                            _ = v30
                        }
                    case "a":
                        if ((32767 == -32768) && (42 == 18446744073709551615)) {
                            var v33 uint8 = 255
                            _ = v33
                        }
                    case "\n\t":
                        fmt.Println(v17, fi19)
                    default:
                        var v34 uint64 = 386
                        _ = v34
                    }
                    v35 := GenS4{F0: v5}
                    _ = v35
                }
                _ = fi19
            }
        default:
            v1 = 1.5
        }
    } else {
        v36 := GenS5{F0: -1, F1: 792, F2: 0}
        _ = v36
    }
    v17 = (v17 + 255)
    v3[-298] = ""
    v17 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
