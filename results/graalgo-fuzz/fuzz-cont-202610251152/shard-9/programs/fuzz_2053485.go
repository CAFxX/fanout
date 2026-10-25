package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int16 }

func main() {
    var v1 uint16 = 39
    var v2 int = -668
    var v3 int16 = -881
    var v4 int64 = 42
    v1 += (v1 / ((108) | 1))
    fmt.Println(v4, v4, v1)
    switch v2 {
    case 127:
        if false {
            for fi5 := 0; fi5 < 7; fi5++ {
                if true {
                    v6 := GenS1{F0: (-(v4))}
                    v4 = v4
                    var v7 uint32 = 0
                    _ = v6
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        fmt.Println(v1, v4)
                        _ = fi8
                    }
                    if (!(fi5 > v2)) {
                        fmt.Println(v1, v3, v1)
                        fmt.Println(fi5, fi5, v2)
                    }
                }
                _ = fi5
            }
            fmt.Println(v2, v2)
            if false {
                v9 := []int{(v2 << ((410) & 63)), (-(v2)), v2}
                _ = v9
            }
        }
    case 4294967295:
        switch "\n\t" {
        case "\n\t":
            v10 := GenS2{F0: 54, F1: (v1 >> ((329) & 63)), F2: ("foo bar" + "\n\t")}
            _ = v10
        case "hello":
            fmt.Println(v3, v1)
        case "":
            fmt.Println(v1)
        }
    case 9223372036854775807:
        v11 := GenS3{F0: (v4 + -286)}
        _ = v11
    default:
        v1 = v1
    }
    var v12 int = (-(v2))
    v13 := []int{}
    _ = len(v13)
    for fi14 := 0; fi14 < 8; fi14++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = fi14
    }
    var v15 uint32 = 0
    switch int(v1) {
    case 9223372036854775807:
        v13 = append(v13, (^(v2)))
    default:
        v2 -= (v12 ^ 32767)
    }
    v13 = append(v13, v2)
    switch "foo bar" {
    case "foo bar":
        var v16 int64 = v4
        _ = v16
    case "":
        if (v15 >= 483) {
            v17 := GenS3{F0: v4}
            if ((true || (int(v15) > -1)) || (v3 != -1)) {
                for fi18 := 0; fi18 < 16; fi18++ {
                    v19 := GenS3{F0: 0}
                    _ = fi18
                    _ = v19
                }
            } else {
                fmt.Println(v1)
                if (v1 != v1) {
                    fmt.Println(v4, v1)
                    switch ("foo bar" + ("hello" + "foo bar")) {
                    case "":
                        var v20 int32 = int32(v4)
                        _ = v20
                    case "\n\t":
                        fmt.Println(v2, v1)
                    case "foo bar":
                        if false {
                            v13 = append(v13, (^(v2)))
                            switch 127 {
                            case 42:
                                switch (-(v12)) {
                                case 128:
                                    var v21 int32 = -1
                                    _ = v21
                                default:
                                    var v22 string = "foo bar"
                                    _ = v22
                                }
                            case 32768:
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                            case 127:
                                for fi23 := 0; fi23 < 8; fi23++ {
                                    var v24 float32 = 3.14159
                                    switch (v2 & fi23) {
                                    case -2147483648:
                                        if len(v13) > 0 {
                                            v13[0] = 4294967295
                                        }
                                    case 127:
                                        var v25 uint16 = 32767
                                        _ = v25
                                    case 32768:
                                        var v26 bool = (v1 > 2)
                                        _ = v26
                                    default:
                                        fmt.Println(v2, v2)
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        } else {
                            var v27 int16 = (v3 << ((18446744073709551615) & 63))
                            v12 = v12
                            v28 := GenS2{F0: uint(v1), F1: (^(v1)), F2: "a"}
                            _ = v27
                            _ = v28
                        }
                    default:
                        switch 387 {
                        case 9223372036854775807:
                            for fi29 := 0; fi29 < 18; fi29++ {
                                v1 = (v1 / ((v1) | 1))
                                _ = fi29
                            }
                        case 32768:
                            v15 = 484
                        case 127:
                            v13 = append(v13, v12)
                        }
                    }
                }
                if false {
                    var v30 uint64 = 4294967295
                    var v31 int8 = 42
                    v15 += (v15 - v15)
                    _ = v30
                    _ = v31
                } else {
                    var v32 int64 = (v4 & (v4 * -800))
                    fmt.Println(v2, v3)
                    _ = v32
                }
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v17
        } else {
            if true {
                var v33 int = (v12 + v2)
                _ = v33
            } else {
                var v34 float64 = ((0.1 + float64((0.1 * 1.5))) / 1.5)
                var v35 int = (v12 >> ((42) & 63))
                _ = v34
                _ = v35
            }
            v36 := GenS4{F0: 6, F1: 42, F2: (-(v3))}
            _ = v36
        }
    default:
        _ = len(v13)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
