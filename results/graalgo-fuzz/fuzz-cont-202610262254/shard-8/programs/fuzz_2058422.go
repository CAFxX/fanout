package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0(p0 int8, p1 uint16, p2 uint16) int {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint8 = 78
        fmt.Println(fi5, p1, p1)
        _ = fi5
        _ = v6
    }
    switch 0 {
    case 1:
        p2 = (^(p2))
    default:
        var v7 uint64 = 563
        _ = v7
    }
    v8 := []int{-202, -927, -58}
    fmt.Println(p1, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 139
}
func genFn1(p0 uint16, p1 uint) bool {
    if (0 > -223) {
        v9 := GenS1{F0: 975, F1: ("a" + "hello")}
        _ = v9
    }
    p1 += p1
    switch ("hello" + ("" + "\n\t")) {
    case "":
        switch 32768 {
        case 128:
            switch 159 {
            case 127:
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(p1)
                    var v11 uint8 = 60
                    v12 := GenS2{F0: (v11 % ((v11) | 1)), F1: -1, F2: 248}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            default:
                v13 := []int{2147483647, 239, genFn0(100, 32767, p0)}
                _ = v13
            }
        case 194:
            var v14 uint8 = 0
            _ = v14
        default:
            p1 = (p1 / ((p1) | 1))
        }
    }
    _ = p0
    _ = p1
    return ((p0 + (p0 << ((p1) & 63))) >= (p0 - 739))
}

func main() {
    var v1 int8 = 22
    var v2 uint = 18446744073709551615
    var v3 uint = v2
    v4 := []int{}
    if false {
        v3 = 255
        var v15 int8 = v1
        _ = v15
    } else {
        switch 32767 {
        case -1:
            for fi16 := 0; fi16 < 16; fi16++ {
                fmt.Println(v2, v1, v1)
                _ = fi16
            }
        case -610:
            v17 := GenS1{F0: 506, F1: ""}
            _ = v17
        default:
            v2 -= (v3 << (((v2 << ((v3) & 63))) & 63))
        }
        v18 := GenS1{F0: 925, F1: ("" + ((("" + ("foo bar" + ("\n\t" + ("a" + ("" + ""))))) + ("foo bar" + "hello")) + ("a" + "hello")))}
        fmt.Println(v3)
        _ = v18
    }
    var v19 bool = (false && false)
    for fi20 := 0; fi20 < 17; fi20++ {
        v19 = v19
        _ = fi20
    }
    var v21 uint32 = 943
    var v22 int64 = 1
    v23 := GenS3{F0: 1, F1: 1, F2: 65535}
    v24 := GenS1{F0: 2, F1: "a"}
    var v25 int32 = 42
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint64 = uint64(v2)
        _ = fi26
        _ = v27
    }
    v19 = (!v19)
    switch 9223372036854775807 {
    case 4294967295:
        if false {
            fmt.Println(v1, v21, v21)
            v28 := GenS3{F0: 0, F1: 128, F2: 0}
            v4 = append(v4, 42)
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        switch "a" {
        case "foo bar":
            switch fi29 {
            case 128:
                fmt.Println(v3)
            default:
                var v30 int64 = (v22 >> ((v3) & 63))
                _ = v30
            }
        }
        v19 = v19
        v31 := GenS1{F0: 868, F1: (("" + "hello") + "hello")}
        _ = fi29
        _ = v31
    }
    var v32 int64 = (v22 >> ((v3) & 63))
    v2 = (^(v3))
    if (32767 > -747) {
        var v33 string = "a"
        _ = v33
    } else {
        var v34 int16 = 32767
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
}
