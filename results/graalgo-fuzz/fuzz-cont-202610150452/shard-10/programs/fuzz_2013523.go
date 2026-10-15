package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
func genFn0() uint32 {
    var v6 uint16 = 1
    switch 501 {
    case -2147483648:
        if true {
            var v7 int8 = 0
            v6 -= v6
            _ = v7
        }
    case 32768:
        if ((true && ((-9223372036854775808 < 117) || (42 >= 123))) && false) {
            fmt.Println(v6)
        }
    case -9223372036854775808:
        if false {
            v8 := GenS1{F0: 0}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    _ = v6
    return 65535
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 128
    var v3 uint32 = 2
    var v4 float32 = 3.14159
    var v5 uint32 = v3
    var v11 int8 = 95
    fmt.Println(v5, v3, v11)
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            if false {
                if true {
                    var v14 float64 = (1.5 + -2.25)
                    _ = v14
                } else {
                    v15 := GenS1{F0: -32}
                    if false {
                        v2 -= v2
                        v5 = 65535
                        var v16 float64 = -2.25
                        _ = v16
                    } else {
                        var v17 uint16 = 32767
                        _ = v17
                    }
                    _ = v15
                }
            } else {
                v18 := GenS1{F0: -370}
                _ = v18
            }
            var v19 float64 = ((3.14159 - float64(-2.25)) + 3.14159)
            _ = fi13
            _ = v19
        }
        _ = fi12
    }
    v4 = float32(float64(0.0))
    fmt.Println(v5, v3, v11)
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            v3 += (v3 & v5)
            v2 -= (^(v2))
            if ((!(!true)) && (32767 < 0)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v22 := GenS1{F0: -1}
            v23 := GenS1{F0: -855}
            switch fi20 {
            case 939:
                v2 += 140
            default:
                var v24 float64 = 3.14159
                _ = v24
            }
            v4 += 1.5
            var v25 int = (fi21 * fi20)
            v3 = (^(v3))
            var v26 uint8 = 138
            v5 -= (v5 >> ((uint(v26)) & 63))
            v26 -= v2
            _ = fi21
            _ = v22
            _ = v23
            _ = v25
            _ = v26
        }
        _ = fi20
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            v3 = v5
            v29 := GenS1{F0: -2147483648}
            v30 := GenS1{F0: -2147483648}
            _ = fi28
            _ = v29
            _ = v30
        }
        v1 = append(v1, v5)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi27
    }
    fmt.Println(v5, v11)
    var v31 int = 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v32 := GenS1{F0: -2147483648}
    v33 := GenS2{F0: v31}
    var v34 int16 = -1
    v5 = genFn0()
    fmt.Println(v3, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
