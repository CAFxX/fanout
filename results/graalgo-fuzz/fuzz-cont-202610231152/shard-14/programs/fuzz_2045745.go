package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    v1 := []uint32{42, 45}
    var v2 string = "hello"
    var v3 string = ((("hello" + v2) + (v2 + v2)) + "foo bar")
    v4 := []string{}
    v5 := make(map[string]int)
    fmt.Println(v3)
    v6 := GenS1{F0: 4294967295}
    switch v3 {
    case "foo bar":
        v2 = ""
    }
    var v7 string = ("hello" + ((("hello" + (v2 + v3)) + v3) + (v3 + (v2 + v3))))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v7)
    switch 128 {
    case 416:
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                for fi9 := 0; fi9 < 19; fi9++ {
                    if (128 == 104) {
                        switch (fi8 % ((930) | 1)) {
                        case 42:
                            var v10 string = v3
                            _ = v10
                        case 4294967295:
                            fmt.Println(v2)
                        }
                        var v11 int32 = 1
                        v2 = v7
                        _ = v11
                    }
                    v12 := GenS2{F0: 42, F1: 0}
                    _ = fi9
                    _ = v12
                }
                v2 = (v3 + v7)
                if (!(!(!true))) {
                    if (float32(((0.0 - 1.5) * 3.14159)) >= float32((0.0 - 1.5))) {
                        v3 = "hello"
                    } else {
                        fmt.Println(v3, v3, v7)
                        var v13 float64 = 1.5
                        _ = v13
                    }
                    _ = len(v4)
                    var v14 uint = 42
                    _ = v14
                } else {
                    v7 = v2
                    var v15 int64 = -9223372036854775808
                    _ = v15
                }
            }
            v2 = "a"
            v5[("a" + "hello")] = fi8
            _ = fi8
        }
    case 866:
        var v16 int16 = -461
        _ = v16
    case 2147483647:
        var v17 uint32 = 42
        _ = v17
    }
    v3 = "foo bar"
    v18 := GenS1{F0: 342}
    var v19 uint8 = 127
    if false {
        switch v3 {
        case "\n\t":
            if false {
                if (444 > 63) {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 uint32 = 2147483647
                        v23 := v5[(v7 + v2)]
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
                fmt.Println(v19)
                fmt.Println(v3)
            } else {
                fmt.Println(v19)
            }
        case "hello":
            v4 = append(v4, (("foo bar" + v7) + (v2 + v7)))
        case "a":
            var v24 bool = (!(0.1 >= float32(1e10)))
            _ = v24
        }
        switch 9223372036854775807 {
        case -568:
            if (true && true) {
                v3 = "a"
            }
        case -320:
            v5["a"] = 42
        case 2147483647:
            switch 32768 {
            case 42:
                if (2 < 0) {
                    switch -1 {
                    case 9223372036854775807:
                        var v25 int64 = 0
                        _ = v25
                    case 127:
                        v7 = "a"
                    case 659:
                        var v26 int64 = 102
                        _ = v26
                    }
                }
            case -9223372036854775808:
                v4 = append(v4, "a")
            case -405:
                var v27 uint16 = 606
                _ = v27
            default:
                _ = len(v5)
            }
        default:
            var v28 int32 = 71
            _ = v28
        }
    }
    fmt.Println(v7, v3, v3)
    var v29 int = -1
    var v30 float64 = (float64(-2.25) * 0.0)
    fmt.Println(v7, v30)
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 uint64 = 1
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
