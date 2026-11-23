package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }

func main() {
    var v1 uint = 42
    v2 := make(map[int]int)
    var v3 float32 = (3.14159 + (float32(1e10) + 3.14159))
    var v4 float64 = 3.14159
    switch 0 {
    case -9223372036854775808:
        var v5 int = 759
        _ = v5
    case -1:
        switch -2147483648 {
        case -655:
            v6 := GenS1{F0: (v1 + 0), F1: 2147483647, F2: -481}
            _ = v6
        case 32768:
            v1 += v1
        }
    case -2147483648:
        for fi7 := 0; fi7 < 19; fi7++ {
            v1 += 65535
            var v8 int8 = -113
            _ = fi7
            _ = v8
        }
    default:
        fmt.Println(v1, v4, v3)
    }
    switch "" {
    case "\n\t":
        v1 -= (v1 * v1)
    }
    var v9 int32 = -2147483648
    var v10 int64 = -9223372036854775808
    var v11 int64 = (v10 / (((v10 ^ (v10 | v10))) | 1))
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint16 = 734
        if true {
            switch "foo bar" {
            case "":
                var v14 uint64 = uint64(v1)
                _ = v14
            case "\n\t":
                v1 = v1
            case "hello":
                if (!(uint32(v1) <= 1)) {
                    var v15 uint8 = 1
                    var v16 string = ""
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{(-(fi12)), 32767, fi12}
                    _ = v17
                }
            default:
                v18 := GenS1{F0: (^(v1)), F1: v9, F2: (fi12 + fi12)}
                _ = v18
            }
            if ((v1 - (v1 + v1)) > 42) {
                v11 -= v10
                fmt.Println(v9, fi12, v1)
            }
            fmt.Println(v9, fi12)
        } else {
            if false {
                v19 := GenS1{F0: (^(v1)), F1: int32(v1), F2: 127}
                v20 := GenS1{F0: 18446744073709551615, F1: 175, F2: (fi12 / (((fi12 << ((42) & 63))) | 1))}
                _ = v19
                _ = v20
            } else {
                fmt.Println(v3, v13)
                v21 := []int{fi12, 127}
                _ = v21
            }
        }
        _ = fi12
        _ = v13
    }
    var v22 string = "a"
    if (int8(v9) < -19) {
        var v23 uint = 4294967295
        var v24 string = ((v22 + v22) + ("" + "\n\t"))
        _ = v23
        _ = v24
    } else {
        v25 := []int{}
        v26 := GenS1{F0: 18446744073709551615, F1: v9, F2: 4294967295}
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 1, F1: v9, F2: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
