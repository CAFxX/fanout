package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 uint8 = 2
    v4 := make(map[int]int)
    v5 := []string{"", ""}
    v3 = (^(v3))
    switch "foo bar" {
    case "hello":
        v3 += v3
    case "foo bar":
        switch (("hello" + "hello") + "hello") {
        case "a":
            var v6 uint16 = 65535
            _ = v6
        default:
            v3 -= v3
        }
    }
    switch "hello" {
    case "hello":
        var v7 int = 42
        _ = v7
    case "foo bar":
        if ((-2.25 == (0.0 * -2.25)) && false) {
            var v8 uint32 = 1
            _ = v8
        } else {
            var v9 int = 128
            v9 = v9
            v9 = 4294967295
            _ = v9
        }
    case "a":
        var v10 string = ""
        _ = v10
    default:
        delete(v1, 1)
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int64 = 42
        _ = fi11
        _ = v12
    }
    if false {
        v3 = v3
        var v13 int = -9223372036854775808
        _ = v13
    } else {
        v2["\n\t"] = ((!false) && true)
    }
    v14 := GenS1{F0: 1, F1: 209, F2: -2147483648}
    switch ("hello" + "\n\t") {
    case "a":
        if (!((float32(1.5) / float32(1e10)) == ((-2.25 + 0.0) * 3.14159))) {
            var v15 int64 = -9223372036854775808
            v15 += 66
            _ = v15
        }
    case "":
        fmt.Println(v3)
    case "hello":
        v16 := v1[4294967295]
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v3 -= (^(v3))
        _ = fi17
    }
    if true {
        if (1000 > 1) {
            if (1 == 255) {
                switch "\n\t" {
                case "":
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        if ((3.14159 + (1e10 - 0.1)) != 1.5) {
                            v3 -= (^(v3))
                            if true {
                                var v18 uint = 18446744073709551615
                                _ = v18
                            } else {
                                _ = len(v5)
                                var v19 int16 = 32767
                                _ = v19
                            }
                        } else {
                            fmt.Println(v3)
                            v3 -= v3
                            switch 127 {
                            case 2147483647:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            default:
                                v20 := GenS1{F0: 127, F1: 0, F2: -795}
                                _ = v20
                            }
                        }
                    }
                case "foo bar":
                    var v21 int8 = 122
                    _ = v21
                }
            } else {
                if (1e10 <= 0.1) {
                    v22 := GenS1{F0: 100, F1: 2, F2: -2147483648}
                    v23 := GenS1{F0: 1, F1: 2147483647, F2: 0}
                    var v24 uint = 206
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    v3 += v3
                }
            }
        }
    } else {
        v25 := GenS1{F0: 127, F1: 42, F2: -2147483648}
        _ = v25
    }
    var v26 int32 = 42
    fmt.Println(v26, v26)
    v27 := GenS1{F0: 53, F1: 0, F2: 42}
    if false {
        if true {
            var v28 int64 = 334
            var v29 uint64 = 0
            v30 := GenS2{F0: 0, F1: uint64(v3)}
            _ = v28
            _ = v29
            _ = v30
        }
        var v31 uint32 = 2147483648
        _ = v31
    }
    var v32 int = -153
    v33 := GenS1{F0: 42, F1: 65535, F2: -1}
    v32 = v32
    v26 += -812
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
