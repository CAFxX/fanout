package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16; F1 uint8; F2 string }
func genFn0() bool {
    var v6 float32 = 3.14159
    v6 -= v6
    var v7 int32 = 1
    _ = v6
    _ = v7
    return false
}
func genFn1() int64 {
    if (0 > -9223372036854775808) {
        v8 := []int{469, -1, 4294967295}
        v9 := GenS1{F0: 1, F1: -2147483648}
        var v10 int64 = 117
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 648, F1: 0}
    var v12 uint = 42
    fmt.Println(v12)
    v13 := []int{}
    _ = v11
    _ = v12
    _ = v13
    return 9223372036854775807
}

func main() {
    var v1 bool = (false || (2 <= 65535))
    var v2 int32 = 42
    var v3 uint16 = 0
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    if (2 != 32767) {
        var v14 uint32 = 42
        _ = v14
    } else {
        v2 -= v2
    }
    v15 := []int{}
    fmt.Println(v3, v3)
    v5[128] = "foo bar"
    var v16 int16 = -127
    _ = len(v15)
    v17 := GenS2{F0: 32767, F1: 23, F2: ("hello" + "\n\t")}
    var v18 uint16 = v3
    var v19 int = 456
    v20 := GenS1{F0: (v18 | v3), F1: 1}
    for fi21 := 0; fi21 < 7; fi21++ {
        v3 = v3
        v5[v19] = "a"
        _ = fi21
    }
    v22 := GenS1{F0: (^(v3)), F1: (v2 % (((v2 | 2147483647)) | 1))}
    fmt.Println(v16, v1)
    v23 := GenS2{F0: 1, F1: 227, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
