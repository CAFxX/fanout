package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (!true)
    var v3 uint = 65535
    var v4 int8 = -128
    v5 := []string{"foo bar", (("" + "a") + "a"), "hello", (((("foo bar" + ("\n\t" + "foo bar")) + "foo bar") + "") + "")}
    var v6 string = "hello"
    _ = len(v5)
    switch v6 {
    case "\n\t":
        var v7 uint8 = 11
        _ = v7
    }
    var v8 float32 = (1e10 * (3.14159 + (float32(0.1) * 1e10)))
    if true {
        if (2147483647 >= -760) {
            v4 = -1
        }
    } else {
        v4 += (v4 % (((v4 | (v4 - v4))) | 1))
        fmt.Println(v6, v2)
        var v9 bool = false
        _ = v9
    }
    v3 += v3
    fmt.Println(v8)
    v10 := v1[v6]
    if v2 {
        if true {
            var v11 uint8 = 1
            switch "\n\t" {
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "hello":
                v12 := GenS1{F0: (v3 * v3), F1: 455, F2: (v4 - 0)}
                _ = v12
            case "foo bar":
                for fi13 := 0; fi13 < 9; fi13++ {
                    var v14 string = "hello"
                    for fi15 := 0; fi15 < 5; fi15++ {
                        var v16 int16 = 42
                        v17 := GenS1{F0: v3, F1: -1, F2: v4}
                        if false {
                            v16 = v16
                            fmt.Println(v11)
                            fmt.Println(v14)
                            switch (fi13 / ((fi15) | 1)) {
                            case -169:
                                if (255 < 255) {
                                    switch "\n\t" {
                                    case "":
                                        _ = len(v5)
                                    case "hello":
                                        fmt.Println(v3)
                                    case "\n\t":
                                        switch v6 {
                                        case "a":
                                            v6 = v14
                                        case "foo bar":
                                            v10 += (v10 & (v10 * (v10 >> ((42) & 63))))
                                        case "hello":
                                            var v18 string = (v14 + v14)
                                            _ = v18
                                        }
                                        var v19 int8 = v4
                                        fmt.Println(v3, v4, v2)
                                        switch ((((("\n\t" + v6) + "hello") + ("hello" + ("" + (v14 + "a")))) + (v6 + v14)) + (v14 + v6)) {
                                        case "hello":
                                            fmt.Println(fi13, fi15)
                                        case "":
                                            if (!(!true)) {
                                                var v20 uint32 = (v10 >> (((v3 * (v3 * (v3 / (((^(v3))) | 1))))) & 63))
                                                v5 = append(v5, (v6 + "a"))
                                                _ = v20
                                            }
                                        }
                                        var v21 uint16 = 2
                                        if len(v5) > 0 {
                                            v5[0] = v14
                                        }
                                        v1[(v6 + (v14 + (v6 + "hello")))] = (v10 | (v10 | v10))
                                        var v22 float64 = float64(float32(((1e10 + (float64((0.0 * v8)) * 3.14159)) - ((3.14159 / (-2.25 + 1.5)) / float64((1e10 * v8))))))
                                        switch (fi13 << ((v3) & 63)) {
                                        case 1:
                                            var v23 int16 = (v16 % ((v16) | 1))
                                            _ = v23
                                        case -9223372036854775808:
                                            if len(v5) > 0 {
                                                v5[0] = (("a" + "hello") + ((v14 + "a") + "hello"))
                                            }
                                        case 4294967295:
                                            fmt.Println(v6)
                                        default:
                                            var v24 int = 32768
                                            _ = v24
                                        }
                                        _ = v19
                                        _ = v21
                                        _ = v22
                                    }
                                }
                            }
                            fmt.Println(v10)
                            v25 := GenS1{F0: v3, F1: 1, F2: -128}
                            if ((((("foo bar" + (("foo bar" + v14) + v14)) + "a") + v6) + v6) < (("hello" + v14) + v6)) {
                                var v26 uint32 = v10
                                v4 = -1
                                _ = v26
                            }
                            _ = v25
                        } else {
                            if false {
                                var v27 int = (^(fi15))
                                _ = v27
                            } else {
                                var v28 int32 = -2147483648
                                if v2 {
                                    v29 := GenS1{F0: 0, F1: 784, F2: (-(v4))}
                                    _ = v29
                                }
                                v16 = 42
                                v30 := GenS1{F0: 4294967295, F1: 42, F2: -1}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                fmt.Println(v4, v28)
                                switch (-(fi13)) {
                                case 755:
                                    var v31 uint = v3
                                    _ = v31
                                }
                                _ = v28
                                _ = v30
                            }
                            v10 -= v10
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                }
            default:
                v3 -= 0
            }
            var v32 uint = 255
            _ = v11
            _ = v32
        }
        var v33 int8 = 0
        if (!(!false)) {
            v34 := GenS2{F0: v3, F1: 1, F2: (v4 >> ((4294967295) & 63))}
            _ = v34
        }
        _ = v33
    }
    if len(v5) > 0 {
        v5[0] = v6
    }
    fmt.Println(v6, v3)
    var v35 bool = true
    var v36 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
