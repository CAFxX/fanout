package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint16; F2 int }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint; F1 int8 }

func main() {
    v1 := []int{-698, -856, 0, 544}
    v2 := make(map[uint32]string)
    var v3 uint = 255
    var v4 int32 = 1
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 65535}
    v7 := GenS2{F0: "\n\t"}
    if (false && (false || (42 != 65535))) {
        switch 127 {
        case 1:
            var v8 bool = true
            _ = v8
        case 2147483647:
            if (v4 > (v4 ^ (v4 ^ 86))) {
                for fi9 := 0; fi9 < 5; fi9++ {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        if len(v1) > 0 {
                            v1[0] = (fi10 % (((fi9 ^ fi9)) | 1))
                        }
                        _ = fi10
                    }
                    v11 := GenS3{F0: "a", F1: 283, F2: (fi9 & fi9)}
                    v12 := GenS4{F0: false, F1: fi9}
                    _ = fi9
                    _ = v11
                    _ = v12
                }
            }
        }
        v3 += v3
        var v13 string = "hello"
        _ = v13
    } else {
        if (-221 >= 1) {
            v1 = append(v1, 128)
            fmt.Println(v4, v3)
            if false {
                for fi14 := 0; fi14 < 19; fi14++ {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v3 = v3
                        _ = fi15
                    }
                    _ = fi14
                }
                if (172 >= 42) {
                    fmt.Println(v3, v4)
                    for fi16 := 0; fi16 < 6; fi16++ {
                        var v17 int32 = (v4 | 443)
                        v17 = -2147483648
                        _ = fi16
                        _ = v17
                    }
                    var v18 float32 = ((1e10 / 1.5) / 0.1)
                    _ = v18
                }
            } else {
                if ((true || false) || (32767 <= 86)) {
                    fmt.Println(v4, v3)
                    v3 = (v3 | v3)
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch (^(fi19)) {
                    case -720:
                        if (65535 <= 986) {
                            if len(v1) > 0 {
                                v1[0] = fi19
                            }
                            v20 := GenS3{F0: ("\n\t" + "\n\t"), F1: 255, F2: (fi19 << ((v3) & 63))}
                            _ = v20
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case 536:
                        v21 := GenS1{F0: 842}
                        _ = v21
                    default:
                        v3 = (v3 >> (((v3 / ((1) | 1))) & 63))
                    }
                    delete(v2, 344)
                    v3 = (v3 ^ (v3 >> ((v3) & 63)))
                    _ = fi19
                }
            }
        }
        var v22 int64 = 1
        v23 := GenS5{F0: (32767 <= 787), F1: -100}
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        v4 = (v4 | v4)
        for fi25 := 0; fi25 < 13; fi25++ {
            switch "hello" {
            case "":
                switch fi25 {
                case 32768:
                    if false {
                        switch ("foo bar" + "hello") {
                        case "hello":
                            v4 += (v4 ^ (v4 >> ((18446744073709551615) & 63)))
                        case "foo bar":
                            v3 = (v3 - 639)
                        default:
                            var v26 uint = 2
                            _ = v26
                        }
                        v27 := GenS2{F0: "a"}
                        fmt.Println(v4, v4, v4)
                        v4 = v4
                        _ = v27
                    } else {
                        v3 = (v3 + 276)
                        v28 := v5[("a" + "a")]
                        _ = v28
                    }
                default:
                    var v29 float64 = float64(0.1)
                    _ = v29
                }
            case "\n\t":
                var v30 int = 28
                _ = v30
            }
            v1 = append(v1, (-(fi24)))
            _ = fi25
        }
        _ = fi24
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 string = ("" + "hello")
        _ = fi31
        _ = v32
    }
    switch 42 {
    case -2147483648:
        var v33 int32 = v4
        _ = v33
    }
    fmt.Println(v3, v4)
    fmt.Println(v4, v4)
    v3 = 1
    var v34 int16 = 1000
    switch -627 {
    case 312:
        v34 = 1000
    case 32768:
        v3 = (v3 / ((0) | 1))
    case -2147483648:
        v35 := GenS6{F0: v3, F1: 17}
        _ = v35
    }
    var v36 int = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
