package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0() int16 {
    var v5 uint64 = 2
    var v6 string = ""
    for fi7 := 0; fi7 < 3; fi7++ {
        v5 += 18446744073709551615
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 1, F1: 127, F2: v6}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: 42, F1: 77, F2: ((v6 + v6) + "")}
            fmt.Println(v5, v5, v5)
            v5 += (v5 + v5)
            _ = fi10
            _ = v11
        }
        _ = fi7
    }
    fmt.Println(v6, v6)
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []string{"a"}
    var v2 int = 127
    var v3 int = (v2 * v2)
    v4 := []int{}
    v12 := GenS1{F0: 42, F1: 0, F2: ("\n\t" + "\n\t")}
    fmt.Println(v2)
    switch (v2 - (v3 / (((v3 / ((-467) | 1))) | 1))) {
    case 0:
        if false {
            if (genFn0() < -32768) {
                var v13 int64 = -9223372036854775808
                var v14 uint16 = 65535
                _ = v13
                _ = v14
            } else {
                var v15 int = v2
                v16 := GenS1{F0: 42, F1: 255, F2: ("foo bar" + "foo bar")}
                var v17 uint64 = 18446744073709551615
                _ = v15
                _ = v16
                _ = v17
            }
        }
    case 4294967295:
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS2{F0: 42, F1: 42}
            _ = fi18
            _ = v19
        }
    }
    v2 = (v2 / (((v3 / ((v2) | 1))) | 1))
    if (float32(float64(3.14159)) <= (1.5 * 1e10)) {
        var v20 string = "a"
        fmt.Println(v20, v2)
        _ = v20
    } else {
        v3 -= -1
        for fi21 := 0; fi21 < 17; fi21++ {
            v3 = fi21
            switch "hello" {
            case "foo bar":
                switch "" {
                case "\n\t":
                    switch ("a" + "foo bar") {
                    case "hello":
                        fmt.Println(v3, v2, v2)
                    default:
                        if (-1 >= -504) {
                            var v22 uint32 = 719
                            if false {
                                var v23 int32 = 0
                                _ = v23
                            }
                            var v24 float64 = -2.25
                            _ = v22
                            _ = v24
                        } else {
                            var v25 uint = 4294967295
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v3 -= 32768
                                switch "\n\t" {
                                case "":
                                    v27 := GenS1{F0: 984, F1: 128, F2: ((("a" + ("\n\t" + "hello")) + (("" + ("a" + "\n\t")) + "hello")) + "a")}
                                    _ = v27
                                }
                                _ = fi26
                            }
                            var v28 int16 = genFn0()
                            _ = v25
                            _ = v28
                        }
                    }
                case "a":
                    v29 := GenS2{F0: 4294967295, F1: 0}
                    _ = v29
                }
            case "":
                switch (v3 % ((fi21) | 1)) {
                case -247:
                    var v30 int16 = 32767
                    _ = v30
                }
            }
            _ = fi21
        }
        fmt.Println(v3)
    }
    var v31 int16 = 889
    switch v3 {
    case 32768:
        var v32 string = "a"
        _ = v32
    case -895:
        for fi33 := 0; fi33 < 20; fi33++ {
            for fi34 := 0; fi34 < 18; fi34++ {
                v3 -= v2
                fmt.Println(v2, v3)
                v35 := GenS3{F0: 42, F1: fi34, F2: 438}
                fmt.Println(fi33, v31, fi34)
                v31 -= v31
                v2 -= v3
                v31 = v31
                var v36 string = ""
                _ = fi34
                _ = v35
                _ = v36
            }
            var v37 uint8 = 42
            _ = fi33
            _ = v37
        }
    case 128:
        v2 = (v3 % ((2147483647) | 1))
    default:
        v3 = (v3 * (v3 + (v2 * v3)))
    }
    v38 := GenS4{F0: 42, F1: -49}
    var v39 uint32 = 65535
    v39 = v39
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
}
