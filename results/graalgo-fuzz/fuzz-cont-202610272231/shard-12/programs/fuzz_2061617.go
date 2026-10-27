package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 string; F2 int16 }
type GenS4 struct { F0 string; F1 int16; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint16 = 42
    v2 := []uint32{2147483648, 1, 42}
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: "\n\t"}
    var v5 float32 = 3.14159
    switch "a" {
    case "hello":
        switch 128 {
        case 128:
            v6 := GenS2{F0: 0}
            _ = v6
        case -854:
            var v7 int8 = -94
            _ = v7
        default:
            if (!false) {
                fmt.Println(v5)
                v8 := GenS3{F0: 18446744073709551615, F1: ("foo bar" + ("hello" + "hello")), F2: 42}
                fmt.Println(v5)
                _ = v8
            } else {
                v1 -= 151
                if true {
                    v9 := GenS4{F0: ("foo bar" + ""), F1: 42, F2: -9223372036854775808}
                    if false {
                        for fi10 := 0; fi10 < 2; fi10++ {
                            v11 := GenS3{F0: 2, F1: ("hello" + (("" + "hello") + "a")), F2: -191}
                            var v12 int64 = -1
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        switch ("hello" + (("\n\t" + ("hello" + "\n\t")) + (("hello" + "hello") + "foo bar"))) {
                        case "":
                            var v13 int = 2147483647
                            _ = v13
                        case "\n\t":
                            v5 = ((v5 / ((v5 * v5) / v5)) * ((v5 * (0.1 * v5)) + float32(0.0)))
                        case "a":
                            v14 := GenS5{F0: 2147483647}
                            _ = v14
                        default:
                            fmt.Println(v5, v1)
                        }
                        var v15 int16 = -883
                        _ = v15
                    } else {
                        v16 := GenS6{F0: false}
                        _ = len(v2)
                        fmt.Println(v1)
                        _ = v16
                    }
                    _ = v9
                }
                v5 += float32((1.5 - float64(v5)))
            }
        }
    case "\n\t":
        v5 += v5
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := v3[748]
        delete(v3, 759)
        fmt.Println(fi17)
        _ = fi17
        _ = v18
    }
    var v19 int = -1
    fmt.Println(v19, v5)
    delete(v3, 1)
    var v20 uint32 = 65535
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 float32 = (v5 / v5)
        v22 += (-2.25 + (float32(0.0) / float32((((0.0 / -2.25) * (-2.25 + 1.5)) - -2.25))))
        _ = fi21
        _ = v22
    }
    switch ("foo bar" + "foo bar") {
    case "":
        var v23 bool = ((!(1.5 >= ((0.1 - 3.14159) / -2.25))) && true)
        _ = v23
    }
    var v24 uint32 = (v20 | v20)
    fmt.Println(v20, v20, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
