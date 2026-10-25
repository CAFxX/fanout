package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 string }
type GenS3 struct { F0 uint8; F1 uint; F2 string }
type GenS4 struct { F0 int8 }
func genFn0(p0 int8, p1 uint8) bool {
    if false {
        var v6 uint = 368
        _ = v6
    }
    p1 -= p1
    _ = p0
    _ = p1
    return (!(128 >= 9223372036854775807))
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = 1.5
    var v3 uint8 = 128
    var v4 string = ("" + "\n\t")
    var v5 float32 = 0.1
    v4 = "foo bar"
    fmt.Println(v4, v2)
    var v7 int8 = 127
    var v8 int64 = 9223372036854775807
    var v9 uint32 = 4294967295
    v10 := []int{2147483647, 212, -391, 128}
    if false {
        switch 338 {
        case 32768:
            v5 -= v5
        default:
            v11 := GenS1{F0: (^(v9))}
            _ = v11
        }
        var v12 bool = (v4 < v4)
        _ = v12
    }
    switch int(v8) {
    case 0:
        if (127 > 780) {
            switch ((v4 + v4) + v4) {
            case "a":
                if (!(((!(float64(float32(v2)) > (v2 * -2.25))) || (!((("a" + v4) + "\n\t") != (v4 + "a")))) || false)) {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        _ = len(v10)
                        if (2 != 32767) {
                            v2 = (v2 * float64(float32(v2)))
                            for fi14 := 0; fi14 < 9; fi14++ {
                                v4 = v4
                                fmt.Println(v1)
                                _ = fi14
                            }
                        } else {
                            v15 := GenS2{F0: (^(v7)), F1: 4294967295, F2: "a"}
                            fmt.Println(v4)
                            _ = v15
                        }
                        var v16 uint32 = (v9 >> ((18446744073709551615) & 63))
                        _ = fi13
                        _ = v16
                    }
                } else {
                    v4 = "\n\t"
                }
            }
            v4 = v4
        } else {
            if true {
                for fi17 := 0; fi17 < 18; fi17++ {
                    if false {
                        v18 := GenS2{F0: v7, F1: v1, F2: v4}
                        v1 = 2
                        var v19 uint = 242
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
                if true {
                    v4 = v4
                    if false {
                        v20 := GenS1{F0: (v9 & 1)}
                        for fi21 := 0; fi21 < 3; fi21++ {
                            v3 += (v3 / (((v3 - (^(v3)))) | 1))
                            _ = fi21
                        }
                        _ = v20
                    } else {
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v3 -= 202
                            v23 := GenS2{F0: v7, F1: v1, F2: v4}
                            v24 := GenS1{F0: 4294967295}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        if (!true) {
                            _ = len(v10)
                            for fi25 := 0; fi25 < 11; fi25++ {
                                fmt.Println(v8, v7)
                                v26 := GenS1{F0: (v9 * 2147483647)}
                                v4 = ((v4 + "\n\t") + v4)
                                _ = fi25
                                _ = v26
                            }
                            v7 = 49
                        }
                    }
                } else {
                    switch ("a" + (v4 + "a")) {
                    case "":
                        for fi27 := 0; fi27 < 5; fi27++ {
                            for fi28 := 0; fi28 < 7; fi28++ {
                                if false {
                                    v4 = (v4 + v4)
                                    if (3.14159 != v5) {
                                        fmt.Println(fi28, v9)
                                        v3 = (v3 & (^(v3)))
                                    }
                                }
                                v9 += v9
                                var v29 int64 = -9223372036854775808
                                _ = fi28
                                _ = v29
                            }
                            _ = fi27
                        }
                    default:
                        v10 = append(v10, 2147483647)
                    }
                }
            }
            v30 := GenS3{F0: (v3 >> ((0) & 63)), F1: 4294967295, F2: v4}
            _ = v30
        }
    case -104:
        v5 = v5
    case 127:
        var v31 float32 = float32((v2 + (v2 - v2)))
        _ = v31
    default:
        v32 := GenS4{F0: v7}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
