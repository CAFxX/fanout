package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) int {
    v4 := []int{-242, 0, 32768}
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi6
        }
        _ = fi5
    }
    var v7 float64 = 0.0
    _ = p0
    _ = v4
    _ = v7
    return -40
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]bool)
    var v3 uint8 = 255
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int = int(v3)
        for fi10 := 0; fi10 < 16; fi10++ {
            fmt.Println(fi10, v1, fi10)
            var v11 bool = false
            v1 = 2147483647
            fmt.Println(fi8)
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 float32 = -2.25
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := []int{(fi13 & 736)}
        var v15 int = genFn0(0)
        v1 += 102
        _ = fi13
        _ = v14
        _ = v15
    }
    if true {
        delete(v2, -39)
        var v16 int64 = 296
        _ = v16
    }
    if false {
        if (0 != 65535) {
            v1 = (v1 * (v1 * v1))
        } else {
            v12 -= (((v12 * (0.0 / (v12 - ((float32((0.0 + 0.0)) * ((((v12 / v12) * v12) + v12) / float32((3.14159 * (((3.14159 - 1.5) / (0.0 - 1e10)) + 1.5))))) * v12)))) + 3.14159) / ((v12 * v12) / (v12 - v12)))
            var v17 string = "foo bar"
            _ = v17
        }
        v12 = 0.1
    }
    var v18 uint64 = 762
    if true {
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v12, v3, v1)
            _ = fi19
        }
        var v20 int32 = -1
        _ = v20
    } else {
        v3 += (v3 ^ v3)
        v3 -= (v3 ^ (v3 - (^(v3))))
    }
    if true {
        if false {
            var v21 uint32 = 4294967295
            fmt.Println(v3, v3, v3)
            switch 32767 {
            case 9223372036854775807:
                var v22 uint32 = (^(v1))
                _ = v22
            case 933:
                v3 = v3
            case -117:
                if (4294967295 <= 4294967295) {
                    v1 = 75
                    v3 = v3
                    v12 -= float32((1.5 - 0.1))
                }
            default:
                var v23 string = "foo bar"
                _ = v23
            }
            _ = v21
        }
    } else {
        var v24 uint16 = 259
        var v25 int16 = -782
        v26 := GenS1{F0: 1, F1: (v24 / ((222) | 1)), F2: v1}
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 0, F1: 2, F2: (^(v1))}
    for fi28 := 0; fi28 < 4; fi28++ {
        if true {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 uint16 = 255
                _ = fi29
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := GenS2{F0: v1}
                var v33 int32 = 1
                if true {
                    var v34 uint = 0
                    _ = v34
                } else {
                    fmt.Println(fi31)
                    var v35 string = ""
                    _ = v35
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = fi28
    }
    v1 = 697
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v27
    _ = v3
}
