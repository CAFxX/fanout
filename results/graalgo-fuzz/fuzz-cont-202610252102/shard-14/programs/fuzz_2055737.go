package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := []string{(("" + "") + ("\n\t" + "foo bar")), "\n\t", (("" + ("foo bar" + "foo bar")) + "a"), "\n\t"}
    v2 := make(map[string]int)
    var v3 string = "\n\t"
    if (!true) {
        var v4 uint32 = 1
        _ = v4
    }
    fmt.Println(v3)
    v3 = (v3 + (("\n\t" + v3) + ((v3 + v3) + (v3 + "a"))))
    fmt.Println(v3)
    v2[v3] = -9223372036854775808
    switch 32767 {
    case 137:
        fmt.Println(v3)
    default:
        for fi5 := 0; fi5 < 2; fi5++ {
            if (3.14159 < (-2.25 + 1e10)) {
                if (255 == 65535) {
                    v6 := v2["hello"]
                    _ = v6
                } else {
                    var v7 float32 = (0.1 + 0.1)
                    if (0 <= 567) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v8 bool = true
                        v9 := GenS1{F0: 1, F1: v3}
                        _ = v8
                        _ = v9
                    }
                    v10 := GenS2{F0: 1, F1: 42}
                    _ = v10
                    _ = v7
                }
                if (!true) {
                    var v11 int = (fi5 << ((42) & 63))
                    var v12 int8 = 53
                    v13 := GenS1{F0: 32767, F1: "hello"}
                    _ = v11
                    _ = v12
                    _ = v13
                }
                var v14 float32 = 1e10
                _ = v14
            } else {
                switch "\n\t" {
                case "\n\t":
                    if (!(!(42 >= 65535))) {
                        for fi15 := 0; fi15 < 4; fi15++ {
                            var v16 int8 = -128
                            if (float32(0.0) == (((1.5 - 1.5) / -2.25) * 3.14159)) {
                                v16 += 127
                                var v17 int32 = -2147483648
                                _ = v17
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v18 := GenS3{F0: 55, F1: -128}
                        v3 = ((((v3 + "hello") + v3) + (v3 + "")) + v3)
                        v19 := GenS4{F0: -659}
                        _ = v18
                        _ = v19
                    }
                case "foo bar":
                    var v20 int32 = -364
                    _ = v20
                case "a":
                    switch 1 {
                    case -91:
                        v3 = "a"
                    case -9223372036854775808:
                        fmt.Println(fi5)
                    default:
                        v21 := GenS3{F0: 127, F1: -128}
                        _ = v21
                    }
                default:
                    v22 := GenS3{F0: -1, F1: 0}
                    _ = v22
                }
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int32 = 42
                    v3 = v3
                    _ = fi23
                    _ = v24
                }
            }
            switch ("" + (v3 + "hello")) {
            case "foo bar":
                fmt.Println(fi5)
            case "\n\t":
                v3 = (("" + v3) + "hello")
            case "hello":
                v3 = (("" + v3) + v3)
            default:
                var v25 uint8 = 255
                _ = v25
            }
            _ = fi5
        }
    }
    var v26 int16 = -617
    v3 = v3
    for fi27 := 0; fi27 < 17; fi27++ {
        v26 = (-(v26))
        fmt.Println(v26)
        fmt.Println(v3)
        _ = fi27
    }
    switch v3 {
    case "foo bar":
        var v28 uint8 = 127
        _ = v28
    case "\n\t":
        var v29 int8 = -128
        _ = v29
    default:
        var v30 uint = 0
        _ = v30
    }
    v2[v3] = 2147483647
    if (0 == 65535) {
        v26 -= v26
        v3 = v3
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v31 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
}
