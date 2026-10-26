package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
func genFn0() bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 bool = (fi4 >= (fi4 & fi4))
        for fi6 := 0; fi6 < 17; fi6++ {
            v5 = (42 <= 42)
            _ = fi6
        }
        if v5 {
            var v7 bool = false
            var v8 int = (fi4 + (fi4 / ((32767) | 1)))
            _ = v7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    v9 := GenS1{F0: -1, F1: (((1.5 / 1.0) + (float64((((-2.25 + 0.1) / 0.1) * (3.14159 - -2.25))) * (float64(1.5) / 0.1))) > 3.14159), F2: 0}
    var v10 float64 = 3.14159
    var v11 float64 = 1.5
    _ = v10
    _ = v11
    _ = v9
    return false
}
func genFn1() int {
    var v12 uint = 1
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(v12)
        v14 := GenS1{F0: 417, F1: (1 <= 1), F2: 112}
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 741
    for fi16 := 0; fi16 < 3; fi16++ {
        if false {
            v15 = (v15 & v15)
            switch 0 {
            case 127:
                v12 = v12
            case 32767:
                var v17 int32 = -1
                _ = v17
            case -974:
                var v18 int16 = 318
                _ = v18
            default:
                var v19 int8 = -1
                _ = v19
            }
        }
        v20 := []int{32768, -247, (fi16 >> ((v12) & 63)), (fi16 << (((v12 / ((v12) | 1))) & 63))}
        var v21 uint8 = 93
        _ = fi16
        _ = v20
        _ = v21
    }
    v15 = 65535
    _ = v12
    _ = v15
    return -1
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 2
    var v3 uint = 255
    if (-2.25 != float64(-2.25)) {
        var v22 uint8 = 148
        if false {
            v23 := []int{-1, 0, 32768}
            switch "\n\t" {
            case "\n\t":
                v2 -= v2
            case "foo bar":
                if (315 != 0) {
                    v2 += 1
                    v22 = (^(v22))
                    switch "a" {
                    case "a":
                        v1 += v1
                    case "\n\t":
                        if true {
                            for fi24 := 0; fi24 < 15; fi24++ {
                                fmt.Println(v3, v3)
                                _ = fi24
                            }
                        }
                    case "foo bar":
                        switch 32767 {
                        case -1:
                            switch 724 {
                            case -997:
                                switch "a" {
                                case "foo bar":
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                case "hello":
                                    switch 4294967295 {
                                    case 127:
                                        var v25 uint = v3
                                        _ = v25
                                    case 902:
                                        _ = len(v23)
                                    default:
                                        fmt.Println(v3, v1)
                                    }
                                default:
                                    var v26 int = genFn1()
                                    _ = v26
                                }
                            case 4294967295:
                                v27 := GenS1{F0: 32768, F1: true, F2: 695}
                                _ = v27
                            default:
                                var v28 bool = false
                                _ = v28
                            }
                        }
                    }
                } else {
                    v23 = append(v23, 2147483647)
                    for fi29 := 0; fi29 < 3; fi29++ {
                        fmt.Println(v2, fi29, v2)
                        _ = fi29
                    }
                }
            case "":
                v22 -= 0
            default:
                v3 = 255
            }
            var v30 uint = v3
            _ = v23
            _ = v30
        } else {
            if true {
                fmt.Println(v3, v22)
            } else {
                var v31 uint16 = 333
                _ = v31
            }
        }
        _ = v22
    } else {
        v32 := GenS1{F0: int(v3), F1: (true || ("a" < "hello")), F2: 2}
        _ = v32
    }
    var v33 string = "hello"
    if false {
        if true {
            v34 := []int{611, 389}
            var v35 int64 = int64(v3)
            _ = v34
            _ = v35
        } else {
            var v36 uint8 = 20
            v2 -= (v2 >> ((v3) & 63))
            _ = v36
        }
    }
    var v37 int8 = -1
    v3 = v3
    v33 = v33
    fmt.Println(v33)
    v33 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
}
