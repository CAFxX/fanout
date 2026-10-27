package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
func genFn0() int8 {
    v5 := GenS1{F0: false, F1: 533, F2: "a"}
    var v6 bool = ((9223372036854775807 != 0) || false)
    v6 = v6
    v7 := GenS2{F0: 100}
    _ = v5
    _ = v6
    _ = v7
    return -47
}
func genFn1() uint {
    v8 := GenS1{F0: false, F1: -32768, F2: ("hello" + ("a" + ("a" + "\n\t")))}
    v9 := []int{127, 4294967295, 127}
    v10 := GenS2{F0: 24}
    _ = v10
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 1
    var v3 uint32 = 0
    var v4 float64 = 0.1
    v3 -= (v3 % (((v3 ^ (v3 * 0))) | 1))
    var v11 int8 = 0
    if true {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS2{F0: v11}
            _ = fi12
            _ = v13
        }
    } else {
        var v14 int64 = 304
        var v15 uint = uint(v2)
        _ = v14
        _ = v15
    }
    switch "" {
    case "a":
        var v16 uint16 = 1
        _ = v16
    case "\n\t":
        var v17 uint = 172
        _ = v17
    case "foo bar":
        v18 := GenS1{F0: (((true || (218 >= v3)) && true) && false), F1: 32767, F2: "hello"}
        _ = v18
    }
    var v19 int8 = 4
    if (42 <= -110) {
        var v20 int32 = 75
        v21 := GenS3{F0: 65535, F1: (9223372036854775807 <= 344), F2: 1000}
        v1[-9223372036854775808] = true
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v11 = v19
        for fi23 := 0; fi23 < 5; fi23++ {
            _ = len(v1)
            _ = fi23
        }
        var v24 int = 9223372036854775807
        _ = fi22
        _ = v24
    }
    switch 1 {
    case 219:
        if true {
            v3 -= 2147483647
        }
    default:
        v25 := v1[-2147483648]
        _ = v25
    }
    if (false || true) {
        if (((42 >= 42) || (1000 != 1)) && (42 <= 1000)) {
            for fi26 := 0; fi26 < 7; fi26++ {
                v27 := []int{fi26}
                if (42 >= 2147483647) {
                    v3 -= v3
                    if (!true) {
                        v28 := GenS3{F0: 65535, F1: ((false && false) && true), F2: -1}
                        _ = v28
                    } else {
                        if true {
                            v2 = v2
                            v29 := GenS3{F0: 32767, F1: ((v2 != v2) && false), F2: 492}
                            _ = v29
                        }
                        v27 = append(v27, fi26)
                        var v30 uint8 = 128
                        _ = v30
                    }
                }
                _ = fi26
                _ = v27
            }
        }
        var v31 uint8 = 182
        fmt.Println(v4)
        _ = v31
    } else {
        v32 := GenS3{F0: 42, F1: ((255 != 1) || (1000 != 0)), F2: 0}
        _ = v32
    }
    fmt.Println(v2, v4, v2)
    var v33 int64 = int64(v19)
    v2 += 1
    v11 += v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
