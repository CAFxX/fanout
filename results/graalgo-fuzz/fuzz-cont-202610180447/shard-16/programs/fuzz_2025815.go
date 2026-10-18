package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint16 = 0
    v2 := []string{"\n\t", "foo bar", "a"}
    var v3 int16 = -1
    var v4 float64 = 3.14159
    v5 := make(map[uint32]uint32)
    fmt.Println(v4, v3, v4)
    fmt.Println(v3, v1, v4)
    _ = len(v5)
    v6 := v5[1]
    v4 = v4
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v3)
        fmt.Println(v6, v6, v3)
        _ = fi7
    }
    if false {
        var v8 uint32 = v6
        _ = len(v2)
        _ = v8
    } else {
        var v9 int64 = -1
        v10 := GenS1{F0: 295}
        var v11 int16 = -32768
        _ = v10
        _ = v11
        _ = v9
    }
    v1 -= (v1 >> ((921) & 63))
    if (true || ((true || true) || true)) {
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS2{F0: 0, F1: 1}
            v14 := GenS3{F0: 2}
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(v4)
        v3 += -480
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v6 += v6
        if true {
            for fi16 := 0; fi16 < 5; fi16++ {
                if (!true) {
                    switch (fi16 ^ fi16) {
                    case -1:
                        switch "hello" {
                        case "foo bar":
                            var v17 float64 = v4
                            _ = v17
                        case "":
                            var v18 uint32 = v6
                            _ = v18
                        }
                    case -9223372036854775808:
                        fmt.Println(fi15, fi15, fi16)
                    default:
                        fmt.Println(v1, v4)
                        _ = len(v5)
                        if (2 != (^(v6))) {
                            var v19 uint8 = 2
                            _ = v19
                        } else {
                            v6 += (v6 / (((v6 * v6)) | 1))
                            switch -1 {
                            case 32768:
                                v6 -= (v6 | uint32(v1))
                            }
                        }
                    }
                    var v20 uint = 4294967295
                    _ = v20
                } else {
                    _ = len(v5)
                }
                fmt.Println(fi16, fi15)
                fmt.Println(v1)
                _ = fi16
            }
            var v21 uint8 = 255
            switch "foo bar" {
            case "":
                if (true && (42 >= 83)) {
                    var v22 float64 = 3.14159
                    v22 = 1.5
                    _ = v22
                }
            case "hello":
                if (false || true) {
                    v21 -= v21
                    if false {
                        v3 = (v3 * 32767)
                        switch "" {
                        case "hello":
                            v23 := GenS1{F0: -1}
                            _ = v23
                        default:
                            if true {
                                fmt.Println(v4)
                                v4 -= v4
                            } else {
                                v4 -= -2.25
                                v6 += (v6 - v6)
                            }
                        }
                    }
                    var v24 uint64 = 4294967295
                    _ = v24
                }
            case "\n\t":
                if false {
                    fmt.Println(v3, v4)
                    v3 = (v3 + (v3 << ((279) & 63)))
                    for fi25 := 0; fi25 < 18; fi25++ {
                        fmt.Println(v1, v6, v4)
                        var v26 uint8 = (v21 << ((2) & 63))
                        fmt.Println(fi25, v21, v6)
                        _ = fi25
                        _ = v26
                    }
                }
            }
            _ = v21
        }
        _ = fi15
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v1 -= v1
        v28 := GenS1{F0: -2147483648}
        v29 := GenS3{F0: 1}
        _ = fi27
        _ = v28
        _ = v29
    }
    switch int(v6) {
    case 127:
        var v30 float32 = 1.5
        _ = v30
    case -2147483648:
        var v31 uint64 = 18446744073709551615
        _ = v31
    }
    var v32 bool = (false || (32768 > 32768))
    if true {
        v6 += (^(v6))
    } else {
        v32 = ("\n\t" <= ("foo bar" + ("" + ("foo bar" + ("a" + "foo bar")))))
        fmt.Println(v4, v32)
        v6 = v6
    }
    v4 = float64(float32(float64(1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
