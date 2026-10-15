package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0(p0 uint32, p1 int8) uint8 {
    var v7 int8 = p1
    var v8 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int16 = 1000
    var v2 string = "a"
    var v3 uint64 = 2
    v4 := []string{"foo bar"}
    v5 := make(map[int]string)
    var v6 bool = true
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 9223372036854775807, F1: (v1 | -761), F2: 2147483648}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 = (v3 << ((191) & 63))
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS2{F0: 843, F1: 2}
    for fi13 := 0; fi13 < 20; fi13++ {
        if (1.5 != 3.14159) {
            var v14 uint8 = 88
            _ = v14
        }
        _ = fi13
    }
    var v15 float64 = -2.25
    v6 = (v6 || (false || v6))
    for fi16 := 0; fi16 < 20; fi16++ {
        delete(v5, (-(fi16)))
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 string = "foo bar"
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v2)
    switch (v2 + (v2 + "\n\t")) {
    case "":
        var v19 uint = 4294967295
        _ = v19
    }
    v20 := GenS3{F0: (1 < 1), F1: v1}
    switch (v2 + "a") {
    case "\n\t":
        v21 := GenS1{F0: 1, F1: 0, F2: 2147483647}
        _ = v21
    case "foo bar":
        fmt.Println(v2, v6, v6)
    case "hello":
        if ((0.0 / (0.1 + (-2.25 * ((3.14159 - 3.14159) - v15)))) > v15) {
            v1 += 203
            v3 -= 425
            v15 += v15
        } else {
            var v22 int64 = -1
            switch v2 {
            case "hello":
                v22 = v22
            case "foo bar":
                v23 := GenS2{F0: 2147483648, F1: genFn0(2, 42)}
                _ = v23
            case "\n\t":
                var v24 int64 = v22
                _ = v24
            default:
                for fi25 := 0; fi25 < 15; fi25++ {
                    if v6 {
                        v26 := GenS1{F0: (v22 ^ (v22 ^ (v22 % ((42) | 1)))), F1: (v1 % ((-32768) | 1)), F2: 4294967295}
                        _ = v26
                    } else {
                        var v27 bool = v6
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 19; fi28++ {
                        v29 := GenS2{F0: 42, F1: 128}
                        var v30 uint16 = 65535
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                }
            }
            _ = v22
        }
    }
    v4 = append(v4, v2)
    var v31 int64 = 104
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
