package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint64) uint16 {
    if ((uint8(p0) != 255) && ((-32768 <= 32767) || true)) {
        v5 := GenS1{F0: 255}
        _ = v5
    } else {
        v6 := GenS1{F0: 255}
        v7 := GenS2{F0: -9223372036854775808}
        if false {
            if (288 > 0) {
                if false {
                    var v8 uint32 = 4294967295
                    switch 128 {
                    case 42:
                        v9 := []int{}
                        _ = v9
                    default:
                        v10 := []int{}
                        _ = v10
                    }
                    v11 := []int{737, 1, -2147483648}
                    _ = v11
                    _ = v8
                } else {
                    v12 := GenS2{F0: -608}
                    var v13 int64 = 1
                    var v14 float32 = 0.1
                    _ = v12
                    _ = v13
                    _ = v14
                }
                p0 -= p0
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 int = 42
                    _ = fi15
                    _ = v16
                }
            }
            v17 := GenS2{F0: -815}
            _ = v17
        }
        _ = v6
        _ = v7
    }
    v18 := []int{}
    var v19 uint = 2
    _ = p0
    _ = v18
    _ = v19
    return 327
}

func main() {
    var v1 int32 = -2147483648
    var v2 string = ((("hello" + "") + "\n\t") + "a")
    var v3 int = 638
    v4 := make(map[uint32]int)
    v20 := []int{v3, (^(v3)), int(v1)}
    v2 = (v2 + v2)
    v1 += v1
    if false {
        var v21 int = (v3 / ((40) | 1))
        fmt.Println(v2)
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 int8 = 42
            for fi24 := 0; fi24 < 1; fi24++ {
                v21 = 127
                _ = fi24
            }
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        var v25 uint32 = 65535
        _ = v25
    }
    v2 = (v2 + ((("\n\t" + v2) + v2) + ""))
    v26 := GenS3{F0: 32767, F1: 1}
    var v27 uint32 = 0
    fmt.Println(v2)
    v1 = 1
    v28 := v4[(v27 >> ((18446744073709551615) & 63))]
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            v3 = fi30
            _ = fi30
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint8 = 1
            v33 := GenS2{F0: -1}
            if false {
                v34 := GenS2{F0: 0}
                v27 += v27
                v28 -= (v3 / (((fi31 % (((v28 * v3)) | 1))) | 1))
                v35 := GenS1{F0: 1}
                fmt.Println(v2, v2)
                _ = v34
                _ = v35
            } else {
                v2 = "a"
                fmt.Println(v2, fi29)
            }
            _ = fi31
            _ = v32
            _ = v33
        }
        v27 -= 2147483648
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
