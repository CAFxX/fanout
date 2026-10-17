package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{2147483648, 2}
    var v3 float64 = float64(((float32(0.1) / 1.5) * ((1e10 + -2.25) * float32(0.1))))
    v4 := []int{}
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 uint8 = 128
        _ = fi5
        _ = v6
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    if (1 > 2147483647) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v4) > 0 {
            v4[0] = -1
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            switch fi7 {
            case -1:
                v8 := GenS1{F0: 255, F1: "a", F2: 0}
                _ = v8
            default:
                delete(v1, "hello")
            }
            v3 = float64(-2.25)
            _ = fi7
        }
    }
    v9 := GenS2{F0: 42}
    _ = len(v2)
    v3 -= 1.5
    v1[("" + "")] = ("" + (("\n\t" + "\n\t") + ""))
    v10 := GenS1{F0: 445, F1: ("hello" + ("\n\t" + (("foo bar" + (("hello" + ("a" + "foo bar")) + "foo bar")) + "hello"))), F2: 65535}
    v3 = (1e10 * v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
