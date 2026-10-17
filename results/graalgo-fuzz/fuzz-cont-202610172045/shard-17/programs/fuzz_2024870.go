package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint16 }

func main() {
    v1 := []string{""}
    var v2 int64 = -976
    var v3 int8 = 127
    var v4 uint = 18446744073709551615
    if false {
        v4 -= v4
        v5 := GenS1{F0: v4, F1: 127, F2: 2}
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int = (fi6 | -2147483648)
        var v8 int = fi6
        v7 = 2147483647
        _ = fi6
        _ = v7
        _ = v8
    }
    v1 = append(v1, "hello")
    v2 = (v2 >> ((v4) & 63))
    if (true || true) {
        if true {
            v9 := GenS1{F0: (^(v4)), F1: 628, F2: 42}
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            if len(v1) > 0 {
                v1[0] = ((("foo bar" + "a") + "\n\t") + "")
            }
            _ = fi10
        }
    }
    v3 += v3
    var v11 int8 = v3
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            switch (^(fi12)) {
            case -106:
                var v14 int = (fi12 << ((v4) & 63))
                _ = v14
            case 0:
                var v15 uint32 = 4294967295
                _ = v15
            case 127:
                switch "foo bar" {
                case "foo bar":
                    var v16 int32 = 42
                    _ = v16
                }
            default:
                v2 = (v2 | (v2 * -9223372036854775808))
            }
            _ = fi13
        }
        var v17 int32 = 844
        _ = fi12
        _ = v17
    }
    v11 = 42
    if (0 != 669) {
        v2 = v2
    }
    v2 += (v2 >> ((v4) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 += (v11 % (((v3 - 42)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
