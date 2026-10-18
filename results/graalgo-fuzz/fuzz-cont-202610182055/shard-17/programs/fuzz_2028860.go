package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []string{("\n\t" + ((("" + "foo bar") + ("foo bar" + "hello")) + "hello")), "hello"}
    v2 := []string{"\n\t", ("foo bar" + "foo bar"), "a"}
    var v3 bool = false
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 14; fi5++ {
        if len(v2) > 0 {
            v2[0] = ("" + (("" + "hello") + "\n\t"))
        }
        fmt.Println(v4, v3)
        _ = fi5
    }
    switch ("hello" + "hello") {
    case "foo bar":
        for fi6 := 0; fi6 < 20; fi6++ {
            if (0 <= 122) {
                v3 = (!(0 <= 2))
                v7 := GenS1{F0: 42}
                _ = len(v1)
                _ = v7
            }
            v3 = (!(!true))
            for fi8 := 0; fi8 < 7; fi8++ {
                v3 = ((v4 & 1) == (-(v4)))
                var v9 bool = v3
                var v10 uint16 = 791
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi6
        }
    case "\n\t":
        if ((^(v4)) == v4) {
            v11 := GenS1{F0: 0}
            _ = v11
        } else {
            fmt.Println(v4, v3)
            if true {
                switch -825 {
                case -1:
                    var v12 string = "\n\t"
                    _ = v12
                }
                var v13 int8 = -32
                var v14 float32 = (3.14159 / 1.0)
                _ = v13
                _ = v14
            }
            v3 = (!false)
        }
    case "":
        v15 := GenS1{F0: 1000}
        _ = v15
    }
    var v16 uint16 = 32767
    var v17 float64 = -2.25
    v18 := GenS2{F0: (v4 ^ int32(v16)), F1: 0}
    v19 := GenS3{F0: 197}
    if true {
        var v20 float64 = v17
        _ = v20
    } else {
        var v21 float64 = v17
        _ = v21
    }
    v3 = v3
    v17 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
