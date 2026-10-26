package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 249
    v3 := []string{"foo bar", "\n\t", "foo bar", ("\n\t" + ("\n\t" + "hello"))}
    var v4 string = ((("hello" + "hello") + "hello") + "a")
    v5 := make(map[int]bool)
    var v6 uint32 = 120
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "a" {
        case "":
            if true {
                var v8 uint8 = (v2 ^ v2)
                var v9 uint8 = v2
                var v10 uint16 = 774
                _ = v10
                _ = v8
                _ = v9
            } else {
                switch 1 {
                case 522:
                    _ = len(v5)
                case 2147483647:
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                default:
                    v11 := GenS1{F0: -101, F1: 42, F2: -84}
                    _ = v11
                }
            }
        default:
            var v12 int32 = (v1 >> ((18446744073709551615) & 63))
            _ = v12
        }
        if false {
            var v13 int32 = (v1 ^ v1)
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            var v14 int = 127
            _ = v13
            _ = v14
        }
        fmt.Println(v1)
        _ = fi7
    }
    var v15 int32 = v1
    if (42 != v2) {
        v6 = (v6 << ((42) & 63))
        if ((1.5 * (-2.25 * 1e10)) < 1.5) {
            if false {
                v16 := GenS1{F0: 1000, F1: -1, F2: 1}
                _ = v16
            } else {
                fmt.Println(v2, v2, v15)
                if true {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        if true {
                            fmt.Println(v6, v15)
                            v3 = append(v3, v4)
                        }
                        _ = fi17
                    }
                } else {
                    if false {
                        _ = len(v5)
                        var v18 string = v4
                        var v19 int8 = -1
                        _ = v18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 float64 = -2.25
                        _ = fi20
                        _ = v21
                    }
                    var v22 int8 = 1
                    _ = v22
                }
            }
            v1 -= (v15 / ((v15) | 1))
            v6 += 2147483647
        } else {
            v23 := GenS2{F0: -876}
            v24 := GenS2{F0: 579}
            switch v4 {
            case "foo bar":
                var v25 string = (v4 + ((("hello" + v4) + ("foo bar" + v4)) + v4))
                _ = v25
            case "":
                var v26 float32 = 3.14159
                _ = v26
            case "a":
                var v27 uint8 = 140
                _ = v27
            default:
                switch "a" {
                case "a":
                    var v28 float32 = 3.14159
                    _ = v28
                default:
                    var v29 uint8 = 255
                    _ = v29
                }
            }
            _ = v23
            _ = v24
        }
        _ = len(v5)
    }
    v30 := GenS2{F0: 1000}
    var v31 uint8 = (v2 >> ((0) & 63))
    v32 := GenS2{F0: 333}
    v33 := GenS2{F0: 32767}
    for fi34 := 0; fi34 < 16; fi34++ {
        switch (v4 + v4) {
        case "foo bar":
            fmt.Println(v4, v4)
        default:
            for fi35 := 0; fi35 < 13; fi35++ {
                switch v4 {
                case "":
                    var v36 float32 = 1.5
                    _ = v36
                case "\n\t":
                    v37 := GenS1{F0: int16(v15), F1: -9223372036854775808, F2: 42}
                    _ = v37
                case "hello":
                    var v38 uint16 = 746
                    _ = v38
                }
                v39 := GenS3{F0: -9223372036854775808}
                v40 := GenS3{F0: 534}
                _ = fi35
                _ = v39
                _ = v40
            }
        }
        _ = len(v3)
        _ = fi34
    }
    var v41 uint8 = 1
    var v42 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
