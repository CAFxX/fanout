package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int64; F1 int32 }

func main() {
    v1 := []uint32{42, 115, 4294967295, 2}
    var v2 uint64 = 18446744073709551615
    var v3 uint8 = 128
    v4 := make(map[string]bool)
    var v5 int8 = int8(v2)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 float64 = (1.5 * 1e10)
        switch "hello" {
        case "foo bar":
            var v8 int8 = (^(v5))
            _ = v8
        case "\n\t":
            v5 += v5
        case "":
            var v9 int8 = (v5 ^ v5)
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v2)
    v10 := GenS1{F0: -32768, F1: false}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            if ("hello" != "") {
                fmt.Println(fi12)
                v2 = (v2 & v2)
                v3 = v3
            } else {
                if true {
                    v13 := GenS1{F0: 32767, F1: (1e10 <= -2.25)}
                    if (!false) {
                        v3 -= (v3 * (v3 ^ (v3 & (v3 >> ((830) & 63)))))
                        if false {
                            v14 := GenS1{F0: 884, F1: ((!(-2147483648 >= -531)) || (!(false && (v3 < (^(v3))))))}
                            if true {
                                v5 = v5
                                v4["\n\t"] = (1 < 660)
                                v1 = append(v1, 306)
                            }
                            _ = v14
                        }
                        if ((0 < -32768) && false) {
                            switch 128 {
                            case -242:
                                switch (("hello" + "a") + ("\n\t" + "hello")) {
                                case "a":
                                    if false {
                                        if true {
                                            v15 := GenS1{F0: -1, F1: (-1 > 526)}
                                            switch 2147483647 {
                                            case 203:
                                                switch "" {
                                                case "a":
                                                    v16 := GenS2{F0: 455}
                                                    _ = v16
                                                case "\n\t":
                                                    switch ((("a" + "") + "hello") + ("a" + ("hello" + "\n\t"))) {
                                                    case "a":
                                                        v2 = 2
                                                    }
                                                case "":
                                                    fmt.Println(fi11, v2)
                                                }
                                            case 0:
                                                v2 += (v2 >> ((4294967295) & 63))
                                            }
                                            _ = v15
                                        } else {
                                            fmt.Println(v2)
                                            fmt.Println(v5, fi12)
                                        }
                                        fmt.Println(v2, v5, v2)
                                        if false {
                                            if false {
                                                if false {
                                                    switch "hello" {
                                                    case "a":
                                                        v2 += v2
                                                    case "foo bar":
                                                        var v17 int64 = 42
                                                        _ = v17
                                                    case "":
                                                        delete(v4, "a")
                                                    default:
                                                        fmt.Println(fi11, v5, v2)
                                                    }
                                                    v3 = v3
                                                    var v18 uint32 = 93
                                                    _ = v18
                                                } else {
                                                    v5 = (^(v5))
                                                }
                                                v5 = v5
                                                _ = len(v4)
                                            } else {
                                                v3 = (v3 & 128)
                                            }
                                            v3 += 61
                                            var v19 uint64 = 1
                                            _ = v19
                                        }
                                        var v20 bool = true
                                        _ = v20
                                    } else {
                                        v2 += v2
                                        v4["a"] = true
                                    }
                                default:
                                    switch (fi11 % ((9223372036854775807) | 1)) {
                                    case 32768:
                                        v21 := v4["\n\t"]
                                        _ = v21
                                    default:
                                        var v22 uint64 = (v2 / ((1) | 1))
                                        _ = v22
                                    }
                                }
                            case -2147483648:
                                var v23 uint32 = 582
                                _ = v23
                            case 1:
                                v5 -= 42
                            }
                        }
                        v24 := GenS1{F0: 0, F1: (("foo bar" + "hello") <= "a")}
                        _ = v24
                    } else {
                        fmt.Println(fi12)
                    }
                    _ = v13
                }
                if (false || (!(65535 == 592))) {
                    switch fi12 {
                    case 28:
                        switch (^(fi11)) {
                        case -806:
                            var v25 uint = 778
                            _ = v25
                        case -589:
                            var v26 float32 = (-2.25 / (0.1 + 1.5))
                            _ = v26
                        case 9223372036854775807:
                            v27 := GenS3{F0: (fi12 << ((0) & 63)), F1: -788, F2: 32767}
                            v5 -= -104
                            var v28 int16 = 1
                            _ = v27
                            _ = v28
                        default:
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = 620
                                }
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                    }
                    if true {
                        var v29 int32 = 1
                        var v30 int64 = 253
                        _ = v29
                        _ = v30
                    } else {
                        switch fi11 {
                        case 9223372036854775807:
                            var v31 int32 = 2147483647
                            _ = v31
                        }
                    }
                    fmt.Println(fi12)
                }
                v32 := GenS4{F0: 255}
                _ = v32
            }
            v33 := GenS5{F0: fi12, F1: "hello"}
            var v34 int8 = (v5 & (v5 ^ (v5 % (((v5 & v5)) | 1))))
            _ = fi12
            _ = v33
            _ = v34
        }
        v35 := GenS5{F0: fi11, F1: "foo bar"}
        fmt.Println(v2, v5, fi11)
        _ = fi11
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        fmt.Println(v3, v2, v5)
        _ = fi36
    }
    v37 := GenS6{F0: -1, F1: 42}
    v38 := GenS3{F0: 32768, F1: 43, F2: 1}
    fmt.Println(v3)
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
