package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint64 = 42
    var v2 int64 = 0
    var v3 float32 = 0.0
    switch -382 {
    case 42:
        v4 := GenS1{F0: -2147483648}
        _ = v4
    }
    v3 -= v3
    switch 478 {
    case 0:
        switch ("\n\t" + "") {
        case "\n\t":
            v5 := GenS2{F0: 120}
            _ = v5
        case "a":
            var v6 int32 = -2147483648
            _ = v6
        }
    case 207:
        fmt.Println(v1, v3, v2)
    default:
        v7 := []int{}
        _ = v7
    }
    if (0 >= -128) {
        var v8 int16 = 42
        v9 := []int{2147483647, 9223372036854775807, 1, 300}
        switch "a" {
        case "":
            switch (("hello" + ("a" + "a")) + "\n\t") {
            case "":
                if len(v9) > 0 {
                    v9[0] = 2147483647
                }
            case "foo bar":
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            case "\n\t":
                fmt.Println(v1, v3)
            }
        case "a":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        default:
            switch "a" {
            case "hello":
                if (v8 == -1) {
                    var v10 uint = 1
                    fmt.Println(v10, v1, v3)
                    _ = v10
                }
            case "a":
                var v11 float64 = -2.25
                _ = v11
            case "":
                fmt.Println(v3, v8)
            }
        }
        _ = v8
        _ = v9
    } else {
        var v12 int64 = (v2 % (((^(v2))) | 1))
        _ = v12
    }
    v13 := []int{0, 830, 2147483647}
    _ = len(v13)
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: int(v2)}
        _ = fi14
        _ = v15
    }
    var v16 int8 = -32
    v17 := GenS1{F0: 32767}
    v18 := GenS2{F0: 1000}
    if (false && (-2.25 < 1.5)) {
        switch (("foo bar" + "foo bar") + "hello") {
        case "":
            var v19 uint8 = 180
            _ = v19
        case "\n\t":
            var v20 uint16 = 65535
            _ = v20
        default:
            if (!("" < "hello")) {
                _ = len(v13)
                for fi21 := 0; fi21 < 1; fi21++ {
                    fmt.Println(fi21, v16)
                    _ = len(v13)
                    _ = fi21
                }
                var v22 uint64 = v1
                _ = v22
            }
        }
    } else {
        for fi23 := 0; fi23 < 10; fi23++ {
            v1 += (v1 / ((18446744073709551615) | 1))
            _ = fi23
        }
    }
    v13 = append(v13, 42)
    switch "a" {
    case "a":
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint = 1
            v3 += v3
            _ = fi24
            _ = v25
        }
    case "":
        var v26 uint64 = 42
        _ = v26
    case "hello":
        var v27 float64 = -2.25
        _ = v27
    }
    v28 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
}
