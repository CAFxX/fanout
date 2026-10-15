package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "a"
    var v3 uint8 = 0
    v4 := []uint32{1, 4294967295, 799, 2}
    var v5 uint8 = (^(v3))
    var v6 int8 = -1
    v5 = v3
    switch ((("foo bar" + v2) + ("" + v2)) + (v2 + "a")) {
    case "foo bar":
        fmt.Println(v2, v2, v2)
    }
    fmt.Println(v5)
    v7 := GenS1{F0: 42}
    v6 += 127
    fmt.Println(v5)
    var v8 int64 = -902
    var v9 float64 = 1.5
    var v10 int8 = v6
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v11 uint16 = 1
    fmt.Println(v11)
    if (v8 >= v8) {
        var v12 uint64 = 18446744073709551615
        _ = v12
    } else {
        switch v2 {
        case "\n\t":
            var v13 uint64 = 2
            _ = v13
        }
        var v14 float32 = 0.0
        _ = v14
    }
    var v15 float64 = (((float64(((float32((v9 - ((v9 + 3.14159) / ((0.1 + (0.0 - float64(1e10))) + 1.5)))) + ((0.1 - (0.0 + 1e10)) * 1e10)) / (0.1 * 3.14159))) + v9) + v9) / v9)
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS1{F0: fi16}
        for fi18 := 0; fi18 < 20; fi18++ {
            switch fi16 {
            case -830:
                switch v2 {
                case "hello":
                    if (!true) {
                        v5 -= 42
                    } else {
                        var v19 uint16 = (v11 | v11)
                        _ = len(v4)
                        if len(v4) > 0 {
                            v4[0] = 929
                        }
                        v20 := GenS1{F0: fi18}
                        _ = v19
                        _ = v20
                    }
                }
            case 9223372036854775807:
                var v21 int16 = -471
                _ = v21
            case -9223372036854775808:
                var v22 float32 = 0.0
                _ = v22
            }
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    var v23 uint = 1
    switch v2 {
    case "foo bar":
        fmt.Println(v2, v23)
    default:
        var v24 int = 32767
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
