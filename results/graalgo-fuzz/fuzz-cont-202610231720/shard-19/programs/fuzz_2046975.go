package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 int64; F1 int64; F2 string }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 string = "hello"
    v4 := make(map[uint32]uint32)
    if (!(-1 != -128)) {
        if false {
            switch 2147483647 {
            case 225:
                fmt.Println(v3)
            case -1:
                v3 = "foo bar"
            case 4294967295:
                delete(v2, 2147483647)
            default:
                v5 := []int{4294967295, 1}
                _ = v5
            }
            fmt.Println(v3)
        }
    }
    v3 = v3
    if false {
        var v6 uint = 18446744073709551615
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 string = v3
            v8 = (v8 + ("" + (v3 + v8)))
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int64 = 0
        _ = fi9
        _ = v10
    }
    var v11 int32 = 42
    var v12 int8 = 42
    v13 := GenS1{F0: 379, F1: 100, F2: (123 < (-(v12)))}
    v14 := GenS2{F0: 0, F1: 875, F2: (v3 + v3)}
    v15 := GenS1{F0: 42, F1: (^(v12)), F2: false}
    var v16 int16 = 32767
    fmt.Println(v16, v11)
    v3 = (v3 + v3)
    switch 128 {
    case 0:
        switch "hello" {
        case "a":
            for fi17 := 0; fi17 < 19; fi17++ {
                switch v3 {
                case "":
                    if (1 > 2147483648) {
                        if true {
                            fmt.Println(v3)
                            for fi18 := 0; fi18 < 5; fi18++ {
                                var v19 float32 = (3.14159 - 3.14159)
                                v12 -= (v12 * 100)
                                delete(v2, fi17)
                                _ = fi18
                                _ = v19
                            }
                        } else {
                            var v20 uint = 255
                            v20 = (v20 ^ v20)
                            _ = v20
                        }
                    }
                }
                if true {
                    v2[128] = (fi17 & (fi17 + (fi17 - (fi17 + 59))))
                }
                _ = fi17
            }
        case "\n\t":
            v21 := GenS3{F0: ((((("foo bar" + ("foo bar" + v3)) + "foo bar") + (v3 + v3)) + v3) != ("hello" + v3))}
            _ = v21
        case "":
            for fi22 := 0; fi22 < 15; fi22++ {
                fmt.Println(v16, v12, v11)
                _ = fi22
            }
        default:
            var v23 bool = (65535 == 42)
            _ = v23
        }
    case -9223372036854775808:
        var v24 int16 = v16
        _ = v24
    default:
        v25 := GenS3{F0: true}
        _ = v25
    }
    var v26 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
