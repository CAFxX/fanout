package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 bool; F2 int64 }

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 222
    var v3 int32 = 2147483647
    v4 := make(map[string]string)
    v5 := []uint32{}
    v6 := make(map[string]string)
    fmt.Println(v3)
    v2 = (v2 | 127)
    delete(v4, "a")
    var v7 float64 = 1.5
    v8 := GenS1{F0: 2147483648, F1: (v3 & v3), F2: (v3 - v3)}
    fmt.Println(v2, v7)
    if (255 >= 0) {
        var v9 uint32 = 65535
        var v10 bool = (255 <= 65535)
        _ = v10
        _ = v9
    } else {
        var v11 uint = 2
        var v12 int16 = -778
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS2{F0: 42}
        _ = fi13
        _ = v14
    }
    if true {
        v6["hello"] = ("hello" + ("a" + "\n\t"))
    } else {
        v7 = v7
        var v15 int8 = 23
        _ = v15
    }
    switch -766 {
    case 127:
        if ((!false) && ("foo bar" < ("hello" + ("foo bar" + "foo bar")))) {
            var v16 float32 = 0.1
            var v17 int32 = v3
            _ = v16
            _ = v17
        } else {
            v7 += 1.5
            if (!(!false)) {
                v3 = (v3 / ((480) | 1))
                for fi18 := 0; fi18 < 13; fi18++ {
                    switch int(v3) {
                    case -2147483648:
                        var v19 uint16 = 1
                        _ = v19
                    case -171:
                        v20 := GenS3{F0: v2, F1: (float64((3.14159 + 0.0)) >= ((float64(1.5) + (-2.25 / float64(-2.25))) / 0.1)), F2: -9223372036854775808}
                        _ = v20
                    default:
                        v1 += v1
                    }
                    v21 := GenS1{F0: 2147483647, F1: (v3 & v3), F2: (v3 + (v3 ^ 1))}
                    fmt.Println(v7, v7)
                    _ = fi18
                    _ = v21
                }
                var v22 uint64 = 704
                _ = v22
            } else {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v2 = 106
                    _ = fi23
                }
                var v24 string = "hello"
                _ = v24
            }
        }
    case 4294967295:
        var v25 string = "hello"
        _ = v25
    case 32768:
        v6[""] = (("foo bar" + "hello") + "a")
    default:
        switch "hello" {
        case "":
            v5 = append(v5, 65535)
        case "foo bar":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "\n\t":
            _ = len(v6)
        }
    }
    fmt.Println(v1, v3, v3)
    v1 += 127
    v2 = 120
    v5 = append(v5, 2147483648)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
