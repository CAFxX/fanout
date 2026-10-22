package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 int, p1 int32) int32 {
    fmt.Println(p1)
    v4 := GenS1{F0: ("hello" + "")}
    fmt.Println(p0, p0)
    var v5 uint64 = 497
    var v6 uint8 = 127
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (p1 << ((191) & 63))
}

func main() {
    var v1 uint32 = 42
    var v2 uint = 600
    v3 := []string{"", "hello", "hello", ("" + "hello")}
    if true {
        if true {
            if true {
                v1 = 85
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 float64 = 3.14159
                    _ = fi7
                    _ = v8
                }
                v3 = append(v3, "")
            } else {
                v9 := GenS1{F0: ("\n\t" + "hello")}
                _ = v9
            }
            switch 42 {
            case -9223372036854775808:
                if len(v3) > 0 {
                    v3[0] = ("a" + "\n\t")
                }
            case 2147483647:
                v2 = 986
            default:
                for fi10 := 0; fi10 < 17; fi10++ {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v12 := GenS2{F0: 734, F1: "hello", F2: 0}
                        _ = fi11
                        _ = v12
                    }
                    v13 := GenS2{F0: -520, F1: "a", F2: 9223372036854775807}
                    v2 = 255
                    _ = fi10
                    _ = v13
                }
            }
        }
        var v14 int8 = 105
        v14 = (v14 * v14)
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int8 = 100
        _ = fi15
        _ = v16
    }
    v2 = 2
    switch -644 {
    case -811:
        var v17 float32 = 1.5
        _ = v17
    default:
        var v18 int16 = 0
        _ = v18
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v19 string = ("hello" + "hello")
    var v20 uint16 = uint16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
