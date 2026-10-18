package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0() uint16 {
    var v5 uint16 = 32767
    v5 = v5
    var v6 uint32 = 673
    var v7 uint64 = 906
    fmt.Println(v6, v5)
    _ = v5
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint8 = 110
    v2 := make(map[uint32]bool)
    var v3 uint64 = 2
    var v4 uint32 = 1
    var v8 bool = true
    v2[v4] = false
    v3 -= v3
    v9 := []int{-1, 128, -2147483648, 128}
    fmt.Println(v8)
    if v8 {
        switch "\n\t" {
        case "hello":
            switch 572 {
            case 9223372036854775807:
                switch 529 {
                case 32768:
                    if v8 {
                        fmt.Println(v1, v1)
                        var v10 float64 = (float64((0.1 * 1e10)) + 1.5)
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            var v12 bool = (!true)
                            var v13 int32 = 42
                            var v14 uint64 = 46
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        var v15 uint = 1000
                        _ = v15
                    }
                case -9223372036854775808:
                    v4 = 824
                case -730:
                    v9 = append(v9, -985)
                }
            case 817:
                v16 := GenS1{F0: 764, F1: 831}
                _ = v16
            default:
                v17 := GenS2{F0: 1, F1: -128}
                _ = v17
            }
        case "\n\t":
            var v18 string = ""
            _ = v18
        default:
            if (251 <= 698) {
                var v19 int = 0
                for fi20 := 0; fi20 < 9; fi20++ {
                    if (-1 == 42) {
                        var v21 uint = 714
                        _ = v21
                    } else {
                        _ = len(v9)
                        var v22 uint8 = v1
                        _ = v22
                    }
                    var v23 float32 = 1.5
                    _ = fi20
                    _ = v23
                }
                for fi24 := 0; fi24 < 14; fi24++ {
                    if (-9223372036854775808 < 9223372036854775807) {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            var v26 string = (("foo bar" + "hello") + "")
                            v3 += v3
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        v27 := GenS2{F0: (-(fi24)), F1: 89}
                        if len(v9) > 0 {
                            v9[0] = (^(v19))
                        }
                        for fi28 := 0; fi28 < 16; fi28++ {
                            if true {
                                v29 := GenS1{F0: 1, F1: genFn0()}
                                var v30 uint32 = 526
                                _ = v29
                                _ = v30
                            }
                            if (!v8) {
                                var v31 uint64 = (^(v3))
                                var v32 uint32 = v4
                                switch ("hello" + "hello") {
                                case "\n\t":
                                    v33 := v2[v32]
                                    _ = v33
                                case "foo bar":
                                    fmt.Println(v8, v1)
                                case "a":
                                    delete(v2, (v4 % ((2147483648) | 1)))
                                default:
                                    if v8 {
                                        fmt.Println(v1, fi28, v32)
                                    } else {
                                        v19 = v19
                                        if v8 {
                                            v8 = v8
                                        } else {
                                            var v34 uint8 = v1
                                            _ = v34
                                        }
                                    }
                                }
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(v3)
                            }
                            var v35 bool = (!v8)
                            _ = fi28
                            _ = v35
                        }
                        _ = v27
                    }
                    fmt.Println(v3, v3, fi24)
                    _ = fi24
                }
                _ = v19
            }
        }
    } else {
        v2[65535] = v8
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 16; fi37++ {
            switch (fi37 ^ fi37) {
            case -9223372036854775808:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v38 float64 = (-2.25 + -2.25)
                v39 := GenS1{F0: 4294967295, F1: 489}
                _ = v38
                _ = v39
            }
            _ = fi37
        }
        v8 = ((v3 / ((v3) | 1)) != v3)
        _ = len(v2)
        _ = fi36
    }
    fmt.Println(v3)
    fmt.Println(v4, v4, v4)
    for fi40 := 0; fi40 < 8; fi40++ {
        if (-1 < 0) {
            v4 = v4
            var v41 uint16 = genFn0()
            v9 = append(v9, (fi40 ^ (^(fi40))))
            _ = v41
        }
        v42 := GenS1{F0: 40, F1: 255}
        _ = fi40
        _ = v42
    }
    fmt.Println(v8, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
