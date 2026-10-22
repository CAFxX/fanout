package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
func genFn0(p0 uint8) uint16 {
    p0 = 0
    p0 -= (^(p0))
    _ = p0
    return 65535
}

func main() {
    var v1 bool = (127 < 1)
    v2 := []uint32{}
    var v3 float32 = 0.0
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := GenS1{F0: "a", F1: (fi4 ^ fi4), F2: fi4}
        v6 := GenS1{F0: "\n\t", F1: fi4, F2: 127}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int64 = -411
    if v1 {
        v8 := GenS1{F0: ("a" + (((("hello" + ("foo bar" + "")) + (("foo bar" + ("" + ("a" + "\n\t"))) + "\n\t")) + "hello") + "")), F1: 661, F2: -2147483648}
        var v9 uint8 = 2
        var v10 int8 = 42
        _ = v10
        _ = v8
        _ = v9
    } else {
        v7 = (^(v7))
        v2 = append(v2, 2147483648)
        var v11 int16 = -472
        _ = v11
    }
    var v12 string = (("" + ("\n\t" + "foo bar")) + "hello")
    var v13 uint = 644
    v7 += v7
    var v14 string = (("" + v12) + v12)
    if (v1 && (!false)) {
        if false {
            _ = len(v2)
            var v15 int16 = 0
            _ = v15
        }
        v16 := GenS1{F0: "hello", F1: -43, F2: -735}
        _ = v16
    }
    var v17 uint8 = 158
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
