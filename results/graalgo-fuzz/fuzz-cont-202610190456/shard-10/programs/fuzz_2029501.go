package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0() int64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{fi6, (fi6 >> ((1) & 63))}
        v8 := GenS1{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 65535}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = 42
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 1.5
    var v3 float64 = 1.5
    var v4 int64 = 42
    v5 := []uint32{}
    var v10 string = "hello"
    v2 = v3
    switch ((("a" + "") + v10) + ("a" + v10)) {
    case "a":
        fmt.Println(v10)
    case "":
        v2 = 3.14159
    }
    fmt.Println(v4)
    var v11 int16 = 42
    if false {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint8 = 128
            if false {
                v14 := GenS1{F0: (v10 + v10), F1: 592}
                var v15 uint16 = 919
                _ = v14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        var v16 uint64 = 4294967295
        _ = v16
    } else {
        var v17 uint = 61
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 uint64 = 1
            fmt.Println(v4, v1)
            var v20 float32 = 1e10
            _ = fi18
            _ = v19
            _ = v20
        }
        v17 = (^(v17))
        _ = v17
    }
    switch 42 {
    case 844:
        if (false || false) {
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 string = v10
                _ = fi21
                _ = v22
            }
            if ((!false) && false) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        } else {
            var v23 bool = (((18446744073709551615 == 1) && (true || ((!(true && false)) || false))) || (!true))
            if ((!(2147483647 < 42)) && (v23 && (!v23))) {
                for fi24 := 0; fi24 < 6; fi24++ {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                        _ = fi25
                    }
                    var v26 float64 = 1e10
                    var v27 int64 = genFn0()
                    _ = fi24
                    _ = v26
                    _ = v27
                }
                var v28 float64 = (v2 / 1e10)
                _ = v28
            } else {
                var v29 uint = 1
                if (68 == 1) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v30 int = 4294967295
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    _ = v30
                }
                _ = v29
            }
            v10 = v10
            _ = v23
        }
    case 2147483647:
        for fi31 := 0; fi31 < 9; fi31++ {
            fmt.Println(v1)
            _ = fi31
        }
    case -441:
        v1 = 255
    default:
        v5 = append(v5, 2)
    }
    var v32 uint = 42
    _ = len(v5)
    switch "" {
    case "a":
        fmt.Println(v2, v32, v3)
    case "foo bar":
        fmt.Println(v11, v1)
    case "hello":
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
