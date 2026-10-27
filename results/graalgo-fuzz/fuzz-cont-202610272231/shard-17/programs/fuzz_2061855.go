package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    var v1 int16 = 1
    var v2 float32 = float32(float64(float32((-2.25 - float64((0.0 - 0.0))))))
    var v3 uint64 = 616
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    if (false || true) {
        var v6 float64 = (1.5 * (float64(3.14159) * 0.1))
        v1 += (v1 >> ((255) & 63))
        _ = v6
    } else {
        var v7 uint8 = 98
        v3 -= v3
        _ = v7
    }
    delete(v5, "\n\t")
    fmt.Println(v1, v1)
    if (!(-128 > 0)) {
        switch ("hello" + "") {
        case "":
            if true {
                v3 = (v3 + v3)
            }
        case "hello":
            v8 := GenS1{F0: 9223372036854775807, F1: -387, F2: 42}
            _ = v8
        }
        v9 := GenS2{F0: 2147483648, F1: 1, F2: 65535}
        v3 += v3
        _ = v9
    }
    var v10 int = -7
    v2 = v2
    fmt.Println(v3)
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v2, v1, v1)
        v12 := GenS3{F0: false, F1: ("hello" + "")}
        fmt.Println(v2)
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v10)
    var v13 int8 = 28
    v13 = v13
    v14 := GenS3{F0: false, F1: "\n\t"}
    v3 -= (v3 >> ((1) & 63))
    if true {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS3{F0: true, F1: ""}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS3{F0: true, F1: "a"}
            var v19 uint64 = 42
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := GenS2{F0: 2, F1: -2147483648, F2: 32767}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 bool = false
            _ = fi21
            _ = v22
        }
        var v23 string = "a"
        v23 = v23
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
