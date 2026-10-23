package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint }
type GenS3 struct { F0 string }

func main() {
    var v1 float32 = (1.5 - 0.1)
    var v2 int = 708
    v3 := make(map[string]string)
    var v4 int16 = 0
    switch ("foo bar" + ("hello" + "a")) {
    case "hello":
        _ = len(v3)
    }
    v5 := GenS1{F0: -2147483648}
    var v6 string = ""
    v6 = v6
    v1 = (v1 / v1)
    fmt.Println(v6)
    _ = len(v3)
    v7 := v3[v6]
    fmt.Println(v2)
    if true {
        v4 += 321
        for fi8 := 0; fi8 < 12; fi8++ {
            switch ("foo bar" + v7) {
            case "hello":
                var v9 int = (v2 ^ (fi8 << ((42) & 63)))
                _ = v9
            case "a":
                for fi10 := 0; fi10 < 18; fi10++ {
                    v11 := GenS2{F0: 42, F1: 1, F2: 1}
                    var v12 string = "hello"
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                v1 += (1e10 / v1)
            }
            v6 = v6
            _ = fi8
        }
    }
    var v13 uint64 = 0
    if false {
        if true {
            var v14 uint64 = (v13 * v13)
            for fi15 := 0; fi15 < 8; fi15++ {
                switch (-(v2)) {
                case -1:
                    v16 := GenS1{F0: 42}
                    _ = v16
                case 42:
                    if true {
                        if false {
                            var v17 int64 = 9223372036854775807
                            var v18 uint = 399
                            for fi19 := 0; fi19 < 3; fi19++ {
                                fmt.Println(v18, v7)
                                fmt.Println(v2)
                                var v20 uint64 = (^(v14))
                                _ = fi19
                                _ = v20
                            }
                            _ = v17
                            _ = v18
                        }
                        if true {
                            if (!false) {
                                switch (^(fi15)) {
                                case -155:
                                    for fi21 := 0; fi21 < 4; fi21++ {
                                        fmt.Println(v4, v4, v1)
                                        var v22 int8 = 53
                                        fmt.Println(fi15, v22)
                                        _ = fi21
                                        _ = v22
                                    }
                                case -2147483648:
                                    var v23 string = "hello"
                                    _ = v23
                                }
                                var v24 string = v7
                                v6 = "\n\t"
                                _ = v24
                            }
                            v1 -= v1
                        } else {
                            for fi25 := 0; fi25 < 20; fi25++ {
                                v13 = (^(v13))
                                switch ((v7 + ("\n\t" + "foo bar")) + "") {
                                case "\n\t":
                                    if (fi15 == v2) {
                                        v26 := GenS2{F0: (v13 & (v14 | (v14 * (v14 & 1)))), F1: 255, F2: 65535}
                                        v27 := GenS2{F0: v14, F1: 50, F2: 2}
                                        _ = v26
                                        _ = v27
                                    } else {
                                        var v28 uint32 = 478
                                        v29 := GenS1{F0: -1}
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                switch (-(fi15)) {
                                case 0:
                                    fmt.Println(v2)
                                default:
                                    fmt.Println(v6, v7)
                                }
                                _ = fi25
                            }
                        }
                    } else {
                        v3[v7] = ""
                        switch 42 {
                        case -752:
                            v3["hello"] = (((v6 + ("hello" + (v6 + ("foo bar" + (v7 + (v6 + v7)))))) + v6) + "hello")
                        case -1:
                            delete(v3, (("a" + v7) + "a"))
                        case 128:
                            var v30 int32 = 42
                            _ = v30
                        }
                        var v31 float64 = -2.25
                        _ = v31
                    }
                case 127:
                    v32 := GenS3{F0: v6}
                    _ = v32
                default:
                    var v33 int16 = 1000
                    _ = v33
                }
                _ = fi15
            }
            for fi34 := 0; fi34 < 17; fi34++ {
                var v35 int16 = 32767
                _ = fi34
                _ = v35
            }
            _ = v14
        }
        v36 := GenS3{F0: "hello"}
        fmt.Println(v4)
        _ = v36
    } else {
        var v37 int16 = (v4 * (v4 & v4))
        var v38 int64 = -1
        v39 := GenS2{F0: v13, F1: 215, F2: 695}
        _ = v37
        _ = v38
        _ = v39
    }
    v40 := GenS2{F0: (v13 - v13), F1: 2, F2: 929}
    var v41 int64 = 0
    var v42 float32 = v1
    if true {
        fmt.Println(v1, v2)
        var v43 float64 = ((0.0 + float64(v42)) - (0.0 * ((3.14159 / -2.25) * (((1e10 / (1.5 / 1e10)) + ((0.1 * 1.5) - -2.25)) / (1e10 * 3.14159)))))
        v6 = "hello"
        _ = v43
    }
    switch v2 {
    case 42:
        var v44 float32 = v1
        _ = v44
    case 195:
        delete(v3, (v6 + v6))
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
