package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int; F2 bool }
func genFn0(p0 uint32) string {
    switch 619 {
    case 9223372036854775807:
        if ((!(4294967295 == p0)) && false) {
            v6 := []int{563, 4294967295}
            _ = v6
        }
    case 1:
        v7 := GenS1{F0: ("" + "hello")}
        _ = v7
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS1{F0: ("hello" + (("foo bar" + "foo bar") + (("a" + "a") + "a")))}
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS2{F0: uint(p0), F1: -9223372036854775808, F2: (true || (0.0 > float32(float64((3.14159 + 0.0)))))}
    v11 := []int{32768, -1}
    var v12 uint8 = 2
    switch 4294967295 {
    case -1:
        if ((0 < 1) && (284 > 42)) {
            v12 = (v12 + (v12 | 2))
            p0 = (p0 / ((p0) | 1))
            v12 += v12
        }
    default:
        var v13 uint = 4294967295
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return "a"
}

func main() {
    var v1 int64 = 0
    var v2 int8 = 82
    var v3 int64 = v1
    v4 := []uint32{}
    v5 := make(map[string]int)
    fmt.Println(v2, v2, v2)
    for fi14 := 0; fi14 < 12; fi14++ {
        if ((("\n\t" >= "\n\t") && (2 == 1)) && ("\n\t" != "a")) {
            if true {
                var v15 int8 = (v2 * v2)
                for fi16 := 0; fi16 < 15; fi16++ {
                    _ = len(v5)
                    _ = fi16
                }
                v5[""] = fi14
                _ = v15
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    switch ("" + "") {
                    case "":
                        delete(v5, "foo bar")
                    case "foo bar":
                        var v18 uint32 = 2147483647
                        _ = v18
                    }
                    switch ("\n\t" + (("\n\t" + "") + "\n\t")) {
                    case "\n\t":
                        var v19 string = ""
                        _ = v19
                    case "":
                        fmt.Println(v3, fi14, v2)
                    default:
                        v3 = (v1 << ((42) & 63))
                        v20 := GenS2{F0: 2, F1: -2147483648, F2: ("" > "")}
                        _ = v20
                    }
                    _ = len(v4)
                    _ = fi17
                }
                if true {
                    var v21 uint = 255
                    _ = v21
                }
                if false {
                    var v22 uint64 = 0
                    _ = v22
                } else {
                    if true {
                        var v23 int16 = 1
                        _ = v23
                    }
                    delete(v5, ("foo bar" + ""))
                    var v24 uint = 255
                    _ = v24
                }
            }
        }
        delete(v5, ("\n\t" + "hello"))
        _ = fi14
    }
    if false {
        v1 -= -431
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            var v25 int = 42
            _ = v25
        }
    }
    var v26 uint64 = 4294967295
    for fi27 := 0; fi27 < 17; fi27++ {
        v3 = v1
        v1 = 9223372036854775807
        v26 -= v26
        _ = fi27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        v2 += 29
        _ = fi28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        v2 -= v2
        _ = fi29
    }
    v1 += (v3 - (v1 + int64(v2)))
    v2 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
