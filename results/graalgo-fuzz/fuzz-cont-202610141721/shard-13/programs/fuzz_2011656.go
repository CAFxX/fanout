package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0() int64 {
    v7 := []int{32768}
    var v8 int32 = 0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 int64 = 857
    v5 := make(map[int]string)
    var v6 uint8 = 0
    switch "foo bar" {
    case "a":
        var v9 uint8 = v6
        _ = v9
    }
    if ("foo bar" < "a") {
        switch "a" {
        case "a":
            switch (("a" + "hello") + ("\n\t" + (("foo bar" + "foo bar") + "\n\t"))) {
            case "a":
                v4 = 1
            case "hello":
                v4 = (v4 / (((v4 >> ((507) & 63))) | 1))
            case "":
                v6 += (v6 * 161)
            }
        case "\n\t":
            v10 := []int{9223372036854775807, 9223372036854775807}
            _ = v10
        case "":
            var v11 float64 = 3.14159
            _ = v11
        }
        v4 = genFn0()
    }
    v12 := GenS1{F0: 0, F1: -217, F2: false}
    v6 = 1
    var v13 uint32 = 0
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS2{F0: v6}
        fmt.Println(v13, fi14, v13)
        _ = fi14
        _ = v15
    }
    v16 := []int{2147483647, 0}
    if true {
        v17 := GenS1{F0: int16(v13), F1: -1, F2: (2 > 1)}
        var v18 string = ("hello" + "hello")
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 9; fi19++ {
            v6 += (v6 & 128)
            if (false && (false || false)) {
                v6 = 255
            } else {
                _ = len(v16)
            }
            _ = fi19
        }
        var v20 string = ""
        _ = v20
    }
    var v21 bool = (float32(-2.25) <= float32(1e10))
    switch "foo bar" {
    case "a":
        if true {
            _ = len(v16)
            if false {
                fmt.Println(v6, v13, v6)
                v22 := GenS1{F0: 42, F1: 1, F2: (!true)}
                v23 := GenS1{F0: 0, F1: 717, F2: v21}
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                v6 = v6
                _ = fi24
            }
        } else {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    case "":
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 int8 = int8(v4)
            switch fi25 {
            case 128:
                v27 := GenS3{F0: 4294967295, F1: 42}
                _ = v27
            }
            if v21 {
                fmt.Println(v13, v4)
                var v28 uint8 = (^(v6))
                _ = v28
            }
            _ = fi25
            _ = v26
        }
    default:
        v4 = (v4 + genFn0())
    }
    switch "" {
    case "":
        v6 = (v6 ^ 42)
    case "foo bar":
        switch ((("" + (("\n\t" + (("" + "\n\t") + "\n\t")) + ("hello" + "hello"))) + "\n\t") + "") {
        case "foo bar":
            v29 := v2[32767]
            _ = v29
        }
    default:
        v13 += (v13 + v13)
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v30 int64 = (v4 % ((v4) | 1))
    var v31 uint64 = 25
    fmt.Println(v6, v21, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
