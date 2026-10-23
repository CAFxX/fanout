package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint {
    var v7 string = "hello"
    switch "foo bar" {
    case "hello":
        switch "" {
        case "a":
            fmt.Println(v7)
        case "\n\t":
            for fi8 := 0; fi8 < 18; fi8++ {
                v7 = "a"
                v9 := []int{(fi8 / (((-(fi8))) | 1)), fi8, 1}
                _ = fi8
                _ = v9
            }
        case "hello":
            for fi10 := 0; fi10 < 15; fi10++ {
                for fi11 := 0; fi11 < 15; fi11++ {
                    v7 = v7
                    var v12 int16 = 1000
                    if (625 != 194) {
                        if (9223372036854775807 <= 856) {
                            var v13 int32 = 264
                            v7 = v7
                            _ = v13
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 uint64 = 18446744073709551615
                    var v16 int64 = 1
                    var v17 uint16 = 0
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    v19 := GenS1{F0: false}
                    v7 = "\n\t"
                    _ = fi18
                    _ = v19
                }
                _ = fi10
            }
        default:
            switch "foo bar" {
            case "":
                switch -9223372036854775808 {
                case 0:
                    v7 = (v7 + "foo bar")
                case 102:
                    v20 := []int{126, 32767, 4294967295, -2147483648}
                    _ = v20
                case 32767:
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v7 = "\n\t"
                        _ = fi21
                    }
                }
            case "\n\t":
                v22 := []int{127, 0}
                _ = v22
            case "foo bar":
                v23 := []int{9223372036854775807, -384}
                _ = v23
            default:
                v7 = v7
            }
        }
    default:
        if true {
            var v24 int = 4294967295
            _ = v24
        }
    }
    _ = v7
    return 2
}

func main() {
    var v1 int64 = 0
    var v2 uint8 = 25
    v3 := make(map[string]bool)
    var v4 uint = 18446744073709551615
    var v5 int64 = 1
    var v6 int8 = 32
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := GenS1{F0: true}
        v1 = (v5 - (v5 ^ int64(fi25)))
        _ = fi25
        _ = v26
    }
    var v27 float64 = 1e10
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := []int{(-(fi28)), 4294967295}
        _ = fi28
        _ = v29
    }
    var v30 int16 = 847
    v30 += 1000
    fmt.Println(v27, v5, v4)
    var v31 uint32 = 2147483648
    var v32 float32 = -2.25
    var v33 string = "foo bar"
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 16; fi35++ {
            fmt.Println(v5, v6)
            _ = fi35
        }
        _ = fi34
    }
    v36 := []int{}
    _ = len(v3)
    if true {
        v31 += v31
    } else {
        switch "foo bar" {
        case "\n\t":
            v5 = -1
        case "hello":
            var v37 uint32 = (^(v31))
            _ = v37
        case "":
            if false {
                var v38 string = "\n\t"
                var v39 uint = v4
                var v40 uint32 = (v31 >> ((v39) & 63))
                _ = v38
                _ = v39
                _ = v40
            }
        }
        v6 += -128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
