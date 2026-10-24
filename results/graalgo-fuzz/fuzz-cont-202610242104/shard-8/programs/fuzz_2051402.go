package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint; F2 int64 }

func main() {
    var v1 bool = true
    var v2 int = 127
    var v3 uint32 = 2147483647
    var v4 int = -440
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{(fi5 * -1), fi5, (v2 << ((225) & 63))}
        _ = fi5
        _ = v6
    }
    switch (("\n\t" + ("a" + "a")) + "") {
    case "":
        v7 := GenS1{F0: -2147483648, F1: "hello"}
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int8 = 127
            _ = fi8
            _ = v9
        }
    default:
        v2 += 9223372036854775807
    }
    var v10 int = -9223372036854775808
    var v11 uint64 = 18446744073709551615
    if (((float64(0.0) * (0.1 + 0.1)) <= 1.5) || (v1 || v1)) {
        v12 := []int{v10, (v4 / ((v2) | 1)), v10, (v4 / ((v4) | 1))}
        _ = v12
    }
    if ((((v1 && v1) && true) && (0 != 42)) || (!v1)) {
        for fi13 := 0; fi13 < 17; fi13++ {
            v10 = v2
            fmt.Println(v4)
            _ = fi13
        }
        if (v1 && false) {
            v14 := []int{v4, v4, (v2 % (((-(v4))) | 1))}
            v15 := GenS1{F0: 0, F1: "a"}
            var v16 uint32 = v3
            _ = v14
            _ = v15
            _ = v16
        }
    } else {
        v17 := []int{(v4 & 32767), (v10 ^ (v10 & (v2 & v2))), (-(v10)), (v4 - 42)}
        v4 += v4
        var v18 string = "\n\t"
        _ = v17
        _ = v18
    }
    var v19 uint8 = 255
    v20 := GenS2{F0: (^(v3)), F1: (^(v3))}
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 bool = false
        var v23 string = "hello"
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 890, F1: "\n\t"}
    v11 += v11
    var v25 int = (v4 ^ v2)
    var v26 int8 = 127
    switch 127 {
    case 2147483647:
        for fi27 := 0; fi27 < 3; fi27++ {
            v28 := GenS3{F0: 1, F1: 65535, F2: 42}
            v29 := []int{}
            _ = fi27
            _ = v28
            _ = v29
        }
    case 32767:
        v25 -= v25
    case 127:
        var v30 string = "\n\t"
        _ = v30
    }
    v10 = (v4 * 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
