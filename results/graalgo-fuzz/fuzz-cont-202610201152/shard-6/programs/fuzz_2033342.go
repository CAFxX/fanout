package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint32, p1 uint8) bool {
    v6 := []int{}
    fmt.Println(p1)
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 = 911
        var v8 int8 = -1
        var v9 uint8 = p1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = -1
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return false
}

func main() {
    var v1 int16 = -35
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v5 int8 = 42
    fmt.Println(v1)
    var v11 string = "foo bar"
    switch v11 {
    case "\n\t":
        var v12 float32 = float32(0.0)
        _ = v12
    }
    if false {
        fmt.Println(v1, v1, v5)
    } else {
        var v13 uint16 = 966
        v13 = (v13 * v13)
        delete(v3, 4294967295)
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 float64 = 0.1
        var v16 uint64 = 18446744073709551615
        fmt.Println(v1)
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 2, F1: 1}
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            fmt.Println(v11, v11, v5)
            _ = fi19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 uint = 390
            _ = fi20
            _ = v21
        }
        var v22 int8 = -128
        _ = fi18
        _ = v22
    }
    v2 = append(v2, 2147483647)
    for fi23 := 0; fi23 < 5; fi23++ {
        v1 += v1
        _ = fi23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        fmt.Println(v5)
        switch v11 {
        case "\n\t":
            fmt.Println(v11)
        case "":
            var v25 int32 = 2147483647
            _ = v25
        case "foo bar":
            for fi26 := 0; fi26 < 6; fi26++ {
                v1 += v1
                fmt.Println(v1, v1, v5)
                _ = fi26
            }
        }
        if (42 < 42) {
            _ = len(v2)
            v27 := v3[42]
            for fi28 := 0; fi28 < 4; fi28++ {
                switch v11 {
                case "a":
                    switch (-(fi28)) {
                    case 4294967295:
                        var v29 uint = 255
                        _ = v29
                    }
                case "foo bar":
                    v5 = v5
                default:
                    v5 = v5
                }
                var v30 string = "foo bar"
                if v27 {
                    fmt.Println(fi24, v11, fi28)
                } else {
                    v30 = v11
                    _ = len(v2)
                }
                _ = fi28
                _ = v30
            }
            _ = v27
        }
        _ = fi24
    }
    if (!(true && (-2147483648 != -2147483648))) {
        switch -2147483648 {
        case 42:
            fmt.Println(v5, v5)
        case -581:
            v31 := GenS1{F0: 18446744073709551615, F1: 1}
            _ = v31
        }
    } else {
        v11 = "hello"
        v5 = v5
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v32 int = 0
    var v33 float32 = ((-2.25 / 1.0) / 1.5)
    var v34 int64 = 572
    fmt.Println(v11)
    v5 = (v5 * (-(v5)))
    for fi35 := 0; fi35 < 7; fi35++ {
        if false {
            for fi36 := 0; fi36 < 4; fi36++ {
                fmt.Println(v32, fi35)
                _ = fi36
            }
            if (731 == 65535) {
                v3[42] = false
                v11 = ((v11 + "foo bar") + (v11 + "a"))
            }
            v32 = -1
        }
        v5 = (v5 - v5)
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
