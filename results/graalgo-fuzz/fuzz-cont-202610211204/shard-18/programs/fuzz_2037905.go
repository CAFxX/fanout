package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint }
type GenS4 struct { F0 uint8; F1 uint }
func genFn0(p0 int16) int8 {
    p0 += (p0 | p0)
    p0 -= p0
    v5 := []int{}
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]bool)
    var v3 int16 = 942
    var v4 uint16 = 0
    var v6 uint16 = (v4 % ((335) | 1))
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 string = v1
        for fi9 := 0; fi9 < 15; fi9++ {
            delete(v2, fi7)
            if true {
                v10 := GenS1{F0: "\n\t"}
                if true {
                    fmt.Println(v1)
                    if ((true && (!(false || false))) && true) {
                        if (uint(v6) < 0) {
                            v11 := GenS1{F0: "hello"}
                            var v12 uint = 65535
                            _ = v11
                            _ = v12
                        }
                        var v13 string = v1
                        v14 := GenS2{F0: 966, F1: 1}
                        _ = v13
                        _ = v14
                    }
                }
                var v15 int16 = 145
                _ = v10
                _ = v15
            } else {
                if false {
                    var v16 int = fi7
                    if (!false) {
                        v6 += v6
                        var v17 uint16 = 0
                        var v18 string = v1
                        _ = v17
                        _ = v18
                    }
                    _ = v16
                }
                var v19 bool = false
                _ = v19
            }
            v2[-457] = false
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v20 := GenS2{F0: 0, F1: 2}
    switch v1 {
    case "foo bar":
        var v21 uint32 = 109
        _ = v21
    default:
        for fi22 := 0; fi22 < 20; fi22++ {
            fmt.Println(fi22, v6, v1)
            _ = fi22
        }
    }
    v4 = v6
    delete(v2, 128)
    var v23 uint16 = v4
    switch 560 {
    case 9223372036854775807:
        var v24 string = v1
        _ = v24
    case 2147483647:
        if false {
            fmt.Println(v6, v1)
            v6 -= 32767
        } else {
            fmt.Println(v4, v3, v6)
        }
    case 1:
        v25 := []int{}
        _ = v25
    default:
        if ("a" <= (((v1 + ("a" + v1)) + v1) + v1)) {
            var v26 int16 = v3
            v27 := GenS3{F0: int64(v3), F1: (^(v6)), F2: 398}
            _ = v26
            _ = v27
        } else {
            v28 := GenS3{F0: 9223372036854775807, F1: (^(v4)), F2: 18446744073709551615}
            _ = v28
        }
    }
    switch 0 {
    case -647:
        switch -367 {
        case -2147483648:
            v29 := GenS4{F0: 1, F1: 65535}
            _ = v29
        }
    case 32767:
        if (42 >= 559) {
            v3 += 32767
        }
    default:
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := v2[fi30]
            _ = fi30
            _ = v31
        }
    }
    v4 += (v4 * v23)
    v32 := []int{2147483647, 32767, 42}
    fmt.Println(v1, v3)
    var v33 uint = 807
    var v34 uint = (v33 << ((4294967295) & 63))
    v2[4294967295] = (((351 < 857) && (true && true)) || true)
    v35 := GenS3{F0: -1, F1: 2, F2: v33}
    var v36 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v6
}
