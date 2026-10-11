package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int; F1 int8; F2 int64 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0() uint64 {
    var v6 float32 = (3.14159 * 1.5)
    v7 := GenS1{F0: 2, F1: 19}
    v8 := GenS2{F0: 32767, F1: -114, F2: -935}
    var v9 bool = (!true)
    fmt.Println(v9, v9)
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() int32 {
    var v10 int32 = 2147483647
    var v11 uint32 = 2147483648
    if (2 <= 2) {
        v12 := []int{32768, 32767}
        _ = v12
    }
    var v13 uint16 = 32767
    v10 = -120
    _ = v10
    _ = v11
    _ = v13
    return v10
}

func main() {
    var v1 string = ""
    var v2 int64 = -1
    var v3 string = (v1 + v1)
    var v4 uint8 = 128
    var v5 int32 = 42
    var v14 int32 = (v5 ^ 191)
    var v15 int16 = -331
    v3 = v1
    if false {
        v16 := GenS1{F0: 18446744073709551615, F1: 127}
        var v17 int64 = (v2 % (((^(v2))) | 1))
        _ = v16
        _ = v17
    }
    fmt.Println(v5, v4, v14)
    v1 = v1
    var v18 uint16 = 1
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS3{F0: 42, F1: 32767}
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
