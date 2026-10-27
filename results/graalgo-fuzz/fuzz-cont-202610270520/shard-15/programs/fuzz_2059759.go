package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    var v3 uint = 255
    var v4 uint64 = 63
    var v5 float64 = (1e10 + (float64(3.14159) - (0.0 * -2.25)))
    var v6 int16 = 577
    v4 -= (v4 << ((v3) & 63))
    v7 := GenS1{F0: 1, F1: 42, F2: 934}
    v8 := v1["a"]
    switch (("foo bar" + "\n\t") + "\n\t") {
    case "":
        v8 += (v8 - v8)
    case "hello":
        if true {
            v6 += 910
            if (!true) {
                _ = len(v2)
            } else {
                var v9 float32 = (((((float32(float64(1e10)) / 1.0) - ((1.5 + -2.25) - 1.5)) * (float32(float64(((float32(0.0) * 3.14159) + float32((float64(1e10) * 0.0))))) * (-2.25 + (3.14159 - 1.5)))) + float32(v5)) - (0.1 + ((1.5 * -2.25) * 0.1)))
                v3 = v3
                if true {
                    v10 := GenS1{F0: 1, F1: 127, F2: 32767}
                    var v11 int = -346
                    _ = v10
                    _ = v11
                }
                _ = v9
            }
            fmt.Println(v3, v3)
        } else {
            if true {
                var v12 int16 = v6
                v13 := GenS1{F0: 2, F1: 255, F2: 0}
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                v8 -= (v8 >> (((v3 + (v3 * 1))) & 63))
                _ = fi14
            }
            if true {
                for fi15 := 0; fi15 < 9; fi15++ {
                    switch "a" {
                    case "":
                        if false {
                            var v16 string = "\n\t"
                            if false {
                                v17 := GenS2{F0: "\n\t"}
                                v1[("a" + "hello")] = v8
                                _ = v17
                            } else {
                                if (true || (!(42 <= 32767))) {
                                    var v18 float64 = v5
                                    _ = v18
                                }
                                v3 += 255
                            }
                            _ = v16
                        }
                    case "a":
                        switch fi15 {
                        case 127:
                            v6 = (v6 % (((v6 ^ v6)) | 1))
                        default:
                            switch (("foo bar" + "hello") + ("foo bar" + (("foo bar" + ("foo bar" + "\n\t")) + ""))) {
                            case "foo bar":
                                var v19 float32 = 1e10
                                _ = v19
                            case "a":
                                var v20 uint8 = 255
                                _ = v20
                            case "hello":
                                v3 = v3
                            default:
                                v4 = uint64(v8)
                            }
                        }
                    case "hello":
                        v21 := GenS1{F0: 7, F1: 0, F2: 32767}
                        _ = v21
                    default:
                        if true {
                            var v22 int64 = 9223372036854775807
                            for fi23 := 0; fi23 < 12; fi23++ {
                                v8 += (v8 ^ (v8 + v8))
                                _ = fi23
                            }
                            _ = v22
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = v8
                    }
                    fmt.Println(v6, fi15, v4)
                    _ = fi15
                }
                v24 := GenS3{F0: 4294967295}
                _ = v24
            } else {
                var v25 int64 = 9223372036854775807
                _ = v25
            }
        }
    }
    var v26 bool = false
    fmt.Println(v4, v5)
    var v27 uint = v3
    _ = len(v2)
    switch ("foo bar" + "foo bar") {
    case "hello":
        switch "\n\t" {
        case "\n\t":
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 string = ((("foo bar" + ("" + "")) + ("foo bar" + ("\n\t" + "a"))) + (("\n\t" + "a") + "\n\t"))
                var v30 string = ("" + v29)
                v30 = v29
                _ = fi28
                _ = v29
                _ = v30
            }
        case "hello":
            if ((!(131 > 0)) || (!(v26 && true))) {
                v5 = v5
            }
        }
    case "":
        if true {
            var v31 uint64 = 42
            var v32 float64 = 0.0
            for fi33 := 0; fi33 < 18; fi33++ {
                fmt.Println(v26, v3, v27)
                _ = fi33
            }
            _ = v31
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
