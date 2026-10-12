package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 int8 }
func genFn0(p0 bool) uint8 {
    var v7 float32 = -2.25
    var v8 int8 = 107
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = ("a" + "\n\t")
    var v3 int64 = 0
    var v4 uint32 = 2
    var v5 int64 = v3
    var v6 uint = 0
    fmt.Println(v6, v4, v2)
    v6 -= v6
    v9 := GenS1{F0: 42, F1: (v6 ^ 1)}
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 float32 = 3.14159
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            switch (fi13 ^ fi13) {
            case 291:
                if (((false || false) && (4294967295 != 2)) || (954 > 0)) {
                    v3 += (v3 - (v3 % ((-413) | 1)))
                } else {
                    v4 -= v4
                }
            case -970:
                v5 -= (v3 / ((v3) | 1))
                var v14 uint32 = (^(v4))
                _ = v14
            case 9223372036854775807:
                var v15 uint8 = uint8(v4)
                _ = v15
            }
            var v16 bool = (v4 > uint32(v6))
            _ = fi13
            _ = v16
        }
        if (v5 >= (v3 / (((v5 | (v5 & 1))) | 1))) {
            var v17 uint = (^(v6))
            v18 := []int{fi12, fi12}
            fmt.Println(v4, v17)
            _ = v17
            _ = v18
        }
        _ = fi12
    }
    var v19 string = v2
    fmt.Println(v6)
    var v20 uint8 = 186
    if true {
        v21 := GenS2{F0: 2, F1: (v6 >> ((2) & 63)), F2: 23}
        v5 = (-(v5))
        for fi22 := 0; fi22 < 13; fi22++ {
            v4 = 2
            v23 := GenS1{F0: -2147483648, F1: (v6 << ((uint(v20)) & 63))}
            var v24 uint64 = 177
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = v21
    }
    var v25 float32 = (((1e10 / 1.0) - (float32(1e10) / 3.14159)) * 0.1)
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS1{F0: 1, F1: 83}
        _ = fi26
        _ = v27
    }
    if (112 >= 100) {
        v4 = v4
        if true {
            var v28 int64 = v3
            _ = v28
        }
    } else {
        fmt.Println(v20)
    }
    var v29 uint32 = 2147483648
    var v30 uint64 = 1
    var v31 int8 = 9
    var v32 uint8 = v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
