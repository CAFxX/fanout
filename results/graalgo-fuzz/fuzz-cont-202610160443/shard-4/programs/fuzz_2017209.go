package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := GenS1{F0: "foo bar", F1: 1, F2: "a"}
    var v5 uint16 = 2
    var v6 uint8 = 1
    if true {
        if true {
            if false {
                switch 0 {
                case 32767:
                    v5 += v5
                case 278:
                    v7 := []int{127}
                    _ = v7
                case 2147483647:
                    v5 = v5
                default:
                    v5 = 22
                }
                fmt.Println(v5)
                if (!false) {
                    v5 = (v5 ^ v5)
                }
            }
            fmt.Println(v6, v6)
            if (!(!(("a" == "\n\t") || false))) {
                v6 = (v6 >> ((1) & 63))
            }
        } else {
            var v8 uint8 = v6
            if (!true) {
                v9 := v2[65535]
                v10 := []int{9223372036854775807, -2147483648, 210}
                _ = v10
                _ = v9
            }
            v11 := v1[""]
            _ = v11
            _ = v8
        }
        v6 = v6
        fmt.Println(v5, v5)
    } else {
        var v12 int32 = 594
        if true {
            v5 = (v5 | v5)
            v12 -= v12
        }
        _ = v12
    }
    v13 := GenS1{F0: "", F1: v5, F2: "hello"}
    if (true && (!true)) {
        fmt.Println(v6, v6)
    } else {
        v5 = (^(v5))
        v14 := GenS1{F0: "foo bar", F1: 42, F2: "hello"}
        for fi15 := 0; fi15 < 1; fi15++ {
            v6 = (^(v6))
            v16 := GenS1{F0: "foo bar", F1: v5, F2: ((("foo bar" + ("" + "\n\t")) + "hello") + "hello")}
            fmt.Println(v6, fi15)
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v17 := GenS1{F0: (("" + (("" + "a") + "a")) + ("\n\t" + "foo bar")), F1: (v5 | (v5 * (v5 | 32767))), F2: ("" + "")}
    var v18 int64 = 450
    delete(v1, "\n\t")
    v18 = 9223372036854775807
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS2{F0: 235}
        v5 = (^(v5))
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint8 = (v6 * (v6 | (v6 >> ((487) & 63))))
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v5)
    fmt.Println(v18, v6)
    var v23 string = ("hello" + "hello")
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(v5)
        for fi25 := 0; fi25 < 18; fi25++ {
            v5 -= 661
            var v26 uint16 = (v5 / ((v5) | 1))
            var v27 int64 = (v18 ^ (v18 ^ (v18 / ((v18) | 1))))
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
