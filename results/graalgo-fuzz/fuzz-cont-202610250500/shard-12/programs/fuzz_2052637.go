package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint, p1 string) string {
    if true {
        var v6 int64 = -1
        if false {
            p0 = p0
        }
        _ = v6
    }
    v7 := GenS1{F0: -60, F1: 65535, F2: 169}
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}
func genFn1(p0 int, p1 string) uint16 {
    if true {
        var v8 uint8 = 25
        _ = v8
    }
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 bool = false
        var v12 float32 = (0.0 * 0.1)
        var v13 int32 = -2147483648
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: -327, F1: (683 >= 2), F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    var v3 uint8 = 128
    var v4 uint8 = v3
    var v5 float64 = 3.14159
    var v15 uint8 = v4
    switch 2147483647 {
    case 1:
        _ = len(v1)
    case 128:
        if true {
            if (-1 >= 0) {
                v16 := GenS2{F0: -527, F1: (!(true && (2 <= 42))), F2: 1}
                v17 := GenS3{F0: ((true && true) || ((0.0 - 1e10) <= 0.0)), F1: v3}
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 int64 = int64(v2)
                    _ = fi18
                    _ = v19
                }
            }
            var v20 int64 = 550
            _ = v20
        } else {
            if (!(42 == 0)) {
                switch (("foo bar" + (("foo bar" + ("hello" + "foo bar")) + ("" + "\n\t"))) + "foo bar") {
                case "":
                    v1[2147483647] = 42
                }
                switch ((("hello" + (("foo bar" + "hello") + ("" + "hello"))) + "foo bar") + "") {
                case "\n\t":
                    fmt.Println(v15)
                case "foo bar":
                    var v21 uint8 = 128
                    _ = v21
                }
                var v22 int = -9223372036854775808
                _ = v22
            }
            if false {
                v2 += v2
                v23 := GenS3{F0: (!true), F1: (^(v15))}
                _ = v23
            }
            var v24 uint8 = (v4 - v3)
            _ = v24
        }
    case -743:
        v5 += 3.14159
    }
    switch "a" {
    case "a":
        var v25 int16 = 1000
        _ = v25
    default:
        var v26 uint16 = genFn1(128, ("a" + ("hello" + "")))
        _ = v26
    }
    var v27 int16 = 42
    v28 := []int{2147483647, -795}
    if (975 > 2147483647) {
        switch (("a" + ((("\n\t" + "\n\t") + "a") + "a")) + "hello") {
        case "hello":
            v28 = append(v28, -783)
        case "":
            var v29 bool = true
            _ = v29
        }
    }
    v1[0] = 4294967295
    v15 += v4
    v30 := GenS1{F0: 1, F1: 296, F2: v15}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
