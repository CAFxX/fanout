package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int16; F2 bool }
func genFn0(p0 string, p1 int64, p2 string) int64 {
    v6 := []int{75, -2147483648, 464}
    fmt.Println(p2, p2)
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 uint64) uint64 {
    var v7 string = ""
    switch 32768 {
    case 42:
        v8 := GenS1{F0: int(p0), F1: 42}
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS2{F0: 2, F1: 1}
        v11 := GenS3{F0: 74}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int32 = -62
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int32 = (v12 << ((42) & 63))
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v12
    _ = v7
    return (p0 & (p0 & p0))
}

func main() {
    var v1 uint64 = 454
    var v2 uint16 = 320
    var v3 int = -2147483648
    v4 := []string{"\n\t", "foo bar", "a", "foo bar"}
    var v5 uint32 = 729
    var v15 int64 = genFn0(("foo bar" + "a"), 0, "")
    if true {
        v16 := GenS2{F0: uint16(v5), F1: 0}
        switch (v3 >> ((693) & 63)) {
        case -1:
            v17 := GenS4{F0: 255, F1: 2}
            _ = v17
        case 548:
            if (v15 > -1) {
                v15 += genFn0(("" + ""), int64(v5), ("\n\t" + "foo bar"))
                var v18 uint64 = genFn1((v1 >> ((65535) & 63)))
                _ = v18
            } else {
                if false {
                    v15 = v15
                } else {
                    switch (v3 | v3) {
                    case 9223372036854775807:
                        if len(v4) > 0 {
                            v4[0] = ("hello" + "")
                        }
                    case -2147483648:
                        switch (v3 * (v3 - (v3 >> ((4294967295) & 63)))) {
                        case -1:
                            fmt.Println(v3, v3)
                        }
                    default:
                        v2 = v2
                    }
                }
            }
        case 127:
            v19 := GenS1{F0: (v3 ^ v3), F1: 255}
            _ = v19
        }
        fmt.Println(v3, v2, v1)
        _ = v16
    } else {
        switch "\n\t" {
        case "":
            for fi20 := 0; fi20 < 11; fi20++ {
                fmt.Println(v5, fi20)
                _ = fi20
            }
        }
    }
    v4 = append(v4, "hello")
    fmt.Println(v2, v15, v15)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 = (v2 ^ (^(v2)))
    } else {
        if false {
            var v21 uint32 = v5
            v22 := GenS2{F0: v2, F1: 0}
            _ = v21
            _ = v22
        } else {
            if true {
                v15 = (v15 ^ -241)
                v23 := GenS4{F0: 2, F1: 0}
                _ = v23
            } else {
                var v24 int8 = -128
                _ = v24
            }
            v25 := GenS5{F0: -32768}
            _ = v25
        }
        var v26 bool = false
        v27 := GenS4{F0: 18446744073709551615, F1: 65535}
        _ = v26
        _ = v27
    }
    var v28 int = v3
    v15 = (v15 >> ((118) & 63))
    if (!false) {
        var v29 uint = 2
        v4 = append(v4, "a")
        if false {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch (v3 % (((v28 >> ((v29) & 63))) | 1)) {
                case 129:
                    v30 := GenS6{F0: ((0 >= 1000) || false), F1: int16(v15), F2: (!false)}
                    _ = v30
                }
                fmt.Println(v3, v1)
            }
        }
        _ = v29
    } else {
        _ = len(v4)
    }
    fmt.Println(v3)
    if true {
        var v31 bool = (!false)
        v32 := GenS6{F0: (!(("foo bar" + "\n\t") >= "hello")), F1: 1, F2: (false && (123 > 154))}
        if len(v4) > 0 {
            v4[0] = ""
        }
        _ = v31
        _ = v32
    }
    fmt.Println(v5, v28)
    var v33 uint16 = v2
    fmt.Println(v33, v1)
    _ = len(v4)
    switch (("" + "hello") + "foo bar") {
    case "a":
        if false {
            var v34 int8 = 42
            var v35 uint64 = (v1 * (^(v1)))
            if false {
                _ = len(v4)
            } else {
                for fi36 := 0; fi36 < 14; fi36++ {
                    v4 = append(v4, ("\n\t" + "foo bar"))
                    v4 = append(v4, "a")
                    _ = fi36
                }
                for fi37 := 0; fi37 < 19; fi37++ {
                    var v38 int8 = (^(v34))
                    v34 = (v38 % ((-60) | 1))
                    v39 := GenS2{F0: (v33 % ((v33) | 1)), F1: -2147483648}
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            _ = v34
            _ = v35
        }
    default:
        var v40 int8 = 100
        _ = v40
    }
    var v41 int32 = 1
    switch "\n\t" {
    case "hello":
        if len(v4) > 0 {
            v4[0] = ("foo bar" + ("hello" + "foo bar"))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v41
    _ = v5
}
