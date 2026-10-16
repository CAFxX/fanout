package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }

func main() {
    var v1 int8 = 48
    v2 := []int{4294967295, 9223372036854775807, 1, -9223372036854775808}
    var v3 int8 = (v1 & (v1 & v1))
    fmt.Println(v3, v3)
    for fi4 := 0; fi4 < 19; fi4++ {
        v3 -= (v1 / ((v1) | 1))
        _ = fi4
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        if true {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint32 = 0
                var v8 int = fi6
                var v9 float32 = 0.0
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            if true {
                for fi10 := 0; fi10 < 7; fi10++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v11 int = -373
                    var v12 uint16 = 42
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 9; fi13++ {
                    _ = len(v2)
                    if true {
                        v2 = append(v2, (fi13 * fi5))
                    } else {
                        v14 := GenS1{F0: -2147483648}
                        if true {
                            var v15 string = ("\n\t" + "")
                            _ = v15
                        }
                        _ = v14
                    }
                    _ = fi13
                }
                v16 := GenS1{F0: 2147483647}
                if true {
                    switch "foo bar" {
                    case "a":
                        if true {
                            var v17 bool = (fi5 > -1)
                            _ = v17
                        }
                    case "":
                        v3 = v1
                    }
                    v18 := GenS1{F0: 1}
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        switch fi19 {
                        case 32768:
                            var v20 float64 = 0.1
                            _ = v20
                        case 60:
                            v3 += (v1 & v3)
                        default:
                            if len(v2) > 0 {
                                v2[0] = (fi5 | fi19)
                            }
                        }
                        _ = fi19
                    }
                }
                _ = v16
            }
        } else {
            var v21 float32 = -2.25
            v1 = (^(v3))
            _ = v21
        }
        _ = fi5
    }
    v1 = 127
    v1 -= v3
    var v22 uint32 = 1
    var v23 uint64 = 18446744073709551615
    v24 := GenS2{F0: (^(v23)), F1: 1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
