package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
func genFn0() uint16 {
    v5 := GenS1{F0: 0}
    var v6 int8 = -128
    v6 = -95
    _ = v5
    _ = v6
    return 668
}
func genFn1(p0 int) string {
    if true {
        var v7 int16 = 173
        var v8 int8 = 100
        _ = v7
        _ = v8
    }
    var v9 bool = (!(true || false))
    for fi10 := 0; fi10 < 13; fi10++ {
        if ((!true) || v9) {
            for fi11 := 0; fi11 < 5; fi11++ {
                if ((2 != 873) && (!(!false))) {
                    v12 := GenS1{F0: 42}
                    _ = v12
                }
                p0 += fi10
                _ = fi11
            }
            fmt.Println(p0, p0)
        }
        v13 := []int{(p0 + 2147483647), (p0 | (fi10 | -1))}
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = v9
    return "a"
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 bool = true
    var v4 int8 = 42
    delete(v1, 1)
    v4 = 0
    var v14 int8 = (v4 | (v4 >> ((1) & 63)))
    var v15 int = 0
    var v16 int16 = 813
    switch (v15 - 127) {
    case 1:
        v15 += 487
    case 2147483647:
        v17 := GenS2{F0: (!true)}
        _ = v17
    case 42:
        fmt.Println(v3, v14)
    default:
        v3 = false
    }
    var v18 int64 = -43
    fmt.Println(v18)
    var v19 uint32 = 1
    var v20 uint8 = 0
    if false {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 uint64 = 2
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS1{F0: v18}
            v25 := GenS2{F0: ((v18 - (v18 >> ((1) & 63))) < v18)}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := v1[v15]
        v20 = (v20 ^ v20)
        v16 = -1
        _ = fi26
        _ = v27
    }
    if true {
        var v28 uint64 = uint64(v20)
        fmt.Println(v4, v16, v15)
        _ = v28
    } else {
        var v29 uint32 = 2147483648
        _ = v29
    }
    var v30 uint32 = v19
    switch v15 {
    case -9223372036854775808:
        fmt.Println(v16, v30)
    case 4294967295:
        switch "hello" {
        case "a":
            switch (("foo bar" + "a") + "hello") {
            case "hello":
                if ((v4 - -6) < v14) {
                    v15 = v15
                }
            case "\n\t":
                v31 := []int{v15, 4294967295}
                _ = v31
            default:
                fmt.Println(v19, v14, v4)
            }
        default:
            v18 += v18
        }
    case -253:
        var v32 int = (^(v15))
        _ = v32
    }
    if (42 > 0) {
        if false {
            v30 = v30
        } else {
            var v33 bool = v2
            fmt.Println(v18)
            var v34 float64 = ((((-2.25 + 3.14159) + 0.1) / (((-2.25 + (-2.25 * (0.1 + 1e10))) * 0.0) - (float64(1e10) * (float64(-2.25) + ((1e10 + -2.25) + 3.14159))))) - -2.25)
            _ = v33
            _ = v34
        }
    }
    if (!(!v3)) {
        v19 = 331
        var v35 uint32 = v30
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
}
