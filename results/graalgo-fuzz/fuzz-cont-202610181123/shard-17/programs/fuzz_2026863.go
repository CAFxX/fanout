package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 string; F1 int64; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    v3 := []uint32{65535, 2147483648, 42}
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]uint32)
    v6 := make(map[string]uint32)
    switch 42 {
    case 177:
        if true {
            var v7 int = -183
            var v8 float32 = 0.1
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 1, F1: 65535, F2: 128}
            if false {
                if false {
                    if (9223372036854775807 < v4) {
                        v5[147] = 1
                        if true {
                            switch "\n\t" {
                            case "":
                                var v10 int64 = (v4 | 42)
                                _ = v10
                            default:
                                var v11 uint32 = 1
                                _ = v11
                            }
                            fmt.Println(v4)
                        }
                        var v12 int8 = -1
                        _ = v12
                    }
                }
                var v13 uint = 2
                switch ((("\n\t" + "a") + "a") + "\n\t") {
                case "a":
                    var v14 string = (("" + "a") + ("\n\t" + "a"))
                    _ = v14
                }
                _ = v13
            } else {
                v4 -= 42
                v15 := GenS2{F0: 100, F1: "a"}
                _ = v15
            }
            _ = v9
        }
    default:
        for fi16 := 0; fi16 < 9; fi16++ {
            if (2 == 2) {
                v17 := GenS3{F0: 0, F1: (fi16 + -527)}
                _ = v17
            }
            if (4294967295 > 42) {
                fmt.Println(v4)
                v4 += (v4 * v4)
                var v18 string = "\n\t"
                _ = v18
            } else {
                v4 = v4
                delete(v5, 0)
            }
            _ = fi16
        }
    }
    fmt.Println(v4)
    fmt.Println(v4)
    v19 := v6["a"]
    switch 127 {
    case 975:
        v20 := GenS4{F0: "hello", F1: v4, F2: -61}
        _ = v20
    case -1:
        var v21 uint8 = 2
        _ = v21
    case 32767:
        var v22 float32 = 0.0
        _ = v22
    default:
        fmt.Println(v19, v19)
    }
    var v23 uint16 = 2
    var v24 int32 = 659
    switch "hello" {
    case "hello":
        for fi25 := 0; fi25 < 5; fi25++ {
            _ = len(v1)
            v26 := GenS3{F0: v19, F1: fi25}
            v4 = int64(fi25)
            _ = fi25
            _ = v26
        }
    case "":
        if false {
            var v27 uint32 = v19
            for fi28 := 0; fi28 < 16; fi28++ {
                for fi29 := 0; fi29 < 6; fi29++ {
                    if (!((("a" + ((("a" + "\n\t") + ("foo bar" + "hello")) + "")) + "a") > "")) {
                        var v30 uint16 = (v23 | v23)
                        v19 = 557
                        _ = v30
                    } else {
                        var v31 uint = 18446744073709551615
                        var v32 int16 = 32767
                        var v33 uint16 = (v23 / (((v23 << (((v31 ^ (v31 - (^(v31))))) & 63))) | 1))
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    switch fi29 {
                    case 1:
                        v27 = 477
                    case 127:
                        v19 = (v19 - 0)
                    default:
                        v24 = (v24 ^ (v24 * int32(v23)))
                    }
                    _ = fi29
                }
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v34 := GenS1{F0: 1, F1: v23, F2: 2}
                    _ = v34
                }
                v35 := GenS2{F0: -128, F1: "foo bar"}
                _ = fi28
                _ = v35
            }
            _ = v27
        }
    case "a":
        v24 = -2147483648
    default:
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 int8 = 42
            fmt.Println(v23)
            _ = fi36
            _ = v37
        }
    }
    v19 = v19
    for fi38 := 0; fi38 < 4; fi38++ {
        for fi39 := 0; fi39 < 20; fi39++ {
            v4 -= v4
            var v40 int32 = int32(v4)
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    v19 = (v19 << ((65535) & 63))
    fmt.Println(v4, v24, v24)
    v19 = 65535
    for fi41 := 0; fi41 < 16; fi41++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
