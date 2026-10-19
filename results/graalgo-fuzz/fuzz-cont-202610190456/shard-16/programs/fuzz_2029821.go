package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }

func main() {
    var v1 int32 = -981
    var v2 uint = 0
    var v3 uint64 = 18446744073709551615
    var v4 float64 = 0.1
    var v5 int8 = 0
    fmt.Println(v3, v2, v1)
    switch ("\n\t" + ("a" + "")) {
    case "a":
        v3 = 4294967295
    case "":
        var v6 uint32 = 42
        _ = v6
    case "hello":
        v4 += 1.5
    default:
        fmt.Println(v3)
    }
    var v7 uint = (v2 ^ v2)
    v5 -= (^(v5))
    var v8 int64 = -1
    fmt.Println(v5)
    v7 += (^(v2))
    if (!(false || true)) {
        var v9 int16 = 42
        var v10 int16 = (^(v9))
        _ = v10
        _ = v9
    } else {
        if true {
            v11 := GenS1{F0: 255, F1: v1}
            var v12 int64 = (^(v8))
            var v13 int16 = 1000
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 0, F1: 922}
            if (42 > 616) {
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := GenS2{F0: 65535, F1: 65535, F2: 2}
                    switch (((("foo bar" + "a") + "\n\t") + "") + "a") {
                    case "a":
                        v2 += v7
                    case "hello":
                        var v17 uint16 = 42
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
            } else {
                var v18 bool = false
                _ = v18
            }
            _ = v14
        }
    }
    switch "hello" {
    case "a":
        v19 := []int{}
        _ = v19
    case "foo bar":
        v7 += v7
    }
    if (true && true) {
        var v20 uint = 42
        var v21 uint16 = 0
        _ = v20
        _ = v21
    } else {
        v7 = v7
        var v22 string = ""
        v23 := []int{127}
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 255, F1: (v1 ^ -782)}
    v4 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
