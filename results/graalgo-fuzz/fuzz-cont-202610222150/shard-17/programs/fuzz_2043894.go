package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int; F2 bool }
func genFn0() int {
    v4 := GenS1{F0: 0, F1: -1}
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS2{F0: 1, F1: 18446744073709551615}
        _ = fi5
        _ = v6
    }
    var v7 float32 = (1.5 * 0.0)
    v7 -= float32(0.1)
    _ = v4
    _ = v7
    return -439
}

func main() {
    var v1 float64 = 3.14159
    var v2 bool = (false && false)
    var v3 float32 = 1.5
    fmt.Println(v3)
    var v8 int32 = -2147483648
    var v9 int16 = 32767
    v10 := GenS3{F0: 255, F1: 127, F2: (!(v2 || v2))}
    fmt.Println(v9)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int8 = -1
        if (v8 <= -37) {
            var v13 uint32 = 65535
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v9 += (v9 / (((v9 + (v9 ^ int16(v8)))) | 1))
    var v14 int32 = v8
    v15 := []int{2147483647, 20}
    var v16 string = ("foo bar" + (("" + "") + ("hello" + "a")))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
