package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 bool; F1 int32; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 string; F2 bool }

func main() {
    var v1 bool = (18446744073709551615 <= 1)
    var v2 uint64 = 18446744073709551615
    v3 := []string{}
    var v4 bool = (!(18446744073709551615 < 2))
    for fi5 := 0; fi5 < 6; fi5++ {
        if true {
            switch 42 {
            case 128:
                switch "\n\t" {
                case "a":
                    for fi6 := 0; fi6 < 9; fi6++ {
                        var v7 uint32 = 2
                        _ = fi6
                        _ = v7
                    }
                }
            case 228:
                var v8 uint = 260
                _ = v8
            }
            for fi9 := 0; fi9 < 1; fi9++ {
                v2 = v2
                _ = fi9
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                if false {
                    v11 := GenS1{F0: 370, F1: 222, F2: v1}
                    _ = v11
                } else {
                    var v12 bool = false
                    _ = v12
                }
                if false {
                    var v13 int = (fi5 + -1)
                    v1 = false
                    var v14 uint16 = 37
                    _ = v13
                    _ = v14
                }
                v2 -= v2
                _ = fi10
            }
        }
        if (975 < 42) {
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 int = fi5
                _ = fi15
                _ = v16
            }
            if (1 != -32768) {
                var v17 int8 = -17
                _ = v17
            } else {
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 string = "\n\t"
                    switch ((v19 + v19) + "\n\t") {
                    case "\n\t":
                        v3 = append(v3, v19)
                        switch (("foo bar" + v19) + v19) {
                        case "a":
                            v1 = true
                        default:
                            v20 := GenS1{F0: 1000, F1: 65535, F2: (!(!((338 != 42) && (1000 >= -341))))}
                            _ = v20
                        }
                    }
                    if (810 >= 489) {
                        v21 := GenS1{F0: 32767, F1: 1, F2: v4}
                        v22 := GenS2{F0: v4, F1: 0, F2: 126}
                        if true {
                            fmt.Println(v2, v4, fi18)
                            var v23 float64 = float64((-2.25 * 3.14159))
                            _ = v23
                        }
                        _ = len(v3)
                        switch ("foo bar" + v19) {
                        case "hello":
                            var v24 int = 9223372036854775807
                            _ = v24
                        case "\n\t":
                            switch v19 {
                            case "a":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "foo bar":
                                v1 = false
                            default:
                                fmt.Println(v4, fi5)
                                v25 := GenS3{F0: v19}
                                _ = v25
                            }
                        default:
                            v4 = false
                        }
                        v2 += 1
                        _ = v21
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                }
                if (("a" + "foo bar") == "foo bar") {
                    v26 := GenS2{F0: (1 > 2147483647), F1: 1, F2: 100}
                    _ = v26
                } else {
                    var v27 int16 = -795
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 5; fi28++ {
                v2 += v2
                _ = fi28
            }
        }
        _ = fi5
    }
    fmt.Println(v2, v4, v4)
    for fi29 := 0; fi29 < 17; fi29++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = v1
        for fi30 := 0; fi30 < 14; fi30++ {
            fmt.Println(v1, v4, v4)
            _ = fi30
        }
        _ = fi29
    }
    switch ("a" + ("foo bar" + "foo bar")) {
    case "foo bar":
        var v31 uint32 = 2
        _ = v31
    case "hello":
        var v32 uint16 = 42
        _ = v32
    default:
        for fi33 := 0; fi33 < 17; fi33++ {
            _ = len(v3)
            for fi34 := 0; fi34 < 18; fi34++ {
                var v35 float64 = (0.0 / 0.1)
                switch ("foo bar" + "foo bar") {
                case "hello":
                    v2 = (v2 / ((42) | 1))
                case "\n\t":
                    v36 := GenS3{F0: "a"}
                    _ = v36
                case "a":
                    fmt.Println(v2, v2, v2)
                }
                var v37 int32 = 497
                if (-1 >= 127) {
                    var v38 string = "a"
                    _ = v38
                } else {
                    v37 = (v37 * v37)
                }
                v39 := GenS4{F0: (255 <= 195), F1: ((("foo bar" + "") + ("" + ("foo bar" + ("" + "\n\t")))) + "\n\t"), F2: false}
                _ = fi34
                _ = v35
                _ = v37
                _ = v39
            }
            fmt.Println(v4, fi33, v4)
            _ = fi33
        }
    }
    var v40 int64 = -692
    var v41 int8 = -75
    v42 := GenS1{F0: 900, F1: 42, F2: (!((0.0 - (0.0 + ((1e10 + 1e10) / 1.0))) < float32(0.1)))}
    var v43 bool = v1
    if v43 {
        fmt.Println(v2, v4, v4)
    }
    v40 += v40
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
