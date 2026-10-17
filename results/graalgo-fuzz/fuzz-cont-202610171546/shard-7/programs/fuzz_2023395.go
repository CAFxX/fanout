package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
func genFn0() uint8 {
    switch 127 {
    case 32768:
        var v6 int16 = -1
        _ = v6
    case -2147483648:
        switch ("foo bar" + "hello") {
        case "hello":
            if true {
                var v7 float32 = 0.1
                var v8 uint64 = 42
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := []int{fi9, 9223372036854775807}
                    _ = len(v10)
                    _ = fi9
                    _ = v10
                }
                v11 := []int{-612, 0, -2147483648, 9223372036854775807}
                _ = v11
            }
        }
    case 0:
        var v12 int = 0
        _ = v12
    default:
        if false {
            v13 := []int{32767, 9223372036854775807, 204}
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(fi14)
                _ = fi14
            }
            _ = v13
        }
    }
    v15 := []int{128, 1, 128}
    var v16 int8 = 78
    v16 += 1
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 int16 = 1000
    var v2 uint64 = 0
    var v3 uint = 0
    var v4 float32 = (-2.25 - 0.1)
    var v5 float64 = -2.25
    if (0 >= 32) {
        v17 := []int{32767, 127, 42}
        if len(v17) > 0 {
            v17[0] = 42
        }
        var v18 bool = ((9223372036854775807 == 133) && false)
        _ = v17
        _ = v18
    }
    var v19 int32 = -10
    var v20 int8 = -50
    v21 := []int{-1, 1, 32767}
    var v22 float64 = 0.0
    v5 -= (v22 + (0.0 - (v22 - (v22 / (v5 * float64((float32(v5) + 1.5)))))))
    var v23 int16 = v1
    v24 := GenS1{F0: ("a" + ((("hello" + "foo bar") + "a") + ("foo bar" + (((("\n\t" + "hello") + "\n\t") + "") + "foo bar"))))}
    fmt.Println(v19)
    var v25 uint = v3
    v21 = append(v21, 32767)
    v26 := GenS2{F0: 918}
    fmt.Println(v4, v19)
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            v23 = (v23 + (v23 & (v1 ^ v1)))
            _ = fi28
        }
        _ = fi27
    }
    fmt.Println(v25, v25, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
