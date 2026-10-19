package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
func genFn0() int {
    v5 := GenS1{F0: -128}
    switch 32767 {
    case -636:
        v6 := GenS1{F0: 0}
        _ = v6
    case -9223372036854775808:
        var v7 int = 1
        _ = v7
    default:
        v8 := []int{1}
        _ = v8
    }
    var v9 uint16 = 1
    v9 += (v9 * 476)
    _ = v5
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int = 127
    var v2 string = "foo bar"
    var v3 string = ((v2 + (v2 + ("a" + ((v2 + v2) + "a")))) + "foo bar")
    var v4 float32 = 0.1
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v2, fi10, v3)
        v11 := []int{(v1 & (^(fi10))), fi10}
        v2 = ("foo bar" + (v3 + (((v2 + "foo bar") + v2) + "\n\t")))
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 65535
    var v13 uint64 = 325
    var v14 int64 = 1
    switch -9223372036854775808 {
    case 128:
        v1 += v1
    case 32767:
        v15 := GenS1{F0: 0}
        _ = v15
    case 32768:
        v2 = ((("a" + v3) + v2) + v2)
    }
    v12 += (^(v12))
    v3 = v2
    v3 = v2
    switch (v1 % (((^(v1))) | 1)) {
    case -952:
        v14 -= v14
    case -2147483648:
        v16 := GenS2{F0: v1}
        _ = v16
    }
    var v17 uint64 = (v13 + (v13 - v13))
    fmt.Println(v3, v3)
    var v18 int32 = 254
    if true {
        var v19 string = v3
        v20 := GenS2{F0: (v1 & v1)}
        v17 = v13
        _ = v19
        _ = v20
    } else {
        v21 := []int{}
        fmt.Println(v2, v4, v17)
        var v22 int16 = 1
        _ = v21
        _ = v22
    }
    v23 := []int{v1, -1}
    switch -1 {
    case -9223372036854775808:
        v18 = (v18 | v18)
    case 4294967295:
        if (true && false) {
            var v24 uint8 = 0
            _ = v24
        }
    default:
        v25 := GenS1{F0: 1}
        _ = v25
    }
    v26 := GenS2{F0: genFn0()}
    if true {
        var v27 int16 = -622
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
