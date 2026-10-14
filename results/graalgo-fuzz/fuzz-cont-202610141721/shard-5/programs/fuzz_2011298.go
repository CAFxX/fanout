package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }

func main() {
    var v1 uint16 = 2
    var v2 uint16 = (^(v1))
    var v3 int32 = 374
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    _ = len(v5)
    if false {
        if true {
            var v6 float32 = 0.0
            var v7 string = ("foo bar" + "foo bar")
            v8 := []int{2147483647}
            _ = v6
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float32 = 1e10
        _ = fi9
        _ = v10
    }
    var v11 float32 = 1e10
    v12 := GenS1{F0: (v2 | 65535)}
    for fi13 := 0; fi13 < 5; fi13++ {
        switch -2147483648 {
        case -1:
            var v14 uint16 = 1
            _ = v14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        if true {
            for fi16 := 0; fi16 < 14; fi16++ {
                v17 := []int{(fi15 >> ((740) & 63)), fi16}
                var v18 int = -857
                var v19 int16 = 589
                if (!(!(-1 >= 0))) {
                    fmt.Println(v3, v19)
                    fmt.Println(fi15)
                } else {
                    if (false || true) {
                        fmt.Println(v18, v11)
                    } else {
                        var v20 int8 = 100
                        _ = v20
                    }
                    v21 := GenS2{F0: -128, F1: 84, F2: (v2 + (^(v1)))}
                    _ = v21
                }
                fmt.Println(fi16, v18)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            var v22 uint16 = v2
            var v23 uint32 = 767
            _ = v22
            _ = v23
        } else {
            delete(v4, 121)
            fmt.Println(fi15, fi15)
            var v24 bool = ("foo bar" != ("hello" + "\n\t"))
            _ = v24
        }
        if false {
            fmt.Println(v2, fi15, v11)
            v2 = v2
        }
        _ = fi15
    }
    v2 = v1
    for fi25 := 0; fi25 < 11; fi25++ {
        v2 = 2
        _ = fi25
    }
    v26 := v4[807]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
