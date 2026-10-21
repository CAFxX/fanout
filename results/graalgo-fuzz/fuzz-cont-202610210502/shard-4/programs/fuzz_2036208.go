package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint8 }

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    switch "foo bar" {
    case "hello":
        if true {
            var v4 bool = (!false)
            _ = v4
        }
    case "foo bar":
        v2[2147483647] = 32767
    }
    v2[65535] = -752
    if (true && false) {
        _ = len(v1)
        var v5 int16 = 461
        _ = v5
    } else {
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 float32 = 3.14159
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    }
    if false {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int16 = -32768
            var v11 uint64 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    v3["hello"] = ("" + "a")
    var v12 int64 = 1
    var v13 int16 = 785
    switch -747 {
    case 20:
        v14 := GenS1{F0: 127}
        _ = v14
    case -1:
        var v15 string = "a"
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS2{F0: 1}
        fmt.Println(v12, v13)
        switch (("hello" + "") + "") {
        case "\n\t":
            if false {
                var v18 uint = 4294967295
                _ = v18
            }
        case "hello":
            if (974 == 32767) {
                v1 = append(v1, "a")
                var v19 uint8 = 200
                _ = v19
            } else {
                switch -29 {
                case 0:
                    v20 := GenS1{F0: -1}
                    _ = v20
                case 42:
                    fmt.Println(v12, v13, v13)
                default:
                    v12 = -1
                }
            }
        case "":
            var v21 int8 = -128
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    var v22 bool = true
    var v23 int32 = 2147483647
    v23 = (v23 ^ (v23 ^ (v23 % ((v23) | 1))))
    if false {
        if v22 {
            v24 := GenS3{F0: 44, F1: (v13 << ((4294967295) & 63))}
            _ = v24
        } else {
            v23 += 1
            v25 := GenS2{F0: 252}
            _ = v25
        }
        v26 := GenS4{F0: 2147483648, F1: 255, F2: 255}
        v27 := GenS1{F0: -76}
        _ = v26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
