package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int16; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS6 struct { F0 int16; F1 uint8 }
func genFn0(p0 int32, p1 int64, p2 bool) uint {
    if p2 {
        if ("hello" == "\n\t") {
            p2 = (p0 < (p0 + (p0 << ((255) & 63))))
        } else {
            v7 := GenS1{F0: 0}
            for fi8 := 0; fi8 < 14; fi8++ {
                p1 = -153
                v9 := GenS2{F0: 743, F1: p2, F2: 127}
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        var v10 uint = 4294967295
        _ = v10
    }
    p1 = p1
    var v11 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := []int{576, 121}
    var v4 int8 = 1
    var v5 uint64 = 0
    var v6 bool = (!(128 != 42))
    if (true || (!(!true))) {
        var v12 int8 = (v4 * (v4 >> ((0) & 63)))
        v13 := GenS1{F0: v5}
        _ = v12
        _ = v13
    } else {
        switch 2147483647 {
        case -381:
            v4 -= v4
        }
        var v14 int16 = 0
        delete(v1, "a")
        _ = v14
    }
    v3 = append(v3, 127)
    var v15 int8 = v4
    v16 := GenS3{F0: 240, F1: -1}
    v17 := GenS3{F0: 2, F1: 0}
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS4{F0: 42, F1: genFn0(1, 9223372036854775807, (!true)), F2: (v5 << ((0) & 63))}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 float64 = 1.5
        v22 := GenS5{F0: 42, F1: 42, F2: 42}
        switch ("\n\t" + "foo bar") {
        case "hello":
            v2["\n\t"] = (226 <= 112)
        case "\n\t":
            if (v6 && v6) {
                if (17 >= v4) {
                    v4 -= (v4 ^ (v4 % ((v4) | 1)))
                } else {
                    v23 := v1[("\n\t" + "a")]
                    var v24 int = fi20
                    _ = v23
                    _ = v24
                }
                var v25 float64 = v21
                _ = v25
            } else {
                if false {
                    var v26 string = (("hello" + "foo bar") + "\n\t")
                    v26 = "foo bar"
                    _ = v26
                } else {
                    var v27 float64 = v21
                    delete(v1, "foo bar")
                    _ = v27
                }
                switch ((("foo bar" + ("foo bar" + "hello")) + "foo bar") + "foo bar") {
                case "":
                    fmt.Println(v15, v15, v15)
                }
            }
        default:
            for fi28 := 0; fi28 < 9; fi28++ {
                var v29 int64 = 405
                _ = fi28
                _ = v29
            }
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        v15 -= v4
        switch fi30 {
        case -1:
            fmt.Println(fi30)
        }
        _ = fi30
    }
    v31 := GenS3{F0: 42, F1: 9223372036854775807}
    var v32 uint16 = 42
    v6 = (1e10 != 1.5)
    if len(v3) > 0 {
        v3[0] = -215
    }
    v5 += (^(v5))
    fmt.Println(v6)
    for fi33 := 0; fi33 < 11; fi33++ {
        if v6 {
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 int16 = 1000
                v36 := v2["\n\t"]
                _ = fi34
                _ = v35
                _ = v36
            }
            v37 := GenS6{F0: -1, F1: 2}
            _ = v37
        }
        v38 := GenS6{F0: 1, F1: 42}
        _ = fi33
        _ = v38
    }
    var v39 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
