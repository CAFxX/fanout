package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 uint64; F2 string }
func genFn0() int64 {
    v5 := []int{106, 3}
    var v6 uint = 4294967295
    var v7 float32 = 0.0
    v6 -= v6
    _ = len(v5)
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint = 42
    var v2 int16 = 718
    var v3 int16 = v2
    v4 := make(map[uint32]uint32)
    var v8 int32 = -997
    v9 := []int{778, -2147483648, -60, -239}
    v10 := GenS1{F0: (v1 - (v1 % (((^(v1))) | 1))), F1: -2147483648}
    fmt.Println(v2, v3)
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS2{F0: 127, F1: 4294967295, F2: "foo bar"}
        _ = fi11
        _ = v12
    }
    delete(v4, 913)
    var v13 string = ("a" + "a")
    var v14 int8 = 1
    fmt.Println(v13)
    if false {
        var v15 bool = false
        _ = v15
    }
    switch -2147483648 {
    case 375:
        v13 = "hello"
    default:
        if (v8 < v8) {
            v13 = "a"
            if true {
                if ((32768 >= -9223372036854775808) && true) {
                    var v16 int = -302
                    _ = v16
                } else {
                    if (true || (!true)) {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            var v18 string = ("foo bar" + v13)
                            v9 = append(v9, fi17)
                            _ = fi17
                            _ = v18
                        }
                    }
                    v1 += v1
                    var v19 uint16 = 32767
                    _ = v19
                }
            }
            if (1 <= 4294967295) {
                fmt.Println(v14)
                fmt.Println(v14, v1, v3)
            } else {
                if len(v9) > 0 {
                    v9[0] = 32767
                }
                var v20 int8 = v14
                _ = len(v4)
                _ = v20
            }
        }
    }
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
