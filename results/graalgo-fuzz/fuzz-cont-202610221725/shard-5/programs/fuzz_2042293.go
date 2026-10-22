package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0() int32 {
    switch 2147483647 {
    case 780:
        var v4 uint8 = 2
        _ = v4
    case 0:
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := GenS1{F0: int64(fi5)}
            if false {
                v7 := []int{}
                var v8 uint32 = 2147483648
                fmt.Println(fi5)
                _ = v7
                _ = v8
            } else {
                var v9 uint8 = 0
                var v10 int = fi5
                _ = v10
                _ = v9
            }
            var v11 int64 = -9223372036854775808
            _ = fi5
            _ = v11
            _ = v6
        }
    case 128:
        switch ("hello" + "foo bar") {
        case "hello":
            var v12 uint = 18446744073709551615
            _ = v12
        default:
            var v13 int64 = 1
            _ = v13
        }
    default:
        var v14 uint = 0
        _ = v14
    }
    switch 349 {
    case 2147483647:
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint8 = 2
            fmt.Println(fi15, v16)
            _ = fi15
            _ = v16
        }
    default:
        var v17 bool = (false || false)
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 float32 = 3.14159
        switch (-(fi18)) {
        case -766:
            switch 132 {
            case 0:
                v19 = v19
            }
        case -1:
            switch "hello" {
            case "hello":
                v19 += v19
            case "":
                var v20 float32 = float32(((1.5 / (-2.25 - (1e10 * (0.1 - (1.5 / float64((v19 - 1e10))))))) - 0.1))
                _ = v20
            default:
                for fi21 := 0; fi21 < 15; fi21++ {
                    v19 += ((v19 / 1.5) - 1e10)
                    var v22 bool = (true && true)
                    switch ("" + "foo bar") {
                    case "":
                        v23 := []int{(-(fi18)), 522, (fi21 | fi18), -537}
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
            }
        case 127:
            v24 := []int{(fi18 & (-(fi18))), fi18}
            _ = v24
        }
        _ = fi18
        _ = v19
    }
    switch "a" {
    case "foo bar":
        var v25 int32 = 678
        _ = v25
    }
    v26 := GenS2{F0: 1, F1: 46}
    _ = v26
    return 2147483647
}
func genFn1() uint8 {
    var v27 uint64 = 85
    var v28 int64 = 9223372036854775807
    v29 := []int{1}
    _ = v27
    _ = v28
    _ = v29
    return 209
}

func main() {
    var v1 uint16 = 42
    var v2 float32 = 3.14159
    v3 := make(map[string]string)
    var v30 string = "hello"
    for fi31 := 0; fi31 < 5; fi31++ {
        fmt.Println(v2, fi31)
        _ = fi31
    }
    fmt.Println(v2, v2, v2)
    _ = len(v3)
    fmt.Println(v30, v1, v2)
    v32 := []int{-2147483648, 1}
    var v33 int8 = 127
    var v34 float32 = ((v2 + v2) - ((0.0 * 0.0) + float32(float64((v2 + (0.0 / 0.1))))))
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 int = fi35
        for fi37 := 0; fi37 < 6; fi37++ {
            if len(v32) > 0 {
                v32[0] = v36
            }
            delete(v3, "")
            _ = fi37
        }
        v34 = 0.0
        _ = fi35
        _ = v36
    }
    if false {
        var v38 float32 = -2.25
        _ = v38
    } else {
        var v39 int16 = 1000
        if false {
            fmt.Println(v1, v39, v39)
        } else {
            _ = len(v3)
            switch 42 {
            case -733:
                v40 := GenS2{F0: (v39 + (v39 >> ((65535) & 63))), F1: v39}
                _ = v40
            default:
                fmt.Println(v34, v34, v33)
            }
            var v41 float64 = (-2.25 - (0.0 - 0.0))
            _ = v41
        }
        _ = v39
    }
    switch v30 {
    case "":
        for fi42 := 0; fi42 < 3; fi42++ {
            if (!true) {
                var v43 int16 = -1
                var v44 float32 = v2
                _ = len(v32)
                _ = v43
                _ = v44
            }
            _ = fi42
        }
    case "hello":
        var v45 float64 = 0.0
        _ = v45
    default:
        for fi46 := 0; fi46 < 11; fi46++ {
            v33 = (v33 - v33)
            var v47 string = v30
            switch (v30 + v30) {
            case "\n\t":
                var v48 float64 = 1.5
                _ = v48
            }
            _ = fi46
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
}
