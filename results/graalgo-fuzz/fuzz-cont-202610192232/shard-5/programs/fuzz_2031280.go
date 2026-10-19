package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint32) int8 {
    var v6 int16 = 32767
    if (!p0) {
        fmt.Println(v6, p0, v6)
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 int = -377
    var v5 int8 = 0
    v7 := v2[(v4 << ((65535) & 63))]
    if true {
        v8 := []int{v4, 127}
        v9 := GenS1{F0: (!(489 < 312))}
        var v10 uint64 = 4294967295
        _ = v10
        _ = v8
        _ = v9
    }
    switch ("\n\t" + "") {
    case "\n\t":
        v11 := GenS2{F0: v5, F1: 42}
        _ = v11
    case "":
        if (!(!true)) {
            var v12 int16 = 989
            v13 := GenS3{F0: 255, F1: 127, F2: -49}
            v14 := v1[364]
            _ = v12
            _ = v13
            _ = v14
        } else {
            v15 := v2[-779]
            _ = v15
        }
    }
    v5 = (v5 / ((v5) | 1))
    v5 -= (-(v5))
    v4 = (^(v4))
    if true {
        v5 = (-(v5))
    } else {
        switch (v4 % ((v4) | 1)) {
        case 128:
            if false {
                var v16 float64 = 1.5
                v17 := GenS4{F0: 42}
                for fi18 := 0; fi18 < 18; fi18++ {
                    switch 622 {
                    case -2147483648:
                        v7 -= v7
                    case 858:
                        v19 := GenS1{F0: (42 != 1)}
                        _ = v19
                    default:
                        if (4294967295 < (v7 ^ (v7 | 1))) {
                            v20 := GenS1{F0: true}
                            _ = v20
                        } else {
                            var v21 int16 = 32767
                            v22 := GenS1{F0: true}
                            delete(v1, (v7 * 65535))
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi18
                }
                _ = v16
                _ = v17
            }
        case 2147483647:
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 string = "a"
                _ = fi23
                _ = v24
            }
        }
    }
    fmt.Println(v7, v4, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
