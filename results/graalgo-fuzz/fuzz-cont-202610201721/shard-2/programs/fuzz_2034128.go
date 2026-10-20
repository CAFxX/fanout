package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32, p1 int32) int16 {
    var v4 float64 = 1e10
    v5 := GenS1{F0: 2, F1: ((3.14159 - (0.1 * (1.5 * 0.1))) == 0.1)}
    v6 := []int{}
    p1 += -2147483648
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = (0.1 * 0.0)
    v3 := []int{1, 975, 0}
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS2{F0: -128}
        v9 := GenS1{F0: 0, F1: false}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = -2.25
    var v11 int64 = 985
    var v12 uint32 = 2
    fmt.Println(v12, v12)
    v11 -= 1
    v2 -= (v2 / float32(v10))
    _ = len(v3)
    fmt.Println(v11, v10)
    v13 := GenS2{F0: 42}
    if false {
        v11 = (v11 - v11)
        switch 32768 {
        case -2147483648:
            for fi14 := 0; fi14 < 13; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if ((v10 * v10) == 0.1) {
                        if (!(true && true)) {
                            fmt.Println(v12)
                            v10 += v10
                            v10 += -2.25
                        }
                        if (161 > 127) {
                            var v16 bool = (v2 <= v2)
                            _ = v16
                        } else {
                            fmt.Println(fi15, fi14, v11)
                        }
                        var v17 int = (fi14 - fi15)
                        _ = v17
                    }
                    var v18 bool = (((v2 / v2) - v2) < (v2 * 0.1))
                    _ = fi15
                    _ = v18
                }
                var v19 int32 = -2147483648
                _ = fi14
                _ = v19
            }
        case -80:
            var v20 uint16 = 393
            _ = v20
        case 42:
            var v21 int = 2147483647
            _ = v21
        default:
            var v22 string = (("hello" + "hello") + "\n\t")
            _ = v22
        }
        if (v11 == v11) {
            var v23 bool = (("a" + "") != "\n\t")
            _ = v23
        } else {
            if false {
                switch (("foo bar" + ("hello" + ("a" + (("foo bar" + "") + ((("" + "") + "foo bar") + ("" + ("\n\t" + ("\n\t" + ("" + (("\n\t" + "a") + "hello")))))))))) + "") {
                case "a":
                    var v24 int = 127
                    _ = v24
                case "foo bar":
                    for fi25 := 0; fi25 < 3; fi25++ {
                        switch (fi25 >> ((65535) & 63)) {
                        case 793:
                            v2 -= (float32(v10) * v2)
                        case 216:
                            v11 -= -1
                        case 4294967295:
                            fmt.Println(v2, v12, v2)
                        default:
                            v26 := GenS2{F0: -128}
                            _ = v26
                        }
                        if true {
                            v27 := GenS2{F0: 103}
                            v11 = v11
                            v28 := GenS1{F0: v12, F1: false}
                            _ = v27
                            _ = v28
                        } else {
                            v29 := GenS2{F0: 100}
                            var v30 float64 = 3.14159
                            v30 = (-2.25 + v30)
                            _ = v29
                            _ = v30
                        }
                        v12 += 42
                        _ = fi25
                    }
                case "\n\t":
                    var v31 int8 = 127
                    _ = v31
                default:
                    v11 += v11
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    v12 -= (^(v12))
                    v12 = v12
                    _ = fi32
                }
            }
            _ = len(v3)
        }
    }
    var v33 uint16 = 0
    var v34 int64 = 42
    if false {
        var v35 int32 = 0
        var v36 uint32 = uint32(v33)
        switch 560 {
        case 792:
            switch 99 {
            case 127:
                var v37 int32 = v35
                _ = v37
            default:
                var v38 int16 = 750
                _ = v38
            }
        case 2147483647:
            v3 = append(v3, 127)
        }
        _ = v35
        _ = v36
    } else {
        var v39 int32 = 0
        _ = v39
    }
    v10 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
