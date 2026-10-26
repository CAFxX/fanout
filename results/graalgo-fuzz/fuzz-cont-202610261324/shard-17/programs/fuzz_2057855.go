package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint16) bool {
    v7 := GenS1{F0: 0, F1: -9223372036854775808}
    v8 := GenS1{F0: 47, F1: 1}
    var v9 float32 = 0.0
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int32, p1 uint16) int64 {
    var v10 string = ("a" + ("" + "foo bar"))
    for fi11 := 0; fi11 < 7; fi11++ {
        if true {
            if ("foo bar" < v10) {
                p0 = 2147483647
                v12 := []int{fi11, (-(fi11)), (fi11 >> ((182) & 63))}
                _ = v12
            } else {
                v13 := GenS1{F0: 130, F1: -951}
                p1 -= p1
                var v14 uint32 = 510
                _ = v13
                _ = v14
            }
        } else {
            p0 = (^(p0))
            if false {
                p0 -= (p0 ^ (^(p0)))
                if true {
                    switch (v10 + "\n\t") {
                    case "":
                        p0 -= (p0 * 889)
                    default:
                        var v15 int16 = -428
                        _ = v15
                    }
                }
            } else {
                switch fi11 {
                case 1:
                    if false {
                        v16 := GenS2{F0: 42}
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v18 := []int{763}
                            var v19 bool = false
                            var v20 int = fi17
                            fmt.Println(v10, v20, v19)
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        _ = v16
                    }
                case 9223372036854775807:
                    for fi21 := 0; fi21 < 1; fi21++ {
                        if true {
                            p1 = (p1 * p1)
                            fmt.Println(v10)
                        }
                        switch fi11 {
                        case 0:
                            fmt.Println(v10, fi11, v10)
                        case 9223372036854775807:
                            p0 += (-(p0))
                        case 854:
                            if (!("foo bar" > v10)) {
                                v22 := []int{}
                                v22 = append(v22, 9223372036854775807)
                                var v23 uint8 = 255
                                _ = v22
                                _ = v23
                            }
                        }
                        var v24 uint8 = 128
                        v25 := []int{fi21, fi21, 128, 2147483647}
                        _ = fi21
                        _ = v24
                        _ = v25
                    }
                }
            }
            var v26 bool = false
            _ = v26
        }
        p0 += -2147483648
        _ = fi11
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v10
    return -749
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 65535
    var v3 uint64 = 18446744073709551615
    var v4 bool = true
    v5 := make(map[int]bool)
    var v6 bool = (!((!true) && true))
    switch 0 {
    case 1:
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 int = (fi27 - 748)
            _ = fi27
            _ = v28
        }
    }
    var v29 uint32 = 2
    if (!v6) {
        v30 := GenS3{F0: 0}
        _ = v30
    } else {
        v31 := GenS4{F0: 18446744073709551615}
        v32 := []int{}
        if false {
            v4 = (!v4)
            var v33 uint8 = 128
            _ = v33
        }
        _ = v31
        _ = v32
    }
    var v34 uint32 = v29
    v35 := GenS2{F0: 150}
    v2 = (v2 - 2147483647)
    fmt.Println(v29, v6, v4)
    var v36 float64 = -2.25
    var v37 int16 = 0
    fmt.Println(v4)
    fmt.Println(v34, v34)
    switch 32767 {
    case 9223372036854775807:
        switch "a" {
        case "hello":
            _ = len(v5)
        case "foo bar":
            var v38 int16 = (v37 | v37)
            _ = v38
        case "a":
            v37 -= (v37 >> ((4294967295) & 63))
        default:
            if ("hello" > (("foo bar" + "") + "a")) {
                v37 -= v37
                v39 := GenS1{F0: 60, F1: 1}
                v40 := GenS4{F0: 18446744073709551615}
                _ = v39
                _ = v40
            } else {
                fmt.Println(v37)
                v6 = v4
                v41 := GenS2{F0: 1}
                _ = v41
            }
        }
    case 689:
        if (((^(v3)) <= (^(v3))) || false) {
            v42 := GenS4{F0: 477}
            var v43 uint16 = 833
            v4 = v6
            _ = v42
            _ = v43
        }
    }
    v34 = v29
    var v44 bool = v4
    fmt.Println(v2, v36)
    fmt.Println(v34, v36, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
