package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 string }

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 2
    var v3 int = 128
    v4 := GenS1{F0: v3, F1: -695, F2: 42}
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS2{F0: (v2 / ((v2) | 1))}
        v7 := []int{v3}
        _ = fi5
        _ = v6
        _ = v7
    }
    if ("foo bar" > "hello") {
        if false {
            v1 -= v1
            fmt.Println(v3)
            v1 += v1
        }
        fmt.Println(v1)
    }
    v8 := []int{(v3 / ((v3) | 1)), 2147483647, -1, 32768}
    v9 := GenS2{F0: 403}
    for fi10 := 0; fi10 < 2; fi10++ {
        if true {
            v11 := GenS3{F0: 4294967295, F1: true}
            if false {
                switch (fi10 - (-(fi10))) {
                case 42:
                    v1 += v1
                case 127:
                    fmt.Println(v1, v1)
                default:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        v13 := GenS3{F0: 0, F1: true}
                        _ = fi12
                        _ = v13
                    }
                }
                _ = len(v8)
                var v14 uint32 = v2
                _ = v14
            } else {
                v2 = 2147483648
                switch (("hello" + ("\n\t" + "hello")) + (("foo bar" + "\n\t") + ((("a" + ("a" + "")) + "") + "\n\t"))) {
                case "hello":
                    v1 -= (v1 + 89)
                }
                _ = len(v8)
            }
            _ = v11
        }
        v1 = (v1 % (((v1 * v1)) | 1))
        var v15 uint = 2
        _ = fi10
        _ = v15
    }
    switch (^(v3)) {
    case -700:
        switch 4294967295 {
        case -2147483648:
            v1 = v1
        case 0:
            v16 := GenS1{F0: v3, F1: -1, F2: 4294967295}
            _ = v16
        }
    default:
        v17 := GenS4{F0: ("" + ("\n\t" + ("" + ("\n\t" + "a"))))}
        _ = v17
    }
    fmt.Println(v2)
    var v18 uint8 = 255
    v8 = append(v8, -809)
    if true {
        v8 = append(v8, v3)
    } else {
        _ = len(v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
