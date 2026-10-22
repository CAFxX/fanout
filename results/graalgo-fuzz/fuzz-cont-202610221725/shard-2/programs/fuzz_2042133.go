package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []string{"", ""}
    var v2 uint = 42
    var v3 uint64 = 42
    v3 += (v3 - 0)
    if ((42 == 2) || true) {
        v2 += 0
        var v4 int64 = -9223372036854775808
        switch "\n\t" {
        case "\n\t":
            v2 = v2
        case "":
            if (69 <= 2) {
                v2 = v2
                fmt.Println(v3, v2)
            } else {
                v2 -= 255
            }
        default:
            var v5 int64 = (v4 + (v4 / (((v4 * 919)) | 1)))
            _ = v5
        }
        _ = v4
    }
    var v6 uint32 = 1
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int = (-(fi7))
        if true {
            switch v8 {
            case -1:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v3 += 18446744073709551615
        }
        _ = fi7
        _ = v8
    }
    v6 += (v6 - v6)
    v2 = 65535
    v2 -= (^(v2))
    v9 := GenS1{F0: uint32(v2)}
    var v10 float32 = (-2.25 + 0.0)
    v2 = 255
    v11 := GenS1{F0: v6}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
