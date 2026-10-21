package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 string; F1 int8; F2 bool }

func main() {
    var v1 int8 = 1
    var v2 float64 = 0.0
    var v3 string = "foo bar"
    v4 := []string{("hello" + ((v3 + (("a" + v3) + "a")) + ("foo bar" + (("" + v3) + (v3 + "a"))))), (v3 + ((v3 + v3) + v3))}
    var v5 string = v3
    _ = len(v4)
    var v6 uint8 = 127
    if false {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = 3.14159
            switch fi7 {
            case -9223372036854775808:
                var v9 int32 = 1
                _ = v9
            case -852:
                v2 = v2
            case 128:
                for fi10 := 0; fi10 < 16; fi10++ {
                    v4 = append(v4, v3)
                    _ = fi10
                }
            default:
                for fi11 := 0; fi11 < 3; fi11++ {
                    fmt.Println(v2, v6)
                    var v12 uint32 = 0
                    if (uint16(v12) >= 140) {
                        var v13 uint16 = 0
                        _ = v13
                    } else {
                        var v14 uint16 = 0
                        fmt.Println(v5, v1, v14)
                        fmt.Println(v14)
                        _ = v14
                    }
                    var v15 bool = true
                    _ = fi11
                    _ = v12
                    _ = v15
                }
            }
            v6 += 0
            _ = fi7
            _ = v8
        }
    } else {
        v1 -= v1
        if (-2.25 < (-2.25 - 1.5)) {
            v1 = -128
            switch 2147483647 {
            case -9223372036854775808:
                for fi16 := 0; fi16 < 11; fi16++ {
                    switch fi16 {
                    case -1:
                        var v17 int = fi16
                        _ = v17
                    case 128:
                        if (false && true) {
                            var v18 int32 = 2147483647
                            var v19 uint16 = 65535
                            _ = v18
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 5; fi20++ {
                                fmt.Println(fi20, v6, v2)
                                var v21 int32 = -1
                                v5 = v3
                                _ = fi20
                                _ = v21
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case -653:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        var v22 uint16 = 809
                        _ = v22
                    }
                    _ = fi16
                }
            }
            var v23 uint32 = 42
            _ = v23
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                switch ((v5 + v3) + (("a" + v3) + v3)) {
                case "":
                    fmt.Println(fi24, v1, v2)
                case "foo bar":
                    var v25 uint8 = 42
                    _ = v25
                }
                _ = fi24
            }
            v26 := GenS1{F0: 127, F1: (482 < 864)}
            _ = v26
        }
    }
    var v27 int32 = -298
    var v28 int16 = 187
    var v29 uint64 = 625
    v30 := GenS1{F0: (^(v1)), F1: false}
    v3 = v3
    _ = len(v4)
    if ((v28 & v28) < v28) {
        for fi31 := 0; fi31 < 7; fi31++ {
            if false {
                if (-2.25 <= 3.14159) {
                    v32 := GenS2{F0: -727, F1: 255}
                    for fi33 := 0; fi33 < 13; fi33++ {
                        var v34 uint32 = 42
                        var v35 uint8 = v6
                        v36 := GenS2{F0: -1, F1: 2}
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    v29 += v29
                    _ = v32
                } else {
                    v1 -= int8(v6)
                }
                v29 = v29
                fmt.Println(v29, v2)
            }
            _ = fi31
        }
        v37 := GenS2{F0: -1, F1: 550}
        for fi38 := 0; fi38 < 7; fi38++ {
            fmt.Println(v28)
            if len(v4) > 0 {
                v4[0] = v5
            }
            v39 := GenS3{F0: "hello", F1: 42, F2: (42 < 1)}
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        v40 := GenS2{F0: -530, F1: 4294967295}
        _ = v40
    }
    if false {
        var v41 float32 = 0.0
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        var v43 int16 = v28
        if len(v4) > 0 {
            v4[0] = v5
        }
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
