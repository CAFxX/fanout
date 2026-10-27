package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 string) int32 {
    switch 32767 {
    case -578:
        var v6 bool = true
        _ = v6
    case 4294967295:
        var v7 string = (("\n\t" + "hello") + "hello")
        _ = v7
    case 2147483647:
        var v8 uint32 = 840
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 = p0
        _ = fi9
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -2147483648
    var v3 uint8 = 200
    v4 := make(map[string]int)
    var v5 uint32 = 2
    fmt.Println(v5, v5, v3)
    v10 := GenS1{F0: 1}
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS2{F0: (!false)}
        _ = fi11
        _ = v12
    }
    var v13 uint32 = v5
    v3 = v3
    _ = len(v4)
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v5, v2)
        v13 = (v5 % ((v13) | 1))
        _ = fi14
    }
    var v15 bool = (false || false)
    var v16 string = ("\n\t" + "hello")
    v17 := GenS2{F0: (uint16(v3) != 0)}
    delete(v1, -9223372036854775808)
    v18 := GenS3{F0: 42, F1: (v5 % (((v5 * (v13 / ((v13) | 1)))) | 1))}
    var v19 string = "hello"
    fmt.Println(v2, v3)
    v16 = ((v16 + ((v16 + "\n\t") + "")) + v16)
    switch -414 {
    case 42:
        switch 9223372036854775807 {
        case 2147483647:
            v20 := []int{0, 2147483647, 32767, 32767}
            _ = v20
        default:
            if (4294967295 < 9223372036854775807) {
                if (32767 > 757) {
                    var v21 uint64 = 191
                    if false {
                        v22 := GenS2{F0: true}
                        _ = v22
                    } else {
                        v23 := v1[-2147483648]
                        _ = v23
                    }
                    _ = v21
                }
                v24 := GenS1{F0: 657}
                fmt.Println(v13)
                _ = v24
            }
        }
    case 9223372036854775807:
        fmt.Println(v2, v15, v13)
    case -2147483648:
        for fi25 := 0; fi25 < 7; fi25++ {
            v4["a"] = fi25
            _ = fi25
        }
    default:
        v2 = (^(v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
