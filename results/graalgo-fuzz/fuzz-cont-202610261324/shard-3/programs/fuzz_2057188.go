package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 uint64, p2 int16) string {
    v5 := []int{-2147483648, -2147483648, 2147483647}
    var v6 float64 = float64((float32(1.5) + 3.14159))
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8, p0)
            if (!(false || (!true))) {
                v5 = append(v5, 128)
                if (-128 <= int8(p0)) {
                    var v9 int8 = -128
                    _ = v9
                }
                p1 += 2
                var v10 string = ""
                p0 = p0
                _ = v10
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi8
        }
        _ = fi7
    }
    var v11 float64 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]uint32)
    v3 := []uint32{65535, 2147483648}
    var v4 int16 = 1
    v3 = append(v3, 4294967295)
    if ((v4 >> ((65535) & 63)) != v4) {
        if true {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint64 = 42
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: -676, F1: v4, F2: 491}
            if (!true) {
                delete(v2, "a")
            } else {
                fmt.Println(v1)
                v2[("a" + ("" + v1))] = 65535
            }
            _ = v14
        } else {
            v2[("a" + v1)] = 2147483648
            var v15 int8 = 80
            _ = v15
        }
    }
    fmt.Println(v1)
    if false {
        if false {
            if true {
                if false {
                    v1 = "hello"
                } else {
                    fmt.Println(v4)
                    fmt.Println(v1)
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    if true {
                        fmt.Println(v4, fi16, fi16)
                    } else {
                        fmt.Println(fi16)
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v2[("a" + v1)] = 2147483648
                            _ = fi17
                        }
                    }
                    _ = fi16
                }
            } else {
                fmt.Println(v4, v4)
                if (-9223372036854775808 > 536) {
                    fmt.Println(v1, v1)
                    if (42 <= 0) {
                        _ = len(v3)
                    }
                }
                var v18 bool = false
                _ = v18
            }
        } else {
            var v19 uint = 65535
            if (!false) {
                switch 49 {
                case -2147483648:
                    switch -586 {
                    case -606:
                        switch 4294967295 {
                        case 9223372036854775807:
                            var v20 uint8 = 1
                            _ = v20
                        }
                    case 128:
                        fmt.Println(v4)
                    }
                case 9223372036854775807:
                    switch ("\n\t" + "") {
                    case "":
                        v1 = "\n\t"
                    case "hello":
                        if (v19 == 839) {
                            for fi21 := 0; fi21 < 4; fi21++ {
                                v3 = append(v3, 65535)
                                for fi22 := 0; fi22 < 11; fi22++ {
                                    v23 := GenS1{F0: 42, F1: (v4 >> (((v19 ^ v19)) & 63)), F2: -9223372036854775808}
                                    _ = fi22
                                    _ = v23
                                }
                                _ = fi21
                            }
                            var v24 uint = 292
                            for fi25 := 0; fi25 < 12; fi25++ {
                                if (((2147483648 < 2) && true) && false) {
                                    switch 410 {
                                    case 600:
                                        var v26 int16 = 1
                                        _ = v26
                                    case -1:
                                        var v27 int = (fi25 & (-(fi25)))
                                        _ = v27
                                    case 251:
                                        fmt.Println(v1, v4, fi25)
                                    default:
                                        fmt.Println(fi25, fi25, v4)
                                    }
                                } else {
                                    var v28 int32 = -438
                                    var v29 uint64 = 4294967295
                                    v29 += (v29 & (v29 - v29))
                                    _ = v28
                                    _ = v29
                                }
                                fmt.Println(v24)
                                _ = fi25
                            }
                            _ = v24
                        }
                    default:
                        var v30 uint64 = 540
                        _ = v30
                    }
                case 42:
                    var v31 int8 = 1
                    _ = v31
                }
                v32 := GenS2{F0: 2147483647}
                _ = v32
            }
            var v33 int16 = v4
            _ = v19
            _ = v33
        }
        _ = len(v2)
    } else {
        var v34 int64 = -750
        var v35 uint = 65535
        fmt.Println(v4, v35)
        _ = v34
        _ = v35
    }
    v3 = append(v3, 0)
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(fi36)
        var v37 uint8 = 113
        _ = len(v3)
        _ = fi36
        _ = v37
    }
    switch "hello" {
    case "foo bar":
        switch (("a" + "foo bar") + "foo bar") {
        case "\n\t":
            v38 := GenS1{F0: 0, F1: v4, F2: 1}
            _ = v38
        default:
            fmt.Println(v1)
        }
    }
    var v39 float64 = float64(((1.5 - 0.1) / (((float32(0.0) / float32(0.1)) + 1.5) * 0.1)))
    v40 := v2[(v1 + ("foo bar" + v1))]
    v41 := GenS1{F0: -534, F1: -32768, F2: -9223372036854775808}
    v42 := GenS3{F0: 255}
    v1 = v1
    v40 = v40
    v40 += 42
    var v43 int = 128
    v44 := GenS3{F0: 255}
    fmt.Println(v1, v40, v39)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
