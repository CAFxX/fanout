package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 int16, p1 string, p2 uint64) int16 {
    if false {
        fmt.Println(p0, p2)
        switch 0 {
        case 0:
            p0 += p0
        case -1:
            switch ((p1 + (p1 + "\n\t")) + ((("hello" + (("\n\t" + "\n\t") + "foo bar")) + "\n\t") + ("hello" + p1))) {
            case "\n\t":
                if (-32768 < 32767) {
                    v6 := GenS1{F0: (p0 & p0)}
                    switch "a" {
                    case "hello":
                        for fi7 := 0; fi7 < 20; fi7++ {
                            var v8 int32 = 42
                            _ = fi7
                            _ = v8
                        }
                    case "\n\t":
                        var v9 int64 = -1
                        _ = v9
                    default:
                        switch p1 {
                        case "hello":
                            for fi10 := 0; fi10 < 18; fi10++ {
                                switch (p1 + p1) {
                                case "":
                                    v11 := GenS1{F0: 32767}
                                    _ = v11
                                case "a":
                                    var v12 int8 = 1
                                    _ = v12
                                default:
                                    v13 := GenS1{F0: p0}
                                    _ = v13
                                }
                                _ = fi10
                            }
                        case "":
                            p1 = ((("\n\t" + p1) + p1) + p1)
                        }
                    }
                    p1 = "a"
                    _ = v6
                }
            default:
                var v14 uint64 = (p2 | p2)
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 17; fi15++ {
                p0 = 32767
                v16 := []int{}
                _ = fi15
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        if true {
            var v18 bool = true
            v18 = ((0.1 / 1e10) < (3.14159 - 1e10))
            v19 := []int{}
            _ = v18
            _ = v19
        }
        v20 := GenS1{F0: (p0 + (^(p0)))}
        _ = fi17
        _ = v20
    }
    switch p1 {
    case "a":
        v21 := GenS1{F0: p0}
        _ = v21
    case "hello":
        fmt.Println(p0, p2, p2)
    case "foo bar":
        v22 := []int{0}
        _ = v22
    }
    v23 := []int{0, 8}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return -32768
}
func genFn1(p0 uint32, p1 string) uint8 {
    for fi24 := 0; fi24 < 19; fi24++ {
        p1 = "\n\t"
        _ = fi24
    }
    var v25 uint = 42
    _ = p0
    _ = p1
    _ = v25
    return 255
}

func main() {
    var v1 uint8 = 220
    var v2 uint16 = 32767
    var v3 float32 = -2.25
    v4 := []string{}
    v5 := make(map[string]int)
    for fi26 := 0; fi26 < 5; fi26++ {
        v2 += v2
        v27 := GenS1{F0: -411}
        _ = fi26
        _ = v27
    }
    if (!(2 <= 65535)) {
        var v28 uint64 = 2
        v2 = 32767
        _ = v28
    }
    v4 = append(v4, ("\n\t" + "hello"))
    var v29 uint64 = 2
    var v30 uint32 = 2147483648
    switch 782 {
    case -167:
        v5[""] = -9223372036854775808
    }
    v31 := GenS2{F0: (v1 & (v1 | (v1 >> ((2) & 63)))), F1: 1, F2: ((("\n\t" + ("a" + ("" + "hello"))) + "hello") + "hello")}
    var v32 uint16 = 697
    v29 -= v29
    v1 = 255
    switch 871 {
    case 1:
        v4 = append(v4, "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
