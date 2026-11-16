package main

import (
    "fmt"
)

func genFn0() int32 {
    v4 := []int{}
    var v5 int = -2147483648
    _ = v4
    _ = v5
    return 47
}

func main() {
    var v1 int = 0
    v2 := make(map[int]bool)
    var v3 uint = 4294967295
    fmt.Println(v1, v3)
    v1 -= 1
    var v6 float32 = ((-2.25 + 0.0) + -2.25)
    if true {
        var v7 bool = (!((v6 == ((((v6 + (v6 + 1.5)) - 3.14159) * (1e10 + 1e10)) - 3.14159)) && (83 > 42)))
        _ = v7
    } else {
        if (true || (461 > 32767)) {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 uint64 = 2
                var v10 int32 = 0
                var v11 uint32 = 2
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 uint32 = 2
            _ = v12
        }
    }
    var v13 uint32 = 42
    switch "" {
    case "foo bar":
        var v14 int64 = -922
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int16 = 1
        fmt.Println(v13, v6)
        _ = fi15
        _ = v16
    }
    switch v1 {
    case -9223372036854775808:
        if true {
            var v17 float32 = (v6 * 0.0)
            _ = v17
        }
    case 0:
        v2[v1] = true
    case 32768:
        switch "\n\t" {
        case "":
            fmt.Println(v3)
        default:
            v1 = (v1 >> ((v3) & 63))
        }
    }
    if (("\n\t" < "hello") || true) {
        switch "\n\t" {
        case "hello":
            var v18 int64 = 42
            _ = v18
        }
        fmt.Println(v3, v13)
    } else {
        v3 += (^(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v6
}
