package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int8 }
func genFn0() int8 {
    var v5 uint64 = 551
    v6 := GenS1{F0: 42, F1: 2, F2: 42}
    v5 += 2
    _ = v5
    _ = v6
    return 127
}
func genFn1() int {
    switch -9223372036854775808 {
    case 42:
        v7 := GenS1{F0: 100, F1: 1, F2: 255}
        _ = v7
    }
    var v8 int8 = 0
    var v9 uint32 = 724
    for fi10 := 0; fi10 < 12; fi10++ {
        v9 -= v9
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(fi11)
            switch "\n\t" {
            case "hello":
                v8 = (v8 - v8)
            }
            var v12 uint32 = v9
            _ = fi11
            _ = v12
        }
        switch (fi10 + fi10) {
        case 32768:
            v13 := []int{-377, (fi10 & fi10), fi10, (fi10 * (fi10 << ((255) & 63)))}
            _ = v13
        }
        _ = fi10
    }
    _ = v8
    _ = v9
    return 361
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]uint32)
    var v3 int32 = 1
    var v4 bool = true
    fmt.Println(v3, v3)
    switch 238 {
    case 119:
        v4 = v4
    }
    if ("foo bar" <= "") {
        v4 = v4
    }
    switch 1 {
    case -966:
        v4 = (42 == 2)
    }
    var v14 int32 = -521
    v15 := []int{9223372036854775807, -9223372036854775808}
    if (v4 || (v4 || (1000 <= 0))) {
        v16 := GenS1{F0: 127, F1: 4294967295, F2: 1}
        var v17 int = -2147483648
        _ = v16
        _ = v17
    }
    delete(v2, 0)
    if false {
        v15 = append(v15, -2147483648)
        if true {
            fmt.Println(v1)
            var v18 bool = (264 <= 645)
            _ = v18
        } else {
            var v19 float64 = (1.5 + 0.1)
            _ = v19
        }
    }
    v20 := GenS2{F0: -1}
    v3 = -503
    var v21 float32 = -2.25
    var v22 uint = 255
    var v23 string = "\n\t"
    var v24 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
