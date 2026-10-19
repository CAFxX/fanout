package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    var v1 int16 = 158
    v2 := make(map[uint32]uint32)
    var v3 int64 = 103
    switch ("hello" + ((((((("foo bar" + "hello") + ("a" + "\n\t")) + "hello") + "") + ("" + "hello")) + ("" + "foo bar")) + (("a" + (((((("foo bar" + "") + (("a" + ("a" + ("" + (("" + "\n\t") + "hello")))) + (("foo bar" + "hello") + "foo bar"))) + ("\n\t" + "a")) + "foo bar") + ("foo bar" + "")) + (("foo bar" + ("" + "")) + "hello"))) + ("hello" + ("" + "foo bar"))))) {
    case "hello":
        v3 = (v3 + v3)
    case "a":
        var v4 bool = (128 <= -2147483648)
        _ = v4
    case "\n\t":
        fmt.Println(v1)
    }
    if true {
        if (false || ((-390 <= 165) || (-862 >= 44))) {
            v5 := []int{1}
            var v6 float32 = 0.1
            _ = v5
            _ = v6
        }
        var v7 uint32 = 1
        v1 = (v1 >> ((42) & 63))
        _ = v7
    }
    v3 = -9223372036854775808
    v8 := []int{}
    var v9 int32 = 2147483647
    v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
    if len(v8) > 0 {
        v8[0] = 9223372036854775807
    }
    if (0 >= 32768) {
        for fi10 := 0; fi10 < 12; fi10++ {
            v2[42] = 4294967295
            v9 += -533
            var v11 uint = 2
            _ = fi10
            _ = v11
        }
        v3 = (v3 & v3)
    } else {
        for fi12 := 0; fi12 < 15; fi12++ {
            delete(v2, 0)
            v9 = (v9 % (((v9 % (((v9 ^ v9)) | 1))) | 1))
            var v13 uint64 = 4294967295
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: (v9 - -24), F1: 1, F2: (("\n\t" + (("hello" + "hello") + ("\n\t" + ("hello" + ("" + ("" + "a")))))) + "\n\t")}
        v9 = (v9 << ((305) & 63))
        _ = v14
    }
    var v15 uint = 42
    v15 += 108
    switch 128 {
    case -858:
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int32 = (v9 | 42)
            var v18 int = fi16
            if (42 <= 26) {
                if ((v3 + -9223372036854775808) != v3) {
                    v2[93] = 287
                    var v19 int8 = -61
                    v20 := GenS1{F0: 2147483647, F1: 349, F2: "a"}
                    _ = v19
                    _ = v20
                }
                fmt.Println(v15)
                fmt.Println(v15, fi16, v9)
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    case 32767:
        switch "hello" {
        case "foo bar":
            if false {
                var v21 int8 = -1
                v22 := GenS2{F0: 0, F1: "\n\t"}
                v23 := GenS1{F0: v9, F1: 2, F2: "a"}
                _ = v21
                _ = v22
                _ = v23
            }
        default:
            if (v15 <= 0) {
                fmt.Println(v1, v3, v1)
                var v24 uint64 = 1
                _ = v24
            } else {
                var v25 uint64 = 76
                var v26 uint32 = 261
                v9 = v9
                _ = v25
                _ = v26
            }
        }
    }
    v15 = (^(v15))
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 string = (((((("\n\t" + "foo bar") + ("" + ("" + "hello"))) + "foo bar") + ("" + "a")) + ("hello" + ((("" + "") + ("\n\t" + "hello")) + ""))) + "hello")
        v15 = 1
        _ = fi27
        _ = v28
    }
    v8 = append(v8, 0)
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 uint16 = 32767
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
