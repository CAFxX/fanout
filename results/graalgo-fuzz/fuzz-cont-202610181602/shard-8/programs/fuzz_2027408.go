package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 int16, p1 int16) uint {
    if true {
        var v6 float64 = -2.25
        _ = v6
    }
    switch -1 {
    case -1:
        var v7 uint16 = 2
        _ = v7
    case 420:
        if false {
            var v8 int32 = -2147483648
            _ = v8
        } else {
            var v9 int64 = 0
            _ = v9
        }
    case 128:
        fmt.Println(p1, p1)
    default:
        v10 := GenS1{F0: 0}
        _ = v10
    }
    p0 = p1
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 int16) int {
    fmt.Println(p0)
    var v11 int32 = 541
    fmt.Println(p0)
    _ = p0
    _ = v11
    return 127
}

func main() {
    var v1 int8 = 42
    var v2 int64 = 9223372036854775807
    var v3 int64 = (v2 * v2)
    var v4 uint = 2
    var v5 uint8 = 154
    var v12 int8 = v1
    v13 := []int{128, -1, genFn1(-779), 1}
    if true {
        if (v3 > -9223372036854775808) {
            if (9223372036854775807 <= -9223372036854775808) {
                var v14 bool = false
                v4 -= (v4 | genFn0(-283, 1000))
                _ = v14
            } else {
                fmt.Println(v12, v1)
                if len(v13) > 0 {
                    v13[0] = 452
                }
                v15 := GenS1{F0: (v2 | -1)}
                _ = v15
            }
        } else {
            if (2147483648 <= 1) {
                switch ("" + "hello") {
                case "a":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        if (true || (0 > 42)) {
                            var v17 uint16 = 32767
                            _ = v17
                        }
                        switch (("" + "") + "foo bar") {
                        case "a":
                            var v18 int64 = (^(v2))
                            _ = v18
                        case "\n\t":
                            if (-80 == (v12 % ((-128) | 1))) {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                v3 -= (v2 + v2)
                            } else {
                                fmt.Println(v3, v3, v4)
                            }
                        case "":
                            fmt.Println(fi16, v5, fi16)
                        }
                        _ = fi16
                    }
                default:
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 float64 = 3.14159
                        v21 := GenS2{F0: 4294967295, F1: ("" + (("" + "a") + (("a" + "\n\t") + (("a" + (("a" + "foo bar") + "a")) + ("\n\t" + "foo bar")))))}
                        var v22 float64 = 1.5
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
                var v23 int16 = -86
                if (((-(v1)) <= (v12 % ((42) | 1))) && true) {
                    var v24 uint32 = 65535
                    var v25 uint8 = v5
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint8 = (v5 - (v5 | (^(v5))))
                    _ = v26
                }
                _ = v23
            } else {
                v27 := GenS1{F0: 180}
                _ = v27
            }
        }
        switch ("hello" + "\n\t") {
        case "\n\t":
            var v28 uint = 592
            _ = v28
        case "foo bar":
            var v29 float64 = 1e10
            _ = v29
        case "hello":
            v30 := GenS2{F0: 368, F1: "a"}
            _ = v30
        default:
            switch "" {
            case "":
                var v31 uint64 = 0
                _ = v31
            case "foo bar":
                if true {
                    if len(v13) > 0 {
                        v13[0] = -590
                    }
                    v12 += (v12 / ((v12) | 1))
                    var v32 int8 = (v1 / ((v12) | 1))
                    _ = v32
                } else {
                    v1 += int8(v5)
                    var v33 int = 42
                    _ = v33
                }
            case "\n\t":
                v34 := GenS2{F0: 42, F1: "hello"}
                _ = v34
            }
        }
        switch 4294967295 {
        case -568:
            for fi35 := 0; fi35 < 6; fi35++ {
                fmt.Println(v2)
                fmt.Println(v5)
                var v36 uint8 = v5
                _ = fi35
                _ = v36
            }
        }
    } else {
        switch 32767 {
        case 4294967295:
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        default:
            var v37 int = 4294967295
            _ = v37
        }
        if false {
            v5 -= 42
            v1 = 127
            for fi38 := 0; fi38 < 10; fi38++ {
                _ = len(v13)
                _ = fi38
            }
        } else {
            var v39 int64 = (-(v2))
            v13 = append(v13, 1)
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v39
        }
        var v40 string = "foo bar"
        _ = v40
    }
    var v41 int16 = 878
    switch 1 {
    case 2147483647:
        v12 += -37
    }
    var v42 uint8 = (v5 - v5)
    v4 = genFn0(v41, -32768)
    for fi43 := 0; fi43 < 7; fi43++ {
        var v44 float64 = (-2.25 + -2.25)
        fmt.Println(v3, v3, v3)
        v45 := GenS2{F0: 0, F1: ""}
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
