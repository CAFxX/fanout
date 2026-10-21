package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int64 = -349
    var v2 int32 = 42
    var v3 float64 = 0.0
    if false {
        v4 := GenS1{F0: ("a" + "hello")}
        v3 = 1e10
        for fi5 := 0; fi5 < 15; fi5++ {
            v3 = v3
            var v6 uint32 = 416
            var v7 int16 = 712
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = v4
    } else {
        fmt.Println(v2, v1)
        fmt.Println(v2, v1)
    }
    v8 := GenS2{F0: 164}
    if true {
        var v9 float64 = ((-2.25 * float64(1.5)) + v3)
        _ = v9
    }
    switch 1 {
    case -189:
        fmt.Println(v3, v3)
    case 4294967295:
        var v10 bool = (-1 == -1)
        _ = v10
    case -889:
        v3 -= ((v3 - v3) * 0.0)
    }
    fmt.Println(v1, v1)
    v3 -= 0.0
    var v11 uint16 = 65535
    var v12 int = 32767
    if true {
        if false {
            v13 := []int{v12, v12, (v12 / ((v12) | 1)), v12}
            _ = v13
        } else {
            var v14 uint8 = 0
            v3 -= (v3 * (((float64(float32(v3)) - 1.5) * v3) / 1.0))
            _ = v14
        }
        var v15 int64 = v1
        switch "hello" {
        case "hello":
            switch "\n\t" {
            case "":
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 int8 = 100
                    v3 = (v3 * v3)
                    _ = fi16
                    _ = v17
                }
            case "\n\t":
                var v18 int = (v12 - (v12 ^ v12))
                _ = v18
            default:
                v11 -= (^(v11))
            }
        default:
            v19 := []int{687, -60, v12, v12}
            _ = v19
        }
        _ = v15
    }
    var v20 uint8 = 255
    var v21 string = (("foo bar" + "") + "")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v8
}
