package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool; F2 uint8 }
func genFn0(p0 int32, p1 uint16) string {
    p0 = -363
    var v7 uint = 4294967295
    if true {
        fmt.Println(p0)
        v8 := []int{2147483647, 32767, -1, 89}
        _ = v8
    } else {
        v9 := GenS1{F0: 968}
        switch "foo bar" {
        case "foo bar":
            v10 := GenS1{F0: 65535}
            _ = v10
        default:
            fmt.Println(v7, p1, p1)
        }
        var v11 bool = ("foo bar" > "hello")
        _ = v11
        _ = v9
    }
    switch 243 {
    case 128:
        if true {
            p1 = 735
            var v12 int64 = -1
            var v13 int8 = -128
            _ = v12
            _ = v13
        }
    case 42:
        v14 := GenS1{F0: 1}
        _ = v14
    case -9223372036854775808:
        if false {
            v15 := []int{}
            _ = v15
        } else {
            var v16 int16 = -654
            v16 = v16
            var v17 uint64 = 1
            _ = v16
            _ = v17
        }
    default:
        var v18 uint32 = 156
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v7
    return ("" + "\n\t")
}
func genFn1(p0 bool) bool {
    if true {
        var v19 int8 = 127
        _ = v19
    }
    if false {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := []int{fi20}
            _ = fi20
            _ = v21
        }
        var v22 uint32 = 122
        _ = v22
    }
    fmt.Println(p0)
    v23 := []int{2147483647, 127, 0}
    _ = p0
    _ = v23
    return p0
}

func main() {
    var v1 int = 42
    var v2 uint16 = 234
    var v3 int64 = -1
    var v4 int32 = -137
    var v5 float64 = 0.0
    var v6 int64 = v3
    v24 := []int{(-(v1)), 4294967295}
    var v25 int = v1
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int = v1
        _ = len(v24)
        var v28 uint16 = (v2 - v2)
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        fmt.Println(fi29, v1)
        fmt.Println(v2, v2, v2)
        _ = fi29
    }
    fmt.Println(v6, v4)
    for fi30 := 0; fi30 < 4; fi30++ {
        if true {
            v24 = append(v24, (fi30 >> ((65535) & 63)))
            fmt.Println(v25, v6, fi30)
            fmt.Println(v1)
        }
        switch ("\n\t" + (("foo bar" + ((((("" + "hello") + "") + "") + ("foo bar" + "")) + ("a" + ("hello" + "a")))) + "hello")) {
        case "foo bar":
            var v31 uint32 = 935
            _ = v31
        case "":
            if true {
                v32 := GenS1{F0: 712}
                if (!(42 > 1)) {
                    fmt.Println(v25, v5)
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                }
                v33 := GenS1{F0: 65535}
                _ = v32
                _ = v33
            }
        case "hello":
            if len(v24) > 0 {
                v24[0] = v1
            }
        default:
            for fi34 := 0; fi34 < 11; fi34++ {
                v35 := GenS1{F0: 4294967295}
                _ = fi34
                _ = v35
            }
        }
        if (v4 != 42) {
            v3 = v6
            v1 = fi30
        } else {
            var v36 int64 = v6
            if false {
                _ = len(v24)
                _ = len(v24)
            } else {
                if true {
                    for fi37 := 0; fi37 < 1; fi37++ {
                        var v38 uint64 = 935
                        _ = fi37
                        _ = v38
                    }
                    v3 = v3
                    if false {
                        var v39 float64 = (v5 - v5)
                        if true {
                            v39 -= 3.14159
                        } else {
                            v40 := GenS1{F0: 65535}
                            _ = v40
                        }
                        _ = v39
                    }
                } else {
                    if false {
                        fmt.Println(v5, v3, v3)
                    }
                    var v41 int16 = 395
                    _ = v41
                }
            }
            v42 := GenS1{F0: 783}
            _ = v36
            _ = v42
        }
        _ = fi30
    }
    var v43 float32 = (1e10 + -2.25)
    fmt.Println(v4)
    v44 := GenS2{F0: (("a" + "foo bar") + "foo bar"), F1: (false || true), F2: 255}
    v45 := GenS2{F0: "", F1: (!false), F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
