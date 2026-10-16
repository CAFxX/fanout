package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 bool = (!(112 != 128))
    var v2 bool = (((!false) && true) && (!v1))
    var v3 float64 = 1e10
    var v4 float32 = 1e10
    switch "hello" {
    case "a":
        var v5 bool = (1 == 2147483648)
        _ = v5
    case "foo bar":
        switch 4294967295 {
        case 1:
            fmt.Println(v2, v3, v1)
        case 127:
            var v6 int = -895
            _ = v6
        case 42:
            var v7 int64 = 839
            _ = v7
        default:
            v8 := GenS1{F0: -2147483648}
            _ = v8
        }
    case "\n\t":
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    fmt.Println(v4, v4)
    switch "a" {
    case "a":
        v10 := []int{0, 0, 128, 2147483647}
        _ = v10
    case "hello":
        v11 := GenS1{F0: -516}
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 19; fi12++ {
            switch 2147483647 {
            case -1:
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint = 255
                    _ = fi13
                    _ = v14
                }
            case 2147483647:
                fmt.Println(v3, v3, v3)
            case 42:
                v4 -= 0.0
            default:
                var v15 int16 = 42
                _ = v15
            }
            switch fi12 {
            case -2147483648:
                v16 := []int{fi12, fi12}
                _ = v16
            case 539:
                fmt.Println(v4, v1)
            default:
                for fi17 := 0; fi17 < 17; fi17++ {
                    if true {
                        var v18 float64 = (v3 / ((1e10 / (1e10 * -2.25)) * (((((((v3 + (v3 - v3)) * 1e10) / (v3 - 3.14159)) - 3.14159) * v3) + 1e10) / 3.14159)))
                        switch fi17 {
                        case 4294967295:
                            fmt.Println(v1, fi17, v2)
                        case 128:
                            if v2 {
                                v19 := []int{fi17}
                                _ = v19
                            }
                        }
                        _ = v18
                    } else {
                        v4 += v4
                        v4 = (v4 / -2.25)
                        v4 = (((v4 / float32(v3)) * v4) - v4)
                    }
                    _ = fi17
                }
            }
            _ = fi12
        }
    default:
        var v20 uint16 = 227
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            v2 = (!(true && v2))
            var v23 float32 = (-2.25 / -2.25)
            v3 = 0.0
            _ = fi22
            _ = v23
        }
        var v24 float64 = v3
        v25 := []int{(fi21 - (fi21 / ((359) | 1))), fi21, (fi21 & 128), 32768}
        _ = fi21
        _ = v24
        _ = v25
    }
    var v26 int32 = -1
    fmt.Println(v1, v1, v26)
    switch "hello" {
    case "foo bar":
        v26 -= (v26 + v26)
    default:
        v1 = (165 < 1)
    }
    var v27 int32 = -1
    switch 9223372036854775807 {
    case 0:
        var v28 uint8 = 102
        _ = v28
    }
    var v29 float64 = (1e10 / v3)
    var v30 uint16 = 906
    fmt.Println(v30)
    fmt.Println(v30, v29)
    var v31 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
