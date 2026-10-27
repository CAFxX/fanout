package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }

func main() {
    var v1 int = -346
    v2 := make(map[string]uint32)
    var v3 uint = 255
    var v4 uint16 = 65535
    var v5 int64 = -1
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 0, F1: (v1 - -1), F2: 122}
        _ = fi6
        _ = v7
    }
    if false {
        v4 = v4
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            v5 += v5
            v5 += v5
            _ = fi8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            v2[""] = 65535
            _ = fi9
        }
        if true {
            var v10 int8 = -128
            if false {
                fmt.Println(v1)
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int16 = 926
                    delete(v2, (("a" + "a") + "a"))
                    _ = fi11
                    _ = v12
                }
            } else {
                v3 += v3
            }
            switch v1 {
            case 579:
                v1 -= v1
            case 127:
                v13 := []int{(v1 << ((uint(v4)) & 63)), v1, 9223372036854775807}
                _ = v13
            default:
                v14 := GenS1{F0: 2, F1: (v1 * (v1 * (^(v1)))), F2: -1}
                _ = v14
            }
            _ = v10
        } else {
            if false {
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 string = ""
                    fmt.Println(fi15, v1)
                    v17 := []int{(^(fi15))}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                fmt.Println(v4, v3)
            } else {
                var v18 bool = false
                _ = v18
            }
            var v19 int32 = 2147483647
            _ = v19
        }
    }
    v1 = v1
    v4 -= v4
    var v20 string = "hello"
    if true {
        switch ((v20 + "hello") + "\n\t") {
        case "\n\t":
            var v21 uint8 = 99
            _ = v21
        default:
            var v22 int32 = -2147483648
            _ = v22
        }
    }
    v20 = "foo bar"
    if (false || true) {
        v20 = v20
        v5 -= v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
