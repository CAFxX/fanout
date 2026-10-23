package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int; F1 uint8; F2 int64 }
type GenS4 struct { F0 int64; F1 uint; F2 uint8 }
type GenS5 struct { F0 int64; F1 bool; F2 string }
func genFn0() int8 {
    for fi5 := 0; fi5 < 7; fi5++ {
        if true {
            fmt.Println(fi5)
        }
        v6 := []int{127, (fi5 - fi5)}
        var v7 int16 = -32768
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = 0
    var v9 uint = 0
    _ = v8
    _ = v9
    return -107
}

func main() {
    var v1 int64 = 42
    v2 := []string{}
    var v3 uint8 = 0
    v4 := []int{0}
    var v10 float64 = (0.1 * -2.25)
    fmt.Println(v3)
    var v11 int16 = 1
    if false {
        var v12 uint64 = 2
        if (!(!(!false))) {
            v3 -= v3
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS1{F0: uint16(v3), F1: fi13, F2: 0}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi13
                _ = v14
            }
        } else {
            var v15 uint32 = 173
            switch -9223372036854775808 {
            case 9223372036854775807:
                fmt.Println(v1, v3, v1)
            }
            v16 := GenS2{F0: 275, F1: 0}
            _ = v15
            _ = v16
        }
        _ = v12
    } else {
        v11 += v11
        var v17 int = 32768
        v17 -= v17
        _ = v17
    }
    if false {
        var v18 uint64 = 42
        v18 -= 184
        if true {
            for fi19 := 0; fi19 < 3; fi19++ {
                v3 = (v3 & (v3 % ((v3) | 1)))
                switch fi19 {
                case 1:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case -1:
                    v20 := GenS3{F0: (-(fi19)), F1: (^(v3)), F2: -1}
                    _ = v20
                case 128:
                    v21 := GenS4{F0: (^(v1)), F1: 65535, F2: (^(v3))}
                    _ = v21
                }
                _ = fi19
            }
        }
        _ = v18
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int32 = -447
        if len(v2) > 0 {
            v2[0] = ("foo bar" + ("" + ("\n\t" + "\n\t")))
        }
        var v24 uint16 = 42
        _ = fi22
        _ = v23
        _ = v24
    }
    switch 0 {
    case -628:
        var v25 uint8 = v3
        _ = v25
    default:
        var v26 int16 = v11
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 1; fi28++ {
            if (!(int64(v11) < -9223372036854775808)) {
                fmt.Println(v10, v3, fi28)
                var v29 bool = (!("a" > "\n\t"))
                _ = v29
            }
            var v30 float32 = 1.5
            v31 := GenS2{F0: (v11 << ((42) & 63)), F1: (fi27 | fi28)}
            _ = fi28
            _ = v30
            _ = v31
        }
        _ = fi27
    }
    switch "hello" {
    case "":
        var v32 float32 = 1.5
        _ = v32
    }
    _ = len(v4)
    for fi33 := 0; fi33 < 7; fi33++ {
        for fi34 := 0; fi34 < 5; fi34++ {
            fmt.Println(fi34, fi34)
            fmt.Println(v10, v10, v3)
            _ = fi34
        }
        fmt.Println(v3, v10)
        v35 := GenS5{F0: v1, F1: (!false), F2: "\n\t"}
        _ = fi33
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
