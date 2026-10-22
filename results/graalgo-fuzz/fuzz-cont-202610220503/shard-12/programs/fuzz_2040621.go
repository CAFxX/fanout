package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0() int64 {
    var v6 string = ("foo bar" + ("a" + (("hello" + "hello") + "a")))
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            if (true && false) {
                v6 = ("hello" + ((v6 + v6) + v6))
                switch v6 {
                case "":
                    var v8 int16 = 42
                    _ = v8
                case "foo bar":
                    var v9 string = v6
                    _ = v9
                }
            }
            if (100 != 100) {
                switch (fi7 / ((127) | 1)) {
                case 128:
                    v10 := GenS1{F0: 9223372036854775807}
                    _ = v10
                }
            }
        }
        _ = fi7
    }
    if false {
        for fi11 := 0; fi11 < 5; fi11++ {
            fmt.Println(fi11, fi11)
            _ = fi11
        }
        if (false || (true || false)) {
            v12 := GenS2{F0: -530, F1: 0, F2: (0 <= 44)}
            v6 = (v6 + v6)
            _ = v12
        }
    }
    switch (v6 + "a") {
    case "\n\t":
        v6 = (("" + v6) + "")
    case "a":
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                v6 = ("foo bar" + (v6 + ("a" + "foo bar")))
            }
            _ = fi13
        }
    default:
        v6 = ("" + ((v6 + v6) + ("" + v6)))
    }
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 65535
    v2 := []int{725, -2147483648}
    var v3 float64 = 3.14159
    var v4 int64 = 1
    v5 := []uint32{(v1 ^ v1), (v1 % ((65535) | 1)), 65535, 257}
    v14 := GenS3{F0: "", F1: int64(v1)}
    var v15 uint = 30
    if false {
        if (42 == 18446744073709551615) {
            v16 := GenS1{F0: 42}
            v15 = 401
            _ = len(v2)
            _ = v16
        } else {
            fmt.Println(v4, v15)
        }
        v5 = append(v5, 42)
        var v17 int = 0
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint = v15
        fmt.Println(v3)
        v4 = 0
        _ = fi18
        _ = v19
    }
    fmt.Println(v15, v4)
    var v20 bool = false
    fmt.Println(v3, v1)
    v21 := GenS4{F0: -1, F1: 1}
    v4 = 9223372036854775807
    if (v20 && ("hello" != "\n\t")) {
        v15 = 255
        if false {
            fmt.Println(v1, v3)
            var v22 float32 = (-2.25 + 1e10)
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
