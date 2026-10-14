package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint16 = 65535
    v2 := []int{32767, -135}
    var v3 uint64 = 519
    v4 := make(map[string]string)
    switch 128 {
    case 9223372036854775807:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -9223372036854775808:
        v1 -= v1
    default:
        switch 127 {
        case 9223372036854775807:
            v3 -= (v3 | 18446744073709551615)
        case 1:
            v1 -= (^(v1))
        case -1:
            v3 += v3
        }
    }
    switch 25 {
    case -344:
        var v5 float32 = (1.5 - 1e10)
        _ = v5
    case 32768:
        if false {
            switch 704 {
            case -104:
                var v6 int = 32767
                _ = v6
            case -194:
                for fi7 := 0; fi7 < 15; fi7++ {
                    v4["a"] = "\n\t"
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 uint32 = 4294967295
                        _ = fi8
                        _ = v9
                    }
                    if true {
                        v2 = append(v2, (fi7 & (fi7 + fi7)))
                        var v10 int16 = 1000
                        _ = v10
                    }
                    _ = fi7
                }
            }
        }
    }
    var v11 string = "hello"
    delete(v4, v11)
    v11 = v11
    var v12 uint = 2
    var v13 int64 = 42
    v14 := GenS1{F0: (^(v1)), F1: -566}
    switch 1 {
    case -2147483648:
        switch int(v1) {
        case -2147483648:
            v2 = append(v2, int(v13))
        }
    default:
        switch "foo bar" {
        case "\n\t":
            v12 += (v12 & (v12 / ((uint(v1)) | 1)))
        }
    }
    _ = len(v4)
    v11 = "\n\t"
    if len(v2) > 0 {
        v2[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
