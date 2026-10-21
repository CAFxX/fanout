package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 string, p1 int8, p2 int16) uint64 {
    var v4 int16 = int16(p1)
    fmt.Println(v4, p2, v4)
    var v5 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1() int32 {
    v6 := GenS1{F0: 0}
    switch 647 {
    case -466:
        for fi7 := 0; fi7 < 14; fi7++ {
            if false {
                var v8 uint8 = 88
                _ = v8
            } else {
                var v9 int16 = -32768
                _ = v9
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                if true {
                    v11 := []int{32767, fi10, (^(fi7))}
                    if (false && (!(!true))) {
                        var v12 float64 = -2.25
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = v12
                    } else {
                        v11 = append(v11, 753)
                        var v13 int32 = -1
                        if true {
                            var v14 int32 = (v13 % ((v13) | 1))
                            _ = v14
                        }
                        _ = v13
                    }
                    fmt.Println(fi10, fi7)
                    _ = v11
                }
                _ = fi10
            }
            _ = fi7
        }
    case 1:
        if true {
            v15 := []int{-2147483648, -9223372036854775808, 42, -9223372036854775808}
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            var v16 int = 219
            _ = v15
            _ = v16
        }
    case 32768:
        var v17 uint16 = 42
        _ = v17
    default:
        var v18 int = -2147483648
        _ = v18
    }
    var v19 uint16 = 152
    _ = v19
    _ = v6
    return 915
}

func main() {
    var v1 int8 = 42
    var v2 int32 = 1
    var v3 float64 = 0.0
    v3 -= ((1e10 + 1.5) + ((v3 * v3) + (0.0 - (((v3 * 0.1) * v3) / (float64(0.0) * v3)))))
    var v20 float32 = (0.0 * (0.0 * 1.5))
    if (42 > -1) {
        for fi21 := 0; fi21 < 19; fi21++ {
            if (!((!(!false)) || (true && (42 < 4294967295)))) {
                v22 := GenS1{F0: 2147483647}
                for fi23 := 0; fi23 < 6; fi23++ {
                    v20 -= v20
                    v24 := GenS1{F0: 65535}
                    fmt.Println(fi21, v3)
                    _ = fi23
                    _ = v24
                }
                _ = v22
            } else {
                v25 := []int{(-(fi21)), -9223372036854775808}
                v26 := GenS1{F0: 826}
                _ = v25
                _ = v26
            }
            fmt.Println(v1, v3, v3)
            _ = fi21
        }
        v27 := GenS1{F0: 2147483648}
        var v28 int16 = 42
        _ = v27
        _ = v28
    } else {
        var v29 int8 = -110
        v2 += genFn1()
        fmt.Println(v3)
        _ = v29
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        fmt.Println(v2, v20, v3)
        var v31 float64 = (v3 * v3)
        v31 += 1.5
        _ = fi30
        _ = v31
    }
    v1 = v1
    v2 += v2
    v3 = ((-2.25 - float64((v20 / v20))) / (v3 / float64(3.14159)))
    var v32 int = -1
    var v33 uint16 = 72
    if false {
        var v34 int32 = (v2 * genFn1())
        _ = v34
    }
    v35 := GenS1{F0: 287}
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 uint16 = 32767
        var v38 uint32 = 0
        var v39 int = -9223372036854775808
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v35
}
