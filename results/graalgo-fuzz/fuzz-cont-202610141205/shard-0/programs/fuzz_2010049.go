package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 uint16 = 42
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    if (!((42 < -974) || (true || true))) {
        for fi6 := 0; fi6 < 4; fi6++ {
            switch (^(fi6)) {
            case -195:
                v3 += 42
            case 0:
                v7 := GenS1{F0: false}
                _ = v7
            case -2147483648:
                v3 -= 65535
            default:
                v8 := GenS2{F0: 295, F1: 895, F2: 112}
                _ = v8
            }
            var v9 uint16 = 119
            switch "a" {
            case "\n\t":
                v10 := GenS3{F0: 866, F1: 32767}
                _ = v10
            case "foo bar":
                if true {
                    switch (fi6 * fi6) {
                    case 32768:
                        var v11 uint32 = 4294967295
                        _ = v11
                    }
                    v9 -= v9
                    for fi12 := 0; fi12 < 4; fi12++ {
                        switch "" {
                        case "foo bar":
                            if (!(!true)) {
                                var v13 string = "a"
                                fmt.Println(v9, v13)
                                if true {
                                    v14 := v5[""]
                                    _ = v14
                                }
                                _ = v13
                            }
                        default:
                            var v15 int64 = -9223372036854775808
                            _ = v15
                        }
                        v9 += (v9 / (((^(v9))) | 1))
                        var v16 float64 = float64((-2.25 * 1.5))
                        _ = fi12
                        _ = v16
                    }
                }
            }
            _ = fi6
            _ = v9
        }
    } else {
        v17 := []int{-2147483648}
        if (42 < -128) {
            for fi18 := 0; fi18 < 4; fi18++ {
                switch "a" {
                case "foo bar":
                    var v19 float64 = 0.0
                    _ = v19
                case "\n\t":
                    var v20 int16 = 0
                    _ = v20
                case "":
                    switch "\n\t" {
                    case "\n\t":
                        var v21 uint32 = 4294967295
                        _ = v21
                    case "foo bar":
                        switch "hello" {
                        case "foo bar":
                            var v22 int32 = -1
                            _ = v22
                        case "":
                            v3 -= v3
                        case "hello":
                            v3 = (v3 << ((120) & 63))
                        default:
                            var v23 uint16 = (v3 % ((2) | 1))
                            _ = v23
                        }
                    case "":
                        for fi24 := 0; fi24 < 10; fi24++ {
                            var v25 int32 = -1
                            _ = fi24
                            _ = v25
                        }
                    default:
                        fmt.Println(fi18, fi18)
                    }
                }
                v26 := v4["foo bar"]
                fmt.Println(v3)
                _ = fi18
                _ = v26
            }
            var v27 uint8 = 255
            if true {
                v28 := GenS4{F0: 102, F1: 4294967295, F2: int8(v3)}
                fmt.Println(v3, v3)
                _ = v28
            } else {
                v3 -= (v3 + 32767)
            }
            _ = v27
        }
        v1["a"] = 4294967295
        _ = v17
    }
    var v29 string = ((("foo bar" + "hello") + "") + "")
    var v30 int32 = 2147483647
    delete(v4, "hello")
    var v31 int64 = -522
    v32 := GenS5{F0: 100}
    if true {
        for fi33 := 0; fi33 < 7; fi33++ {
            fmt.Println(v30, fi33)
            _ = fi33
        }
    }
    v30 = 42
    v3 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
