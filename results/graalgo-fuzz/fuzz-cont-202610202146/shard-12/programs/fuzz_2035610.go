package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string) uint32 {
    v6 := []int{32768, 32767, 0, 1}
    fmt.Println(p0)
    v7 := GenS1{F0: 127, F1: 2147483647}
    var v8 int32 = -2147483648
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint16, p1 uint64, p2 uint8) int8 {
    v9 := GenS1{F0: -1, F1: 2147483647}
    p1 = 254
    var v10 int = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []int{1, 1}
    v2 := []int{1}
    var v3 uint16 = 920
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v3 = v3
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint8 = 31
        var v13 uint32 = 608
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint64 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 float32 = (1.5 * 0.0)
    if (-126 != -846) {
        switch "" {
        case "":
            v16 := GenS1{F0: 65, F1: 1}
            _ = v16
        case "\n\t":
            v17 := GenS1{F0: -32, F1: -1}
            _ = v17
        }
        v18 := GenS2{F0: 18446744073709551615}
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
        _ = v18
    }
    var v19 uint = 1
    var v20 string = (((("\n\t" + ("hello" + "")) + "hello") + "\n\t") + ("a" + "hello"))
    fmt.Println(v19, v15)
    v14 -= v14
    if true {
        v19 = 4294967295
        _ = len(v5)
        var v21 int16 = 1
        _ = v21
    }
    v22 := GenS2{F0: v14}
    if false {
        var v23 uint = 18446744073709551615
        _ = v23
    }
    if true {
        for fi24 := 0; fi24 < 17; fi24++ {
            _ = len(v2)
            var v25 uint = (v19 ^ v19)
            _ = fi24
            _ = v25
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            switch ("hello" + "") {
            case "hello":
                switch (fi26 % ((fi26) | 1)) {
                case -501:
                    var v27 string = "foo bar"
                    _ = v27
                case 128:
                    _ = len(v4)
                default:
                    v28 := GenS1{F0: 93, F1: -2147483648}
                    _ = v28
                }
            case "foo bar":
                v19 -= 0
            case "":
                if false {
                    switch ((v20 + (((v20 + "foo bar") + v20) + "\n\t")) + (("a" + "") + "foo bar")) {
                    case "":
                        var v29 int = (fi26 % (((fi26 * int(v14))) | 1))
                        _ = v29
                    case "foo bar":
                        v20 = (((v20 + ((v20 + "foo bar") + (v20 + ((v20 + "") + v20)))) + "") + (v20 + v20))
                    case "a":
                        v3 += (v3 * 117)
                    }
                    _ = len(v2)
                    var v30 uint8 = 128
                    _ = v30
                }
            }
            _ = fi26
        }
        var v31 int = 14
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
