package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }

func main() {
    var v1 int = 128
    var v2 float64 = ((0.1 / 1e10) - (0.0 * 1e10))
    var v3 uint64 = 1
    v4 := make(map[string]string)
    var v5 uint8 = 128
    if (2147483648 <= 0) {
        fmt.Println(v5, v5)
    } else {
        v2 += 0.0
    }
    v6 := []int{v1}
    switch (("hello" + "a") + ("hello" + ("" + ""))) {
    case "foo bar":
        var v7 float64 = (v2 * (((v2 / (float64((-2.25 / 1.5)) - 1.5)) / 1.5) / (v2 + 1e10)))
        _ = v7
    case "\n\t":
        v3 -= v3
    default:
        if (true && true) {
            var v8 int64 = -9223372036854775808
            var v9 int = v1
            _ = v8
            _ = v9
        } else {
            v5 = v5
            v3 = v3
            switch (v1 & v1) {
            case 128:
                _ = len(v6)
            case -9223372036854775808:
                _ = len(v4)
            case -2147483648:
                v4["a"] = "a"
            }
        }
    }
    v10 := GenS1{F0: 0, F1: 42, F2: "\n\t"}
    if true {
        var v11 float32 = (0.1 - 1e10)
        v12 := GenS2{F0: 0}
        _ = v11
        _ = v12
    }
    var v13 uint16 = uint16(v3)
    if false {
        if false {
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS1{F0: v5, F1: 1, F2: "foo bar"}
                _ = fi14
                _ = v15
            }
        } else {
            var v16 int64 = -570
            v13 = 65535
            if (!(false || false)) {
                if true {
                    v17 := GenS3{F0: v16, F1: 1}
                    v13 = (v13 | v13)
                    v18 := GenS3{F0: -216, F1: 0}
                    _ = v17
                    _ = v18
                }
            }
            _ = v16
        }
    }
    _ = len(v6)
    if (v3 >= v3) {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 int = v1
            if true {
                v2 = 1e10
                var v21 uint8 = (^(v5))
                _ = v21
            } else {
                v22 := GenS4{F0: 109}
                v23 := GenS5{F0: false}
                _ = v22
                _ = v23
            }
            if true {
                v20 = (v1 / ((v1) | 1))
            }
            _ = fi19
            _ = v20
        }
        var v24 int8 = 42
        v13 += (v13 & 32767)
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
