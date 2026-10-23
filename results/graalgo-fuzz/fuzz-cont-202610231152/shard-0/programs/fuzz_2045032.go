package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0() uint8 {
    var v5 float32 = 3.14159
    var v6 int16 = 624
    switch "foo bar" {
    case "\n\t":
        var v7 uint64 = 4294967295
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 1; fi8++ {
            if (983 != 2147483647) {
                if (0 < 906) {
                    v5 = (0.1 - (v5 - v5))
                    if false {
                        if (false && true) {
                            fmt.Println(v6)
                            v6 = v6
                            v9 := GenS1{F0: 1, F1: 712, F2: 2147483648}
                            _ = v9
                        } else {
                            v5 = v5
                            v5 = float32((0.1 + 1.5))
                        }
                        if true {
                            var v10 int64 = 0
                            _ = v10
                        } else {
                            v5 = (v5 / v5)
                            v11 := []int{fi8, -252, -643, fi8}
                            _ = v11
                        }
                    }
                }
                fmt.Println(v5, v5, v6)
            }
            _ = fi8
        }
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{fi12}
        fmt.Println(fi12, v6)
        if false {
            if false {
                fmt.Println(fi12)
            }
        } else {
            if true {
                var v14 uint = 2
                v14 = (v14 - v14)
                _ = v14
            } else {
                v5 = ((float32(1e10) / v5) - v5)
                fmt.Println(fi12)
            }
            v5 = v5
        }
        _ = fi12
        _ = v13
    }
    v15 := []int{}
    _ = v15
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 int8 = -128
    fmt.Println(v4)
    var v16 string = ""
    if (!(genFn0() > 42)) {
        v17 := GenS2{F0: 622, F1: 0}
        _ = v17
    } else {
        _ = len(v2)
        fmt.Println(v4)
    }
    if true {
        v4 = (-(v4))
    } else {
        var v18 int = 2147483647
        v19 := GenS3{F0: false, F1: 610}
        var v20 bool = true
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 uint32 = 2147483648
    var v22 int = 9223372036854775807
    var v23 string = ((v16 + "hello") + ((v16 + v16) + (("hello" + v16) + "")))
    var v24 float64 = float64(float32(((1e10 + 1.5) - 1.5)))
    var v25 uint16 = 917
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := GenS1{F0: -930, F1: v25, F2: v21}
        v28 := []int{(^(fi26)), fi26, (fi26 % ((v22) | 1))}
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := GenS4{F0: v25}
            v28 = append(v28, (fi26 / ((412) | 1)))
            fmt.Println(v22, fi26, v16)
            _ = fi29
            _ = v30
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v21, v25, v22)
    switch v16 {
    case "a":
        v4 += v4
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 int16 = 1000
        v21 = uint32(v25)
        v16 = "foo bar"
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
