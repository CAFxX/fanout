package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint32 = 375
    var v2 float64 = 1e10
    var v3 int16 = 298
    var v4 int64 = 1
    v3 = (v3 * (v3 ^ v3))
    v5 := []int{-308}
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint = 65535
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        if len(v5) > 0 {
            v5[0] = 0
        }
        var v9 int32 = 954
        if (3.14159 >= (-2.25 + (-2.25 + 0.1))) {
            var v10 uint8 = 237
            var v11 float32 = 3.14159
            var v12 int8 = 1
            _ = v10
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: 1, F1: -1}
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    if false {
        switch (("a" + ("foo bar" + "")) + "hello") {
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = 128
            }
        case "hello":
            v14 := GenS1{F0: 42, F1: 1}
            _ = v14
        case "foo bar":
            var v15 bool = false
            _ = v15
        default:
            if (65535 <= 1) {
                v1 = (v1 * 42)
            } else {
                var v16 int = -9223372036854775808
                if true {
                    var v17 int = v16
                    switch v17 {
                    case 343:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v3 -= (-(v3))
                            _ = fi18
                        }
                    default:
                        for fi19 := 0; fi19 < 8; fi19++ {
                            fmt.Println(v16, v3, v2)
                            _ = fi19
                        }
                    }
                    if ((255 <= 1) || (141 < 2)) {
                        for fi20 := 0; fi20 < 8; fi20++ {
                            switch "\n\t" {
                            case "foo bar":
                                switch "" {
                                case "hello":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case "\n\t":
                                    v17 = (v17 >> ((586) & 63))
                                case "":
                                    v16 -= fi20
                                }
                            case "hello":
                                if (!(!(true || ((v17 <= (^(v17))) && ((^(v4)) != (v4 ^ (v4 >> ((122) & 63)))))))) {
                                    v21 := GenS2{F0: v1, F1: 42, F2: 0}
                                    _ = v21
                                } else {
                                    v1 = v1
                                    v22 := GenS3{F0: 0, F1: v3}
                                    _ = v22
                                }
                            }
                            _ = fi20
                        }
                    } else {
                        if (false || true) {
                            v23 := GenS4{F0: 549}
                            v3 -= v3
                            _ = v23
                        } else {
                            var v24 int16 = (-(v3))
                            for fi25 := 0; fi25 < 19; fi25++ {
                                switch "a" {
                                case "":
                                    if true {
                                        var v26 int8 = 1
                                        _ = v26
                                    }
                                case "a":
                                    fmt.Println(v17, v17, v17)
                                case "\n\t":
                                    switch "foo bar" {
                                    case "foo bar":
                                        if len(v5) > 0 {
                                            v5[0] = v17
                                        }
                                    case "\n\t":
                                        v27 := GenS5{F0: v24, F1: (v1 * v1)}
                                        _ = v27
                                    }
                                default:
                                    _ = len(v5)
                                }
                                switch fi25 {
                                case 974:
                                    var v28 string = ("" + "a")
                                    _ = v28
                                }
                                _ = fi25
                            }
                            _ = v24
                        }
                    }
                    _ = v17
                }
                v1 = (^(v1))
                _ = v16
            }
        }
    } else {
        v1 = v1
        switch (("\n\t" + "hello") + "foo bar") {
        case "hello":
            v2 += float64(1.5)
        case "a":
            for fi29 := 0; fi29 < 1; fi29++ {
                v1 += v1
                v3 += -1
                _ = fi29
            }
        default:
            v3 = v3
        }
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        switch ("a" + ("a" + "foo bar")) {
        case "\n\t":
            v31 := GenS4{F0: v4}
            _ = v31
        default:
            var v32 float64 = v2
            _ = v32
        }
        _ = fi30
    }
    var v33 uint32 = v1
    fmt.Println(v33)
    v2 += 1.5
    fmt.Println(v33, v1, v4)
    var v34 int64 = v4
    var v35 uint8 = 181
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
