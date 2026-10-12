package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int16) int32 {
    var v7 int64 = 309
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    fmt.Println(v7, v7)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: p0}
        v12 := []int{(fi10 % ((fi10) | 1))}
        var v13 uint64 = 18446744073709551615
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v7
    return -299
}
func genFn1() uint64 {
    var v14 int32 = -1
    fmt.Println(v14)
    v14 += (v14 * -2147483648)
    var v15 uint64 = 4294967295
    _ = v14
    _ = v15
    return (v15 & v15)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (((!false) || (2147483647 != 294)) || false)
    var v3 float32 = 3.14159
    v4 := make(map[uint32]bool)
    v5 := []string{"hello", "", "foo bar", ("" + "\n\t")}
    v6 := []string{}
    if true {
        switch ("\n\t" + "hello") {
        case "hello":
            v1[2] = -156
        default:
            switch 32767 {
            case 128:
                switch 868 {
                case 85:
                    v16 := GenS1{F0: 42}
                    _ = v16
                case 1:
                    var v17 int = -981
                    _ = v17
                case 9223372036854775807:
                    var v18 int64 = 1
                    _ = v18
                default:
                    if v2 {
                        fmt.Println(v3, v3)
                        v2 = v2
                        v3 = (v3 + 0.0)
                    } else {
                        switch "foo bar" {
                        case "":
                            var v19 uint8 = 42
                            _ = v19
                        case "\n\t":
                            var v20 int64 = 0
                            _ = v20
                        }
                        switch "a" {
                        case "":
                            switch 2147483647 {
                            case -9223372036854775808:
                                var v21 int = 0
                                _ = v21
                            case 202:
                                v2 = (-1 != -58)
                            case 1:
                                fmt.Println(v2, v3)
                            }
                        }
                        var v22 uint32 = 717
                        _ = v22
                    }
                }
            }
        }
        v23 := GenS1{F0: 42}
        v24 := GenS2{F0: 65535, F1: "\n\t", F2: 0}
        _ = v23
        _ = v24
    } else {
        var v25 uint16 = 0
        switch 2147483647 {
        case 128:
            var v26 int16 = 0
            _ = v26
        case 0:
            var v27 bool = true
            _ = v27
        default:
            for fi28 := 0; fi28 < 1; fi28++ {
                switch (-(fi28)) {
                case -617:
                    for fi29 := 0; fi29 < 3; fi29++ {
                        v30 := GenS1{F0: 0}
                        var v31 int16 = 0
                        v2 = (96 != 1)
                        v32 := GenS1{F0: v31}
                        v31 -= -781
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                }
                v33 := GenS2{F0: v25, F1: (("hello" + (("hello" + "\n\t") + ("" + ""))) + "a"), F2: 127}
                fmt.Println(v3, v2, fi28)
                _ = fi28
                _ = v33
            }
        }
        _ = v25
    }
    switch 32767 {
    case 594:
        v2 = (!(71 != -128))
    }
    switch -2147483648 {
    case -2147483648:
        v34 := GenS1{F0: 32767}
        _ = v34
    case 128:
        v35 := GenS1{F0: -1}
        _ = v35
    case 42:
        v3 += float32(1.5)
    default:
        _ = len(v4)
    }
    var v36 uint64 = 1
    v3 -= (v3 - float32(0.1))
    v37 := GenS3{F0: 65535, F1: (^(v36))}
    var v38 int16 = -258
    var v39 bool = (v2 || (-1 <= 42))
    if true {
        v36 += 565
    }
    v40 := GenS1{F0: v38}
    var v41 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
