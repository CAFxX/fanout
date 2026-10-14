package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 bool; F2 bool }
type GenS4 struct { F0 uint }
func genFn0() int8 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{fi7}
        var v9 float32 = 0.1
        v10 := GenS1{F0: 2147483647, F1: 2}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if false {
        switch (("foo bar" + ("hello" + "foo bar")) + "\n\t") {
        case "a":
            var v11 uint32 = 1
            _ = v11
        case "\n\t":
            v12 := GenS1{F0: 993, F1: 0}
            _ = v12
        case "hello":
            var v13 int = 459
            _ = v13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint64 = 772
            v16 := []int{fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    if (2 < 2) {
        var v17 uint = 2
        fmt.Println(v17)
        _ = v17
    } else {
        v18 := GenS1{F0: -2147483648, F1: 17}
        if (42 >= 70) {
            var v19 uint = 255
            _ = v19
        }
        _ = v18
    }
    var v20 uint = 65535
    _ = v20
    return 1
}
func genFn1(p0 uint16) bool {
    if (0.0 >= (1.5 - 1e10)) {
        var v21 uint = 4294967295
        switch 127 {
        case -100:
            for fi22 := 0; fi22 < 15; fi22++ {
                v21 -= 4294967295
                v23 := []int{(fi22 % (((-(fi22))) | 1)), fi22}
                _ = fi22
                _ = v23
            }
        case -2147483648:
            fmt.Println(p0, v21)
        case 128:
            if false {
                v24 := GenS2{F0: 1}
                _ = v24
            } else {
                v25 := []int{32768, 1, 2147483647}
                p0 = p0
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 7; fi26++ {
                p0 -= (^(p0))
                switch "foo bar" {
                case "\n\t":
                    v21 += v21
                case "foo bar":
                    v27 := GenS2{F0: 2}
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 float32 = (1e10 - (3.14159 * 1.5))
                        _ = fi28
                        _ = v29
                    }
                }
                var v30 int8 = genFn0()
                _ = fi26
                _ = v30
            }
        }
        _ = v21
    } else {
        v31 := []int{}
        var v32 uint = 2
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 0, F1: 319}
    v34 := GenS3{F0: (("hello" + "foo bar") + "hello"), F1: true, F2: (!(!true))}
    for fi35 := 0; fi35 < 14; fi35++ {
        switch ("a" + ("hello" + "\n\t")) {
        case "\n\t":
            var v36 bool = (0 <= 42)
            _ = v36
        case "hello":
            v37 := GenS2{F0: 18446744073709551615}
            _ = v37
        }
        _ = fi35
    }
    if ((-2.25 / 0.1) <= 1e10) {
        v38 := []int{793}
        _ = v38
    } else {
        p0 -= 0
        var v39 uint = 65535
        var v40 int16 = -32768
        _ = v39
        _ = v40
    }
    _ = p0
    _ = v33
    _ = v34
    return true
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = "foo bar"
    var v3 int = 4294967295
    var v4 uint8 = 135
    v5 := make(map[string]string)
    var v6 uint16 = 30
    v1 = ((v1 / v1) / v1)
    var v41 float64 = 1e10
    var v42 bool = false
    if (((1 == -472) && ("\n\t" > "a")) && v42) {
        var v43 uint32 = 4294967295
        _ = v43
    } else {
        var v44 uint32 = 314
        _ = v44
    }
    v45 := GenS4{F0: 18446744073709551615}
    switch v2 {
    case "":
        var v46 float64 = 1e10
        _ = v46
    case "\n\t":
        v5[v2] = v2
    default:
        _ = len(v5)
    }
    var v47 float32 = 0.1
    if v42 {
        v3 += v3
    }
    var v48 int16 = 0
    var v49 float64 = (3.14159 + 0.0)
    v3 -= v3
    var v50 uint64 = 938
    var v51 int64 = 42
    var v52 float32 = ((0.1 - 0.1) / v47)
    var v53 int8 = int8(v48)
    v54 := []int{v3, v3}
    v6 = v6
    var v55 string = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
