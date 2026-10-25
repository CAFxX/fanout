package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0(p0 uint16, p1 int16, p2 int16) int8 {
    p1 = p2
    v4 := []int{127, 9223372036854775807, 1, -2147483648}
    v4 = append(v4, 32768)
    switch ("\n\t" + "a") {
    case "":
        var v5 float32 = 1.5
        _ = v5
    case "\n\t":
        var v6 bool = true
        _ = v6
    case "a":
        var v7 int8 = -1
        _ = v7
    default:
        var v8 int32 = -2147483648
        _ = v8
    }
    _ = len(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 39
}
func genFn1(p0 int8, p1 int64, p2 uint64) bool {
    fmt.Println(p1)
    p0 = p0
    if true {
        var v9 uint64 = (^(p2))
        var v10 uint64 = p2
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 32767, F1: 771}
        if true {
            p1 -= (p1 ^ p1)
            p2 -= 0
            switch "a" {
            case "a":
                p2 += p2
            }
        }
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        p0 += (p0 >> ((65535) & 63))
        v13 := []int{fi12}
        v14 := GenS1{F0: 0, F1: 63}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return (false && false)
}

func main() {
    v1 := []string{"foo bar", ("a" + "foo bar")}
    var v2 int16 = -414
    var v3 uint16 = 1
    var v15 int32 = 853
    var v16 int8 = -1
    _ = len(v1)
    switch 911 {
    case 882:
        switch "hello" {
        case "hello":
            var v17 uint64 = 0
            _ = v17
        case "a":
            v18 := GenS1{F0: (v3 ^ (v3 << ((299) & 63))), F1: 665}
            _ = v18
        case "":
            v2 -= v2
        }
    }
    var v19 int = 812
    v16 -= genFn0((v3 << ((1) & 63)), 42, 0)
    if len(v1) > 0 {
        v1[0] = ((("" + "\n\t") + "hello") + "hello")
    }
    v2 -= 1000
    v15 = -705
    for fi20 := 0; fi20 < 18; fi20++ {
        v15 += (-(v15))
        v2 -= -32768
        _ = fi20
    }
    var v21 uint64 = 0
    if (true || (!(-2.25 > 3.14159))) {
        switch (v19 % ((v19) | 1)) {
        case 9223372036854775807:
            v22 := GenS1{F0: (v3 / ((v3) | 1)), F1: 2}
            _ = v22
        case 645:
            var v23 uint8 = 255
            _ = v23
        case -268:
            fmt.Println(v3, v19, v2)
        default:
            if (true || (v15 <= (v15 << ((1) & 63)))) {
                fmt.Println(v16, v3, v19)
                fmt.Println(v15, v16)
                v3 += v3
            }
        }
        v19 = 127
        if (0 < 1) {
            v24 := GenS1{F0: v3, F1: 105}
            if len(v1) > 0 {
                v1[0] = ("foo bar" + "foo bar")
            }
            switch (v19 << ((uint(v3)) & 63)) {
            case -1:
                var v25 int8 = genFn0(42, 1000, v2)
                _ = v25
            case 128:
                var v26 uint = 255
                _ = v26
            }
            _ = v24
        } else {
            if true {
                var v27 float32 = (((0.0 / ((float32(1.5) * 3.14159) / ((((3.14159 + 0.1) + float32(0.0)) + ((3.14159 - 3.14159) - -2.25)) * 1e10))) / float32(0.1)) + 0.1)
                _ = v27
            }
            var v28 float32 = 0.0
            var v29 uint16 = v3
            _ = v28
            _ = v29
        }
    }
    var v30 int8 = (-(v16))
    v31 := GenS1{F0: v3, F1: 1}
    for fi32 := 0; fi32 < 19; fi32++ {
        if false {
            var v33 uint32 = 2147483647
            for fi34 := 0; fi34 < 5; fi34++ {
                if len(v1) > 0 {
                    v1[0] = ("foo bar" + "a")
                }
                _ = fi34
            }
            _ = v33
        } else {
            var v35 int64 = -9223372036854775808
            fmt.Println(v30, v2)
            _ = v35
        }
        v15 += -1
        _ = fi32
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
}
