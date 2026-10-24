package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int; F1 int32; F2 uint8 }

func main() {
    var v1 uint = 0
    v2 := make(map[string]string)
    var v3 int = 128
    v4 := []uint32{438}
    var v5 bool = (true && true)
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint16 = 1
        switch ("hello" + "hello") {
        case "hello":
            var v8 string = ("a" + (("hello" + ("" + (((("a" + "\n\t") + "") + ("foo bar" + "foo bar")) + "foo bar"))) + (("" + "") + ("\n\t" + ("\n\t" + ("foo bar" + "a"))))))
            _ = v8
        case "a":
            var v9 uint16 = (v7 | (v7 % (((v7 - 2)) | 1)))
            _ = v9
        case "":
            var v10 string = "\n\t"
            _ = v10
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v11 := GenS1{F0: -1}
        _ = fi6
        _ = v11
        _ = v7
    }
    v3 -= (v3 << ((v1) & 63))
    var v12 int64 = 1
    if v5 {
        if (!v5) {
            switch "a" {
            case "hello":
                var v13 int8 = 45
                _ = v13
            case "foo bar":
                fmt.Println(v1, v3)
            case "\n\t":
                for fi14 := 0; fi14 < 7; fi14++ {
                    v1 = 65535
                    _ = fi14
                }
            }
            if v5 {
                switch "a" {
                case "foo bar":
                    v12 = (v12 % ((v12) | 1))
                case "":
                    _ = len(v2)
                default:
                    switch (-(v3)) {
                    case 1000:
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v5 = v5
                            _ = fi15
                        }
                    case 2147483647:
                        if (v5 || (!(42 == 127))) {
                            v1 += 4294967295
                            if (32767 >= 1000) {
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                                fmt.Println(v5, v3)
                            } else {
                                var v16 uint = (v1 % ((v1) | 1))
                                fmt.Println(v16, v3, v12)
                                v5 = false
                                _ = v16
                            }
                            fmt.Println(v3)
                        }
                    case 665:
                        var v17 int64 = v12
                        _ = v17
                    default:
                        if (((0 != -48) && true) && v5) {
                            switch 626 {
                            case 127:
                                v1 = (v1 % (((^(v1))) | 1))
                            case 2147483647:
                                var v18 uint8 = 103
                                _ = v18
                            case -383:
                                fmt.Println(v12)
                            }
                            var v19 uint32 = 352
                            v5 = (618 <= 32767)
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 14; fi20++ {
                                v5 = v5
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    fmt.Println(v12, v3)
                                    v1 -= 689
                                    _ = fi21
                                }
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    v3 = (^(fi22))
                                    if ("a" >= "hello") {
                                        if v5 {
                                            switch -1 {
                                            case 219:
                                                var v23 int8 = 1
                                                _ = v23
                                            case 1:
                                                fmt.Println(v3)
                                            }
                                        } else {
                                            var v24 uint16 = 2
                                            fmt.Println(fi22, v24, fi20)
                                            v5 = v5
                                            _ = v24
                                        }
                                        fmt.Println(fi22, v5)
                                        switch v3 {
                                        case 0:
                                            var v25 int32 = 1
                                            _ = v25
                                        case 4294967295:
                                            if false {
                                                switch "a" {
                                                case "hello":
                                                    var v26 int64 = 110
                                                    _ = v26
                                                case "\n\t":
                                                    v27 := GenS1{F0: -574}
                                                    _ = v27
                                                }
                                            }
                                        case 9223372036854775807:
                                            v12 = v12
                                        default:
                                            v28 := GenS2{F0: 554, F1: 100}
                                            _ = v28
                                        }
                                    }
                                    _ = fi22
                                }
                                _ = fi20
                            }
                            for fi29 := 0; fi29 < 5; fi29++ {
                                switch (-(fi29)) {
                                case -2147483648:
                                    v1 += (^(v1))
                                }
                                var v30 uint64 = 440
                                _ = len(v2)
                                _ = fi29
                                _ = v30
                            }
                        }
                    }
                }
                v31 := GenS3{F0: 537, F1: 0}
                _ = v31
            }
        } else {
            v32 := GenS4{F0: (^(v3)), F1: -209, F2: 3}
            var v33 int8 = 127
            _ = len(v4)
            _ = v32
            _ = v33
        }
        _ = len(v4)
    } else {
        if false {
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 int8 = 1
                switch ((((("" + "") + (("hello" + (("\n\t" + "\n\t") + ("\n\t" + ""))) + "hello")) + "") + (("a" + "a") + "a")) + "") {
                case "a":
                    var v36 int8 = v35
                    _ = v36
                case "hello":
                    var v37 uint8 = 255
                    _ = v37
                }
                var v38 float64 = 1.5
                _ = fi34
                _ = v35
                _ = v38
            }
            var v39 int64 = 9223372036854775807
            v12 = (v12 * v12)
            _ = v39
        }
        var v40 int32 = 0
        _ = v40
    }
    v1 -= (v1 >> ((0) & 63))
    if (9223372036854775807 != -146) {
        var v41 bool = (v12 >= (-(v12)))
        _ = v41
    } else {
        var v42 float64 = -2.25
        _ = v42
    }
    var v43 int = 4294967295
    var v44 int8 = 42
    v3 = (v3 % (((v3 * v3)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
