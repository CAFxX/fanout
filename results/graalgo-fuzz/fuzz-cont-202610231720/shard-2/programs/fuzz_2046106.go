package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint64 }

func main() {
    var v1 int8 = 84
    var v2 int64 = -822
    v3 := make(map[string]uint32)
    v4 := []uint32{0, 2147483647, 311, 297}
    var v5 int32 = int32(v1)
    var v6 int32 = v5
    if false {
        v6 += (v6 ^ (v6 * v5))
        fmt.Println(v6, v6)
        if true {
            v6 = 453
            var v7 int64 = 42
            v4 = append(v4, 2)
            _ = v7
        }
    }
    v2 -= (v2 & (v2 & v2))
    var v8 uint = 0
    if true {
        v9 := GenS1{F0: "\n\t", F1: 998}
        var v10 bool = true
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            switch 0 {
            case -716:
                _ = len(v4)
            case 127:
                switch ("" + "\n\t") {
                case "\n\t":
                    v6 = v6
                default:
                    if true {
                        fmt.Println(fi11)
                        var v12 int32 = (v6 - (-(v6)))
                        var v13 uint64 = 1
                        _ = v12
                        _ = v13
                    }
                }
            case -273:
                switch fi11 {
                case -1:
                    if true {
                        var v14 int16 = 42
                        if false {
                            v15 := GenS1{F0: "foo bar", F1: v2}
                            _ = v15
                        }
                        _ = v14
                    } else {
                        switch ("" + "\n\t") {
                        case "":
                            v1 = (v1 + v1)
                        case "a":
                            switch "" {
                            case "a":
                                if (256 == 204) {
                                    var v16 uint32 = 363
                                    v2 += v2
                                    var v17 bool = false
                                    _ = v16
                                    _ = v17
                                } else {
                                    _ = len(v3)
                                    v1 -= v1
                                    var v18 uint8 = 88
                                    _ = v18
                                }
                            default:
                                v8 -= (v8 | (v8 * v8))
                            }
                        default:
                            var v19 uint16 = uint16(v8)
                            _ = v19
                        }
                    }
                default:
                    if false {
                        v2 = -62
                        var v20 uint16 = 255
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v20
                    } else {
                        if true {
                            for fi21 := 0; fi21 < 13; fi21++ {
                                v22 := GenS1{F0: "", F1: 42}
                                _ = fi21
                                _ = v22
                            }
                        }
                        for fi23 := 0; fi23 < 16; fi23++ {
                            v8 = (v8 ^ v8)
                            fmt.Println(v8, v5, v8)
                            var v24 string = "\n\t"
                            _ = fi23
                            _ = v24
                        }
                    }
                }
            default:
                switch (("foo bar" + ((("foo bar" + "") + ("" + (("a" + ("hello" + "\n\t")) + "a"))) + "foo bar")) + (("foo bar" + ("\n\t" + "")) + ("foo bar" + "\n\t"))) {
                case "":
                    if true {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v1 = int8(v8)
                            _ = fi25
                        }
                    }
                case "\n\t":
                    var v26 uint8 = 148
                    _ = v26
                case "foo bar":
                    v2 = (v2 & (v2 | 1))
                }
            }
            if len(v4) > 0 {
                v4[0] = 42
            }
            v27 := GenS2{F0: v1, F1: int8(v6), F2: 2}
            _ = fi11
            _ = v27
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 int = 32768
            v30 := GenS3{F0: 272, F1: v1, F2: 18446744073709551615}
            _ = fi28
            _ = v29
            _ = v30
        }
        v6 += (v5 / ((v5) | 1))
    }
    v6 = v6
    fmt.Println(v6, v2, v2)
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    var v31 uint = (v8 >> (((v8 >> ((v8) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
