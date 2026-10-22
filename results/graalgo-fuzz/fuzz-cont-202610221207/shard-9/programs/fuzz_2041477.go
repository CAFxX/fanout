package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8; F1 string; F2 uint }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint64 }

func main() {
    v1 := []int{32767, 1, -9223372036854775808, 128}
    v2 := []string{}
    var v3 bool = true
    var v4 uint = 42
    v5 := []string{"a", (("foo bar" + "hello") + "a"), ("foo bar" + "hello"), "a"}
    v6 := make(map[string]string)
    v4 += v4
    var v7 int8 = -51
    delete(v6, "a")
    if (532 <= -2147483648) {
        switch ("\n\t" + "\n\t") {
        case "a":
            v7 = v7
        case "foo bar":
            var v8 uint8 = 128
            _ = v8
        }
    }
    if (v3 || v3) {
        v9 := GenS1{F0: 2147483648, F1: -1}
        v3 = true
        _ = v9
    }
    if (!true) {
        v10 := GenS1{F0: 0, F1: 127}
        switch "" {
        case "":
            switch 32768 {
            case 4294967295:
                var v11 uint8 = 127
                _ = v11
            }
        case "a":
            v12 := GenS2{F0: v4, F1: ""}
            _ = v12
        }
        v3 = false
        _ = v10
    }
    switch -2147483648 {
    case 2147483647:
        v7 += v7
    default:
        if false {
            switch 922 {
            case -1:
                fmt.Println(v7, v3, v4)
            case 32767:
                fmt.Println(v4, v7, v3)
            }
            v13 := GenS3{F0: 51, F1: "hello", F2: (v4 ^ v4)}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS4{F0: 1, F1: 65535, F2: 42}
                _ = fi14
                _ = v15
            }
        }
    }
    v1 = append(v1, 775)
    fmt.Println(v3, v3, v4)
    if (!v3) {
        v7 = (v7 >> ((216) & 63))
        v5 = append(v5, ((("foo bar" + "") + "a") + ""))
    }
    v16 := GenS4{F0: 2, F1: 2, F2: 369}
    v17 := GenS4{F0: 181, F1: 209, F2: 4294967295}
    switch 0 {
    case -9223372036854775808:
        v18 := GenS2{F0: v4, F1: "\n\t"}
        _ = v18
    default:
        v4 = (v4 % ((v4) | 1))
    }
    v7 -= -16
    var v19 uint = 343
    v7 = (-(v7))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
