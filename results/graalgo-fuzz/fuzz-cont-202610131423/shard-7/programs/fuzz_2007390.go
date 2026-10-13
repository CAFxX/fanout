package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int64, p2 int8) int {
    if true {
        v7 := GenS1{F0: 42}
        var v8 int32 = -2147483648
        _ = v7
        _ = v8
    }
    v9 := []int{0, 2147483647, 184}
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: 0}
        fmt.Println(p0, p1)
        p2 -= p2
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 32768
}
func genFn1(p0 uint) bool {
    v12 := GenS1{F0: 42}
    var v13 int32 = 1
    for fi14 := 0; fi14 < 4; fi14++ {
        v13 -= v13
        _ = fi14
    }
    p0 = (p0 * (^(p0)))
    _ = p0
    _ = v12
    _ = v13
    return (128 == genFn0(65535, -9223372036854775808, -128))
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint64 = 4294967295
    var v3 float64 = (((float64(-2.25) + float64(float32(1.5))) / (float64((0.0 / 0.1)) + -2.25)) / 1.0)
    var v4 int64 = -166
    var v5 uint = 1
    v6 := make(map[uint32]bool)
    var v15 uint32 = v1
    v6[v1] = false
    _ = len(v6)
    var v16 string = "\n\t"
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint8 = 157
        fmt.Println(v15, v16, v1)
        _ = fi17
        _ = v18
    }
    var v19 uint64 = v2
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 bool = (!false)
        fmt.Println(v21)
        var v22 uint64 = (v19 | v19)
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 string = v16
    v24 := v6[v15]
    if v24 {
        if ((523 > 0) || v24) {
            switch -9223372036854775808 {
            case -2147483648:
                var v25 uint32 = (^(v1))
                _ = v25
            case 9223372036854775807:
                var v26 int = -1
                _ = v26
            case -161:
                for fi27 := 0; fi27 < 14; fi27++ {
                    v28 := []int{int(v15), fi27, (fi27 | fi27), genFn0(32767, v4, -128)}
                    var v29 uint8 = 255
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            default:
                var v30 uint = 20
                _ = v30
            }
            fmt.Println(v3, v16)
        } else {
            fmt.Println(v19)
        }
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 int64 = v4
            if v24 {
                var v33 float32 = 1.5
                fmt.Println(v2, v15)
                var v34 uint = v5
                _ = v33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
    }
    switch 850 {
    case -9223372036854775808:
        _ = len(v6)
    }
    switch -897 {
    case 127:
        v35 := GenS1{F0: 1}
        _ = v35
    }
    if true {
        v36 := GenS1{F0: 193}
        _ = v36
    }
    if false {
        var v37 int32 = -128
        for fi38 := 0; fi38 < 3; fi38++ {
            if false {
                fmt.Println(v3)
            }
            _ = fi38
        }
        _ = v37
    } else {
        if ((!((2 != 0) && (-91 > 127))) && (!(((v2 + 0) != v2) && v24))) {
            if (-1 > -520) {
                v23 = ((v16 + "\n\t") + "\n\t")
                fmt.Println(v1, v5)
                var v39 int32 = -434
                _ = v39
            }
        }
    }
    var v40 int8 = 42
    v3 = (0.1 / (v3 * (v3 + (0.0 * 1e10))))
    fmt.Println(v4)
    var v41 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
