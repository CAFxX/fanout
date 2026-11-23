package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = 9223372036854775807
    v3 := []int{-203, -876, 128, 32767}
    switch ("foo bar" + ("" + "a")) {
    case "a":
        v4 := GenS1{F0: -128, F1: 848, F2: 2147483647}
        _ = v4
    case "foo bar":
        v5 := GenS1{F0: 42, F1: -390, F2: 1}
        _ = v5
    case "hello":
        var v6 int16 = 1000
        _ = v6
    }
    var v7 string = ("a" + "\n\t")
    v2 += (v2 % ((1) | 1))
    var v8 string = "hello"
    v1 = v1
    v7 = v8
    if ((128 < 29) || false) {
        v9 := GenS2{F0: 1}
        _ = v9
    } else {
        var v10 float32 = (0.1 + float32(-2.25))
        switch 32768 {
        case 32768:
            var v11 int16 = -32768
            _ = v11
        case 32767:
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        case -1:
            for fi12 := 0; fi12 < 11; fi12++ {
                if true {
                    switch v7 {
                    case "a":
                        if false {
                            v8 = ((v8 + v8) + v8)
                            v1 = 2
                        }
                    case "\n\t":
                        var v13 uint = 134
                        _ = v13
                    case "foo bar":
                        v14 := GenS1{F0: -1, F1: 0, F2: -2147483648}
                        _ = v14
                    }
                }
                _ = fi12
            }
        default:
            switch (((v7 + (((v7 + ("foo bar" + "a")) + v8) + "hello")) + v7) + v7) {
            case "\n\t":
                if false {
                    v3 = append(v3, -2147483648)
                    if true {
                        _ = len(v3)
                        if true {
                            if (true || true) {
                                for fi15 := 0; fi15 < 17; fi15++ {
                                    var v16 uint = 65535
                                    var v17 uint = v16
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    switch ("hello" + v7) {
                                    case "hello":
                                        for fi19 := 0; fi19 < 6; fi19++ {
                                            v2 -= v2
                                            _ = fi19
                                        }
                                    case "a":
                                        if false {
                                            v2 -= v2
                                            if len(v3) > 0 {
                                                v3[0] = 42
                                            }
                                        }
                                    }
                                    switch (fi18 | fi18) {
                                    case -2147483648:
                                        if (!(!true)) {
                                            if len(v3) > 0 {
                                                v3[0] = fi18
                                            }
                                        } else {
                                            v2 -= v2
                                        }
                                    case 4294967295:
                                        var v20 uint64 = v1
                                        _ = v20
                                    default:
                                        if false {
                                            switch ("\n\t" + v8) {
                                            case "foo bar":
                                                fmt.Println(v7)
                                            case "a":
                                                for fi21 := 0; fi21 < 20; fi21++ {
                                                    var v22 uint8 = 24
                                                    v23 := GenS1{F0: 109, F1: -1, F2: 2147483647}
                                                    _ = fi21
                                                    _ = v22
                                                    _ = v23
                                                }
                                            default:
                                                if false {
                                                    if (255 >= 128) {
                                                        v24 := GenS3{F0: -2147483648}
                                                        _ = v24
                                                    }
                                                }
                                            }
                                            v7 = (v8 + "hello")
                                            v25 := GenS2{F0: -2147483648}
                                            _ = v25
                                        }
                                    }
                                    _ = fi18
                                }
                                fmt.Println(v10, v8, v7)
                            } else {
                                var v26 uint32 = 852
                                v8 = ("" + "")
                                _ = v26
                            }
                            var v27 int = -2147483648
                            for fi28 := 0; fi28 < 20; fi28++ {
                                var v29 int64 = (-(v2))
                                _ = fi28
                                _ = v29
                            }
                            _ = v27
                        } else {
                            var v30 uint = 42
                            v30 = (v30 & (v30 & (v30 >> ((v30) & 63))))
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 13; fi31++ {
                            for fi32 := 0; fi32 < 1; fi32++ {
                                var v33 uint = 0
                                fmt.Println(fi32, v8, v8)
                                switch (v8 + "a") {
                                case "hello":
                                    if (uint8(v33) != 42) {
                                        var v34 float64 = 0.0
                                        var v35 float64 = 1e10
                                        var v36 uint = v33
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    } else {
                                        fmt.Println(v33, fi32, fi32)
                                    }
                                case "a":
                                    var v37 float32 = 0.1
                                    _ = v37
                                }
                                _ = fi32
                                _ = v33
                            }
                            var v38 float32 = 3.14159
                            _ = fi31
                            _ = v38
                        }
                    } else {
                        fmt.Println(v8, v7, v2)
                        v39 := GenS3{F0: -1}
                        _ = v39
                    }
                    v1 -= v1
                } else {
                    var v40 uint = 2
                    _ = v40
                }
            case "hello":
                if ((128 != 207) && true) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v3 = append(v3, 402)
                    v1 += 0
                }
            default:
                fmt.Println(v8)
            }
        }
        for fi41 := 0; fi41 < 15; fi41++ {
            for fi42 := 0; fi42 < 11; fi42++ {
                var v43 int16 = -1
                switch fi41 {
                case 127:
                    v43 = (v43 * v43)
                case -411:
                    var v44 bool = false
                    _ = v44
                case -613:
                    var v45 string = "a"
                    _ = v45
                default:
                    if (!(287 == 32767)) {
                        var v46 int16 = (v43 % ((v43) | 1))
                        _ = v46
                    }
                }
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
        _ = v10
    }
    var v47 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v7
    _ = v8
}
