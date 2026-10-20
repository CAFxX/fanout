package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() uint {
    switch "hello" {
    case "hello":
        var v5 uint32 = 4294967295
        _ = v5
    }
    var v6 uint = 42
    if false {
        v6 += v6
        v7 := []int{1, 127, -701, 128}
        _ = v7
    } else {
        fmt.Println(v6)
        v8 := []int{969, -2147483648}
        _ = v8
    }
    v9 := []int{-9223372036854775808, 636, 4294967295, 4294967295}
    v10 := GenS1{F0: 0}
    _ = v10
    _ = v6
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    v4 := []string{"", ("foo bar" + ("\n\t" + "hello")), ("\n\t" + "foo bar")}
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v1)
    }
    switch -489 {
    case 0:
        fmt.Println(v1)
    case 123:
        switch 913 {
        case -48:
            for fi11 := 0; fi11 < 5; fi11++ {
                v1 += (v1 + v1)
                _ = fi11
            }
        case 510:
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: 4294967295}
                if (776 <= 0) {
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "\n\t")
                    }
                    v1 += 2
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v15 := GenS1{F0: 2}
                        if false {
                            var v16 int = 32767
                            fmt.Println(fi12)
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                } else {
                    var v17 int16 = 42
                    _ = v17
                }
                _ = fi12
                _ = v13
            }
        case 0:
            delete(v2, 4294967295)
        default:
            fmt.Println(v1)
        }
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        if (!true) {
            if (0 == 1) {
                switch 4294967295 {
                case 1:
                    delete(v2, 564)
                case -1:
                    v19 := GenS1{F0: 1}
                    _ = v19
                case -926:
                    var v20 int = fi18
                    _ = v20
                }
            }
        }
        fmt.Println(v1, fi18)
        _ = fi18
    }
    var v21 int8 = 127
    fmt.Println(v1)
    v1 += (v1 >> ((42) & 63))
    for fi22 := 0; fi22 < 1; fi22++ {
        v21 -= (v21 / (((^(v21))) | 1))
        v23 := GenS1{F0: 65535}
        _ = fi22
        _ = v23
    }
    v24 := v2[4294967295]
    v2[2147483648] = 4294967295
    v24 = v24
    var v25 uint16 = 864
    delete(v2, 925)
    v26 := GenS1{F0: 4294967295}
    if ((v25 > 32767) || false) {
        v1 -= v1
        var v27 bool = false
        if true {
            var v28 float64 = -2.25
            v27 = true
            _ = v28
        }
        _ = v27
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        fmt.Println(v25)
        var v30 uint16 = 1
        var v31 int8 = v21
        _ = fi29
        _ = v30
        _ = v31
    }
    if false {
        v32 := GenS1{F0: 42}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
