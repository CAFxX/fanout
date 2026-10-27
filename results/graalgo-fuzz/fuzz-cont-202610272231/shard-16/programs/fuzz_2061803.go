package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 bool) int32 {
    if ((727 <= 255) && (0.0 < 0.0)) {
        var v4 int8 = 29
        _ = v4
    }
    p0 = (p0 && (-2147483648 >= 0))
    _ = p0
    return -2147483648
}

func main() {
    var v1 int32 = 704
    var v2 uint64 = 1
    var v3 int = 4294967295
    var v5 int16 = -32768
    if (42 < 1) {
        for fi6 := 0; fi6 < 16; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{(v3 & v3), (fi7 % ((fi7) | 1)), (-(fi7))}
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        fmt.Println(v1, v2, v5)
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    var v10 string = ("hello" + (("foo bar" + "\n\t") + "hello"))
    if (!(410 != 65535)) {
        v11 := GenS1{F0: v3, F1: true, F2: v2}
        _ = v11
    }
    fmt.Println(v2, v10)
    for fi12 := 0; fi12 < 19; fi12++ {
        v5 = v5
        _ = fi12
    }
    v13 := GenS1{F0: v3, F1: true, F2: v2}
    for fi14 := 0; fi14 < 12; fi14++ {
        if true {
            var v15 float64 = 3.14159
            v2 += (v2 + 4294967295)
            _ = v15
        }
        _ = fi14
    }
    var v16 int8 = -77
    if true {
        fmt.Println(v3, v5)
        var v17 float64 = 0.1
        _ = v17
    }
    var v18 float32 = (float32((((1e10 + 3.14159) / 1e10) / -2.25)) / 1.5)
    var v19 string = ((v10 + (v10 + "foo bar")) + v10)
    switch "" {
    case "foo bar":
        v1 = (v1 >> ((954) & 63))
    default:
        if false {
            for fi20 := 0; fi20 < 3; fi20++ {
                if false {
                    switch v19 {
                    case "":
                        var v21 int64 = 1
                        _ = v21
                    case "a":
                        switch (fi20 / ((v3) | 1)) {
                        case 127:
                            if true {
                                for fi22 := 0; fi22 < 18; fi22++ {
                                    v23 := []int{469}
                                    if ((((((float32(3.14159) + 0.1) / v18) + -2.25) - (float32(((3.14159 / -2.25) - -2.25)) + 0.1)) != 1.5) && true) {
                                        _ = len(v23)
                                        v24 := GenS1{F0: (fi22 >> ((4294967295) & 63)), F1: false, F2: v2}
                                        _ = v24
                                    } else {
                                        fmt.Println(v3, fi22, v5)
                                        v25 := GenS2{F0: (v19 + ((v10 + v19) + (v10 + v19)))}
                                        _ = v25
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        case 42:
                            v26 := []int{560, v3}
                            _ = v26
                        default:
                            var v27 float64 = ((((1.5 * -2.25) / (1.5 * (0.1 * (3.14159 * float64((v18 + (0.1 - (v18 - (v18 * v18))))))))) + 1e10) * float64((0.1 + (-2.25 + (v18 - 1e10)))))
                            _ = v27
                        }
                    case "hello":
                        switch ((v10 + "hello") + ("" + v19)) {
                        case "a":
                            fmt.Println(v1, v19)
                        default:
                            for fi28 := 0; fi28 < 6; fi28++ {
                                var v29 uint = 255
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                } else {
                    v30 := []int{1}
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 int64 = 0
                        if len(v30) > 0 {
                            v30[0] = fi20
                        }
                        _ = fi31
                        _ = v32
                    }
                    var v33 int = (-(fi20))
                    _ = v30
                    _ = v33
                }
                _ = fi20
            }
        } else {
            fmt.Println(v19, v2)
            if ((v18 + v18) == (-2.25 / ((-2.25 - (v18 - v18)) + 3.14159))) {
                v34 := GenS1{F0: v3, F1: false, F2: 2}
                fmt.Println(v2, v1)
                var v35 uint8 = 128
                _ = v34
                _ = v35
            }
            var v36 uint = 42
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
}
