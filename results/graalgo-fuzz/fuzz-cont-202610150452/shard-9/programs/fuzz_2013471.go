package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    v1 := []string{(("\n\t" + ("a" + "hello")) + "foo bar"), (("\n\t" + "foo bar") + "\n\t"), "\n\t", "a"}
    var v2 uint64 = 501
    var v3 int16 = 0
    v4 := make(map[int]string)
    if (0 > uint16(v2)) {
        v5 := GenS1{F0: 705}
        fmt.Println(v2, v2)
        _ = v5
    } else {
        v2 -= (v2 & (v2 / ((v2) | 1)))
        fmt.Println(v3)
        v2 = v2
    }
    var v6 int32 = 2147483647
    v7 := GenS1{F0: 2}
    _ = len(v4)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v8 := GenS1{F0: 1}
        v9 := GenS1{F0: 2}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            v6 = v6
            var v11 float32 = (0.0 / (1e10 * ((float32(1e10) - 3.14159) + 0.0)))
            _ = fi10
            _ = v11
        }
        fmt.Println(v2)
    }
    fmt.Println(v3, v6, v3)
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint16 = 0
            switch "a" {
            case "hello":
                switch (fi13 % (((-(fi12))) | 1)) {
                case 9223372036854775807:
                    var v15 bool = (42 > -1)
                    v16 := GenS2{F0: 338, F1: -32768}
                    var v17 int8 = 1
                    delete(v4, (fi13 + fi13))
                    v17 = v17
                    var v18 bool = (v2 < (v2 << ((835) & 63)))
                    if true {
                        fmt.Println(v3)
                    } else {
                        fmt.Println(v18, v18, v6)
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        if v18 {
                            if (!(!false)) {
                                v17 += (^(v17))
                            } else {
                                v19 := GenS3{F0: 122}
                                v3 -= v3
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                _ = v19
                            }
                            if false {
                                switch "hello" {
                                case "":
                                    v2 = v2
                                case "hello":
                                    v14 -= (v14 ^ 32767)
                                }
                            } else {
                                if true {
                                    v17 = v17
                                    if v18 {
                                        _ = len(v4)
                                    } else {
                                        fmt.Println(v18, v15, v18)
                                    }
                                    if v15 {
                                        var v20 int16 = v3
                                        fmt.Println(v2, v17, v15)
                                        var v21 int64 = -459
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 uint = 2
                                    _ = len(v4)
                                    _ = v22
                                } else {
                                    switch 32767 {
                                    case 42:
                                        _ = len(v4)
                                    }
                                }
                                var v23 uint32 = 390
                                v24 := GenS1{F0: 65535}
                                _ = v23
                                _ = v24
                            }
                        } else {
                            var v25 int = fi13
                            v17 = (v17 ^ 100)
                            _ = v25
                        }
                    }
                    _ = len(v1)
                    v26 := GenS3{F0: 42}
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v26
                case 2147483647:
                    if false {
                        switch ("\n\t" + (("foo bar" + "hello") + "foo bar")) {
                        case "\n\t":
                            if (0.1 >= 0.1) {
                                var v27 uint64 = (^(v2))
                                fmt.Println(v6, v2)
                                _ = v27
                            } else {
                                var v28 uint64 = v2
                                _ = v28
                            }
                        default:
                            fmt.Println(v2, v14, v3)
                        }
                    } else {
                        v4[fi12] = "hello"
                        if (100 >= -128) {
                            var v29 uint32 = 4294967295
                            _ = v29
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v14, v6)
                        v6 = (v6 | -2147483648)
                    }
                default:
                    v6 = (v6 / ((v6) | 1))
                    if (true && ((-2.25 + -2.25) != 1e10)) {
                        fmt.Println(v14, v14)
                        v30 := GenS4{F0: true}
                        _ = len(v4)
                        _ = v30
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                }
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            }
            v31 := GenS4{F0: (false && false)}
            _ = fi13
            _ = v14
            _ = v31
        }
        v32 := GenS4{F0: true}
        _ = fi12
        _ = v32
    }
    v3 += (v3 & v3)
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int16 = (v3 / ((v3) | 1))
        _ = fi33
        _ = v34
    }
    v3 = (v3 & v3)
    if (false || true) {
        v6 += (v6 - 2147483647)
    }
    v35 := GenS5{F0: 1, F1: 669, F2: 255}
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v6
    _ = v7
}
