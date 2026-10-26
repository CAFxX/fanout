package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 int16, p1 uint16, p2 int) int {
    v5 := []int{-9223372036854775808}
    var v6 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 127
}
func genFn1() uint32 {
    switch ("a" + "") {
    case "hello":
        v7 := []int{-2147483648, -2147483648, 873, 0}
        _ = v7
    case "":
        var v8 int16 = -32768
        _ = v8
    case "\n\t":
        var v9 uint64 = 799
        _ = v9
    }
    var v10 uint32 = 2147483648
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: (v10 * (^(v10)))}
        v10 = (v10 & (v10 % ((v10) | 1)))
        _ = fi11
        _ = v12
    }
    _ = v10
    return 2147483648
}

func main() {
    var v1 int32 = 45
    var v2 int32 = (v1 & 1)
    v3 := make(map[int]string)
    var v4 uint64 = 42
    v13 := v3[-2147483648]
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 bool = (false || (!false))
            switch v13 {
            case "\n\t":
                v17 := GenS2{F0: 155, F1: 0, F2: 127}
                _ = v17
            }
            v1 += (v1 << ((255) & 63))
            if false {
                if (-2.25 != -2.25) {
                    var v18 uint64 = (v4 ^ v4)
                    var v19 int32 = (v1 ^ (v1 >> ((255) & 63)))
                    switch v13 {
                    case "a":
                        var v20 bool = v16
                        _ = v20
                    case "hello":
                        fmt.Println(v13)
                    }
                    _ = v18
                    _ = v19
                }
                var v21 uint8 = 2
                _ = v21
            } else {
                v22 := GenS3{F0: 896, F1: 1}
                _ = len(v3)
                _ = v22
            }
            var v23 float32 = 3.14159
            _ = fi15
            _ = v16
            _ = v23
        }
        var v24 int16 = -32768
        _ = fi14
        _ = v24
    }
    var v25 string = ""
    v2 += 2147483647
    v26 := []int{0, 42, 685, 798}
    var v27 int16 = 1
    v28 := GenS2{F0: 107, F1: 2, F2: 255}
    var v29 int16 = v27
    if true {
        _ = len(v3)
    }
    v30 := GenS4{F0: -2147483648}
    v4 = (v4 | (^(v4)))
    for fi31 := 0; fi31 < 8; fi31++ {
        fmt.Println(v13, v1)
        if (32768 < fi31) {
            if (true && false) {
                v26 = append(v26, 9223372036854775807)
            } else {
                if (0 < (v27 << ((881) & 63))) {
                    switch v25 {
                    case "hello":
                        var v32 float64 = 1.5
                        _ = v32
                    }
                    v3[32768] = v25
                }
                v25 = v25
            }
            var v33 float64 = 1e10
            var v34 bool = (255 > 128)
            _ = v33
            _ = v34
        } else {
            v35 := v3[(fi31 + (fi31 >> ((0) & 63)))]
            if (127 != -128) {
                var v36 uint8 = 255
                v2 += v2
                _ = v36
            }
            v13 = v13
            _ = v35
        }
        _ = fi31
    }
    v13 = v25
    fmt.Println(v27, v27, v2)
    var v37 uint8 = 2
    var v38 bool = false
    var v39 float32 = ((0.1 * (1e10 - float32(0.0))) / (0.1 + 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
