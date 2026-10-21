package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }

func main() {
    var v1 int = -1
    var v2 float64 = 0.0
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    var v5 int8 = 1
    var v6 int16 = -964
    v1 -= v1
    if true {
        var v7 int64 = 715
        _ = v7
    }
    v5 = 100
    _ = len(v4)
    if ((42 == 42) && false) {
        for fi8 := 0; fi8 < 16; fi8++ {
            delete(v3, fi8)
            _ = fi8
        }
        var v9 int64 = 9223372036854775807
        for fi10 := 0; fi10 < 7; fi10++ {
            v2 = ((0.0 - 1e10) / (float64((1e10 * -2.25)) / v2))
            _ = fi10
        }
        _ = v9
    } else {
        if ((2 >= 255) || ((v1 * (v1 & v1)) > 32767)) {
            fmt.Println(v5, v1)
            var v11 int = v1
            var v12 int32 = 436
            _ = v11
            _ = v12
        } else {
            var v13 float64 = v2
            for fi14 := 0; fi14 < 4; fi14++ {
                v6 += (v6 * v6)
                _ = fi14
            }
            _ = v13
        }
        v15 := GenS1{F0: v5, F1: 0, F2: (^(v5))}
        v16 := GenS1{F0: v5, F1: 42, F2: (v5 | 115)}
        _ = v15
        _ = v16
    }
    var v17 uint16 = 2
    if (!false) {
        v1 += (v1 << ((0) & 63))
        v18 := GenS1{F0: -112, F1: 42, F2: v5}
        _ = v18
    } else {
        v19 := []int{v1, v1, -2147483648, (v1 % (((v1 / ((-31) | 1))) | 1))}
        var v20 bool = (65535 < 59)
        v2 += v2
        _ = v19
        _ = v20
    }
    var v21 int64 = -9223372036854775808
    var v22 int16 = (v6 ^ (v6 * v6))
    var v23 uint64 = 18446744073709551615
    v24 := []int{v1, (v1 >> ((65535) & 63)), (v1 + 128)}
    var v25 bool = (3.14159 > 1e10)
    var v26 string = "foo bar"
    var v27 string = (v26 + ("\n\t" + "\n\t"))
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 int16 = v6
        var v30 uint16 = uint16(v23)
        _ = len(v24)
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 string = ((((v26 + ("" + v26)) + "hello") + v27) + ((("a" + (v27 + v26)) + "a") + ((v27 + v27) + v27)))
        _ = fi31
        _ = v32
    }
    switch (("a" + ((v27 + (("foo bar" + (("a" + "a") + (v26 + v26))) + "")) + (v26 + ""))) + ("\n\t" + (v26 + ((v27 + v26) + (v27 + v27))))) {
    case "":
        for fi33 := 0; fi33 < 7; fi33++ {
            delete(v3, v1)
            var v34 float64 = 0.0
            v6 += v22
            _ = fi33
            _ = v34
        }
    case "\n\t":
        fmt.Println(v25, v22, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
