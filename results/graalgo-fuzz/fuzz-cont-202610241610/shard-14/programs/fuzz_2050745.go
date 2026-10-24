package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []uint32{}
    v2 := []uint32{}
    var v3 int = -1
    v4 := make(map[int]string)
    fmt.Println(v3)
    v3 += v3
    var v5 int8 = 1
    for fi6 := 0; fi6 < 16; fi6++ {
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 int32 = -2147483648
                var v9 uint32 = 1
                if false {
                    switch 32767 {
                    case 128:
                        var v10 int64 = -1
                        _ = v10
                    }
                } else {
                    v5 = (v5 - v5)
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            switch (fi6 & (-(fi6))) {
            case 711:
                var v11 int64 = -9223372036854775808
                _ = v11
            case -9223372036854775808:
                switch (("hello" + "") + "hello") {
                case "":
                    var v12 uint32 = 2147483647
                    _ = v12
                case "foo bar":
                    var v13 float32 = 0.1
                    _ = v13
                case "\n\t":
                    var v14 bool = false
                    _ = v14
                default:
                    var v15 int = -498
                    _ = v15
                }
            case 2147483647:
                v5 += int8(v3)
            }
        }
        switch v3 {
        case 0:
            v16 := v4[239]
            _ = v16
        case -484:
            _ = len(v4)
        }
        _ = fi6
    }
    v5 += -110
    v17 := v4[-558]
    fmt.Println(v5, v5)
    v18 := GenS1{F0: false, F1: 2}
    if (true || false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v19 := GenS1{F0: true, F1: 42}
        _ = len(v2)
        _ = v19
    } else {
        var v20 int16 = 1
        if (v3 > (v3 + 32767)) {
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 float32 = 0.0
                v23 := GenS1{F0: (4294967295 < 18446744073709551615), F1: 4294967295}
                _ = fi21
                _ = v22
                _ = v23
            }
            v5 -= -128
        }
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
