package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int32; F1 int32; F2 int32 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint; F1 uint32; F2 int64 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 int8; F1 uint64; F2 int64 }
func genFn0() uint8 {
    var v6 int16 = 1000
    v6 = (v6 + -49)
    _ = v6
    return 67
}
func genFn1(p0 int16, p1 bool) uint16 {
    v7 := GenS1{F0: 1, F1: 42, F2: -9223372036854775808}
    v8 := []int{-9223372036854775808}
    var v9 int = 2147483647
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 111
}

func main() {
    var v1 int = 845
    var v2 int16 = 32767
    var v3 uint32 = 2
    v4 := []uint32{v3}
    var v5 int32 = 778
    var v10 bool = (false && (0 >= 234))
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v3, v2)
        var v12 int64 = 693
        _ = fi11
        _ = v12
    }
    switch ((("\n\t" + "a") + "hello") + ("a" + "a")) {
    case "hello":
        var v13 int64 = 9223372036854775807
        _ = v13
    case "foo bar":
        switch 1 {
        case 481:
            if len(v4) > 0 {
                v4[0] = v3
            }
        case 4294967295:
            if ("" >= "hello") {
                fmt.Println(v5, v1)
                for fi14 := 0; fi14 < 11; fi14++ {
                    if false {
                        if (float64(0.1) != float64(-2.25)) {
                            var v15 float32 = 0.0
                            fmt.Println(v15, v15, v1)
                            _ = v15
                        }
                        if (v10 && v10) {
                            v16 := GenS2{F0: fi14, F1: v5}
                            var v17 int64 = 42
                            var v18 float32 = 0.0
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    } else {
                        switch (("\n\t" + "hello") + "") {
                        case "foo bar":
                            var v19 string = (("" + (("" + "") + "foo bar")) + "\n\t")
                            _ = v19
                        }
                        if len(v4) > 0 {
                            v4[0] = (v3 | 2147483648)
                        }
                    }
                    v10 = false
                    _ = fi14
                }
                if len(v4) > 0 {
                    v4[0] = (v3 << ((18446744073709551615) & 63))
                }
            }
        default:
            switch v1 {
            case 128:
                v10 = (v10 || v10)
            case 2147483647:
                var v20 bool = (0 <= -1)
                _ = v20
            default:
                switch (v1 >> ((65535) & 63)) {
                case -531:
                    if (0.1 < (((1e10 / -2.25) * 1e10) - 1.5)) {
                        v21 := GenS3{F0: v5, F1: -715, F2: (v5 + v5)}
                        _ = v21
                    }
                case -9223372036854775808:
                    v22 := GenS4{F0: -1, F1: v5}
                    _ = v22
                case 1:
                    v10 = (42 == 587)
                default:
                    if (42 >= 38) {
                        var v23 string = ("a" + "foo bar")
                        _ = v23
                    }
                }
            }
        }
    default:
        switch (^(v1)) {
        case -2147483648:
            v10 = (false || v10)
        default:
            var v24 bool = false
            _ = v24
        }
    }
    var v25 int32 = 42
    var v26 int16 = (-(v2))
    v27 := GenS2{F0: (v1 * 32768), F1: (v25 & (v5 * v5))}
    switch (v1 << ((42) & 63)) {
    case 32767:
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS5{F0: 2, F1: 0, F2: -787}
            if len(v4) > 0 {
                v4[0] = v3
            }
            if (1 == 127) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                var v30 bool = (v25 == -2147483648)
                if v30 {
                    switch "hello" {
                    case "\n\t":
                        for fi31 := 0; fi31 < 6; fi31++ {
                            v30 = (-139 < 1)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi31
                        }
                    case "foo bar":
                        v10 = (v2 == (v2 << ((1) & 63)))
                    case "a":
                        switch fi28 {
                        case 42:
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    fmt.Println(v3)
                                    if true {
                                        if false {
                                            v2 = int16(fi28)
                                            var v33 bool = false
                                            v4 = append(v4, (v3 & v3))
                                            var v34 uint8 = 240
                                            _ = v33
                                            _ = v34
                                        }
                                        var v35 uint = 255
                                        _ = v35
                                    }
                                    var v36 int = (v1 / (((^(fi28))) | 1))
                                    _ = fi32
                                    _ = v36
                                }
                            } else {
                                fmt.Println(fi28, v10)
                                fmt.Println(v3, v2, v26)
                            }
                        case 4294967295:
                            v4 = append(v4, 1)
                        }
                    }
                    v37 := GenS6{F0: "", F1: v3}
                    v2 = v26
                    _ = v37
                }
                _ = v30
            }
            _ = fi28
            _ = v29
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        if (!v10) {
            v39 := GenS3{F0: v5, F1: (^(v25)), F2: v25}
            _ = v39
        } else {
            v26 = (-(v26))
            v5 += v5
            fmt.Println(v1, v1, v26)
        }
        _ = len(v4)
        v40 := GenS4{F0: 42, F1: v5}
        _ = fi38
        _ = v40
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        v42 := GenS7{F0: -82, F1: 2, F2: -9223372036854775808}
        for fi43 := 0; fi43 < 20; fi43++ {
            _ = len(v4)
            _ = fi43
        }
        var v44 uint16 = 255
        _ = fi41
        _ = v42
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
