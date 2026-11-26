package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    var v4 int8 = 42
    var v5 bool = (0 >= 255)
    var v6 uint16 = 2
    v4 += v4
    _ = v4
    _ = v5
    _ = v6
    return "a"
}

func main() {
    var v1 int8 = 1
    var v2 uint16 = 32767
    var v3 uint16 = (v2 & 1)
    var v7 string = "hello"
    if true {
        v8 := []int{32767}
        v8 = append(v8, 4294967295)
        _ = v8
    } else {
        v3 += (v3 * 65535)
        fmt.Println(v1, v1, v3)
    }
    var v9 int16 = 0
    v10 := GenS1{F0: 42}
    for fi11 := 0; fi11 < 15; fi11++ {
        switch "\n\t" {
        case "a":
            v1 = -112
        default:
            v9 -= int16(v1)
        }
        _ = fi11
    }
    var v12 uint = uint(v3)
    var v13 int = 42
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := []int{fi14}
        _ = fi14
        _ = v15
    }
    var v16 int8 = (v1 + -1)
    for fi17 := 0; fi17 < 14; fi17++ {
        switch "" {
        case "\n\t":
            v18 := []int{v13, (-(fi17)), -55}
            _ = v18
        case "":
            v19 := GenS1{F0: 2147483647}
            _ = v19
        case "hello":
            var v20 uint8 = 127
            _ = v20
        }
        _ = fi17
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS1{F0: -749}
        _ = fi21
        _ = v22
    }
    var v23 uint32 = 2147483648
    var v24 uint64 = 0
    fmt.Println(v9, v16, v24)
    v23 = v23
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v9
}
