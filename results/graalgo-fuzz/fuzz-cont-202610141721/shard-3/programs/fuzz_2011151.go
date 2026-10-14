package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int8 }
func genFn0(p0 int64) int {
    if true {
        switch "a" {
        case "hello":
            fmt.Println(p0)
        case "foo bar":
            p0 -= (p0 ^ p0)
        case "\n\t":
            var v4 float32 = 1.5
            _ = v4
        }
        switch ("a" + "\n\t") {
        case "":
            var v5 int32 = 80
            _ = v5
        case "\n\t":
            v6 := []int{}
            _ = v6
        case "foo bar":
            v7 := GenS1{F0: false, F1: -409}
            _ = v7
        }
    } else {
        v8 := GenS2{F0: 65535}
        var v9 int8 = 0
        v10 := GenS1{F0: false, F1: 347}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int8 = -128
        if true {
            for fi13 := 0; fi13 < 7; fi13++ {
                switch (^(fi11)) {
                case -1:
                    if false {
                        var v14 float32 = 0.0
                        _ = v14
                    } else {
                        var v15 uint32 = 65535
                        var v16 int32 = -804
                        if (!false) {
                            v16 -= v16
                        }
                        _ = v15
                        _ = v16
                    }
                case 9223372036854775807:
                    v17 := GenS3{F0: 158}
                    _ = v17
                case 127:
                    v18 := []int{(fi13 / ((fi11) | 1))}
                    if ((fi11 > (fi13 & 95)) || true) {
                        fmt.Println(p0)
                        var v19 int64 = (^(p0))
                        var v20 uint32 = 740
                        _ = v19
                        _ = v20
                    }
                    _ = v18
                }
                _ = fi13
            }
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v12 = (v12 >> ((4294967295) & 63))
            v22 := []int{(fi21 | (fi21 << ((546) & 63))), 840}
            var v23 int8 = 68
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    return 127
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1e10
    var v3 float64 = (float64((0.0 - 1e10)) - 0.0)
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := []int{127}
        fmt.Println(v2)
        _ = fi24
        _ = v25
    }
    delete(v1, (("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))) + "\n\t"))
    v26 := GenS2{F0: 364}
    v3 = v3
    var v27 uint16 = 2
    var v28 string = (("" + "\n\t") + ("foo bar" + ""))
    v29 := []int{-155, -322, -200}
    switch ("" + v28) {
    case "a":
        v30 := GenS2{F0: 65535}
        _ = v30
    case "foo bar":
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS1{F0: (!(true || (89 == -67))), F1: fi31}
            for fi33 := 0; fi33 < 10; fi33++ {
                switch ("" + v28) {
                case "\n\t":
                    v29 = append(v29, fi33)
                case "foo bar":
                    v34 := GenS2{F0: 921}
                    _ = v34
                default:
                    v35 := GenS3{F0: -866}
                    _ = v35
                }
                v28 = ("a" + (v28 + v28))
                if true {
                    v36 := GenS4{F0: v27}
                    _ = v36
                }
                _ = fi33
            }
            _ = fi31
            _ = v32
        }
    }
    if len(v29) > 0 {
        v29[0] = 127
    }
    switch 9223372036854775807 {
    case -1:
        switch -2147483648 {
        case 127:
            _ = len(v29)
        }
    case 9223372036854775807:
        fmt.Println(v27)
    default:
        v29 = append(v29, 419)
    }
    fmt.Println(v3)
    switch 32767 {
    case -2147483648:
        var v37 uint8 = 255
        _ = v37
    case 641:
        if (1 >= 2) {
            _ = len(v1)
            if true {
                var v38 int64 = 0
                v1[((((v28 + v28) + v28) + ("hello" + "")) + (v28 + ("foo bar" + "")))] = true
                v39 := GenS5{F0: 562, F1: 65}
                _ = v38
                _ = v39
            } else {
                fmt.Println(v28)
            }
            v29 = append(v29, 493)
        } else {
            if false {
                v40 := GenS5{F0: -343, F1: 127}
                _ = v40
            } else {
                v3 = ((-2.25 / float64(((3.14159 + (1e10 * ((3.14159 - 1.5) / float32((((float64(0.0) / v3) / v2) + 1e10))))) - (1.5 + 3.14159)))) * (0.1 + 0.0))
                v3 = float64((float32((v3 / 1e10)) - -2.25))
            }
            var v41 int8 = 1
            fmt.Println(v27, v28, v27)
            _ = v41
        }
    }
    fmt.Println(v28, v2, v3)
    var v42 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v42
}
