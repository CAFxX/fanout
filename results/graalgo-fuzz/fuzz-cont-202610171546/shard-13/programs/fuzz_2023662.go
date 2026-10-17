package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 int64, p1 bool, p2 int32) uint8 {
    var v5 int = 32768
    v6 := []int{(v5 + v5), 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1() int32 {
    v7 := []int{128}
    switch ("hello" + ("hello" + "foo bar")) {
    case "hello":
        var v8 int16 = -1
        _ = v8
    case "a":
        var v9 int32 = 1
        _ = v9
    case "foo bar":
        if len(v7) > 0 {
            v7[0] = -642
        }
    default:
        if (224 == -2147483648) {
            if true {
                var v10 bool = (!true)
                v10 = v10
                v10 = v10
                _ = v10
            } else {
                v11 := GenS1{F0: -128, F1: 9, F2: (-874 == -1)}
                var v12 bool = (true && true)
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 float32 = 0.1
                    v14 = (1.5 - -2.25)
                    var v15 int32 = 684
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v11
                _ = v12
            }
        }
    }
    if false {
        if (-9223372036854775808 != 42) {
            var v16 bool = false
            _ = v16
        } else {
            v7 = append(v7, 148)
            var v17 int32 = 1
            for fi18 := 0; fi18 < 16; fi18++ {
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 int = 127
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            _ = v17
        }
        var v21 string = "foo bar"
        _ = len(v7)
        _ = v21
    }
    _ = v7
    return -2147483648
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{("hello" + "foo bar"), "\n\t", "\n\t", ((("\n\t" + "\n\t") + "") + ("foo bar" + ""))}
    var v3 float32 = 3.14159
    v4 := make(map[string]bool)
    v22 := GenS1{F0: 42, F1: -128, F2: (18446744073709551615 <= 474)}
    _ = len(v4)
    switch ("foo bar" + "foo bar") {
    case "hello":
        if false {
            v3 += 3.14159
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint8 = 110
                v3 = -2.25
                fmt.Println(v3, v24, v24)
                _ = fi23
                _ = v24
            }
            v25 := GenS2{F0: 9223372036854775807, F1: 42}
            _ = v25
        } else {
            var v26 uint8 = 42
            var v27 int32 = -1
            _ = v26
            _ = v27
        }
    case "":
        v1 = (v1 - v1)
    }
    var v28 float64 = 1.5
    switch ((((((("foo bar" + "a") + "hello") + "\n\t") + "hello") + "foo bar") + ((("foo bar" + "foo bar") + ("" + "\n\t")) + "foo bar")) + "\n\t") {
    case "hello":
        fmt.Println(v3, v3, v28)
    }
    v28 -= float64(1.5)
    if (!(false || false)) {
        fmt.Println(v3, v3, v28)
    }
    v1 = (v1 * 1)
    var v29 float64 = float64(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
