package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    var v1 int64 = 0
    var v2 uint = 2
    var v3 uint32 = 42
    var v4 int = 4294967295
    var v5 bool = (((465 < 65535) && true) || true)
    switch v4 {
    case 4294967295:
        v4 += (v4 ^ 679)
    case 2147483647:
        if (v5 && (100 <= 42)) {
            v4 = (v4 & 0)
            var v6 float64 = 1.5
            v7 := []int{v4, (-(v4))}
            _ = v6
            _ = v7
        }
    default:
        var v8 int64 = v1
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v2)
        _ = fi9
    }
    switch ("\n\t" + ("a" + "hello")) {
    case "a":
        v10 := []int{32767, -367}
        _ = v10
    case "\n\t":
        if v5 {
            for fi11 := 0; fi11 < 13; fi11++ {
                for fi12 := 0; fi12 < 17; fi12++ {
                    fmt.Println(v4, v2)
                    fmt.Println(fi11)
                    fmt.Println(fi11, fi12, fi12)
                    _ = fi12
                }
                _ = fi11
            }
            if v5 {
                v13 := GenS1{F0: "hello", F1: 111}
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 uint8 = 255
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        } else {
            var v16 float64 = 0.1
            _ = v16
        }
    case "foo bar":
        v3 = (v3 ^ (v3 * v3))
    }
    v5 = (v5 || v5)
    v4 = v4
    fmt.Println(v2, v5, v3)
    v1 -= v1
    var v17 float64 = ((float64((0.1 * (0.0 * (1.5 + 0.1)))) * 0.1) + 0.1)
    if (-51 == 42) {
        fmt.Println(v5, v4)
        v18 := GenS1{F0: "", F1: 2}
        fmt.Println(v5, v1)
        _ = v18
    } else {
        if (false && true) {
            fmt.Println(v3, v2, v2)
            fmt.Println(v1, v5, v5)
            fmt.Println(v3, v4, v17)
        } else {
            v17 += 3.14159
        }
        v19 := []int{}
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v2 = v2
        if v5 {
            for fi21 := 0; fi21 < 12; fi21++ {
                v2 -= (v2 - v2)
                var v22 uint32 = 139
                var v23 int32 = -1
                _ = fi21
                _ = v22
                _ = v23
            }
            if (127 > 0) {
                v24 := GenS2{F0: "a"}
                for fi25 := 0; fi25 < 7; fi25++ {
                    v26 := GenS1{F0: "hello", F1: 2}
                    var v27 uint64 = 0
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = v24
            }
        } else {
            v2 = (v2 / ((v2) | 1))
        }
        v3 += 858
        _ = fi20
    }
    v1 -= -715
    var v28 int = v4
    if v5 {
        for fi29 := 0; fi29 < 4; fi29++ {
            v4 = (v28 + (-(fi29)))
            _ = fi29
        }
        v30 := GenS1{F0: "a", F1: 65535}
        _ = v30
    } else {
        if true {
            v28 = -2147483648
        } else {
            switch "hello" {
            case "":
                fmt.Println(v4, v2, v1)
            case "hello":
                v31 := GenS1{F0: (("foo bar" + ("foo bar" + "hello")) + "\n\t"), F1: 0}
                _ = v31
            case "a":
                v32 := []int{}
                _ = v32
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
