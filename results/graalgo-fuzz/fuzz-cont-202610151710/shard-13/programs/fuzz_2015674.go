package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{878, 1, 65535}
    v3 := []string{}
    var v4 uint32 = 1
    var v5 uint64 = 159
    v6 := make(map[uint32]string)
    if true {
        v2 = append(v2, v4)
        switch 32767 {
        case -607:
            v7 := GenS1{F0: 98, F1: 1, F2: 0}
            _ = v7
        case 127:
            var v8 int = -9223372036854775808
            _ = v8
        }
    } else {
        v9 := GenS2{F0: 678, F1: v4, F2: false}
        _ = v9
    }
    fmt.Println(v5)
    fmt.Println(v5)
    v5 -= v5
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            if true {
                if ((("" + "a") + ("" + ("" + ("\n\t" + "")))) != ("foo bar" + "hello")) {
                    if (!false) {
                        v4 = 65535
                        if (true || ((v4 * 65535) == v4)) {
                            var v12 int8 = 127
                            v12 += v12
                            v4 += 2147483648
                            _ = v12
                        }
                        v5 += (v5 * 1)
                    }
                } else {
                    v5 = v5
                }
                switch ("hello" + "foo bar") {
                case "foo bar":
                    v13 := GenS2{F0: 255, F1: v4, F2: false}
                    _ = v13
                case "hello":
                    var v14 bool = (false || (3.14159 != (1.5 - (float32((1.5 * 0.1)) * 3.14159))))
                    _ = v14
                case "a":
                    switch "foo bar" {
                    case "foo bar":
                        v3 = append(v3, "a")
                    case "hello":
                        var v15 string = "\n\t"
                        _ = v15
                    case "":
                        v4 += v4
                        if false {
                            v16 := v6[v4]
                            _ = v16
                        } else {
                            fmt.Println(v4)
                            var v17 float64 = (1.5 - 0.1)
                            _ = v17
                        }
                        switch fi10 {
                        case 2147483647:
                            switch "a" {
                            case "hello":
                                v18 := v1[1]
                                _ = v18
                            case "a":
                                var v19 string = ("" + "")
                                _ = v19
                            default:
                                var v20 uint16 = 32767
                                _ = v20
                            }
                        }
                    default:
                        v21 := GenS2{F0: 32767, F1: (v4 | v4), F2: true}
                        _ = v21
                    }
                }
            } else {
                var v22 uint64 = (v5 | uint64(v4))
                fmt.Println(fi10, v5)
                _ = v22
            }
            v23 := GenS3{F0: 1000}
            v24 := GenS3{F0: 32767}
            _ = fi11
            _ = v23
            _ = v24
        }
        _ = fi10
    }
    var v25 float32 = 0.1
    v5 += (v5 + (^(v5)))
    v25 -= ((v25 - v25) * -2.25)
    var v26 float32 = (v25 - 1e10)
    v4 = v4
    for fi27 := 0; fi27 < 3; fi27++ {
        v4 -= (v4 / (((v4 ^ 483)) | 1))
        v28 := v6[(v4 << ((454) & 63))]
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        fmt.Println(v25, v25)
        if (true && (42 < 0)) {
            var v30 int16 = -1
            var v31 int = fi29
            switch (("a" + ("" + "\n\t")) + ("a" + "")) {
            case "\n\t":
                v30 -= v30
            }
            _ = v30
            _ = v31
        } else {
            v26 = v26
            var v32 uint64 = (v5 / ((v5) | 1))
            v33 := GenS1{F0: 15, F1: 4294967295, F2: 443}
            _ = v32
            _ = v33
        }
        fmt.Println(v26, v25, v5)
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
