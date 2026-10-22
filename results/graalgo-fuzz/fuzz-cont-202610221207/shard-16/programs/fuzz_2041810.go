package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint = 65535
    var v2 uint8 = 157
    v3 := []string{(("\n\t" + "\n\t") + "\n\t"), "foo bar", "foo bar"}
    v4 := []uint32{305, 42, 1}
    v2 = v2
    for fi5 := 0; fi5 < 4; fi5++ {
        v1 = 4294967295
        var v6 uint16 = 1
        _ = fi5
        _ = v6
    }
    if (32 > (v2 >> ((18446744073709551615) & 63))) {
        var v7 int16 = 1000
        _ = v7
    } else {
        if (-32768 != -32768) {
            if len(v3) > 0 {
                v3[0] = (("a" + ("a" + "")) + "hello")
            }
            for fi8 := 0; fi8 < 10; fi8++ {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v1 = (v1 % ((v1) | 1))
                    _ = fi9
                }
                v1 = 4294967295
                _ = fi8
            }
            v4 = append(v4, 1)
        }
        fmt.Println(v2)
    }
    if false {
        v2 = v2
    } else {
        v2 -= v2
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = 589
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            v1 -= (^(v1))
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println(v1, v2)
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v1, v2)
        var v13 int8 = 0
        _ = fi12
        _ = v13
    }
    switch "a" {
    case "hello":
        switch "a" {
        case "foo bar":
            fmt.Println(v1, v1)
        case "":
            v14 := GenS1{F0: v2}
            _ = v14
        }
    default:
        if true {
            for fi15 := 0; fi15 < 7; fi15++ {
                switch fi15 {
                case 707:
                    v1 = (^(v1))
                case -984:
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                default:
                    v2 -= (v2 % ((0) | 1))
                }
                v1 += v1
                var v16 string = ""
                _ = fi15
                _ = v16
            }
            v2 += 42
        } else {
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS1{F0: (v2 + 2)}
                _ = fi17
                _ = v18
            }
            var v19 int32 = 0
            switch (("" + "a") + "hello") {
            case "hello":
                if (!false) {
                    var v20 uint16 = 0
                    v20 = (v20 / (((v20 % ((v20) | 1))) | 1))
                    _ = v20
                } else {
                    switch "a" {
                    case "a":
                        for fi21 := 0; fi21 < 13; fi21++ {
                            v2 = v2
                            v19 = 1
                            _ = fi21
                        }
                    case "":
                        if (826 < 1) {
                            v4 = append(v4, 528)
                        }
                    case "hello":
                        for fi22 := 0; fi22 < 6; fi22++ {
                            v23 := GenS1{F0: (v2 / (((v2 - 230)) | 1))}
                            _ = fi22
                            _ = v23
                        }
                    default:
                        var v24 float64 = 0.0
                        _ = v24
                    }
                    var v25 string = "\n\t"
                    var v26 uint32 = 300
                    _ = v25
                    _ = v26
                }
            }
            _ = v19
        }
    }
    if (!(!(!(false && true)))) {
        var v27 uint16 = 120
        v2 = uint8(v1)
        _ = v27
    } else {
        var v28 string = ((("" + "a") + "hello") + "foo bar")
        _ = v28
    }
    if (!(false && true)) {
        v29 := GenS1{F0: 242}
        _ = v29
    } else {
        v2 += v2
        v4 = append(v4, 0)
    }
    switch "foo bar" {
    case "hello":
        var v30 float32 = 1.5
        _ = v30
    case "foo bar":
        if false {
            v1 -= 42
            switch 32768 {
            case 9223372036854775807:
                fmt.Println(v1, v2)
            default:
                var v31 float64 = 1.5
                _ = v31
            }
        } else {
            var v32 uint64 = 2
            v32 = v32
            var v33 int = 32768
            _ = v32
            _ = v33
        }
    default:
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "foo bar")
        }
    }
    var v34 int8 = -128
    fmt.Println(v2, v34)
    if false {
        v35 := GenS2{F0: 719, F1: 4294967295}
        _ = v35
    }
    v36 := GenS2{F0: -761, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
}
