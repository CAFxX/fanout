package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 bool) uint32 {
    v7 := []int{4294967295, 433, 637, 229}
    var v8 bool = true
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 uint16 = (v2 - v2)
    v4 := []int{-465, -525}
    var v5 uint = 4294967295
    var v6 string = ("foo bar" + "")
    switch int(v3) {
    case 356:
        if true {
            v9 := GenS1{F0: -128, F1: (v6 + "\n\t"), F2: 42}
            if ((1e10 + 3.14159) >= 0.1) {
                var v10 bool = (909 == genFn0(836, true))
                var v11 int8 = 127
                var v12 int = -754
                _ = v10
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 6; fi13++ {
                    fmt.Println(v6)
                    v14 := GenS1{F0: 127, F1: v6, F2: 32767}
                    _ = fi13
                    _ = v14
                }
            }
            _ = v9
        }
    default:
        var v15 float32 = 0.0
        _ = v15
    }
    v2 = (v3 * v2)
    var v16 float64 = -2.25
    if true {
        if ((!(1 < 117)) || (-1 < 42)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (v2 == 469) {
                for fi17 := 0; fi17 < 18; fi17++ {
                    v18 := GenS1{F0: 1, F1: v6, F2: 32767}
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            v19 := GenS1{F0: 42, F1: "", F2: 1}
            _ = v19
        }
        if (42 >= -201) {
            v5 = 65535
            v3 -= v2
            var v20 string = v6
            _ = v20
        } else {
            var v21 int8 = -1
            _ = v21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := GenS1{F0: -128, F1: "a", F2: 42}
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    } else {
        v4 = append(v4, int(v5))
        if (true || false) {
            var v25 uint32 = 2147483647
            _ = v25
        } else {
            v5 -= 183
            for fi26 := 0; fi26 < 9; fi26++ {
                for fi27 := 0; fi27 < 13; fi27++ {
                    if false {
                        if true {
                            var v28 uint16 = v2
                            _ = v28
                        } else {
                            v3 = (^(v3))
                        }
                        var v29 uint = 326
                        _ = v29
                    }
                    _ = fi27
                }
                delete(v1, (v6 + v6))
                _ = fi26
            }
        }
    }
    var v30 float64 = 1.5
    _ = len(v4)
    v2 -= (^(v3))
    for fi31 := 0; fi31 < 3; fi31++ {
        v6 = (("" + (v6 + v6)) + v6)
        _ = fi31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        fmt.Println(v6, v5)
        v4 = append(v4, fi32)
        v6 = ("hello" + v6)
        _ = fi32
    }
    v33 := GenS1{F0: 63, F1: "", F2: -436}
    v5 = (v5 % (((^(v5))) | 1))
    switch v6 {
    case "\n\t":
        var v34 uint = (v5 & 37)
        _ = v34
    }
    delete(v1, v6)
    if (0 < v2) {
        fmt.Println(v5, v5, v6)
        switch "\n\t" {
        case "\n\t":
            v3 += v2
        case "":
            v35 := GenS2{F0: 1000}
            _ = v35
        case "foo bar":
            var v36 uint8 = 15
            _ = v36
        default:
            if false {
                fmt.Println(v30)
            }
        }
        v6 = v6
    } else {
        switch -2147483648 {
        case -9223372036854775808:
            v5 = v5
        default:
            _ = len(v1)
        }
    }
    var v37 float64 = (3.14159 - 0.1)
    if true {
        v5 = 0
    } else {
        var v38 uint32 = 687
        v39 := GenS1{F0: -128, F1: "", F2: 42}
        if (0 < 42) {
            v40 := GenS2{F0: 328}
            var v41 string = v6
            _ = len(v4)
            _ = v40
            _ = v41
        }
        _ = v38
        _ = v39
    }
    fmt.Println(v37, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
