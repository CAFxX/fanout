package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint) uint16 {
    fmt.Println(p0)
    if false {
        if false {
            v4 := GenS1{F0: 9223372036854775807}
            p0 = (^(p0))
            _ = v4
        } else {
            p0 = 18446744073709551615
        }
    }
    fmt.Println(p0)
    _ = p0
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    var v3 float64 = 1.5
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 float64 = v3
            _ = fi6
            _ = v7
        }
        var v8 uint64 = 87
        if (v3 != 1.5) {
            var v9 float64 = ((((v3 / -2.25) + v3) / (0.0 - -2.25)) / (1.5 * ((float64((1e10 / (-2.25 - ((3.14159 + 3.14159) * 1e10)))) + float64((3.14159 + 3.14159))) - 1.5)))
            _ = v9
        } else {
            var v10 uint16 = 2
            if (!false) {
                switch fi5 {
                case 127:
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 int32 = 0
                        var v13 float32 = 1.5
                        var v14 int = -996
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case 32768:
                    switch fi5 {
                    case 2147483647:
                        v10 += 494
                    }
                case 1:
                    var v15 uint8 = 128
                    _ = v15
                default:
                    fmt.Println(v3)
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 int8 = 100
                    _ = fi16
                    _ = v17
                }
            } else {
                v8 = v8
                if true {
                    var v18 uint16 = 42
                    _ = v18
                } else {
                    v1[fi5] = ((("\n\t" + v2) + "a") + v2)
                    v19 := []int{fi5, int(v10), fi5}
                    _ = v19
                }
                v20 := GenS1{F0: -9223372036854775808}
                _ = v20
            }
            v3 += ((float64(((float32((v3 * v3)) - 1e10) * float32(v3))) - v3) + v3)
            _ = v10
        }
        _ = fi5
        _ = v8
    }
    var v21 uint16 = 42
    v22 := v1[922]
    v23 := GenS1{F0: 350}
    v22 = v2
    v2 = v22
    var v24 uint64 = 933
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v24)
        _ = fi25
    }
    fmt.Println(v3)
    v24 = v24
    v26 := v1[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
}
