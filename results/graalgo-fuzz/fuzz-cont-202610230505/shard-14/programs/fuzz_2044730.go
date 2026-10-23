package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0() uint {
    var v6 uint16 = 20
    v6 = (^(v6))
    v7 := []int{-2147483648, 42}
    v7 = append(v7, 42)
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint64 = 172
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    var v4 int32 = 425
    var v5 bool = v3
    v8 := []int{40, 42}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 uint8 = 127
    if v3 {
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                v5 = (2 != 1)
                for fi11 := 0; fi11 < 19; fi11++ {
                    v5 = v3
                    fmt.Println(v4)
                    var v12 uint8 = (v9 ^ v9)
                    _ = fi11
                    _ = v12
                }
            } else {
                v5 = v3
                v1 = v1
            }
            var v13 int8 = 1
            for fi14 := 0; fi14 < 12; fi14++ {
                if (true || (!v3)) {
                    v15 := GenS1{F0: 42, F1: (-(v4)), F2: (v9 - (^(v9)))}
                    _ = v15
                } else {
                    var v16 int64 = 806
                    switch fi14 {
                    case -2147483648:
                        fmt.Println(v5, v1, v4)
                    case 0:
                        v16 = v16
                    case -371:
                        var v17 float32 = 0.0
                        _ = v17
                    default:
                        var v18 int64 = 864
                        _ = v18
                    }
                    if len(v8) > 0 {
                        v8[0] = fi10
                    }
                    v19 := GenS1{F0: v16, F1: -2147483648, F2: v9}
                    _ = v16
                    _ = v19
                }
                _ = fi14
            }
            _ = fi10
            _ = v13
        }
    } else {
        var v20 float32 = 1e10
        for fi21 := 0; fi21 < 4; fi21++ {
            v20 -= 1e10
            v5 = v5
            v8 = append(v8, (fi21 % ((fi21) | 1)))
            _ = fi21
        }
        _ = v20
    }
    fmt.Println(v5, v4)
    v2[911] = 426
    v3 = v3
    v1 = v1
    v4 += -2147483648
    var v22 float32 = 0.0
    var v23 float64 = (1.5 - (0.1 + 0.1))
    v24 := GenS2{F0: v4}
    for fi25 := 0; fi25 < 14; fi25++ {
        if (-1 >= 42) {
            v3 = v5
        } else {
            var v26 int64 = int64(v1)
            _ = v26
        }
        _ = fi25
    }
    v27 := GenS3{F0: 2147483647, F1: 1}
    if ((v1 - (v1 << ((18446744073709551615) & 63))) < (v1 * (v1 & (v1 << ((4294967295) & 63))))) {
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 int32 = (v4 % ((v4) | 1))
            v5 = v5
            _ = fi28
            _ = v29
        }
    } else {
        var v30 int8 = 1
        v31 := GenS2{F0: (v4 & (v4 - 269))}
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
