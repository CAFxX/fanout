package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 int32, p1 uint16, p2 uint8) uint64 {
    p0 = (p0 + 42)
    p1 = p1
    var v4 int = 32768
    fmt.Println(v4, v4, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 281
}
func genFn1() int {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    switch ("hello" + "foo bar") {
    case "a":
        var v6 int8 = 1
        _ = v6
    case "foo bar":
        var v7 bool = (((("a" + ("hello" + "\n\t")) + "") >= "foo bar") || false)
        _ = v7
    case "hello":
        v8 := []int{-2147483648, 98, -9223372036854775808, 32767}
        _ = v8
    default:
        var v9 float64 = (3.14159 + -2.25)
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint64 = 807
        if false {
            if (!(1 >= -2147483648)) {
                if (!true) {
                    var v12 bool = false
                    v13 := []int{(-(fi10)), (fi10 << ((255) & 63))}
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = v12
                    _ = v13
                }
                var v14 uint16 = 42
                _ = v14
            } else {
                v15 := []int{(fi10 << ((18446744073709551615) & 63))}
                for fi16 := 0; fi16 < 20; fi16++ {
                    switch "foo bar" {
                    case "hello":
                        switch 2147483647 {
                        case 2147483647:
                            _ = len(v15)
                        case -852:
                            v17 := GenS1{F0: 0, F1: ""}
                            _ = v17
                        }
                    case "a":
                        if true {
                            v18 := GenS2{F0: -2147483648}
                            v19 := GenS3{F0: 42, F1: 42}
                            _ = v18
                            _ = v19
                        } else {
                            var v20 int64 = -1
                            _ = v20
                        }
                    }
                    v15 = append(v15, fi16)
                    v21 := GenS2{F0: -2147483648}
                    _ = fi16
                    _ = v21
                }
                for fi22 := 0; fi22 < 19; fi22++ {
                    v11 += genFn0(745, 255, 0)
                    switch ("\n\t" + ("hello" + ("a" + ""))) {
                    case "hello":
                        if true {
                            v11 -= (v11 + genFn0(-2147483648, 65535, 128))
                            var v23 int8 = 42
                            _ = v23
                        }
                    default:
                        v11 += 18446744073709551615
                    }
                    _ = fi22
                }
                _ = v15
            }
            var v24 uint8 = 128
            var v25 uint8 = v24
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 uint16 = 811
                v27 = 916
                _ = fi26
                _ = v27
            }
        }
        _ = fi10
        _ = v11
    }
    v28 := GenS3{F0: 1, F1: 255}
    if (2 < 751) {
        v29 := GenS2{F0: 2147483647}
        v30 := []int{755, 128}
        _ = v29
        _ = v30
    }
    _ = v28
    return 861
}

func main() {
    var v1 uint32 = 65535
    var v2 float64 = ((1.5 / float64(float32(float64(1.5)))) + (1e10 + -2.25))
    v3 := []int{127, -364, -397, 1}
    var v31 string = (("hello" + "") + "hello")
    v31 = "hello"
    v32 := GenS3{F0: 0, F1: 42}
    var v33 uint32 = 1
    if true {
        v33 += v1
        var v34 float64 = -2.25
        v2 = 0.1
        _ = v34
    }
    if (!true) {
        var v35 uint64 = 4294967295
        _ = v35
    }
    var v36 int8 = -44
    v3 = append(v3, -9223372036854775808)
    fmt.Println(v2, v31, v2)
    switch 4294967295 {
    case 1:
        v2 = 1e10
    case 32767:
        v37 := GenS2{F0: -1}
        _ = v37
    }
    var v38 uint64 = 42
    v2 = 0.1
    v39 := GenS2{F0: -1}
    if true {
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 uint64 = (v38 + (v38 | (v38 - v38)))
            _ = len(v3)
            _ = fi40
            _ = v41
        }
    } else {
        var v42 int64 = -9223372036854775808
        for fi43 := 0; fi43 < 10; fi43++ {
            var v44 int64 = v42
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) > 0 {
                v3[0] = fi43
            }
            _ = fi43
            _ = v44
        }
        fmt.Println(v42, v42)
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v38
    _ = v39
}
