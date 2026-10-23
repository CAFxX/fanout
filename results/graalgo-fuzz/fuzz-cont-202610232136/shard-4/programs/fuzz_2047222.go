package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint = 1
    v2 := []string{}
    v3 := make(map[string]string)
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    _ = len(v2)
    var v4 int64 = 9223372036854775807
    var v5 int64 = (v4 ^ 42)
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: ("\n\t" + ""), F1: -1, F2: 4294967295}
        v1 -= 255
        _ = fi7
        _ = v8
    }
    if (!(0.1 == 1.5)) {
        var v9 uint32 = 1
        _ = v9
    }
    v10 := GenS2{F0: 1, F1: 2147483647}
    var v11 string = ""
    v4 = (v5 % (((v4 * (v5 - -1))) | 1))
    v11 = "foo bar"
    if ((217 == (v6 >> ((uint(v6)) & 63))) || false) {
        switch -1 {
        case 998:
            v12 := GenS3{F0: -9223372036854775808}
            _ = v12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v3[(v11 + "hello")] = (v11 + v11)
            _ = fi13
        }
        var v14 uint16 = 0
        _ = v14
    }
    fmt.Println(v11, v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
