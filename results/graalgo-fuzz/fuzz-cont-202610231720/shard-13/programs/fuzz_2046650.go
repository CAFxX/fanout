package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 32767
    var v3 int = 4294967295
    var v4 int32 = 898
    for fi5 := 0; fi5 < 10; fi5++ {
        switch 127 {
        case -645:
            for fi6 := 0; fi6 < 11; fi6++ {
                fmt.Println(fi6)
                var v7 float64 = float64(0.1)
                if true {
                    v8 := []int{}
                    v4 = -850
                    v2 = v2
                    _ = v8
                }
                _ = fi6
                _ = v7
            }
        case 715:
            v1 = (^(v1))
        case 128:
            var v9 bool = (-2.25 > (float64((0.0 * float32(0.1))) * 1e10))
            _ = v9
        default:
            v3 = int(v1)
        }
        _ = fi5
    }
    fmt.Println(v3)
    v10 := GenS1{F0: v2, F1: ("hello" + "")}
    v4 = (-(v4))
    v11 := GenS1{F0: v2, F1: "\n\t"}
    if false {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v2, v4)
        case "":
            var v12 uint32 = 1
            _ = v12
        case "hello":
            v13 := []int{v3, v3}
            _ = v13
        }
    } else {
        v14 := GenS1{F0: 2, F1: ("hello" + ("a" + "a"))}
        _ = v14
    }
    var v15 uint8 = 0
    var v16 uint16 = v2
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            fmt.Println(v4)
            _ = fi18
        }
        var v19 int32 = v4
        v20 := []int{(fi17 & (v3 / (((v3 + int(v19))) | 1))), fi17, (fi17 & 88)}
        _ = fi17
        _ = v19
        _ = v20
    }
    switch "foo bar" {
    case "foo bar":
        v21 := GenS1{F0: v16, F1: "a"}
        _ = v21
    case "":
        v22 := []int{(^(v3)), (v3 % ((v3) | 1))}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
