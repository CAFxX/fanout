package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint16 = 686
    var v2 string = "hello"
    var v3 bool = (-32768 <= 1)
    var v4 uint = 4294967295
    var v5 float32 = 1e10
    if true {
        var v6 uint = (^(v4))
        v7 := GenS1{F0: int32(v1)}
        var v8 int32 = 2147483647
        _ = v6
        _ = v7
        _ = v8
    }
    if (v3 || false) {
        v9 := GenS1{F0: 42}
        v4 = (v4 % ((16) | 1))
        _ = v9
    }
    var v10 uint8 = 1
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{fi11, fi11}
        fmt.Println(v2, fi11, v10)
        _ = fi11
        _ = v12
    }
    var v13 uint32 = 917
    if v3 {
        var v14 uint32 = (v13 * (^(v13)))
        _ = v14
    } else {
        v15 := []int{1, -98, 117}
        fmt.Println(v2, v3)
        _ = v15
    }
    var v16 float64 = 0.0
    var v17 float32 = v5
    if (32768 >= -942) {
        v4 += (v4 / ((v4) | 1))
        switch v2 {
        case "a":
            for fi18 := 0; fi18 < 8; fi18++ {
                switch v2 {
                case "":
                    for fi19 := 0; fi19 < 20; fi19++ {
                        if false {
                            if v3 {
                                v2 = v2
                                v20 := []int{(-(fi19)), (^(fi18)), (fi18 + fi18)}
                                var v21 uint16 = v1
                                _ = v20
                                _ = v21
                            }
                            var v22 uint8 = v10
                            _ = v22
                        } else {
                            fmt.Println(v5, v3, v2)
                        }
                        fmt.Println(v17)
                        v10 += v10
                        _ = fi19
                    }
                case "foo bar":
                    v23 := []int{(fi18 & fi18)}
                    _ = v23
                default:
                    var v24 int64 = -9223372036854775808
                    _ = v24
                }
                _ = fi18
            }
        case "":
            if (!v3) {
                for fi25 := 0; fi25 < 3; fi25++ {
                    var v26 int64 = 42
                    _ = fi25
                    _ = v26
                }
                var v27 int16 = 32767
                _ = v27
            } else {
                v16 = (v16 + v16)
            }
        }
        if (v3 && v3) {
            v5 = (float32(v16) - ((float32(((float64(v5) / v16) / float64(3.14159))) / 3.14159) / 1.5))
        } else {
            v2 = v2
        }
    }
    v5 = float32((v16 + v16))
    v3 = v3
    v17 += v5
    v28 := GenS2{F0: v1}
    var v29 int64 = 1
    var v30 float32 = -2.25
    v31 := GenS1{F0: -1}
    fmt.Println(v17, v29, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
