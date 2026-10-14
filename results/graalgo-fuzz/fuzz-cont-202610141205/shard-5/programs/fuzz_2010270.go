package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32) int8 {
    if true {
        var v6 int32 = (^(p0))
        fmt.Println(v6)
        fmt.Println(p0, v6)
        _ = v6
    }
    v7 := []int{}
    v7 = append(v7, 367)
    var v8 uint = 42
    v9 := GenS1{F0: ("" + "foo bar"), F1: -1, F2: -257}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint16 = 431
    var v2 bool = true
    var v3 int16 = -264
    var v4 uint32 = 2147483648
    v5 := []int{125, 4294967295, 32767}
    v1 -= (^(v1))
    var v10 uint = 879
    switch -552 {
    case 128:
        v3 = -1
    case 4294967295:
        v3 = -203
    default:
        var v11 uint32 = (^(v4))
        _ = v11
    }
    if v2 {
        switch 0 {
        case 976:
            v12 := GenS1{F0: "hello", F1: 100, F2: 358}
            _ = v12
        case -666:
            v4 -= (v4 / (((v4 % ((v4) | 1))) | 1))
        case -30:
            var v13 uint8 = 2
            _ = v13
        }
        if len(v5) > 0 {
            v5[0] = 42
        }
        var v14 int8 = 0
        _ = v14
    }
    var v15 int16 = v3
    if len(v5) > 0 {
        v5[0] = -1
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v15 = (^(v15))
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int64 = 1
        _ = fi16
        _ = v17
    }
    var v18 string = "foo bar"
    if (0 <= 42) {
        if (42 >= -1) {
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 float64 = 0.1
                _ = fi19
                _ = v20
            }
        } else {
            switch v18 {
            case "foo bar":
                for fi21 := 0; fi21 < 20; fi21++ {
                    v22 := GenS1{F0: v18, F1: genFn0(0), F2: (fi21 >> ((v10) & 63))}
                    _ = fi21
                    _ = v22
                }
            case "\n\t":
                for fi23 := 0; fi23 < 11; fi23++ {
                    v10 -= 2
                    _ = fi23
                }
            }
            v24 := GenS1{F0: (v18 + v18), F1: -41, F2: 32768}
            _ = v24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 int64 = 42
                switch v18 {
                case "":
                    fmt.Println(v10, v1)
                case "foo bar":
                    var v28 int = fi26
                    _ = v28
                case "a":
                    v29 := GenS1{F0: "\n\t", F1: -98, F2: fi25}
                    _ = v29
                }
                v30 := GenS2{F0: v27}
                _ = fi26
                _ = v27
                _ = v30
            }
            switch "hello" {
            case "":
                var v31 uint16 = 255
                _ = v31
            case "a":
                fmt.Println(v10)
            }
            _ = fi25
        }
    }
    v32 := GenS1{F0: "hello", F1: genFn0(2147483647), F2: 1}
    for fi33 := 0; fi33 < 7; fi33++ {
        for fi34 := 0; fi34 < 11; fi34++ {
            fmt.Println(v3, v18)
            _ = len(v5)
            _ = len(v5)
            _ = fi34
        }
        _ = fi33
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        v36 := GenS3{F0: (v1 + v1)}
        _ = fi35
        _ = v36
    }
    var v37 uint8 = 145
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
}
