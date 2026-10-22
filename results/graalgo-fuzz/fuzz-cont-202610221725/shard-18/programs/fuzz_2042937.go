package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int; F1 int32; F2 int32 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint8 = 0
    v2 := []int{}
    var v3 int8 = -128
    var v4 uint8 = v1
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v5 uint32 = 2147483648
        v6 := GenS1{F0: uint8(v5)}
        switch "" {
        case "hello":
            v4 = (^(v4))
        case "":
            if (v3 > v3) {
                v1 = 255
                v7 := GenS1{F0: (v4 ^ 2)}
                v3 -= (v3 | -1)
                _ = v7
            }
        default:
            v8 := GenS1{F0: v4}
            _ = v8
        }
        _ = v5
        _ = v6
    }
    v9 := GenS1{F0: v4}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 := GenS2{F0: 18446744073709551615, F1: 1}
    v3 = (v3 / (((v3 ^ v3)) | 1))
    if len(v2) > 0 {
        v2[0] = -240
    }
    var v11 bool = (3.14159 <= -2.25)
    v3 = 42
    v2 = append(v2, 128)
    if len(v2) > 0 {
        v2[0] = 0
    }
    if v11 {
        v3 = (v3 + v3)
        v12 := GenS2{F0: 18446744073709551615, F1: 0}
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: 1, F1: 0}
            _ = fi13
            _ = v14
        }
        _ = v12
    } else {
        switch 32768 {
        case -51:
            v1 += (^(v4))
        case 0:
            v3 = v3
        }
        if len(v2) > 0 {
            v2[0] = 128
        }
        _ = len(v2)
    }
    var v15 bool = (((-2.25 * (0.1 + 3.14159)) <= -2.25) || false)
    var v16 bool = v15
    switch 0 {
    case 2147483647:
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                switch "" {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = fi18
                    }
                case "\n\t":
                    fmt.Println(v11)
                case "a":
                    if v11 {
                        _ = len(v2)
                        var v19 int64 = -9223372036854775808
                        _ = v19
                    }
                }
                _ = fi18
            }
            _ = fi17
        }
    default:
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int64 = 1
            v4 = v4
            switch fi20 {
            case 127:
                if (1000 != 921) {
                    var v22 int64 = (^(v21))
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 float64 = -2.25
                        v11 = false
                        v24 -= (0.1 + 0.1)
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
            case 128:
                var v25 bool = true
                _ = v25
            default:
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := GenS2{F0: 541, F1: (v21 * int64(v3))}
                    if len(v2) > 0 {
                        v2[0] = (fi26 >> ((263) & 63))
                    }
                    if (1 >= 1) {
                        if v15 {
                            if true {
                                if true {
                                    v3 -= (-(v3))
                                    v1 = (v4 + v4)
                                }
                            }
                            var v28 string = ("\n\t" + "hello")
                            _ = v28
                        } else {
                            v29 := GenS3{F0: -1, F1: 0, F2: -2147483648}
                            if (!(861 < 2147483647)) {
                                var v30 int32 = -1
                                _ = v30
                            } else {
                                fmt.Println(v3)
                            }
                            _ = v29
                        }
                        v31 := GenS4{F0: 2147483647}
                        _ = v31
                    }
                    if v16 {
                        var v32 uint8 = v1
                        v32 = (v4 >> ((948) & 63))
                        switch (("\n\t" + "hello") + "\n\t") {
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = -2147483648
                            }
                        case "\n\t":
                            if (("hello" + "\n\t") == "hello") {
                                v33 := GenS4{F0: 314}
                                if (("hello" + ("foo bar" + "foo bar")) >= "a") {
                                    var v34 float32 = 1.5
                                    v1 = (v1 << ((18446744073709551615) & 63))
                                    if v11 {
                                        v16 = v16
                                    }
                                    _ = v34
                                } else {
                                    v32 = 210
                                    fmt.Println(v1)
                                }
                                v11 = (uint(v1) != 454)
                                _ = v33
                            }
                        case "":
                            fmt.Println(v3)
                        default:
                            switch (fi20 >> ((uint(v1)) & 63)) {
                            case -9223372036854775808:
                                fmt.Println(v11, v16)
                                fmt.Println(v1, v1)
                            default:
                                switch -1 {
                                case 127:
                                    var v35 uint = 2
                                    _ = v35
                                case -219:
                                    v36 := GenS4{F0: 0}
                                    _ = v36
                                default:
                                    fmt.Println(v11, v32, fi26)
                                }
                            }
                        }
                        _ = v32
                    } else {
                        _ = len(v2)
                        v21 = -914
                    }
                    v4 += v1
                    v21 -= 0
                    if len(v2) > 0 {
                        v2[0] = fi26
                    }
                    _ = fi26
                    _ = v27
                }
            }
            _ = fi20
            _ = v21
        }
    }
    var v37 int64 = -683
    for fi38 := 0; fi38 < 3; fi38++ {
        switch fi38 {
        case 32767:
            var v39 int = fi38
            _ = v39
        default:
            var v40 string = (("a" + "a") + (("hello" + "foo bar") + "\n\t"))
            _ = v40
        }
        var v41 int = fi38
        var v42 string = ""
        _ = fi38
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v9
}
