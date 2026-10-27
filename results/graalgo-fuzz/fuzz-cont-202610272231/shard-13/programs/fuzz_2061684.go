package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 int32) uint64 {
    for fi4 := 0; fi4 < 18; fi4++ {
        switch -615 {
        case -1:
            fmt.Println(p0, fi4)
        case -9223372036854775808:
            v5 := GenS1{F0: 1}
            _ = v5
        case 4294967295:
            var v6 int = -412
            _ = v6
        default:
            switch fi4 {
            case -366:
                var v7 int16 = 42
                _ = v7
            default:
                v8 := GenS1{F0: 476}
                _ = v8
            }
        }
        _ = fi4
    }
    var v9 int16 = 616
    var v10 float64 = 0.1
    _ = p0
    _ = v10
    _ = v9
    return 858
}

func main() {
    var v1 float32 = -2.25
    var v2 string = (("foo bar" + "\n\t") + (("a" + "foo bar") + ("foo bar" + "hello")))
    var v3 int8 = 42
    switch -316 {
    case -452:
        var v11 int8 = v3
        _ = v11
    case 9223372036854775807:
        fmt.Println(v3, v1)
    case 1:
        var v12 bool = false
        _ = v12
    }
    v1 = v1
    switch "a" {
    case "hello":
        v13 := []int{32768}
        _ = v13
    case "foo bar":
        v14 := GenS1{F0: 0}
        _ = v14
    case "\n\t":
        var v15 float32 = v1
        _ = v15
    default:
        v2 = ("foo bar" + "foo bar")
    }
    if (v3 >= -1) {
        var v16 bool = true
        switch (("a" + v2) + v2) {
        case "\n\t":
            fmt.Println(v1, v1, v3)
        case "":
            v17 := []int{9223372036854775807, 1}
            _ = v17
        default:
            v18 := GenS2{F0: -9223372036854775808}
            _ = v18
        }
        _ = v16
    }
    v19 := []int{127}
    switch 128 {
    case 128:
        switch v2 {
        case "foo bar":
            for fi20 := 0; fi20 < 19; fi20++ {
                if false {
                    v21 := GenS2{F0: (fi20 + fi20)}
                    fmt.Println(v1, v2, v1)
                    _ = v21
                }
                var v22 uint8 = 86
                v19 = append(v19, fi20)
                _ = fi20
                _ = v22
            }
        }
    case 419:
        var v23 uint64 = 2
        _ = v23
    case 9223372036854775807:
        v3 = (v3 % ((v3) | 1))
    }
    fmt.Println(v1, v1)
    v24 := GenS2{F0: 42}
    var v25 uint = 2
    if false {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    }
    if ((false || true) && (true && (false && (!(4294967295 == uint32(v25)))))) {
        if (!false) {
            fmt.Println(v25)
        }
        var v26 uint32 = 42
        _ = v26
    }
    fmt.Println(v2)
    var v27 uint32 = 65535
    fmt.Println(v25)
    if true {
        if false {
            v28 := GenS3{F0: (v3 >> ((v25) & 63)), F1: v3, F2: 675}
            if false {
                switch v2 {
                case "":
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                case "\n\t":
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                case "foo bar":
                    if (4294967295 != 1) {
                        for fi29 := 0; fi29 < 4; fi29++ {
                            v30 := GenS1{F0: 397}
                            _ = fi29
                            _ = v30
                        }
                    }
                }
                v31 := GenS1{F0: 65535}
                _ = v31
            }
            v3 -= v3
            _ = v28
        }
        _ = len(v19)
        var v32 uint16 = 65535
        _ = v32
    }
    v25 = (v25 - v25)
    v19 = append(v19, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
}
