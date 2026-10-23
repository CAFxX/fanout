package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int; F2 bool }
type GenS5 struct { F0 bool }
func genFn0() int8 {
    var v5 bool = false
    if v5 {
        var v6 float64 = ((-2.25 - (float64((-2.25 * float32((1e10 - (((float64(0.1) / float64(((float32(3.14159) + 1.5) / 1.0))) / float64(float32(1.5))) + (float64((0.1 / 1e10)) - 1e10)))))) / (-2.25 / ((1e10 + 0.0) - (((float64(0.1) - 0.1) - (0.0 * -2.25)) + ((0.0 * 1e10) + 0.0)))))) - float64(0.1))
        _ = v6
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch fi7 {
            case 127:
                var v8 uint32 = 0
                _ = v8
            case 736:
                v9 := []int{fi7, 128}
                _ = v9
            }
            v5 = (9223372036854775807 < 985)
            _ = fi7
        }
        v5 = v5
    }
    switch ("a" + ("foo bar" + ("hello" + "hello"))) {
    case "":
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(fi11, fi10, fi10)
                var v12 float32 = float32(0.0)
                v13 := []int{127}
                _ = fi11
                _ = v12
                _ = v13
            }
            v5 = true
            for fi14 := 0; fi14 < 11; fi14++ {
                v5 = v5
                _ = fi14
            }
            _ = fi10
        }
    }
    _ = v5
    return 1
}

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    v3 := []string{}
    var v4 uint8 = 128
    v1 = v1
    var v15 bool = v1
    if (float64((((0.1 - 0.0) * (((0.0 + 0.0) / float32(-2.25)) + (0.1 + 0.1))) / 1.5)) < 0.0) {
        switch 127 {
        case 32768:
            var v16 int64 = -1
            _ = v16
        default:
            v17 := GenS1{F0: 4294967295}
            _ = v17
        }
        v18 := v2[32768]
        _ = v18
    } else {
        v19 := GenS1{F0: -9223372036854775808}
        v4 = v4
        for fi20 := 0; fi20 < 16; fi20++ {
            fmt.Println(fi20, v1)
            _ = fi20
        }
        _ = v19
    }
    fmt.Println(v15)
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v1 = v1
    v21 := GenS1{F0: -93}
    v22 := GenS2{F0: 42, F1: 201, F2: 2}
    if v1 {
        v4 -= (v4 + (v4 ^ 1))
    } else {
        var v23 uint64 = 211
        _ = v23
    }
    v15 = v1
    switch (("\n\t" + (("" + ("hello" + "hello")) + "a")) + "\n\t") {
    case "a":
        v15 = v1
    case "hello":
        if v15 {
            if v1 {
                v24 := GenS3{F0: v1}
                v15 = true
                _ = v24
            } else {
                v1 = v1
                for fi25 := 0; fi25 < 7; fi25++ {
                    v1 = (!(!v15))
                    v26 := GenS3{F0: false}
                    v4 = (v4 >> ((2) & 63))
                    _ = fi25
                    _ = v26
                }
            }
        }
    default:
        for fi27 := 0; fi27 < 5; fi27++ {
            v4 += 220
            _ = fi27
        }
    }
    switch (("hello" + "foo bar") + "\n\t") {
    case "foo bar":
        v1 = v1
    default:
        v4 = (v4 * v4)
    }
    var v28 uint32 = 756
    v29 := GenS4{F0: -299, F1: 9223372036854775807, F2: (1 != 255)}
    for fi30 := 0; fi30 < 11; fi30++ {
        if (42 == 1) {
            for fi31 := 0; fi31 < 5; fi31++ {
                v4 = (v4 - 76)
                fmt.Println(fi31, fi31)
                var v32 int64 = 653
                switch ("" + (("hello" + "hello") + (("" + "a") + "\n\t"))) {
                case "\n\t":
                    var v33 int8 = 100
                    if len(v3) > 0 {
                        v3[0] = ("a" + "foo bar")
                    }
                    v3 = append(v3, "\n\t")
                    _ = v33
                default:
                    var v34 int = 875
                    _ = v34
                }
                var v35 bool = v1
                _ = len(v3)
                _ = fi31
                _ = v32
                _ = v35
            }
            switch "a" {
            case "":
                switch "foo bar" {
                case "a":
                    if (!v1) {
                        v1 = false
                        for fi36 := 0; fi36 < 13; fi36++ {
                            if (true || v1) {
                                var v37 int = fi36
                                fmt.Println(v28, fi30)
                                switch "" {
                                case "":
                                    v37 = (v37 - fi30)
                                }
                                _ = v37
                            }
                            _ = fi36
                        }
                        v4 -= 255
                    } else {
                        var v38 uint16 = 255
                        var v39 float32 = 0.0
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        _ = v38
                        _ = v39
                    }
                case "\n\t":
                    v28 = 65535
                case "":
                    v40 := GenS5{F0: (1 >= 255)}
                    _ = v40
                }
            case "foo bar":
                fmt.Println(v4)
            case "\n\t":
                var v41 uint32 = v28
                _ = v41
            default:
                if ((0.1 < 3.14159) || (v15 || (!(!v15)))) {
                    var v42 int = fi30
                    var v43 string = ""
                    v4 = (v4 | v4)
                    _ = v42
                    _ = v43
                }
            }
        } else {
            v4 = 255
        }
        _ = fi30
    }
    delete(v2, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
