package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
func genFn0(p0 uint16) bool {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 7; fi4++ {
        if true {
            var v5 bool = (0 < 42)
            v5 = (false || ((1 < 60) && (-128 > 42)))
            _ = v5
        }
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 bool = (true && false)
            if ((-2147483648 < 0) || ((false && v7) || (1 >= 128))) {
                var v8 string = ""
                v9 := GenS1{F0: 42, F1: 462, F2: (fi6 | (fi6 / ((fi4) | 1)))}
                v10 := GenS1{F0: 42, F1: 202, F2: fi6}
                _ = v10
                _ = v8
                _ = v9
            } else {
                var v11 float64 = 1.5
                v12 := []int{fi6, (fi4 >> ((42) & 63))}
                _ = v11
                _ = v12
            }
            if v7 {
                var v13 uint32 = 572
                _ = v13
            }
            _ = fi6
            _ = v7
        }
        _ = fi4
    }
    _ = p0
    return (true || true)
}
func genFn1() int {
    var v14 int16 = 916
    v15 := []int{9223372036854775807, -796}
    var v16 float64 = 1e10
    var v17 string = "\n\t"
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 128
}

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    var v18 int16 = (v2 - v2)
    var v19 int64 = -265
    switch "" {
    case "a":
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(v2)
            _ = fi20
        }
    case "":
        v21 := GenS1{F0: 4294967295, F1: 142, F2: genFn1()}
        _ = v21
    }
    fmt.Println(v2)
    if (v18 > 42) {
        if true {
            if true {
                var v22 int32 = v1
                switch 9223372036854775807 {
                case 32767:
                    var v23 int32 = v1
                    _ = v23
                }
                _ = v22
            } else {
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 bool = false
                    _ = fi24
                    _ = v25
                }
            }
            v3[1] = (803 >= genFn1())
            var v26 uint32 = 453
            _ = v26
        }
    }
    delete(v3, 0)
    var v27 float32 = float32((1e10 + (3.14159 + 0.1)))
    for fi28 := 0; fi28 < 2; fi28++ {
        switch (fi28 * fi28) {
        case 32768:
            v19 = (v19 ^ v19)
        case 242:
            for fi29 := 0; fi29 < 10; fi29++ {
                switch "hello" {
                case "":
                    var v30 uint16 = 672
                    _ = v30
                default:
                    var v31 uint32 = 0
                    _ = v31
                }
                fmt.Println(fi28, v19)
                var v32 uint32 = 42
                _ = fi29
                _ = v32
            }
        case 1:
            v1 -= v1
        }
        v3[2147483647] = (!(927 < 2))
        v33 := GenS1{F0: 2, F1: 4294967295, F2: (^(fi28))}
        _ = fi28
        _ = v33
    }
    v18 -= -550
    for fi34 := 0; fi34 < 10; fi34++ {
        if true {
            v2 -= (-(v18))
            fmt.Println(v27, fi34)
            v19 = v19
        }
        if false {
            var v35 int64 = 9223372036854775807
            v1 -= v1
            v3[42] = false
            _ = v35
        } else {
            v3[2147483648] = (118 > -43)
            var v36 uint32 = 0
            _ = v36
        }
        if (fi34 > (fi34 / ((genFn1()) | 1))) {
            for fi37 := 0; fi37 < 17; fi37++ {
                _ = len(v3)
                v38 := []int{fi37, genFn1()}
                _ = fi37
                _ = v38
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v19, v18)
            default:
                delete(v3, 2147483647)
            }
            v39 := v3[2147483647]
            _ = v39
        }
        _ = fi34
    }
    v19 = (v19 + v19)
    for fi40 := 0; fi40 < 10; fi40++ {
        var v41 float64 = 0.0
        _ = fi40
        _ = v41
    }
    var v42 string = "a"
    v3[4294967295] = false
    v1 += (v1 - v1)
    var v43 uint64 = 990
    v44 := v3[291]
    v45 := []int{-598}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
