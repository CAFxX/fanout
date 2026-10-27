package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 42
    v3 := []string{"\n\t", "hello"}
    switch 0 {
    case 42:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v4 int32 = v2
        _ = v4
    }
    for fi5 := 0; fi5 < 17; fi5++ {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = v2
            v6 := GenS1{F0: (^(v2)), F1: 2}
            _ = v6
        } else {
            if true {
                var v7 int64 = 9223372036854775807
                if (false && (("foo bar" + "foo bar") > "foo bar")) {
                    v8 := GenS1{F0: (v2 & (-(v2))), F1: 4294967295}
                    fmt.Println(fi5, fi5, v2)
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        v2 += (-(v2))
                        var v10 uint16 = 1
                        switch -9223372036854775808 {
                        case 4294967295:
                            var v11 uint8 = 6
                            _ = v11
                        case -9223372036854775808:
                            v12 := GenS1{F0: 0, F1: 255}
                            fmt.Println(v2, fi9)
                            v2 += (v2 << ((0) & 63))
                            _ = v12
                        case 127:
                            var v13 int32 = (-(v2))
                            _ = v13
                        default:
                            var v14 uint = 18446744073709551615
                            _ = v14
                        }
                        _ = fi9
                        _ = v10
                    }
                }
                var v15 uint64 = 18446744073709551615
                _ = v15
                _ = v7
            }
        }
        if false {
            switch 366 {
            case 547:
                var v16 uint16 = 65535
                _ = v16
            default:
                var v17 int32 = (v2 ^ 0)
                _ = v17
            }
            v18 := GenS1{F0: v2, F1: 0}
            _ = v18
        } else {
            var v19 uint = 0
            _ = v19
        }
        var v20 bool = false
        _ = fi5
        _ = v20
    }
    switch (("foo bar" + "hello") + "a") {
    case "hello":
        var v21 float64 = 0.0
        _ = v21
    case "":
        var v22 int64 = 320
        _ = v22
    case "\n\t":
        if false {
            var v23 int = 4294967295
            _ = v23
        } else {
            v24 := GenS2{F0: ((v2 == v2) && false)}
            var v25 bool = (true && false)
            _ = v24
            _ = v25
        }
    default:
        var v26 uint16 = 449
        _ = v26
    }
    var v27 float32 = (float32(3.14159) * -2.25)
    var v28 int16 = 32767
    v3 = append(v3, "hello")
    if (true && (((704 > 349) && (-417 > 232)) || ((v28 >> ((2) & 63)) <= (^(v28))))) {
        var v29 uint = 1
        _ = v29
    } else {
        v30 := GenS1{F0: v2, F1: 42}
        _ = v30
    }
    v2 += -1
    var v31 uint = 694
    var v32 int16 = v28
    var v33 int = 2147483647
    v33 += -351
    var v34 int64 = -1
    v35 := GenS3{F0: -9223372036854775808, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
