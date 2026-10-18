package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 int = -9223372036854775808
    v5 := []string{"hello", "hello", "foo bar", ("foo bar" + ("a" + "hello"))}
    v6 := make(map[string]int)
    var v7 int = 32767
    var v8 string = "\n\t"
    var v9 int32 = -2147483648
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            if (!(!false)) {
                if ("a" > "foo bar") {
                    delete(v6, v8)
                    var v12 int = v7
                    var v13 int8 = (v1 << ((732) & 63))
                    _ = v12
                    _ = v13
                } else {
                    switch (v8 + "\n\t") {
                    case "\n\t":
                        v14 := GenS1{F0: -1, F1: 4294967295, F2: 199}
                        fmt.Println(v8, fi10)
                        _ = v14
                    case "foo bar":
                        var v15 int32 = (-(v9))
                        _ = v15
                    case "hello":
                        var v16 uint32 = 4294967295
                        _ = v16
                    }
                    if false {
                        var v17 float32 = 0.1
                        var v18 int64 = 874
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS1{F0: 0, F1: 631, F2: 1}
                        fmt.Println(v4)
                        _ = v19
                    }
                    if true {
                        var v20 int8 = (v1 + (v1 / (((v1 ^ v1)) | 1)))
                        v20 = (v1 >> ((0) & 63))
                        if true {
                            var v21 uint64 = 851
                            _ = v21
                        }
                        _ = v20
                    } else {
                        fmt.Println(v7, v4, fi11)
                        var v22 uint32 = 2147483648
                        var v23 uint16 = 623
                        _ = v22
                        _ = v23
                    }
                    switch ((v8 + v8) + (v8 + v8)) {
                    case "hello":
                        var v24 uint64 = 848
                        _ = v24
                    case "a":
                        v25 := GenS2{F0: 796, F1: 1000}
                        _ = v25
                    case "":
                        v7 += -655
                    default:
                        switch v8 {
                        case "a":
                            _ = len(v5)
                        case "foo bar":
                            var v26 uint32 = 2147483647
                            _ = v26
                        default:
                            v27 := GenS3{F0: 428}
                            _ = v27
                        }
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (v8 > "\n\t") {
                    v7 = fi11
                }
                fmt.Println(fi11, v8, v4)
            } else {
                v28 := GenS2{F0: 502, F1: 4294967295}
                fmt.Println(fi11)
                _ = v28
            }
            if false {
                v29 := GenS4{F0: 348, F1: v9}
                var v30 int8 = v1
                if len(v5) > 0 {
                    v5[0] = (("a" + v8) + "foo bar")
                }
                _ = v29
                _ = v30
            }
            _ = fi11
        }
        v4 = v7
        var v31 string = (((v8 + "") + "a") + (v8 + v8))
        _ = fi10
        _ = v31
    }
    delete(v3, (v4 | (v4 * v7)))
    var v32 int64 = -55
    if false {
        for fi33 := 0; fi33 < 17; fi33++ {
            if true {
                var v34 int32 = (v9 << ((744) & 63))
                delete(v6, v8)
                v5 = append(v5, "foo bar")
                _ = v34
            }
            _ = fi33
        }
        if (2147483647 == 2147483647) {
            var v35 int = (v7 ^ v7)
            if (507 >= 989) {
                v36 := GenS3{F0: 0}
                if (((!(255 == 42)) && (1.5 >= ((0.1 * (1e10 + 0.1)) + (-2.25 + 1.5)))) && (!true)) {
                    if false {
                        v37 := GenS5{F0: 0}
                        var v38 string = (v8 + ("" + ((v8 + v8) + v8)))
                        _ = v37
                        _ = v38
                    }
                } else {
                    var v39 int = v35
                    var v40 int8 = 127
                    _ = v39
                    _ = v40
                }
                _ = v36
            }
            _ = v35
        }
        v4 += 127
    }
    fmt.Println(v8, v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
