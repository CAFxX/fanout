package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }

func main() {
    var v1 bool = false
    var v2 string = ("a" + ("a" + ("foo bar" + "")))
    v3 := make(map[int]string)
    fmt.Println(v2, v2)
    var v4 int16 = -807
    var v5 string = v2
    v1 = v1
    var v6 bool = v1
    var v7 uint64 = 1
    var v8 string = ""
    var v9 uint16 = 42
    var v10 uint8 = 2
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 127, F1: (^(v7)), F2: v4}
        v13 := GenS1{F0: 42, F1: (v7 * v7), F2: (v4 | (v4 + v4))}
        if (v6 || true) {
            for fi14 := 0; fi14 < 8; fi14++ {
                v7 += v7
                v2 = "foo bar"
                var v15 float32 = (3.14159 - 3.14159)
                _ = fi14
                _ = v15
            }
            if true {
                if v6 {
                    if (1 != 127) {
                        v16 := GenS1{F0: -1, F1: v7, F2: 784}
                        _ = v16
                    } else {
                        fmt.Println(v5, fi11)
                    }
                } else {
                    v8 = ("foo bar" + v5)
                }
                fmt.Println(v10)
            }
        } else {
            fmt.Println(v7, v7)
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v17 int64 = -1
    for fi18 := 0; fi18 < 7; fi18++ {
        if v1 {
            fmt.Println(v6, v2)
            v19 := GenS1{F0: 1, F1: v7, F2: 1}
            _ = v19
        } else {
            for fi20 := 0; fi20 < 10; fi20++ {
                fmt.Println(v2, fi18, v2)
                fmt.Println(v1, v5)
                fmt.Println(fi18, v9, v6)
                _ = fi20
            }
            v1 = (true || v1)
        }
        var v21 uint64 = v7
        _ = fi18
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 int32 = 2147483647
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 25, F1: v7, F2: 552}
    v25 := GenS1{F0: int8(v10), F1: v7, F2: (v4 >> ((1) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
