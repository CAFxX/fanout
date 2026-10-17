package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 bool }
type GenS2 struct { F0 string; F1 uint32; F2 uint16 }
func genFn0(p0 int) int64 {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int32 = 42
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS1{F0: "hello", F1: -360, F2: false}
    _ = p0
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int32, p1 uint64, p2 int) uint32 {
    v8 := GenS1{F0: "a", F1: -380, F2: (!(true && (405 == 306)))}
    v9 := []int{p2, 4294967295, p2, (-(p2))}
    var v10 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1e10
    var v3 string = "\n\t"
    for fi11 := 0; fi11 < 5; fi11++ {
        if true {
            for fi12 := 0; fi12 < 8; fi12++ {
                v3 = v3
                fmt.Println(v2, v2, fi12)
                _ = fi12
            }
            v1[0] = genFn1(14, 622, fi11)
        } else {
            v13 := []int{}
            _ = v13
        }
        if (0 == 29) {
            var v14 uint8 = 127
            if false {
                v15 := v1[1]
                _ = v15
            }
            var v16 uint = 4294967295
            _ = v14
            _ = v16
        } else {
            var v17 int64 = 0
            var v18 float64 = -2.25
            _ = v17
            _ = v18
        }
        _ = fi11
    }
    var v19 int32 = 1
    var v20 int = 942
    var v21 uint = 18446744073709551615
    v22 := []int{578, (v20 << (((v21 + (v21 << ((v21) & 63)))) & 63))}
    var v23 int8 = 55
    v24 := GenS1{F0: v3, F1: int64(v20), F2: (1e10 >= (float32((v2 / ((v2 * -2.25) + (v2 - ((float64((-2.25 - 3.14159)) - v2) * (v2 * v2)))))) + 0.1))}
    v25 := GenS1{F0: "\n\t", F1: -572, F2: (v3 < "hello")}
    v21 = v21
    v2 += -2.25
    if true {
        switch 1 {
        case 42:
            var v26 uint64 = 0
            _ = v26
        case -949:
            fmt.Println(v20)
        case 4294967295:
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := v1[2]
                _ = fi27
                _ = v28
            }
        default:
            if false {
                if false {
                    v29 := v1[65535]
                    fmt.Println(v23, v29, v3)
                    _ = v29
                } else {
                    switch (v20 * (v20 * (v20 | (v20 % (((^(v20))) | 1))))) {
                    case 128:
                        var v30 uint8 = 255
                        _ = v30
                    case 4294967295:
                        v22 = append(v22, v20)
                    case -1:
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v32 := GenS1{F0: "hello", F1: -9223372036854775808, F2: true}
                            _ = fi31
                            _ = v32
                        }
                    default:
                        v19 = (v19 + v19)
                    }
                    v33 := GenS1{F0: (v3 + v3), F1: 465, F2: (genFn1((v19 >> ((18446744073709551615) & 63)), 292, -288) > 880)}
                    _ = v33
                }
                v34 := GenS2{F0: "hello", F1: uint32(v21), F2: 0}
                _ = v34
            } else {
                v21 += v21
                v19 += v19
            }
        }
        v22 = append(v22, v20)
    }
    var v35 uint = (v21 - 65535)
    v35 -= v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
}
