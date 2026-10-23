package main

import (
    "fmt"
)


func main() {
    var v1 float32 = 0.1
    var v2 uint = 132
    v3 := []uint32{2147483647, 42, 0, 4294967295}
    var v4 uint64 = 18446744073709551615
    v4 = 2
    var v5 uint16 = 255
    if ((false || (((("hello" + "hello") + "\n\t") + (("foo bar" + "a") + "hello")) < (("" + "") + ("foo bar" + "a")))) || true) {
        var v6 uint16 = v5
        v4 = 42
        _ = v6
    } else {
        var v7 uint32 = 611
        _ = v7
    }
    var v8 float64 = (float64(v1) + 1e10)
    v3 = append(v3, 2147483647)
    v4 += 18446744073709551615
    _ = len(v3)
    if true {
        v3 = append(v3, 39)
        for fi9 := 0; fi9 < 19; fi9++ {
            switch fi9 {
            case 127:
                var v10 int = int(v5)
                _ = v10
            case -163:
                var v11 int16 = 32767
                _ = v11
            default:
                v8 = float64(((float32((float64(v1) * float64(v1))) * v1) * v1))
            }
            _ = fi9
        }
        var v12 float64 = 3.14159
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
