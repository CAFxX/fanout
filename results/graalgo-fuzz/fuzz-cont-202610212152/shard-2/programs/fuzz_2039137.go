package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 string; F2 uint16 }
type GenS4 struct { F0 int32; F1 int; F2 int8 }

func main() {
    var v1 int32 = 849
    var v2 uint32 = 269
    v3 := make(map[int]bool)
    v4 := []uint32{v2}
    var v5 uint8 = 1
    v6 := make(map[string]uint32)
    switch -949 {
    case 128:
        v5 = v5
    }
    v7 := GenS1{F0: ("" + ((("\n\t" + "foo bar") + "\n\t") + "a")), F1: -745}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if (!true) {
            _ = len(v6)
        }
        v8 := GenS1{F0: ("foo bar" + "foo bar"), F1: 1}
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                switch (("foo bar" + "\n\t") + "hello") {
                case "foo bar":
                    for fi10 := 0; fi10 < 15; fi10++ {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v2, v5, v2)
                            if true {
                                v11 := GenS2{F0: 2}
                                _ = v11
                            }
                            switch 127 {
                            case 32767:
                                v12 := GenS3{F0: 42, F1: "foo bar", F2: 654}
                                _ = v12
                            case -9223372036854775808:
                                v13 := GenS4{F0: v1, F1: -1, F2: 42}
                                _ = v13
                            default:
                                var v14 float64 = 3.14159
                                _ = v14
                            }
                            var v15 uint32 = (^(v2))
                            _ = v15
                        case "a":
                            v16 := GenS3{F0: (fi10 >> ((18446744073709551615) & 63)), F1: ("\n\t" + ("foo bar" + "hello")), F2: 1}
                            _ = v16
                        default:
                            fmt.Println(v2)
                        }
                        fmt.Println(v2)
                        _ = fi10
                    }
                case "":
                    var v17 uint8 = v5
                    _ = v17
                case "hello":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 string = "a"
                        v1 = v1
                        v5 -= v5
                        _ = fi18
                        _ = v19
                    }
                default:
                    v1 = v1
                }
                v2 = (v2 - v2)
                var v20 int32 = (v1 << ((42) & 63))
                _ = fi9
                _ = v20
            }
            v2 = v2
        } else {
            if false {
                fmt.Println(v5)
                for fi21 := 0; fi21 < 19; fi21++ {
                    v1 = 2147483647
                    _ = fi21
                }
                for fi22 := 0; fi22 < 20; fi22++ {
                    v5 += (v5 / ((v5) | 1))
                    _ = fi22
                }
            } else {
                v5 = (v5 | v5)
            }
        }
        _ = v8
    }
    switch "hello" {
    case "a":
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                v2 += 1
                if (!(((v1 << ((905) & 63)) <= (^(v1))) || ("hello" == "\n\t"))) {
                    v25 := v3[32767]
                    v26 := GenS3{F0: fi23, F1: ("" + "hello"), F2: 0}
                    v5 += (v5 * (^(v5)))
                    v3[(fi23 - fi23)] = (1 <= 0)
                    var v27 uint64 = 504
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(fi24, fi23, fi23)
                }
                var v28 string = "foo bar"
                var v29 uint8 = uint8(v2)
                _ = fi24
                _ = v28
                _ = v29
            }
            _ = fi23
        }
    case "":
        if (!false) {
            switch "a" {
            case "a":
                var v30 uint32 = (^(v2))
                _ = v30
            case "hello":
                for fi31 := 0; fi31 < 10; fi31++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v32 float64 = float64(0.0)
                        _ = v32
                    default:
                        for fi33 := 0; fi33 < 10; fi33++ {
                            v1 = v1
                            var v34 int = 104
                            delete(v3, -611)
                            _ = fi33
                            _ = v34
                        }
                    }
                    var v35 float32 = (0.0 / 1.0)
                    _ = fi31
                    _ = v35
                }
            }
            var v36 int = -1
            _ = v36
        }
    }
    var v37 int16 = -79
    var v38 int8 = 42
    v2 -= (^(v2))
    v6["a"] = (^(v2))
    v1 = v1
    var v39 string = ""
    v1 = (v1 & (v1 << ((1) & 63)))
    if (1 < -9223372036854775808) {
        v5 = 0
        delete(v3, -2147483648)
    } else {
        fmt.Println(v1, v39)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
