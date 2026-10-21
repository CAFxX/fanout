package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint, p1 int32, p2 uint) uint {
    var v4 int8 = -43
    p2 = p0
    v5 := []int{9223372036854775807, 0}
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (^(p0))
}
func genFn1(p0 int64) int32 {
    p0 = 82
    v6 := GenS1{F0: 1}
    var v7 float64 = 0.1
    var v8 int8 = 1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 898
}

func main() {
    var v1 float64 = 1e10
    var v2 bool = ("foo bar" != "foo bar")
    var v3 float64 = 1.5
    v9 := []int{-44, 32768, 861, 127}
    if false {
        if (-881 != -2147483648) {
            var v10 int64 = -68
            if v2 {
                v1 = ((v3 - (((float64((3.14159 / (float32(v3) * ((0.0 + 0.1) + -2.25)))) / v1) - v1) - (v3 * 1e10))) - ((v1 * v3) / float64((3.14159 - (float32(v3) / (0.1 / 1.5))))))
                fmt.Println(v3, v10, v2)
                fmt.Println(v1)
            } else {
                switch 42 {
                case -1:
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 bool = (255 >= 235)
                        v13 := GenS1{F0: 2}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                case 127:
                    v14 := GenS1{F0: 0}
                    _ = v14
                case 564:
                    if v2 {
                        var v15 uint32 = 1
                        _ = v15
                    } else {
                        var v16 uint8 = 137
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v9 = append(v9, (fi17 >> ((1) & 63)))
                            _ = fi17
                        }
                        if v2 {
                            fmt.Println(v16, v1, v16)
                            fmt.Println(v3, v2)
                            var v18 uint = 1
                            _ = v18
                        }
                        _ = v16
                    }
                }
                if (1 == 0) {
                    v3 = v3
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v9 = append(v9, (fi19 >> ((2) & 63)))
                        _ = fi19
                    }
                }
            }
            if true {
                switch 32767 {
                case 1:
                    var v20 float32 = (3.14159 - ((-2.25 - 1.5) * (float32(v3) - (0.0 - 1.5))))
                    _ = v20
                case 0:
                    fmt.Println(v10)
                case 292:
                    v3 = float64(1.5)
                }
            } else {
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 bool = v2
                    v3 += 0.0
                    _ = fi21
                    _ = v22
                }
                switch ((("\n\t" + "foo bar") + "foo bar") + ("hello" + ("hello" + ("" + ("a" + ""))))) {
                case "foo bar":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        switch ("" + "a") {
                        case "":
                            v1 = ((3.14159 / float64(1e10)) / 3.14159)
                        case "a":
                            switch -822 {
                            case 128:
                                v24 := GenS2{F0: 42}
                                _ = v24
                            case 0:
                                var v25 int32 = genFn1(-9223372036854775808)
                                _ = v25
                            case 32767:
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    v27 := GenS1{F0: 844}
                                    _ = fi26
                                    _ = v27
                                }
                            }
                        case "hello":
                            v3 += v3
                        }
                        _ = fi23
                    }
                default:
                    fmt.Println(v3, v10)
                }
            }
            _ = v10
        }
    }
    v2 = (111 >= 123)
    if len(v9) > 0 {
        v9[0] = 827
    }
    var v28 string = ("hello" + "a")
    fmt.Println(v1, v3)
    var v29 int8 = -128
    v30 := GenS1{F0: 1}
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 int8 = v29
        fmt.Println(v1)
        _ = fi31
        _ = v32
    }
    if (genFn1(1) == 0) {
        v9 = append(v9, -9223372036854775808)
    }
    v33 := GenS2{F0: -914}
    _ = len(v9)
    switch -2147483648 {
    case 902:
        v29 += v29
    case -2147483648:
        fmt.Println(v29, v3, v2)
    default:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v9
}
