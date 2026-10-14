package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int16) int64 {
    switch "\n\t" {
    case "\n\t":
        var v5 string = "\n\t"
        _ = v5
    }
    v6 := GenS1{F0: "\n\t"}
    var v7 int = 2147483647
    v8 := GenS1{F0: ("foo bar" + "foo bar")}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 387
}
func genFn1(p0 uint, p1 int) bool {
    if true {
        fmt.Println(p0, p1)
    }
    v9 := GenS1{F0: ("a" + "hello")}
    fmt.Println(p0)
    v10 := GenS1{F0: ("foo bar" + ("hello" + ("hello" + "\n\t")))}
    p0 -= 2
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (!false)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 int32 = 1
    v4 := make(map[string]uint32)
    fmt.Println(v3)
    v11 := GenS2{F0: 9223372036854775807, F1: 326}
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: ""}
        switch ("a" + "\n\t") {
        case "a":
            v3 -= (v3 / (((v3 >> ((2) & 63))) | 1))
        case "\n\t":
            var v14 int32 = (v3 - (v3 << ((2) & 63)))
            _ = v14
        case "hello":
            switch (fi12 % ((2147483647) | 1)) {
            case 2147483647:
                var v15 string = ("a" + ("" + "\n\t"))
                _ = v15
            }
        }
        _ = fi12
        _ = v13
    }
    if false {
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS3{F0: 1, F1: 4294967295}
            _ = fi16
            _ = v17
        }
        v18 := GenS4{F0: 2}
        v3 += v3
        _ = v18
    }
    v3 += (-(v3))
    if true {
        fmt.Println(v3)
    }
    fmt.Println(v3)
    var v19 string = (("a" + (("\n\t" + "") + ("hello" + "foo bar"))) + (("hello" + "foo bar") + "a"))
    var v20 bool = (v3 > (v3 | 679))
    fmt.Println(v19, v19, v20)
    var v21 string = "foo bar"
    var v22 string = v21
    v3 -= v3
    var v23 int32 = 959
    v24 := GenS3{F0: -1, F1: 18446744073709551615}
    v25 := GenS1{F0: "\n\t"}
    if false {
        v21 = v22
        v22 = v19
    } else {
        v26 := GenS3{F0: -752, F1: 65535}
        v27 := []int{32767, 482, 676, 127}
        var v28 uint16 = 396
        _ = v26
        _ = v27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
