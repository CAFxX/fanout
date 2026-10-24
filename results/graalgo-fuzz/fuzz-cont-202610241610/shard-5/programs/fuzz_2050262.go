package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int64 = 0
    var v2 uint32 = 739
    v3 := []int{-2147483648, 128, -1}
    v4 := []uint32{(v2 - 42), (v2 & (^(v2))), 537}
    var v5 uint64 = uint64(v2)
    var v6 int = 32767
    v7 := GenS1{F0: 399, F1: -1, F2: 2}
    for fi8 := 0; fi8 < 20; fi8++ {
        v5 = (v5 | v5)
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint32 = v2
            v10 -= v10
            if false {
                v3 = append(v3, fi8)
                var v11 uint8 = 127
                _ = v11
            }
            v12 := GenS1{F0: 1, F1: 0, F2: 942}
            if false {
                var v13 int = 128
                _ = len(v4)
                fmt.Println(fi8, v2, v1)
                _ = v13
            }
            v14 := GenS1{F0: 32767, F1: int8(v10), F2: 255}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi9
            _ = v10
            _ = v12
            _ = v14
        }
        v2 -= v2
        _ = fi8
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: 255, F1: 127, F2: 255}
        v3 = append(v3, fi15)
        switch (v6 & (fi15 + (v6 & fi15))) {
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        if (255 <= 0) {
            switch ("\n\t" + ("a" + "\n\t")) {
            case "foo bar":
                if false {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        if false {
                            v5 = (v5 >> ((0) & 63))
                            if len(v3) > 0 {
                                v3[0] = (fi17 - fi17)
                            }
                            var v19 uint32 = (v2 * v2)
                            _ = v19
                        } else {
                            var v20 uint64 = (v5 - v5)
                            _ = v20
                        }
                        if (true || false) {
                            _ = len(v4)
                            var v21 int32 = 42
                            var v22 int32 = (v21 + 892)
                            _ = v21
                            _ = v22
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi18
                    }
                    var v23 float64 = ((((float64(((0.1 + (((3.14159 * -2.25) + (1.5 - float32(0.1))) * (0.0 * 0.0))) - 3.14159)) - (3.14159 + (float64((float32(3.14159) - ((0.1 - 0.1) * 0.1))) - 0.1))) + float64(0.0)) * 3.14159) - ((0.0 + 3.14159) + 0.0))
                    _ = v23
                }
            case "hello":
                v24 := GenS1{F0: 255, F1: 96, F2: 0}
                _ = v24
            default:
                var v25 int = (v6 << ((65535) & 63))
                _ = v25
            }
            fmt.Println(v1, v1, v6)
        }
        _ = fi17
    }
    v1 += -945
    switch (("foo bar" + "foo bar") + ("foo bar" + "a")) {
    case "a":
        if true {
            if (2147483647 < -201) {
                var v26 uint64 = 258
                fmt.Println(v26)
                _ = v26
            } else {
                fmt.Println(v1, v2, v5)
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v27 bool = true
                    _ = len(v3)
                    _ = v27
                } else {
                    v5 = v5
                    v6 = v6
                }
                switch ("\n\t" + "a") {
                case "a":
                    v5 = v5
                case "foo bar":
                    v5 = v5
                case "\n\t":
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v29 := GenS2{F0: (v6 ^ fi28), F1: 4294967295, F2: 1}
                        fmt.Println(fi28, v6)
                        _ = fi28
                        _ = v29
                    }
                default:
                    v30 := GenS3{F0: (v5 ^ uint64(v2))}
                    _ = v30
                }
            }
            v31 := GenS3{F0: 42}
            var v32 string = "\n\t"
            _ = v31
            _ = v32
        } else {
            v6 += v6
            var v33 float32 = 0.1
            switch "" {
            case "a":
                for fi34 := 0; fi34 < 11; fi34++ {
                    v1 = (v1 | v1)
                    var v35 uint32 = 2147483647
                    var v36 uint = 255
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case "foo bar":
                v2 = (v2 ^ (v2 + v2))
            case "\n\t":
                v1 += (v1 >> ((0) & 63))
            default:
                switch "foo bar" {
                case "foo bar":
                    v37 := GenS1{F0: 65535, F1: -57, F2: uint16(v5)}
                    _ = v37
                case "a":
                    var v38 float32 = v33
                    _ = v38
                default:
                    var v39 uint64 = 4294967295
                    _ = v39
                }
            }
            _ = v33
        }
    default:
        v2 = 4294967295
    }
    _ = len(v4)
    fmt.Println(v2, v2, v6)
    _ = len(v3)
    var v40 uint = 18446744073709551615
    if (187 <= 39) {
        v6 += v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
