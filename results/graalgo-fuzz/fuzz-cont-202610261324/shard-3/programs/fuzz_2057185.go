package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 bool; F2 int32 }
func genFn0() bool {
    v5 := []int{-9223372036854775808, -1, 32768, 257}
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 float64 = 1e10
        var v8 uint = 42
        v7 = v7
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = 1e10
    _ = v5
    _ = v9
    return false
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-9223372036854775808, 32768, 128}
    var v3 uint = 65535
    var v4 uint16 = 165
    if false {
        v10 := GenS1{F0: -32768}
        v3 = v3
        if false {
            var v11 int16 = 32767
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := GenS2{F0: true}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = v10
    }
    var v14 uint = (v3 | (^(v3)))
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS3{F0: (v14 + 1), F1: (!false), F2: 42}
        _ = fi15
        _ = v16
    }
    var v17 int32 = 0
    var v18 float32 = 1.5
    if (196 != 994) {
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 float32 = (3.14159 / ((v18 / (float32(1.5) + v18)) * 3.14159))
                    v18 += v21
                    v17 = (^(v17))
                    var v22 float32 = v21
                    v3 = (^(v14))
                    var v23 uint32 = 247
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    v25 := GenS3{F0: (v14 >> ((v14) & 63)), F1: true, F2: v17}
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 17; fi26++ {
                    fmt.Println(fi19, v17, fi26)
                    v27 := GenS2{F0: true}
                    if (0 < 139) {
                        var v28 uint16 = (v4 + (v4 & 42))
                        v28 = 291
                        var v29 float64 = 1.5
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                }
            }
            var v30 string = "a"
            v1[-9223372036854775808] = 2147483647
            _ = fi19
            _ = v30
        }
        v17 = 2147483647
        var v31 uint64 = 42
        _ = v31
    }
    if false {
        var v32 uint = (v14 * (v3 ^ 1))
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 uint64 = 42
            if ((!(!true)) || true) {
                v35 := GenS2{F0: (4294967295 >= 784)}
                v34 = (v34 - (v34 >> (((v14 - v32)) & 63)))
                if false {
                    v3 += 65535
                    v32 = 65535
                }
                _ = v35
            }
            _ = fi33
            _ = v34
        }
        if true {
            if (true && true) {
                v36 := GenS3{F0: v3, F1: (false && (-45 >= 401)), F2: int32(v32)}
                var v37 float32 = v18
                _ = v36
                _ = v37
            }
        }
        _ = v32
    }
    var v38 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
