package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int; F2 string }
func genFn0() string {
    v7 := []int{2147483647}
    if (float32((0.1 - 1e10)) < (0.0 / 1e10)) {
        var v8 int64 = 42
        _ = v8
    }
    if ((0 == -723) && (!(862 >= 0))) {
        v7 = append(v7, 32767)
        var v9 bool = false
        _ = v9
    } else {
        switch "foo bar" {
        case "":
            if true {
                _ = len(v7)
                v10 := GenS1{F0: 65535}
                v11 := GenS1{F0: 537}
                _ = v10
                _ = v11
            } else {
                var v12 uint32 = 1
                var v13 bool = true
                fmt.Println(v13)
                _ = v12
                _ = v13
            }
        case "a":
            var v14 int16 = 1000
            _ = v14
        case "\n\t":
            var v15 float32 = -2.25
            _ = v15
        }
        var v16 float64 = 1e10
        var v17 int8 = -128
        _ = v16
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 10; fi18++ {
            if (65535 < 4294967295) {
                v19 := GenS1{F0: 65535}
                v20 := GenS2{F0: 128, F1: 621, F2: "hello"}
                var v21 bool = true
                _ = v19
                _ = v20
                _ = v21
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi18
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v22 uint64 = 265
        _ = v22
    }
    if (true && (1000 > 1000)) {
        v7 = append(v7, 32767)
        var v23 int64 = 55
        _ = v23
    }
    _ = v7
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = (1.5 / (float32((float64(((1e10 * -2.25) + -2.25)) * 0.0)) - 3.14159))
    var v3 uint32 = 4294967295
    var v4 int16 = -964
    var v5 bool = false
    v6 := []string{}
    if v5 {
        for fi24 := 0; fi24 < 10; fi24++ {
            switch (fi24 >> ((0) & 63)) {
            case -427:
                v2 = v2
            case -1:
                v25 := GenS1{F0: 1}
                _ = v25
            case 9223372036854775807:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                if (0 >= -1) {
                    v2 -= 1.5
                    fmt.Println(fi24, v2)
                } else {
                    v5 = (34 == 127)
                    v2 = float32(0.0)
                    if (-2.25 == -2.25) {
                        var v26 int64 = -1
                        _ = v26
                    } else {
                        v27 := GenS1{F0: 42}
                        _ = v27
                    }
                }
            }
            _ = fi24
        }
        v4 += v4
        switch ("\n\t" + "hello") {
        case "\n\t":
            v28 := GenS1{F0: 42}
            _ = v28
        default:
            var v29 uint8 = 128
            _ = v29
        }
    }
    v4 = v4
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(v5, fi30)
        _ = fi30
    }
    v1[v3] = v5
    v5 = v5
    for fi31 := 0; fi31 < 8; fi31++ {
        switch "foo bar" {
        case "hello":
            v32 := GenS2{F0: v3, F1: (fi31 | fi31), F2: "\n\t"}
            _ = v32
        case "foo bar":
            var v33 uint8 = 2
            _ = v33
        case "\n\t":
            v4 = -1
        }
        _ = fi31
    }
    var v34 float64 = 1e10
    fmt.Println(v5, v3, v34)
    var v35 string = "hello"
    if v5 {
        var v36 bool = false
        _ = v36
    } else {
        for fi37 := 0; fi37 < 9; fi37++ {
            if (v5 && (!false)) {
                switch fi37 {
                case -431:
                    var v38 uint64 = 988
                    _ = v38
                case 1:
                    var v39 int64 = -9223372036854775808
                    _ = v39
                default:
                    var v40 int8 = -112
                    _ = v40
                }
                var v41 uint32 = (v3 & (v3 % ((v3) | 1)))
                _ = v41
            }
            v42 := v1[(v3 | 4294967295)]
            _ = fi37
            _ = v42
        }
    }
    delete(v1, (v3 << ((65535) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
