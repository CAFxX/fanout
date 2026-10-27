package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }

func main() {
    var v1 int16 = 1000
    var v2 int16 = 338
    var v3 bool = (0.1 < -2.25)
    v4 := []string{"\n\t"}
    var v5 float32 = ((0.1 - 1e10) - 1e10)
    var v6 string = "a"
    switch v6 {
    case "a":
        switch v6 {
        case "\n\t":
            for fi7 := 0; fi7 < 6; fi7++ {
                fmt.Println(v1)
                fmt.Println(v6, v3, fi7)
                var v8 uint8 = 128
                _ = fi7
                _ = v8
            }
        case "foo bar":
            v3 = (v3 || (!v3))
        }
    case "":
        var v9 int64 = 0
        _ = v9
    default:
        for fi10 := 0; fi10 < 18; fi10++ {
            v3 = (true || (!v3))
            var v11 float64 = 0.0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi10
            _ = v11
        }
    }
    v5 += v5
    v1 += v1
    var v12 uint32 = 0
    var v13 uint16 = 32767
    if v3 {
        switch (v6 + "") {
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = (v6 + ("" + v6))
            }
        case "hello":
            if (1.5 >= 1.5) {
                var v14 int8 = 42
                _ = v14
            }
        }
    } else {
        v6 = v6
        var v15 float64 = 3.14159
        _ = v15
    }
    v2 = (v1 / ((0) | 1))
    var v16 int = 32767
    v17 := GenS1{F0: (v5 <= (float32(0.0) / 3.14159)), F1: false, F2: v16}
    if (9223372036854775807 > 0) {
        var v18 uint8 = 128
        var v19 uint32 = v12
        _ = v18
        _ = v19
    }
    switch (v16 ^ (v16 - (v16 | v16))) {
    case -1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        var v20 int32 = 546
        _ = v20
    }
    v16 = 32768
    var v21 uint32 = v12
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 int64 = 207
    var v23 uint8 = 2
    var v24 float64 = float64((3.14159 - v5))
    fmt.Println(v6, v21, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
