package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int32 = 205
    var v2 uint8 = 5
    var v3 int = -9223372036854775808
    var v4 uint = 1
    if ((v1 >= (v1 & 2147483647)) || (v3 != -9223372036854775808)) {
        for fi5 := 0; fi5 < 1; fi5++ {
            v6 := GenS1{F0: "hello"}
            _ = fi5
            _ = v6
        }
        v4 = 4294967295
        var v7 int8 = -24
        _ = v7
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            if true {
                v1 = (^(v1))
            } else {
                var v9 int8 = 100
                v10 := GenS1{F0: ("hello" + "hello")}
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    }
    fmt.Println(v1, v3, v1)
    var v11 int64 = 0
    switch ("foo bar" + "foo bar") {
    case "":
        v12 := GenS2{F0: (v1 ^ int32(v11)), F1: 1, F2: "a"}
        _ = v12
    case "foo bar":
        var v13 string = "foo bar"
        _ = v13
    }
    v3 = -115
    for fi14 := 0; fi14 < 1; fi14++ {
        v4 += (v4 - (v4 << ((v4) & 63)))
        if (!(false || (((true || true) || true) || (2 < 42)))) {
            v3 = (v3 | -2147483648)
            switch v3 {
            case 1:
                v15 := GenS3{F0: 65535, F1: v2}
                _ = v15
            case 439:
                v16 := []int{}
                _ = v16
            default:
                v3 -= v3
            }
        } else {
            v11 -= v11
        }
        fmt.Println(v11, v1)
        _ = fi14
    }
    v17 := GenS2{F0: v1, F1: -2147483648, F2: ((("a" + ("a" + "hello")) + ((("foo bar" + ("hello" + ((("foo bar" + ("hello" + "foo bar")) + ((("foo bar" + ("" + ("hello" + ""))) + "hello") + "\n\t")) + ("hello" + "\n\t")))) + "foo bar") + "hello")) + ("\n\t" + "foo bar"))}
    v2 -= v2
    v18 := GenS3{F0: 358, F1: (v2 - 74)}
    for fi19 := 0; fi19 < 8; fi19++ {
        fmt.Println(fi19)
        v2 -= 1
        _ = fi19
    }
    v20 := []int{(v3 - 128), v3, (v3 ^ 841), v3}
    var v21 bool = (("\n\t" + ("" + (("hello" + "\n\t") + ("hello" + "foo bar")))) != "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
