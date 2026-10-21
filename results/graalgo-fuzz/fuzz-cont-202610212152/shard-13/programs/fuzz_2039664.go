package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    v1 := []uint32{1, 4294967295}
    var v2 float32 = ((float32(3.14159) * -2.25) * (-2.25 - -2.25))
    var v3 int64 = -1
    var v4 bool = true
    var v5 int16 = 1
    for fi6 := 0; fi6 < 8; fi6++ {
        if ((v2 >= (float32(float64(1e10)) * (0.0 * (float32(1e10) * 0.0)))) || (fi6 > int(v3))) {
            var v7 uint = 949
            var v8 uint = 5
            switch fi6 {
            case -9223372036854775808:
                v7 -= 255
            default:
                v9 := GenS1{F0: v8}
                _ = v9
            }
            _ = v7
            _ = v8
        }
        var v10 float32 = v2
        _ = fi6
        _ = v10
    }
    if v4 {
        switch "" {
        case "":
            fmt.Println(v4, v5, v3)
        case "foo bar":
            var v11 int32 = 0
            _ = v11
        }
        if false {
            _ = len(v1)
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int = -412
                var v14 float64 = (((3.14159 - (0.0 / -2.25)) * 1.5) - -2.25)
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v2 += (v2 / v2)
        _ = fi15
    }
    var v16 int64 = (v3 | v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v4 {
        switch (((("\n\t" + "") + ("" + "")) + "\n\t") + "") {
        case "a":
            fmt.Println(v2)
        case "foo bar":
            var v17 int16 = v5
            _ = v17
        default:
            v4 = true
        }
        v18 := GenS2{F0: 4294967295, F1: 362}
        _ = v18
    }
    switch ("hello" + "hello") {
    case "hello":
        if (2 == 42) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v19 uint = 2
            var v20 uint32 = 1
            _ = v19
            _ = v20
        } else {
            fmt.Println(v5)
            fmt.Println(v16, v5, v3)
        }
    case "foo bar":
        var v21 uint = 42
        _ = v21
    case "":
        switch 127 {
        case -9223372036854775808:
            var v22 string = "foo bar"
            _ = v22
        case 2147483647:
            var v23 int64 = 9223372036854775807
            _ = v23
        }
    default:
        v1 = append(v1, 227)
    }
    if (337 <= 255) {
        v2 -= float32((float64(v2) / -2.25))
    }
    v16 = v16
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 float32 = v2
            var v27 float32 = 1e10
            switch (fi24 >> ((771) & 63)) {
            case 9223372036854775807:
                v5 = v5
            default:
                v28 := GenS1{F0: 1}
                _ = v28
            }
            v2 = v2
            _ = fi25
            _ = v26
            _ = v27
        }
        var v29 bool = (true && true)
        _ = fi24
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
