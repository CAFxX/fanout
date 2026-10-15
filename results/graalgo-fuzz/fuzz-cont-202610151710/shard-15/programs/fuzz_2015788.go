package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 bool = true
    var v4 uint16 = 698
    var v5 float32 = float32(float64(float32(((1e10 * ((((1.5 - 0.1) - (-2.25 - float64((0.0 / (3.14159 + 1.5))))) / 1.0) + float64((1e10 + 0.1)))) / 1.0))))
    v6 := []string{"\n\t"}
    if ((9223372036854775807 != 216) && v3) {
        var v7 uint32 = 152
        var v8 int64 = -274
        if len(v6) > 0 {
            v6[0] = (("foo bar" + "") + ("foo bar" + ("\n\t" + (("\n\t" + "hello") + ("a" + (("\n\t" + "foo bar") + ("" + "foo bar")))))))
        }
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        if v3 {
            v3 = (100 < 127)
            var v10 int = (fi9 / ((-2147483648) | 1))
            _ = v10
        }
        if v3 {
            for fi11 := 0; fi11 < 12; fi11++ {
                switch "foo bar" {
                case "a":
                    v3 = v3
                default:
                    var v12 string = ""
                    _ = v12
                }
                _ = fi11
            }
            v5 += -2.25
            if v3 {
                if v3 {
                    v13 := GenS1{F0: 42, F1: 1}
                    _ = v13
                }
                delete(v2, "foo bar")
                switch "foo bar" {
                case "foo bar":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "a":
                    v4 += v4
                case "\n\t":
                    switch (fi9 + fi9) {
                    case 127:
                        fmt.Println(v3, fi9, v4)
                    default:
                        switch 174 {
                        case 127:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            v5 += 1.5
                        }
                    }
                default:
                    v4 -= v4
                }
            }
        }
        _ = fi9
    }
    if ("hello" < "hello") {
        if len(v6) > 0 {
            v6[0] = ("foo bar" + "foo bar")
        }
        v3 = false
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        if (-2.25 >= float64(float32(1e10))) {
            fmt.Println(v5)
        } else {
            fmt.Println(v4, v3, v5)
            v5 -= 0.1
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 4; fi15++ {
                v5 -= 0.0
                var v16 float64 = 1e10
                if true {
                    if (v3 || v3) {
                        var v17 uint32 = 65535
                        var v18 uint64 = 42
                        switch ("hello" + "\n\t") {
                        case "a":
                            fmt.Println(v4, fi15, v18)
                        case "\n\t":
                            v19 := GenS1{F0: -612, F1: (^(v18))}
                            _ = v19
                        default:
                            var v20 float64 = 0.0
                            _ = v20
                        }
                        _ = v17
                        _ = v18
                    } else {
                        v21 := GenS1{F0: 0, F1: 754}
                        var v22 int16 = 368
                        switch (fi15 ^ (^(fi15))) {
                        case -2147483648:
                            v23 := GenS2{F0: 0, F1: true, F2: 2}
                            _ = v23
                        case 814:
                            v5 = v5
                        }
                        _ = v21
                        _ = v22
                    }
                    var v24 string = "a"
                    v4 = 2
                    _ = v24
                } else {
                    fmt.Println(v3, v16)
                    switch fi15 {
                    case 328:
                        v25 := GenS3{F0: 18446744073709551615, F1: 4294967295}
                        _ = v25
                    case 127:
                        v5 = (((v5 / 1e10) + (v5 + v5)) + v5)
                    }
                    if v3 {
                        delete(v1, fi14)
                        v3 = (v3 || ("hello" != (("hello" + "hello") + "hello")))
                        v16 += (float64(v5) * -2.25)
                    } else {
                        _ = len(v1)
                    }
                }
                _ = fi15
                _ = v16
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                v3 = v3
                delete(v1, -16)
                v27 := GenS1{F0: 42, F1: 0}
                _ = fi26
                _ = v27
            }
            _ = fi14
        }
    }
    switch 4294967295 {
    case 42:
        var v28 int64 = -9223372036854775808
        _ = v28
    case 128:
        if len(v6) > 0 {
            v6[0] = ("" + "")
        }
    default:
        v3 = (!v3)
    }
    if len(v6) > 0 {
        v6[0] = ((("\n\t" + "\n\t") + "hello") + "\n\t")
    }
    v4 -= v4
    if (255 != 55) {
        for fi29 := 0; fi29 < 16; fi29++ {
            if (!(0 != 1)) {
                _ = len(v6)
                v2["hello"] = (fi29 / (((-(fi29))) | 1))
            }
            _ = fi29
        }
    }
    var v30 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
