package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
func genFn0() uint {
    var v6 bool = true
    v6 = (!v6)
    v7 := GenS1{F0: 238, F1: 42, F2: 1}
    _ = v6
    _ = v7
    return 586
}
func genFn1() int {
    var v8 bool = (-9223372036854775808 >= 819)
    if (127 > 255) {
        v8 = true
        var v9 int16 = 32767
        var v10 int16 = v9
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch ("foo bar" + "hello") {
        case "":
            var v12 uint = 1000
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 uint8 = 42
                if v8 {
                    v15 := GenS1{F0: 0, F1: 32767, F2: (fi13 | fi11)}
                    _ = v15
                }
                if v8 {
                    switch (^(fi13)) {
                    case 9223372036854775807:
                        switch ("foo bar" + ("" + "foo bar")) {
                        case "a":
                            v14 -= v14
                        default:
                            v16 := []int{(fi11 * fi11), 127, fi11}
                            _ = v16
                        }
                    }
                    v8 = (255 == 321)
                }
                _ = fi13
                _ = v14
            }
        default:
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint8 = 77
                if v8 {
                    switch 42 {
                    case 810:
                        if v8 {
                            v18 -= 0
                        }
                    case 9223372036854775807:
                        var v19 uint32 = 343
                        _ = v19
                    }
                }
                v20 := GenS1{F0: 255, F1: (fi17 ^ fi17), F2: fi11}
                _ = fi17
                _ = v18
                _ = v20
            }
        }
        v8 = v8
        _ = fi11
    }
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 int16 = 0
    var v5 float64 = (0.0 / -2.25)
    v4 += (^(v4))
    if false {
        v21 := []int{}
        var v22 int32 = -912
        _ = v21
        _ = v22
    }
    v5 -= (3.14159 + -2.25)
    v23 := GenS1{F0: 456, F1: genFn1(), F2: 42}
    if false {
        v1[-437] = 127
        v5 = v5
        v1[465] = genFn1()
    } else {
        for fi24 := 0; fi24 < 3; fi24++ {
            if false {
                if false {
                    fmt.Println(v4)
                    v4 += (v4 - (-(v4)))
                    v5 += 0.0
                }
                var v25 int32 = -1
                _ = v25
            } else {
                v26 := []int{(fi24 << ((499) & 63)), (fi24 * genFn1()), (fi24 << ((42) & 63))}
                _ = v26
            }
            fmt.Println(v5, v5)
            v27 := GenS2{F0: -128, F1: 18446744073709551615, F2: -120}
            _ = fi24
            _ = v27
        }
        var v28 uint16 = 726
        _ = v28
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        v4 -= 32767
        v30 := GenS2{F0: 100, F1: 18446744073709551615, F2: 0}
        _ = fi29
        _ = v30
    }
    v5 = 1.5
    v3[genFn1()] = "hello"
    if false {
        v5 = (v5 + (0.0 + v5))
        var v31 float64 = v5
        for fi32 := 0; fi32 < 5; fi32++ {
            switch fi32 {
            case 1:
                for fi33 := 0; fi33 < 6; fi33++ {
                    var v34 uint64 = 603
                    var v35 int16 = 1
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case 0:
                v36 := GenS2{F0: -15, F1: 1, F2: 100}
                _ = v36
            case -343:
                v37 := GenS1{F0: 381, F1: (fi32 & (fi32 >> ((1) & 63))), F2: (fi32 << ((658) & 63))}
                _ = v37
            }
            v38 := GenS3{F0: 2, F1: -2147483648, F2: 0}
            var v39 float64 = ((0.1 + -2.25) - (3.14159 - float64(float32(v31))))
            _ = fi32
            _ = v38
            _ = v39
        }
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
