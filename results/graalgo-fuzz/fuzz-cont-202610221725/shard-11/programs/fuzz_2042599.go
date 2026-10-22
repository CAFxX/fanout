package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
type GenS3 struct { F0 int; F1 uint32; F2 int }

func main() {
    var v1 uint16 = 165
    var v2 int32 = 42
    v3 := []uint32{4294967295, 42, 660}
    var v4 int8 = -46
    _ = len(v3)
    v3 = append(v3, 42)
    var v5 float32 = -2.25
    _ = len(v3)
    var v6 int32 = (v2 - int32(v1))
    v4 -= -55
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                if true {
                    if true {
                        var v9 int32 = 469
                        v10 := GenS1{F0: 2147483647, F1: 456}
                        _ = v10
                        _ = v9
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                    }
                    var v11 float32 = (3.14159 / float32((1e10 - 1.5)))
                    if (false && (("hello" + "") < (("a" + "\n\t") + "hello"))) {
                        if len(v3) > 0 {
                            v3[0] = 877
                        }
                        v12 := GenS1{F0: 407, F1: 399}
                        _ = v12
                    }
                    var v13 int64 = 9223372036854775807
                    v14 := GenS1{F0: 780, F1: 921}
                    _ = v11
                    _ = v13
                    _ = v14
                } else {
                    if false {
                        if (!false) {
                            switch fi7 {
                            case 42:
                                _ = len(v3)
                            default:
                                var v15 uint32 = uint32(v1)
                                _ = v15
                            }
                            if ((true && true) && (!(2 >= 255))) {
                                var v16 int64 = -131
                                v17 := GenS1{F0: 0, F1: 72}
                                _ = v16
                                _ = v17
                            } else {
                                switch fi8 {
                                case -552:
                                    var v18 uint64 = 681
                                    v19 := GenS2{F0: 42, F1: 281, F2: -59}
                                    v2 = v6
                                    _ = v18
                                    _ = v19
                                case 138:
                                    switch ((("" + "a") + "hello") + ("\n\t" + "\n\t")) {
                                    case "":
                                        fmt.Println(v6, v6, fi8)
                                    case "\n\t":
                                        switch "a" {
                                        case "a":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "foo bar":
                                            v3 = append(v3, 402)
                                        }
                                    case "hello":
                                        var v20 uint8 = 110
                                        _ = v20
                                    default:
                                        v4 -= 0
                                    }
                                default:
                                    switch ("\n\t" + ("a" + "\n\t")) {
                                    case "a":
                                        fmt.Println(v2, v6, fi7)
                                    case "foo bar":
                                        v21 := GenS3{F0: 2147483647, F1: 913, F2: (fi8 & fi8)}
                                        _ = v21
                                    case "hello":
                                        _ = len(v3)
                                        v6 = (v2 + (v2 | v6))
                                    default:
                                        v3 = append(v3, 42)
                                    }
                                }
                            }
                            fmt.Println(v6)
                        }
                    }
                }
                _ = fi8
            }
            var v22 float32 = (v5 * v5)
            v6 -= -2147483648
            _ = fi7
            _ = v22
        }
        fmt.Println(v1, v5, v6)
        _ = len(v3)
    } else {
        var v23 uint = 935
        v24 := GenS3{F0: 128, F1: 4294967295, F2: 32767}
        var v25 uint32 = 2
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println(v4, v1, v5)
    v1 = v1
    _ = len(v3)
    v2 = v6
    v6 = v2
    var v26 uint32 = 728
    if ("hello" >= "foo bar") {
        if true {
            switch "foo bar" {
            case "":
                var v27 string = (("foo bar" + "\n\t") + "")
                _ = v27
            default:
                if (1000 < 32767) {
                    fmt.Println(v5, v5, v5)
                } else {
                    var v28 float64 = float64(v5)
                    var v29 int8 = (-(v4))
                    fmt.Println(v2, v2, v4)
                    _ = v28
                    _ = v29
                }
            }
        } else {
            if true {
                switch "\n\t" {
                case "foo bar":
                    for fi30 := 0; fi30 < 2; fi30++ {
                        var v31 int16 = -404
                        switch (fi30 << ((4294967295) & 63)) {
                        case 127:
                            v1 += (^(v1))
                        }
                        v2 = (v6 / ((-1) | 1))
                        _ = fi30
                        _ = v31
                    }
                case "a":
                    var v32 uint = 65535
                    _ = v32
                default:
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch "a" {
                        case "hello":
                            var v33 bool = (false || ((!true) || false))
                            _ = v33
                        case "foo bar":
                            for fi34 := 0; fi34 < 11; fi34++ {
                                if false {
                                    v5 += (float32(0.1) - -2.25)
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = (v26 & (v26 % ((622) | 1)))
                                    }
                                    var v35 uint64 = 18446744073709551615
                                    _ = v35
                                }
                                _ = fi34
                            }
                        }
                    }
                }
                v3 = append(v3, 42)
                var v36 uint64 = 2
                _ = v36
            } else {
                var v37 uint64 = 700
                if false {
                    var v38 uint32 = (v26 << ((548) & 63))
                    var v39 int64 = 405
                    var v40 uint8 = 0
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = v37
            }
            var v41 uint = 18446744073709551615
            _ = v41
        }
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "hello" {
            case "a":
                v2 += v6
            case "hello":
                if false {
                    for fi42 := 0; fi42 < 5; fi42++ {
                        switch fi42 {
                        case 127:
                            v43 := GenS3{F0: -1, F1: 610, F2: (fi42 | fi42)}
                            _ = v43
                        case -664:
                            if false {
                                if (true || true) {
                                    var v44 int8 = v4
                                    _ = v44
                                } else {
                                    var v45 uint = 42
                                    _ = v45
                                }
                                fmt.Println(v4, v6)
                            } else {
                                fmt.Println(v26, v1, v26)
                                v46 := GenS3{F0: int(v2), F1: (^(v26)), F2: fi42}
                                _ = v46
                            }
                        default:
                            var v47 uint = 2
                            _ = v47
                        }
                        switch "foo bar" {
                        case "":
                            var v48 int16 = 42
                            _ = v48
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = v26
                            }
                        case "hello":
                            var v49 int8 = -128
                            _ = v49
                        default:
                            v1 += 32767
                        }
                        _ = fi42
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    fmt.Println(v6, v26)
                }
            default:
                var v50 bool = ((!(2 <= v26)) || true)
                _ = v50
            }
        }
        switch -819 {
        case 0:
            fmt.Println(v5)
        }
    }
    v2 -= (v6 + (v2 & 1))
    var v51 string = ("" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
