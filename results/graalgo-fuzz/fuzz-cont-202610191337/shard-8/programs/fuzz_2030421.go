package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 int32 }

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{161, 42, 0, 2147483647}
    var v3 float32 = float32((0.1 + (0.0 + (1.5 - float64(-2.25)))))
    v4 := make(map[uint32]uint32)
    v5 := []string{(("a" + "a") + "\n\t")}
    for fi6 := 0; fi6 < 15; fi6++ {
        if true {
            fmt.Println(fi6, v1)
            var v7 int = fi6
            _ = v7
        } else {
            switch (("\n\t" + "") + "\n\t") {
            case "hello":
                v1 -= v1
            case "foo bar":
                switch (("\n\t" + (("a" + "") + ("" + "foo bar"))) + "hello") {
                case "":
                    var v8 uint16 = 255
                    _ = v8
                case "a":
                    var v9 int8 = 1
                    _ = v9
                case "\n\t":
                    switch "a" {
                    case "a":
                        var v10 bool = false
                        _ = v10
                    case "":
                        var v11 int16 = 1
                        _ = v11
                    case "hello":
                        fmt.Println(fi6)
                    default:
                        switch "" {
                        case "hello":
                            v12 := GenS1{F0: 2147483647, F1: 1, F2: 255}
                            _ = v12
                        case "a":
                            var v13 float32 = v3
                            _ = v13
                        }
                    }
                default:
                    v1 += v1
                }
            case "":
                _ = len(v4)
            default:
                var v14 int = 42
                _ = v14
            }
            v3 += 0.1
            delete(v4, 42)
        }
        _ = fi6
    }
    v15 := GenS2{F0: "", F1: 2147483647}
    var v16 int = -1
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 float64 = v1
        if true {
            if true {
                _ = len(v5)
                v19 := GenS2{F0: "hello", F1: 992}
                v20 := GenS3{F0: 42, F1: -1, F2: 42}
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 13; fi21++ {
                    _ = len(v5)
                    v22 := GenS1{F0: 2, F1: 9223372036854775807, F2: 2}
                    var v23 int32 = int32(fi17)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                fmt.Println(fi17)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        _ = fi17
        _ = v18
    }
    v3 = v3
    v16 = v16
    var v24 float64 = v1
    v3 = ((float32(v1) * (v3 / ((v3 / (v3 + (0.1 + 0.1))) * 1e10))) / -2.25)
    var v25 uint = 1
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint = 0
        var v28 float32 = 3.14159
        var v29 uint64 = 274
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println(v24, v24)
    fmt.Println(v16, v25)
    v24 += 3.14159
    for fi30 := 0; fi30 < 5; fi30++ {
        fmt.Println(v25, v24)
        v31 := GenS3{F0: v25, F1: 346, F2: 42}
        _ = fi30
        _ = v31
    }
    if (true || (843 <= 4294967295)) {
        v32 := GenS2{F0: "hello", F1: 738}
        for fi33 := 0; fi33 < 19; fi33++ {
            v1 = float64((((v3 / v3) * v3) / 3.14159))
            _ = fi33
        }
        if len(v2) > 0 {
            v2[0] = 553
        }
        _ = v32
    }
    fmt.Println(v16, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
