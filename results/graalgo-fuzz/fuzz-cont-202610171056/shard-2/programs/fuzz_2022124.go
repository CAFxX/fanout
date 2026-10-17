package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string }

func main() {
    var v1 uint16 = 255
    var v2 int = 32767
    v3 := make(map[string]uint32)
    var v4 int32 = -2147483648
    var v5 string = ("" + "a")
    var v6 uint16 = 65535
    var v7 uint64 = 282
    v5 = ((v5 + v5) + "\n\t")
    var v8 int8 = 100
    v9 := GenS1{F0: true, F1: false, F2: 713}
    v2 -= (^(v2))
    v10 := GenS1{F0: false, F1: false, F2: 42}
    var v11 uint16 = (^(v1))
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := v3[((v5 + "a") + v5)]
            var v15 bool = true
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS2{F0: 127, F1: 1}
            if true {
                if false {
                    var v18 int8 = (v8 / ((v8) | 1))
                    v7 = 18446744073709551615
                    var v19 uint = 42
                    _ = v18
                    _ = v19
                }
            } else {
                v5 = v5
            }
            _ = fi16
            _ = v17
        }
        switch (((v5 + (v5 + ("foo bar" + v5))) + "hello") + v5) {
        case "a":
            var v20 int64 = 842
            _ = v20
        case "\n\t":
            fmt.Println(v5)
        case "":
            var v21 float32 = 1.5
            _ = v21
        }
        _ = fi12
    }
    if (879 <= 1) {
        fmt.Println(v1, v7)
    } else {
        fmt.Println(v8, v5, v7)
        v22 := GenS3{F0: v5}
        var v23 bool = false
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
