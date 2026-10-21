package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint) string {
    var v5 uint16 = 665
    var v6 float32 = float32(float64(0.0))
    if (float32(1.5) >= ((((1e10 / float32(((((1e10 - 0.1) / -2.25) * (3.14159 + 0.0)) * ((3.14159 - (1e10 + (0.0 * float64(-2.25)))) / 0.1)))) / (float32(((0.1 * 1.5) + (float64(v6) + -2.25))) * float32(1e10))) - float32(0.1)) + (v6 - (v6 * -2.25)))) {
        var v7 int8 = 1
        var v8 uint32 = 2147483648
        v9 := GenS1{F0: -1}
        _ = v7
        _ = v8
        _ = v9
    } else {
        v5 -= 1
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: 100}
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: 0}
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v5
    _ = v6
    return "a"
}
func genFn1() uint8 {
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(fi14)
        v15 := []int{fi14, (fi14 << ((0) & 63))}
        var v16 uint8 = 53
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int32 = 828
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint64 = 42
        switch "foo bar" {
        case "a":
            fmt.Println(v17)
        }
        var v20 int8 = 100
        _ = fi18
        _ = v19
        _ = v20
    }
    switch "" {
    case "a":
        fmt.Println(v17)
    }
    _ = v17
    return 118
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{32768, -9223372036854775808, 960}
    var v3 uint64 = 173
    var v4 uint16 = 479
    if false {
        v1 = ((v1 + v1) + v1)
        v1 = "a"
    }
    v4 -= 1
    var v21 float32 = 0.1
    var v22 int = 9223372036854775807
    var v23 int8 = 1
    switch v22 {
    case -171:
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS2{F0: int(v3)}
            var v26 uint64 = 313
            v27 := GenS3{F0: 0}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    case -1:
        fmt.Println(v23, v4)
    case 889:
        var v28 int = v22
        _ = v28
    default:
        fmt.Println(v22)
    }
    var v29 uint32 = 2147483647
    var v30 bool = true
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            v33 := GenS1{F0: v23}
            var v34 int32 = 42
            v34 += -208
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
