package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 int, p1 bool, p2 uint16) int32 {
    fmt.Println(p1, p0, p2)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}
func genFn1(p0 uint16, p1 string) int {
    switch ("a" + ((p1 + ("" + (p1 + p1))) + "a")) {
    case "a":
        fmt.Println(p1, p0)
    case "\n\t":
        var v7 int = 523
        _ = v7
    }
    if false {
        v8 := GenS1{F0: 42, F1: 1000}
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: 4294967295, F1: 32767}
            if (!(!false)) {
                switch ("a" + ("hello" + p1)) {
                case "hello":
                    switch p1 {
                    case "hello":
                        p1 = "\n\t"
                    case "":
                        v11 := []int{fi9, 4294967295, -1}
                        _ = v11
                    case "a":
                        fmt.Println(p1, p0, fi9)
                    }
                }
            } else {
                if true {
                    p1 = p1
                    var v12 float32 = 0.0
                    if (fi9 == 1) {
                        p1 = "a"
                        v13 := []int{fi9, fi9}
                        if false {
                            var v14 uint16 = (p0 / (((^(p0))) | 1))
                            var v15 int64 = 0
                            if false {
                                p0 += (^(p0))
                            } else {
                                var v16 uint8 = 0
                                v14 = v14
                                if (42 != 2) {
                                    var v17 uint32 = 309
                                    _ = v17
                                }
                                _ = v16
                            }
                            _ = v14
                            _ = v15
                        }
                        _ = v13
                    }
                    _ = v12
                }
            }
            v18 := []int{}
            _ = fi9
            _ = v10
            _ = v18
        }
        _ = v8
    } else {
        fmt.Println(p1, p1)
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := []int{(^(fi20)), fi19, (fi20 + (-(fi20))), -445}
            p0 = (^(p0))
            var v22 int32 = genFn0(fi20, (fi20 != (fi20 << ((42) & 63))), 2)
            _ = fi20
            _ = v21
            _ = v22
        }
        if true {
            if (128 >= 1) {
                var v23 int8 = 100
                _ = v23
            }
        } else {
            v24 := []int{fi19, 833, 60}
            var v25 int = (fi19 << ((2) & 63))
            _ = v24
            _ = v25
        }
        _ = fi19
    }
    var v26 int = 2147483647
    var v27 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 4294967295
    v3 := make(map[int]uint32)
    v4 := []int{4294967295, 1, 128, -9223372036854775808}
    var v5 bool = true
    var v6 uint = v2
    v28 := GenS2{F0: -17, F1: -32768, F2: "a"}
    delete(v1, 32768)
    v29 := GenS1{F0: 881, F1: 126}
    var v30 int8 = 42
    var v31 int32 = -1
    fmt.Println(v6, v6)
    var v32 bool = v5
    _ = len(v4)
    switch -271 {
    case 2147483647:
        var v33 uint64 = 706
        _ = v33
    case 457:
        fmt.Println(v31, v6)
    }
    fmt.Println(v2, v31, v32)
    var v34 int16 = 1000
    fmt.Println(v2, v30)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
