package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 bool, p1 uint8) uint32 {
    v7 := []int{128}
    var v8 uint8 = 27
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 int = 797
    var v2 bool = ((v1 - -9223372036854775808) >= v1)
    v3 := make(map[uint32]int)
    var v4 string = ((((("foo bar" + "a") + "\n\t") + (("" + "hello") + "a")) + ("foo bar" + "hello")) + (("" + (("foo bar" + "a") + ("foo bar" + "\n\t"))) + "hello"))
    v5 := make(map[uint32]bool)
    var v6 float64 = float64(float32(-2.25))
    if v2 {
        var v9 uint = 1
        var v10 uint = (v9 + 149)
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := v5[715]
            _ = fi11
            _ = v12
        }
        var v13 float32 = (((1.5 * -2.25) + ((float32((1.5 - 1e10)) / ((0.0 + -2.25) + float32(v6))) * (1e10 - 1e10))) - 1e10)
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int16 = 32767
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v1 -= v1
    var v16 string = ""
    var v17 bool = v2
    var v18 uint8 = 22
    if (-1 != 1) {
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                switch -9223372036854775808 {
                case 127:
                    fmt.Println(v17)
                case 968:
                    v18 += 1
                case 32767:
                    v3[638] = fi19
                }
                _ = fi20
            }
            if false {
                v21 := []int{}
                var v22 uint8 = 1
                switch v4 {
                case "foo bar":
                    v22 = (^(v22))
                case "":
                    for fi23 := 0; fi23 < 12; fi23++ {
                        var v24 float32 = 1.5
                        switch (v16 + ((v4 + "hello") + "a")) {
                        case "a":
                            if v2 {
                                switch "" {
                                case "foo bar":
                                    v25 := GenS1{F0: -128, F1: fi19, F2: 42}
                                    _ = v25
                                }
                                v26 := v3[genFn0(true, (v22 >> ((1) & 63)))]
                                v1 = v1
                                _ = v26
                            } else {
                                v18 = v22
                            }
                            if v2 {
                                var v27 uint32 = 2
                                if (!(847 == 0)) {
                                    var v28 uint8 = v18
                                    v24 = ((0.1 + 1e10) + v24)
                                    if (65535 == 359) {
                                        if true {
                                            v29 := GenS1{F0: -111, F1: fi23, F2: -89}
                                            _ = v29
                                        }
                                    }
                                    _ = v28
                                }
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                _ = v27
                            } else {
                                v18 += v18
                                v16 = "foo bar"
                                var v30 float64 = v6
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                _ = v30
                            }
                            var v31 float64 = (-2.25 * v6)
                            _ = v31
                        }
                        v3[65535] = -368
                        _ = fi23
                        _ = v24
                    }
                default:
                    if (!((true || (!(v17 || (4294967295 <= 1)))) || v2)) {
                        for fi32 := 0; fi32 < 3; fi32++ {
                            if len(v21) > 0 {
                                v21[0] = 375
                            }
                            _ = fi32
                        }
                        switch (-(v1)) {
                        case 127:
                            delete(v5, 2)
                        }
                    }
                }
                _ = v21
                _ = v22
            }
            var v33 uint = 1
            _ = fi19
            _ = v33
        }
    }
    var v34 float32 = float32(v6)
    var v35 uint16 = 807
    var v36 uint16 = 642
    v37 := GenS2{F0: 1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
