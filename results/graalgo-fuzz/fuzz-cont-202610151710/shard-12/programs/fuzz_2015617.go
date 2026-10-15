package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 int8, p1 uint64, p2 uint16) uint8 {
    if false {
        var v6 int = 32767
        switch v6 {
        case -372:
            var v7 int16 = 1000
            _ = v7
        }
        _ = v6
    } else {
        var v8 uint = 42
        _ = v8
    }
    if (int64(p2) > 0) {
        var v9 float32 = (3.14159 + (-2.25 + 1e10))
        var v10 float64 = ((1e10 / (float64(v9) + 0.0)) * float64((v9 + 0.1)))
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        switch (("" + "foo bar") + "") {
        case "foo bar":
            if (1 < fi11) {
                var v12 bool = true
                _ = v12
            }
        case "hello":
            v13 := []int{}
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS1{F0: 910, F1: (p0 % (((p0 & (-(p0)))) | 1))}
                _ = fi14
                _ = v15
            }
        }
        if ((true && true) && false) {
            var v16 int64 = int64(p1)
            _ = v16
        }
        v17 := GenS2{F0: 1000}
        _ = fi11
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 string = ("a" + "")
    var v5 uint16 = 42
    v18 := GenS3{F0: 710, F1: 0}
    switch 1 {
    case 240:
        v19 := GenS2{F0: -1}
        _ = v19
    case 128:
        switch v4 {
        case "a":
            if (true || true) {
                if true {
                    v4 = (v4 + ((v4 + v4) + (v4 + "foo bar")))
                    v4 = v4
                } else {
                    for fi20 := 0; fi20 < 14; fi20++ {
                        fmt.Println(v5)
                        var v21 int16 = -149
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v23 := GenS1{F0: 96, F1: 100}
                        var v24 uint16 = v5
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := []int{42, 9223372036854775807, 32767, 1}
                    _ = v25
                }
            }
        case "foo bar":
            if (!false) {
                v26 := v3["\n\t"]
                _ = v26
            }
        default:
            v4 = ("hello" + v4)
        }
    }
    v27 := GenS3{F0: 159, F1: 255}
    delete(v2, -1)
    if (0 > 65535) {
        v4 = v4
        v28 := []int{128, 32768}
        _ = v28
    }
    v4 = (v4 + v4)
    if false {
        v4 = v4
    } else {
        v5 = (v5 & (^(v5)))
        v4 = "hello"
        v4 = v4
    }
    var v29 uint32 = 42
    v4 = ""
    if false {
        v29 = 489
        fmt.Println(v5)
        for fi30 := 0; fi30 < 3; fi30++ {
            v5 += v5
            v4 = (v4 + v4)
            fmt.Println(fi30)
            _ = fi30
        }
    } else {
        var v31 uint16 = v5
        var v32 bool = true
        _ = v31
        _ = v32
    }
    var v33 uint = 18446744073709551615
    v2[-1] = (true && true)
    if true {
        for fi34 := 0; fi34 < 8; fi34++ {
            if (999 <= 1) {
                v35 := v3[((((v4 + ((v4 + v4) + "\n\t")) + v4) + (v4 + (v4 + (v4 + v4)))) + v4)]
                var v36 float32 = 0.1
                _ = v35
                _ = v36
            } else {
                v5 = v5
            }
            _ = fi34
        }
        v37 := GenS2{F0: 1000}
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
