package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 142
    var v3 uint8 = (v2 | (v2 - 2))
    var v4 uint = 903
    fmt.Println(v4, v3, v3)
    v5 := GenS1{F0: 0, F1: 540, F2: v4}
    v6 := GenS1{F0: 2147483648, F1: 2, F2: (^(v4))}
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            var v8 int32 = 2147483647
            fmt.Println(fi7, v2)
            _ = v8
        } else {
            if (true || false) {
                var v9 bool = (fi7 >= fi7)
                _ = v9
            } else {
                fmt.Println(v3, v4, v3)
                switch "a" {
                case "\n\t":
                    v10 := []int{fi7, fi7}
                    _ = v10
                default:
                    if true {
                        var v11 int8 = 42
                        v12 := GenS1{F0: 5, F1: 641, F2: 199}
                        switch ("hello" + "hello") {
                        case "\n\t":
                            delete(v1, (-(fi7)))
                        case "":
                            var v13 int64 = 0
                            _ = v13
                        default:
                            v14 := GenS2{F0: 1}
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                    } else {
                        if false {
                            fmt.Println(fi7, v2, v3)
                            v15 := GenS1{F0: 4294967295, F1: uint16(v3), F2: 255}
                            var v16 int16 = 1
                            _ = v15
                            _ = v16
                        }
                    }
                }
                if true {
                    v17 := GenS1{F0: 65535, F1: uint16(v4), F2: v4}
                    var v18 uint32 = 842
                    for fi19 := 0; fi19 < 7; fi19++ {
                        var v20 int = -2147483648
                        if true {
                            v2 += 128
                        } else {
                            var v21 uint8 = (v3 ^ (v2 << (((v4 ^ (v4 / ((65535) | 1)))) & 63)))
                            switch ("\n\t" + "hello") {
                            case "hello":
                                var v22 int = (-(v20))
                                _ = v22
                            case "\n\t":
                                if ((^(v2)) >= (v3 << (((^(v4))) & 63))) {
                                    fmt.Println(v21, v18)
                                    fmt.Println(v3, v21, fi7)
                                    v23 := []int{}
                                    _ = v23
                                }
                            case "a":
                                if true {
                                    v24 := []int{}
                                    v25 := GenS2{F0: 0}
                                    _ = v24
                                    _ = v25
                                }
                            }
                            _ = v21
                        }
                        if ((v3 - v2) != (v3 & (v3 % ((65) | 1)))) {
                            var v26 int = (v20 / (((fi19 << ((v4) & 63))) | 1))
                            v27 := GenS2{F0: 0}
                            var v28 float64 = (float64(1e10) * -2.25)
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = v17
                    _ = v18
                } else {
                    v29 := GenS2{F0: -181}
                    v3 = 1
                    v4 = (v4 % (((v4 % ((v4) | 1))) | 1))
                    _ = v29
                }
            }
        }
        _ = fi7
    }
    var v30 float32 = 3.14159
    var v31 uint8 = (v3 - 228)
    v32 := GenS1{F0: 32, F1: 2, F2: v4}
    v31 -= (v2 ^ (v3 % ((v31) | 1)))
    var v33 uint = 18446744073709551615
    switch 158 {
    case 1:
        v1[127] = (42 == 1)
    case 32768:
        if true {
            v34 := GenS1{F0: 2, F1: 1, F2: 544}
            _ = v34
        } else {
            var v35 uint32 = 332
            v31 = v31
            v2 = (v3 ^ (v31 >> (((^(v4))) & 63)))
            _ = v35
        }
    }
    if (!(0.0 >= 0.1)) {
        var v36 uint32 = 1
        v37 := GenS1{F0: v36, F1: 42, F2: uint(v3)}
        if false {
            if false {
                var v38 uint32 = v36
                for fi39 := 0; fi39 < 14; fi39++ {
                    v2 -= v3
                    var v40 int16 = -1
                    _ = fi39
                    _ = v40
                }
                v38 = (^(v36))
                _ = v38
            }
        }
        _ = v36
        _ = v37
    } else {
        v41 := GenS1{F0: 49, F1: 151, F2: (v33 >> ((v4) & 63))}
        _ = v41
    }
    fmt.Println(v30)
    v3 = (^(v3))
    var v42 uint64 = 211
    var v43 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
