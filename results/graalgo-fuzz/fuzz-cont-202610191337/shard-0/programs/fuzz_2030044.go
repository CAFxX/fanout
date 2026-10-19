package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint, p1 int16) string {
    v5 := []int{9223372036854775807, -9223372036854775808}
    for fi6 := 0; fi6 < 13; fi6++ {
        if (!false) {
            p0 = p0
            var v7 int32 = 1
            _ = v7
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            if false {
                p1 = p1
            } else {
                var v9 bool = true
                switch "a" {
                case "hello":
                    var v10 bool = (!v9)
                    _ = v10
                case "":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v9 = v9
                }
                var v11 int64 = -1
                _ = v11
                _ = v9
            }
            if len(v5) > 0 {
                v5[0] = fi6
            }
            _ = fi8
        }
        _ = fi6
    }
    var v12 int8 = 42
    switch -1 {
    case 2147483647:
        fmt.Println(p0)
    case 0:
        v13 := GenS1{F0: v12, F1: 42}
        _ = v13
    }
    v14 := GenS1{F0: v12, F1: 0}
    _ = p0
    _ = p1
    _ = v12
    _ = v14
    _ = v5
    return "\n\t"
}

func main() {
    v1 := []uint32{2147483647, 2147483648}
    v2 := []uint32{487, 65535, 65535, 747}
    var v3 uint32 = 74
    var v4 uint16 = 415
    v1 = append(v1, 269)
    for fi15 := 0; fi15 < 14; fi15++ {
        if true {
            if (false && true) {
                v4 = v4
            }
            if false {
                v16 := GenS1{F0: -1, F1: 0}
                var v17 uint = 18446744073709551615
                fmt.Println(v17, v17)
                _ = v16
                _ = v17
            } else {
                v2 = append(v2, (v3 * 2))
            }
        }
        _ = fi15
    }
    var v18 int32 = 703
    if false {
        switch "hello" {
        case "foo bar":
            v19 := GenS2{F0: true}
            _ = v19
        case "\n\t":
            fmt.Println(v18)
        }
    } else {
        switch ((("\n\t" + ((("\n\t" + "hello") + (("\n\t" + "\n\t") + "foo bar")) + "hello")) + ("a" + "foo bar")) + "") {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            v20 := GenS2{F0: ((42 != 1) && true)}
            _ = v20
        case "hello":
            v3 = (v3 % ((v3) | 1))
        default:
            if false {
                v21 := GenS1{F0: -120, F1: v4}
                _ = v21
            }
        }
    }
    var v22 uint = 4294967295
    v23 := GenS1{F0: 18, F1: (v4 + (^(v4)))}
    v4 -= 65535
    fmt.Println(v22, v22, v22)
    var v24 string = ("\n\t" + (((("\n\t" + ("" + ((((("hello" + "hello") + "a") + ("a" + "hello")) + "a") + ""))) + "foo bar") + ("a" + "\n\t")) + "a"))
    var v25 bool = true
    if (v24 >= "a") {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 int32 = -2147483648
            _ = fi26
            _ = v27
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v28 uint32 = v3
        v3 = v28
        _ = v28
    }
    v24 = "a"
    var v29 string = (v24 + ("hello" + v24))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
