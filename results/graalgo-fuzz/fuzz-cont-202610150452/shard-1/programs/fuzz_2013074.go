package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool }

func main() {
    v1 := []uint32{1, 42, 4294967295}
    var v2 int64 = -376
    v3 := []uint32{1, 851, 2}
    v4 := []string{("hello" + "foo bar")}
    var v5 uint8 = 255
    var v6 uint = 4294967295
    v2 -= int64(v6)
    fmt.Println(v5, v6, v5)
    for fi7 := 0; fi7 < 11; fi7++ {
        if (!false) {
            switch "hello" {
            case "hello":
                var v8 int64 = (v2 << (((v6 * v6)) & 63))
                _ = v8
            default:
                v2 -= v2
            }
        } else {
            var v9 float32 = (-2.25 * float32((0.0 + (3.14159 * (0.1 - -2.25)))))
            _ = v9
        }
        switch ("" + ("hello" + "")) {
        case "a":
            if true {
                v6 = v6
                var v10 uint16 = 1
                if false {
                    if false {
                        v2 = -137
                    }
                }
                _ = v10
            } else {
                var v11 float32 = 3.14159
                fmt.Println(fi7, v2)
                v12 := GenS1{F0: 2, F1: -2147483648}
                _ = v11
                _ = v12
            }
        case "":
            v5 -= 128
        case "hello":
            var v13 float32 = -2.25
            _ = v13
        default:
            var v14 uint16 = 457
            _ = v14
        }
        fmt.Println(v5, v5)
        _ = fi7
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int64 = v2
        v17 := GenS1{F0: 908, F1: -893}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int16 = -8
    var v19 float64 = 1.5
    var v20 float32 = 1.5
    if (42 < -948) {
        var v21 int32 = int32(v18)
        _ = v21
    }
    if (!(1 >= 0)) {
        v2 += -22
        v5 = 2
    } else {
        v2 -= -418
        var v22 uint16 = 476
        v22 = (^(v22))
        _ = v22
    }
    var v23 int16 = (v18 * 42)
    switch 0 {
    case 917:
        fmt.Println(v6)
    case 106:
        v20 -= v20
    default:
        var v24 int64 = v2
        _ = v24
    }
    var v25 uint32 = 65535
    fmt.Println(v19, v5, v19)
    v26 := GenS2{F0: 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        if true {
            switch "" {
            case "hello":
                fmt.Println(v6, v25)
            case "a":
                v23 -= (v18 << ((v6) & 63))
            case "foo bar":
                v4 = append(v4, ("a" + ""))
            }
        }
        fmt.Println(v18, v20, v5)
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 uint8 = v5
            var v29 uint32 = v25
            v23 -= -309
            _ = fi27
            _ = v28
            _ = v29
        }
    } else {
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := GenS3{F0: 255, F1: (v25 << ((v6) & 63)), F2: -128}
            v32 := GenS2{F0: 42}
            if true {
                var v33 bool = ((^(v6)) <= v6)
                for fi34 := 0; fi34 < 2; fi34++ {
                    switch -2147483648 {
                    case 127:
                        if (!(!((18446744073709551615 < 2) && ((v18 ^ 0) == -401)))) {
                            v35 := GenS4{F0: 986, F1: 37, F2: -1}
                            _ = v35
                        }
                    case -1:
                        var v36 int8 = -93
                        _ = v36
                    case 32768:
                        v37 := GenS5{F0: 873}
                        _ = v37
                    }
                    _ = fi34
                }
                v6 -= v6
                _ = v33
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        for fi38 := 0; fi38 < 5; fi38++ {
            v25 = v25
            _ = fi38
        }
    }
    if (!(2147483647 >= 2147483647)) {
        var v39 int16 = (v23 % (((v23 % (((v18 % (((-(v18))) | 1))) | 1))) | 1))
        fmt.Println(v20)
        switch ((("a" + "a") + "") + ("a" + "hello")) {
        case "hello":
            v2 -= v2
        case "\n\t":
            var v40 uint8 = 1
            _ = v40
        case "foo bar":
            for fi41 := 0; fi41 < 11; fi41++ {
                switch (-(fi41)) {
                case -2147483648:
                    var v42 int = (fi41 & (^(fi41)))
                    _ = v42
                case 525:
                    if false {
                        for fi43 := 0; fi43 < 16; fi43++ {
                            var v44 int64 = v2
                            var v45 int16 = (v39 & v18)
                            if false {
                                fmt.Println(v19, v25, v18)
                                if (!((v5 % ((v5) | 1)) != v5)) {
                                    if len(v3) > 0 {
                                        v3[0] = (v25 & (v25 & (v25 | uint32(v5))))
                                    }
                                } else {
                                    switch (("" + ("\n\t" + "")) + "\n\t") {
                                    case "hello":
                                        v3 = append(v3, (^(v25)))
                                    }
                                }
                            } else {
                                var v46 uint16 = 55
                                _ = len(v3)
                                _ = v46
                            }
                            var v47 uint32 = (v25 | 65535)
                            if ((-(v44)) >= -9223372036854775808) {
                                fmt.Println(v20)
                                switch fi41 {
                                case 127:
                                    var v48 string = ("a" + ("a" + "foo bar"))
                                    _ = v48
                                case 4294967295:
                                    v49 := GenS6{F0: "foo bar"}
                                    _ = v49
                                case 223:
                                    var v50 int8 = -128
                                    _ = v50
                                default:
                                    if true {
                                        v51 := GenS3{F0: 255, F1: v47, F2: 0}
                                        _ = v51
                                    } else {
                                        if true {
                                            fmt.Println(v47, v6)
                                            var v52 uint8 = (v5 >> (((v6 ^ (v6 << ((v6) & 63)))) & 63))
                                            _ = v52
                                        } else {
                                            if false {
                                                v25 = (^(v25))
                                            } else {
                                                v25 = uint32(v5)
                                            }
                                        }
                                    }
                                }
                            } else {
                                v20 = v20
                            }
                            if len(v3) > 0 {
                                v3[0] = v25
                            }
                            _ = fi43
                            _ = v44
                            _ = v45
                            _ = v47
                        }
                    } else {
                        if true {
                            switch "a" {
                            case "\n\t":
                                if len(v3) > 0 {
                                    v3[0] = (v25 ^ v25)
                                }
                            case "hello":
                                if false {
                                    switch ("\n\t" + "\n\t") {
                                    case "hello":
                                        _ = len(v1)
                                    case "\n\t":
                                        var v53 int16 = (^(v23))
                                        _ = v53
                                    }
                                    fmt.Println(v20)
                                    fmt.Println(v2, v6)
                                }
                            }
                        }
                        for fi54 := 0; fi54 < 20; fi54++ {
                            var v55 int = fi54
                            v55 = 32768
                            _ = fi54
                            _ = v55
                        }
                        var v56 float32 = (v20 - v20)
                        _ = v56
                    }
                }
                v57 := GenS7{F0: true}
                var v58 string = "hello"
                _ = fi41
                _ = v57
                _ = v58
            }
        default:
            var v59 uint = 297
            _ = v59
        }
        _ = v39
    } else {
        if (v18 >= -65) {
            if true {
                var v60 int = -856
                var v61 uint64 = 2
                var v62 int32 = -1
                _ = v60
                _ = v61
                _ = v62
            }
        }
        for fi63 := 0; fi63 < 15; fi63++ {
            var v64 uint16 = uint16(v6)
            _ = fi63
            _ = v64
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
