package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 uint16) uint64 {
    if false {
        fmt.Println(p0)
    }
    if (1 == 255) {
        var v7 uint16 = p0
        _ = v7
    }
    v8 := GenS1{F0: 146}
    switch 385 {
    case 61:
        var v9 string = "hello"
        _ = v9
    }
    var v10 int32 = -664
    _ = p0
    _ = v10
    _ = v8
    return 0
}
func genFn1(p0 int8) int16 {
    switch "hello" {
    case "\n\t":
        p0 = 42
    default:
        v11 := GenS1{F0: 42}
        _ = v11
    }
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 int = 128
    v4 := make(map[int]string)
    var v5 int32 = -1
    v6 := []string{"a", "a"}
    v12 := GenS2{F0: 2, F1: "\n\t"}
    switch (v3 << ((4294967295) & 63)) {
    case 1:
        var v13 uint32 = 2
        _ = v13
    case 127:
        v14 := GenS3{F0: 2}
        _ = v14
    default:
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int16 = -32768
            _ = fi15
            _ = v16
        }
    }
    var v17 float64 = 0.1
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 int = (v3 / (((^(v3))) | 1))
        if (1 == 1) {
            if true {
                if false {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := GenS4{F0: (42 > 763), F1: v5, F2: "foo bar"}
                        _ = len(v6)
                        var v22 uint8 = 175
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            }
            var v23 uint8 = 255
            v24 := GenS1{F0: 1}
            _ = v23
            _ = v24
        }
        _ = fi18
        _ = v19
    }
    v3 += (v3 ^ 0)
    if false {
        var v25 bool = (!((!true) && false))
        var v26 float64 = v17
        _ = len(v4)
        _ = v25
        _ = v26
    }
    var v27 uint = 1
    switch ("a" + "foo bar") {
    case "":
        switch "\n\t" {
        case "\n\t":
            if true {
                if false {
                    v17 = 3.14159
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v28 uint64 = 712
                    _ = v28
                }
            } else {
                var v29 uint32 = 42
                _ = v29
            }
        }
    case "foo bar":
        v5 += v5
    }
    var v30 uint8 = 127
    for fi31 := 0; fi31 < 14; fi31++ {
        v3 -= fi31
        _ = fi31
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := GenS4{F0: (32767 == -3), F1: v5, F2: ""}
        switch fi32 {
        case 32767:
            if true {
                if (!true) {
                    fmt.Println(fi32)
                    switch ("a" + "a") {
                    case "hello":
                        var v34 uint8 = (v30 * 127)
                        _ = v34
                    }
                }
            }
        case 128:
            if true {
                for fi35 := 0; fi35 < 20; fi35++ {
                    _ = len(v6)
                    v5 = v5
                    v36 := GenS2{F0: 0, F1: (((((("" + ((("\n\t" + "\n\t") + ("a" + (("\n\t" + "\n\t") + ("hello" + "a")))) + "foo bar")) + (("foo bar" + "a") + "\n\t")) + "hello") + ((("a" + "hello") + "a") + "a")) + ("foo bar" + ("a" + "foo bar"))) + ("a" + ("foo bar" + "hello")))}
                    _ = fi35
                    _ = v36
                }
            } else {
                var v37 int = fi32
                _ = v37
            }
        }
        if false {
            fmt.Println(v5, v30)
        }
        _ = fi32
        _ = v33
    }
    var v38 uint8 = 208
    var v39 uint64 = 620
    v27 = 255
    fmt.Println(v27, v30, v3)
    v3 = 636
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
