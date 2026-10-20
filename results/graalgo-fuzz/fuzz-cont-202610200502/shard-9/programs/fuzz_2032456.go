package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 string; F2 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 int8, p1 string) uint32 {
    fmt.Println(p1, p0)
    v7 := GenS1{F0: 1}
    var v8 bool = false
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 408
}
func genFn1(p0 uint8, p1 string) int64 {
    var v9 string = ("foo bar" + ("foo bar" + "foo bar"))
    v10 := []int{128}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]string)
    var v3 uint32 = 42
    var v4 float64 = 3.14159
    var v5 int = -2147483648
    var v6 bool = false
    fmt.Println(v1, v4, v4)
    var v11 int64 = 9223372036854775807
    fmt.Println(v1, v11)
    v12 := GenS1{F0: 18446744073709551615}
    v13 := []int{v5}
    v1 = -1
    switch "hello" {
    case "foo bar":
        v3 = v3
    case "\n\t":
        var v14 int16 = (^(v1))
        _ = v14
    }
    if (127 <= 42) {
        if v6 {
            fmt.Println(v11, v4)
            for fi15 := 0; fi15 < 7; fi15++ {
                fmt.Println(v6, v11, v6)
                _ = fi15
            }
        }
        v16 := GenS1{F0: 65535}
        for fi17 := 0; fi17 < 9; fi17++ {
            if false {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v11 -= v11
                    v19 := GenS2{F0: 752, F1: "a", F2: 42}
                    _ = fi18
                    _ = v19
                }
            }
            var v20 bool = v6
            _ = fi17
            _ = v20
        }
        _ = v16
    } else {
        for fi21 := 0; fi21 < 17; fi21++ {
            v3 = v3
            v5 = fi21
            _ = fi21
        }
    }
    v5 -= (^(v5))
    v1 -= v1
    switch (v5 ^ (-(v5))) {
    case 128:
        var v22 int = v5
        _ = v22
    default:
        switch "foo bar" {
        case "hello":
            if v6 {
                fmt.Println(v4)
            } else {
                var v23 bool = (!v6)
                _ = v23
            }
        case "":
            v24 := GenS3{F0: 4294967295}
            _ = v24
        }
    }
    var v25 uint = 4294967295
    v3 -= v3
    switch 4294967295 {
    case 42:
        fmt.Println(v25)
    case -382:
        switch ("" + (("\n\t" + "foo bar") + (("\n\t" + ("\n\t" + "")) + ""))) {
        case "\n\t":
            var v26 uint8 = 128
            _ = v26
        case "":
            if ((v25 & (^(v25))) < v25) {
                v27 := GenS2{F0: (^(v25)), F1: "hello", F2: 42}
                _ = v27
            }
        case "a":
            v25 += 42
        default:
            var v28 int = int(v1)
            _ = v28
        }
    case -712:
        var v29 uint8 = 255
        _ = v29
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        v31 := GenS2{F0: (v25 - (v25 << ((722) & 63))), F1: "a", F2: -2147483648}
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
