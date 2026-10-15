package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint64 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{599}
    var v3 int16 = 1
    v4 := make(map[uint32]int)
    var v5 string = "a"
    var v6 uint = 0
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            delete(v4, 0)
            _ = fi9
        }
        _ = fi8
    }
    v4[65535] = 2147483647
    v6 = v6
    switch v5 {
    case "a":
        v5 = "a"
    default:
        v10 := GenS1{F0: v6, F1: 128, F2: v7}
        _ = v10
    }
    fmt.Println(v6)
    if false {
        for fi11 := 0; fi11 < 17; fi11++ {
            v5 = v5
            if (!false) {
                switch fi11 {
                case 32768:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 9223372036854775807:
                    var v12 string = v5
                    _ = v12
                default:
                    var v13 float64 = 3.14159
                    _ = v13
                }
                if false {
                    if len(v2) > 0 {
                        v2[0] = (fi11 >> (((v6 + (v6 | v6))) & 63))
                    }
                }
            } else {
                v7 = 0
                for fi14 := 0; fi14 < 3; fi14++ {
                    if true {
                        fmt.Println(v5)
                        var v15 uint32 = 2147483647
                        v7 -= v7
                        _ = v15
                    } else {
                        _ = len(v1)
                        fmt.Println(fi11, fi11, v3)
                        v7 = 4294967295
                        v16 := GenS1{F0: (v6 | v6), F1: fi14, F2: (v7 << ((v6) & 63))}
                        switch fi11 {
                        case 127:
                            v3 = v3
                            switch -9223372036854775808 {
                            case -2147483648:
                                v5 = "\n\t"
                            case 299:
                                var v17 uint32 = 65535
                                _ = v17
                            case -1:
                                v5 = v5
                            }
                        }
                        _ = v16
                    }
                    v18 := GenS2{F0: ("hello" + (v5 + v5)), F1: 9223372036854775807}
                    fmt.Println(v3, v5)
                    _ = fi14
                    _ = v18
                }
            }
            _ = fi11
        }
        var v19 uint = v6
        var v20 uint64 = v7
        _ = v19
        _ = v20
    } else {
        v5 = (v5 + ("" + (v5 + "a")))
        var v21 string = v5
        _ = v21
    }
    var v22 int = -327
    var v23 float64 = ((1e10 / float64((float32((1e10 * (3.14159 / 1e10))) * 0.1))) * float64(3.14159))
    if false {
        v24 := v4[0]
        v5 = ((v5 + ((v5 + (v5 + (v5 + "a"))) + v5)) + v5)
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
