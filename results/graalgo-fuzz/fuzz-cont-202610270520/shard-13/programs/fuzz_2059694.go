package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }

func main() {
    var v1 uint64 = 42
    v2 := []uint32{uint32(v1), 76, 2147483647, 1}
    var v3 int32 = 2147483647
    var v4 int = -2147483648
    var v5 int64 = -1
    fmt.Println(v1)
    var v6 uint64 = 4294967295
    switch (v4 + -73) {
    case 128:
        var v7 uint = 4294967295
        _ = v7
    case 0:
        v8 := GenS1{F0: -810, F1: 866, F2: 347}
        _ = v8
    case 341:
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int8 = 42
            for fi11 := 0; fi11 < 18; fi11++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    if (!true) {
        switch "" {
        case "hello":
            v6 = v1
        }
    } else {
        v5 -= (v5 << ((0) & 63))
        if false {
            var v12 bool = (!true)
            _ = v12
        } else {
            var v13 uint64 = (v6 | v1)
            v14 := GenS1{F0: (v4 / ((v4) | 1)), F1: 4294967295, F2: v6}
            _ = v13
            _ = v14
        }
        if true {
            fmt.Println(v1, v4, v5)
            v2 = append(v2, 486)
        } else {
            var v15 uint32 = 65535
            var v16 uint64 = v1
            _ = v15
            _ = v16
        }
    }
    v3 -= (v3 >> ((18446744073709551615) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("hello" + ((("" + "\n\t") + ("\n\t" + ("" + "foo bar"))) + (("foo bar" + "hello") + "hello"))) {
    case "foo bar":
        v5 -= 9223372036854775807
    default:
        var v17 float64 = (3.14159 / (1.5 - float64(float32(((3.14159 * (0.1 / 3.14159)) * -2.25)))))
        _ = v17
    }
    v6 = 1
    var v18 float32 = 0.1
    var v19 int32 = v3
    var v20 float64 = 3.14159
    var v21 float64 = v20
    v2 = append(v2, 2147483648)
    v22 := GenS1{F0: (v4 << ((65535) & 63)), F1: 255, F2: v6}
    fmt.Println(v5, v6, v6)
    var v23 int64 = 0
    v24 := GenS1{F0: 4294967295, F1: 323, F2: (v1 << ((0) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
