package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
func genFn0(p0 int8) string {
    v7 := []int{-173, 0, 32767}
    var v8 float64 = float64(((0.0 / float32(3.14159)) - ((3.14159 - -2.25) * ((-2.25 - 0.0) * -2.25))))
    _ = p0
    _ = v7
    _ = v8
    return ("" + "a")
}

func main() {
    var v1 string = (("a" + "") + "foo bar")
    var v2 int32 = 786
    var v3 int = 32768
    var v4 string = (("hello" + v1) + v1)
    var v5 bool = false
    v6 := []uint32{66, 760, 453, 0}
    for fi9 := 0; fi9 < 13; fi9++ {
        if v5 {
            var v10 int64 = 392
            _ = v10
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS1{F0: (1 == 739), F1: v5, F2: ((v2 & (v2 & (v2 & v2))) > v2)}
                v13 := GenS1{F0: (v5 && (v5 && false)), F1: (v5 || true), F2: (v2 >= v2)}
                v1 = "a"
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        if (0 >= 758) {
            switch ((v1 + "") + ("hello" + "a")) {
            case "":
                switch v4 {
                case "":
                    v3 = (fi9 % ((v3) | 1))
                case "a":
                    if (!(true && (v5 && (float64(-2.25) < float64((0.1 - 0.0)))))) {
                        var v14 uint8 = 0
                        v15 := GenS1{F0: true, F1: (3.14159 > (0.0 + 1.5)), F2: true}
                        switch 751 {
                        case 4294967295:
                            for fi16 := 0; fi16 < 1; fi16++ {
                                var v17 uint = 231
                                v2 += 1
                                switch fi9 {
                                case 114:
                                    var v18 string = v4
                                    _ = v18
                                case 4294967295:
                                    var v19 uint64 = 42
                                    _ = v19
                                }
                                v20 := GenS1{F0: (false && v5), F1: (v2 < (v2 ^ v2)), F2: (1 <= fi16)}
                                _ = fi16
                                _ = v17
                                _ = v20
                            }
                        case 32768:
                            if len(v6) > 0 {
                                v6[0] = 2
                            }
                        default:
                            if len(v6) > 0 {
                                v6[0] = 515
                            }
                        }
                        _ = v14
                        _ = v15
                    }
                case "hello":
                    v4 = v4
                default:
                    for fi21 := 0; fi21 < 11; fi21++ {
                        switch v3 {
                        case 128:
                            v22 := GenS1{F0: v5, F1: v5, F2: false}
                            _ = v22
                        default:
                            var v23 int16 = -1
                            var v24 int8 = 0
                            switch "" {
                            case "\n\t":
                                var v25 uint32 = 801
                                _ = v25
                            default:
                                var v26 string = (v1 + ("\n\t" + v1))
                                _ = v26
                            }
                            _ = v23
                            _ = v24
                        }
                        fmt.Println(fi21)
                        fmt.Println(fi21, v2, v1)
                        _ = fi21
                    }
                }
            default:
                var v27 int64 = 42
                _ = v27
            }
            v28 := GenS1{F0: true, F1: false, F2: false}
            if len(v6) > 0 {
                v6[0] = 949
            }
            _ = v28
        }
        _ = fi9
    }
    v29 := GenS1{F0: v5, F1: true, F2: v5}
    fmt.Println(v1)
    var v30 float64 = 0.1
    fmt.Println(v4)
    if ("hello" > "hello") {
        v1 = (("" + v1) + (v4 + v1))
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS1{F0: v5, F1: ((v5 || v5) || (v5 || true)), F2: (false && ((1e10 <= float32(v30)) || (78 != 100)))}
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 int16 = -32768
                v35 := GenS1{F0: (false && (1e10 != (0.1 + ((3.14159 / (((0.1 + (1e10 - (1.5 + (((0.0 - 3.14159) - (3.14159 / 1.0)) / (-2.25 + 0.0))))) * float32(float64(0.1))) * 1.5)) / 3.14159)))), F1: (!v5), F2: v5}
                v36 := GenS1{F0: v5, F1: true, F2: true}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi31
            _ = v32
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if v5 {
            v4 = "\n\t"
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v37 int32 = v2
    v38 := GenS1{F0: true, F1: (-1 != v37), F2: false}
    switch ("" + v1) {
    case "a":
        v30 = v30
    case "\n\t":
        fmt.Println(v3, v37)
    default:
        var v39 int64 = 42
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
