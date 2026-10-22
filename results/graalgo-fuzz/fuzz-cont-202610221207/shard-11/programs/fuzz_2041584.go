package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 uint64) uint16 {
    v6 := GenS1{F0: (16 == 127), F1: 1}
    p1 -= (p1 % (((^(p1))) | 1))
    var v7 uint16 = 32767
    var v8 int64 = 112
    var v9 bool = false
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v7 + v7)
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 572
    var v3 int16 = -381
    var v4 string = (("hello" + ("hello" + "a")) + "\n\t")
    var v5 int16 = 1
    v4 = ((v4 + "\n\t") + ((v4 + v4) + v4))
    v5 = v5
    v10 := GenS1{F0: (((0.0 / 1.0) - 1e10) > ((0.1 * 3.14159) * 0.0)), F1: (v5 + v5)}
    if true {
        v5 = v3
        switch 2147483647 {
        case 32768:
            v11 := GenS2{F0: false, F1: v4}
            _ = v11
        case 1:
            v3 += v5
        }
        var v12 uint8 = 2
        _ = v12
    } else {
        v13 := GenS1{F0: (0 != -1), F1: (v3 ^ (v5 / ((32767) | 1)))}
        var v14 uint = 255
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: true, F1: "hello"}
    v16 := GenS2{F0: (false && true), F1: ((("foo bar" + v4) + (v4 + v4)) + "foo bar")}
    fmt.Println(v5, v2)
    var v17 int8 = -28
    v4 = "a"
    if false {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS2{F0: (0 < v5), F1: v4}
            _ = fi18
            _ = v19
        }
        var v20 uint64 = 4294967295
        _ = v20
    } else {
        v21 := GenS2{F0: (!false), F1: v4}
        fmt.Println(v5)
        var v22 int64 = v2
        _ = v21
        _ = v22
    }
    v5 += 784
    v23 := v1[-99]
    v24 := GenS3{F0: 2147483647}
    var v25 uint32 = 831
    delete(v1, 470)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
