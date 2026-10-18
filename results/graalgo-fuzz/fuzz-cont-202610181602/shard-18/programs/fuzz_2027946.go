package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    v3 := []uint32{2147483647, 1, 1}
    v4 := make(map[string]string)
    var v5 bool = (128 == 680)
    var v6 uint64 = 2
    var v7 int64 = -532
    _ = len(v2)
    v7 = (-(v7))
    var v8 string = ""
    v9 := GenS1{F0: 42, F1: (v7 | v7)}
    v6 += v6
    if (4294967295 < -555) {
        v8 = v8
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 string = v8
        var v12 int32 = 901
        var v13 int32 = 42
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(v8, v7, v5)
    var v14 uint8 = 2
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            v7 = -9223372036854775808
            fmt.Println(v5, fi16, v14)
            _ = fi16
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            v7 -= v7
            if true {
                if (!(uint32(v6) <= uint32(v6))) {
                    switch v8 {
                    case "foo bar":
                        if (v5 && ((!(v5 || (!false))) || false)) {
                            v6 = (v6 - v6)
                            var v18 int = (fi17 - (fi15 + -9223372036854775808))
                            var v19 int = 85
                            _ = v18
                            _ = v19
                        } else {
                            var v20 uint = 42
                            v6 = v6
                            if true {
                                v2 = append(v2, uint32(v20))
                                v7 -= (^(v7))
                            }
                            _ = v20
                        }
                        v21 := GenS2{F0: 4294967295, F1: 2}
                        if v5 {
                            v22 := v4["a"]
                            var v23 string = v22
                            _ = v22
                            _ = v23
                        }
                        _ = v21
                    case "a":
                        switch ("a" + v8) {
                        case "a":
                            var v24 int = (fi15 / (((fi15 / ((int(v6)) | 1))) | 1))
                            _ = v24
                        }
                    }
                }
            } else {
                var v25 string = "foo bar"
                v5 = true
                v26 := GenS1{F0: 998, F1: v7}
                _ = v25
                _ = v26
            }
            fmt.Println(v8, v6)
            _ = fi17
        }
        _ = fi15
    }
    switch "" {
    case "":
        v8 = (v8 + v8)
    case "foo bar":
        v14 -= (v14 | 1)
    case "\n\t":
        if (v7 > v7) {
            var v27 int64 = (v7 - (v7 % (((v7 ^ (v7 + -9223372036854775808))) | 1)))
            delete(v1, (((v8 + ("foo bar" + v8)) + ("a" + v8)) + "foo bar"))
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        fmt.Println(v8)
        if (!v5) {
            var v29 uint16 = 255
            var v30 uint32 = 2147483647
            fmt.Println(v30)
            _ = v29
            _ = v30
        } else {
            var v31 uint64 = (v6 * (v6 >> ((921) & 63)))
            _ = v31
        }
        _ = fi28
    }
    var v32 uint32 = 65535
    v6 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
