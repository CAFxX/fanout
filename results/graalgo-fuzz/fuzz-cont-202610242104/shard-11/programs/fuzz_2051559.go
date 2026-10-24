package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
func genFn0() int16 {
    if false {
        v4 := []int{42, 1, -9223372036854775808}
        switch 90 {
        case 385:
            var v5 float64 = 1.5
            _ = v5
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            if false {
                var v6 uint64 = 286
                v7 := GenS1{F0: 1}
                v6 = (^(v6))
                _ = v6
                _ = v7
            }
        }
        _ = v4
    }
    if (-1 >= 127) {
        var v8 float64 = (((((0.1 + -2.25) + (3.14159 + float64(((1.5 * 1.5) - -2.25)))) - 0.0) - (1e10 / 1.0)) + 1e10)
        var v9 int = 4294967295
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 string = "a"
            v12 := []int{fi10, fi10, (-(fi10)), (fi10 & -2147483648)}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v8
        _ = v9
    } else {
        var v13 uint = 65535
        _ = v13
    }
    var v14 uint64 = 18446744073709551615
    _ = v14
    return 456
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = -1
    v3 := []int{976, 42, -45}
    switch ("" + "\n\t") {
    case "":
        var v15 bool = false
        _ = v15
    case "\n\t":
        if false {
            for fi16 := 0; fi16 < 1; fi16++ {
                if (int16(v1) <= 1) {
                    v2 = v2
                    v2 = v2
                }
                _ = fi16
            }
            v17 := GenS1{F0: 127}
            var v18 int16 = 32767
            _ = v17
            _ = v18
        }
    }
    _ = len(v3)
    v19 := GenS2{F0: 820}
    for fi20 := 0; fi20 < 3; fi20++ {
        v1 = 0
        v1 += 255
        v21 := GenS1{F0: uint8(v1)}
        _ = fi20
        _ = v21
    }
    if true {
        if true {
            for fi22 := 0; fi22 < 20; fi22++ {
                v23 := GenS2{F0: (fi22 ^ (fi22 - (fi22 / ((fi22) | 1))))}
                v1 = 0
                _ = fi22
                _ = v23
            }
        }
    } else {
        var v24 float64 = 1e10
        if (v1 < v1) {
            v2 = v2
            v1 += (v1 - v1)
        } else {
            v1 += 32767
            var v25 uint32 = 517
            if false {
                v2 += 2147483647
                var v26 int8 = 127
                if true {
                    var v27 uint32 = (v25 - v25)
                    v28 := GenS3{F0: v26}
                    fmt.Println(v27, v25, v27)
                    _ = v27
                    _ = v28
                }
                _ = v26
            } else {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                if false {
                    _ = len(v3)
                    v25 = (v25 / ((1) | 1))
                } else {
                    var v29 uint8 = 2
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 uint = 4294967295
                        _ = fi30
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 11; fi32++ {
                        if len(v3) > 0 {
                            v3[0] = 9223372036854775807
                        }
                        _ = fi32
                    }
                    _ = v29
                }
            }
            _ = v25
        }
        _ = v24
    }
    var v33 int32 = 1
    fmt.Println(v33, v2, v1)
    if true {
        v34 := GenS1{F0: 0}
        _ = v34
    } else {
        v35 := GenS1{F0: 1}
        v1 = 255
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
}
