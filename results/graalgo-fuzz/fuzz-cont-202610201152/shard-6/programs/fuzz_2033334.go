package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint16 = 22
    var v2 string = "foo bar"
    var v3 int32 = 0
    var v4 int32 = (v3 + (-(v3)))
    v5 := []uint32{4294967295, 213, 42}
    v6 := []string{}
    var v7 uint16 = (v1 << ((1) & 63))
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v10 int32 = -505
            if ((v3 * (v10 + v3)) == -2147483648) {
                var v11 uint8 = 1
                fmt.Println(v4)
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        var v12 uint8 = 2
        _ = fi8
        _ = v12
    }
    v13 := GenS1{F0: 32767, F1: 1}
    switch ((v2 + "") + (("" + v2) + (("\n\t" + ("\n\t" + (v2 + "\n\t"))) + v2))) {
    case "":
        var v14 uint64 = 141
        _ = v14
    case "hello":
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint8 = 127
            var v17 string = "a"
            _ = fi15
            _ = v16
            _ = v17
        }
    default:
        v18 := GenS1{F0: -1, F1: 42}
        _ = v18
    }
    v19 := GenS1{F0: -1, F1: -843}
    v2 = (((v2 + ((("foo bar" + ("\n\t" + v2)) + ((v2 + (v2 + "")) + "\n\t")) + ("\n\t" + ""))) + v2) + (v2 + v2))
    v20 := GenS1{F0: int16(v4), F1: -716}
    switch v2 {
    case "hello":
        _ = len(v5)
    }
    v7 = v7
    switch ((("hello" + "foo bar") + "\n\t") + "hello") {
    case "foo bar":
        if true {
            var v21 uint16 = 274
            _ = v21
        } else {
            v1 -= (v1 << ((924) & 63))
            v2 = "hello"
            v6 = append(v6, (v2 + (((v2 + v2) + v2) + ((v2 + (("\n\t" + v2) + v2)) + (("hello" + "") + (v2 + ((((v2 + "a") + v2) + ("\n\t" + ("hello" + "hello"))) + "a")))))))
        }
    case "a":
        var v22 int64 = 0
        _ = v22
    case "hello":
        _ = len(v6)
    }
    v6 = append(v6, (((v2 + v2) + v2) + (v2 + ("hello" + ("" + v2)))))
    v3 += (^(v3))
    switch v2 {
    case "hello":
        v2 = v2
    case "foo bar":
        _ = len(v5)
    }
    var v23 int16 = 1000
    for fi24 := 0; fi24 < 2; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            if ((!false) && false) {
                if (2147483647 > 0) {
                    var v26 bool = (3.14159 >= (float32((-2.25 - 0.1)) - ((0.1 * 1e10) + 0.0)))
                    fmt.Println(fi24, fi25, v23)
                    _ = v26
                } else {
                    v27 := GenS1{F0: (-(v23)), F1: 1000}
                    var v28 uint64 = 4294967295
                    _ = len(v5)
                    _ = v27
                    _ = v28
                }
                fmt.Println(v2)
                fmt.Println(v3)
            }
            _ = fi25
        }
        v29 := GenS2{F0: 1, F1: "a"}
        if false {
            fmt.Println(v23, v23, v7)
        }
        _ = fi24
        _ = v29
    }
    v30 := GenS3{F0: 255}
    var v31 uint16 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
