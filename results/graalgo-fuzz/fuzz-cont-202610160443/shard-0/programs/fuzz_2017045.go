package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 int, p1 int, p2 int64) int16 {
    v4 := GenS1{F0: 595, F1: p0, F2: 0}
    fmt.Println(p2, p0)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{42, fi5, (fi5 / ((fi5) | 1))}
        v7 := GenS1{F0: 2147483648, F1: p1, F2: 18446744073709551615}
        var v8 int32 = 1
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{128, (-(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return -503
}
func genFn1() int32 {
    var v10 bool = false
    for fi11 := 0; fi11 < 7; fi11++ {
        if v10 {
            if true {
                fmt.Println(fi11, v10)
            }
        } else {
            v12 := GenS2{F0: 127, F1: 4}
            _ = v12
        }
        fmt.Println(v10)
        _ = fi11
    }
    _ = v10
    return 2147483647
}

func main() {
    var v1 uint = 290
    v2 := make(map[uint32]int)
    var v3 bool = ((!true) && (4294967295 == 409))
    v3 = true
    var v13 string = ""
    v13 = v13
    if (-2147483648 >= -806) {
        fmt.Println(v1, v13, v3)
        fmt.Println(v1)
        fmt.Println(v3, v1, v1)
    }
    v3 = (!(v13 <= ("hello" + v13)))
    delete(v2, 42)
    delete(v2, 793)
    fmt.Println(v1, v1)
    for fi14 := 0; fi14 < 4; fi14++ {
        if v3 {
            var v15 float32 = 1e10
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 bool = (!false)
                var v18 int8 = -64
                var v19 uint8 = 2
                var v20 float32 = 1.5
                switch -614 {
                case 127:
                    v21 := GenS2{F0: 42, F1: v19}
                    _ = v21
                }
                v22 := GenS1{F0: 4294967295, F1: fi14, F2: 967}
                if true {
                    v23 := GenS1{F0: 967, F1: fi16, F2: 848}
                    _ = v23
                } else {
                    fmt.Println(v13, v1, v1)
                    v18 += v18
                    switch fi16 {
                    case -9223372036854775808:
                        if ((v13 + v13) != v13) {
                            v24 := []int{}
                            v1 = v1
                            _ = v24
                        }
                    case 1:
                        v1 -= 482
                    default:
                        fmt.Println(v13, v13, v3)
                    }
                }
                v25 := GenS3{F0: (v19 & (v19 >> ((v1) & 63))), F1: fi16}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
                _ = v20
                _ = v22
                _ = v25
            }
            v1 += (^(v1))
            _ = v15
        } else {
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 float64 = float64(3.14159)
                var v28 float32 = (0.1 * 0.0)
                _ = fi26
                _ = v27
                _ = v28
            }
            v13 = v13
        }
        _ = fi14
    }
    v29 := GenS1{F0: 0, F1: 0, F2: 759}
    var v30 uint32 = 65535
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 4; fi32++ {
            switch -2147483648 {
            case 42:
                fmt.Println(fi31, v1)
            case 4294967295:
                var v33 uint32 = (v30 / ((2147483647) | 1))
                _ = v33
            }
            _ = fi32
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
