package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 string; F1 int16; F2 int8 }

func main() {
    var v1 bool = true
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    v4 := []string{("foo bar" + "hello"), "hello", "foo bar"}
    v5 := make(map[int]string)
    var v6 uint64 = 882
    delete(v5, -2147483648)
    v4 = append(v4, "hello")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float32 = 0.1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v1 = (42 > 42)
        _ = fi9
    }
    if v1 {
        v10 := GenS1{F0: v1, F1: -2147483648}
        switch "\n\t" {
        case "":
            fmt.Println(v1, v6)
        case "a":
            if (float32(1.5) >= 1.5) {
                var v11 int64 = -9223372036854775808
                for fi12 := 0; fi12 < 17; fi12++ {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        var v14 string = "hello"
                        _ = fi13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 1; fi15++ {
                        switch "a" {
                        case "":
                            fmt.Println(fi12)
                        case "hello":
                            v1 = (v1 && (v1 || (2 >= 2147483648)))
                        case "a":
                            if v1 {
                                var v16 uint16 = 42
                                var v17 bool = (1.5 != 0.1)
                                fmt.Println(v17, v1)
                                v6 += (v6 << ((4294967295) & 63))
                                _ = v16
                                _ = v17
                            }
                        default:
                            var v18 uint64 = v6
                            _ = v18
                        }
                        _ = len(v4)
                        var v19 uint64 = 1
                        var v20 int8 = -30
                        switch "" {
                        case "\n\t":
                            var v21 uint = 214
                            var v22 bool = (fi15 != (fi12 + (fi12 << (((^(v21))) & 63))))
                            _ = v21
                            _ = v22
                        case "foo bar":
                            var v23 int64 = 42
                            _ = v23
                        default:
                            var v24 int64 = v11
                            _ = v24
                        }
                        _ = fi15
                        _ = v19
                        _ = v20
                    }
                    v25 := GenS2{F0: 398, F1: 42}
                    _ = fi12
                    _ = v25
                }
                _ = v11
            }
        }
        _ = v10
    } else {
        fmt.Println(v6, v6)
        var v26 bool = (349 == -288)
        _ = v26
    }
    v1 = v1
    v1 = ((-426 != -9223372036854775808) && (1 > 0))
    fmt.Println(v6, v1)
    if true {
        v27 := GenS1{F0: true, F1: -713}
        for fi28 := 0; fi28 < 12; fi28++ {
            switch (fi28 >> ((4294967295) & 63)) {
            case 2147483647:
                var v29 float64 = 1e10
                _ = v29
            case 4294967295:
                v30 := GenS3{F0: "foo bar", F1: -451, F2: -128}
                _ = v30
            case 32768:
                var v31 float64 = 1e10
                _ = v31
            }
            var v32 uint32 = 4294967295
            _ = fi28
            _ = v32
        }
        _ = v27
    } else {
        v6 -= (^(v6))
        for fi33 := 0; fi33 < 17; fi33++ {
            var v34 uint16 = 179
            _ = fi33
            _ = v34
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
