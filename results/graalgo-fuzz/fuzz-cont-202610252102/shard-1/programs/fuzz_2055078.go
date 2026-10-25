package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }

func main() {
    var v1 int32 = -689
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "\n\t"
        var v7 int16 = 1
        fmt.Println(v1, v2)
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = (59 != 55)
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{(fi10 | (fi9 >> ((478) & 63)))}
            fmt.Println(v1)
            if v8 {
                v2 = v2
            } else {
                v3[(v2 - v2)] = v8
                fmt.Println(v2, v1)
            }
            v2 = (v2 & (v2 * 686))
            fmt.Println(fi9, fi10)
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            if (33 == 2) {
                v13 := GenS1{F0: fi9, F1: -1, F2: fi9}
                var v14 int8 = 42
                v15 := []int{fi12, 325, (fi12 ^ (fi12 - (-(fi12)))), (fi9 - (fi12 | fi9))}
                _ = v13
                _ = v14
                _ = v15
            } else {
                if ((2 >= 110) || v8) {
                    v1 = (^(v1))
                    _ = len(v3)
                    if (18446744073709551615 == 4294967295) {
                        v16 := []int{}
                        _ = v16
                    } else {
                        v17 := []int{}
                        v18 := v4[((("" + "foo bar") + "") + (("foo bar" + "") + ("hello" + ("\n\t" + ""))))]
                        _ = v17
                        _ = v18
                    }
                    switch "\n\t" {
                    case "\n\t":
                        var v19 int16 = 1
                        _ = v19
                    case "hello":
                        v20 := []int{fi9}
                        _ = v20
                    case "a":
                        var v21 uint8 = 2
                        switch "a" {
                        case "a":
                            var v22 int16 = -726
                            _ = v22
                        case "foo bar":
                            fmt.Println(v21, v1, v8)
                        case "hello":
                            v23 := GenS1{F0: fi9, F1: v1, F2: fi9}
                            _ = v23
                        default:
                            delete(v3, (v2 >> ((0) & 63)))
                        }
                        fmt.Println(v21)
                        _ = v21
                    }
                } else {
                    fmt.Println(fi12)
                    v24 := GenS2{F0: 42}
                    _ = v24
                }
                var v25 int32 = (v1 >> ((65535) & 63))
                var v26 float64 = (((3.14159 * (0.0 / 1.0)) + 3.14159) - 0.1)
                _ = v25
                _ = v26
            }
            v3[434] = (2 != 65535)
            var v27 string = "hello"
            if false {
                fmt.Println(v1, v8, v8)
                fmt.Println(v1, v2, v2)
                var v28 float32 = (-2.25 / -2.25)
                _ = v28
            }
            switch "hello" {
            case "a":
                v2 -= (v2 - 4294967295)
            default:
                var v29 uint64 = 18446744073709551615
                _ = v29
            }
            _ = fi12
            _ = v27
        }
        _ = fi9
    }
    switch (("foo bar" + "a") + "hello") {
    case "\n\t":
        fmt.Println(v2, v8)
    }
    var v30 string = ("" + ("foo bar" + "hello"))
    var v31 int = 4294967295
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := []int{v31, 4294967295, (fi32 & v31), 1}
        _ = fi32
        _ = v33
    }
    if v8 {
        var v34 uint8 = 255
        v35 := []int{v31, v31, v31}
        switch (v30 + v30) {
        case "":
            var v36 uint32 = 427
            _ = v36
        case "foo bar":
            var v37 float64 = 3.14159
            _ = v37
        default:
            switch (v30 + v30) {
            case "hello":
                var v38 uint = 359
                _ = v38
            case "\n\t":
                v34 = (v34 ^ (^(v34)))
            case "":
                var v39 int32 = (v1 << ((198) & 63))
                _ = v39
            }
        }
        _ = v34
        _ = v35
    }
    switch (v31 ^ v31) {
    case -103:
        v31 = 457
    case 2147483647:
        var v40 float32 = -2.25
        _ = v40
    case 706:
        for fi41 := 0; fi41 < 7; fi41++ {
            v31 = fi41
            _ = fi41
        }
    default:
        switch v30 {
        case "a":
            var v42 uint32 = v2
            _ = v42
        }
    }
    fmt.Println(v31, v2, v30)
    if v8 {
        _ = len(v3)
        var v43 float32 = 3.14159
        _ = v43
    }
    fmt.Println(v31)
    v44 := GenS3{F0: ((v30 + v30) + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v44
    _ = v8
}
