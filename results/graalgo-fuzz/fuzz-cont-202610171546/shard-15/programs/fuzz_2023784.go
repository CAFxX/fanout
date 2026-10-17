package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 string; F1 int64; F2 bool }

func main() {
    var v1 float64 = (3.14159 / -2.25)
    var v2 int64 = 9223372036854775807
    var v3 float32 = ((1.5 + (0.1 * -2.25)) * -2.25)
    v4 := []string{"a", ("foo bar" + ("" + "foo bar")), ("\n\t" + "foo bar"), ((("a" + "a") + ("foo bar" + "\n\t")) + (("foo bar" + ("\n\t" + ((("hello" + "hello") + "") + "foo bar"))) + (("a" + "hello") + "a")))}
    if len(v4) > 0 {
        v4[0] = ""
    }
    var v5 string = ("" + "a")
    v5 = v5
    if (!false) {
        v4 = append(v4, v5)
        if (!("hello" < v5)) {
            v6 := GenS1{F0: 18446744073709551615, F1: 32767}
            v7 := GenS1{F0: 0, F1: 32767}
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            if (fi8 != fi8) {
                var v9 uint16 = 1
                switch fi8 {
                case 127:
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                default:
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v11 := GenS2{F0: 166}
                        _ = fi10
                        _ = v11
                    }
                }
                fmt.Println(v1, v9)
                _ = v9
            }
            if true {
                fmt.Println(v2, fi8)
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                var v12 uint16 = 255
                _ = v12
            } else {
                fmt.Println(v1, fi8)
            }
            _ = fi8
        }
    } else {
        if len(v4) > 0 {
            v4[0] = "a"
        }
        if false {
            var v13 int32 = 2147483647
            var v14 uint32 = 73
            _ = v13
            _ = v14
        } else {
            v3 -= (1.5 + v3)
            if false {
                v2 = (v2 - v2)
            }
        }
        v15 := GenS3{F0: false, F1: 0}
        _ = v15
    }
    var v16 uint8 = 127
    var v17 uint16 = 0
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    var v18 int8 = -44
    v19 := GenS4{F0: v5, F1: 42, F2: true}
    v16 -= (v16 + 1)
    v16 += (v16 - v16)
    v20 := GenS1{F0: 612, F1: 42}
    var v21 uint16 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
