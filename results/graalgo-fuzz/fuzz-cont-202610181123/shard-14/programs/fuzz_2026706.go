package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 uint64 }
func genFn0(p0 int16, p1 int64) uint {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 float32 = 1e10
        var v6 bool = (80 >= 92)
        v7 := []int{}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 634
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 448}
    var v3 uint16 = 114
    fmt.Println(v3)
    var v8 float32 = 0.1
    switch ("" + "a") {
    case "foo bar":
        if (((!(((v8 * (v8 * (v8 / 3.14159))) < v8) || true)) || (3.14159 >= 1.5)) && (!false)) {
            v8 = v8
            for fi9 := 0; fi9 < 1; fi9++ {
                if false {
                    switch ("a" + "") {
                    case "a":
                        if false {
                            switch "a" {
                            case "hello":
                                var v10 int = 32768
                                _ = v10
                            case "foo bar":
                                var v11 int8 = 0
                                _ = v11
                            }
                            if true {
                                v2 = append(v2, 42)
                            } else {
                                fmt.Println(v8, v8, fi9)
                            }
                        }
                    case "\n\t":
                        if ((("" + "a") < "hello") && true) {
                            var v12 uint64 = 916
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            v13 := v1[1]
                            _ = v12
                            _ = v13
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            v8 -= v8
                            for fi14 := 0; fi14 < 12; fi14++ {
                                fmt.Println(fi14, fi14)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi14
                            }
                        }
                    case "":
                        v8 -= -2.25
                    }
                    var v15 uint16 = (v3 ^ (v3 & v3))
                    _ = v15
                } else {
                    var v16 bool = (false && (!true))
                    var v17 int = (-(fi9))
                    _ = v16
                    _ = v17
                }
                _ = fi9
            }
            fmt.Println(v8)
        }
    default:
        _ = len(v2)
    }
    var v18 float32 = 0.0
    v19 := GenS1{F0: 255}
    _ = len(v2)
    switch "\n\t" {
    case "a":
        switch 114 {
        case -673:
            var v20 bool = (190 <= -47)
            _ = v20
        }
    case "\n\t":
        var v21 int32 = 42
        _ = v21
    default:
        var v22 bool = false
        _ = v22
    }
    if (-38 != 0) {
        var v23 uint = genFn0(-836, -1)
        _ = v23
    } else {
        var v24 int8 = 42
        if false {
            var v25 int32 = -2147483648
            var v26 string = (("" + "\n\t") + "")
            _ = v25
            _ = v26
        } else {
            fmt.Println(v8, v18)
            if (!(1 > 42)) {
                for fi27 := 0; fi27 < 8; fi27++ {
                    if (-2147483648 >= -553) {
                        v28 := GenS2{F0: 254, F1: 2147483647, F2: 4294967295}
                        _ = v28
                    }
                    var v29 int64 = 9223372036854775807
                    var v30 int16 = 1
                    _ = fi27
                    _ = v29
                    _ = v30
                }
                v1[1] = 99
                switch (("hello" + ("a" + "")) + "") {
                case "a":
                    if (32767 > 0) {
                        v8 = 3.14159
                    }
                default:
                    var v31 uint64 = 1
                    _ = v31
                }
            }
        }
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
