package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0(p0 uint16, p1 uint64, p2 int64) string {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 42}
        var v7 int32 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 0}
    fmt.Println(p0, p2, p2)
    v9 := GenS1{F0: 28}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1() uint8 {
    switch ((("foo bar" + "hello") + "hello") + "a") {
    case "\n\t":
        var v10 int64 = 42
        _ = v10
    }
    v11 := GenS2{F0: -215, F1: 1}
    if (!(42 > 0)) {
        var v12 uint16 = 255
        var v13 uint16 = v12
        v14 := GenS3{F0: 42, F1: 1}
        _ = v12
        _ = v13
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 17; fi15++ {
            switch "foo bar" {
            case "\n\t":
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 uint64 = 18
                    _ = fi16
                    _ = v17
                }
            case "hello":
                v18 := []int{-2147483648, (fi15 / ((127) | 1)), 9223372036854775807}
                _ = v18
            }
            _ = fi15
        }
        var v19 int32 = -2147483648
        v20 := []int{4294967295, 2147483647, 42, 226}
        _ = v19
        _ = v20
    }
    _ = v11
    return 255
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    var v3 uint64 = (^(v1))
    v4 := []string{"foo bar", "hello", (("a" + (("" + "foo bar") + "foo bar")) + ("" + (("foo bar" + "foo bar") + "foo bar"))), "hello"}
    var v21 float32 = 3.14159
    var v22 int64 = -898
    var v23 bool = (!false)
    v1 -= v3
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 int8 = 127
        _ = fi24
        _ = v25
    }
    if (!((!(920 >= 1)) && (false || v23))) {
        v1 = (v1 & (v3 % (((v1 * v3)) | 1)))
    } else {
        v26 := GenS1{F0: -128}
        delete(v2, "")
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := GenS2{F0: 128, F1: 2}
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    var v29 int32 = -1
    v29 = v29
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 int64 = int64(v3)
        var v32 uint = 42
        _ = fi30
        _ = v31
        _ = v32
    }
    v22 = v22
    var v33 int32 = (^(v29))
    var v34 uint16 = 255
    fmt.Println(v1)
    for fi35 := 0; fi35 < 2; fi35++ {
        v3 -= v1
        v3 += v1
        _ = fi35
    }
    fmt.Println(v33, v23)
    fmt.Println(v3, v21)
    var v36 uint64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
}
