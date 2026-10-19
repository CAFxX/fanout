package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint32 }
func genFn0(p0 uint64, p1 uint8) int8 {
    v5 := []int{9223372036854775807, 1, 9223372036854775807}
    p1 += p1
    for fi6 := 0; fi6 < 19; fi6++ {
        switch (("" + "hello") + "") {
        case "":
            for fi7 := 0; fi7 < 1; fi7++ {
                v5 = append(v5, fi6)
                var v8 uint8 = (p1 << ((502) & 63))
                _ = fi7
                _ = v8
            }
        case "a":
            v9 := GenS1{F0: "", F1: 127, F2: 389}
            _ = v9
        case "\n\t":
            var v10 int16 = 32767
            _ = v10
        default:
            var v11 string = ("" + ("hello" + "foo bar"))
            _ = v11
        }
        if (("a" != "a") || true) {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 string = (("hello" + "hello") + (((("\n\t" + "a") + "foo bar") + ((("" + "\n\t") + "hello") + "\n\t")) + ("\n\t" + "")))
                _ = fi12
                _ = v13
            }
            fmt.Println(p0, fi6)
            fmt.Println(fi6)
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 46
}

func main() {
    var v1 uint16 = 735
    var v2 uint8 = 63
    var v3 uint = 4294967295
    var v4 int = 159
    v4 -= -335
    var v14 int32 = 42
    v2 += v2
    v1 = (v1 % ((v1) | 1))
    v15 := []int{(v4 ^ v4), v4, v4, v4}
    v3 = 18446744073709551615
    if len(v15) > 0 {
        v15[0] = v4
    }
    switch "a" {
    case "a":
        if false {
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(v4, v4, v4)
                v3 -= (^(v3))
                _ = len(v15)
                _ = fi16
            }
            v1 += v1
        } else {
            switch v4 {
            case 0:
                v14 -= (v14 & v14)
            case -419:
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            }
            switch (-(v4)) {
            case 33:
                v15 = append(v15, 0)
            }
            var v17 int16 = 32767
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
