package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
func genFn0() int16 {
    v4 := []int{127, 0, 0, 0}
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    _ = v4
    return -32768
}
func genFn1(p0 int, p1 uint8, p2 uint) uint8 {
    v5 := []int{808, (p0 ^ 4294967295), p0}
    p2 = (p2 | 591)
    if (201 <= 65535) {
        p1 += (p1 ^ p1)
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: 193, F1: 169, F2: false}
            _ = fi6
            _ = v7
        }
        var v8 int16 = genFn0()
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            if (-9223372036854775808 <= -245) {
                var v10 string = (("\n\t" + "a") + "\n\t")
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 int = 1
                    v12 += v12
                    switch (("a" + "") + ("a" + "a")) {
                    case "":
                        p0 = fi11
                    case "a":
                        var v13 int16 = 453
                        _ = v13
                    case "hello":
                        _ = len(v5)
                    }
                    _ = fi11
                    _ = v12
                }
                for fi14 := 0; fi14 < 4; fi14++ {
                    p0 -= p0
                    _ = fi14
                }
                _ = v10
            } else {
                if (p2 != p2) {
                    p2 -= 65535
                    v15 := GenS1{F0: 1, F1: -32768, F2: true}
                    v16 := GenS1{F0: 65535, F1: 693, F2: (1 <= -1)}
                    _ = v15
                    _ = v16
                } else {
                    p1 += p1
                }
            }
            p2 = p2
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    v1 := []uint32{114, 2147483648, 1}
    var v2 uint8 = 1
    v3 := make(map[string]string)
    var v17 uint64 = 42
    fmt.Println(v2, v2)
    v17 -= 0
    for fi18 := 0; fi18 < 12; fi18++ {
        delete(v3, "hello")
        _ = len(v1)
        v3["hello"] = ""
        _ = fi18
    }
    delete(v3, "foo bar")
    if (0.0 < 1e10) {
        _ = len(v3)
        var v19 float64 = 1.5
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = v19
    }
    var v20 float32 = float32(1e10)
    v21 := v3[(("" + "\n\t") + "")]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
