package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 int16) uint16 {
    switch 9223372036854775807 {
    case 13:
        switch ("" + ("" + ("" + "foo bar"))) {
        case "":
            if false {
                for fi4 := 0; fi4 < 10; fi4++ {
                    v5 := GenS1{F0: 194, F1: 371}
                    _ = fi4
                    _ = v5
                }
                v6 := []int{1}
                for fi7 := 0; fi7 < 14; fi7++ {
                    v8 := GenS1{F0: 127, F1: 0}
                    p0 = p0
                    var v9 bool = true
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                _ = v6
            } else {
                var v10 uint = 1
                _ = v10
            }
        case "a":
            v11 := GenS2{F0: 18446744073709551615, F1: 75}
            _ = v11
        default:
            var v12 uint64 = 2
            _ = v12
        }
    default:
        p0 = p0
    }
    p0 = 32767
    fmt.Println(p0)
    _ = p0
    return 255
}

func main() {
    var v1 uint = 4294967295
    var v2 int = -119
    v3 := []string{"hello", ("a" + "a")}
    v1 = 732
    v2 = (v2 / (((^(v2))) | 1))
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println(v2)
    switch "foo bar" {
    case "foo bar":
        for fi13 := 0; fi13 < 19; fi13++ {
            v2 -= 128
            var v14 bool = (3.14159 >= 0.0)
            _ = fi13
            _ = v14
        }
    case "\n\t":
        if (!false) {
            var v15 int16 = 1000
            _ = v15
        } else {
            switch (-(v2)) {
            case 128:
                var v16 uint64 = 1
                _ = v16
            case -2147483648:
                v17 := GenS2{F0: 1, F1: -128}
                _ = v17
            }
            if true {
                _ = len(v3)
                v2 = (v2 >> (((v1 << ((768) & 63))) & 63))
                var v18 uint32 = 971
                _ = v18
            }
        }
    case "a":
        var v19 string = "hello"
        _ = v19
    }
    var v20 int32 = 381
    if (!(!(-896 >= -1))) {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    fmt.Println(v1, v1, v1)
    _ = len(v3)
    v1 += 0
    if len(v3) > 0 {
        v3[0] = ("a" + ((("foo bar" + "foo bar") + "") + ""))
    }
    var v21 int = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
