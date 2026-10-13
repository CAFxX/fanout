package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 int16; F1 uint64; F2 string }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]string)
    v3 := []uint32{889}
    var v4 uint8 = 1
    v5 := v2[65535]
    v1 += 9223372036854775807
    v5 = "\n\t"
    var v6 uint = 2
    switch "hello" {
    case "foo bar":
        if true {
            v2[4294967295] = v5
        }
    case "\n\t":
        v4 -= 127
    case "hello":
        if false {
            v4 = 42
            for fi7 := 0; fi7 < 5; fi7++ {
                for fi8 := 0; fi8 < 10; fi8++ {
                    if (!true) {
                        fmt.Println(v4, v1)
                        fmt.Println(fi8, v6)
                    } else {
                        v2[942] = v5
                        var v9 int = fi8
                        if (true || false) {
                            var v10 uint16 = 2
                            _ = v10
                        } else {
                            if true {
                                if (fi7 == v9) {
                                    if false {
                                        if false {
                                            v11 := GenS1{F0: (v9 + 127), F1: (v9 | v9)}
                                            if len(v3) > 0 {
                                                v3[0] = 477
                                            }
                                            _ = v11
                                        }
                                        var v12 uint32 = 0
                                        v5 = "a"
                                        _ = v12
                                    }
                                    fmt.Println(v1, v4)
                                }
                                v13 := GenS1{F0: -483, F1: (fi7 * (-(fi7)))}
                                v14 := GenS2{F0: 0, F1: v6}
                                _ = v13
                                _ = v14
                            } else {
                                if (!(!(false || false))) {
                                    var v15 string = "a"
                                    if (!(3.14159 < 1e10)) {
                                        var v16 uint = 4294967295
                                        fmt.Println(v6, v16, fi7)
                                        _ = v16
                                    }
                                    v17 := GenS2{F0: 127, F1: v6}
                                    _ = v15
                                    _ = v17
                                } else {
                                    fmt.Println(fi7, v1)
                                    var v18 float32 = (0.1 - 1e10)
                                    if len(v3) > 0 {
                                        v3[0] = 2
                                    }
                                    _ = v18
                                }
                                if (((("hello" + v5) + "hello") + v5) < v5) {
                                    fmt.Println(v9)
                                    if (2147483647 != 2) {
                                        v5 = ((v5 + v5) + v5)
                                        switch 233 {
                                        case 32768:
                                            v4 = v4
                                        case 2147483647:
                                            v19 := GenS2{F0: (v4 | v4), F1: v6}
                                            _ = v19
                                        }
                                    }
                                }
                            }
                            v6 = (v6 >> ((656) & 63))
                            var v20 int8 = -128
                            _ = v20
                        }
                        _ = v9
                    }
                    _ = fi8
                }
                var v21 float32 = 1e10
                _ = fi7
                _ = v21
            }
        }
    default:
        v4 += 2
    }
    fmt.Println(v4, v1)
    v22 := GenS1{F0: 0, F1: 1}
    var v23 int16 = 975
    v24 := GenS3{F0: int16(v6), F1: v6, F2: "hello"}
    switch 0 {
    case 904:
        fmt.Println(v1, v1)
    case 128:
        if (1 < v23) {
            if (42 < int(v1)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            v25 := GenS2{F0: v4, F1: v6}
            v2[969] = (v5 + "foo bar")
            var v26 int = 42
            _ = v25
            _ = v26
        }
    case 42:
        var v27 bool = (2147483648 >= 2147483647)
        _ = v27
    }
    switch (v5 + ("foo bar" + v5)) {
    case "foo bar":
        v4 += 128
    case "hello":
        var v28 uint = (^(v6))
        _ = v28
    case "":
        v2[186] = "\n\t"
    }
    v29 := GenS3{F0: 0, F1: v6, F2: ((v5 + v5) + ((v5 + "") + v5))}
    switch "hello" {
    case "\n\t":
        for fi30 := 0; fi30 < 17; fi30++ {
            if (255 < (v4 << ((v6) & 63))) {
                v31 := GenS4{F0: v23, F1: 2, F2: (((v5 + ("foo bar" + v5)) + "hello") + ("hello" + ""))}
                fmt.Println(v6)
                var v32 uint8 = (v4 ^ 42)
                _ = v31
                _ = v32
            }
            v4 = (^(v4))
            _ = fi30
        }
    case "":
        if (v4 > v4) {
            v4 -= (^(v4))
            v23 = v23
        }
    }
    v4 += (^(v4))
    switch 527 {
    case 34:
        v23 += 545
    case 393:
        if false {
            v2[2147483647] = v5
        }
    case 32768:
        if true {
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 uint64 = 92
                var v35 uint = (v6 << ((v6) & 63))
                v34 -= (v34 % ((uint64(v4)) | 1))
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    default:
        for fi36 := 0; fi36 < 7; fi36++ {
            v2[2] = v5
            _ = fi36
        }
    }
    v4 = (^(v4))
    v37 := GenS3{F0: v23, F1: v6, F2: (v5 + ("hello" + "\n\t"))}
    for fi38 := 0; fi38 < 10; fi38++ {
        v39 := v2[645]
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
