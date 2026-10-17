package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0() int {
    var v5 int32 = 42
    var v6 int64 = -1
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 2147483648
    var v3 int32 = 0
    var v4 int32 = -1
    v3 -= (v3 & v3)
    var v7 int16 = -321
    v8 := GenS1{F0: 1}
    v1 += 107
    v7 -= v1
    v9 := []int{2147483647, 765, -448}
    for fi10 := 0; fi10 < 3; fi10++ {
        if ((v4 + v3) >= v4) {
            for fi11 := 0; fi11 < 17; fi11++ {
                v4 -= v4
                fmt.Println(v4)
                if (306 > 0) {
                    if false {
                        v3 -= 0
                    } else {
                        var v12 string = "hello"
                        fmt.Println(v12, fi11, v4)
                        var v13 uint16 = 115
                        _ = v12
                        _ = v13
                    }
                    if false {
                        if (!true) {
                            _ = len(v9)
                            v3 = int32(fi11)
                        }
                    } else {
                        var v14 uint64 = 0
                        _ = v14
                    }
                    var v15 float64 = (-2.25 * 1e10)
                    _ = v15
                } else {
                    if false {
                        switch ("a" + "") {
                        case "":
                            v16 := GenS1{F0: 1}
                            _ = v16
                        case "hello":
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        case "\n\t":
                            v2 -= (^(v2))
                        }
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    _ = len(v9)
                }
                _ = fi11
            }
        } else {
            v17 := GenS1{F0: 819}
            if ("" != "") {
                switch (fi10 & 42) {
                case 9223372036854775807:
                    if (2147483648 < (^(v2))) {
                        if (!false) {
                            v18 := GenS2{F0: (v7 / ((275) | 1)), F1: false, F2: 0}
                            for fi19 := 0; fi19 < 12; fi19++ {
                                var v20 bool = ((0.0 + 3.14159) == ((((0.0 + (1.5 - ((((1e10 + 0.0) / 0.1) + ((1.5 / 1.5) - 3.14159)) + 1e10))) * -2.25) - 1.5) + (((1.5 * 1e10) + 0.1) - (float64((-2.25 + -2.25)) * 3.14159))))
                                v20 = (((^(fi10)) > (fi19 ^ (fi10 % ((-9223372036854775808) | 1)))) && v20)
                                if len(v9) > 0 {
                                    v9[0] = 1
                                }
                                var v21 int32 = v4
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            v2 -= (v2 << ((24) & 63))
                            _ = v18
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                    }
                }
            } else {
                var v22 bool = (1 > 9223372036854775807)
                _ = v22
            }
            _ = v17
        }
        _ = fi10
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if ((-2.25 > 1.5) && true) {
        var v23 int = -958
        for fi24 := 0; fi24 < 2; fi24++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            switch (v23 & (fi24 - v23)) {
            case 4294967295:
                var v25 uint32 = (v2 << ((2) & 63))
                _ = v25
            case -1:
                for fi26 := 0; fi26 < 10; fi26++ {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v27 int8 = 42
                    switch genFn0() {
                    case 0:
                        if (!(v7 > 32767)) {
                            var v28 float32 = -2.25
                            _ = v28
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            v9 = append(v9, (^(v23)))
                        }
                    case -2147483648:
                        v27 += (v27 & 0)
                    }
                    _ = fi26
                    _ = v27
                }
            case 0:
                switch "a" {
                case "":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                default:
                    var v29 int64 = 1
                    _ = v29
                }
            default:
                switch v23 {
                case -359:
                    var v30 bool = false
                    _ = v30
                case -756:
                    var v31 uint32 = (v2 << ((706) & 63))
                    _ = v31
                case 183:
                    v32 := GenS1{F0: 255}
                    _ = v32
                }
            }
            _ = fi24
        }
        v33 := GenS2{F0: (-(v7)), F1: (false && (!true)), F2: 42}
        _ = v23
        _ = v33
    }
    var v34 string = "\n\t"
    v3 -= 2147483647
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
