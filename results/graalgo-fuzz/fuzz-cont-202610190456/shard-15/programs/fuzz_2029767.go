package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint64 }
func genFn0(p0 uint16) uint {
    p0 -= p0
    p0 -= 65535
    switch "\n\t" {
    case "\n\t":
        v4 := []int{447, -102, 127}
        _ = v4
    }
    switch "a" {
    case "":
        v5 := []int{585, 723, 0, -815}
        _ = v5
    case "a":
        var v6 int = 32768
        _ = v6
    case "\n\t":
        var v7 uint32 = 2147483648
        _ = v7
    }
    _ = p0
    return 1
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int64 = -1
        _ = fi8
        _ = v9
    }
    var v10 float64 = -2.25
    switch -1 {
    case 4294967295:
        var v11 float64 = -2.25
        _ = v11
    }
    v10 = 1.5
    for fi12 := 0; fi12 < 11; fi12++ {
        v10 -= (v10 / v10)
        if false {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 int64 = -892
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: 2147483647, F1: 0, F2: 42}
            v16 := v3[2]
            _ = v15
            _ = v16
        } else {
            if (genFn0(0) != 628) {
                v10 -= 0.0
                for fi17 := 0; fi17 < 3; fi17++ {
                    if true {
                        if true {
                            var v18 float32 = ((3.14159 / 1.5) * -2.25)
                            var v19 float32 = v18
                            if (!true) {
                                var v20 uint8 = 219
                                _ = v20
                            } else {
                                v18 = 0.0
                                v19 = v19
                                var v21 uint32 = 2
                                if (756 != 245) {
                                    switch fi12 {
                                    case 32767:
                                        var v22 int32 = -742
                                        _ = v22
                                    }
                                }
                                var v23 int8 = 100
                                var v24 int32 = 42
                                v10 = (v10 * 1e10)
                                v21 = v21
                                _ = v21
                                _ = v23
                                _ = v24
                            }
                            _ = v18
                            _ = v19
                        }
                    } else {
                        switch fi12 {
                        case 441:
                            var v25 string = "hello"
                            _ = v25
                        case -2147483648:
                            var v26 int16 = -32768
                            _ = v26
                        }
                    }
                    _ = fi17
                }
                v1 = append(v1, ("hello" + "hello"))
            }
        }
        v10 = (v10 / (0.1 / (3.14159 / ((((v10 / (1e10 - v10)) + v10) + 0.1) / v10))))
        _ = fi12
    }
    v10 = v10
    fmt.Println(v10)
    v1 = append(v1, "\n\t")
    var v27 uint8 = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
}
