package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint, p1 uint, p2 int32) int8 {
    p2 -= -2147483648
    if false {
        for fi4 := 0; fi4 < 9; fi4++ {
            for fi5 := 0; fi5 < 3; fi5++ {
                switch ("foo bar" + "\n\t") {
                case "hello":
                    var v6 int16 = int16(p2)
                    _ = v6
                case "\n\t":
                    if false {
                        p2 -= -676
                        var v7 int32 = 861
                        _ = v7
                    }
                case "a":
                    v8 := []int{(fi5 | (fi4 - (fi4 * 0))), (fi5 + (-(fi5)))}
                    _ = v8
                }
                _ = fi5
            }
            switch "" {
            case "hello":
                var v9 int16 = 352
                _ = v9
            default:
                if true {
                    v10 := []int{fi4, int(p0), fi4, (fi4 & fi4)}
                    fmt.Println(p2, p1, p2)
                    var v11 uint16 = 254
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v13 := []int{(fi4 >> ((p0) & 63)), fi12}
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 9; fi14++ {
                        fmt.Println(p1, p1, p1)
                        _ = fi14
                    }
                }
            }
            switch int(p2) {
            case -771:
                fmt.Println(p1)
            }
            _ = fi4
        }
        switch ("hello" + ("foo bar" + "\n\t")) {
        case "hello":
            v15 := GenS1{F0: -32768, F1: -1, F2: -1}
            _ = v15
        case "":
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 uint32 = 2
                _ = fi16
                _ = v17
            }
        default:
            var v18 uint = p0
            _ = v18
        }
    } else {
        p0 += 517
    }
    switch 9223372036854775807 {
    case -627:
        p2 -= (p2 % ((p2) | 1))
    case -890:
        v19 := []int{-881, 0, 32767}
        _ = v19
    case -9223372036854775808:
        for fi20 := 0; fi20 < 20; fi20++ {
            fmt.Println(p1, fi20)
            p1 -= (p0 & p1)
            _ = fi20
        }
    default:
        p1 = p1
    }
    p1 -= (p0 | p0)
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := []int{4294967295}
            if len(v23) > 0 {
                v23[0] = 127
            }
            if (!true) {
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                p0 = (p1 << ((p0) & 63))
                v24 := GenS2{F0: (p2 - (p2 * (p2 + -2147483648))), F1: 963}
                _ = v24
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = p2
    return 100
}
func genFn1(p0 string) uint64 {
    var v25 uint8 = 91
    fmt.Println(p0)
    _ = p0
    _ = v25
    return 42
}

func main() {
    var v1 int64 = 481
    var v2 string = ("a" + ((("" + ((("\n\t" + ("\n\t" + "a")) + "a") + "a")) + "\n\t") + (("foo bar" + ("foo bar" + "a")) + "foo bar")))
    var v3 string = v2
    var v26 int32 = 42
    v27 := GenS1{F0: 32767, F1: 100, F2: -32768}
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 uint = 0
        var v30 uint64 = 18446744073709551615
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := []int{-2147483648}
    v3 = v2
    v32 := GenS2{F0: (v26 - v26), F1: 32767}
    v3 = (("hello" + "") + "")
    var v33 bool = (genFn0(42, 2, v26) <= -1)
    fmt.Println(v33, v3, v26)
    var v34 uint64 = 0
    v34 = v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
