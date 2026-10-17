package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int16 = -32768
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{1, 4294967295, 2147483647, 2147483648}
    var v4 string = "hello"
    v5 := GenS1{F0: 0}
    var v6 int32 = 0
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int32 = 42
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v1)
    fmt.Println(v1, v2, v2)
    if (!false) {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch "" {
            case "\n\t":
                v4 = "foo bar"
            case "":
                var v10 float32 = 0.0
                _ = v10
            }
            switch ("\n\t" + (v4 + "hello")) {
            case "hello":
                var v11 bool = (2147483648 == 1)
                _ = v11
            }
            var v12 string = "hello"
            _ = fi9
            _ = v12
        }
        if true {
            v13 := GenS1{F0: 42}
            var v14 int8 = 42
            v1 = v1
            _ = v13
            _ = v14
        } else {
            var v15 bool = true
            _ = v15
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v16 := GenS1{F0: 324}
        _ = v16
    }
    fmt.Println(v2, v6, v2)
    var v17 uint64 = v2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
