package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0() string {
    v5 := GenS1{F0: 1, F1: 215, F2: -9223372036854775808}
    var v6 uint64 = 2
    v6 += (v6 % ((v6) | 1))
    var v7 float32 = (1.5 + 1e10)
    _ = v5
    _ = v6
    _ = v7
    return (("hello" + (("hello" + "\n\t") + "a")) + "a")
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 560
    v3 := make(map[string]bool)
    var v4 float64 = 3.14159
    v2 = (^(v2))
    v3[""] = (2 < 255)
    var v8 uint = 4294967295
    var v9 uint32 = 30
    v10 := GenS1{F0: (v2 | (v2 / (((^(v2))) | 1))), F1: 255, F2: -1}
    fmt.Println(v9, v4, v2)
    fmt.Println(v2)
    v4 = (v4 - 3.14159)
    fmt.Println(v2, v8)
    for fi11 := 0; fi11 < 11; fi11++ {
        switch 32767 {
        case 127:
            v1[(((("a" + "a") + ("foo bar" + "")) + "a") + (("\n\t" + "hello") + ("hello" + "hello")))] = fi11
        case 1:
            v12 := []int{int(v8), fi11, fi11, 128}
            _ = v12
        case 128:
            var v13 uint8 = 42
            _ = v13
        default:
            var v14 uint64 = 42
            _ = v14
        }
        v15 := GenS1{F0: v2, F1: 255, F2: 9223372036854775807}
        _ = fi11
        _ = v15
    }
    if (1 != uint8(v8)) {
        var v16 uint32 = (v9 & (v9 & 907))
        v2 += 2147483647
        v4 += v4
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
