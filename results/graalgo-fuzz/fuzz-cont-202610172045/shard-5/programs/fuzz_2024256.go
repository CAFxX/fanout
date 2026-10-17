package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
func genFn0() string {
    var v4 int8 = -1
    switch 655 {
    case 4294967295:
        fmt.Println(v4)
    default:
        v5 := []int{-2147483648, 321, 1}
        _ = v5
    }
    _ = v4
    return ""
}
func genFn1() uint16 {
    v6 := []int{42, -9223372036854775808, -705, -788}
    if true {
        var v7 uint16 = 479
        switch ("foo bar" + (("foo bar" + "hello") + "a")) {
        case "a":
            if (-892 > 9223372036854775807) {
                var v8 int = -184
                _ = v8
            } else {
                var v9 uint = 2
                var v10 uint32 = 42
                _ = v10
                _ = v9
            }
        case "":
            switch 1 {
            case 32767:
                var v11 int64 = -1
                _ = v11
            case 4294967295:
                var v12 string = "hello"
                _ = v12
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        case "hello":
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint16 = (v7 << ((238) & 63))
                v14 = 255
                _ = fi13
                _ = v14
            }
        }
        _ = len(v6)
        _ = v7
    } else {
        _ = len(v6)
    }
    switch (("\n\t" + "") + "a") {
    case "":
        v15 := GenS1{F0: 924, F1: 2, F2: 32767}
        _ = v15
    }
    _ = v6
    return 274
}

func main() {
    v1 := []int{512, 0, 2147483647, 630}
    var v2 string = "\n\t"
    v3 := []uint32{}
    switch "" {
    case "\n\t":
        v2 = ""
    case "foo bar":
        switch 127 {
        case 233:
            v2 = v2
        case 0:
            if (129 == 4294967295) {
                var v16 float32 = 1e10
                fmt.Println(v2)
                _ = v16
            } else {
                v2 = v2
                v17 := GenS1{F0: 42, F1: 1, F2: 42}
                fmt.Println(v2)
                _ = v17
            }
        }
    }
    fmt.Println(v2)
    v1 = append(v1, 914)
    v2 = v2
    if ((!true) || (false && (true && false))) {
        v18 := GenS2{F0: (false || (true || true)), F1: 1, F2: 18446744073709551615}
        _ = v18
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    if true {
        v2 = v2
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v19 uint64 = 42
    v1 = append(v1, 32767)
    fmt.Println(v19, v19)
    v19 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
