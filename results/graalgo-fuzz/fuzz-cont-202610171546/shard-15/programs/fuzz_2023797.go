package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint8, p1 bool) uint32 {
    switch 127 {
    case -823:
        switch ("foo bar" + "a") {
        case "\n\t":
            fmt.Println(p0, p1)
        }
    case 47:
        if (32767 == -660) {
            v5 := GenS1{F0: 813}
            var v6 uint32 = 200
            _ = v5
            _ = v6
        }
    case -9223372036854775808:
        v7 := GenS2{F0: ("" + ("foo bar" + ""))}
        _ = v7
    }
    v8 := []int{42, 2147483647, 32767, 490}
    if false {
        fmt.Println(p0)
    }
    switch 42 {
    case 499:
        p0 = 2
    case 738:
        if len(v8) > 0 {
            v8[0] = -973
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 uint16) uint64 {
    if (false && true) {
        switch ("" + "") {
        case "foo bar":
            p0 = (p0 | (^(p0)))
        default:
            switch -393 {
            case 32768:
                if (((("a" + ((("" + "") + ("\n\t" + ("hello" + "a"))) + "\n\t")) + ("foo bar" + "foo bar")) + "hello") != ((("\n\t" + ("a" + "foo bar")) + (("\n\t" + "hello") + ("\n\t" + "foo bar"))) + "\n\t")) {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        p0 += (p0 / ((p0) | 1))
                        var v10 uint64 = 42
                        _ = fi9
                        _ = v10
                    }
                    v11 := []int{-1, -743, -2147483648}
                    _ = v11
                }
            case 0:
                var v12 uint8 = 42
                _ = v12
            }
        }
    }
    if true {
        fmt.Println(p0)
    } else {
        p0 = p0
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 int16 = -1
    var v4 string = "\n\t"
    v4 = v4
    if (false || true) {
        var v13 uint8 = 134
        _ = v13
    } else {
        _ = len(v1)
    }
    switch -2147483648 {
    case 662:
        v1["a"] = 9223372036854775807
    case 32767:
        for fi14 := 0; fi14 < 1; fi14++ {
            v4 = "\n\t"
            v3 = -1
            v15 := GenS1{F0: 0}
            _ = fi14
            _ = v15
        }
    case -973:
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := []int{-1, fi16}
            var v18 int = fi16
            v3 -= (v3 << ((42) & 63))
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        switch (fi19 & fi19) {
        case 568:
            v20 := []int{(fi19 >> ((18446744073709551615) & 63)), fi19, -1, -1}
            _ = v20
        }
        v4 = ((v4 + ("\n\t" + ("a" + "foo bar"))) + v4)
        _ = fi19
    }
    v3 -= (-(v3))
    fmt.Println(v4)
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 int64 = 1
        var v23 uint = 767
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v3)
    v24 := GenS2{F0: (((((("\n\t" + "hello") + v4) + "foo bar") + ("hello" + "")) + v4) + (v4 + v4))}
    if true {
        var v25 uint = 243
        fmt.Println(v25, v25, v4)
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
