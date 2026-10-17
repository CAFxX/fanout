package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 string }
func genFn0() int32 {
    v7 := GenS1{F0: 32767, F1: 1}
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{(fi8 & 1), fi8, fi8, -9223372036854775808}
        v10 := GenS1{F0: 65535, F1: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int = 1
    v12 := GenS2{F0: "\n\t"}
    _ = v11
    _ = v12
    _ = v7
    return 1
}
func genFn1(p0 int) uint {
    if (42 > 0) {
        var v13 int = p0
        for fi14 := 0; fi14 < 5; fi14++ {
            v13 = p0
            var v15 int = 0
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v16 int32 = 42
        var v17 int16 = 32767
        _ = v16
        _ = v17
    }
    p0 -= (p0 | -616)
    p0 -= (p0 | (p0 ^ p0))
    v18 := []int{-9223372036854775808, 4294967295, p0, (^(p0))}
    _ = p0
    _ = v18
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint = 615
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    v6 := make(map[uint32]int)
    switch ((("a" + "hello") + "\n\t") + (("foo bar" + "hello") + "foo bar")) {
    case "foo bar":
        v3 = (v3 | 255)
    case "a":
        var v19 uint8 = 2
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        if false {
            var v21 uint32 = 412
            _ = len(v5)
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 string = "hello"
                var v24 uint16 = 721
                var v25 uint64 = 18446744073709551615
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v21
        }
        _ = fi20
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 int8 = 42
        if ((v27 ^ v27) <= v27) {
            v3 -= (v3 % ((v3) | 1))
        }
        _ = fi26
        _ = v27
    }
    switch "foo bar" {
    case "":
        v28 := []int{}
        _ = v28
    case "foo bar":
        v29 := GenS2{F0: "foo bar"}
        _ = v29
    case "a":
        fmt.Println(v3)
    default:
        v30 := []int{119, -1, 32768, 127}
        _ = v30
    }
    if (1 < 255) {
        v3 += genFn1(127)
        fmt.Println(v3)
        fmt.Println(v3)
    }
    var v31 int64 = -9223372036854775808
    if true {
        switch 32767 {
        case -464:
            for fi32 := 0; fi32 < 3; fi32++ {
                for fi33 := 0; fi33 < 18; fi33++ {
                    v31 = v31
                    _ = fi33
                }
                v2[4294967295] = 1
                _ = fi32
            }
        case 32768:
            v34 := GenS2{F0: ("foo bar" + "a")}
            _ = v34
        }
        if true {
            v3 = 2
        }
    } else {
        _ = len(v6)
        var v35 int = 42
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := v1[0]
            switch ("hello" + ("\n\t" + (("\n\t" + ("hello" + "\n\t")) + (((((("a" + ("hello" + "foo bar")) + (("" + "hello") + "\n\t")) + "hello") + (("hello" + "\n\t") + "")) + "foo bar") + ("a" + "foo bar"))))) {
            case "\n\t":
                if ((1e10 - (0.0 - 1e10)) > 0.0) {
                    v37 = false
                    v35 -= (fi36 ^ (fi36 ^ fi36))
                }
            }
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    var v38 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
