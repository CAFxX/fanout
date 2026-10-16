package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int8; F2 int8 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0(p0 uint64) bool {
    v5 := []int{0, 0, 1}
    v6 := GenS1{F0: 2, F1: 1}
    _ = p0
    _ = v5
    _ = v6
    return true
}
func genFn1() int {
    var v7 uint8 = 2
    switch "hello" {
    case "a":
        fmt.Println(v7)
    }
    if (0 <= 617) {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint8 = (^(v7))
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := []int{(fi8 & (fi8 % ((fi10) | 1))), fi10}
                _ = fi10
                _ = v11
            }
            v12 := GenS2{F0: ""}
            _ = fi8
            _ = v12
            _ = v9
        }
        var v13 float32 = (((((3.14159 * 0.1) - (float32(((1e10 * 0.1) - 1.5)) / 3.14159)) + (3.14159 * 1e10)) + (1.5 - (-2.25 / 1.5))) * 0.0)
        v14 := []int{-9223372036854775808, 32768, 128}
        _ = v13
        _ = v14
    } else {
        if (4294967295 < 328) {
            switch ((("hello" + "a") + ("a" + "foo bar")) + "") {
            case "foo bar":
                v7 = v7
            }
        }
        if false {
            v7 = v7
            switch "" {
            case "":
                if (1 >= -9223372036854775808) {
                    v7 += (^(v7))
                } else {
                    fmt.Println(v7)
                    switch (("a" + ("hello" + "hello")) + "a") {
                    case "\n\t":
                        if false {
                            v7 = (v7 / ((v7) | 1))
                            for fi15 := 0; fi15 < 2; fi15++ {
                                if true {
                                    if (1 != int64(v7)) {
                                        v7 = (v7 * v7)
                                    }
                                    v16 := GenS2{F0: ""}
                                    _ = v16
                                } else {
                                    v7 = v7
                                    v7 -= 127
                                }
                                switch "hello" {
                                case "":
                                    v17 := GenS1{F0: 1, F1: 54}
                                    _ = v17
                                }
                                _ = fi15
                            }
                        } else {
                            v7 = v7
                            var v18 uint8 = 0
                            _ = v18
                        }
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        if (((42 >= 1) && false) || false) {
                            v20 := GenS2{F0: "a"}
                            v21 := []int{(fi19 & fi19), fi19}
                            var v22 float32 = 1.5
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        _ = fi19
                    }
                }
            default:
                var v23 bool = (-834 <= -9223372036854775808)
                _ = v23
            }
        }
        v7 = (v7 % ((0) | 1))
    }
    v7 = v7
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 22
    v2 := make(map[int]string)
    var v3 int64 = 9223372036854775807
    v4 := []string{}
    _ = len(v2)
    if len(v4) > 0 {
        v4[0] = (("a" + "a") + ("" + ("a" + (("foo bar" + "a") + "a"))))
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        if false {
            var v25 float64 = ((0.0 + (0.0 / 3.14159)) / 0.1)
            _ = len(v4)
            _ = v25
        }
        v26 := GenS3{F0: 167, F1: v1, F2: v1}
        v3 += (v3 << ((1) & 63))
        _ = fi24
        _ = v26
    }
    switch ("\n\t" + ("\n\t" + ((("a" + ("" + "foo bar")) + "") + "hello"))) {
    case "":
        for fi27 := 0; fi27 < 8; fi27++ {
            _ = len(v2)
            v3 += -1
            _ = fi27
        }
    case "a":
        v4 = append(v4, "a")
    case "hello":
        v3 -= v3
    default:
        v1 += (v1 << ((2) & 63))
    }
    var v28 int8 = v1
    v29 := GenS4{F0: 0, F1: genFn1()}
    v1 += v28
    var v30 uint64 = 18446744073709551615
    _ = len(v2)
    v28 = (v1 * 0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
