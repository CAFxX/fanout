package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string; F1 string; F2 bool }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 uint64 }
func genFn0(p0 string) int64 {
    var v4 int8 = -1
    var v5 bool = false
    var v6 uint32 = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = false
    var v3 int16 = 952
    if v2 {
        v2 = false
        v7 := []int{128, -1, -405, 32767}
        fmt.Println(v3, v2, v1)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{}
            v10 := GenS1{F0: 1, F1: 0}
            v2 = true
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 uint = 96
    fmt.Println(v11, v1)
    v12 := []int{}
    if (42 <= 2) {
        fmt.Println(v1, v3)
        v13 := GenS2{F0: "a", F1: "hello", F2: (42 < 162)}
        v12 = append(v12, 32767)
        _ = v13
    } else {
        if ((-70 != int8(v11)) || (v2 && (1.5 != -2.25))) {
            for fi14 := 0; fi14 < 19; fi14++ {
                v3 -= v3
                v3 = v3
                _ = fi14
            }
        }
    }
    if (-1 > -16) {
        var v15 int8 = 127
        var v16 uint32 = 2147483647
        v2 = ((((v2 || v2) || true) && (-1 <= 544)) && v2)
        _ = v15
        _ = v16
    } else {
        v2 = v2
        v17 := GenS2{F0: "", F1: ((("" + "") + "\n\t") + "a"), F2: v2}
        if true {
            v18 := GenS2{F0: "a", F1: "hello", F2: (!v2)}
            var v19 uint64 = 108
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS1{F0: 584, F1: (v19 & (v19 ^ 1))}
                _ = fi20
                _ = v21
            }
            _ = v18
            _ = v19
        }
        _ = v17
    }
    _ = len(v12)
    v22 := GenS2{F0: ("hello" + "hello"), F1: "foo bar", F2: (((false && (v2 || v2)) && v2) || v2)}
    v23 := GenS3{F0: (("\n\t" + "hello") + ("a" + "\n\t")), F1: (v11 >> ((4294967295) & 63)), F2: v11}
    v24 := GenS4{F0: 4294967295}
    v12 = append(v12, -9223372036854775808)
    v11 += 643
    for fi25 := 0; fi25 < 1; fi25++ {
        if v2 {
            switch 819 {
            case 913:
                v1 += v1
            }
        } else {
            fmt.Println(v3)
            v26 := GenS4{F0: 1}
            switch -9223372036854775808 {
            case -403:
                v27 := GenS1{F0: genFn0("foo bar"), F1: 2}
                _ = v27
            case 2147483647:
                for fi28 := 0; fi28 < 8; fi28++ {
                    var v29 float64 = -2.25
                    _ = fi28
                    _ = v29
                }
            default:
                switch -595 {
                case 32767:
                    var v30 uint16 = 2
                    _ = v30
                default:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
            }
            _ = v26
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
