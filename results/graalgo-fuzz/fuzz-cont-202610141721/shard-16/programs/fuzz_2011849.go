package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int16 }
func genFn0(p0 int, p1 bool) int16 {
    v5 := GenS1{F0: 65535}
    switch p0 {
    case -716:
        p1 = (!p1)
    case 1:
        var v6 float64 = 3.14159
        _ = v6
    default:
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := []int{fi7, (fi7 | fi7)}
            var v9 uint = 255
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    p0 = 32768
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint32) string {
    var v10 string = "\n\t"
    v11 := GenS1{F0: 0}
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            p0 = (p0 ^ p0)
            _ = fi13
        }
        _ = fi12
    }
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS1{F0: 426}
            v10 = (v10 + "foo bar")
            var v16 int16 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        v17 := []int{0, -926}
        _ = v17
    }
    if true {
        p0 -= 2
    }
    _ = p0
    _ = v10
    _ = v11
    return "foo bar"
}

func main() {
    var v1 bool = (-1 == 42)
    var v2 int = -1
    v3 := []string{}
    var v4 int = (v2 & 0)
    var v18 uint64 = 1
    var v19 string = "a"
    var v20 uint16 = 1
    var v21 int64 = 451
    if true {
        var v22 uint32 = 390
        var v23 float64 = ((0.1 + 0.0) - 3.14159)
        var v24 int32 = 42
        _ = v22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: v20}
    for fi26 := 0; fi26 < 13; fi26++ {
        v27 := GenS1{F0: (^(v20))}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi26
        _ = v27
    }
    if false {
        v28 := GenS1{F0: (^(v20))}
        v1 = (2 <= uint8(v18))
        for fi29 := 0; fi29 < 8; fi29++ {
            if ((((0.0 - 3.14159) + 0.1) + 0.1) > 0.0) {
                v18 = v18
                for fi30 := 0; fi30 < 3; fi30++ {
                    v3 = append(v3, "hello")
                    _ = fi30
                }
                if true {
                    v31 := GenS2{F0: 0, F1: v2, F2: int16(fi29)}
                    _ = len(v3)
                    for fi32 := 0; fi32 < 14; fi32++ {
                        var v33 float64 = -2.25
                        _ = fi32
                        _ = v33
                    }
                    _ = v31
                } else {
                    var v34 int64 = 0
                    _ = len(v3)
                    _ = v34
                }
            }
            _ = fi29
        }
        _ = v28
    }
    if false {
        var v35 int8 = int8(v21)
        _ = v35
    } else {
        var v36 float32 = 3.14159
        var v37 string = v19
        v38 := GenS2{F0: 42, F1: (v2 / ((v2) | 1)), F2: genFn0(v2, ((v21 | -334) > -9223372036854775808))}
        _ = v36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        v19 = "hello"
        var v40 uint32 = 65535
        _ = fi39
        _ = v40
    }
    switch ("foo bar" + ("" + ("foo bar" + (v19 + v19)))) {
    case "":
        var v41 int64 = v21
        _ = v41
    case "foo bar":
        fmt.Println(v20)
    default:
        v4 = (v4 % ((-1) | 1))
    }
    v2 = v4
    v3 = append(v3, v19)
    switch 32768 {
    case 32767:
        _ = len(v3)
    case 1:
        var v42 uint64 = v18
        _ = v42
    default:
        v1 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
}
