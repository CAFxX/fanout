package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int32 }
func genFn0(p0 uint32) uint8 {
    var v4 uint16 = 736
    var v5 uint8 = 0
    if false {
        v4 += v4
        var v6 int = -454
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint = 42
    var v2 bool = (!(-7 >= 105))
    v3 := make(map[uint32]int)
    v2 = true
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint64 = 18446744073709551615
            v9 += (v9 ^ v9)
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: 1, F1: -1}
            v2 = v2
            _ = fi10
            _ = v11
        }
        _ = fi7
    }
    _ = len(v3)
    v1 = v1
    var v12 int32 = -383
    var v13 float64 = 3.14159
    fmt.Println(v1, v1)
    if (42 > 4294967295) {
        var v14 bool = false
        _ = v14
    }
    v13 = 0.1
    if (-32768 > 1000) {
        if (true && v2) {
            fmt.Println(v2, v2, v2)
        }
        _ = len(v3)
        if true {
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS1{F0: -567, F1: v12}
                if false {
                    var v17 float32 = (((3.14159 / (((3.14159 - -2.25) * 0.0) - float32(1e10))) / 3.14159) * (0.1 * (((0.0 / ((-2.25 + -2.25) + 0.0)) - (-2.25 * (1e10 / 1.0))) - 3.14159)))
                    if v2 {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            v19 := GenS2{F0: -1}
                            _ = fi18
                            _ = v19
                        }
                    }
                    _ = v17
                }
                var v20 bool = ((v13 * (((v13 - v13) + v13) + (v13 + 1.5))) != 1e10)
                _ = fi15
                _ = v16
                _ = v20
            }
        } else {
            v21 := GenS1{F0: 1, F1: 1}
            v22 := []int{}
            v23 := GenS3{F0: 2}
            _ = v21
            _ = v22
            _ = v23
        }
    } else {
        fmt.Println(v2, v1, v2)
        fmt.Println(v12, v12, v13)
        v12 = v12
    }
    v24 := GenS4{F0: "hello", F1: 1}
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 bool = true
        _ = fi25
        _ = v26
    }
    var v27 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v27
    _ = v3
}
