package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
func genFn0() uint16 {
    var v7 int8 = 22
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint8 = 190
        v10 := []int{}
        switch fi8 {
        case 42:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        case 0:
            var v11 int8 = (-(v7))
            _ = v11
        case -965:
            fmt.Println(fi8, v9)
        default:
            if ((!(!false)) && true) {
                fmt.Println(v9)
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 128
    v3 := []uint32{2, 4294967295, 4294967295, 113}
    var v4 int = 42
    var v5 float64 = ((1.5 + 0.0) * (1e10 - (((0.0 * -2.25) - -2.25) * float64(3.14159))))
    var v6 bool = false
    switch v4 {
    case 430:
        v6 = v6
    }
    var v12 int = 32767
    var v13 float32 = float32(v5)
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v13, v13)
        var v15 float64 = 1.5
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: -2147483648, F1: v6}
    var v17 string = "foo bar"
    var v18 uint8 = v2
    for fi19 := 0; fi19 < 17; fi19++ {
        _ = len(v1)
        var v20 int64 = -642
        var v21 float64 = ((3.14159 * -2.25) / ((float64((v13 / v13)) * v5) - v5))
        _ = fi19
        _ = v20
        _ = v21
    }
    v12 += (v4 - v4)
    for fi22 := 0; fi22 < 14; fi22++ {
        v17 = ("a" + (v17 + ("a" + (v17 + (v17 + v17)))))
        _ = fi22
    }
    switch (v12 >> ((18446744073709551615) & 63)) {
    case 32767:
        var v23 uint16 = 10
        _ = v23
    case -1:
        v24 := GenS2{F0: -855, F1: false}
        _ = v24
    case 638:
        v25 := GenS3{F0: -1}
        _ = v25
    default:
        for fi26 := 0; fi26 < 8; fi26++ {
            switch 0 {
            case 0:
                var v27 int = -2147483648
                _ = v27
            case 1:
                for fi28 := 0; fi28 < 20; fi28++ {
                    switch "hello" {
                    case "":
                        var v29 uint64 = 0
                        _ = v29
                    }
                    switch (v17 + v17) {
                    case "foo bar":
                        if (-9223372036854775808 != 1) {
                            var v30 uint16 = 465
                            _ = v30
                        }
                    default:
                        switch 32768 {
                        case -1:
                            var v31 uint8 = 128
                            _ = v31
                        case -164:
                            v17 = "hello"
                        default:
                            v32 := GenS4{F0: v12}
                            _ = v32
                        }
                    }
                    if v6 {
                        var v33 uint32 = 2
                        v6 = true
                        _ = v33
                    } else {
                        fmt.Println(v4)
                        if (-330 > -9223372036854775808) {
                            fmt.Println(fi28, v17, v17)
                            var v34 uint = 18446744073709551615
                            _ = v34
                        } else {
                            v35 := GenS1{F0: v4, F1: (1 >= 687)}
                            fmt.Println(v12)
                            fmt.Println(v5, fi28)
                            _ = v35
                        }
                    }
                    _ = fi28
                }
            case 215:
                v36 := GenS5{F0: 0}
                _ = v36
            }
            _ = fi26
        }
    }
    var v37 int8 = 0
    v38 := GenS5{F0: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
