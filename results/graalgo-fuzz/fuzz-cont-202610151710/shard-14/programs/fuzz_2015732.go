package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }
type GenS4 struct { F0 int; F1 int16; F2 uint }
func genFn0(p0 int32, p1 uint64, p2 string) int {
    if (!false) {
        if false {
            var v4 int64 = 0
            _ = v4
        } else {
            p2 = (((("a" + p2) + p2) + ((p2 + ((p2 + p2) + p2)) + "\n\t")) + ("hello" + (p2 + "hello")))
            p2 = "\n\t"
        }
        var v5 uint32 = 8
        _ = v5
    }
    p0 = (p0 & (p0 << ((319) & 63)))
    switch (p2 + "a") {
    case "a":
        v6 := GenS1{F0: 32767}
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: 773}
        _ = v7
    case "\n\t":
        var v8 float64 = (3.14159 / 1.5)
        _ = v8
    default:
        for fi9 := 0; fi9 < 13; fi9++ {
            if false {
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS1{F0: 933}
                    p0 -= (p0 % (((p0 | (p0 * (-(p0))))) | 1))
                    switch (fi10 | fi9) {
                    case -2147483648:
                        fmt.Println(fi10)
                    case 128:
                        fmt.Println(p2, p1)
                    case -1:
                        fmt.Println(p2, p0)
                    default:
                        p2 = p2
                    }
                    if false {
                        var v12 int16 = -1
                        v12 = (v12 >> ((42) & 63))
                        fmt.Println(fi9)
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
                v13 := GenS2{F0: p0, F1: p2}
                var v14 float64 = float64(((3.14159 / ((1e10 + 1.5) * -2.25)) - 1.5))
                _ = v13
                _ = v14
            }
            if (true || (42 == 1)) {
                var v15 int8 = 100
                v16 := GenS3{F0: 42, F1: -26, F2: 208}
                _ = v15
                _ = v16
            } else {
                v17 := []int{(fi9 * (^(fi9))), fi9}
                _ = v17
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := []int{}
                fmt.Println(fi9, fi9)
                p1 -= (p1 & 4294967295)
                _ = fi18
                _ = v19
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := []string{"a", "hello", "hello", "hello"}
    v2 := make(map[uint32]bool)
    var v3 uint16 = 1
    switch ("\n\t" + "a") {
    case "a":
        v20 := GenS2{F0: 241, F1: ""}
        _ = v20
    }
    v21 := v2[2147483648]
    var v22 float32 = -2.25
    if (("hello" > "\n\t") || v21) {
        delete(v2, 189)
        var v23 float64 = 0.0
        _ = v23
    } else {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        v3 -= (^(v3))
        v24 := v2[65535]
        _ = v24
    }
    fmt.Println(v21)
    v25 := GenS1{F0: 1}
    if v21 {
        delete(v2, 4294967295)
        v26 := GenS4{F0: 2147483647, F1: -280, F2: 2}
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        fmt.Println(v22, v21, v22)
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v21, v3)
            v22 = v22
            _ = fi28
        }
        fmt.Println(v21, v22)
        _ = fi27
    }
    var v29 int32 = -260
    v30 := GenS4{F0: 4294967295, F1: -767, F2: 222}
    switch 127 {
    case -9223372036854775808:
        switch int(v29) {
        case 128:
            var v31 uint16 = v3
            _ = v31
        default:
            v3 += 481
        }
    case 2147483647:
        _ = len(v1)
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        v2[42] = v21
    case 32767:
        if (false && v21) {
            fmt.Println(v22, v22)
            switch "hello" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v32 int64 = -868
                _ = v32
            }
        }
    case 4294967295:
        var v33 string = "hello"
        _ = v33
    }
    switch "" {
    case "hello":
        var v34 bool = (("" + "") != (("foo bar" + "a") + ""))
        _ = v34
    default:
        switch 42 {
        case -1:
            v3 = (v3 ^ (v3 * v3))
        case 127:
            var v35 int32 = (v29 / ((v29) | 1))
            _ = v35
        case 4294967295:
            var v36 int8 = 42
            _ = v36
        }
    }
    if v21 {
        fmt.Println(v21, v22, v29)
        if (((1 >= 0) && ((v29 & v29) <= (v29 % (((^(v29))) | 1)))) && v21) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v29
    _ = v3
    _ = v30
}
