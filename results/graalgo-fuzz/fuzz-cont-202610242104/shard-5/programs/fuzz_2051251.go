package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string; F1 bool; F2 int16 }

func main() {
    v1 := []uint32{332}
    v2 := make(map[string]string)
    v3 := []int{-9223372036854775808, -804, 32768}
    v4 := []int{127}
    var v5 float64 = 3.14159
    var v6 int16 = 42
    v5 += 1e10
    fmt.Println(v5)
    v2["\n\t"] = (("foo bar" + "hello") + "")
    v2[""] = "a"
    v7 := GenS1{F0: 32768, F1: -2147483648, F2: 65535}
    if (false || false) {
        v8 := GenS2{F0: 42}
        switch "" {
        case "foo bar":
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 string = ("\n\t" + "hello")
                _ = fi9
                _ = v10
            }
        case "hello":
            var v11 float32 = 3.14159
            _ = v11
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            v5 = v5
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(fi12, v6)
                v14 := GenS3{F0: 1, F1: 1}
                _ = fi13
                _ = v14
            }
            v15 := GenS4{F0: "\n\t", F1: ((42 < 65535) && true), F2: -1}
            _ = fi12
            _ = v15
        }
        _ = v8
    }
    v3 = append(v3, 77)
    fmt.Println(v6, v6)
    v6 -= v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
