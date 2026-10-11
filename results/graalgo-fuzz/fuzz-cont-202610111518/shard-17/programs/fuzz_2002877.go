package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool }
func genFn0() int8 {
    var v6 int8 = -1
    v7 := []int{-1, -838, -652, -2147483648}
    if (false && false) {
        v6 = -1
        v8 := GenS1{F0: 127, F1: ((("a" + ("" + "\n\t")) + ("" + "\n\t")) + "")}
        var v9 int8 = v6
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: 2147483647, F1: ""}
        _ = v10
    }
    var v11 uint16 = 65535
    fmt.Println(v6)
    _ = v11
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int16) uint {
    v12 := GenS1{F0: 2, F1: "a"}
    var v13 int = 9223372036854775807
    v13 = (v13 >> ((2) & 63))
    _ = p0
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 69
    var v3 float64 = (1e10 / 1.5)
    var v4 float64 = ((v3 / v3) - v3)
    var v5 uint8 = 2
    fmt.Println(v5)
    v14 := GenS2{F0: 2147483647, F1: "foo bar"}
    var v15 string = ""
    v16 := GenS1{F0: (^(v2)), F1: ((("a" + "") + ("a" + "a")) + v15)}
    var v17 uint8 = v2
    v18 := GenS2{F0: 65535, F1: "hello"}
    var v19 int8 = -114
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := []int{(fi20 + 543), 1}
        _ = fi20
        _ = v21
    }
    var v22 bool = true
    if (v22 && ((-(v19)) != -128)) {
        if true {
            v23 := GenS3{F0: (true || v22)}
            _ = v23
        } else {
            v24 := []int{32767, -943, 127}
            _ = v24
        }
    }
    v2 -= v5
    v2 = (v2 >> ((412) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
