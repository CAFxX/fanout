package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int = 1
    v2 := make(map[string]uint32)
    var v3 int32 = 752
    for fi4 := 0; fi4 < 20; fi4++ {
        switch ("foo bar" + ("a" + ("\n\t" + ""))) {
        case "a":
            v5 := GenS1{F0: -9223372036854775808, F1: ((^(v3)) >= v3)}
            _ = v5
        case "foo bar":
            switch 0 {
            case 240:
                var v6 uint8 = 0
                _ = v6
            }
        case "\n\t":
            var v7 float32 = 0.1
            _ = v7
        default:
            if (2147483647 >= -1) {
                var v8 uint = 590
                _ = v8
            } else {
                v9 := GenS2{F0: 1, F1: 867, F2: fi4}
                _ = v9
            }
        }
        switch (v1 * 32768) {
        case 4294967295:
            v10 := []int{4294967295}
            _ = v10
        case 127:
            v11 := GenS2{F0: 127, F1: 18446744073709551615, F2: int(v3)}
            _ = v11
        }
        v12 := GenS1{F0: 931, F1: (!false)}
        _ = fi4
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 bool = false
        if (!(v14 && v14)) {
            v3 = (v3 | (v3 + (v3 - v3)))
        } else {
            v15 := []int{}
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 int8 = -95
    var v17 bool = false
    v16 += (v16 % ((v16) | 1))
    v18 := GenS2{F0: 127, F1: 0, F2: (v1 - v1)}
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int8 = v16
        v21 := GenS1{F0: -1, F1: true}
        if v17 {
            for fi22 := 0; fi22 < 6; fi22++ {
                fmt.Println(fi22, v17)
                v23 := v2["a"]
                if (!v17) {
                    v24 := []int{fi22}
                    var v25 int = (v1 | (v1 >> ((18446744073709551615) & 63)))
                    _ = v24
                    _ = v25
                } else {
                    v23 = (v23 % (((v23 - (v23 * v23))) | 1))
                    var v26 string = "a"
                    if v17 {
                        v16 -= v16
                    }
                    _ = v26
                }
                v2["\n\t"] = 2147483647
                if v17 {
                    if (((0.0 - float64(1e10)) > 3.14159) && (((0.1 * float32(-2.25)) / 0.1) <= 1.5)) {
                        v3 = 1
                        var v27 uint64 = 419
                        if (!(494 < uint16(v27))) {
                            v28 := GenS3{F0: 1, F1: (v16 >> ((0) & 63)), F2: 2}
                            _ = v28
                        } else {
                            v29 := GenS4{F0: (2 > 872), F1: 0}
                            v30 := []int{}
                            if v17 {
                                fmt.Println(v27, v23)
                                fmt.Println(v3, v20)
                                var v31 int32 = (^(v3))
                                _ = v31
                            } else {
                                var v32 int64 = 0
                                v30 = append(v30, fi22)
                                fmt.Println(fi19, v1)
                                _ = v32
                            }
                            _ = v29
                            _ = v30
                        }
                        v16 -= v20
                        _ = v27
                    }
                }
                _ = fi22
                _ = v23
            }
            v20 -= (^(v16))
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v3 += 1
    v33 := []int{-586, (^(v1)), v1, v1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
}
