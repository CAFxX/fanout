package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS4 struct { F0 bool; F1 int; F2 bool }

func main() {
    v1 := []string{""}
    var v2 int = 32768
    var v3 uint64 = 42
    var v4 uint64 = (^(v3))
    var v5 int = (v2 + v2)
    switch (((("" + "foo bar") + "foo bar") + "hello") + "hello") {
    case "":
        v3 -= v3
    case "hello":
        v2 = (v5 & 9223372036854775807)
    }
    v6 := GenS1{F0: 225, F1: v4}
    v3 -= 18446744073709551615
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v7 := GenS2{F0: 0, F1: (v3 + 42), F2: 2147483647}
        _ = v7
    } else {
        if false {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int = (v5 + -9223372036854775808)
                _ = fi8
                _ = v9
            }
            var v10 uint = 0
            _ = v10
        } else {
            v4 -= 18446744073709551615
            v4 = (v3 % ((v3) | 1))
            var v11 int32 = -283
            _ = v11
        }
    }
    switch (("\n\t" + "\n\t") + "") {
    case "a":
        var v12 uint64 = (v4 % ((v3) | 1))
        _ = v12
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        if (true && (!false)) {
            var v13 uint16 = 915
            if (228 > 2147483647) {
                v14 := GenS2{F0: 38, F1: 0, F2: 2147483647}
                _ = v14
            }
            _ = v13
        } else {
            var v15 int32 = 0
            if (42 >= -126) {
                fmt.Println(v3, v3)
                v3 = v4
            }
            _ = v15
        }
    }
    var v16 float64 = ((0.0 / 1e10) * 3.14159)
    var v17 int16 = 0
    if (!false) {
        v4 -= v3
    } else {
        fmt.Println(v2, v2, v17)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v18 uint32 = 2147483647
        _ = v18
    }
    v19 := GenS3{F0: 2, F1: v17, F2: 65535}
    var v20 int8 = 119
    _ = len(v1)
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            v4 += (v3 % ((18446744073709551615) | 1))
            v23 := GenS4{F0: (42 <= 42), F1: 42, F2: true}
            var v24 uint32 = 2147483648
            _ = fi22
            _ = v23
            _ = v24
        }
        if len(v1) > 0 {
            v1[0] = ("hello" + ("hello" + "hello"))
        }
        _ = fi21
    }
    v16 = (v16 / v16)
    var v25 bool = false
    fmt.Println(v2, v16)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
