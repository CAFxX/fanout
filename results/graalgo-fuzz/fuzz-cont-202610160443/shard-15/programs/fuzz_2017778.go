package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    v1 := []int{32768, -240, -9223372036854775808}
    var v2 string = ("foo bar" + "foo bar")
    var v3 int8 = 42
    v4 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808, F2: 247}
    var v5 uint32 = 4294967295
    fmt.Println(v2, v3)
    if false {
        v6 := GenS2{F0: (v5 ^ 103), F1: 0}
        switch v2 {
        case "hello":
            v3 = (-(v3))
        case "":
            v5 += v5
        }
        fmt.Println(v3, v2, v3)
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        if false {
            v8 := GenS1{F0: 524, F1: -1, F2: 0}
            if len(v1) > 0 {
                v1[0] = (fi7 & (fi7 | 42))
            }
            v2 = (v2 + ((v2 + "\n\t") + v2))
            _ = v8
        }
        fmt.Println(fi7, v2)
        switch (fi7 / ((int(v5)) | 1)) {
        case -374:
            v5 += (v5 - (v5 >> ((923) & 63)))
        case 132:
            fmt.Println(v5, v5, v2)
        }
        _ = fi7
    }
    var v9 float32 = 0.0
    v9 += (((1e10 - 1.5) + (v9 * 1.5)) - float32((3.14159 + (0.0 * (-2.25 - 0.1)))))
    switch (((v2 + v2) + v2) + v2) {
    case "\n\t":
        var v10 uint16 = 255
        _ = v10
    case "":
        _ = len(v1)
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
