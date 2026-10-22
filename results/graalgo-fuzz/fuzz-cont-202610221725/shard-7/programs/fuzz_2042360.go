package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint, p2 uint) string {
    var v5 int32 = (p0 * (p0 * (^(p0))))
    var v6 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (("a" + "\n\t") + ("" + ("hello" + "\n\t")))
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{2, 2147483648, 65535, 480}
    var v3 uint64 = 0
    var v4 string = "foo bar"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = v4
    fmt.Println(v3, v3, v4)
    v1 = v1
    switch 127 {
    case 880:
        if len(v2) > 0 {
            v2[0] = 2
        }
    case 1:
        var v7 uint = 2
        _ = v7
    case 154:
        if false {
            switch 1 {
            case -1:
                v1 += -175
            case 9223372036854775807:
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 int64 = 9223372036854775807
                    v3 += (^(v3))
                    _ = fi8
                    _ = v9
                }
            default:
                for fi10 := 0; fi10 < 20; fi10++ {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        if true {
                            v2 = append(v2, 2147483647)
                        }
                        var v12 bool = true
                        fmt.Println(v1)
                        _ = fi11
                        _ = v12
                    }
                    _ = fi10
                }
            }
            v13 := GenS1{F0: 42}
            _ = v13
        }
    }
    var v14 int64 = 0
    var v15 uint16 = 1
    fmt.Println(v14, v4)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int32 = int32(v15)
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
