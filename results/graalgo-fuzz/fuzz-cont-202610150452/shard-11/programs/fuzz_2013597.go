package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64 }
func genFn0() int64 {
    v4 := []int{9223372036854775807, 32768, 1}
    var v5 int16 = 0
    v6 := GenS1{F0: ("foo bar" + (("foo bar" + "") + ("\n\t" + (((("foo bar" + "hello") + "foo bar") + "") + ("\n\t" + "foo bar")))))}
    _ = v4
    _ = v5
    _ = v6
    return -649
}

func main() {
    var v1 int16 = 1000
    var v2 int64 = 496
    v3 := []int{969, 42}
    var v7 uint = 2
    fmt.Println(v7)
    _ = len(v3)
    if true {
        switch "" {
        case "":
            var v8 uint = (v7 * (^(v7)))
            _ = v8
        default:
            if true {
                v9 := GenS1{F0: ("a" + (("\n\t" + ("hello" + "foo bar")) + "foo bar"))}
                v10 := GenS2{F0: 18446744073709551615, F1: 118}
                _ = v10
                _ = v9
            }
        }
    } else {
        var v11 uint64 = 2
        _ = v11
    }
    if (255 < 42) {
        var v12 string = "hello"
        _ = v12
    } else {
        _ = len(v3)
    }
    v13 := GenS3{F0: v2}
    var v14 int16 = (^(v1))
    for fi15 := 0; fi15 < 20; fi15++ {
        v14 = (v14 | v14)
        _ = fi15
    }
    var v16 uint32 = 65535
    if (v16 <= (v16 >> ((42) & 63))) {
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(v1, fi17, v2)
            _ = fi17
        }
    } else {
        v16 = (^(v16))
        var v18 int = 32767
        fmt.Println(v1, v7, v16)
        _ = v18
    }
    v1 += (v14 >> (((v7 | (v7 >> ((540) & 63)))) & 63))
    v19 := GenS2{F0: 42, F1: 127}
    var v20 bool = ((!true) && true)
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 bool = v20
        v23 := GenS2{F0: 24, F1: 127}
        _ = fi21
        _ = v22
        _ = v23
    }
    switch -2147483648 {
    case 1:
        if len(v3) > 0 {
            v3[0] = 9223372036854775807
        }
    case 0:
        _ = len(v3)
    case 2147483647:
        fmt.Println(v1, v14, v2)
    }
    v2 = genFn0()
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
