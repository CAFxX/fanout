package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
func genFn0() uint8 {
    if (!((true || (true && (1 != -1))) || (!false))) {
        var v5 int16 = -993
        _ = v5
    }
    var v6 uint64 = 1
    if false {
        v7 := []int{42, 32767, -552, 127}
        var v8 int64 = 42
        if (!("a" > "a")) {
            if true {
                if (("\n\t" < ("foo bar" + (("foo bar" + "hello") + ""))) && ((!false) || true)) {
                    var v9 string = "foo bar"
                    var v10 float32 = 0.0
                    _ = v10
                    _ = v9
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v11 int = 684
                        _ = v11
                    case "a":
                        v12 := GenS1{F0: 649, F1: 1000, F2: 128}
                        _ = v12
                    case "hello":
                        if (!(!true)) {
                            switch ("a" + ("hello" + ("\n\t" + (("\n\t" + (("foo bar" + "") + "a")) + "a")))) {
                            case "\n\t":
                                var v13 uint32 = 2147483647
                                _ = v13
                            case "a":
                                v8 -= 919
                            default:
                                v8 = v8
                            }
                            if ((((127 == 128) && (9223372036854775807 != -288)) || ((!(0 < 1)) && false)) || (!(("foo bar" + "foo bar") != "foo bar"))) {
                                var v14 int64 = (-(v8))
                                v6 = 4294967295
                                _ = v14
                            }
                        }
                    }
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 int8 = 100
                        _ = fi15
                        _ = v16
                    }
                    var v17 int16 = 32767
                    _ = v17
                }
                var v18 int8 = 42
                _ = v18
            } else {
                fmt.Println(v8, v6)
                fmt.Println(v6, v6)
            }
            if true {
                v19 := GenS2{F0: 65535, F1: (v6 * (v6 ^ (v6 & (v6 << ((4294967295) & 63)))))}
                _ = v19
            }
        }
        _ = v7
        _ = v8
    } else {
        v20 := GenS3{F0: 42}
        var v21 string = ("\n\t" + ("foo bar" + ("hello" + ("foo bar" + "foo bar"))))
        if true {
            v22 := []int{619, -363, 1}
            if true {
                var v23 float64 = ((float64((0.1 + 0.0)) * 3.14159) / 1e10)
                _ = v23
            }
            if len(v22) > 0 {
                v22[0] = 32767
            }
            _ = v22
        } else {
            fmt.Println(v21, v6)
            var v24 int = 32768
            _ = v24
        }
        _ = v20
        _ = v21
    }
    fmt.Println(v6)
    var v25 float64 = 0.0
    _ = v25
    _ = v6
    return 255
}
func genFn1() int32 {
    v26 := []int{162, 559, 562, 1}
    var v27 bool = true
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    v26 = append(v26, 27)
    var v28 int16 = -32768
    _ = v26
    _ = v27
    _ = v28
    return -1
}

func main() {
    var v1 uint16 = 2
    var v2 uint16 = 543
    v3 := []int{}
    var v4 uint8 = 0
    v4 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v29 := GenS2{F0: (v2 / ((v1) | 1)), F1: 4294967295}
    v1 = uint16(v4)
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    v30 := GenS4{F0: -889}
    var v31 uint8 = genFn0()
    fmt.Println(v1)
    var v32 float32 = 1.5
    if false {
        v2 = 42
    } else {
        fmt.Println(v2)
        for fi33 := 0; fi33 < 9; fi33++ {
            var v34 uint16 = 32767
            fmt.Println(v31, v34, v31)
            for fi35 := 0; fi35 < 18; fi35++ {
                switch (fi35 ^ fi33) {
                case 9223372036854775807:
                    var v36 int8 = 102
                    var v37 float32 = v32
                    fmt.Println(v37, fi33, v31)
                    _ = v36
                    _ = v37
                case 128:
                    var v38 int16 = 1000
                    _ = v38
                }
                _ = fi35
            }
            _ = fi33
            _ = v34
        }
        var v39 uint64 = 42
        _ = v39
    }
    fmt.Println(v4, v4, v4)
    v3 = append(v3, 2147483647)
    fmt.Println(v4)
    var v40 float32 = v32
    if len(v3) > 0 {
        v3[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v40
}
