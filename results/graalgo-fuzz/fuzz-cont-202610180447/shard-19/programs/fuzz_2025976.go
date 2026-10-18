package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint; F2 int8 }
func genFn0(p0 uint8, p1 int) uint {
    v5 := GenS1{F0: 464}
    for fi6 := 0; fi6 < 7; fi6++ {
        switch (p1 << ((255) & 63)) {
        case 0:
            v7 := []int{}
            _ = v7
        case 128:
            var v8 uint16 = 255
            _ = v8
        default:
            var v9 bool = true
            _ = v9
        }
        v10 := []int{(p1 | (fi6 - 32767)), (fi6 & 910), 9223372036854775807, p1}
        _ = fi6
        _ = v10
    }
    v11 := []int{}
    if true {
        p0 = p0
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    } else {
        fmt.Println(p1)
    }
    if true {
        p1 = p1
        switch (p1 * -267) {
        case 4294967295:
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 bool = false
                v14 := GenS2{F0: (p1 | int(p0)), F1: (^(p0)), F2: 0}
                var v15 string = "hello"
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        case 42:
            p1 = (p1 * 449)
        case 32768:
            for fi16 := 0; fi16 < 19; fi16++ {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 uint64 = 42
                    var v19 uint16 = 65535
                    switch "hello" {
                    case "foo bar":
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case "hello":
                        var v20 int = fi16
                        _ = v20
                    case "a":
                        v21 := GenS3{F0: (!(((false || (true || false)) && ((v18 - 4294967295) >= (^(v18)))) || true)), F1: 0, F2: -1}
                        _ = v21
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 295
}
func genFn1() uint8 {
    v22 := GenS3{F0: false, F1: 0, F2: -425}
    var v23 string = ""
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(fi24, v23)
        v25 := []int{(fi24 >> ((255) & 63)), fi24, -1, (fi24 | fi24)}
        _ = fi24
        _ = v25
    }
    _ = v22
    _ = v23
    return 127
}

func main() {
    var v1 int8 = -30
    v2 := []int{-9223372036854775808, 32767}
    var v3 int64 = 0
    var v4 int8 = -128
    var v26 uint8 = 1
    if true {
        if false {
            var v27 uint8 = 128
            v1 += v1
            _ = v27
        }
    }
    if true {
        v28 := GenS4{F0: -32768, F1: 2147483647}
        _ = v28
    }
    v1 -= (^(v4))
    if true {
        v4 = v1
    }
    var v29 uint16 = 0
    v2 = append(v2, -894)
    v30 := GenS3{F0: ((true && (!true)) && ((998 != 42) && (!(!("hello" > ("a" + (("hello" + "\n\t") + "foo bar"))))))), F1: 1000, F2: 42}
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 uint64 = 394
        var v33 float64 = (1e10 - (0.1 - 0.1))
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 11; fi35++ {
            if (!(18446744073709551615 < 42)) {
                v36 := GenS5{F0: 212, F1: genFn0(128, (fi34 * (-(fi35)))), F2: (v4 & 100)}
                v37 := GenS1{F0: 420}
                v3 = v3
                fmt.Println(v4)
                _ = v36
                _ = v37
            }
            _ = fi35
        }
        var v38 string = "hello"
        fmt.Println(v4, v29)
        _ = fi34
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
