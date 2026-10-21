package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 string, p1 int16) int32 {
    fmt.Println(p1, p0)
    if false {
        var v4 uint32 = 42
        _ = v4
    } else {
        var v5 uint = 42
        _ = v5
    }
    _ = p0
    _ = p1
    return -2147483648
}
func genFn1(p0 int16) uint {
    for fi6 := 0; fi6 < 18; fi6++ {
        p0 -= 1000
        _ = fi6
    }
    v7 := []int{0, 127}
    _ = p0
    _ = v7
    return 211
}

func main() {
    v1 := []int{}
    var v2 int16 = 32767
    var v3 int8 = -128
    v8 := GenS1{F0: 199}
    v2 -= v2
    var v9 float32 = float32((1e10 + float64((0.1 / ((0.1 * 3.14159) + 0.0)))))
    if (((1 > 0) || (!(0 >= 2))) || false) {
        if (87 == 1) {
            _ = len(v1)
            if len(v1) > 0 {
                v1[0] = 1
            }
        }
    } else {
        fmt.Println(v9, v2)
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 float64 = float64(3.14159)
            _ = fi10
            _ = v11
        }
    }
    var v12 int32 = -589
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 float32 = (3.14159 * -2.25)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int32 = v12
        if true {
            if (255 > 503) {
                _ = len(v1)
                var v17 bool = (!false)
                _ = v17
            } else {
                var v18 uint = genFn1(1)
                if len(v1) > 0 {
                    v1[0] = fi15
                }
                v19 := GenS1{F0: -125}
                _ = v18
                _ = v19
            }
        } else {
            v1 = append(v1, fi15)
            if (true || true) {
                if true {
                    v1 = append(v1, (-(fi15)))
                }
                v12 += (v12 + v12)
                v9 = (v9 + 1e10)
            } else {
                for fi20 := 0; fi20 < 9; fi20++ {
                    switch "foo bar" {
                    case "":
                        fmt.Println(v3)
                    }
                    var v21 float64 = 0.1
                    if true {
                        v21 = 3.14159
                    } else {
                        var v22 bool = true
                        var v23 uint16 = 255
                        v24 := GenS1{F0: (^(fi15))}
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                }
                var v25 uint16 = 42
                if (!(false || true)) {
                    var v26 int64 = 9223372036854775807
                    _ = v26
                }
                _ = v25
            }
        }
        fmt.Println(v9, v12, v12)
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
