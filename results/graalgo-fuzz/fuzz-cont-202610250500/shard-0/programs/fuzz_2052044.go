package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint16 }
func genFn0() int64 {
    var v5 uint = 0
    v6 := []int{4294967295, 304, -433, 4294967295}
    fmt.Println(v5)
    var v7 uint64 = 2
    fmt.Println(v5)
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := []string{}
    var v2 uint16 = 2
    var v3 float64 = (1e10 * float64(-2.25))
    v4 := []string{(("" + "foo bar") + ("a" + ("foo bar" + ""))), ("foo bar" + "a"), "foo bar"}
    var v8 float32 = ((((3.14159 / (1.5 - -2.25)) / 1.5) / 1e10) - (-2.25 / 1.0))
    fmt.Println(v2, v3, v8)
    fmt.Println(v2, v2)
    if (!true) {
        var v9 string = "a"
        _ = v9
    }
    if len(v4) > 0 {
        v4[0] = ("" + ("" + (("foo bar" + "a") + "\n\t")))
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if true {
            for fi11 := 0; fi11 < 9; fi11++ {
                v3 = (float64(-2.25) / v3)
                _ = fi11
            }
            v8 += ((1.5 - float32((v3 - 1.5))) / v8)
        } else {
            v12 := GenS1{F0: 4294967295, F1: ("\n\t" + ("a" + "a")), F2: 32767}
            v8 = ((v8 - (float32(float64(((v8 - float32(v3)) + v8))) - 0.0)) / v8)
            v2 = 32767
            _ = v12
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi10
    }
    var v13 uint = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "a" {
    case "":
        var v14 uint16 = v2
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 10; fi15++ {
            if true {
                v16 := GenS1{F0: 4294967295, F1: "", F2: 32767}
                v3 -= (3.14159 / v3)
                var v17 int64 = -532
                _ = v16
                _ = v17
            }
            v8 = ((3.14159 + v8) * -2.25)
            var v18 uint32 = 2
            _ = fi15
            _ = v18
        }
    case "foo bar":
        v4 = append(v4, "hello")
    }
    v8 = 0.0
    v13 = 18446744073709551615
    v4 = append(v4, "\n\t")
    v19 := GenS2{F0: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
