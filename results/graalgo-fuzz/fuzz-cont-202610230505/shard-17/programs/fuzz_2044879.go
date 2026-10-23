package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int32; F2 uint64 }
func genFn0() uint64 {
    v6 := GenS1{F0: 232, F1: 42}
    var v7 int32 = 187
    v7 -= v7
    var v8 uint64 = 0
    _ = v6
    _ = v7
    _ = v8
    return (v8 & (^(v8)))
}
func genFn1() uint64 {
    if true {
        if (127 > 72) {
            v9 := GenS1{F0: 128, F1: -1}
            var v10 uint32 = 42
            var v11 int32 = -2147483648
            _ = v10
            _ = v11
            _ = v9
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 11; fi13++ {
                    var v14 uint = 18446744073709551615
                    v15 := []int{fi12, -2147483648}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := GenS1{F0: 127, F1: 1}
                    var v18 uint32 = 65535
                    v18 = (v18 / ((v18) | 1))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi12
            }
        }
    }
    switch 128 {
    case 2147483647:
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := []int{fi19, fi19}
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS2{F0: 1}
                v23 := GenS3{F0: -9223372036854775808, F1: -1, F2: 466}
                _ = fi21
                _ = v22
                _ = v23
            }
            if (!(0 < 2)) {
                var v24 uint64 = 18446744073709551615
                v20 = append(v20, (^(fi19)))
                _ = v24
            }
            _ = fi19
            _ = v20
        }
    case -9223372036854775808:
        v25 := GenS1{F0: 97, F1: -2147483648}
        _ = v25
    }
    var v26 int = 628
    _ = v26
    return 474
}

func main() {
    var v1 int8 = 49
    var v2 int64 = 42
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint = 4294967295
    v2 = (v2 + (v2 % ((v2) | 1)))
    v27 := []int{9223372036854775807}
    delete(v4, -637)
    v1 -= v1
    var v28 int64 = v2
    if ((true || (!((0.1 * 1e10) <= 3.14159))) && ((255 != uint16(v5)) && true)) {
        if ((1e10 * 3.14159) == (3.14159 * 3.14159)) {
            _ = len(v27)
        }
        fmt.Println(v5, v28)
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v1)
        _ = fi29
    }
    _ = len(v27)
    var v30 int8 = (v1 | v1)
    v1 = v1
    if (!false) {
        v28 = v28
        if false {
            v30 = 10
            if ((0.0 - -2.25) == (-2.25 / 1e10)) {
                v5 += v5
                for fi31 := 0; fi31 < 6; fi31++ {
                    if false {
                        var v32 uint64 = 2
                        v4[fi31] = int(v5)
                        _ = len(v27)
                        _ = v32
                    }
                    _ = fi31
                }
                v1 -= 42
            }
        } else {
            v1 = v30
        }
    }
    v1 = (^(v1))
    v28 -= v28
    var v33 uint8 = 33
    var v34 int64 = (v28 / ((0) | 1))
    for fi35 := 0; fi35 < 15; fi35++ {
        fmt.Println(v1, v2, v2)
        _ = fi35
    }
    var v36 uint8 = (^(v33))
    v36 -= 144
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
