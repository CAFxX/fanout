package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 string; F1 int8 }
func genFn0(p0 bool, p1 uint64, p2 uint) int8 {
    var v5 int64 = 42
    var v6 int16 = -864
    v7 := GenS1{F0: 4294967295, F1: 471, F2: (p2 | (p2 | p2))}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int32 = 21
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 uint8, p2 int32) uint {
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: 4294967295, F1: 419, F2: 18446744073709551615}
        v12 := GenS1{F0: 894, F1: 1, F2: 4294967295}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ("foo bar" + ("" + ("foo bar" + ("\n\t" + "hello")))) {
    case "\n\t":
        var v13 uint8 = p1
        _ = v13
    case "foo bar":
        var v14 int64 = int64(p0)
        _ = v14
    default:
        var v15 uint32 = 564
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(fi16, p2, fi16)
        _ = fi16
    }
    p0 += p0
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 1
    v3 := make(map[string]int)
    v4 := []uint32{}
    fmt.Println(v2)
    var v17 uint8 = (^(v2))
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                var v20 float32 = 0.0
                _ = v20
            } else {
                if false {
                    var v21 float64 = -2.25
                    _ = v21
                } else {
                    v17 = 128
                    v22 := GenS1{F0: 0, F1: 4294967295, F2: 482}
                    _ = v22
                }
            }
            v23 := GenS1{F0: 2, F1: 2147483648, F2: 42}
            _ = fi19
            _ = v23
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            v2 = v17
            switch (-(fi24)) {
            case -2147483648:
                var v25 uint32 = 1
                _ = v25
            case 32767:
                v17 = (v2 ^ 127)
            case 9223372036854775807:
                v2 += v2
            }
            _ = fi24
        }
        v17 = v2
        _ = fi18
    }
    var v26 float32 = -2.25
    var v27 uint = genFn1(886, v2, 0)
    fmt.Println(v26)
    switch 32767 {
    case 9223372036854775807:
        v26 = (v26 / 1e10)
    case -800:
        var v28 int32 = 1
        _ = v28
    default:
        if true {
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 int64 = -318
                switch ((("a" + "\n\t") + "") + ("\n\t" + "a")) {
                case "\n\t":
                    v1[(^(fi29))] = "foo bar"
                case "hello":
                    if (("foo bar" + "foo bar") > "") {
                        v2 -= (^(v2))
                    } else {
                        switch "hello" {
                        case "\n\t":
                            v30 = (^(v30))
                        case "a":
                            fmt.Println(v27, v26, fi29)
                        }
                    }
                case "a":
                    for fi31 := 0; fi31 < 10; fi31++ {
                        if true {
                            v17 += 42
                        } else {
                            v17 = 128
                            fmt.Println(v17, v17)
                            var v32 string = "foo bar"
                            var v33 int64 = (v30 ^ 0)
                            fmt.Println(v27, fi29)
                            _ = v32
                            _ = v33
                        }
                        var v34 int = fi29
                        var v35 uint16 = 2
                        _ = fi31
                        _ = v34
                        _ = v35
                    }
                }
                switch fi29 {
                case 5:
                    for fi36 := 0; fi36 < 17; fi36++ {
                        var v37 uint8 = (^(v2))
                        _ = fi36
                        _ = v37
                    }
                case 32768:
                    delete(v1, 128)
                default:
                    _ = len(v1)
                }
                _ = fi29
                _ = v30
            }
            delete(v1, 4294967295)
            v27 = 4294967295
        }
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        fmt.Println(v26, v27, fi38)
        _ = fi38
    }
    if (true && false) {
        var v39 uint = 0
        var v40 bool = (-2.25 == 0.1)
        _ = v39
        _ = v40
    } else {
        v26 += v26
        v41 := GenS2{F0: 1, F1: -32768}
        _ = v41
    }
    v27 += (^(v27))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v17 -= v17
    fmt.Println(v17, v2)
    v17 += 213
    v42 := GenS3{F0: (((((("foo bar" + "foo bar") + "foo bar") + "a") + "a") + "hello") + ("foo bar" + "a")), F1: 93}
    var v43 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
