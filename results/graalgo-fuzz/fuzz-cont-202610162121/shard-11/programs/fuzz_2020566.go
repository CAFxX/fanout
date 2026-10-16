package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
func genFn0(p0 int8, p1 bool) uint16 {
    if p1 {
        v6 := []int{2147483647, 128, 362, 804}
        v7 := GenS1{F0: (18446744073709551615 < 2), F1: 9223372036854775807, F2: "hello"}
        p1 = (847 == 42)
        _ = v6
        _ = v7
    } else {
        v8 := []int{}
        _ = v8
    }
    p1 = true
    _ = p0
    _ = p1
    return 883
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{"\n\t"}
    var v3 int32 = 866
    var v4 int64 = 0
    v5 := []uint32{925, 4294967295, 0}
    for fi9 := 0; fi9 < 18; fi9++ {
        v4 = (v4 / ((v4) | 1))
        var v10 uint32 = 1
        _ = fi9
        _ = v10
    }
    v1 -= (((v1 + v1) * v1) + 0.0)
    var v11 uint16 = 604
    fmt.Println(v11)
    for fi12 := 0; fi12 < 17; fi12++ {
        v4 = v4
        _ = fi12
    }
    v1 = ((v1 * float32(3.14159)) * v1)
    if (920 > (v3 * v3)) {
        v3 = 42
    }
    if false {
        v1 -= v1
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int = 4294967295
            if ((2 != 4294967295) || true) {
                if false {
                    switch (("" + "a") + "\n\t") {
                    case "":
                        v15 := GenS2{F0: 42, F1: 32767, F2: 0}
                        _ = v15
                    case "a":
                        var v16 string = (("a" + ("\n\t" + "hello")) + ((("hello" + "") + "a") + (("foo bar" + "hello") + "\n\t")))
                        _ = v16
                    default:
                        v5 = append(v5, 2147483647)
                    }
                } else {
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        switch ("hello" + "a") {
                        case "\n\t":
                            v17 := GenS2{F0: 98, F1: -1, F2: -54}
                            _ = v17
                        case "hello":
                            v11 += v11
                        }
                    case "":
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                    case "a":
                        fmt.Println(v1, v11, v3)
                    }
                }
                var v18 int16 = -466
                _ = v18
            } else {
                fmt.Println(v4, v4, v1)
            }
            _ = fi13
            _ = v14
        }
    }
    var v19 uint32 = 4294967295
    v20 := GenS2{F0: -128, F1: 1000, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
