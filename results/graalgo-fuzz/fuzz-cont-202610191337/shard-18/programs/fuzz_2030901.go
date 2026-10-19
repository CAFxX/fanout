package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint64 = 932
    v2 := make(map[int]int)
    var v3 float32 = 3.14159
    var v4 int32 = 0
    var v5 int32 = 367
    v6 := make(map[string]int)
    switch "hello" {
    case "foo bar":
        v5 = 42
    }
    v7 := GenS1{F0: 42}
    var v8 int = 32768
    v9 := GenS1{F0: 1000}
    v2[v8] = 127
    for fi10 := 0; fi10 < 19; fi10++ {
        switch "hello" {
        case "":
            if (3.14159 == float64((0.1 - v3))) {
                v2[202] = 52
                v8 = 32767
                var v11 uint64 = v1
                _ = v11
            }
        default:
            fmt.Println(fi10, v5)
        }
        _ = fi10
    }
    v3 -= 1.5
    v12 := GenS1{F0: 1000}
    var v13 int16 = 1000
    for fi14 := 0; fi14 < 7; fi14++ {
        if (!true) {
            var v15 bool = (65535 <= 56)
            var v16 bool = (float64(v3) < 0.0)
            _ = v15
            _ = v16
        }
        v13 = (v13 << ((18446744073709551615) & 63))
        _ = fi14
    }
    fmt.Println(v5, v13)
    v4 = -1
    switch "\n\t" {
    case "hello":
        v17 := GenS1{F0: v13}
        _ = v17
    case "foo bar":
        v8 += v8
    case "":
        fmt.Println(v1, v8, v4)
    }
    switch 42 {
    case -424:
        var v18 int = v8
        _ = v18
    case 32768:
        for fi19 := 0; fi19 < 3; fi19++ {
            if ((-2.25 == 3.14159) && (2 < v1)) {
                v20 := GenS1{F0: (v13 * v13)}
                _ = v20
            } else {
                var v21 int8 = 0
                var v22 string = "foo bar"
                v21 = 1
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    case 521:
        _ = len(v6)
    }
    var v23 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
