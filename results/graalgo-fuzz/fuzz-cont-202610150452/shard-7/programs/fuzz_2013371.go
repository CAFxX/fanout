package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
func genFn0() int64 {
    v5 := []int{1}
    v6 := GenS1{F0: ("\n\t" + ("" + "\n\t")), F1: 890}
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"\n\t", ""}
    v3 := make(map[string]uint32)
    var v4 int = 9223372036854775807
    fmt.Println(v4)
    v4 -= -710
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: ("hello" + "hello"), F1: 186}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        _ = len(v1)
        if true {
            switch v4 {
            case 2147483647:
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint = 0
                    _ = fi10
                    _ = v11
                }
            case 42:
                fmt.Println(fi9, v4)
            }
            switch 440 {
            case 2147483647:
                var v12 uint32 = 159
                _ = v12
            case 0:
                fmt.Println(fi9)
            case 32768:
                v13 := GenS2{F0: 42, F1: 0}
                _ = v13
            }
            if false {
                v14 := GenS2{F0: 32767, F1: 42}
                fmt.Println(fi9)
                _ = v14
            } else {
                if true {
                    fmt.Println(fi9, v4)
                } else {
                    v4 -= (fi9 | v4)
                    var v15 uint32 = 4294967295
                    _ = v15
                }
            }
        } else {
            var v16 int8 = 1
            v4 += 128
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS1{F0: "foo bar", F1: -32768}
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        _ = fi9
    }
    v19 := GenS2{F0: -222, F1: 504}
    if (110 != 127) {
        switch "hello" {
        case "foo bar":
            for fi20 := 0; fi20 < 16; fi20++ {
                v4 = (fi20 * (v4 + -80))
                _ = fi20
            }
        }
        fmt.Println(v4)
        fmt.Println(v4)
    } else {
        var v21 uint = 390
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := v1[-899]
        fmt.Println(v23)
        v4 = fi22
        _ = fi22
        _ = v23
    }
    if (!(255 >= 0)) {
        v24 := GenS2{F0: -1, F1: 4294967295}
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(v4, v4)
            _ = fi25
        }
        switch -158 {
        case 127:
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 int32 = 0
                v28 := GenS2{F0: 1000, F1: 0}
                v27 = v27
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            switch (-(v4)) {
            case 42:
                for fi29 := 0; fi29 < 1; fi29++ {
                    v2 = append(v2, "a")
                    fmt.Println(fi29)
                    _ = fi29
                }
            }
        }
        _ = v24
    }
    v30 := GenS3{F0: ("\n\t" + "")}
    var v31 int8 = 109
    v32 := GenS4{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
