package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
func genFn0(p0 int16, p1 uint8) uint64 {
    v4 := []int{1, 9223372036854775807, int(p1)}
    for fi5 := 0; fi5 < 18; fi5++ {
        if ((-112 >= 127) && false) {
            fmt.Println(p1, fi5)
        } else {
            var v6 uint32 = 65535
            switch (fi5 / (((fi5 * fi5)) | 1)) {
            case 4294967295:
                v7 := GenS1{F0: (v6 & (v6 / ((v6) | 1)))}
                _ = v7
            }
            v8 := GenS1{F0: v6}
            _ = v6
            _ = v8
        }
        var v9 int = (fi5 | fi5)
        _ = fi5
        _ = v9
    }
    var v10 int8 = 21
    var v11 float32 = 1.5
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v4
    return 42
}
func genFn1(p0 int64, p1 int32, p2 uint16) uint {
    var v12 float32 = 1e10
    v13 := []int{}
    switch -687 {
    case 9223372036854775807:
        p0 = p0
    default:
        var v14 int = 42
        _ = v14
    }
    _ = len(v13)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 71
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 2
    v3 := []string{"a", "a", ("foo bar" + "")}
    var v15 uint16 = 32767
    switch ("\n\t" + "\n\t") {
    case "a":
        if true {
            var v16 uint32 = 2147483648
            switch "foo bar" {
            case "\n\t":
                switch -1 {
                case 319:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 32767:
                    fmt.Println(v16, v2, v16)
                case -9223372036854775808:
                    v16 += (v16 & v16)
                default:
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 int32 = 42
                        _ = fi17
                        _ = v18
                    }
                }
            case "foo bar":
                fmt.Println(v1, v16)
            case "a":
                v2 = (v2 >> ((genFn1(-610, 1, v15)) & 63))
            }
            _ = v16
        }
    case "hello":
        if (!(!((v2 / ((v2) | 1)) < (v2 & 2)))) {
            v3 = append(v3, ("" + "a"))
        }
    case "\n\t":
        if (v1 < (^(v1))) {
            switch -2147483648 {
            case 810:
                v19 := GenS1{F0: 466}
                _ = v19
            case 290:
                v15 += v15
            case 1:
                if (!false) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v2, v15, v15)
                    fmt.Println(v2, v1)
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v3 = append(v3, ("foo bar" + "a"))
            var v20 bool = false
            _ = v20
        }
    }
    var v21 int64 = 42
    if false {
        v15 = 0
        _ = len(v3)
    }
    v3 = append(v3, "\n\t")
    var v22 string = ""
    v23 := GenS2{F0: 1, F1: 90, F2: 1}
    v15 -= (v15 & 65535)
    var v24 uint64 = 1
    var v25 bool = (((0 <= uint32(v1)) || (v22 < (v22 + v22))) || true)
    var v26 uint32 = 852
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
