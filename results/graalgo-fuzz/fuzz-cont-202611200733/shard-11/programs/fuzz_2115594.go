package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 int32 = -2147483648
    var v4 int16 = 32767
    fmt.Println(v4, v4)
    if true {
        if false {
            if ((v4 | (v4 % (((v4 >> ((1) & 63))) | 1))) > v4) {
                for fi5 := 0; fi5 < 7; fi5++ {
                    for fi6 := 0; fi6 < 6; fi6++ {
                        switch 2147483647 {
                        case 9223372036854775807:
                            var v7 string = ("\n\t" + "\n\t")
                            if ((1e10 * (0.0 + float32(1e10))) >= (((0.1 * 0.1) + 1.5) + 0.1)) {
                                v7 = (("" + ("" + "a")) + v7)
                                v3 = (v3 & (v3 / (((v3 / ((v3) | 1))) | 1)))
                            } else {
                                v3 += (v3 % ((v3) | 1))
                                var v8 uint8 = 0
                                _ = v8
                            }
                            v1[v7] = fi6
                            _ = v7
                        }
                        v9 := GenS1{F0: true, F1: (false && (-2147483648 >= fi5)), F2: 255}
                        _ = fi6
                        _ = v9
                    }
                    _ = fi5
                }
                v10 := GenS1{F0: (!(18446744073709551615 == 679)), F1: (!(128 == 42)), F2: 684}
                _ = v10
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 string = "foo bar"
                v4 -= (v4 & -1)
                _ = fi11
                _ = v12
            }
            v13 := []int{}
            _ = v13
        } else {
            v3 += 914
        }
    }
    v14 := GenS1{F0: false, F1: true, F2: 65535}
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 float32 = 0.0
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 int64 = 625
            fmt.Println(fi17, v18)
            v19 := []int{(fi17 % ((-166) | 1)), (fi15 >> ((255) & 63)), fi17}
            _ = fi17
            _ = v18
            _ = v19
        }
        if ("hello" < "\n\t") {
            var v20 int16 = (^(v4))
            if ("foo bar" == "\n\t") {
                var v21 int16 = (^(v4))
                _ = v21
            }
            switch "hello" {
            case "\n\t":
                v20 += (v20 - -265)
            case "":
                switch "foo bar" {
                case "a":
                    if (-2147483648 <= (v3 & (v3 << ((2) & 63)))) {
                        if true {
                            var v22 uint16 = 255
                            var v23 string = "foo bar"
                            switch ("\n\t" + ("foo bar" + v23)) {
                            case "a":
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v25 uint8 = 42
                                        _ = v25
                                    }
                                    delete(v1, v23)
                                    _ = fi24
                                }
                            case "foo bar":
                                var v26 int32 = (v3 | (v3 / ((v3) | 1)))
                                _ = v26
                            default:
                                switch (^(fi15)) {
                                case -1:
                                    var v27 float64 = (float64((v16 + 1e10)) - 3.14159)
                                    _ = v27
                                case 4294967295:
                                    for fi28 := 0; fi28 < 15; fi28++ {
                                        v29 := []int{0}
                                        v4 = (v4 ^ (v20 % ((-32768) | 1)))
                                        delete(v2, 67)
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                            }
                            _ = v22
                            _ = v23
                        } else {
                            v3 += v3
                        }
                        v30 := []int{(fi15 / (((^(fi15))) | 1)), fi15, (fi15 << ((255) & 63))}
                        _ = v30
                    } else {
                        if true {
                            var v31 uint32 = 0
                            v32 := GenS1{F0: (32767 <= (v4 << ((327) & 63))), F1: false, F2: 42}
                            v33 := []int{fi15}
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                if ((true || (v3 == v3)) || false) {
                                    v34 := []int{fi15}
                                    var v35 int = fi15
                                    _ = v34
                                    _ = v35
                                }
                            case "\n\t":
                                v1["a"] = fi15
                            }
                        }
                        _ = len(v1)
                    }
                default:
                    v36 := []int{703}
                    _ = v36
                }
            case "hello":
                _ = len(v2)
            }
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        v38 := []int{213, fi37, (^(fi37))}
        var v39 int16 = v4
        _ = fi37
        _ = v38
        _ = v39
    }
    delete(v2, 42)
    v2[1] = true
    v4 = v4
    delete(v2, 332)
    v3 = (v3 - v3)
    v40 := GenS1{F0: true, F1: true, F2: 355}
    v41 := GenS1{F0: false, F1: (3.14159 != 0.1), F2: 1}
    if false {
        v42 := []int{464, -26, 32767, -743}
        _ = v42
    }
    v3 -= 1
    v43 := GenS2{F0: 100, F1: -14}
    for fi44 := 0; fi44 < 16; fi44++ {
        v4 -= (^(v4))
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
}
