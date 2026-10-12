package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int32 = -11
    v2 := []string{"a", (("hello" + "a") + ((("foo bar" + "\n\t") + (((("foo bar" + "a") + "a") + (("hello" + "a") + "hello")) + ("foo bar" + (("foo bar" + ((("hello" + "\n\t") + (("foo bar" + "") + "hello")) + "")) + "foo bar")))) + "foo bar")), "a"}
    var v3 int = -301
    v4 := []uint32{65535, 839, 0, 464}
    switch -1 {
    case -9223372036854775808:
        v3 = (v3 & 32768)
    case 1:
        v1 -= v1
    default:
        _ = len(v2)
    }
    fmt.Println(v3, v3)
    var v5 float32 = -2.25
    v3 = v3
    v6 := GenS1{F0: "\n\t", F1: 88}
    switch ("foo bar" + ((("" + "a") + ("\n\t" + "a")) + ("" + "a"))) {
    case "\n\t":
        v5 += 0.0
    case "hello":
        _ = len(v4)
    default:
        v5 -= v5
    }
    v1 -= v1
    v7 := GenS2{F0: 2}
    var v8 int = (-(v3))
    v9 := GenS3{F0: 1, F1: 32767}
    v1 = 1
    v8 = v3
    switch "a" {
    case "hello":
        fmt.Println(v1, v8, v1)
    case "":
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(fi10, v1)
            _ = fi10
        }
    case "foo bar":
        _ = len(v4)
    default:
        var v11 string = "a"
        _ = v11
    }
    switch (-(v3)) {
    case 1:
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 uint = 0
                var v15 int16 = 1
                fmt.Println(v15)
                _ = fi13
                _ = v14
                _ = v15
            }
            if len(v4) > 0 {
                v4[0] = 2
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi12
        }
    case 396:
        if (0 < 42) {
            if (("hello" + "hello") > (("foo bar" + "") + "a")) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    v1 += 2147483647
                    v16 := GenS4{F0: 127}
                    _ = v16
                }
                fmt.Println(v8)
            } else {
                v17 := GenS5{F0: 42, F1: 42}
                _ = v17
            }
        } else {
            var v18 uint32 = 65535
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
