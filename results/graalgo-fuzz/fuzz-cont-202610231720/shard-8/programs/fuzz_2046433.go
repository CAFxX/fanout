package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -713
    var v3 uint8 = 0
    var v4 uint16 = 784
    v5 := []int{42, -156, 4294967295}
    fmt.Println(v4)
    v4 = 2
    if (!false) {
        v2 = v2
    } else {
        v6 := GenS1{F0: 4294967295, F1: 0}
        _ = v6
    }
    var v7 int16 = 32767
    if false {
        if len(v5) > 0 {
            v5[0] = 42
        }
        var v8 int16 = (v7 + (v7 & (v7 + (-(v7)))))
        if false {
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS2{F0: -34}
                _ = fi9
                _ = v10
            }
            v11 := GenS2{F0: 114}
            _ = v11
        } else {
            _ = len(v5)
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := v1[2147483647]
                if (v4 < v4) {
                    v7 -= (v7 & (v7 >> ((18446744073709551615) & 63)))
                }
                _ = fi12
                _ = v13
            }
            v8 += v8
        }
        _ = v8
    }
    fmt.Println(v3, v3, v2)
    var v14 int8 = 127
    _ = len(v5)
    switch (("a" + "hello") + "hello") {
    case "a":
        v2 += (-(v2))
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            v15 := GenS2{F0: v14}
            _ = v15
        case "hello":
            var v16 int32 = v2
            _ = v16
        }
    }
    v1[726] = ""
    v4 -= 42
    if ("foo bar" <= (("a" + "") + "hello")) {
        v4 = 32767
        v4 = 65535
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS1{F0: 0, F1: 42}
            delete(v1, 2)
            var v19 int16 = (v7 % (((v7 / ((v7) | 1))) | 1))
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    v1[883] = "foo bar"
    var v20 uint8 = (v3 >> ((0) & 63))
    v21 := GenS2{F0: v14}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
