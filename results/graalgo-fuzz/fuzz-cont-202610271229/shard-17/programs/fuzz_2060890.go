package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 17; fi6++ {
        switch ("foo bar" + "hello") {
        case "hello":
            var v7 string = "a"
            _ = v7
        default:
            var v8 int32 = -573
            _ = v8
        }
        switch "foo bar" {
        case "\n\t":
            if true {
                if (32767 >= 32767) {
                    switch "a" {
                    case "foo bar":
                        var v9 float32 = 1.5
                        _ = v9
                    }
                    if false {
                        var v10 string = ("hello" + (((("" + "a") + "hello") + "a") + "\n\t"))
                        v10 = v10
                        _ = v10
                    } else {
                        var v11 int16 = -32768
                        _ = v11
                    }
                    fmt.Println(fi6)
                } else {
                    v12 := GenS1{F0: 1}
                    var v13 float32 = float32(0.1)
                    _ = v12
                    _ = v13
                }
                v14 := GenS1{F0: 1}
                _ = v14
            }
        case "":
            fmt.Println(fi6)
        case "a":
            var v15 int = fi6
            _ = v15
        }
        v16 := GenS2{F0: fi6}
        _ = fi6
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 float64 = 1e10
        _ = fi17
        _ = v18
    }
    v19 := []int{4294967295}
    _ = v19
    return 969
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{"", "\n\t"}
    var v3 uint32 = 371
    var v4 uint64 = uint64(v3)
    var v5 uint32 = 970
    for fi20 := 0; fi20 < 5; fi20++ {
        v5 -= v5
        var v21 int8 = -35
        _ = fi20
        _ = v21
    }
    fmt.Println(v3)
    v2 = append(v2, (("foo bar" + "hello") + ("" + "hello")))
    var v22 int64 = (-(v1))
    fmt.Println(v4, v22, v1)
    var v23 bool = true
    v22 = (v22 | v22)
    v24 := GenS3{F0: 2}
    switch 9223372036854775807 {
    case 1:
        fmt.Println(v4)
    default:
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 string = ""
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (v1 > -9223372036854775808) {
                fmt.Println(v5, v4, v1)
            }
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 9; fi28++ {
            fmt.Println(v3, v3)
            _ = fi28
        }
        v3 = uint32(v4)
        _ = fi27
    }
    switch "a" {
    case "\n\t":
        v1 -= (-(v22))
    case "":
        v29 := GenS1{F0: 1}
        _ = v29
    case "hello":
        fmt.Println(v4)
    default:
        switch -352 {
        case 127:
            var v30 uint16 = 2
            _ = v30
        default:
            var v31 float64 = (1e10 - (0.1 * (3.14159 + (1e10 - 0.1))))
            _ = v31
        }
    }
    var v32 int64 = (v1 << ((uint(v4)) & 63))
    var v33 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
