package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int; F1 int8; F2 string }
func genFn0() uint16 {
    var v4 int8 = 0
    v4 = v4
    v5 := []int{42, 4294967295, -2147483648}
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 uint, p1 int, p2 int8) uint64 {
    var v6 int32 = -939
    fmt.Println(v6, p2)
    p1 -= p1
    v6 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return uint64(p0)
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    v3 := []int{391, 127, 32768, 9223372036854775807}
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            var v8 float64 = 3.14159
            var v9 uint = 255
            var v10 int8 = -128
            _ = v10
            _ = v8
            _ = v9
        }
        if len(v3) > 0 {
            v3[0] = 9223372036854775807
        }
        var v11 int8 = -128
        _ = fi7
        _ = v11
    }
    if (!false) {
        var v12 uint32 = 789
        switch v1 {
        case "":
            v13 := GenS1{F0: 32767}
            _ = v13
        case "a":
            v1 = v1
        }
        var v14 int64 = -9223372036854775808
        _ = v12
        _ = v14
    } else {
        v15 := GenS1{F0: 1}
        v16 := GenS2{F0: (0 >= 0)}
        _ = v15
        _ = v16
    }
    fmt.Println(v1)
    v17 := GenS3{F0: 2, F1: 40, F2: 127}
    fmt.Println(v1)
    switch v1 {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 0
        }
    case "":
        switch v1 {
        case "\n\t":
            for fi18 := 0; fi18 < 9; fi18++ {
                v1 = ("\n\t" + (v1 + ("hello" + ((("foo bar" + (v1 + (v1 + v1))) + v1) + v1))))
                switch (v1 + v1) {
                case "a":
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v1 = v1
                        switch (fi19 + fi19) {
                        case 4294967295:
                            var v20 int64 = 1
                            _ = v20
                        case 2147483647:
                            v1 = "a"
                        default:
                            if (42 == 65535) {
                                var v21 float32 = (0.1 - (0.1 + (0.1 / (float32((-2.25 / 0.1)) * 1.5))))
                                fmt.Println(v21, fi19, v1)
                                if true {
                                    var v22 uint64 = 406
                                    _ = v22
                                }
                                _ = v21
                            } else {
                                var v23 bool = false
                                switch 824 {
                                case 127:
                                    _ = len(v3)
                                case 821:
                                    v24 := GenS2{F0: false}
                                    _ = v24
                                case 2147483647:
                                    if (-2147483648 < 1) {
                                        var v25 uint8 = 205
                                        fmt.Println(v23)
                                        v23 = (!v23)
                                        _ = v25
                                    } else {
                                        v1 = v1
                                    }
                                default:
                                    if (!((true || v23) || (!false))) {
                                        fmt.Println(v1, fi18)
                                    }
                                }
                                v23 = (100 == 127)
                                _ = v23
                            }
                        }
                        fmt.Println(fi18)
                        _ = fi19
                    }
                case "\n\t":
                    v1 = (v1 + "foo bar")
                }
                v26 := GenS2{F0: (128 == 128)}
                _ = fi18
                _ = v26
            }
        case "foo bar":
            var v27 int32 = 12
            _ = v27
        case "a":
            if (4294967295 == -725) {
                v28 := GenS1{F0: 32767}
                if (42 <= genFn1(0, -1, 127)) {
                    var v29 int64 = 0
                    fmt.Println(v29, v1)
                    _ = v29
                } else {
                    v1 = ""
                }
                _ = v28
            } else {
                for fi30 := 0; fi30 < 19; fi30++ {
                    var v31 uint8 = 87
                    _ = fi30
                    _ = v31
                }
            }
        default:
            v1 = v1
        }
    case "a":
        v32 := GenS3{F0: 127, F1: 151, F2: 255}
        _ = v32
    default:
        if true {
            if true {
                var v33 float32 = 3.14159
                v34 := GenS3{F0: 127, F1: 207, F2: 189}
                v33 = (v33 + (v33 - v33))
                _ = v33
                _ = v34
            }
            fmt.Println(v1)
        } else {
            var v35 uint32 = 2147483647
            _ = v35
        }
    }
    v36 := GenS1{F0: 42}
    var v37 uint = 42
    var v38 bool = false
    if v38 {
        for fi39 := 0; fi39 < 3; fi39++ {
            v40 := GenS4{F0: (fi39 >> ((v37) & 63)), F1: int8(v37), F2: (v1 + ("" + (("hello" + v1) + v1)))}
            _ = fi39
            _ = v40
        }
        var v41 uint64 = 1
        var v42 uint = (v37 & v37)
        _ = v41
        _ = v42
    } else {
        v37 = (v37 - 2)
        var v43 float64 = (float64((1e10 / float32(1e10))) - 0.0)
        if (!(!(1 != 0))) {
            v1 = ("hello" + "foo bar")
        } else {
            var v44 bool = false
            v44 = (42 == 1)
            _ = v44
        }
        _ = v43
    }
    v2[0] = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
