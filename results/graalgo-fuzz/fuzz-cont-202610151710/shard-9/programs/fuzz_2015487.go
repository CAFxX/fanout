package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0() uint8 {
    var v6 float32 = 3.14159
    if (!true) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{(fi7 + (fi7 * fi7)), (fi7 >> ((848) & 63)), fi7}
            _ = fi7
            _ = v8
        }
        v9 := []int{128}
        _ = v9
    }
    _ = v6
    return 128
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v5 := []uint32{146, 4294967295}
    var v10 bool = true
    var v11 bool = v10
    v10 = v10
    if ((((v1 - 1e10) >= v1) && v10) && (-812 >= 1)) {
        if len(v5) > 0 {
            v5[0] = 686
        }
        var v12 int16 = -32768
        if v10 {
            var v13 float64 = (1e10 + -2.25)
            _ = v13
        }
        _ = v12
    } else {
        if v11 {
            var v14 float32 = v1
            v15 := GenS1{F0: 42, F1: -1, F2: 100}
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 uint64 = 2
                _ = fi16
                _ = v17
            }
        }
        var v18 uint = 0
        v11 = (1 <= genFn0())
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 float32 = float32(-2.25)
        if (((!("\n\t" >= "")) && (true && (!(v10 || v10)))) || (583 > 255)) {
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        }
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if v10 {
            if (!v11) {
                var v22 int = 1
                v4[v22] = (("hello" + "foo bar") + ("a" + ""))
                v23 := GenS1{F0: -1, F1: fi21, F2: 100}
                _ = v22
                _ = v23
            }
            v10 = true
        } else {
            delete(v4, fi21)
            var v24 int = 4294967295
            _ = len(v5)
            _ = v24
        }
        var v25 int = (-(fi21))
        _ = fi21
        _ = v25
    }
    var v26 int64 = -708
    fmt.Println(v10, v26, v10)
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 float32 = v1
        fmt.Println(fi27)
        var v29 string = "hello"
        _ = fi27
        _ = v28
        _ = v29
    }
    v2[0] = 2147483647
    var v30 uint8 = 1
    if true {
        for fi31 := 0; fi31 < 18; fi31++ {
            v32 := GenS1{F0: 0, F1: (fi31 & fi31), F2: 127}
            v30 = (^(v30))
            _ = fi31
            _ = v32
        }
    } else {
        var v33 int16 = -549
        _ = v33
    }
    if true {
        v5 = append(v5, 42)
        fmt.Println(v1, v26, v10)
        var v34 float64 = -2.25
        _ = v34
    }
    fmt.Println(v1, v30)
    switch 415 {
    case 4294967295:
        v11 = (false && (!(706 >= 163)))
    case 127:
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := GenS2{F0: -9223372036854775808}
            _ = fi35
            _ = v36
        }
    case 128:
        v37 := GenS1{F0: 80, F1: 9223372036854775807, F2: 1}
        _ = v37
    default:
        v38 := GenS3{F0: 527, F1: 2}
        _ = v38
    }
    v39 := v2[600]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v5
}
