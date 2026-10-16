package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }

func main() {
    var v1 uint64 = 995
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    fmt.Println(v1)
    fmt.Println(v1)
    v1 = 307
    switch 32768 {
    case 1:
        v1 = (^(v1))
    case 935:
        if (468 == -437) {
            for fi4 := 0; fi4 < 15; fi4++ {
                var v5 float64 = 0.1
                _ = fi4
                _ = v5
            }
            v6 := GenS1{F0: -969, F1: 42}
            _ = v6
        } else {
            var v7 uint32 = 403
            var v8 int = -342
            _ = v7
            _ = v8
        }
    }
    v1 = 18446744073709551615
    var v9 int32 = -796
    var v10 bool = ((true || false) || false)
    var v11 string = "hello"
    fmt.Println(v1)
    switch (v11 + "a") {
    case "a":
        fmt.Println(v1, v11)
    case "foo bar":
        if true {
            if false {
                v12 := GenS2{F0: 3, F1: 0}
                _ = v12
            } else {
                _ = len(v2)
                v9 += (v9 | (v9 & v9))
            }
            fmt.Println(v10, v10)
        }
    }
    var v13 uint = 566
    if (3.14159 == 0.1) {
        var v14 uint8 = 2
        _ = v14
    }
    if true {
        if false {
            v15 := v2[-9223372036854775808]
            var v16 int64 = 0
            _ = v15
            _ = v16
        } else {
            v17 := GenS3{F0: 65535, F1: 342, F2: 128}
            v18 := []int{2147483647, -9223372036854775808}
            _ = v17
            _ = v18
        }
        v10 = ((118 == 4294967295) && (-2.25 == 0.0))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
