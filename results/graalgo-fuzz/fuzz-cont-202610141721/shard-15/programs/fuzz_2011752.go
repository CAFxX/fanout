package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 0
    var v3 int32 = 1
    var v4 uint = 459
    v5 := []uint32{864, 880}
    v6 := []string{("hello" + ("foo bar" + ("" + "\n\t"))), "a", (("" + "") + "\n\t"), "hello"}
    delete(v1, "foo bar")
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = -1
        _ = fi7
        _ = v8
    }
    v3 -= v3
    v2 = v2
    v2 += -128
    v4 = (v4 * (v4 % ((v4) | 1)))
    for fi9 := 0; fi9 < 14; fi9++ {
        v4 = (v4 * v4)
        _ = fi9
    }
    v1[("foo bar" + (("" + "foo bar") + "a"))] = (("a" + "foo bar") + "hello")
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: 318}
        for fi12 := 0; fi12 < 15; fi12++ {
            v2 = 42
            v13 := GenS1{F0: 453}
            v2 -= 100
            _ = fi12
            _ = v13
        }
        var v14 uint32 = 26
        _ = fi10
        _ = v11
        _ = v14
    }
    v15 := GenS1{F0: 2147483647}
    switch "" {
    case "foo bar":
        switch (("" + "foo bar") + "hello") {
        case "a":
            v4 += v4
        }
    }
    if true {
        var v16 uint8 = 42
        _ = v16
    } else {
        switch 1 {
        case 2147483647:
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 float32 = -2.25
                var v19 int16 = 131
                v2 += 1
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        var v20 uint = v4
        v2 -= (^(v2))
        _ = v20
    }
    fmt.Println(v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v21 float32 = 1e10
    switch "\n\t" {
    case "\n\t":
        for fi22 := 0; fi22 < 8; fi22++ {
            if true {
                if true {
                    var v23 int32 = (v3 * v3)
                    v2 += (v2 * v2)
                    _ = v23
                }
                var v24 string = ""
                var v25 int64 = 1
                _ = v24
                _ = v25
            } else {
                fmt.Println(fi22)
                if (!false) {
                    v26 := GenS1{F0: 215}
                    for fi27 := 0; fi27 < 6; fi27++ {
                        var v28 uint8 = 128
                        if ((!true) && true) {
                            var v29 int32 = (v3 | (v3 % ((v3) | 1)))
                            v6 = append(v6, "a")
                            var v30 float32 = (((0.0 * 3.14159) - v21) - 1.5)
                            _ = len(v1)
                            _ = v29
                            _ = v30
                        }
                        fmt.Println(fi27, v21)
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                }
            }
            v21 += float32((0.1 - 0.1))
            for fi31 := 0; fi31 < 13; fi31++ {
                v3 = v3
                var v32 bool = false
                v32 = (1 <= 0)
                _ = fi31
                _ = v32
            }
            _ = fi22
        }
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
