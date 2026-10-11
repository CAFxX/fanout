package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int8 = 42
    var v2 float64 = 3.14159
    v3 := make(map[uint32]bool)
    v4 := []string{"\n\t", "\n\t"}
    v4 = append(v4, ("\n\t" + ""))
    var v5 int = 9223372036854775807
    v3[42] = (!true)
    v5 -= v5
    v2 = v2
    if (-803 != -602) {
        var v6 string = "a"
        delete(v3, 2147483648)
        _ = v6
    }
    v5 = v5
    v1 -= 1
    var v7 uint = 101
    if (219 != (v5 * (^(v5)))) {
        switch v5 {
        case -266:
            v7 += v7
        case -750:
            var v8 uint64 = 42
            _ = v8
        default:
            if true {
                v9 := GenS1{F0: 782}
                _ = v9
            } else {
                fmt.Println(v5)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
