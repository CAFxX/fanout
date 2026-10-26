package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
func genFn0(p0 string, p1 int8) int8 {
    if (65535 == 4294967295) {
        p1 = 42
        var v4 int64 = -771
        var v5 uint64 = 1
        _ = v4
        _ = v5
    } else {
        var v6 uint32 = 701
        for fi7 := 0; fi7 < 6; fi7++ {
            p1 = p1
            _ = fi7
        }
        _ = v6
    }
    switch 4294967295 {
    case 32768:
        p0 = "a"
    case 2147483647:
        var v8 int64 = 9223372036854775807
        _ = v8
    case 667:
        v9 := []int{-2147483648, 483, 32768, 931}
        _ = v9
    default:
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int16 = 42
            _ = fi10
            _ = v11
        }
    }
    var v12 uint8 = 42
    p1 = p1
    _ = p0
    _ = p1
    _ = v12
    return p1
}

func main() {
    var v1 uint = 453
    var v2 int64 = 0
    var v3 int8 = 100
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 string = (("hello" + (("a" + "hello") + "\n\t")) + "hello")
        _ = fi13
        _ = v14
    }
    if (true || false) {
        v15 := GenS1{F0: 18446744073709551615, F1: 527, F2: true}
        var v16 uint8 = 0
        _ = v15
        _ = v16
    } else {
        if (!(true && (1e10 > 1e10))) {
            v17 := GenS1{F0: 434, F1: 4294967295, F2: ((!((1.5 / (3.14159 + 1e10)) < 1.5)) || false)}
            _ = v17
        }
    }
    if true {
        v1 = 0
        v18 := GenS1{F0: uint64(v1), F1: 1, F2: true}
        if false {
            v19 := []int{}
            var v20 uint64 = 618
            if len(v19) > 0 {
                v19[0] = 4294967295
            }
            _ = v19
            _ = v20
        }
        _ = v18
    }
    v1 = 0
    if true {
        v21 := GenS1{F0: 18446744073709551615, F1: 0, F2: true}
        for fi22 := 0; fi22 < 8; fi22++ {
            v23 := []int{-191, (fi22 * fi22), (fi22 >> ((v1) & 63))}
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int32 = 42
                v23 = append(v23, (fi22 ^ 0))
                _ = fi24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    v3 += v3
    v26 := GenS1{F0: 75, F1: 1, F2: false}
    v27 := []int{1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
