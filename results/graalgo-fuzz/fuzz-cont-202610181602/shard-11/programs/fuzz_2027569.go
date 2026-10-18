package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int16; F1 int }
func genFn0(p0 int16, p1 uint16) uint16 {
    var v6 bool = true
    v7 := []int{9223372036854775807, -9223372036854775808}
    var v8 int = 690
    p0 -= -32768
    for fi9 := 0; fi9 < 12; fi9++ {
        p1 = 32767
        switch v8 {
        case 2147483647:
            for fi10 := 0; fi10 < 17; fi10++ {
                v8 = 2147483647
                var v11 uint = 2
                _ = fi10
                _ = v11
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 string = (("foo bar" + ((("\n\t" + "") + "\n\t") + "\n\t")) + "")
    v2 := []int{4294967295}
    var v3 int32 = -1
    var v4 bool = false
    var v5 uint64 = 238
    switch v1 {
    case "hello":
        v12 := GenS1{F0: -9223372036854775808, F1: (v5 * (v5 - 42))}
        _ = v12
    case "\n\t":
        v1 = "foo bar"
    }
    if (1 > 255) {
        if ((((-2.25 + 0.1) / (3.14159 + 1.5)) + (-2.25 - 1e10)) >= (((0.1 - (float32((1e10 - 3.14159)) / -2.25)) + (0.0 - ((((0.1 + 0.0) - (((-2.25 / -2.25) + float32((1.5 / 1.0))) / ((0.0 + -2.25) * 1.5))) * (((-2.25 + 0.0) / 1.5) / 3.14159)) + 0.0))) * (0.1 + -2.25))) {
            var v13 int = -9223372036854775808
            v5 -= (v5 & (v5 / ((v5) | 1)))
            _ = v13
        } else {
            _ = len(v2)
            v1 = (v1 + ("a" + (v1 + v1)))
        }
        switch 540 {
        case -2147483648:
            v14 := GenS2{F0: 4294967295, F1: (v4 || (false && (-921 < 42)))}
            _ = v14
        default:
            fmt.Println(v5, v5)
        }
    } else {
        fmt.Println(v3)
        if v4 {
            v4 = v4
            v15 := GenS2{F0: 506, F1: (genFn0(-652, 626) != 65535)}
            v4 = (v4 && false)
            _ = v15
        }
        v3 -= (v3 | (^(v3)))
    }
    var v16 int64 = 992
    fmt.Println(v4, v16)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v1 {
    case "":
        switch "" {
        case "":
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(v3, v1)
                var v18 uint16 = genFn0(32767, uint16(v5))
                _ = fi17
                _ = v18
            }
        case "\n\t":
            v1 = v1
        case "foo bar":
            var v19 string = "a"
            _ = v19
        default:
            var v20 uint = 630
            _ = v20
        }
    case "foo bar":
        v21 := GenS2{F0: 0, F1: (!(v4 && (false || false)))}
        _ = v21
    case "hello":
        if (2 < 2) {
            v22 := GenS3{F0: 397, F1: 100}
            v2 = append(v2, 42)
            var v23 int16 = 32767
            _ = v22
            _ = v23
        } else {
            switch -2147483648 {
            case 379:
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 uint64 = v5
                    v26 := GenS4{F0: 750, F1: v1}
                    v1 = (("a" + ("hello" + v1)) + "hello")
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            default:
                fmt.Println(v5)
            }
            var v27 bool = (false || v4)
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 int64 = v16
            _ = fi28
            _ = v29
        }
    }
    var v30 uint = 18446744073709551615
    fmt.Println(v16, v5)
    fmt.Println(v16)
    if ((v30 >> ((2) & 63)) <= v30) {
        if len(v2) > 0 {
            v2[0] = -526
        }
        var v31 string = v1
        v31 = ("foo bar" + ((v31 + v31) + "\n\t"))
        _ = v31
    } else {
        v4 = v4
        v1 = v1
    }
    v1 = v1
    fmt.Println(v5, v1)
    var v32 uint8 = 22
    v2 = append(v2, 128)
    switch (v1 + v1) {
    case "a":
        v33 := GenS5{F0: 262, F1: -2147483648}
        _ = v33
    case "":
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 int32 = v3
            _ = fi34
            _ = v35
        }
    }
    v16 += 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
