package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 string }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 188
    v3 := []int{-678, int(v2), 272, 2147483647}
    var v4 uint64 = 42
    var v5 uint16 = 1
    var v6 uint = v1
    if ((false || true) || ("\n\t" > "a")) {
        _ = len(v3)
    } else {
        var v7 uint = v6
        for fi8 := 0; fi8 < 3; fi8++ {
            if (53 >= -26) {
                var v9 float32 = 1.5
                switch 2147483647 {
                case 42:
                    v3 = append(v3, fi8)
                case 32768:
                    _ = len(v3)
                }
                _ = v9
            } else {
                var v10 uint64 = uint64(v5)
                switch ("\n\t" + "") {
                case "hello":
                    v3 = append(v3, fi8)
                case "a":
                    for fi11 := 0; fi11 < 15; fi11++ {
                        fmt.Println(v4)
                        _ = fi11
                    }
                }
                _ = v10
            }
            _ = fi8
        }
        v12 := GenS1{F0: "hello"}
        _ = v12
        _ = v7
    }
    v4 = (v4 / ((1) | 1))
    v3 = append(v3, -9223372036854775808)
    v6 = uint(v5)
    switch ("\n\t" + "\n\t") {
    case "a":
        fmt.Println(v5)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v13 := GenS2{F0: -54, F1: ("\n\t" + "a")}
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
