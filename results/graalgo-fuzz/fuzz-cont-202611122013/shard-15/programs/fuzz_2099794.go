package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16) int {
    v6 := []int{128, 127, 42, -9223372036854775808}
    _ = len(v6)
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    var v1 float64 = 1e10
    var v2 int16 = -860
    var v3 string = "hello"
    var v4 uint = 632
    v5 := []int{127, 911}
    v1 += -2.25
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 = ("" + "")
        fmt.Println(v2)
        _ = fi7
    }
    var v8 int16 = 1
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int16 = 204
        _ = fi9
        _ = v10
    }
    if true {
        for fi11 := 0; fi11 < 20; fi11++ {
            v8 += (v2 & v2)
            _ = fi11
        }
        var v12 uint32 = 1
        fmt.Println(v8)
        _ = v12
    } else {
        fmt.Println(v8, v8, v2)
        v3 = "\n\t"
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v1 -= v1
        v1 = -2.25
        _ = fi13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int16 = v8
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 128
    for fi17 := 0; fi17 < 15; fi17++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v18 string = v3
        _ = fi17
        _ = v18
    }
    if (9223372036854775807 > 110) {
        v16 = v16
    } else {
        var v19 uint64 = 394
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v5 = append(v5, 127)
        _ = fi20
    }
    var v21 uint = (v4 + v4)
    fmt.Println(v16)
    if false {
        fmt.Println(v8)
        for fi22 := 0; fi22 < 1; fi22++ {
            v2 = int16(v16)
            switch ((v3 + "") + ("a" + v3)) {
            case "foo bar":
                var v23 bool = true
                _ = v23
            case "\n\t":
                for fi24 := 0; fi24 < 3; fi24++ {
                    v2 -= v2
                    v21 = 442
                    fmt.Println(v16, fi22, v21)
                    _ = fi24
                }
            case "":
                v25 := GenS1{F0: -23}
                _ = v25
            default:
                if false {
                    if false {
                        var v26 int = fi22
                        _ = v26
                    }
                    if false {
                        switch (fi22 % (((fi22 ^ (fi22 + genFn0(42)))) | 1)) {
                        case 0:
                            if ((!(!false)) || (!false)) {
                                v3 = "foo bar"
                                v21 = v21
                                v5 = append(v5, genFn0(1))
                            } else {
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    v1 -= ((float64(3.14159) + v1) + v1)
                                    v3 = v3
                                    _ = fi27
                                }
                            }
                        case 292:
                            v16 = (^(v16))
                        default:
                            var v28 uint64 = 633
                            _ = v28
                        }
                    } else {
                        var v29 float64 = 1e10
                        _ = v29
                    }
                }
            }
            _ = fi22
        }
        v16 -= 0
    } else {
        var v30 int8 = -128
        fmt.Println(v8)
        v31 := GenS2{F0: v4}
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
