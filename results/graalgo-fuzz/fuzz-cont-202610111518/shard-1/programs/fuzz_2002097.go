package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 uint8, p1 bool, p2 uint64) uint16 {
    v4 := []int{-9223372036854775808, -1, 0}
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: fi5, F1: p2}
        _ = fi5
        _ = v6
    }
    v4 = append(v4, -9223372036854775808)
    for fi7 := 0; fi7 < 9; fi7++ {
        switch fi7 {
        case 171:
            switch (fi7 + fi7) {
            case 32767:
                v8 := GenS2{F0: fi7, F1: fi7, F2: (p0 << ((506) & 63))}
                _ = v8
            default:
                if p1 {
                    if len(v4) > 0 {
                        v4[0] = fi7
                    }
                }
            }
        case 117:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 834
}

func main() {
    v1 := []string{"hello", ("hello" + (("\n\t" + ("\n\t" + "a")) + "a")), "foo bar"}
    v2 := make(map[uint32]bool)
    v3 := []string{("" + "foo bar"), "\n\t", (("a" + "hello") + "foo bar"), "\n\t"}
    for fi9 := 0; fi9 < 11; fi9++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi9
    }
    switch 9223372036854775807 {
    case -650:
        if true {
            var v10 int32 = -150
            _ = v10
        }
    }
    v11 := GenS3{F0: 2147483648, F1: 0}
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS1{F0: fi12, F1: 627}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint8 = 128
        _ = fi14
        _ = v15
    }
    switch "" {
    case "foo bar":
        switch "a" {
        case "hello":
            var v16 uint64 = 208
            _ = v16
        default:
            var v17 int8 = 1
            _ = v17
        }
    default:
        var v18 uint8 = 1
        _ = v18
    }
    var v19 float64 = 3.14159
    fmt.Println(v19)
    var v20 float64 = -2.25
    v21 := GenS3{F0: 246, F1: 1}
    if false {
        var v22 uint8 = 0
        _ = v22
    } else {
        v19 = (v20 / -2.25)
    }
    fmt.Println(v19, v20)
    if false {
        fmt.Println(v20, v19)
    } else {
        var v23 uint32 = 65535
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(v19, fi24)
            _ = fi24
        }
        _ = v23
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v25 uint8 = 1
    switch 560 {
    case 9223372036854775807:
        switch 96 {
        case 32768:
            fmt.Println(v25, v20, v19)
        case 440:
            switch 2147483647 {
            case 418:
                v26 := GenS4{F0: 0, F1: "a"}
                _ = v26
            case 32768:
                v27 := GenS3{F0: 2147483647, F1: 2}
                _ = v27
            }
        case 589:
            var v28 int8 = 127
            _ = v28
        }
    case -9223372036854775808:
        v20 = float64(0.0)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
}
