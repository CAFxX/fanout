package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint, p1 int16) bool {
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(fi4, p1)
        if (false || false) {
            if (true && (0 == 481)) {
                fmt.Println(p1, p1)
                p0 += p0
                v5 := GenS1{F0: (false || false), F1: true, F2: 42}
                _ = v5
            }
        }
        _ = fi4
    }
    switch "\n\t" {
    case "foo bar":
        v6 := GenS1{F0: (!true), F1: (!true), F2: 0}
        _ = v6
    default:
        var v7 int32 = -379
        _ = v7
    }
    switch ("hello" + "\n\t") {
    case "hello":
        v8 := []int{}
        _ = v8
    default:
        p0 = 743
    }
    _ = p0
    _ = p1
    return (-794 <= 676)
}

func main() {
    var v1 float32 = (1.5 / ((1.5 / 1e10) * -2.25))
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    v9 := GenS1{F0: (0.0 != (3.14159 + 1e10)), F1: (true || false), F2: 0}
    var v10 int = -9223372036854775808
    switch ("foo bar" + ("hello" + v2)) {
    case "hello":
        var v11 bool = true
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := []int{(v10 - fi12)}
            _ = fi12
            _ = v13
        }
    case "foo bar":
        v3[v10] = (!false)
    }
    switch v2 {
    case "\n\t":
        if false {
            var v14 int = (v10 & -2147483648)
            var v15 uint = 138
            _ = v14
            _ = v15
        } else {
            v1 = float32(1e10)
        }
    case "foo bar":
        v2 = (v2 + v2)
    case "":
        v10 = (v10 << ((65535) & 63))
    default:
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS2{F0: -32768, F1: 1}
            v18 := []int{}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v1 = float32(0.1)
    var v19 int16 = -284
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(fi20, v1, v1)
        switch v2 {
        case "a":
            v1 = v1
        default:
            v21 := GenS2{F0: (v19 >> ((18446744073709551615) & 63)), F1: 0}
            _ = v21
        }
        _ = fi20
    }
    switch (-(v10)) {
    case 455:
        v19 -= (v19 / ((v19) | 1))
    case 1:
        switch v10 {
        case 0:
            var v22 float64 = (1.5 + 1.5)
            _ = v22
        case 742:
            var v23 uint64 = 639
            _ = v23
        default:
            v24 := v3[(v10 ^ -2147483648)]
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        switch v2 {
        case "":
            var v26 uint8 = 248
            _ = v26
        case "hello":
            switch 187 {
            case 886:
                v27 := []int{}
                _ = v27
            case 655:
                v1 += v1
            case -830:
                v28 := GenS1{F0: (0.0 >= ((3.14159 / 3.14159) + v1)), F1: (-906 != 0), F2: 32767}
                _ = v28
            }
        case "foo bar":
            for fi29 := 0; fi29 < 15; fi29++ {
                fmt.Println(v2)
                var v30 bool = true
                var v31 int32 = 679
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        _ = fi25
    }
    fmt.Println(v10, v1)
    v2 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
