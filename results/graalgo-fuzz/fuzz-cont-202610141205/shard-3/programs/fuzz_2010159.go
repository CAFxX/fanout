package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    v1 := []int{32767, 4294967295, 32767, 42}
    v2 := []string{"a", "foo bar"}
    var v3 int64 = 1
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    v3 = (v3 << ((319) & 63))
    if (true && (32767 > 127)) {
        switch 127 {
        case -9223372036854775808:
            var v6 bool = (42 < 1)
            _ = v6
        case -935:
            fmt.Println(v3)
        default:
            if len(v2) > 0 {
                v2[0] = ((("foo bar" + ("a" + "\n\t")) + "\n\t") + "")
            }
        }
    } else {
        switch 2147483647 {
        case 0:
            fmt.Println(v3)
        case 42:
            _ = len(v5)
        case -77:
            if true {
                fmt.Println(v3)
                v3 = 0
                fmt.Println(v3)
            } else {
                v7 := GenS1{F0: 115, F1: 42}
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 uint64 = 1
                    _ = fi8
                    _ = v9
                }
                _ = v7
            }
        default:
            v3 -= (v3 % (((v3 >> ((65535) & 63))) | 1))
        }
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS2{F0: (fi10 << ((42) & 63)), F1: 0}
        v3 -= (v3 + v3)
        _ = fi10
        _ = v11
    }
    v3 += v3
    var v12 float64 = 1e10
    var v13 bool = true
    var v14 uint16 = 0
    switch "" {
    case "a":
        switch -819 {
        case -1:
            if v13 {
                var v15 int64 = v3
                var v16 int8 = 100
                var v17 int32 = 1
                _ = v15
                _ = v16
                _ = v17
            }
        case 127:
            v1 = append(v1, -9223372036854775808)
        }
    }
    v12 += (0.0 / ((float64(1.5) + (((((((((0.1 / 1e10) + v12) * 0.0) / float64((1.5 / 3.14159))) / v12) * 3.14159) + v12) - v12) + ((float64(0.1) - v12) * (1.5 - (((3.14159 + v12) + ((v12 - (v12 * v12)) / v12)) + (float64((0.0 - ((1e10 - 1.5) * 0.0))) * (v12 - 1e10))))))) - v12))
    _ = len(v2)
    _ = len(v4)
    switch "" {
    case "foo bar":
        v3 += v3
    case "hello":
        var v18 int64 = v3
        _ = v18
    default:
        _ = len(v5)
    }
    switch "" {
    case "foo bar":
        v14 += v14
    case "hello":
        var v19 int64 = (v3 / (((v3 + 955)) | 1))
        _ = v19
    }
    var v20 uint = 255
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
