package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    v1 := []string{(("hello" + "") + (("\n\t" + ("hello" + "")) + "a"))}
    var v2 int32 = 1
    var v3 int32 = v2
    v2 = v3
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    switch ("" + "\n\t") {
    case "":
        v4 := GenS1{F0: 4294967295, F1: 593}
        _ = v4
    case "\n\t":
        if false {
            v2 -= (v2 ^ (v2 + v2))
        } else {
            fmt.Println(v2)
        }
    case "a":
        v3 = (v2 - -453)
    default:
        fmt.Println(v3, v2)
    }
    if (32767 <= -832) {
        if (-2147483648 >= 569) {
            fmt.Println(v2, v2)
        }
        v5 := GenS1{F0: 42, F1: 362}
        v1 = append(v1, "")
        _ = v5
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            v3 += (v3 + v2)
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            _ = fi6
        }
        fmt.Println(v3, v2)
    }
    var v7 uint = 2
    if ((43 > 42) || (1 != -921)) {
        if false {
            v8 := GenS1{F0: 401, F1: 2147483648}
            _ = v8
        }
        v7 = v7
        var v9 float64 = 0.1
        _ = v9
    } else {
        switch "" {
        case "\n\t":
            var v10 int16 = -1
            _ = v10
        case "a":
            v11 := GenS2{F0: -1}
            _ = v11
        case "":
            var v12 uint8 = 1
            _ = v12
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            if (-805 > -187) {
                v14 := GenS3{F0: (v7 + (v7 / ((v7) | 1))), F1: int16(v2)}
                _ = v14
            } else {
                var v15 int64 = 0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = (-(v3))
                _ = v15
            }
            v16 := GenS2{F0: (fi13 + fi13)}
            var v17 float32 = 1.5
            _ = fi13
            _ = v16
            _ = v17
        }
        if (!(4294967295 < 42)) {
            fmt.Println(v2, v7)
            for fi18 := 0; fi18 < 19; fi18++ {
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS1{F0: 65535, F1: 2147483647}
                    if true {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        fmt.Println(fi18, v3, v2)
                    } else {
                        switch "" {
                        case "\n\t":
                            var v21 int8 = 38
                            _ = v21
                        case "hello":
                            v22 := GenS3{F0: v7, F1: 1000}
                            _ = v22
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        }
                        v23 := GenS3{F0: 42, F1: 42}
                        _ = v23
                    }
                    var v24 int8 = -115
                    _ = fi19
                    _ = v20
                    _ = v24
                }
                _ = fi18
            }
            v7 = 255
        }
    }
    v2 = (^(v3))
    for fi25 := 0; fi25 < 1; fi25++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi25
    }
    var v26 float32 = 3.14159
    var v27 uint16 = 406
    if true {
        v2 += (v3 | 2147483647)
        if len(v1) > 0 {
            v1[0] = ("hello" + "foo bar")
        }
        v1 = append(v1, "\n\t")
    }
    v2 = (v3 * -2147483648)
    var v28 string = "\n\t"
    v2 = v2
    switch -1 {
    case -537:
        v7 = 949
    case 0:
        fmt.Println(v28, v27)
    default:
        fmt.Println(v26, v27)
    }
    if true {
        switch 42 {
        case 32767:
            var v29 string = ""
            _ = v29
        case -602:
            switch 956 {
            case -159:
                var v30 int = 127
                _ = v30
            case 4294967295:
                v27 += v27
            case -2147483648:
                var v31 int64 = 9223372036854775807
                _ = v31
            default:
                v3 = v2
            }
        default:
            fmt.Println(v28, v26)
        }
        fmt.Println(v28)
    } else {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v28, v26, v27)
        case "foo bar":
            v2 += 42
        case "a":
            var v32 uint = 2
            _ = v32
        }
        switch (("foo bar" + (v28 + (v28 + v28))) + v28) {
        case "foo bar":
            v26 = (0.1 + v26)
        default:
            for fi33 := 0; fi33 < 5; fi33++ {
                for fi34 := 0; fi34 < 19; fi34++ {
                    var v35 uint16 = v27
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        }
        var v36 string = "\n\t"
        _ = v36
    }
    v1 = append(v1, "hello")
    v27 += v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v7
}
