package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint64; F1 int }
func genFn0() int {
    if false {
        v4 := []int{}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v4
    }
    var v5 int8 = 72
    _ = v5
    return -1
}

func main() {
    v1 := []string{}
    v2 := []string{"foo bar", "", (("" + "foo bar") + (("\n\t" + "hello") + ("\n\t" + "foo bar"))), ((("foo bar" + "hello") + "foo bar") + "hello")}
    var v3 uint32 = 2
    v6 := GenS1{F0: (false && true)}
    v7 := GenS2{F0: ("\n\t" + "\n\t"), F1: -525}
    var v8 uint64 = 0
    v1 = append(v1, ((((("\n\t" + "a") + ("a" + "")) + "hello") + "\n\t") + "hello"))
    var v9 int64 = 1
    v10 := GenS3{F0: -75, F1: 42}
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS4{F0: (v8 ^ v8), F1: (fi11 / (((fi11 | (fi11 / ((128) | 1)))) | 1))}
        switch "a" {
        case "a":
            v9 += v9
        case "foo bar":
            for fi13 := 0; fi13 < 16; fi13++ {
                fmt.Println(fi13, v3, fi11)
                if (784 >= 531) {
                    if (-649 < 198) {
                        var v14 uint32 = (v3 - (v3 % (((v3 << ((304) & 63))) | 1)))
                        var v15 uint64 = (v8 ^ (v8 << ((49) & 63)))
                        _ = v14
                        _ = v15
                    }
                    v3 = v3
                } else {
                    var v16 int16 = -167
                    v9 -= v9
                    _ = v16
                }
                v3 = 42
                _ = fi13
            }
        case "\n\t":
            var v17 uint8 = 42
            _ = v17
        default:
            fmt.Println(v8, v9, v3)
        }
        v1 = append(v1, ("hello" + ("hello" + "\n\t")))
        _ = fi11
        _ = v12
    }
    var v18 int64 = (-(v9))
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v19 int16 = 0
    for fi20 := 0; fi20 < 9; fi20++ {
        switch (fi20 ^ (-(fi20))) {
        case -1:
            fmt.Println(v19, v18)
        default:
            var v21 uint64 = (^(v8))
            _ = v21
        }
        _ = fi20
    }
    fmt.Println(v19)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    switch "\n\t" {
    case "hello":
        v3 = v3
    case "":
        switch (("a" + "\n\t") + "a") {
        case "foo bar":
            for fi22 := 0; fi22 < 15; fi22++ {
                v18 -= 9223372036854775807
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                v19 = (v19 / ((v19) | 1))
                _ = fi22
            }
        case "":
            if ((-953 > 2147483647) || (!(-2.25 > (3.14159 * 1e10)))) {
                fmt.Println(v8, v3)
            }
        }
    case "a":
        v18 += 42
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v8 -= v8
        _ = fi23
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v24 := GenS1{F0: (false || false)}
    var v25 int64 = (v18 << ((255) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
