package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0() int {
    if (-2147483648 > 32768) {
        var v4 int64 = -9223372036854775808
        _ = v4
    }
    var v5 uint = 1
    v6 := GenS1{F0: -1, F1: 65535}
    if (!false) {
        v7 := []int{-74, 32767, 419}
        v5 = 1
        _ = v7
    } else {
        v8 := GenS1{F0: -1, F1: v5}
        fmt.Println(v5)
        _ = v8
    }
    _ = v5
    _ = v6
    return -920
}
func genFn1() int64 {
    var v9 bool = (((3.14159 / float64(-2.25)) / -2.25) < 1e10)
    for fi10 := 0; fi10 < 12; fi10++ {
        if true {
            var v11 uint8 = 42
            var v12 int32 = -966
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint16 = 255
                v11 = (^(v11))
                v9 = (v9 && true)
                _ = fi13
                _ = v14
            }
            _ = v11
            _ = v12
        }
        v9 = (!("a" == "hello"))
        v9 = v9
        _ = fi10
    }
    v9 = ((!true) || (!(true || true)))
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint = 65535
    var v2 int16 = -888
    var v3 string = "foo bar"
    fmt.Println(v1, v1, v1)
    v15 := []int{genFn0(), 127, 1}
    var v16 uint = 0
    var v17 uint16 = 1
    v15 = append(v15, -874)
    for fi18 := 0; fi18 < 7; fi18++ {
        _ = len(v15)
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v1, v1, v16)
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v17 -= v17
            _ = fi19
        }
        _ = fi18
    }
    var v20 uint = (v1 ^ v1)
    v2 -= v2
    var v21 uint32 = 72
    var v22 int16 = 152
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
