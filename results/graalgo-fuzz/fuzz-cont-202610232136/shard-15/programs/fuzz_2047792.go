package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
func genFn0() uint8 {
    var v7 uint32 = 415
    fmt.Println(v7)
    var v8 int = -716
    var v9 int16 = 42
    var v10 int8 = 33
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1() int8 {
    v11 := []int{-1, 1, 189, 9223372036854775807}
    var v12 string = ""
    if (false && (v12 <= v12)) {
        v13 := GenS1{F0: 42, F1: false}
        fmt.Println(v12)
        fmt.Println(v12)
        _ = v13
    }
    _ = v11
    _ = v12
    return -25
}

func main() {
    var v1 uint8 = 1
    var v2 string = "a"
    var v3 int8 = 42
    var v4 uint = 65535
    v5 := make(map[uint32]string)
    v6 := make(map[string]bool)
    var v14 float64 = 0.0
    var v15 string = v2
    if (!false) {
        for fi16 := 0; fi16 < 14; fi16++ {
            v3 -= (v3 / ((v3) | 1))
            _ = fi16
        }
    }
    v17 := GenS1{F0: 0, F1: false}
    for fi18 := 0; fi18 < 4; fi18++ {
        _ = len(v6)
        for fi19 := 0; fi19 < 20; fi19++ {
            v2 = (v2 + v15)
            var v20 int = fi19
            var v21 bool = (v3 >= -1)
            _ = fi19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            _ = len(v5)
            var v23 int = fi18
            v24 := []int{-212, v23, 9223372036854775807, -2147483648}
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi18
    }
    if false {
        var v25 float64 = (0.0 / ((1e10 * (v14 * float64(1e10))) - float64(-2.25)))
        _ = v25
    }
    v4 = v4
    switch v15 {
    case "\n\t":
        v26 := []int{-542, 1, 4294967295}
        _ = v26
    case "":
        switch (v15 + v2) {
        case "\n\t":
            var v27 int32 = 2147483647
            _ = v27
        }
    }
    v3 += v3
    if (!false) {
        for fi28 := 0; fi28 < 12; fi28++ {
            v2 = (v15 + v15)
            if false {
                switch -50 {
                case 1:
                    var v29 uint32 = 2147483648
                    _ = v29
                case 4294967295:
                    v14 = v14
                }
                v30 := v5[65535]
                var v31 float64 = float64(0.1)
                _ = v30
                _ = v31
            } else {
                var v32 float64 = (v14 / v14)
                _ = v32
            }
            var v33 uint8 = v1
            _ = fi28
            _ = v33
        }
        v34 := GenS1{F0: 403, F1: (1e10 == v14)}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
