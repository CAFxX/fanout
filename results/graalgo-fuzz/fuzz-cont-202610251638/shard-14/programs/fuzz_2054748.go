package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int64; F2 uint64 }
func genFn0(p0 int8, p1 uint16) uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            p1 += p1
            _ = fi7
        }
        _ = fi6
    }
    var v8 uint16 = p1
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := []int{fi9, fi9, -1}
            v8 += v8
            var v12 int32 = 174
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    if (-9223372036854775808 != 650) {
        var v13 int = 4294967295
        _ = v13
    } else {
        var v14 uint64 = 42
        v15 := []int{-2147483648, 32768, 32767}
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v8
    return 834
}
func genFn1(p0 uint8, p1 uint64, p2 uint16) int8 {
    v16 := GenS1{F0: -128}
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS2{F0: 9223372036854775807, F1: -671, F2: p1}
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 42
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    var v5 uint = 18446744073709551615
    var v19 uint = 766
    v19 += (v5 + (^(v5)))
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    fmt.Println(v19, v19)
    _ = len(v1)
    v19 -= 4294967295
    var v20 uint8 = 70
    for fi21 := 0; fi21 < 5; fi21++ {
        v5 -= (^(v19))
        var v22 int = (fi21 + -720)
        if false {
            var v23 uint64 = 18446744073709551615
            v2["\n\t"] = 1
            _ = len(v4)
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    switch "a" {
    case "a":
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 int32 = 724
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 int8 = -1
                if len(v1) > 0 {
                    v1[0] = 984
                }
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    }
    if true {
        if false {
            v19 = (v5 + v19)
        }
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            if (true || false) {
                switch ("foo bar" + ("hello" + ((("foo bar" + "hello") + (("a" + "foo bar") + ("foo bar" + "foo bar"))) + ("\n\t" + (((((("a" + ("\n\t" + ("a" + ""))) + "hello") + "hello") + ("" + ((("" + "a") + "hello") + "foo bar"))) + "a") + "hello"))))) {
                case "hello":
                    v20 = v20
                }
                if false {
                    for fi29 := 0; fi29 < 4; fi29++ {
                        fmt.Println(v20)
                        var v30 bool = true
                        var v31 bool = (1 <= 9223372036854775807)
                        if (true || v31) {
                            var v32 int32 = -2147483648
                            if (!v31) {
                                var v33 float64 = 3.14159
                                _ = v33
                            }
                            _ = v32
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                } else {
                    var v34 int = (fi28 % ((32768) | 1))
                    _ = v34
                }
                var v35 string = "hello"
                _ = v35
            } else {
                _ = len(v4)
                fmt.Println(v5)
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v20 = (v20 | v20)
            _ = fi28
        }
    }
    v5 += v5
    var v36 uint64 = 753
    fmt.Println(v19, v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
