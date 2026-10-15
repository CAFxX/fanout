package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64; F2 string }

func main() {
    var v1 float32 = (0.1 / 1.0)
    v2 := []string{"\n\t"}
    var v3 uint64 = 0
    var v4 uint8 = uint8(v3)
    v5 := make(map[int]string)
    if (513 >= 32767) {
        var v6 int16 = 1000
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int64 = 0
            fmt.Println(v3, v6)
            _ = fi7
            _ = v8
        }
        var v9 int8 = -98
        _ = v6
        _ = v9
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint8 = (v4 * (v4 - 255))
                v13 := GenS1{F0: ""}
                var v14 int32 = -1
                fmt.Println(v1, v3, fi10)
                v4 += v12
                var v15 int8 = 1
                var v16 int16 = 1
                fmt.Println(v4, v12, v14)
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 uint8 = 42
            var v18 uint32 = 4
            _ = fi10
            _ = v17
            _ = v18
        }
    }
    var v19 string = "foo bar"
    switch (v19 + (v19 + "foo bar")) {
    case "foo bar":
        fmt.Println(v1)
    }
    if len(v2) > 0 {
        v2[0] = v19
    }
    _ = len(v5)
    v4 -= v4
    fmt.Println(v3)
    v5[-994] = (v19 + v19)
    fmt.Println(v19, v1)
    fmt.Println(v3)
    for fi20 := 0; fi20 < 14; fi20++ {
        v21 := GenS1{F0: v19}
        var v22 uint32 = 771
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 float32 = ((1e10 + (1.5 + v1)) + 0.1)
    v23 += v1
    var v24 int16 = 0
    for fi25 := 0; fi25 < 18; fi25++ {
        v2 = append(v2, "\n\t")
        _ = fi25
    }
    if (128 < 4294967295) {
        v1 -= v23
        v24 = (v24 ^ (v24 & v24))
    } else {
        switch v19 {
        case "\n\t":
            _ = len(v5)
        case "a":
            v24 = 1
        case "":
            var v26 int32 = 0
            _ = v26
        }
        switch v19 {
        case "foo bar":
            if true {
                var v27 uint = 65535
                v1 -= 0.0
                var v28 int = 128
                _ = v27
                _ = v28
            } else {
                v29 := GenS2{F0: v19, F1: -9223372036854775808, F2: v19}
                var v30 uint8 = v4
                _ = v29
                _ = v30
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
