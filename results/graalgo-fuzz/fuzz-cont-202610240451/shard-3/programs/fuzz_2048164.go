package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool }
func genFn0() bool {
    var v4 uint8 = 25
    v5 := GenS1{F0: 0}
    _ = v4
    _ = v5
    return (0 <= -417)
}
func genFn1(p0 int8, p1 bool) uint8 {
    if (!(!(false && (-60 < -128)))) {
        var v6 int64 = 448
        var v7 uint = 1
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: -880}
        _ = v8
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 string = "hello"
    var v2 int = 32767
    v3 := make(map[uint32]int)
    v2 = v2
    var v9 int64 = 1
    delete(v3, 2147483647)
    if false {
        fmt.Println(v1)
        var v10 float64 = 1e10
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := []int{9223372036854775807, fi11, v2, v2}
        if false {
            var v13 float64 = 1.5
            _ = v13
        } else {
            var v14 int32 = -2147483648
            _ = len(v3)
            if false {
                if (780 <= 2) {
                    v14 += v14
                    var v15 int = (fi11 >> ((255) & 63))
                    _ = v15
                } else {
                    v2 = fi11
                    v3[2147483648] = -1
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v2 = fi16
                        v14 -= (v14 + v14)
                        fmt.Println(v14, v14, fi16)
                        _ = fi16
                    }
                }
                var v17 uint16 = 1
                var v18 bool = ((v1 < (v1 + v1)) || false)
                _ = v17
                _ = v18
            }
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v19 uint32 = 318
    switch v2 {
    case 32767:
        switch (v2 >> ((1) & 63)) {
        case -424:
            var v20 int32 = -334
            _ = v20
        case -2147483648:
            var v21 int32 = -184
            _ = v21
        }
    }
    v2 = v2
    if true {
        var v22 int = v2
        switch (v1 + v1) {
        case "hello":
            if (464 >= 146) {
                var v23 string = v1
                v2 += (^(v22))
                v24 := GenS2{F0: 0, F1: ((false && (true && true)) || true)}
                _ = v23
                _ = v24
            }
        case "":
            v9 -= (v9 | v9)
        default:
            switch (v2 - 1) {
            case -9223372036854775808:
                var v25 uint8 = 127
                _ = v25
            default:
                v26 := v3[v19]
                _ = v26
            }
        }
        _ = v22
    } else {
        switch v2 {
        case -228:
            v27 := GenS2{F0: 818, F1: true}
            _ = v27
        }
        v28 := GenS2{F0: -32768, F1: false}
        _ = v28
    }
    var v29 int32 = 42
    delete(v3, v19)
    if ((v19 <= v19) && true) {
        var v30 int32 = (-(v29))
        _ = v30
    }
    var v31 float32 = 3.14159
    v31 += 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v9
}
