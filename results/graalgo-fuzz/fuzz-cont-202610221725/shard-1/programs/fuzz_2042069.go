package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint32; F1 bool; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 string }
func genFn0() string {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int32 = 44
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: false, F1: 128}
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint8 = 1
        _ = fi8
        _ = v9
    }
    _ = v7
    return ("hello" + "a")
}
func genFn1(p0 string) uint16 {
    if false {
        if true {
            if (919 <= -1) {
                var v10 uint8 = 169
                v10 = v10
                var v11 int = 1
                _ = v10
                _ = v11
            } else {
                switch ("" + "\n\t") {
                case "":
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := GenS2{F0: 2, F1: ((true || (((3.14159 - 1.5) < 1e10) || (!(!(32767 < 42))))) && true), F2: 0}
                        fmt.Println(fi12, fi12)
                        p0 = "a"
                        _ = fi12
                        _ = v13
                    }
                case "\n\t":
                    var v14 uint = 42
                    _ = v14
                }
                var v15 uint64 = 144
                _ = v15
            }
            fmt.Println(p0)
        } else {
            v16 := GenS3{F0: 536}
            for fi17 := 0; fi17 < 15; fi17++ {
                switch "" {
                case "a":
                    var v18 string = p0
                    _ = v18
                }
                v19 := []int{}
                _ = fi17
                _ = v19
            }
            _ = v16
        }
    }
    var v20 uint16 = 325
    if (-1 > 396) {
        var v21 uint8 = uint8(v20)
        v22 := []int{}
        _ = v21
        _ = v22
    } else {
        fmt.Println(p0, v20)
        v23 := GenS1{F0: ((1e10 / -2.25) == (1e10 + float32(1e10))), F1: -9223372036854775808}
        _ = v23
    }
    _ = p0
    _ = v20
    return (v20 % ((1) | 1))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 int32 = -2147483648
    var v4 uint16 = 58
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int64 = 9223372036854775807
        _ = fi24
        _ = v25
    }
    var v26 uint8 = 128
    var v27 int = 4294967295
    for fi28 := 0; fi28 < 9; fi28++ {
        fmt.Println(v26, fi28)
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS1{F0: false, F1: (fi29 % (((^(fi29))) | 1))}
            fmt.Println(v4)
            var v31 uint16 = 65535
            _ = fi29
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            fmt.Println(v26, fi32, fi28)
            fmt.Println(v4, v4)
            var v33 string = (((((("foo bar" + "foo bar") + "hello") + (("foo bar" + (("hello" + ("hello" + "")) + "hello")) + "a")) + "a") + "hello") + ("foo bar" + ("foo bar" + "a")))
            v34 := GenS4{F0: v33, F1: v33}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi28
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        if (("a" + ("" + "a")) == ((("foo bar" + ("\n\t" + "\n\t")) + "") + "hello")) {
            v26 += v26
            v3 = -1
            v36 := []int{}
            _ = v36
        } else {
            v27 = fi35
        }
        _ = len(v1)
        _ = fi35
    }
    v3 = -2147483648
    var v37 uint8 = v26
    v3 = v3
    v38 := GenS1{F0: false, F1: v27}
    v4 -= (v4 * (v4 / ((v4) | 1)))
    v1[-602] = 42
    for fi39 := 0; fi39 < 20; fi39++ {
        v40 := v2[65535]
        _ = fi39
        _ = v40
    }
    if ((!false) || (true && (754 > 823))) {
        for fi41 := 0; fi41 < 10; fi41++ {
            v27 = 128
            _ = fi41
        }
        var v42 int64 = 1
        _ = v42
    } else {
        var v43 int64 = 42
        v26 += 2
        v4 = 119
        _ = v43
    }
    fmt.Println(v3, v4, v37)
    for fi44 := 0; fi44 < 3; fi44++ {
        for fi45 := 0; fi45 < 10; fi45++ {
            fmt.Println(v26, v3)
            _ = fi45
        }
        _ = fi44
    }
    fmt.Println(v4, v37)
    if (!false) {
        if (795 >= 65535) {
            v27 = v27
            v26 = 25
            var v46 uint32 = uint32(v37)
            _ = v46
        }
        v4 = (v4 / (((^(v4))) | 1))
    }
    var v47 int = v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v47
}
