package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]string)
    var v3 bool = true
    if ((1e10 / 3.14159) == 0.0) {
        if false {
            if false {
                v1 = v1
                fmt.Println(v1, v1)
                if (false || v3) {
                    switch "hello" {
                    case "a":
                        if v3 {
                            delete(v2, 2)
                        } else {
                            switch (("hello" + (("" + "foo bar") + "hello")) + "hello") {
                            case "foo bar":
                                fmt.Println(v1, v1)
                            case "":
                                var v4 int8 = 100
                                _ = v4
                            case "\n\t":
                                v1 += (v1 * 0.0)
                            default:
                                var v5 uint64 = 18446744073709551615
                                _ = v5
                            }
                        }
                    case "foo bar":
                        delete(v2, 2147483647)
                    default:
                        v1 = v1
                    }
                } else {
                    switch (("\n\t" + ("hello" + "")) + "") {
                    case "hello":
                        for fi6 := 0; fi6 < 14; fi6++ {
                            v7 := []int{(fi6 << ((65535) & 63)), (fi6 << ((2) & 63)), fi6}
                            fmt.Println(v1, v1)
                            _ = fi6
                            _ = v7
                        }
                    }
                }
            }
        } else {
            if (!(934 < 1)) {
                v3 = v3
                fmt.Println(v1, v1)
            } else {
                fmt.Println(v3, v3)
                v8 := GenS1{F0: -9223372036854775808, F1: 1}
                _ = v8
            }
        }
        v1 = (v1 + v1)
        for fi9 := 0; fi9 < 17; fi9++ {
            v1 = (0.1 / v1)
            _ = fi9
        }
    } else {
        v1 = (1e10 / v1)
        if (18446744073709551615 != 508) {
            var v10 uint32 = 258
            var v11 uint = 2
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v1, v3)
        _ = fi12
    }
    var v13 bool = v3
    if (true && (((("a" + "") + "") + "") >= "\n\t")) {
        if (307 > 327) {
            if ((32 <= 1) && v13) {
                v1 += v1
                fmt.Println(v3, v13, v3)
                v14 := GenS1{F0: 479, F1: -1}
                _ = v14
            }
        } else {
            v13 = (v3 && (v3 || (0.0 <= (float64((0.0 * v1)) / 1e10))))
        }
        v3 = v13
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int8 = 127
            var v17 float64 = 1.5
            switch "" {
            case "hello":
                var v18 float64 = (v17 + 0.0)
                _ = v18
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    } else {
        fmt.Println(v1, v3)
        v1 -= 3.14159
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 int = 527
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint16 = 558
            fmt.Println(fi21, v1)
            var v23 string = "hello"
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    v13 = true
    if true {
        switch ("" + "a") {
        case "hello":
            v24 := GenS1{F0: 42, F1: -1}
            _ = v24
        default:
            v13 = (216 == 157)
        }
        var v25 bool = (v1 < (v1 / (1.5 * 0.1)))
        var v26 int8 = 0
        _ = v25
        _ = v26
    } else {
        if (128 > 2) {
            v1 = (v1 - float32((1e10 + 0.0)))
        }
    }
    var v27 uint16 = 0
    if (true && v3) {
        v13 = (0 == -1)
    } else {
        v28 := []int{2147483647, 127, -958, -2147483648}
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 uint32 = 2147483648
            delete(v2, (v30 - (v30 << ((2) & 63))))
            fmt.Println(v3, v27)
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint8 = 96
        for fi33 := 0; fi33 < 15; fi33++ {
            v34 := []int{(-(fi33)), (fi31 - fi31), (fi31 ^ (^(fi33)))}
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    v13 = v13
    if (255 < 151) {
        delete(v2, 321)
        var v35 bool = false
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
}
