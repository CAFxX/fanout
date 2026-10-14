package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float32 = float32((float64(3.14159) * 1e10))
    var v2 string = ("hello" + "")
    var v3 int = 268
    var v4 float32 = v1
    var v5 string = ((v2 + v2) + v2)
    var v6 uint = 1
    if (!false) {
        var v7 int32 = 42
        _ = v7
    }
    switch (v3 / ((-342) | 1)) {
    case 2147483647:
        fmt.Println(v5, v3)
    case 9223372036854775807:
        fmt.Println(v1)
    default:
        if false {
            v8 := []int{9223372036854775807, (^(v3))}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v1)
        if true {
            v1 -= 3.14159
            v4 += -2.25
            for fi10 := 0; fi10 < 13; fi10++ {
                if false {
                    v11 := GenS1{F0: fi9}
                    _ = v11
                }
                v6 += v6
                var v12 uint = v6
                _ = fi10
                _ = v12
            }
        }
        _ = fi9
    }
    if (!true) {
        for fi13 := 0; fi13 < 8; fi13++ {
            v4 = (v4 / float32(float64((v4 * -2.25))))
            var v14 int32 = 452
            var v15 int32 = v14
            _ = fi13
            _ = v14
            _ = v15
        }
        v4 = 3.14159
    } else {
        v16 := []int{v3, (v3 | 127), (^(v3))}
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 float32 = (v1 + (v4 * (3.14159 - v1)))
        switch (("foo bar" + "foo bar") + v5) {
        case "a":
            fmt.Println(v6)
        case "hello":
            fmt.Println(v1, v6)
        case "":
            v19 := []int{fi17}
            _ = v19
        }
        if (!(!false)) {
            v18 += 3.14159
            switch 42 {
            case -862:
                var v20 uint8 = 127
                _ = v20
            }
            fmt.Println(v5, v1)
        } else {
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 uint64 = 42
                v23 := []int{fi21}
                v18 = v4
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        _ = fi17
        _ = v18
    }
    var v24 int64 = 9223372036854775807
    var v25 int = v3
    var v26 int = 0
    v26 = (v26 / ((v26) | 1))
    v4 -= (1.5 * ((1e10 * 1e10) - float32((0.1 * -2.25))))
    for fi27 := 0; fi27 < 11; fi27++ {
        if (!false) {
            fmt.Println(fi27, fi27)
        }
        v26 += v26
        _ = fi27
    }
    v6 += (v6 >> (((v6 | (v6 ^ v6))) & 63))
    var v28 uint = v6
    var v29 int32 = -152
    fmt.Println(v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
