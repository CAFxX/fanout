package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int16; F1 bool; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint, p1 uint64) int16 {
    fmt.Println(p0, p0)
    v6 := GenS1{F0: 42, F1: 18446744073709551615}
    p1 = p1
    if true {
        var v7 float32 = ((0.1 / -2.25) / 1.0)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return -32768
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 int32 = -286
    var v4 uint8 = 255
    v5 := []string{"\n\t", "a"}
    if (42 <= 42) {
        v5 = append(v5, "\n\t")
        v3 = (v3 + (v3 % (((-(v3))) | 1)))
        var v8 uint = 2
        _ = v8
    }
    if (v1 && v1) {
        v9 := GenS1{F0: (v4 - v4), F1: 42}
        switch "" {
        case "\n\t":
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 int64 = -9223372036854775808
                _ = fi10
                _ = v11
            }
        case "foo bar":
            for fi12 := 0; fi12 < 15; fi12++ {
                for fi13 := 0; fi13 < 13; fi13++ {
                    delete(v2, 65535)
                    var v14 uint32 = 4294967295
                    fmt.Println(v1, fi12)
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        case "":
            _ = len(v2)
        }
        _ = v9
    } else {
        var v15 uint64 = 42
        _ = v15
    }
    var v16 uint32 = 2
    if true {
        v3 = 42
        v17 := GenS2{F0: -2147483648}
        switch "foo bar" {
        case "\n\t":
            v1 = v1
        case "hello":
            delete(v2, v16)
        case "":
            v4 = (^(v4))
        default:
            if (127 <= -128) {
                v18 := GenS2{F0: (^(v3))}
                v4 -= v4
                _ = v18
            }
        }
        _ = v17
    }
    _ = len(v2)
    v1 = (0.1 < 0.1)
    v19 := GenS3{F0: "a", F1: 455}
    v20 := GenS2{F0: (v3 / (((v3 * v3)) | 1))}
    var v21 int8 = -1
    if v1 {
        var v22 float32 = float32((-2.25 - ((float64(3.14159) / 1e10) / -2.25)))
        fmt.Println(v3, v22)
        if (2 < 2) {
            v22 = v22
            _ = len(v5)
        }
        _ = v22
    } else {
        v23 := GenS2{F0: (v3 << ((4294967295) & 63))}
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 int8 = v21
            var v26 uint8 = uint8(v16)
            v3 = v3
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 string = (("hello" + "hello") + "foo bar")
            v29 := GenS4{F0: (fi27 << ((1) & 63)), F1: uint64(v4)}
            var v30 uint32 = 306
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = v23
    }
    var v31 int64 = -9223372036854775808
    if (((("hello" + "hello") + "\n\t") != "") || (565 == 32767)) {
        if false {
            v32 := GenS5{F0: -351, F1: true, F2: v21}
            fmt.Println(v3)
            for fi33 := 0; fi33 < 8; fi33++ {
                if len(v5) > 0 {
                    v5[0] = (((("foo bar" + ("\n\t" + ("hello" + (("hello" + "foo bar") + "a")))) + "a") + "hello") + "")
                }
                _ = fi33
            }
            _ = v32
        } else {
            fmt.Println(v3, v1)
            v34 := GenS6{F0: 2, F1: (v16 - 227)}
            _ = v34
        }
        var v35 uint64 = 0
        _ = v35
    }
    var v36 int32 = (v3 + (v3 - (v3 & (v3 & (v3 / ((v3) | 1))))))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v5
}
