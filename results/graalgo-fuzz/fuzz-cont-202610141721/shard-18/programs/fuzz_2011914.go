package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = ((1.5 / 1.0) * -2.25)
    var v3 uint32 = 360
    v4 := GenS1{F0: 18446744073709551615, F1: 214}
    var v5 int64 = -1
    var v6 int = -2147483648
    var v7 uint32 = 716
    var v8 bool = (1 <= -412)
    v7 = (^(v7))
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v8)
        _ = fi9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: 4294967295, F1: 255}
        var v12 bool = (!v8)
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v8)
    var v13 uint8 = 255
    fmt.Println(v5, v6)
    if (v2 >= (1e10 + 1.5)) {
        var v14 bool = v8
        _ = v14
    }
    v5 += v5
    for fi15 := 0; fi15 < 16; fi15++ {
        v6 = v6
        v6 += (-(fi15))
        _ = fi15
    }
    if (v6 >= v6) {
        switch "" {
        case "\n\t":
            v16 := v1["foo bar"]
            _ = v16
        default:
            v6 -= -907
        }
        if ((v8 || v8) && (549 == 42)) {
            v1[""] = 2
            if (901 < -32768) {
                v8 = v8
                if v8 {
                    switch (^(v6)) {
                    case 2147483647:
                        var v17 int32 = 1
                        _ = v17
                    case -2147483648:
                        switch "hello" {
                        case "foo bar":
                            v18 := GenS2{F0: 12}
                            _ = v18
                        }
                    default:
                        switch "foo bar" {
                        case "a":
                            v19 := GenS2{F0: 1}
                            _ = v19
                        case "\n\t":
                            var v20 float32 = (v2 + 1e10)
                            _ = v20
                        default:
                            var v21 string = "a"
                            _ = v21
                        }
                    }
                }
            }
        }
        var v22 int8 = -52
        _ = v22
    }
    v23 := GenS2{F0: 404}
    var v24 int16 = 829
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
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
