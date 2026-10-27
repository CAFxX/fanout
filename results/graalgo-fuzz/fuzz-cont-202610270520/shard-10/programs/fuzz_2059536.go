package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint32 = 65535
    var v4 int8 = 100
    v5 := GenS1{F0: 1, F1: "", F2: 127}
    v6 := []int{4294967295, -9223372036854775808, 128, 32768}
    v7 := v1["foo bar"]
    switch ("\n\t" + "foo bar") {
    case "hello":
        if (364 < 32767) {
            v8 := GenS1{F0: 161, F1: "", F2: 42}
            _ = v8
        } else {
            fmt.Println(v4)
            var v9 int32 = 42
            _ = v9
        }
    case "foo bar":
        var v10 uint = 656
        _ = v10
    case "":
        switch 2147483647 {
        case 128:
            _ = len(v1)
        case -649:
            var v11 string = v7
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 17; fi12++ {
            if ((1 < 65535) || true) {
                fmt.Println(v3, v4)
            }
            var v13 uint = 0
            _ = fi12
            _ = v13
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v7)
    if len(v6) > 0 {
        v6[0] = 32767
    }
    if false {
        if true {
            fmt.Println(v7)
            v7 = v7
            v7 = ((((v7 + "") + (v7 + v7)) + "hello") + ((v7 + v7) + ""))
        } else {
            v4 += 1
        }
        if true {
            v14 := GenS2{F0: 0, F1: 1, F2: 32767}
            v7 = "hello"
            v3 = (v3 ^ v3)
            _ = v14
        }
        v7 = v7
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        switch "" {
        case "":
            delete(v2, (v7 + "hello"))
        case "a":
            switch fi15 {
            case -1:
                v4 -= 1
            case 295:
                v16 := GenS2{F0: 1, F1: -2147483648, F2: 1000}
                _ = v16
            case -2147483648:
                var v17 float64 = (1e10 * (3.14159 / ((3.14159 + 0.1) / 0.1)))
                _ = v17
            }
        case "foo bar":
            var v18 int32 = 0
            _ = v18
        }
        var v19 uint16 = 2
        _ = fi15
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
