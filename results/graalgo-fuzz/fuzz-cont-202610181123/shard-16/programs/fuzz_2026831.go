package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 459
    var v3 uint8 = 1
    v4 := make(map[string]int)
    var v5 float32 = 0.0
    var v6 int = -865
    if true {
        _ = len(v1)
    }
    fmt.Println(v6, v2, v3)
    var v7 float64 = 3.14159
    fmt.Println(v7)
    var v8 uint8 = v3
    switch ("" + "") {
    case "\n\t":
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS1{F0: ("foo bar" + "foo bar"), F1: (v2 * v2), F2: -320}
            for fi11 := 0; fi11 < 15; fi11++ {
                delete(v1, (^(fi11)))
                v6 = -1
                var v12 int = fi9
                _ = fi11
                _ = v12
            }
            v13 := []int{(v6 << ((v2) & 63)), fi9}
            _ = fi9
            _ = v10
            _ = v13
        }
    default:
        if true {
            var v14 int64 = 9223372036854775807
            fmt.Println(v8, v7, v2)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                if (147 != 2147483647) {
                    v8 -= 42
                    v2 = (v2 >> (((^(v2))) & 63))
                }
                _ = fi15
            }
        }
    }
    v16 := []int{}
    fmt.Println(v3, v2, v7)
    if (!(v5 != (0.1 - v5))) {
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint8 = v8
            _ = fi17
            _ = v18
        }
    } else {
        v19 := GenS2{F0: 1, F1: v2, F2: 145}
        _ = v19
    }
    fmt.Println(v3, v8)
    v5 = (3.14159 - v5)
    switch (("hello" + ("hello" + ("foo bar" + "a"))) + "\n\t") {
    case "":
        v5 = v5
    }
    var v20 uint = 364
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
