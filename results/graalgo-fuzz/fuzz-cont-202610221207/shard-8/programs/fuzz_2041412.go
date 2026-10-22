package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0() uint8 {
    if false {
        var v5 int = 66
        fmt.Println(v5)
        _ = v5
    }
    v6 := []int{}
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 uint16, p2 uint) int {
    p0 = -1
    var v7 int32 = -490
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 128
    var v3 string = ("foo bar" + ("a" + ""))
    v4 := make(map[int]string)
    v3 = v3
    v3 = (v3 + (v3 + v3))
    v2 = (v2 << ((2) & 63))
    v8 := GenS1{F0: v3, F1: "hello", F2: true}
    switch -819 {
    case -2147483648:
        v9 := GenS1{F0: "\n\t", F1: "foo bar", F2: ((v2 + (v2 | v2)) < (^(v2)))}
        _ = v9
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            if (1 == 0) {
                v3 = v3
            } else {
                var v11 bool = false
                var v12 uint = 1
                _ = v11
                _ = v12
            }
            if (-2.25 == 0.0) {
                var v13 float32 = (((0.1 + 0.1) + 0.1) + 0.0)
                if false {
                    var v14 int8 = 1
                    v15 := GenS2{F0: 4294967295, F1: v2, F2: 389}
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v17 := GenS3{F0: "foo bar", F1: (v2 ^ genFn0()), F2: (v2 << ((65535) & 63))}
                        _ = fi16
                        _ = v17
                    }
                    var v18 int64 = 9223372036854775807
                    var v19 uint32 = 4294967295
                    _ = v18
                    _ = v19
                }
                _ = v13
            }
            _ = fi10
        }
    }
    if (!(128 >= 32767)) {
        if (42 != 32767) {
            var v20 float64 = (1e10 + (((((-2.25 + float64(0.1)) * 1e10) / 1e10) * 1.5) * (-2.25 - 1.5)))
            _ = v20
        } else {
            v2 -= (v2 * (v2 | v2))
            switch v3 {
            case "":
                v21 := []int{genFn1(42, 622, 4294967295), 4294967295, 37}
                _ = v21
            }
            if false {
                v22 := []int{}
                v23 := GenS4{F0: 0}
                v2 -= v2
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 20; fi24++ {
                    v2 -= genFn0()
                    _ = fi24
                }
            }
        }
        var v25 string = v3
        _ = v25
    }
    v2 += v2
    var v26 float64 = 1e10
    var v27 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
}
