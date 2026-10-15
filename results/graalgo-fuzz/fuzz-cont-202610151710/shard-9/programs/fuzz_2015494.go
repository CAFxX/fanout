package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
func genFn0(p0 int8, p1 uint32, p2 uint16) int {
    v4 := GenS1{F0: "hello", F1: -2147483648, F2: (4294967295 < 68)}
    if (-430 >= -573) {
        v5 := []int{1, 32768, 1}
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int64 = 1
            fmt.Println(p2, p2)
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        var v8 uint = 42
        _ = v8
    }
    fmt.Println(p1, p2)
    v9 := []int{32768, 9223372036854775807, 0, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return 622
}
func genFn1() int8 {
    v10 := []int{127, 4294967295}
    var v11 uint16 = 1
    for fi12 := 0; fi12 < 8; fi12++ {
        v11 = 2
        v11 = (v11 << ((255) & 63))
        _ = fi12
    }
    switch -729 {
    case 32767:
        v13 := GenS1{F0: ("\n\t" + ""), F1: 2147483647, F2: (2 < 255)}
        _ = v13
    case 866:
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int16 = 0
            _ = fi14
            _ = v15
        }
    default:
        fmt.Println(v11)
    }
    if true {
        v16 := GenS1{F0: ("" + ""), F1: 2147483647, F2: (((-2.25 + 0.0) * ((0.0 * 0.1) + (0.1 / 1e10))) >= 1.5)}
        if (9223372036854775807 == 4294967295) {
            var v17 bool = (!(!true))
            _ = v17
        } else {
            for fi18 := 0; fi18 < 16; fi18++ {
                v11 -= v11
                _ = fi18
            }
            v19 := GenS1{F0: "a", F1: 42, F2: (56 > 127)}
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 8; fi20++ {
                v11 += 42
                _ = fi20
            }
            if (1 < 463) {
                v21 := GenS1{F0: ("hello" + (("foo bar" + ((("hello" + "") + "") + "hello")) + "hello")), F1: -1, F2: true}
                var v22 float64 = 0.0
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                if false {
                    if false {
                        var v24 int32 = 0
                        _ = v24
                    }
                }
                _ = fi23
            }
        }
        _ = v16
    }
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    if (!(165 >= 4294967295)) {
        var v25 int16 = -883
        _ = v25
    } else {
        v26 := []int{32768, 4294967295, genFn0(-30, 2, 32767), -309}
        var v27 int16 = -560
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 uint16 = 65535
            var v30 bool = ("a" == ("foo bar" + ((("a" + "hello") + ("" + "foo bar")) + "a")))
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = v26
        _ = v27
    }
    var v31 uint16 = 65535
    var v32 float32 = 1e10
    var v33 int32 = -2147483648
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 int32 = 2147483647
        fmt.Println(v35, fi34, v31)
        var v36 uint64 = 2
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 int64 = 42
    fmt.Println(v37, v37, v32)
    if false {
        var v38 uint16 = v31
        _ = v38
    }
    _ = len(v2)
    v39 := []int{9223372036854775807}
    if false {
        if (659 < uint32(v31)) {
            v3[("hello" + (("" + "hello") + "hello"))] = false
            var v40 int64 = v37
            _ = v40
        }
        v33 = (v33 / (((v33 + (v33 * (v33 | v33)))) | 1))
        var v41 int64 = -9223372036854775808
        _ = v41
    } else {
        if len(v39) > 0 {
            v39[0] = -940
        }
        var v42 float64 = -2.25
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v39
}
