package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    v3 := make(map[string]string)
    var v4 float64 = -2.25
    v5 := make(map[string]bool)
    v4 += (v2 + v2)
    _ = len(v5)
    fmt.Println(v4, v2)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float64 = (v2 - v2)
        v8 := GenS1{F0: -1}
        v9 := GenS2{F0: 0, F1: "\n\t"}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v2, v4)
    fmt.Println(v4)
    v4 = v2
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{(^(fi10)), (fi10 - (fi10 << ((65535) & 63))), fi10}
        _ = fi10
        _ = v11
    }
    switch -187 {
    case 128:
        if true {
            v2 += (v2 + v4)
            v4 += 0.1
            v12 := GenS1{F0: 100}
            _ = v12
        } else {
            _ = len(v1)
            for fi13 := 0; fi13 < 17; fi13++ {
                if true {
                    if (!true) {
                        for fi14 := 0; fi14 < 6; fi14++ {
                            delete(v1, fi14)
                            delete(v1, (fi13 * -155))
                            _ = fi14
                        }
                        var v15 int16 = 64
                        v16 := GenS2{F0: 65535, F1: ""}
                        _ = v15
                        _ = v16
                    } else {
                        v17 := []int{-1, fi13, fi13, 32767}
                        v4 = float64((float32(v4) - ((float32((v4 / v2)) / (3.14159 / 1e10)) / (-2.25 / -2.25))))
                        _ = v17
                    }
                }
                v2 -= 0.0
                v2 += (v4 / v4)
                _ = fi13
            }
        }
    }
    v18 := GenS1{F0: 127}
    v1[0] = true
    if true {
        v2 -= v4
    }
    var v19 float64 = 3.14159
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            if (141 <= 0) {
                var v22 bool = (!((true && false) && (!(101 <= 42))))
                v23 := GenS2{F0: 42, F1: ""}
                _ = v22
                _ = v23
            }
            if false {
                var v24 bool = true
                var v25 uint = 2
                if false {
                    v2 += float64((0.0 + (0.0 * 1.5)))
                    v26 := []int{fi20}
                    v1[(fi21 & fi21)] = (-1 != -9223372036854775808)
                    var v27 int8 = 100
                    fmt.Println(v27, v19, v19)
                    _ = v26
                    _ = v27
                }
                _ = v24
                _ = v25
            } else {
                v28 := GenS1{F0: 42}
                if false {
                    fmt.Println(v4, v19, fi20)
                    var v29 float64 = v4
                    _ = v29
                }
                var v30 uint16 = 42
                _ = v28
                _ = v30
            }
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
