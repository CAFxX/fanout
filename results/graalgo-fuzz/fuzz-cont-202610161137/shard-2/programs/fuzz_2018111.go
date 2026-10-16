package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 1.5
    v3 := make(map[string]bool)
    v2 += v2
    var v4 int = 2147483647
    delete(v1, "foo bar")
    switch (v4 << ((0) & 63)) {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 9; fi5++ {
            if false {
                v6 := GenS1{F0: 1, F1: 2, F2: 2}
                v4 = v4
                _ = v6
            }
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS2{F0: 0}
                var v9 string = "foo bar"
                var v10 int = (fi7 % (((v4 & fi7)) | 1))
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi5
        }
    }
    v2 = v2
    switch (^(v4)) {
    case 32768:
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v2)
            var v12 uint64 = 4294967295
            var v13 int16 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
    case 42:
        delete(v1, "hello")
    default:
        fmt.Println(v2, v2)
    }
    v2 = (v2 / v2)
    var v14 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
