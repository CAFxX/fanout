package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    var v1 uint32 = 453
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]bool)
    if false {
        v4 := GenS1{F0: int32(v1), F1: (v1 / ((v1) | 1)), F2: 0}
        _ = v4
    }
    if (2147483647 != -1) {
        for fi5 := 0; fi5 < 20; fi5++ {
            switch "" {
            case "foo bar":
                if true {
                    fmt.Println(fi5, v2, v2)
                }
            default:
                for fi6 := 0; fi6 < 14; fi6++ {
                    v2 = (^(v2))
                    _ = fi6
                }
            }
            v1 = (v1 ^ v1)
            v7 := GenS1{F0: 2147483647, F1: v1, F2: 42}
            _ = fi5
            _ = v7
        }
        switch "" {
        case "hello":
            for fi8 := 0; fi8 < 10; fi8++ {
                if ((fi8 != fi8) && (-2.25 == (1.5 + float64(1e10)))) {
                    var v9 uint8 = 2
                    v10 := []int{}
                    v1 -= v1
                    _ = v10
                    _ = v9
                }
                v11 := []int{}
                _ = fi8
                _ = v11
            }
        case "foo bar":
            var v12 string = "a"
            _ = v12
        }
        var v13 int8 = -10
        _ = v13
    } else {
        var v14 float64 = float64(1.5)
        v15 := []int{296, -2147483648, 1, -2147483648}
        v16 := GenS1{F0: 0, F1: v1, F2: 186}
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 int16 = 1
    var v18 uint8 = 153
    var v19 uint16 = 493
    if false {
        fmt.Println(v17)
        var v20 uint = 620
        _ = v20
    }
    var v21 uint8 = (^(v18))
    var v22 uint32 = 1
    v19 = 1
    if true {
        var v23 int8 = 29
        if true {
            v24 := GenS1{F0: 0, F1: 357, F2: v18}
            for fi25 := 0; fi25 < 3; fi25++ {
                _ = len(v3)
                v2 += 0
                v26 := []int{(fi25 ^ 42)}
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        _ = v23
    }
    v18 = 0
    v1 += v22
    for fi27 := 0; fi27 < 1; fi27++ {
        if true {
            var v28 uint16 = (v19 + v19)
            fmt.Println(v28, v18)
            _ = v28
        } else {
            if true {
                var v29 uint64 = 255
                _ = v29
            }
            v30 := v3[v22]
            if true {
                delete(v3, (^(v22)))
            }
            _ = v30
        }
        _ = fi27
    }
    if false {
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := GenS1{F0: -1, F1: (v22 & v22), F2: v21}
            _ = fi31
            _ = v32
        }
        var v33 uint64 = 0
        for fi34 := 0; fi34 < 17; fi34++ {
            v35 := GenS2{F0: (false && false), F1: 4294967295}
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    var v36 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v36
}
