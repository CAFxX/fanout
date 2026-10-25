package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0(p0 bool) string {
    var v7 uint32 = 574
    var v8 bool = (1.5 < 1.5)
    v9 := []int{}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1() string {
    v10 := []int{0}
    if true {
        if (!(!(0 != 2))) {
            if true {
                var v11 uint64 = 42
                v11 = v11
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint16 = 32767
                _ = fi12
                _ = v13
            }
            var v14 string = "\n\t"
            v14 = "hello"
            _ = v14
        }
        var v15 uint16 = 32767
        v16 := GenS1{F0: 42, F1: -1}
        _ = v15
        _ = v16
    }
    if (42 > 1) {
        switch -700 {
        case 0:
            var v17 uint8 = 1
            _ = v17
        case 42:
            var v18 int64 = 42
            _ = v18
        case 525:
            var v19 int64 = -345
            _ = v19
        default:
            var v20 int64 = 120
            _ = v20
        }
        if true {
            var v21 int = 763
            v21 -= 128
            v21 = v21
            _ = v21
        } else {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v22 int32 = 0
            _ = v22
        }
    } else {
        if false {
            var v23 int16 = 678
            for fi24 := 0; fi24 < 6; fi24++ {
                fmt.Println(v23)
                fmt.Println(v23)
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 uint = 255
                    if false {
                        var v27 int32 = -1
                        v28 := GenS1{F0: -9223372036854775808, F1: -1}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 string = "\n\t"
                        var v30 uint8 = 128
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
            fmt.Println(v23)
            _ = v23
        }
        var v31 string = "a"
        fmt.Println(v31)
        _ = v31
    }
    _ = v10
    return ("" + "")
}

func main() {
    v1 := []string{}
    var v2 uint = 2
    var v3 float32 = (((1e10 * (0.1 - (-2.25 / 0.1))) + (1.5 + 3.14159)) - float32(0.0))
    v4 := []string{""}
    var v5 float32 = (float32(0.1) - v3)
    var v6 uint16 = 2
    v5 += v5
    if false {
        var v32 uint64 = 2
        _ = v32
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v33 int = 32767
        _ = v33
    case 32767:
        var v34 uint64 = 42
        _ = v34
    case 32768:
        v35 := GenS2{F0: 502, F1: 127}
        _ = v35
    default:
        v5 += v5
    }
    switch 31 {
    case 42:
        v3 -= float32(1.5)
    case -9223372036854775808:
        if false {
            v36 := GenS2{F0: 0, F1: 1}
            var v37 uint64 = 1
            _ = v36
            _ = v37
        }
    case -1:
        v38 := GenS1{F0: -1, F1: 9223372036854775807}
        _ = v38
    }
    v6 += (v6 | v6)
    v5 -= v5
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 bool = true
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        var v41 int32 = 861
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v2, v3, v3)
    if false {
        v42 := GenS2{F0: -9223372036854775808, F1: 128}
        _ = len(v4)
        _ = v42
    } else {
        fmt.Println(v3, v3, v5)
    }
    var v43 float64 = 1.5
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    var v44 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
