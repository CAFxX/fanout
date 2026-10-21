package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v4 := []string{"foo bar", "\n\t"}
    var v5 uint16 = 403
    var v6 int32 = v1
    v4 = append(v4, "hello")
    if (100 <= -128) {
        v5 = 0
        if ((v5 & 1) < v5) {
            if (!false) {
                delete(v2, (("hello" + "hello") + (("\n\t" + "hello") + ("\n\t" + ("" + "a")))))
                var v7 uint16 = 255
                _ = v7
            }
            v6 = -2147483648
        } else {
            var v8 int8 = 127
            _ = v8
        }
    } else {
        if false {
            var v9 uint64 = 750
            _ = v9
        } else {
            v1 = v6
        }
        fmt.Println(v6, v6, v5)
    }
    if (true || ((42 == 1000) || true)) {
        v5 -= (v5 * v5)
        v1 += v1
    }
    v10 := GenS1{F0: 128, F1: 9223372036854775807, F2: (9223372036854775807 > -1)}
    v11 := GenS2{F0: 1, F1: 1}
    if ("\n\t" <= ("hello" + (("a" + ("\n\t" + ("\n\t" + ""))) + "\n\t"))) {
        var v12 float64 = (((((0.1 * (float64((1e10 * -2.25)) + float64((float32((3.14159 / (-2.25 / ((1.5 * 3.14159) - 3.14159)))) * 0.0)))) / (3.14159 * -2.25)) * -2.25) + float64(1e10)) * ((0.0 + 0.0) / (((3.14159 / 1e10) - (float64(1e10) + ((-2.25 + 1.5) / float64(-2.25)))) / (float64(3.14159) + (0.0 / (1.5 * 0.1))))))
        for fi13 := 0; fi13 < 9; fi13++ {
            v1 += v6
            if (172 >= 42) {
                switch "foo bar" {
                case "foo bar":
                    v5 = (v5 | v5)
                case "a":
                    var v14 float32 = 1.5
                    _ = v14
                case "":
                    v15 := GenS2{F0: 42, F1: 128}
                    _ = v15
                }
            } else {
                v12 = float64((float32(0.0) * -2.25))
                v5 = (v5 << ((18446744073709551615) & 63))
            }
            fmt.Println(v1)
            _ = fi13
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v12
    }
    v16 := v3[-859]
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 uint = 0
        _ = fi17
        _ = v18
    }
    var v19 uint = 18446744073709551615
    var v20 uint16 = 253
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            v19 -= v19
            if ((1.5 - 0.0) != 0.0) {
                var v23 int8 = 1
                fmt.Println(v5, v20, fi22)
                _ = v23
            }
            fmt.Println(v6)
            _ = fi22
        }
        _ = fi21
    }
    v24 := GenS2{F0: -32768, F1: 42}
    fmt.Println(v19)
    for fi25 := 0; fi25 < 3; fi25++ {
        if true {
            if ((!(!(!(0.0 > ((1.5 + ((-2.25 * (float32(1e10) * float32(0.1))) / 3.14159)) + 3.14159))))) && ((v16 ^ v16) != (v16 | (v16 + 1)))) {
                var v26 uint16 = v20
                var v27 uint16 = v5
                switch "\n\t" {
                case "hello":
                    var v28 int32 = (v6 / ((1) | 1))
                    _ = v28
                }
                _ = v26
                _ = v27
            } else {
                for fi29 := 0; fi29 < 1; fi29++ {
                    if false {
                        v6 -= (v1 * -337)
                        v5 -= (v20 * 32767)
                        if (v1 < -1) {
                            delete(v3, -9223372036854775808)
                            v6 = v1
                        }
                        v4 = append(v4, "")
                    } else {
                        var v30 string = "hello"
                        switch v30 {
                        case "\n\t":
                            fmt.Println(v16, v19)
                        }
                        v30 = ""
                        _ = v30
                    }
                    switch (fi29 ^ (fi25 * 9223372036854775807)) {
                    case -2147483648:
                        v31 := GenS3{F0: -1, F1: -2147483648}
                        _ = v31
                    }
                    var v32 int64 = -1
                    _ = fi29
                    _ = v32
                }
            }
        }
        _ = fi25
    }
    v16 = uint32(v19)
    fmt.Println(v19, v19)
    var v33 int32 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
