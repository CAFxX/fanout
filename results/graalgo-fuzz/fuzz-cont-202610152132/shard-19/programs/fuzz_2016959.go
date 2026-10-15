package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    v3 := []uint32{0, 565, 0, 4294967295}
    v4 := make(map[int]string)
    var v5 uint8 = 2
    var v6 string = "foo bar"
    var v7 uint = 42
    v5 = v5
    v8 := GenS1{F0: false, F1: 65535}
    v5 = v5
    var v9 string = "foo bar"
    switch -2147483648 {
    case 9223372036854775807:
        var v10 string = v9
        _ = v10
    case -381:
        v7 = (v7 & (v7 | (v7 + (v7 / ((v7) | 1)))))
    }
    if ((("a" + (((v6 + "") + ("\n\t" + (v6 + (v9 + v9)))) + v6)) + "hello") > v9) {
        switch 2147483647 {
        case 0:
            for fi11 := 0; fi11 < 13; fi11++ {
                if false {
                    v5 = v5
                }
                _ = fi11
            }
        case 297:
            v5 += 1
        }
        if false {
            v6 = v9
            var v12 int32 = 315
            var v13 uint32 = 4294967295
            _ = v12
            _ = v13
        } else {
            v9 = "a"
            var v14 uint32 = 4294967295
            if (-1 != 42) {
                if true {
                    v14 = 4294967295
                }
            } else {
                v7 -= v7
                if true {
                    v14 = v14
                    v5 = (v5 - (v5 & 1))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                var v15 int32 = -271
                _ = v15
            }
            _ = v14
        }
    }
    v4[-925] = ""
    v16 := GenS1{F0: false, F1: v7}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
