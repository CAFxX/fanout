package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0() uint16 {
    switch -608 {
    case -21:
        v5 := GenS1{F0: 485}
        _ = v5
    case 0:
        var v6 string = "\n\t"
        _ = v6
    case 802:
        var v7 float64 = 3.14159
        _ = v7
    default:
        v8 := GenS1{F0: 216}
        _ = v8
    }
    v9 := []int{}
    var v10 uint = 18446744073709551615
    v10 += (v10 * (v10 - 2))
    _ = v10
    _ = v9
    return 872
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]int)
    var v3 uint = 255
    var v4 uint = v3
    v11 := GenS1{F0: 1}
    v12 := GenS1{F0: 27}
    if (!false) {
        _ = len(v2)
    }
    switch 901 {
    case 32767:
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v13 uint16 = 1
            _ = v13
        case 127:
            if true {
                if (0 < -128) {
                    v1 = v1
                    v4 = 65535
                } else {
                    if ((^(v3)) < v3) {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            v15 := GenS1{F0: 42}
                            fmt.Println(v4, fi14, v1)
                            _ = len(v2)
                            _ = fi14
                            _ = v15
                        }
                        var v16 int16 = 591
                        _ = v16
                    }
                }
                if (false || false) {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v3 = v3
                        switch fi17 {
                        case 127:
                            fmt.Println(fi17, fi17)
                        case 128:
                            if (100 != -46) {
                                v18 := GenS2{F0: 948}
                                v19 := []int{}
                                _ = v18
                                _ = v19
                            }
                        }
                        v1 = v1
                        _ = fi17
                    }
                } else {
                    _ = len(v2)
                }
                v20 := GenS2{F0: 4294967295}
                _ = v20
            }
        }
    case 128:
        v21 := GenS1{F0: 32767}
        _ = v21
    case -214:
        var v22 int16 = 42
        _ = v22
    }
    v4 += v3
    var v23 uint32 = 931
    v3 -= 2
    fmt.Println(v3, v23)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
