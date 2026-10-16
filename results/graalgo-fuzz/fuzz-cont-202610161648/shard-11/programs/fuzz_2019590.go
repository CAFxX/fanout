package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 bool; F1 bool; F2 uint32 }
func genFn0(p0 uint16, p1 uint64, p2 uint16) int {
    if (!(1e10 < -2.25)) {
        v7 := GenS1{F0: 1, F1: true}
        _ = v7
    }
    if false {
        p1 = (p1 % (((p1 ^ p1)) | 1))
        if true {
            p1 = (p1 | 715)
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := []int{fi8, (fi8 ^ (fi8 & (fi8 / ((fi8) | 1)))), -9223372036854775808}
                var v10 float32 = 0.0
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS1{F0: (^(p1)), F1: (1 < 2147483647)}
                p1 -= p1
                v13 := GenS2{F0: ((p2 % ((p0) | 1)) < 42), F1: (1 < 6), F2: 2147483647}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    if (853 >= 1) {
        var v14 uint64 = (^(p1))
        _ = v14
    } else {
        p0 = 60
        p0 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]bool)
    var v3 int8 = 127
    var v4 uint64 = 1
    var v5 bool = (1 == 1)
    var v6 uint64 = v4
    v15 := GenS1{F0: v6, F1: (!v5)}
    fmt.Println(v3, v6, v4)
    v1 = 2147483647
    var v16 uint32 = 42
    var v17 uint16 = 522
    switch 9223372036854775807 {
    case 32768:
        v1 = -2147483648
    }
    v18 := GenS2{F0: false, F1: v5, F2: (v16 ^ (^(v16)))}
    for fi19 := 0; fi19 < 19; fi19++ {
        v4 = v6
        v20 := []int{genFn0((^(v17)), (^(v6)), v17), int(v1)}
        var v21 string = "foo bar"
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int32 = -1
    var v23 uint64 = (^(v4))
    v2[1] = v5
    var v24 int8 = 42
    for fi25 := 0; fi25 < 3; fi25++ {
        if v5 {
            var v26 uint32 = (v16 & v16)
            v1 -= (v22 / ((v22) | 1))
            _ = v26
        }
        if true {
            for fi27 := 0; fi27 < 1; fi27++ {
                v16 -= v16
                if ((-(v24)) < v24) {
                    _ = len(v2)
                    var v28 string = "hello"
                    _ = v28
                } else {
                    switch genFn0(v17, v23, (v17 % ((v17) | 1))) {
                    case -233:
                        v6 += v6
                    }
                    v29 := GenS1{F0: uint64(v17), F1: v5}
                    _ = v29
                }
                _ = fi27
            }
            if (!v5) {
                fmt.Println(v22, v1, v17)
                v3 -= -1
                var v30 float64 = (float64(0.1) - 1.5)
                _ = v30
            } else {
                v4 = v23
            }
        }
        _ = fi25
    }
    var v31 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
