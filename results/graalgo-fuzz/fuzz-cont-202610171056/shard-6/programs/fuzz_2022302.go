package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint8 }
func genFn0(p0 int64, p1 uint8, p2 int64) bool {
    p0 += p2
    v5 := GenS1{F0: (true || true)}
    var v6 uint = 4294967295
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 uint = 594
    v2 := make(map[int]bool)
    var v3 int16 = 42
    v4 := make(map[string]int)
    fmt.Println(v1, v3)
    if (2147483647 <= -9223372036854775808) {
        var v7 bool = false
        _ = v7
    }
    var v8 int8 = -128
    if true {
        v1 += (^(v1))
    } else {
        v1 -= v1
        v9 := []int{2147483647, 127}
        var v10 float32 = 1e10
        _ = v10
        _ = v9
    }
    var v11 bool = true
    v1 = (^(v1))
    switch -1 {
    case 127:
        v12 := GenS2{F0: v8, F1: 128}
        _ = v12
    default:
        v13 := GenS3{F0: 763, F1: (^(v1)), F2: 1}
        _ = v13
    }
    var v14 float64 = (0.1 - 0.0)
    fmt.Println(v1)
    switch 149 {
    case 32767:
        v15 := []int{-516}
        _ = v15
    }
    var v16 uint = v1
    for fi17 := 0; fi17 < 9; fi17++ {
        if v11 {
            switch fi17 {
            case 9223372036854775807:
                var v18 uint16 = 32767
                _ = v18
            default:
                fmt.Println(v11)
            }
            v8 -= v8
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            delete(v2, (fi19 | -2147483648))
            _ = fi19
        }
        v3 -= v3
        _ = fi17
    }
    switch -473 {
    case 128:
        v11 = true
    case -74:
        if (-454 > 2147483647) {
            v20 := v4[((("foo bar" + "a") + "hello") + "a")]
            v3 = (v3 - (v3 & 42))
            _ = v20
        } else {
            if v11 {
                fmt.Println(v3, v3)
                for fi21 := 0; fi21 < 17; fi21++ {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 int = fi22
                        var v24 bool = false
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v11 = false
                        _ = fi25
                    }
                    for fi26 := 0; fi26 < 5; fi26++ {
                        var v27 float64 = v14
                        v14 -= 1e10
                        _ = fi26
                        _ = v27
                    }
                    _ = fi21
                }
                switch -682 {
                case 9223372036854775807:
                    v2[2147483647] = true
                default:
                    var v28 bool = true
                    _ = v28
                }
            }
            var v29 float64 = v14
            for fi30 := 0; fi30 < 13; fi30++ {
                v11 = ("a" > "hello")
                fmt.Println(fi30)
                _ = fi30
            }
            _ = v29
        }
    case 32768:
        v31 := []int{-955, 32767}
        _ = v31
    default:
        v2[-952] = v11
    }
    v32 := []int{42, -539}
    v16 -= 42
    for fi33 := 0; fi33 < 5; fi33++ {
        if len(v32) > 0 {
            v32[0] = fi33
        }
        v11 = v11
        for fi34 := 0; fi34 < 14; fi34++ {
            v32 = append(v32, (fi33 >> ((v16) & 63)))
            if len(v32) > 0 {
                v32[0] = -724
            }
            v35 := v2[(fi34 / ((705) | 1))]
            _ = len(v32)
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    v36 := GenS3{F0: 42, F1: 65535, F2: 227}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v8
}
