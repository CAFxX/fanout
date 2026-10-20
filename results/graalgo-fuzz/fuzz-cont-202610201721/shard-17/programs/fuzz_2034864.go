package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string }

func main() {
    v1 := []string{"foo bar", (("\n\t" + ("a" + ("a" + "foo bar"))) + "foo bar"), ((("" + "a") + "a") + ("" + ("foo bar" + "hello")))}
    v2 := make(map[uint32]string)
    var v3 uint32 = 276
    v4 := []int{42, -2147483648, 674, 57}
    var v5 bool = true
    var v6 int64 = 9223372036854775807
    var v7 string = ""
    fmt.Println(v6, v6)
    v5 = v5
    v8 := GenS1{F0: v3, F1: (v7 + "a"), F2: -650}
    if (1e10 >= 3.14159) {
        switch 4294967295 {
        case -9223372036854775808:
            v2[v3] = "hello"
        case -2147483648:
            fmt.Println(v3)
        case -413:
            v5 = v5
        default:
            switch 163 {
            case 585:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 2147483647:
                v9 := GenS1{F0: v3, F1: (("a" + (v7 + "")) + ("foo bar" + v7)), F2: 1}
                _ = v9
            default:
                var v10 int16 = 1000
                _ = v10
            }
        }
    } else {
        v11 := GenS2{F0: 0, F1: v3}
        switch v7 {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v12 uint8 = 73
            _ = v12
        }
        _ = v11
    }
    v13 := GenS3{F0: "a"}
    var v14 float64 = (3.14159 * 3.14159)
    var v15 string = v7
    if (18 < 593) {
        fmt.Println(v15, v3)
        v16 := v2[(v3 + v3)]
        _ = v16
    }
    fmt.Println(v14)
    fmt.Println(v5)
    v15 = (v15 + v15)
    if false {
        v15 = v15
    } else {
        v17 := v2[(v3 << ((0) & 63))]
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
