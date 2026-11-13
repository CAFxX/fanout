package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0() uint {
    var v4 int32 = 2147483647
    v5 := GenS1{F0: 9223372036854775807}
    var v6 float32 = 0.0
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS2{F0: 2}
            var v10 int64 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 2
    var v3 uint16 = v2
    for fi11 := 0; fi11 < 2; fi11++ {
        switch -2147483648 {
        case 32767:
            v2 -= 1
        default:
            var v12 uint16 = (^(v2))
            _ = v12
        }
        var v13 uint32 = 2147483648
        var v14 float32 = 1.5
        _ = fi11
        _ = v13
        _ = v14
    }
    var v15 string = "foo bar"
    _ = len(v1)
    v15 = (v15 + ("hello" + (("\n\t" + v15) + v15)))
    switch "a" {
    case "":
        v16 := []int{9223372036854775807, -885, 32767}
        _ = v16
    default:
        fmt.Println(v15)
    }
    if true {
        var v17 uint64 = 42
        switch v15 {
        case "hello":
            v17 = v17
        case "a":
            v18 := GenS3{F0: 232}
            _ = v18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                if false {
                    v17 += (v17 << ((genFn0()) & 63))
                    var v21 uint32 = 849
                    var v22 float32 = float32(0.1)
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
            _ = fi19
        }
        _ = v17
    }
    v3 -= v3
    v23 := []int{}
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
}
