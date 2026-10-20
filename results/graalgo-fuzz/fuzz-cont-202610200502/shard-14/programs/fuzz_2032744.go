package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 int32) uint16 {
    v4 := []int{-410, 32767, 4294967295}
    switch 128 {
    case 128:
        v5 := GenS1{F0: 343, F1: 195}
        _ = v5
    default:
        var v6 bool = (!((1 < 825) && false))
        _ = v6
    }
    switch "a" {
    case "\n\t":
        v7 := GenS1{F0: 42, F1: 1}
        _ = v7
    case "":
        fmt.Println(p0)
    case "foo bar":
        if true {
            switch ("" + ((("\n\t" + "a") + "\n\t") + "hello")) {
            case "":
                fmt.Println(p0)
            case "foo bar":
                _ = len(v4)
            case "hello":
                p0 += 12
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
    }
    var v8 uint64 = 391
    var v9 string = "a"
    _ = p0
    _ = v4
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 bool) uint64 {
    p0 = (-2.25 != (0.0 / 0.1))
    var v10 uint8 = 128
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 float64 = 1.5
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    return 4294967295
}

func main() {
    var v1 uint64 = 42
    var v2 bool = (true || true)
    var v3 bool = (-430 >= -1)
    fmt.Println(v3, v3)
    var v13 string = ((("a" + "") + "\n\t") + "a")
    switch "hello" {
    case "a":
        v14 := GenS1{F0: 1000, F1: 2}
        _ = v14
    case "":
        for fi15 := 0; fi15 < 2; fi15++ {
            switch (v13 + v13) {
            case "a":
                v1 += (v1 | (v1 - (v1 * 4294967295)))
            case "foo bar":
                if (0 >= 1) {
                    v16 := GenS1{F0: 1000, F1: 128}
                    var v17 uint = 0
                    v18 := []int{fi15, (fi15 - fi15)}
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi15
        }
    default:
        var v19 uint = 245
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            if v2 {
                var v22 uint32 = 71
                var v23 uint8 = 116
                _ = v22
                _ = v23
            } else {
                v24 := GenS1{F0: 1, F1: 1}
                if (127 <= 2) {
                    v25 := []int{(fi20 / ((fi21) | 1)), (fi21 << ((642) & 63)), (fi20 - fi20)}
                    var v26 uint = 4294967295
                    _ = v25
                    _ = v26
                } else {
                    switch (fi20 << ((uint(v1)) & 63)) {
                    case 42:
                        if true {
                            var v27 uint = 4294967295
                            var v28 uint = v27
                            v3 = (!v3)
                            _ = v27
                            _ = v28
                        } else {
                            v13 = v13
                            v2 = (201 == 103)
                        }
                    case -2147483648:
                        v29 := []int{32767, fi21}
                        _ = v29
                    }
                    var v30 uint16 = 255
                    var v31 bool = (!v3)
                    _ = v30
                    _ = v31
                }
                _ = v24
            }
            v32 := GenS2{F0: 957, F1: (v13 + (v13 + ((v13 + "foo bar") + v13)))}
            _ = fi21
            _ = v32
        }
        var v33 uint = 2
        var v34 uint16 = 65535
        _ = fi20
        _ = v33
        _ = v34
    }
    var v35 float64 = ((-2.25 / (1e10 / -2.25)) + 0.0)
    v36 := []int{}
    switch ("foo bar" + "hello") {
    case "a":
        if (v13 <= v13) {
            var v37 uint32 = 2147483647
            var v38 int16 = 42
            _ = v37
            _ = v38
        } else {
            v3 = v3
        }
    case "\n\t":
        var v39 int = 189
        _ = v39
    }
    v3 = v2
    if (v35 < v35) {
        if false {
            v13 = v13
            v36 = append(v36, -332)
        }
        v40 := GenS3{F0: 154, F1: 142}
        v2 = (v2 || (1 >= 974))
        _ = v40
    }
    var v41 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v41
}
