package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 string, p1 string) int32 {
    var v5 uint = 42
    var v6 uint32 = 2147483648
    v6 = (^(v6))
    var v7 uint16 = 900
    switch p0 {
    case "":
        p1 = p1
    default:
        v6 += (^(v6))
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -812
}

func main() {
    var v1 float64 = -2.25
    var v2 int64 = 9223372036854775807
    v3 := []string{(("foo bar" + (("foo bar" + "foo bar") + ((("\n\t" + ("" + "")) + "") + ("a" + "hello")))) + "hello"), (("a" + ("\n\t" + ("" + ""))) + ("hello" + ""))}
    v4 := []int{751, -481, 255, 9223372036854775807}
    v2 += (v2 >> ((42) & 63))
    var v8 int = 2147483647
    if false {
        var v9 int = v8
        _ = v9
    } else {
        var v10 string = "a"
        _ = v10
    }
    var v11 int64 = (^(v2))
    switch v8 {
    case 127:
        for fi12 := 0; fi12 < 20; fi12++ {
            switch (fi12 >> ((368) & 63)) {
            case 2147483647:
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS1{F0: 2147483648, F1: (-94 < 43), F2: 42}
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = fi13
                    _ = v14
                }
            case 0:
                fmt.Println(fi12)
            case -354:
                v11 += v2
            }
            var v15 string = ""
            var v16 uint64 = 439
            _ = fi12
            _ = v15
            _ = v16
        }
    case 0:
        for fi17 := 0; fi17 < 6; fi17++ {
            if true {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v1 -= ((v1 * float64(0.0)) + (v1 * v1))
                case 128:
                    var v18 int = 539
                    _ = v18
                default:
                    if (255 <= 2) {
                        v2 = (^(v2))
                        v11 += v2
                        var v19 bool = false
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v21 := GenS1{F0: 559, F1: false, F2: 42}
                            _ = fi20
                            _ = v21
                        }
                        _ = len(v3)
                    }
                }
                v22 := GenS2{F0: 2}
                _ = v22
            } else {
                v3 = append(v3, (("" + (("foo bar" + "") + "hello")) + "a"))
                if (1 == 46) {
                    fmt.Println(fi17, v11, fi17)
                    v2 -= (v11 >> ((0) & 63))
                }
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                if len(v4) > 0 {
                    v4[0] = (v8 << ((4294967295) & 63))
                }
                _ = fi23
            }
            _ = fi17
        }
    case 1:
        v24 := GenS3{F0: 267, F1: (false && (!(!(!false))))}
        _ = v24
    default:
        v2 = (v11 - 302)
    }
    if len(v4) > 0 {
        v4[0] = (^(v8))
    }
    var v25 float64 = v1
    v26 := GenS4{F0: 32767, F1: 255}
    var v27 float64 = (v1 - float64(3.14159))
    fmt.Println(v27, v2)
    switch "a" {
    case "\n\t":
        if false {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v11, v25, v11)
            case "hello":
                var v28 float64 = (3.14159 + v1)
                _ = v28
            }
            v29 := GenS5{F0: 128}
            var v30 uint32 = 0
            _ = v29
            _ = v30
        } else {
            fmt.Println(v11, v8, v27)
            v25 += v27
            v1 -= (v1 + v27)
        }
    case "":
        v25 += ((1.5 - 0.1) - v27)
    case "a":
        v31 := GenS2{F0: 32767}
        _ = v31
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 15; fi33++ {
            switch 128 {
            case 32767:
                var v34 uint32 = 45
                _ = v34
            }
            var v35 uint = 38
            if len(v4) > 0 {
                v4[0] = fi32
            }
            _ = fi33
            _ = v35
        }
        if false {
            v36 := GenS1{F0: 592, F1: false, F2: 2147483648}
            if true {
                var v37 float32 = 3.14159
                _ = v37
            }
            _ = v36
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        fmt.Println(v2)
        _ = fi32
    }
    var v38 int = int(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v8
}
