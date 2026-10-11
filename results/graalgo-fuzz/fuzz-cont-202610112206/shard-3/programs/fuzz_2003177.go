package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{699, 2147483648}
    var v3 float64 = 0.1
    v4 := make(map[string]int)
    fmt.Println(v3, v3)
    var v5 bool = (!(255 >= 0))
    var v6 uint16 = 0
    v7 := v4["\n\t"]
    var v8 uint64 = 1
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint32 = 2147483648
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: -842, F1: 1, F2: 1}
            switch (^(fi11)) {
            case -1:
                switch (("hello" + "") + "a") {
                case "hello":
                    var v13 float32 = 1.5
                    _ = v13
                case "foo bar":
                    v1 = 3.14159
                }
            }
            switch "a" {
            case "":
                fmt.Println(v6)
            case "a":
                var v14 uint64 = 0
                _ = v14
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(fi11, v8, fi11)
            _ = fi11
            _ = v12
        }
        if false {
            var v15 int16 = 1
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    if false {
        v3 += 1e10
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v16 := v4["hello"]
        v17 := GenS1{F0: -1, F1: 35, F2: uint8(v8)}
        _ = v16
        _ = v17
    }
    fmt.Println(v7)
    v18 := GenS1{F0: 0, F1: 0, F2: 236}
    _ = len(v2)
    v3 += 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
