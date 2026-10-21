package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint64 = 0
    v2 := []int{1}
    v3 := []uint32{2147483648, 65535}
    var v4 float64 = float64((float32(((((1e10 / float64((1e10 / -2.25))) - float64(1.5)) * 1.5) + 0.1)) - 0.1))
    v5 := GenS1{F0: "\n\t"}
    v6 := GenS1{F0: "foo bar"}
    v7 := GenS1{F0: "\n\t"}
    v8 := GenS2{F0: -291, F1: 2147483647, F2: 0}
    v2 = append(v2, -2147483648)
    switch 32768 {
    case 4294967295:
        v2 = append(v2, 4294967295)
    case 42:
        switch 4294967295 {
        case 4294967295:
            v2 = append(v2, -962)
        case -673:
            var v9 int32 = -2147483648
            _ = v9
        default:
            fmt.Println(v1)
        }
    }
    if ((("" + "hello") + "foo bar") <= "") {
        v10 := GenS3{F0: 505, F1: 4294967295}
        v2 = append(v2, 154)
        switch (("hello" + "hello") + ("a" + "\n\t")) {
        case "":
            v11 := GenS3{F0: -425, F1: 2147483647}
            _ = v11
        default:
            fmt.Println(v1, v4)
        }
        _ = v10
    }
    var v12 uint = 18446744073709551615
    if len(v3) > 0 {
        v3[0] = 136
    }
    switch "" {
    case "\n\t":
        if ((^(v1)) >= (^(v1))) {
            var v13 int = 2147483647
            _ = v13
        }
    case "a":
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    default:
        v14 := GenS2{F0: 937, F1: 1, F2: -1}
        _ = v14
    }
    v12 += v12
    var v15 int32 = -932
    var v16 int16 = 1
    switch (("a" + "") + "a") {
    case "foo bar":
        _ = len(v2)
    case "":
        switch 249 {
        case 32768:
            var v17 uint = (^(v12))
            _ = v17
        case -115:
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint32 = 0
                fmt.Println(v12, v19)
                v15 = v15
                _ = fi18
                _ = v19
            }
        }
    case "a":
        var v20 int64 = 9223372036854775807
        _ = v20
    default:
        var v21 int8 = 0
        _ = v21
    }
    var v22 uint16 = 820
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
