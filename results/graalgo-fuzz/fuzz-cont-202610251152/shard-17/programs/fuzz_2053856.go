package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int16, p1 int, p2 int32) uint64 {
    var v6 bool = false
    v6 = v6
    fmt.Println(p2)
    v6 = ((696 < 0) || ((233 != 0) && (2147483647 != 65535)))
    v7 := GenS1{F0: 2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 bool) uint16 {
    fmt.Println(p0)
    v8 := GenS1{F0: 790}
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 137
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -668
    var v3 uint64 = 4294967295
    var v4 float64 = ((1e10 * (0.0 - ((3.14159 - float64((1e10 - 1.5))) * float64(((((float32(float64(-2.25)) + 1e10) + 1.5) * 1e10) / ((3.14159 - float32(0.0)) - (3.14159 + (-2.25 - 1e10)))))))) * ((-2.25 - (-2.25 / 1e10)) / 0.1))
    var v5 bool = true
    v2 += (v2 % (((v2 << ((4294967295) & 63))) | 1))
    v4 += v4
    v3 = 804
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int32 = (v2 - v2)
        fmt.Println(v3, v3, v10)
        if true {
            v3 = v3
            for fi11 := 0; fi11 < 18; fi11++ {
                switch "hello" {
                case "hello":
                    var v12 uint16 = genFn1(true)
                    _ = v12
                case "\n\t":
                    v4 += v4
                default:
                    var v13 float64 = v4
                    _ = v13
                }
                if (65535 == 32767) {
                    fmt.Println(v3)
                    switch 1 {
                    case 32768:
                        var v14 int16 = 42
                        _ = v14
                    }
                    switch fi11 {
                    case 0:
                        v2 += v10
                    case 690:
                        var v15 uint8 = 112
                        _ = v15
                    default:
                        v16 := GenS1{F0: 42}
                        var v17 int32 = v10
                        v18 := GenS1{F0: 42}
                        v17 = (v2 | (v17 ^ v10))
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                }
                var v19 int32 = v2
                if true {
                    v20 := GenS2{F0: -9223372036854775808}
                    if (!(65535 > 2)) {
                        v21 := []int{}
                        v22 := GenS3{F0: 20}
                        _ = v21
                        _ = v22
                    }
                    var v23 uint64 = genFn0(907, fi11, (v19 % ((0) | 1)))
                    _ = v20
                    _ = v23
                }
                _ = fi11
                _ = v19
            }
        } else {
            v24 := GenS3{F0: 1}
            _ = v24
        }
        _ = fi9
        _ = v10
    }
    v2 = v2
    v1[4294967295] = (2 <= 128)
    v25 := v1[32767]
    if false {
        var v26 int = 32768
        var v27 int16 = 32767
        _ = v26
        _ = v27
    }
    var v28 uint64 = v3
    if (v5 && v5) {
        var v29 bool = true
        fmt.Println(v2)
        _ = v29
    } else {
        v30 := GenS1{F0: 67}
        _ = v30
    }
    v31 := GenS1{F0: 0}
    delete(v1, 42)
    v32 := GenS4{F0: 134}
    v33 := GenS3{F0: 223}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
