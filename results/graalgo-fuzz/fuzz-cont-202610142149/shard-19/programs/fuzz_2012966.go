package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() int {
    v5 := GenS1{F0: 853, F1: 4294967295}
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(fi6)
        switch (fi6 / (((fi6 & fi6)) | 1)) {
        case 528:
            var v7 uint64 = 0
            _ = v7
        }
        _ = fi6
    }
    switch "" {
    case "":
        var v8 int64 = 42
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: 255, F1: 50}
        _ = v9
    case "hello":
        var v10 int32 = 793
        _ = v10
    }
    var v11 uint16 = 559
    v12 := GenS2{F0: 751, F1: 933}
    _ = v11
    _ = v12
    _ = v5
    return 509
}
func genFn1() int {
    v13 := GenS2{F0: 504, F1: 255}
    var v14 bool = (736 > 614)
    var v15 float32 = (3.14159 + 0.1)
    if (42 < 0) {
        var v16 uint16 = 1
        _ = v16
    }
    _ = v13
    _ = v14
    _ = v15
    return -962
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 uint16 = 1
    var v4 int64 = 749
    fmt.Println(v4)
    if (false && (true && true)) {
        v3 -= (v3 ^ v3)
    } else {
        v3 = (v3 + 1)
        if true {
            v4 += 0
            var v17 int32 = 0
            _ = v17
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                switch (fi18 * fi18) {
                case 32767:
                    var v19 int = fi18
                    _ = v19
                case 127:
                    if (("a" > ("" + (("hello" + "hello") + "a"))) && true) {
                        v4 += 1
                    }
                }
                _ = fi18
            }
        }
        fmt.Println(v4, v4)
    }
    fmt.Println(v3, v4)
    v3 = v3
    for fi20 := 0; fi20 < 13; fi20++ {
        v3 -= 255
        v3 -= (v3 ^ (v3 + v3))
        v3 = (v3 ^ (^(v3)))
        _ = fi20
    }
    v3 -= (^(v3))
    var v21 string = "foo bar"
    if (!((v21 + (("hello" + ("hello" + "\n\t")) + v21)) < v21)) {
        v4 = 1
        var v22 uint32 = 65535
        for fi23 := 0; fi23 < 9; fi23++ {
            for fi24 := 0; fi24 < 3; fi24++ {
                v21 = "foo bar"
                _ = fi24
            }
            v3 = (v3 / ((100) | 1))
            _ = fi23
        }
        _ = v22
    } else {
        var v25 uint32 = 2147483648
        var v26 bool = ((0.1 - float32(1.5)) == 3.14159)
        var v27 uint8 = 42
        _ = v25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := GenS1{F0: v3, F1: 4294967295}
        _ = fi28
        _ = v29
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
