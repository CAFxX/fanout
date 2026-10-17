package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 string; F2 uint16 }
func genFn0() int32 {
    if false {
        var v4 uint8 = 1
        _ = v4
    } else {
        var v5 bool = ((1.5 > 3.14159) || (true && false))
        if ((0.1 - 1e10) <= (1.5 + -2.25)) {
            v6 := []int{128, 9223372036854775807, -2147483648}
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := GenS1{F0: 18446744073709551615, F1: 127}
                var v9 int64 = -790
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := GenS2{F0: 1, F1: v9}
                    _ = fi10
                    _ = v11
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v6
        }
        _ = v5
    }
    switch (("foo bar" + "\n\t") + "\n\t") {
    case "":
        var v12 int16 = -1
        _ = v12
    case "a":
        var v13 bool = ((((0.0 + (-2.25 - float32((1.5 * float64((3.14159 + ((float32(1.5) + (3.14159 / 0.1)) + 0.1))))))) * 1.5) + -2.25) < 0.1)
        _ = v13
    }
    switch ("" + ("a" + "a")) {
    case "a":
        var v14 int8 = 0
        _ = v14
    default:
        v15 := GenS3{F0: -828}
        _ = v15
    }
    v16 := []int{4294967295, -9223372036854775808}
    var v17 int64 = 42
    _ = v16
    _ = v17
    return -2147483648
}

func main() {
    var v1 string = "hello"
    var v2 int8 = 0
    var v3 int32 = 2147483647
    switch 438 {
    case -912:
        v2 = v2
    case 2147483647:
        fmt.Println(v2, v1, v2)
    case 0:
        v2 -= (v2 * (^(v2)))
    default:
        fmt.Println(v1)
    }
    var v18 string = "a"
    if (v2 > v2) {
        v19 := GenS4{F0: v2, F1: v18, F2: 65535}
        v2 += v2
        v20 := []int{46, 0, 32767, -2147483648}
        _ = v19
        _ = v20
    }
    var v21 float64 = 1e10
    v21 = v21
    v18 = (v18 + ("\n\t" + v1))
    switch int(v2) {
    case 4294967295:
        fmt.Println(v3, v18, v21)
    }
    var v22 int8 = (v2 >> ((42) & 63))
    v22 -= (^(v2))
    if false {
        for fi23 := 0; fi23 < 6; fi23++ {
            if false {
                var v24 float32 = 0.0
                _ = v24
            }
            switch 2147483647 {
            case -32:
                v25 := GenS3{F0: fi23}
                _ = v25
            }
            if false {
                v26 := []int{}
                v1 = "a"
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 int32 = v3
                    var v29 int = fi27
                    v21 = 1e10
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            } else {
                fmt.Println(v18, v21, v21)
                fmt.Println(v18, v22)
            }
            _ = fi23
        }
    } else {
        v30 := []int{128}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
