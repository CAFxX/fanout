package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
func genFn0() uint64 {
    switch -549 {
    case -979:
        v5 := GenS1{F0: 182}
        _ = v5
    }
    var v6 bool = false
    v7 := []int{284, 32767, 42, 1}
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int64, p1 int8) string {
    var v8 uint8 = 53
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v8
    return "hello"
}

func main() {
    var v1 uint32 = 748
    v2 := []uint32{v1, v1}
    var v3 int32 = 1
    var v4 float32 = float32((3.14159 - -2.25))
    fmt.Println(v3, v1)
    var v9 uint64 = 42
    v10 := GenS2{F0: 685, F1: 975, F2: 0}
    v9 = (v9 * (v9 & (^(v9))))
    switch (("" + "") + "hello") {
    case "foo bar":
        v1 += v1
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        if false {
            if false {
                v1 -= 4294967295
            } else {
                if (1 != 127) {
                    v11 := GenS1{F0: v3}
                    _ = v11
                }
                if (0 >= 255) {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        fmt.Println(v1, fi12)
                        var v13 float32 = (v4 + (float32((float64(v4) * (3.14159 - 1e10))) - v4))
                        _ = fi12
                        _ = v13
                    }
                    switch "a" {
                    case "\n\t":
                        v3 = 384
                    case "a":
                        if false {
                            fmt.Println(v4, v9, v1)
                        }
                    case "hello":
                        var v14 uint = 42
                        _ = v14
                    default:
                        if (((v9 | genFn0()) > v9) && (2 > 2)) {
                            var v15 int64 = -481
                            var v16 string = "a"
                            _ = v15
                            _ = v16
                        }
                    }
                    fmt.Println(v1, v9)
                } else {
                    var v17 string = "a"
                    _ = v17
                }
            }
            v2 = append(v2, (v1 | v1))
        } else {
            var v18 bool = (1.5 < (1.5 + 1e10))
            v9 += genFn0()
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        fmt.Println(v4)
        _ = fi19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        if (127 == -128) {
            var v21 uint32 = (^(v1))
            v9 = (v9 - (v9 & (v9 - (^(v9)))))
            _ = v21
        }
        _ = fi20
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            fmt.Println(v1, v4)
            _ = fi23
        }
        var v24 int8 = 100
        v9 = (v9 & v9)
        _ = fi22
        _ = v24
    }
    var v25 uint = 65535
    fmt.Println(v1, v25)
    v26 := GenS2{F0: 480, F1: (^(v1)), F2: 255}
    for fi27 := 0; fi27 < 15; fi27++ {
        fmt.Println(v3, v3)
        _ = fi27
    }
    v1 = (v1 + 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
