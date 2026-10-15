package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 127
    v3 := []string{"a"}
    var v4 int16 = -32768
    var v5 int16 = v4
    v5 = -438
    var v6 uint16 = 32767
    fmt.Println(v2, v5)
    fmt.Println(v6, v4)
    v2 += v2
    var v7 bool = false
    v8 := GenS1{F0: 42, F1: -1, F2: 2}
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 string = "foo bar"
        switch (^(fi9)) {
        case -467:
            v11 := GenS1{F0: 4294967295, F1: fi9, F2: 0}
            _ = v11
        case 32768:
            if (!(3.14159 == float32(0.0))) {
                v12 := GenS1{F0: 255, F1: fi9, F2: 210}
                if (v7 && (-1 != -9223372036854775808)) {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        var v14 int = 9223372036854775807
                        var v15 float32 = -2.25
                        if (true || (42 > 4294967295)) {
                            fmt.Println(v5, v5, v10)
                            v6 = v6
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    if (!false) {
                        v7 = false
                        for fi16 := 0; fi16 < 11; fi16++ {
                            fmt.Println(v10, v7, v7)
                            var v17 uint = 65535
                            fmt.Println(v17)
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        switch (fi9 >> ((102) & 63)) {
                        case 128:
                            var v18 int16 = v4
                            _ = v18
                        default:
                            v19 := v1[v10]
                            _ = v19
                        }
                    }
                    v20 := GenS2{F0: 0}
                    _ = v20
                }
                switch fi9 {
                case -168:
                    _ = len(v3)
                case 32767:
                    v21 := GenS2{F0: 820}
                    _ = v21
                case -699:
                    v22 := GenS2{F0: 1}
                    _ = v22
                }
                _ = v12
            } else {
                if len(v3) > 0 {
                    v3[0] = v10
                }
                fmt.Println(v10, v7, v10)
            }
        case 127:
            if v7 {
                v10 = (v10 + v10)
            }
        }
        if false {
            v2 = v2
        }
        _ = fi9
        _ = v10
    }
    v23 := GenS3{F0: uint64(v6), F1: 128, F2: (("a" + "foo bar") + "")}
    fmt.Println(v6, v5, v6)
    var v24 uint8 = 127
    v7 = (v7 && true)
    v6 += (^(v6))
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(v5)
        v5 -= 1000
        fmt.Println(v4, v4)
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
