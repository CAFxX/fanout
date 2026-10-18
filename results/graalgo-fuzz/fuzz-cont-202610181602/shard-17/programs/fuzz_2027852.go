package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 string }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint; F1 int8; F2 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 string }
func genFn0(p0 bool, p1 bool) uint32 {
    fmt.Println(p0, p0)
    switch ("" + "\n\t") {
    case "foo bar":
        for fi7 := 0; fi7 < 1; fi7++ {
            p1 = (false && p0)
            v8 := GenS1{F0: 255, F1: 42, F2: "a"}
            _ = fi7
            _ = v8
        }
    case "\n\t":
        p0 = (1 == 65535)
    case "":
        var v9 uint16 = 4
        _ = v9
    default:
        var v10 float32 = -2.25
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        switch (fi11 * 9223372036854775807) {
        case -2147483648:
            p0 = true
        }
        v12 := GenS2{F0: true, F1: -974}
        _ = fi11
        _ = v12
    }
    p1 = true
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 uint8 = 1
    var v2 int = 0
    var v3 float64 = float64((-2.25 / float32(0.1)))
    v4 := []int{(^(v2)), v2}
    var v5 uint64 = 1
    var v6 int16 = -32768
    v13 := GenS1{F0: v1, F1: 448, F2: ""}
    if (-761 == (-(v2))) {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS2{F0: (-1 >= -1), F1: 9223372036854775807}
            _ = fi14
            _ = v15
        }
    }
    if true {
        if true {
            var v16 string = "\n\t"
            _ = v16
        } else {
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := GenS3{F0: 42, F1: 100, F2: 1}
                _ = fi17
                _ = v18
            }
            v6 = (v6 % ((1) | 1))
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := GenS2{F0: true, F1: -9223372036854775808}
            _ = fi19
            _ = v20
        }
    } else {
        v6 = 1000
    }
    _ = len(v4)
    for fi21 := 0; fi21 < 14; fi21++ {
        switch "foo bar" {
        case "":
            v22 := GenS4{F0: -36, F1: (v5 * v5), F2: "hello"}
            _ = v22
        }
        if false {
            var v23 uint32 = 2
            _ = len(v4)
            v2 += (fi21 | (-(v2)))
            _ = v23
        }
        var v24 uint = 348
        _ = fi21
        _ = v24
    }
    var v25 uint16 = 65535
    switch "hello" {
    case "hello":
        v2 = (v2 | (v2 - 128))
    case "a":
        v26 := GenS1{F0: (^(v1)), F1: 777, F2: (("foo bar" + "a") + "foo bar")}
        _ = v26
    default:
        v2 = -12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = v3
    _ = len(v4)
    fmt.Println(v6, v6, v25)
    switch 233 {
    case 4294967295:
        v2 += v2
    default:
        fmt.Println(v2)
    }
    v3 += v3
    var v27 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
