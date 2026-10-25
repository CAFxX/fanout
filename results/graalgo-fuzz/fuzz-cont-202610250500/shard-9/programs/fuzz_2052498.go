package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 int32 }
func genFn0() string {
    var v4 int8 = 0
    var v5 uint8 = 42
    switch "\n\t" {
    case "":
        v6 := GenS1{F0: -1, F1: 44}
        _ = v6
    default:
        var v7 uint64 = 2
        _ = v7
    }
    _ = v4
    _ = v5
    return (("hello" + (("" + ("" + "a")) + "\n\t")) + ("a" + (("" + "") + ("hello" + "\n\t"))))
}
func genFn1() int8 {
    v8 := []int{}
    if (602 != 9223372036854775807) {
        if (4294967295 != 313) {
            switch 32767 {
            case -833:
                if len(v8) > 0 {
                    v8[0] = -361
                }
            case -289:
                var v9 uint16 = 0
                _ = v9
            }
            if (true || true) {
                var v10 int = -669
                _ = v10
            }
        } else {
            var v11 int64 = 821
            _ = v11
        }
        var v12 uint32 = 2147483648
        var v13 int16 = 1000
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if ((3.14159 + 3.14159) > -2.25) {
            var v15 int16 = -351
            v15 = (v15 ^ v15)
            _ = v15
        }
        fmt.Println(fi14)
        v16 := GenS2{F0: 549}
        _ = fi14
        _ = v16
    }
    _ = v8
    return 127
}

func main() {
    var v1 int = -2147483648
    v2 := []string{}
    v3 := make(map[string]bool)
    if (-254 >= (v1 % ((v1) | 1))) {
        v17 := GenS3{F0: "a", F1: false, F2: 541}
        _ = v17
    } else {
        var v18 string = ""
        _ = v18
    }
    v1 -= v1
    fmt.Println(v1)
    var v19 int8 = 0
    v20 := GenS1{F0: (^(v1)), F1: 62}
    var v21 uint32 = 65535
    v1 = v1
    v21 -= (v21 - 2)
    v1 -= 1
    var v22 uint = 1
    v19 += v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
