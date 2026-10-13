package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint32, p1 uint, p2 int16) uint8 {
    v7 := GenS1{F0: 0, F1: (p2 - 32767), F2: 1}
    var v8 int32 = -270
    switch "hello" {
    case "hello":
        v8 = (v8 - (-(v8)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int16 = 1
    v2 := []uint32{1, 348, 477, 788}
    v3 := make(map[string]bool)
    v4 := []string{"\n\t", "hello", ("" + "hello"), ""}
    var v5 int = -2147483648
    var v6 int16 = v1
    v6 -= (v1 - v6)
    if true {
        _ = len(v3)
        fmt.Println(v5)
    } else {
        v6 -= v6
        var v9 bool = true
        _ = v9
    }
    fmt.Println(v1, v6, v6)
    v10 := GenS2{F0: v6, F1: false}
    fmt.Println(v6, v6, v5)
    var v11 float32 = 3.14159
    fmt.Println(v1)
    v1 = (v6 & (v1 >> ((2) & 63)))
    if (183 != 128) {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        var v12 int16 = v6
        if false {
            switch "a" {
            case "\n\t":
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint32 = 0
                    if true {
                        if (255 == 18446744073709551615) {
                            v6 = (^(v6))
                        }
                        v14 = v14
                        v5 = fi13
                    } else {
                        v5 = v5
                    }
                    _ = fi13
                    _ = v14
                }
            case "":
                switch "" {
                case "foo bar":
                    var v15 bool = false
                    _ = v15
                case "\n\t":
                    var v16 int64 = 9223372036854775807
                    _ = v16
                case "a":
                    for fi17 := 0; fi17 < 10; fi17++ {
                        if false {
                            v1 = v6
                        } else {
                            if false {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                for fi18 := 0; fi18 < 3; fi18++ {
                                    fmt.Println(v6, v6)
                                    v19 := GenS1{F0: 0, F1: 1, F2: 1}
                                    v20 := GenS3{F0: v6, F1: 255}
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            } else {
                                var v21 int = 4294967295
                                _ = v21
                            }
                            var v22 uint8 = 127
                            _ = v22
                        }
                        v3[("hello" + ("hello" + "a"))] = true
                        _ = fi17
                    }
                default:
                    v12 -= v1
                }
            }
            if (true && false) {
                v2 = append(v2, 643)
                var v23 uint8 = genFn0(547, 1, v1)
                _ = v23
            }
            v11 = v11
        }
        _ = v12
    } else {
        delete(v3, (("hello" + "") + "foo bar"))
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := GenS2{F0: (v1 >> ((65535) & 63)), F1: false}
                if (false || (v6 >= v6)) {
                    fmt.Println(v5)
                }
                v11 = (1e10 - v11)
                fmt.Println(fi24, v11, v6)
                _ = fi25
                _ = v26
            }
            if true {
                v6 += v1
                for fi27 := 0; fi27 < 17; fi27++ {
                    delete(v3, "")
                    _ = fi27
                }
                var v28 uint8 = 213
                _ = v28
            }
            fmt.Println(v11)
            _ = fi24
        }
    }
    if (v6 == (v6 ^ (v6 ^ v6))) {
        var v29 uint32 = 406
        _ = v29
    } else {
        var v30 string = "a"
        _ = v30
    }
    var v31 int8 = 0
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int32 = -1
        _ = fi32
        _ = v33
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 += v1
    if false {
        var v34 uint8 = 128
        _ = len(v4)
        _ = v34
    }
    if true {
        fmt.Println(v6, v31, v5)
    }
    var v35 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
