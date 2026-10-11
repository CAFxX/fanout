package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int64) int8 {
    p0 -= (p0 ^ p0)
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 string = "\n\t"
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: ((fi4 * fi4) != fi4), F1: 255, F2: (^(fi4))}
        _ = fi4
        _ = v7
    }
    var v8 uint64 = 42
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    var v1 float64 = 0.0
    var v2 int16 = 1
    var v3 uint64 = 637
    v9 := GenS2{F0: 42, F1: 255}
    if (18 > 103) {
        v10 := []int{int(v2), -482, 2147483647, -2147483648}
        if false {
            for fi11 := 0; fi11 < 7; fi11++ {
                v3 = 0
                _ = fi11
            }
        }
        v3 += 18446744073709551615
        _ = v10
    } else {
        v12 := []int{2147483647, -1, -1}
        if len(v12) > 0 {
            v12[0] = 596
        }
        _ = v12
    }
    if (42 < 652) {
        v13 := []int{128, 2147483647, 128, 918}
        _ = v13
    }
    fmt.Println(v3, v3, v2)
    if true {
        if (true || (!false)) {
            var v14 uint16 = 255
            _ = v14
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            v1 -= v1
            var v16 int = fi15
            _ = fi15
            _ = v16
        }
        fmt.Println(v1, v2)
    } else {
        v1 += v1
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(v1, fi17)
            v1 = v1
            var v18 int64 = -92
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            fmt.Println(v2, v1)
            if true {
                switch ("hello" + ("" + "a")) {
                case "a":
                    v21 := GenS2{F0: 0, F1: 42}
                    _ = v21
                case "foo bar":
                    v3 += 4294967295
                }
                v22 := []int{fi20}
                if false {
                    var v23 uint16 = 42
                    _ = v23
                } else {
                    if len(v22) > 0 {
                        v22[0] = fi20
                    }
                    v1 += ((v1 * 0.1) + v1)
                    v1 += (v1 * 1.5)
                    var v24 uint16 = 65535
                    _ = v24
                }
                _ = v22
            }
            v3 += 18446744073709551615
            _ = fi20
        }
        v2 -= v2
        v25 := []int{fi19}
        _ = fi19
        _ = v25
    }
    v26 := GenS2{F0: 849, F1: 2}
    var v27 uint64 = v3
    v2 += -1
    fmt.Println(v3, v2, v3)
    v28 := GenS3{F0: 2}
    if true {
        var v29 uint32 = 836
        for fi30 := 0; fi30 < 16; fi30++ {
            if (42 != -750) {
                v2 = (v2 * (v2 / ((v2) | 1)))
                v29 = (v29 + (^(v29)))
            } else {
                var v31 string = "a"
                _ = v31
            }
            v32 := GenS3{F0: 47}
            v33 := []int{2147483647, fi30}
            _ = fi30
            _ = v32
            _ = v33
        }
        _ = v29
    }
    v34 := GenS1{F0: true, F1: 42, F2: -593}
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v9
}
