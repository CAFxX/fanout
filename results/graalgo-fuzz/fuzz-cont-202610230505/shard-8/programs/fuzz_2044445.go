package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{v1, (v1 | v1), v1}
    v3 := make(map[string]uint32)
    v1 = (v1 + 561)
    for fi4 := 0; fi4 < 18; fi4++ {
        v1 = 42
        if (((262 < 963) || false) && false) {
            v1 = (^(v1))
            v1 = (v1 >> ((65535) & 63))
            if len(v2) > 0 {
                v2[0] = v1
            }
        } else {
            var v5 uint = 351
            v1 = v1
            delete(v3, "a")
            _ = v5
        }
        _ = fi4
    }
    var v6 int = -1
    v6 = 128
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint = 65535
        v6 = fi7
        _ = fi7
        _ = v8
    }
    if (!(false && (!(100 == 100)))) {
        var v9 uint16 = 410
        _ = v9
    }
    if (((4294967295 == 128) && false) && ((866 <= uint(v1)) || true)) {
        switch v6 {
        case 32768:
            var v10 int = 88
            _ = v10
        case 127:
            for fi11 := 0; fi11 < 3; fi11++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v12 := GenS1{F0: 42}
                _ = fi11
                _ = v12
            }
        default:
            switch ("a" + "a") {
            case "":
                fmt.Println(v6, v6)
            case "\n\t":
                v13 := GenS2{F0: 42}
                _ = v13
            }
        }
        var v14 float64 = 1e10
        v14 += float64(1e10)
        _ = v14
    } else {
        fmt.Println(v6, v6)
    }
    v6 = 128
    v1 = v1
    var v15 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v6
}
