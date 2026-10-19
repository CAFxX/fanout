package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() uint64 {
    if (!(false && (false || ((476 > 2) || true)))) {
        var v4 float64 = 1e10
        _ = v4
    }
    var v5 int16 = -817
    v6 := []int{-9223372036854775808}
    fmt.Println(v5)
    v5 += 323
    _ = v5
    _ = v6
    return 834
}
func genFn1(p0 int16) int64 {
    v7 := []int{32767, 639}
    if false {
        p0 -= (p0 + (p0 << ((269) & 63)))
        _ = len(v7)
    }
    var v8 float32 = 3.14159
    if false {
        v8 -= 0.0
        v9 := GenS1{F0: -2147483648}
        v8 = float32(((1.5 / 1e10) / (0.1 / 3.14159)))
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]string)
    var v3 float64 = 1.5
    for fi10 := 0; fi10 < 12; fi10++ {
        _ = len(v2)
        var v11 int16 = 42
        var v12 float32 = 1.5
        _ = fi10
        _ = v11
        _ = v12
    }
    if false {
        if (381 <= 0) {
            v1 += v1
        }
        var v13 int8 = 42
        if false {
            fmt.Println(v3)
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint16 = 1
                _ = fi14
                _ = v15
            }
            fmt.Println(v1, v1)
        }
        _ = v13
    }
    v3 = (0.0 * v3)
    delete(v2, 1)
    _ = len(v2)
    delete(v2, -9223372036854775808)
    v1 = -32768
    v16 := []int{}
    fmt.Println(v1, v1)
    v17 := v2[32767]
    if true {
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 3; fi19++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi19
            }
            var v20 string = (v17 + "\n\t")
            if true {
                fmt.Println(v3, v17, v17)
            } else {
                fmt.Println(v17)
                var v21 uint16 = 32767
                v22 := GenS1{F0: -1}
                _ = v21
                _ = v22
            }
            _ = fi18
            _ = v20
        }
        v23 := GenS1{F0: 2147483647}
        for fi24 := 0; fi24 < 14; fi24++ {
            v3 += float64(1.5)
            var v25 uint32 = 0
            v2[(fi24 | (fi24 ^ fi24))] = (v17 + v17)
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v17 = (v17 + (v17 + "hello"))
    if (false && false) {
        for fi26 := 0; fi26 < 10; fi26++ {
            v27 := GenS1{F0: -147}
            _ = fi26
            _ = v27
        }
        var v28 uint16 = 255
        fmt.Println(v3, v28, v3)
        _ = v28
    }
    var v29 float64 = v3
    var v30 uint64 = genFn0()
    var v31 float64 = 0.0
    fmt.Println(v30, v29)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
