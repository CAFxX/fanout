package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int8; F2 int }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 int64, p2 uint8) int8 {
    var v4 uint8 = (p2 & (p2 + p2))
    if true {
        if (42 != 255) {
            var v5 int16 = 1
            var v6 string = ""
            _ = v5
            _ = v6
        } else {
            fmt.Println(v4, p0, p0)
            var v7 string = "\n\t"
            _ = v7
        }
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}

func main() {
    var v1 float64 = ((1.5 * (((3.14159 / 1.0) + (1.5 - 1.5)) - (1e10 - 1e10))) * 1.5)
    var v2 int16 = -32768
    var v3 int16 = v2
    var v8 float64 = (0.1 + float64(1e10))
    v9 := GenS1{F0: 1}
    for fi10 := 0; fi10 < 10; fi10++ {
        switch (fi10 % ((fi10) | 1)) {
        case -2147483648:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int16 = 32767
                v1 -= ((v1 / 1.0) / 0.1)
                v13 := []int{}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        if (v8 > (1e10 * (((0.0 + v1) - (float64(float32(-2.25)) - (1e10 + v1))) + 1e10))) {
            v3 -= -32768
            for fi14 := 0; fi14 < 8; fi14++ {
                v2 = (-(v3))
                var v15 bool = false
                fmt.Println(v8, fi14)
                _ = fi14
                _ = v15
            }
        }
        var v16 bool = (!true)
        _ = fi10
        _ = v16
    }
    if true {
        v17 := GenS2{F0: 115, F1: 1, F2: -9223372036854775808}
        v8 = float64((((-2.25 / 0.1) * 1.5) + -2.25))
        _ = v17
    }
    var v18 int8 = -49
    v19 := GenS1{F0: 65535}
    for fi20 := 0; fi20 < 2; fi20++ {
        v8 += 0.1
        if (!false) {
            v1 -= (((v1 * ((((((1e10 - v1) - (((v1 / (v8 + 0.1)) * v1) + v1)) * (((v8 / v8) / float64(0.1)) + v1)) / v8) - v8) * v8)) - 1e10) / (((((((v8 / 1.5) * v8) + 0.0) - (v8 / v8)) * v1) - (v8 + (v1 * v8))) * float64(0.1)))
            var v21 uint = 853
            _ = v21
        } else {
            if false {
                v18 += genFn0(1, 0, 33)
                v22 := []int{}
                _ = v22
            }
            var v23 int8 = -128
            var v24 int16 = v3
            _ = v23
            _ = v24
        }
        if false {
            var v25 uint16 = 2
            v8 = 1e10
            v8 = v1
            _ = v25
        } else {
            var v26 int64 = 42
            _ = v26
        }
        _ = fi20
    }
    if (v18 >= v18) {
        v3 += 244
        if (32768 == 32767) {
            var v27 uint16 = 655
            _ = v27
        } else {
            v28 := GenS3{F0: ""}
            if false {
                var v29 uint = 0
                _ = v29
            }
            _ = v28
        }
    }
    var v30 int32 = 42
    v18 -= v18
    switch 1 {
    case 32767:
        var v31 bool = (790 != 4294967295)
        _ = v31
    case -2147483648:
        v32 := []int{-214, 32768, -9223372036854775808, -2147483648}
        _ = v32
    }
    var v33 bool = false
    v18 += v18
    v34 := []int{128, 0}
    for fi35 := 0; fi35 < 6; fi35++ {
        for fi36 := 0; fi36 < 15; fi36++ {
            v8 = float64(1e10)
            var v37 uint64 = 1
            _ = fi36
            _ = v37
        }
        if v33 {
            var v38 uint = 65535
            var v39 string = "\n\t"
            v34 = append(v34, fi35)
            _ = v38
            _ = v39
        } else {
            v2 = -948
            var v40 int16 = v2
            _ = v40
        }
        v8 = v8
        _ = fi35
    }
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v8
    _ = v9
}
