package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 string; F2 int32 }
func genFn0(p0 string) bool {
    fmt.Println(p0)
    p0 = p0
    v7 := GenS1{F0: 797, F1: 841}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return ("foo bar" <= p0)
}
func genFn1(p0 int16) int {
    v8 := []int{127}
    if (!(!(!true))) {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                p0 = p0
                if ((int16(fi9) == (^(p0))) || true) {
                    var v11 string = "\n\t"
                    v11 = ("a" + "hello")
                    var v12 float64 = (float64((float32(float64(1e10)) / 3.14159)) - 3.14159)
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
            p0 = (p0 + 135)
            _ = fi9
        }
        var v13 bool = false
        p0 -= p0
        _ = v13
    } else {
        p0 -= (^(p0))
        switch ((("a" + "a") + ("hello" + ((("hello" + "\n\t") + "foo bar") + ""))) + "foo bar") {
        case "foo bar":
            var v14 bool = false
            _ = v14
        case "\n\t":
            p0 -= p0
        case "":
            var v15 int16 = (p0 >> ((724) & 63))
            _ = v15
        default:
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := GenS1{F0: 0, F1: 255}
                fmt.Println(p0)
                _ = fi16
                _ = v17
            }
        }
    }
    _ = p0
    _ = v8
    return -1
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]string)
    var v3 int8 = 127
    var v4 uint64 = 4294967295
    var v5 float32 = 0.1
    var v6 int64 = -9223372036854775808
    v3 -= (v3 + int8(v4))
    switch "\n\t" {
    case "\n\t":
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: 2, F1: 42}
            var v20 string = (("a" + "a") + "")
            _ = fi18
            _ = v19
            _ = v20
        }
    case "hello":
        var v21 uint16 = 508
        _ = v21
    case "foo bar":
        v3 += -1
    }
    v6 -= -9223372036854775808
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int16 = v1
        _ = fi22
        _ = v23
    }
    _ = len(v2)
    v24 := GenS2{F0: "foo bar", F1: "foo bar", F2: 0}
    var v25 uint32 = 1
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
