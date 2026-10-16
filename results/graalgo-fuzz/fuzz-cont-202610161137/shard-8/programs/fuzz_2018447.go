package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
func genFn0() bool {
    var v4 uint8 = 218
    v4 = 1
    var v5 string = "\n\t"
    fmt.Println(v5)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 float64 = (((-2.25 / 1.0) / 1.0) * ((-2.25 / 3.14159) - 1e10))
        var v8 bool = false
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v4
    _ = v5
    return (false || (true && (("" + (((v5 + v5) + v5) + v5)) != v5)))
}
func genFn1(p0 int16) int {
    var v9 string = "a"
    switch v9 {
    case "a":
        v9 = "hello"
    case "":
        v10 := GenS1{F0: 42, F1: 425}
        _ = v10
    }
    var v11 uint = 4294967295
    v12 := GenS2{F0: 839, F1: "foo bar", F2: 65535}
    _ = p0
    _ = v11
    _ = v12
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2}
    var v3 string = "\n\t"
    var v13 int64 = 1
    v13 = -1
    var v14 uint8 = 0
    fmt.Println(v13, v13, v13)
    switch ((v3 + v3) + v3) {
    case "a":
        var v15 int32 = -964
        _ = v15
    }
    if true {
        switch 2147483647 {
        case 4294967295:
            v13 = 0
        }
        var v16 int16 = 1
        _ = v16
    }
    _ = len(v2)
    var v17 string = v3
    var v18 uint16 = 65535
    if (4294967295 <= 380) {
        var v19 bool = (141 != genFn1(250))
        _ = v19
    }
    var v20 uint8 = (v14 & v14)
    v21 := GenS1{F0: 0, F1: 2}
    for fi22 := 0; fi22 < 14; fi22++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi22
    }
    var v23 uint64 = 1
    var v24 uint = 4294967295
    v24 += (v24 + uint(v18))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
}
