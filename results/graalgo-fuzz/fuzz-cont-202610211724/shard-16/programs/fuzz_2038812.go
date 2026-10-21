package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int }

func main() {
    var v1 string = ("hello" + "")
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    switch 4294967295 {
    case 9223372036854775807:
        switch -9223372036854775808 {
        case 42:
            switch 128 {
            case 418:
                for fi4 := 0; fi4 < 11; fi4++ {
                    var v5 bool = (!true)
                    _ = fi4
                    _ = v5
                }
            case -12:
                v1 = (v1 + (((v1 + (v1 + "hello")) + v1) + ("\n\t" + "\n\t")))
            case 32768:
                var v6 int32 = -2147483648
                _ = v6
            }
        case 128:
            v1 = "\n\t"
        default:
            fmt.Println(v1)
        }
    case -459:
        v1 = ""
    case 1:
        fmt.Println(v1)
    default:
        v7 := GenS1{F0: 0, F1: v1, F2: 1}
        _ = v7
    }
    var v8 bool = true
    switch 4294967295 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 float64 = 3.14159
            _ = fi9
            _ = v10
        }
    case 4294967295:
        var v11 uint16 = 32767
        _ = v11
    case 127:
        v8 = (((0.1 / (1.5 / 1.0)) >= float64((float32(((float64(0.1) - 1.5) + 0.1)) * float32(0.1)))) && v8)
    default:
        v12 := GenS2{F0: 0, F1: 9223372036854775807, F2: 225}
        _ = v12
    }
    switch v1 {
    case "foo bar":
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS3{F0: 377, F1: (v1 > "hello")}
                _ = fi14
                _ = v15
            }
            var v16 string = (v1 + ((v1 + (v1 + ((v1 + ((v1 + "\n\t") + (v1 + ("" + (v1 + "foo bar"))))) + v1))) + v1))
            v17 := []int{1}
            _ = fi13
            _ = v16
            _ = v17
        }
    case "\n\t":
        for fi18 := 0; fi18 < 5; fi18++ {
            v1 = "\n\t"
            fmt.Println(v8, fi18, fi18)
            _ = fi18
        }
    default:
        fmt.Println(v8)
    }
    v8 = ((127 < 348) || (1e10 > (3.14159 - 0.0)))
    v19 := GenS3{F0: 0, F1: ((0.1 / (3.14159 * (3.14159 + ((-2.25 * float32(0.1)) + 0.0)))) != -2.25)}
    v20 := GenS4{F0: 0}
    v21 := GenS3{F0: 255, F1: v8}
    v8 = false
    v22 := GenS2{F0: 42, F1: 0, F2: 57}
    v1 = "hello"
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v8
}
