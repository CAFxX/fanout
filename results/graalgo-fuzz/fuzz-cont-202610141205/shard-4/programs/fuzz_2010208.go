package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint) int16 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 128
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 bool = true
    v2 := []uint32{65535}
    v3 := make(map[int]bool)
    var v4 int16 = -1
    var v5 uint64 = 2
    v6 := make(map[string]string)
    _ = len(v2)
    if false {
        v4 -= v4
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        switch (fi8 | -810) {
        case 4294967295:
            switch 541 {
            case 128:
                var v9 uint64 = v5
                _ = v9
            case 43:
                if (724 <= 65535) {
                    v1 = v1
                    var v10 float64 = 1.5
                    _ = v10
                }
            }
        case -154:
            for fi11 := 0; fi11 < 10; fi11++ {
                v4 -= 1
                if (566 < 506) {
                    if (false || false) {
                        fmt.Println(v4, v4)
                        if (2 >= 1) {
                            fmt.Println(v4)
                            var v12 int64 = 9223372036854775807
                            var v13 uint16 = 65535
                            _ = v12
                            _ = v13
                        }
                        if v1 {
                            var v14 int8 = -128
                            var v15 bool = false
                            v14 -= v14
                            _ = v14
                            _ = v15
                        }
                    }
                    switch ("\n\t" + ("foo bar" + "hello")) {
                    case "hello":
                        fmt.Println(fi11)
                        fmt.Println(v5, v1)
                        v5 += (v5 ^ 42)
                    }
                }
                v1 = (-1 < -1)
                _ = fi11
            }
        case 1:
            var v16 int16 = (v4 - (v4 >> ((2) & 63)))
            _ = v16
        default:
            v1 = false
        }
        v4 += v4
        switch -2147483648 {
        case 128:
            v4 -= -602
        case 2147483647:
            v1 = (!(53 < -1))
        case -1:
            v17 := GenS1{F0: v1, F1: v4, F2: ("hello" + ("" + ""))}
            _ = v17
        }
        _ = fi8
    }
    v4 += genFn0(18446744073709551615)
    v18 := GenS1{F0: v1, F1: genFn0(4294967295), F2: ("" + ("foo bar" + ("foo bar" + ("a" + "a"))))}
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v19 int8 = 100
        _ = v19
    case 1:
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint32 = 0
            v22 := GenS2{F0: 42, F1: 1}
            switch "hello" {
            case "a":
                v6["a"] = ("foo bar" + "")
            case "\n\t":
                switch fi20 {
                case 32767:
                    v23 := GenS1{F0: (!(v5 != (v5 + (v5 << ((840) & 63))))), F1: int16(v5), F2: ("foo bar" + ("foo bar" + (("foo bar" + "") + "foo bar")))}
                    _ = v23
                }
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    case -1:
        fmt.Println(v4)
    default:
        var v24 string = ("a" + "a")
        _ = v24
    }
    v25 := GenS2{F0: 597, F1: 32767}
    var v26 int64 = 0
    if v1 {
        for fi27 := 0; fi27 < 5; fi27++ {
            v26 += v26
            if v1 {
                v28 := GenS1{F0: (!(true || false)), F1: (v4 + v4), F2: "hello"}
                v1 = (("\n\t" + "") < "hello")
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                switch "\n\t" {
                case "a":
                    fmt.Println(v5)
                case "hello":
                    var v30 int = 32768
                    _ = v30
                default:
                    v31 := GenS1{F0: v1, F1: 32767, F2: "hello"}
                    _ = v31
                }
                _ = fi29
            }
            _ = fi27
        }
    }
    var v32 int8 = 85
    fmt.Println(v4, v1, v4)
    v4 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
