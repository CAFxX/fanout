package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint64; F1 string }

func main() {
    v1 := []int{9223372036854775807}
    var v2 int = 9223372036854775807
    var v3 string = "a"
    var v4 bool = (42 < -2147483648)
    v3 = (v3 + v3)
    if v4 {
        var v5 uint8 = 114
        _ = v5
    }
    var v6 bool = v4
    if v6 {
        fmt.Println(v4, v2, v3)
        if (!(103 == -70)) {
            v7 := GenS1{F0: -1, F1: 1}
            _ = v7
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS2{F0: 65535, F1: 0}
            _ = fi8
            _ = v9
        }
    }
    var v10 float64 = (0.0 / 0.1)
    for fi11 := 0; fi11 < 2; fi11++ {
        v10 = v10
        _ = fi11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v3 = v3
        if (1 >= 107) {
            v10 = v10
            fmt.Println(fi12, v4, fi12)
            v13 := GenS3{F0: 0, F1: 65535, F2: v3}
            _ = v13
        }
        v4 = true
        _ = fi12
    }
    var v14 uint32 = 2147483647
    var v15 string = ("hello" + (v3 + v3))
    if (1e10 != 1e10) {
        var v16 float64 = 3.14159
        v10 = ((float64(((-2.25 - 0.1) + (1e10 / (1.5 * float32((v10 / (v16 * 0.0))))))) - 1e10) * v10)
        if (!(0 != 0)) {
            var v17 uint8 = 1
            v17 -= v17
            if false {
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if v4 {
                            v20 := GenS2{F0: 380, F1: v17}
                            _ = v20
                        } else {
                            var v21 string = v15
                            _ = v21
                        }
                        v14 -= 65535
                        _ = fi19
                    }
                    var v22 float64 = 0.1
                    v23 := GenS4{F0: (true || v6), F1: 42}
                    _ = fi18
                    _ = v22
                    _ = v23
                }
                if v6 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v24 int32 = -2147483648
                    var v25 uint64 = 70
                    _ = v24
                    _ = v25
                }
                var v26 int = (v2 & (v2 - (v2 * v2)))
                _ = v26
            }
            _ = v17
        }
        _ = v16
    }
    if v6 {
        for fi27 := 0; fi27 < 2; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                v1 = append(v1, fi27)
                v29 := GenS1{F0: 0, F1: -9223372036854775808}
                _ = fi28
                _ = v29
            }
            var v30 bool = (1 < fi27)
            _ = fi27
            _ = v30
        }
        var v31 int8 = -1
        if len(v1) > 0 {
            v1[0] = v2
        }
        _ = v31
    } else {
        v32 := GenS5{F0: 2, F1: (v3 + v3)}
        var v33 uint64 = 1
        _ = v32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
