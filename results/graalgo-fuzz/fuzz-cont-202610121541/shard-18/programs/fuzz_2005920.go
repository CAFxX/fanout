package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
func genFn0() int32 {
    v6 := GenS1{F0: 149}
    v7 := []int{-9223372036854775808, 0}
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 630
    var v2 int8 = 1
    var v3 int = 32767
    var v4 int = -9223372036854775808
    v5 := make(map[uint32]bool)
    if false {
        var v8 string = ""
        var v9 uint32 = (v1 << ((888) & 63))
        v4 = v3
        _ = v8
        _ = v9
    } else {
        switch "" {
        case "foo bar":
            v4 = (v3 / (((v4 + (v3 + (v4 << ((65535) & 63))))) | 1))
        case "a":
            switch -620 {
            case 9223372036854775807:
                v2 -= v2
            case 1:
                for fi10 := 0; fi10 < 20; fi10++ {
                    for fi11 := 0; fi11 < 19; fi11++ {
                        var v12 uint64 = 18446744073709551615
                        var v13 uint16 = 0
                        var v14 int = fi10
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                }
            case -787:
                v15 := GenS1{F0: 0}
                _ = v15
            default:
                v16 := GenS1{F0: 111}
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v1 = v1
            _ = len(v5)
            v1 -= v1
            _ = fi17
        }
        if (1 > -9223372036854775808) {
            v18 := GenS1{F0: 255}
            _ = v18
        }
    }
    switch ("" + "\n\t") {
    case "\n\t":
        var v19 bool = (42 < 42)
        _ = v19
    case "foo bar":
        var v20 uint8 = 2
        _ = v20
    }
    v1 -= 65535
    v4 += (v3 << ((1) & 63))
    v21 := GenS2{F0: v4}
    v2 = v2
    fmt.Println(v3, v2)
    var v22 uint16 = 371
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
