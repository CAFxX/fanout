package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint; F2 uint }
type GenS4 struct { F0 string; F1 bool; F2 int }

func main() {
    v1 := []string{("foo bar" + ""), "", (("a" + ("hello" + "\n\t")) + (("" + ("\n\t" + "")) + "a")), "foo bar"}
    var v2 int8 = -128
    v3 := []string{"a", "foo bar", "\n\t"}
    var v4 uint8 = 127
    fmt.Println(v2, v2)
    v4 = (v4 + 127)
    var v5 float64 = ((-2.25 + 1.5) * 1e10)
    _ = len(v3)
    v6 := GenS1{F0: 42}
    var v7 string = ""
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 string = ("hello" + "\n\t")
        v9 = v9
        v1 = append(v1, "foo bar")
        _ = fi8
        _ = v9
    }
    if false {
        var v10 int32 = 2147483647
        for fi11 := 0; fi11 < 6; fi11++ {
            v7 = "foo bar"
            v12 := GenS2{F0: ("" + (v7 + "hello"))}
            var v13 int8 = (v2 & (v2 % ((-82) | 1)))
            _ = fi11
            _ = v12
            _ = v13
        }
        v10 = v10
        _ = v10
    } else {
        v14 := GenS1{F0: 32767}
        var v15 string = v7
        _ = v14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS3{F0: v4, F1: 355, F2: 1}
            var v18 int = (fi16 / ((9223372036854775807) | 1))
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        v4 = v4
        if false {
            v2 = (v2 & (v2 << ((7) & 63)))
        } else {
            for fi19 := 0; fi19 < 17; fi19++ {
                v4 = v4
                _ = fi19
            }
            v4 += v4
            v4 = 255
        }
        _ = len(v1)
    }
    v5 = v5
    v20 := GenS4{F0: (v7 + "foo bar"), F1: true, F2: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
