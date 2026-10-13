package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 int64, p1 int32) int {
    p0 += (^(p0))
    var v4 float32 = (3.14159 - 0.1)
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 int32 = (^(p1))
            if (2147483647 == 1) {
                p1 += -179
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 127
}

func main() {
    v1 := []uint32{321, 736, 42}
    v2 := make(map[string]int)
    var v3 int64 = 1
    if false {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int16 = 42
            var v10 int64 = (v3 * (-(v3)))
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        switch ("foo bar" + ((("foo bar" + ("foo bar" + "hello")) + "a") + ("a" + (("hello" + "foo bar") + "\n\t")))) {
        case "a":
            v1 = append(v1, 4294967295)
        case "\n\t":
            if false {
                switch 553 {
                case 2147483647:
                    v11 := GenS1{F0: -59, F1: -625, F2: v3}
                    _ = v11
                case 32767:
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v3 -= (v3 | v3)
                        if (false && false) {
                            v13 := GenS1{F0: 98, F1: (fi12 + fi12), F2: v3}
                            if true {
                                v3 += v3
                            } else {
                                if true {
                                    v3 = v3
                                    v3 -= v3
                                } else {
                                    _ = len(v2)
                                    var v14 uint64 = 2
                                    if false {
                                        v15 := GenS1{F0: -9, F1: genFn0((^(v3)), 42), F2: v3}
                                        v16 := GenS2{F0: -32768, F1: (fi12 - (-(fi12))), F2: 595}
                                        v14 = v14
                                        _ = v15
                                        _ = v16
                                    }
                                    _ = v14
                                }
                            }
                            _ = v13
                        }
                        switch (^(fi12)) {
                        case 127:
                            var v17 uint = 4294967295
                            _ = v17
                        }
                        _ = fi12
                    }
                case 0:
                    fmt.Println(v3)
                default:
                    var v18 bool = ((true && true) && true)
                    _ = v18
                }
            }
        }
        _ = len(v1)
    }
    if (!false) {
        fmt.Println(v3)
        if (-2.25 <= -2.25) {
            fmt.Println(v3)
        } else {
            v19 := GenS1{F0: 77, F1: 127, F2: (v3 & v3)}
            var v20 bool = (!true)
            for fi21 := 0; fi21 < 17; fi21++ {
                v1 = append(v1, 2147483647)
                fmt.Println(fi21, fi21, v3)
                _ = fi21
            }
            _ = v19
            _ = v20
        }
    } else {
        v3 = v3
        switch -716 {
        case 201:
            var v22 int16 = 42
            _ = v22
        case -962:
            switch "foo bar" {
            case "\n\t":
                v3 = v3
            }
        }
        v3 += v3
    }
    if true {
        switch "\n\t" {
        case "a":
            fmt.Println(v3)
        default:
            fmt.Println(v3)
        }
        switch 32768 {
        case 42:
            v23 := GenS2{F0: -1, F1: 32767, F2: 143}
            _ = v23
        case 9223372036854775807:
            var v24 uint16 = 2
            _ = v24
        case 127:
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 string = ""
                v1 = append(v1, 987)
                v3 += (^(v3))
                _ = fi25
                _ = v26
            }
        default:
            var v27 int8 = int8(v3)
            _ = v27
        }
    }
    v3 = v3
    fmt.Println(v3)
    v3 += v3
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        fmt.Println(fi28)
        v3 = (v3 % (((v3 % ((v3) | 1))) | 1))
        _ = fi28
    }
    switch (("" + "foo bar") + "a") {
    case "a":
        var v29 uint = 587
        _ = v29
    case "\n\t":
        v3 -= (-(v3))
    }
    if (0 > 42) {
        v3 -= (v3 - v3)
        var v30 uint8 = 68
        _ = v30
    } else {
        if false {
            v31 := GenS1{F0: 1, F1: 816, F2: -593}
            fmt.Println(v3)
            fmt.Println(v3)
            _ = v31
        }
        var v32 bool = true
        switch 72 {
        case 4294967295:
            v32 = true
        case 32768:
            var v33 uint16 = 42
            _ = v33
        case 0:
            if true {
                v34 := GenS3{F0: -1, F1: -1}
                _ = v34
            }
        }
        _ = v32
    }
    var v35 bool = true
    var v36 bool = ((288 > 4294967295) && (!(((!v35) && v35) && ((0.0 * ((0.1 / 1.5) + 1e10)) > (((1.5 / ((0.1 + 0.0) - 1.5)) + 3.14159) - (0.0 * ((3.14159 * 1.5) * (1.5 + 1e10))))))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
