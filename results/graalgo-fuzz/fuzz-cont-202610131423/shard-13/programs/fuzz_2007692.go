package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }

func main() {
    var v1 uint64 = 42
    var v2 uint = 13
    var v3 uint32 = 4294967295
    v2 = (v2 & (^(v2)))
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := []int{}
        switch "foo bar" {
        case "a":
            var v6 int = -570
            _ = v6
        case "":
            var v7 uint = v2
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    if ((("\n\t" + ("a" + ("\n\t" + "a"))) + (((("hello" + ("" + ((("hello" + "\n\t") + "a") + ("a" + ("\n\t" + "hello"))))) + "a") + "") + ("a" + (("hello" + "\n\t") + ("\n\t" + "foo bar"))))) < "a") {
        if (!(true || true)) {
            fmt.Println(v3, v2, v2)
            var v8 uint = (v2 << ((uint(v3)) & 63))
            var v9 uint64 = v1
            _ = v8
            _ = v9
        } else {
            var v10 int32 = 2147483647
            v3 = v3
            if false {
                v3 -= (v3 / ((v3) | 1))
                fmt.Println(v3)
            } else {
                var v11 string = (("hello" + "") + ("foo bar" + "a"))
                fmt.Println(v10, v1)
                fmt.Println(v2)
                _ = v11
            }
            _ = v10
        }
        v2 = v2
    }
    var v12 float32 = (1e10 + 0.1)
    v3 -= (v3 & v3)
    v2 = v2
    if false {
        v13 := []int{9223372036854775807}
        v3 = (v3 & (v3 ^ v3))
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int32 = 636
            var v16 string = ""
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        v17 := []int{0, -1}
        v1 -= v1
        var v18 uint64 = v1
        _ = v17
        _ = v18
    }
    if (((!(-670 != 0)) || (v2 == v2)) && true) {
        v3 = 8
        var v19 uint64 = 1
        var v20 uint64 = v19
        _ = v19
        _ = v20
    } else {
        v1 += (v1 % (((v1 % ((42) | 1))) | 1))
        for fi21 := 0; fi21 < 12; fi21++ {
            v1 += v1
            _ = fi21
        }
    }
    if (!true) {
        v22 := GenS1{F0: -1, F1: "\n\t"}
        var v23 int32 = -1
        _ = v22
        _ = v23
    } else {
        var v24 int16 = int16(v3)
        var v25 int32 = -2147483648
        v2 = (v2 % ((v2) | 1))
        _ = v24
        _ = v25
    }
    var v26 float64 = 0.0
    v26 += (v26 + (v26 * (-2.25 * v26)))
    var v27 int8 = -128
    var v28 uint16 = 42
    var v29 uint16 = (v28 % (((v28 / ((v28) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
