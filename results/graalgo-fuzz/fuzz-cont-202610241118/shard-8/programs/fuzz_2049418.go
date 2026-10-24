package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []string{"a", "hello", "\n\t"}
    var v2 string = ("a" + "foo bar")
    var v3 int64 = -527
    var v4 uint32 = 2147483647
    var v5 int = 902
    var v6 uint = 0
    fmt.Println(v5, v5)
    v1 = append(v1, ((v2 + (v2 + v2)) + ((v2 + v2) + v2)))
    var v7 string = (v2 + "foo bar")
    if true {
        v8 := GenS1{F0: 434}
        _ = v8
    }
    v9 := GenS2{F0: 1}
    fmt.Println(v5, v7, v5)
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(fi10)
        if (v6 > v6) {
            v5 = -572
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 float32 = 1.5
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(fi10)
                _ = fi11
                _ = v12
            }
        }
        v13 := GenS1{F0: 65535}
        _ = fi10
        _ = v13
    }
    v14 := GenS3{F0: -32768, F1: 65535}
    v1 = append(v1, "foo bar")
    var v15 float64 = 1e10
    var v16 string = v2
    v17 := GenS3{F0: -1, F1: 887}
    var v18 uint16 = 42
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
