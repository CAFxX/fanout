package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }

func main() {
    v1 := []int{245, -9223372036854775808, 127, -9223372036854775808}
    v2 := []string{"a"}
    var v3 bool = ((0.0 > 3.14159) || true)
    var v4 int32 = -1
    var v5 string = (("\n\t" + "\n\t") + "foo bar")
    for fi6 := 0; fi6 < 14; fi6++ {
        v4 = 2147483647
        var v7 uint = 65535
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(v3, v4)
            v3 = true
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int32 = v4
        for fi11 := 0; fi11 < 8; fi11++ {
            if v3 {
                if len(v2) > 0 {
                    v2[0] = v5
                }
                v5 = "foo bar"
                var v12 uint32 = 157
                var v13 int32 = (^(v10))
                var v14 int = 1
                var v15 uint8 = 77
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi11
        }
        switch 4294967295 {
        case 32767:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 4294967295:
            var v16 int64 = 1
            _ = v16
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi9
        _ = v10
    }
    switch v5 {
    case "":
        if ((0 >= -120) && (202 <= 2)) {
            switch ((v5 + ((("foo bar" + "a") + ((v5 + "a") + (((v5 + v5) + (v5 + v5)) + v5))) + (v5 + (((v5 + v5) + v5) + "")))) + v5) {
            case "a":
                var v17 int16 = -1
                _ = v17
            }
            fmt.Println(v3, v4, v5)
            v5 = v5
        }
    default:
        switch 1 {
        case 127:
            var v18 int16 = 0
            _ = v18
        case 652:
            fmt.Println(v4, v3, v5)
        }
    }
    var v19 int64 = 42
    v20 := GenS1{F0: 42, F1: 42, F2: 2147483648}
    var v21 bool = (("" + v5) <= (("\n\t" + (v5 + v5)) + v5))
    _ = len(v1)
    v22 := GenS1{F0: 42, F1: -606, F2: 192}
    v1 = append(v1, -2147483648)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
