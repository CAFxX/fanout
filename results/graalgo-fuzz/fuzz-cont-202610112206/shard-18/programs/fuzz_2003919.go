package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
func genFn0() int32 {
    if false {
        if (true && false) {
            var v7 int64 = 9223372036854775807
            _ = v7
        } else {
            var v8 uint32 = 0
            _ = v8
        }
        var v9 int = -628
        _ = v9
    } else {
        var v10 uint8 = 76
        fmt.Println(v10)
        v11 := []int{}
        _ = v10
        _ = v11
    }
    var v12 uint16 = 660
    var v13 string = "a"
    _ = v12
    _ = v13
    return -2147483648
}

func main() {
    var v1 string = "a"
    var v2 int64 = -533
    var v3 uint = 0
    var v4 float64 = 1e10
    var v5 bool = (606 <= (v2 / ((v2) | 1)))
    var v6 uint16 = 2
    v1 = v1
    fmt.Println(v1, v6, v6)
    var v14 int64 = v2
    v15 := GenS1{F0: (v5 && (1 < 128)), F1: (v3 >> ((732) & 63))}
    v16 := []int{127, 32768}
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v4, fi17, v14)
        fmt.Println(v5, v1)
        _ = fi17
    }
    v18 := GenS1{F0: v5, F1: v3}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    switch ((v1 + v1) + (v1 + v1)) {
    case "\n\t":
        v2 += 42
    default:
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    var v19 uint8 = 255
    var v20 int = int(v19)
    var v21 bool = v5
    if (!v5) {
        v22 := GenS1{F0: v5, F1: v3}
        _ = v22
    }
    fmt.Println(v4)
    _ = len(v16)
    var v23 int64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
