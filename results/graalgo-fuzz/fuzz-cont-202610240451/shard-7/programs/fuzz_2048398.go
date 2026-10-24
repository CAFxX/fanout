package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint64; F1 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 int32 = 1
    v4 := []string{("" + (((("\n\t" + "\n\t") + "foo bar") + "hello") + (("a" + "foo bar") + "\n\t"))), "hello", ""}
    var v5 string = ""
    if len(v4) > 0 {
        v4[0] = v5
    }
    _ = len(v2)
    fmt.Println(v3, v3)
    v3 += (v3 & (v3 - (v3 / ((42) | 1))))
    v3 -= (^(v3))
    v5 = ("foo bar" + v5)
    switch v5 {
    case "foo bar":
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            switch "hello" {
            case "foo bar":
                v5 = ""
            case "a":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "":
                v8 := GenS1{F0: (v3 & v3), F1: -4, F2: 18446744073709551615}
                _ = v8
            default:
                for fi9 := 0; fi9 < 5; fi9++ {
                    v2[32767] = ("hello" + (v5 + v5))
                    if true {
                        fmt.Println(v5, fi9, fi9)
                        delete(v2, (^(fi9)))
                        if true {
                            v3 = v3
                            var v10 int16 = 162
                            if true {
                                switch ("a" + v5) {
                                case "":
                                    v11 := GenS2{F0: 233, F1: 74}
                                    _ = v11
                                case "foo bar":
                                    if true {
                                        var v12 float32 = (3.14159 - float32(0.0))
                                        _ = v12
                                    } else {
                                        switch fi7 {
                                        case 748:
                                            v10 += (-(v10))
                                        }
                                    }
                                    v13 := GenS3{F0: 255, F1: 1}
                                    if false {
                                        v14 := GenS4{F0: 0, F1: v10}
                                        _ = v14
                                    } else {
                                        if true {
                                            switch v5 {
                                            case "a":
                                                if true {
                                                    if false {
                                                        var v15 uint32 = 65535
                                                        v10 += int16(fi7)
                                                        fmt.Println(v5)
                                                        _ = v15
                                                    }
                                                }
                                            case "":
                                                var v16 int64 = 1
                                                _ = v16
                                            }
                                        } else {
                                            if true {
                                                var v17 int16 = v10
                                                _ = v17
                                            } else {
                                                var v18 uint = 70
                                                if false {
                                                    var v19 int32 = (v3 >> ((v18) & 63))
                                                    _ = v19
                                                } else {
                                                    if len(v4) > 0 {
                                                        v4[0] = v5
                                                    }
                                                    v1[fi7] = ""
                                                }
                                                _ = v18
                                            }
                                            var v20 uint8 = 1
                                            v3 = -504
                                            _ = v20
                                        }
                                    }
                                    _ = v13
                                default:
                                    var v21 int64 = 1
                                    _ = v21
                                }
                                var v22 int16 = v10
                                _ = v22
                            }
                            _ = v10
                        } else {
                            v3 = v3
                            v23 := GenS3{F0: 1, F1: v3}
                            var v24 uint = 1
                            v25 := GenS3{F0: 42, F1: -1}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        v26 := GenS2{F0: 128, F1: 2}
                        _ = v26
                    } else {
                        fmt.Println(v5, fi9, v5)
                        if false {
                            v5 = (v5 + v5)
                        } else {
                            v3 = -2147483648
                        }
                        v3 += v3
                    }
                    _ = fi9
                }
            }
            _ = fi7
        }
        fmt.Println(v3, v5)
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
