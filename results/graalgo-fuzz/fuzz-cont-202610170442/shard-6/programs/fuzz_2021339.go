package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }

func main() {
    var v1 int64 = 0
    var v2 float32 = ((float32((1.5 - float64(1e10))) * (0.1 + 1.5)) + 3.14159)
    var v3 int64 = -9223372036854775808
    var v4 string = ("hello" + "")
    var v5 uint32 = 2147483648
    if true {
        v4 = v4
        v1 += v1
        fmt.Println(v3, v5, v3)
    } else {
        var v6 bool = true
        _ = v6
    }
    var v7 int16 = int16(v3)
    v8 := GenS1{F0: 65535}
    if false {
        v5 -= v5
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            if false {
                switch fi9 {
                case 128:
                    var v10 string = v4
                    _ = v10
                case -9223372036854775808:
                    v11 := []int{32768, fi9, (^(fi9)), fi9}
                    _ = v11
                }
                if (0 <= 4294967295) {
                    var v12 uint32 = v5
                    if false {
                        var v13 string = v4
                        _ = v13
                    }
                    _ = v12
                }
            } else {
                if ((!(!(true || false))) && ((v7 / ((v7) | 1)) <= 42)) {
                    v14 := GenS1{F0: 328}
                    switch fi9 {
                    case 32767:
                        var v15 uint8 = 42
                        _ = v15
                    case -60:
                        v16 := []int{fi9, (fi9 << ((18446744073709551615) & 63))}
                        _ = v16
                    case 42:
                        if true {
                            v5 = (^(v5))
                            var v17 float64 = float64(v2)
                            v18 := []int{(fi9 ^ (fi9 % ((fi9) | 1))), fi9, fi9, fi9}
                            _ = v17
                            _ = v18
                        } else {
                            v19 := []int{}
                            _ = v19
                        }
                    }
                    var v20 float64 = 0.0
                    _ = v14
                    _ = v20
                } else {
                    var v21 float64 = 1e10
                    _ = v21
                }
                var v22 uint = 2
                _ = v22
            }
            if ((!false) && (!((v1 ^ 217) < v3))) {
                v23 := []int{(fi9 % (((fi9 + fi9)) | 1))}
                _ = v23
            }
            _ = fi9
        }
    }
    var v24 uint8 = 39
    v25 := GenS1{F0: 42}
    v26 := GenS1{F0: 32}
    v27 := GenS2{F0: 15}
    if (0 < 802) {
        var v28 bool = false
        v7 -= (v7 << ((2) & 63))
        _ = v28
    }
    v29 := []int{-1, 42, -9223372036854775808, 49}
    v30 := GenS2{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
