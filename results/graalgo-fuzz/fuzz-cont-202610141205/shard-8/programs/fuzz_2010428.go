package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0(p0 string, p1 int, p2 uint8) uint {
    v4 := GenS1{F0: 2, F1: (p2 + p2), F2: 0}
    v5 := GenS1{F0: 1, F1: p2, F2: 32767}
    p0 = (p0 + p0)
    var v6 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint) int32 {
    p0 = p0
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            p0 += p0
            var v8 int = -1
            _ = fi7
            _ = v8
        }
        var v9 int64 = 0
        var v10 uint8 = 128
        _ = v10
        _ = v9
    }
    switch 128 {
    case 1:
        var v11 float32 = 0.1
        _ = v11
    default:
        var v12 bool = (true && true)
        _ = v12
    }
    var v13 int = 32768
    var v14 uint32 = 2147483648
    _ = p0
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 int8 = -20
    var v2 float64 = float64(-2.25)
    var v3 int32 = -183
    v15 := GenS1{F0: 1, F1: 2, F2: -996}
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := []int{(fi16 | (^(fi17))), -1}
            _ = fi17
            _ = v18
        }
        v1 -= (v1 | v1)
        _ = fi16
    }
    switch -9223372036854775808 {
    case -653:
        v19 := GenS1{F0: genFn0((("a" + (("foo bar" + "\n\t") + "hello")) + "\n\t"), 9223372036854775807, 1), F1: 42, F2: -921}
        _ = v19
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            if (42 > 1) {
                v2 += 1e10
                switch fi20 {
                case -2147483648:
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if ((-2.25 * (1.5 - (0.0 + -2.25))) > 0.0) {
                            var v22 string = ""
                            _ = v22
                        } else {
                            fmt.Println(fi20, v2, v3)
                            switch ("foo bar" + "foo bar") {
                            case "":
                                var v23 int32 = (v3 << ((42) & 63))
                                _ = v23
                            }
                        }
                        var v24 int8 = (v1 - 127)
                        var v25 uint32 = 22
                        _ = fi21
                        _ = v24
                        _ = v25
                    }
                case 0:
                    v26 := GenS1{F0: 65535, F1: 0, F2: 1}
                    _ = v26
                default:
                    var v27 int64 = 42
                    _ = v27
                }
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := []int{}
                    if true {
                        v30 := GenS1{F0: 4294967295, F1: 2, F2: 1}
                        fmt.Println(v1)
                        v1 = v1
                        _ = v30
                    } else {
                        if true {
                            v1 += 88
                        }
                        v29 = append(v29, (fi28 - (fi28 + -770)))
                        v1 = (v1 & v1)
                        v2 -= v2
                    }
                    _ = len(v29)
                    var v31 float32 = (-2.25 - 0.1)
                    v29 = append(v29, fi28)
                    v32 := GenS1{F0: 42, F1: 255, F2: -1}
                    if true {
                        v3 = v3
                        var v33 int = -633
                        var v34 float64 = 1e10
                        _ = v33
                        _ = v34
                    }
                    v2 += (v2 / v2)
                    _ = fi28
                    _ = v29
                    _ = v31
                    _ = v32
                }
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v1, fi20, v1)
            case "":
                v2 = (v2 * v2)
            }
            _ = fi20
        }
    }
    v35 := GenS1{F0: genFn0((((((("\n\t" + ("" + ("hello" + ""))) + "hello") + (((("" + "\n\t") + "hello") + ("foo bar" + "foo bar")) + ("foo bar" + ("a" + "\n\t")))) + ("a" + "\n\t")) + "hello") + "a"), -1, 42), F1: 124, F2: -658}
    v36 := GenS1{F0: 1, F1: 36, F2: 1}
    fmt.Println(v3, v2, v2)
    var v37 int32 = (v3 | v3)
    v38 := []int{}
    var v39 uint16 = 255
    if (18446744073709551615 >= 2) {
        v37 = v3
        v1 = (v1 & 0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
