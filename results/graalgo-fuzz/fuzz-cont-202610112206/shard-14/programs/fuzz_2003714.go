package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 bool }
func genFn0(p0 int) uint64 {
    var v4 int32 = -2147483648
    if (94 >= 100) {
        v4 -= 627
    } else {
        var v5 uint8 = 255
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := []int{1, (^(fi6)), 127, 780}
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int8 = -128
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        if true {
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 float64 = 0.1
                if true {
                    var v12 uint16 = 2
                    var v13 int32 = v4
                    v14 := []int{(fi10 - (p0 / (((fi10 - 127)) | 1))), 2147483647}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                if false {
                    var v15 int64 = -9223372036854775808
                    for fi16 := 0; fi16 < 10; fi16++ {
                        fmt.Println(fi10, v5)
                        _ = fi16
                    }
                    fmt.Println(p0, v15)
                    _ = v15
                } else {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v18 := []int{(fi17 / ((p0) | 1)), (fi17 * 32767), (p0 | fi17)}
                        if len(v18) > 0 {
                            v18[0] = (fi10 + 2147483647)
                        }
                        _ = fi17
                        _ = v18
                    }
                }
                _ = fi10
                _ = v11
            }
        }
        _ = v5
    }
    v4 = (-(v4))
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := []int{fi19}
        _ = fi19
        _ = v20
    }
    p0 += 9223372036854775807
    _ = p0
    _ = v4
    return 918
}

func main() {
    var v1 int = 32767
    v2 := make(map[string]uint32)
    v3 := []int{513, (^(v1))}
    _ = len(v2)
    fmt.Println(v1)
    v3 = append(v3, (v1 & 802))
    var v21 int16 = -362
    v1 += 131
    var v22 float64 = 3.14159
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v23 := GenS1{F0: 287, F1: 65535, F2: v1}
    v22 += ((float64(-2.25) - v22) * 0.0)
    switch v1 {
    case -1:
        switch v1 {
        case -9223372036854775808:
            var v24 int16 = 450
            _ = v24
        case 127:
            v25 := v2["foo bar"]
            _ = v25
        }
    case 42:
        v26 := GenS1{F0: 0, F1: 852, F2: (v1 + 128)}
        _ = v26
    default:
        v21 = v21
    }
    var v27 int8 = -128
    if (!((1000 >= 1) || false)) {
        switch ((("hello" + "\n\t") + "") + "") {
        case "a":
            for fi28 := 0; fi28 < 10; fi28++ {
                fmt.Println(v1)
                if (!false) {
                    v1 = (-(v1))
                }
                v22 -= ((v22 - (v22 + float64(0.0))) / float64(float32(((float64(-2.25) * (v22 * float64(1.5))) / v22))))
                _ = fi28
            }
        default:
            var v29 float32 = -2.25
            _ = v29
        }
    } else {
        if (65535 >= 255) {
            v21 = (v21 + 136)
            v30 := GenS2{F0: 0, F1: 0, F2: (2 <= 2)}
            _ = v30
        }
        fmt.Println(v27, v21)
        v21 -= v21
    }
    var v31 int64 = 9223372036854775807
    var v32 uint32 = 728
    var v33 int8 = v27
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 int64 = v31
            delete(v2, "")
            _ = fi34
            _ = v35
        }
    case "\n\t":
        for fi36 := 0; fi36 < 8; fi36++ {
            for fi37 := 0; fi37 < 15; fi37++ {
                var v38 int16 = v21
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
