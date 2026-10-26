package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0() int16 {
    switch 0 {
    case 9223372036854775807:
        var v7 float64 = float64((1.5 / 3.14159))
        _ = v7
    case 369:
        switch "foo bar" {
        case "\n\t":
            if false {
                var v8 int16 = 32767
                var v9 float64 = 0.1
                _ = v8
                _ = v9
            } else {
                if false {
                    v10 := GenS1{F0: -598, F1: 645, F2: -2147483648}
                    var v11 uint = 0
                    switch 299 {
                    case 205:
                        v12 := []int{-16, 0}
                        _ = v12
                    }
                    _ = v10
                    _ = v11
                }
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 int = (fi13 - fi13)
                    if ((false || (1 < 42)) || (!(65535 != 961))) {
                        var v15 uint32 = 42
                        _ = v15
                    } else {
                        var v16 uint16 = 1
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
            }
        case "":
            var v17 uint = 0
            _ = v17
        }
    case 42:
        var v18 int32 = -1
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(fi19)
        }
        _ = fi19
    }
    return -865
}

func main() {
    var v1 float64 = (1e10 + (3.14159 * -2.25))
    v2 := make(map[int]string)
    var v3 uint = 255
    var v4 int32 = 0
    v5 := make(map[string]int)
    v6 := make(map[uint32]int)
    var v20 bool = false
    v20 = (1 != -28)
    fmt.Println(v20)
    v21 := GenS1{F0: 371, F1: 65535, F2: v4}
    v1 += v1
    v20 = (820 != -32768)
    if v20 {
        v22 := GenS1{F0: -9223372036854775808, F1: 42, F2: v4}
        v5["foo bar"] = -346
        var v23 int32 = (-(v4))
        _ = v22
        _ = v23
    }
    if v20 {
        var v24 float32 = 0.0
        _ = v24
    }
    v25 := GenS2{F0: -128}
    var v26 bool = v20
    if ((935 >= 486) || v26) {
        v3 = v3
        if (v26 && v20) {
            if (42 != -85) {
                var v27 int = 42
                v28 := []int{-2147483648, v27}
                var v29 float64 = 3.14159
                _ = v27
                _ = v28
                _ = v29
            }
            if true {
                delete(v6, 2)
                var v30 int = 9223372036854775807
                for fi31 := 0; fi31 < 14; fi31++ {
                    v32 := GenS2{F0: 127}
                    _ = fi31
                    _ = v32
                }
                _ = v30
            } else {
                v33 := GenS1{F0: 53, F1: 65535, F2: v4}
                _ = v33
            }
            v34 := GenS3{F0: false, F1: 0, F2: 42}
            _ = v34
        }
    } else {
        v35 := GenS2{F0: -128}
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        for fi37 := 0; fi37 < 8; fi37++ {
            if false {
                v20 = false
                var v38 float32 = ((0.1 - (0.0 * -2.25)) / -2.25)
                var v39 int = fi36
                _ = v38
                _ = v39
            }
            _ = fi37
        }
        v40 := []int{816, fi36}
        _ = fi36
        _ = v40
    }
    var v41 int = 127
    delete(v5, "a")
    var v42 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
