package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 bool = (2147483648 <= 1)
    var v2 uint64 = 740
    var v3 string = "hello"
    v4 := []int{}
    var v5 int64 = -873
    v6 := GenS1{F0: v2, F1: (!(-607 > -914)), F2: 32767}
    var v7 string = ""
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint64 = v2
        _ = fi8
        _ = v9
    }
    var v10 int = 32767
    v1 = (1000 != -731)
    if (!false) {
        v11 := GenS2{F0: (v3 + (((v7 + (("a" + "\n\t") + (("foo bar" + v3) + (v7 + "\n\t")))) + v3) + ("foo bar" + (v7 + ((("hello" + v7) + "") + (v7 + "\n\t")))))), F1: 42, F2: -22}
        _ = v11
    }
    switch (((v3 + v7) + "foo bar") + v3) {
    case "":
        var v12 int32 = -1
        _ = v12
    case "a":
        v13 := GenS3{F0: 1, F1: v2, F2: (v2 + (v2 ^ (v2 % ((v2) | 1))))}
        _ = v13
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    switch (v10 - (v10 / (((v10 & 82)) | 1))) {
    case -2147483648:
        if ((false || (1e10 < 0.1)) || (-1 >= -32768)) {
            v3 = "foo bar"
            var v14 int32 = 2147483647
            _ = v14
        } else {
            v15 := GenS4{F0: 813}
            _ = v15
        }
    case 2147483647:
        v16 := GenS5{F0: false, F1: 42}
        _ = v16
    case 9223372036854775807:
        var v17 string = v7
        _ = v17
    }
    v10 = 32767
    fmt.Println(v3, v5)
    v2 = (v2 | v2)
    var v18 uint8 = 42
    v2 = 4294967295
    if (v1 || ((!v1) && true)) {
        var v19 float64 = 3.14159
        var v20 int16 = 1
        v21 := GenS6{F0: int8(v10)}
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: v2, F1: (203 >= v18), F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
