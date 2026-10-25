package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int; F1 string; F2 int8 }
func genFn0(p0 uint8, p1 string, p2 uint) uint {
    var v6 uint32 = 0
    if (2147483647 <= -1) {
        v7 := GenS1{F0: 1000, F1: -2147483648}
        switch p1 {
        case "foo bar":
            v8 := GenS1{F0: 32767, F1: 28}
            _ = v8
        }
        v9 := GenS1{F0: -1, F1: 9223372036854775807}
        _ = v7
        _ = v9
    } else {
        var v10 string = ("\n\t" + p1)
        for fi11 := 0; fi11 < 19; fi11++ {
            if true {
                var v12 uint64 = 4294967295
                _ = v12
            } else {
                v13 := GenS1{F0: 0, F1: -735}
                _ = v13
            }
            v14 := GenS1{F0: 32767, F1: (fi11 >> ((p2) & 63))}
            if false {
                if true {
                    var v15 float32 = 3.14159
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 int16 = 899
                        v18 := GenS2{F0: fi11, F1: ("a" + "\n\t"), F2: 1}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    if false {
                        var v19 uint = (p2 ^ (p2 >> ((p2) & 63)))
                        v6 = v6
                        var v20 int16 = 1000
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint = p2
                        _ = v21
                    }
                    _ = v15
                } else {
                    var v22 float64 = 1.5
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 uint = 1
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
                switch (p1 + ("\n\t" + (v10 + v10))) {
                case "\n\t":
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v26 := []int{(-(fi25))}
                        _ = fi25
                        _ = v26
                    }
                case "foo bar":
                    v27 := GenS2{F0: 1, F1: "\n\t", F2: 1}
                    _ = v27
                case "hello":
                    switch 864 {
                    case 4294967295:
                        fmt.Println(v10, v6, p2)
                    case 61:
                        v28 := GenS2{F0: (-(fi11)), F1: v10, F2: 0}
                        _ = v28
                    default:
                        var v29 float32 = 0.1
                        _ = v29
                    }
                default:
                    fmt.Println(p0)
                }
                p2 = 686
            }
            _ = fi11
            _ = v14
        }
        v30 := GenS2{F0: 2147483647, F1: p1, F2: -121}
        _ = v10
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 776
}

func main() {
    var v1 int16 = 1
    var v2 int32 = 301
    var v3 float32 = 3.14159
    v4 := make(map[string]string)
    var v5 uint32 = 632
    _ = len(v4)
    if false {
        v31 := GenS2{F0: 1, F1: ("a" + "\n\t"), F2: 95}
        v1 = v1
        _ = v31
    }
    var v32 int8 = 1
    if (616 >= genFn0(127, "a", 536)) {
        for fi33 := 0; fi33 < 11; fi33++ {
            v32 -= v32
            var v34 bool = (!false)
            var v35 uint = 4294967295
            _ = fi33
            _ = v34
            _ = v35
        }
        v3 = (v3 + 0.1)
        var v36 int64 = -1
        _ = v36
    }
    v4["foo bar"] = "a"
    var v37 float64 = -2.25
    v38 := GenS2{F0: 42, F1: (("" + "a") + "\n\t"), F2: v32}
    var v39 int32 = (-(v2))
    if (!(42 <= 18446744073709551615)) {
        var v40 string = "\n\t"
        switch 127 {
        case -1:
            fmt.Println(v40, v1, v32)
        default:
            var v41 uint64 = 42
            _ = v41
        }
        v1 -= v1
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
