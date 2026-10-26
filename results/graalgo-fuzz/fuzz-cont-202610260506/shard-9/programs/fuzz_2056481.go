package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 bool; F2 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
func genFn0() uint8 {
    v7 := GenS1{F0: 1}
    var v8 int64 = -761
    v8 -= 1
    v9 := GenS2{F0: 42}
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 uint64, p1 uint16) string {
    var v10 uint64 = p0
    v11 := []int{-2147483648, 737, -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2147483647
    v4 := []int{-9223372036854775808, 1}
    var v5 int8 = -38
    var v6 int32 = 2147483647
    switch "hello" {
    case "a":
        v5 += 0
    default:
        switch ("foo bar" + "hello") {
        case "a":
            v5 = (^(v5))
        }
    }
    fmt.Println(v5, v3, v6)
    fmt.Println(v6)
    var v12 int16 = -637
    switch "hello" {
    case "foo bar":
        for fi13 := 0; fi13 < 19; fi13++ {
            v6 = v6
            v12 -= v12
            _ = fi13
        }
    case "a":
        var v14 uint = 65535
        _ = v14
    }
    fmt.Println(v6)
    v15 := GenS1{F0: 738}
    if (true && false) {
        var v16 uint16 = 32767
        v17 := GenS3{F0: 9223372036854775807}
        v5 += v5
        _ = v16
        _ = v17
    } else {
        v5 -= -1
    }
    v6 += -2147483648
    var v18 uint16 = 65535
    v18 = (v18 - (^(v18)))
    for fi19 := 0; fi19 < 15; fi19++ {
        if (3.14159 < ((-2.25 * 0.0) * 1.5)) {
            var v20 uint16 = 651
            delete(v2, v3)
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(v12, fi21)
                var v22 float64 = ((3.14159 * 0.1) + -2.25)
                _ = fi21
                _ = v22
            }
            _ = v20
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            if (true && false) {
                v24 := GenS1{F0: -1}
                var v25 uint32 = 0
                _ = v24
                _ = v25
            }
            _ = fi23
        }
        _ = fi19
    }
    if (true || true) {
        switch 32767 {
        case 313:
            var v26 uint32 = uint32(v18)
            _ = v26
        default:
            v27 := GenS4{F0: -25, F1: true, F2: (^(v5))}
            _ = v27
        }
        if (v18 >= (^(v18))) {
            if true {
                if len(v4) > 0 {
                    v4[0] = 0
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                for fi28 := 0; fi28 < 19; fi28++ {
                    fmt.Println(v6, v5, v6)
                    _ = fi28
                }
                v18 -= (v18 + 32767)
            }
            v6 = v6
            fmt.Println(v3, v5, v6)
        }
    }
    switch "hello" {
    case "\n\t":
        v29 := GenS1{F0: 970}
        _ = v29
    case "":
        for fi30 := 0; fi30 < 12; fi30++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi30
        }
    case "hello":
        v3 = v3
    }
    v31 := GenS5{F0: 127}
    switch "hello" {
    case "":
        var v32 uint64 = 519
        _ = v32
    }
    var v33 bool = (!(v6 <= (-(v6))))
    if ((^(v6)) > -1) {
        v6 = (v6 - v6)
        for fi34 := 0; fi34 < 2; fi34++ {
            v35 := GenS4{F0: 5, F1: v33, F2: (v5 - 127)}
            switch "\n\t" {
            case "\n\t":
                for fi36 := 0; fi36 < 1; fi36++ {
                    v37 := GenS6{F0: 65535}
                    v12 = (-(v12))
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi34
            _ = v35
        }
        var v38 int = 2147483647
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
