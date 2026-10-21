package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }

func main() {
    v1 := []int{}
    var v2 int = 2147483647
    v3 := []string{("hello" + "a"), ""}
    v4 := make(map[uint32]int)
    v5 := []int{(v2 / ((1) | 1))}
    v6 := GenS1{F0: 1, F1: 32767}
    var v7 int = v2
    fmt.Println(v2)
    fmt.Println(v7, v2)
    for fi8 := 0; fi8 < 3; fi8++ {
        _ = len(v4)
        v5 = append(v5, (v7 % (((v7 << ((0) & 63))) | 1)))
        _ = fi8
    }
    var v9 string = ("foo bar" + "a")
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int8 = 0
        if true {
            var v12 float64 = -2.25
            switch (v7 * 4294967295) {
            case -719:
                for fi13 := 0; fi13 < 17; fi13++ {
                    v2 = 32767
                    fmt.Println(v7, v11, v9)
                    _ = len(v1)
                    var v14 uint32 = 2147483647
                    _ = fi13
                    _ = v14
                }
            case 1:
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(v12, fi10, v9)
                    fmt.Println(v9, v12, v12)
                    _ = fi15
                }
            }
            _ = v12
        } else {
            switch -863 {
            case 32767:
                v4[65535] = (fi10 + v7)
            case 4294967295:
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int8 = (v11 ^ v11)
                    switch "foo bar" {
                    case "":
                        var v18 int8 = v11
                        _ = v18
                    case "hello":
                        v9 = v9
                    case "foo bar":
                        var v19 uint8 = 2
                        _ = v19
                    default:
                        var v20 int16 = 687
                        _ = v20
                    }
                    var v21 uint = 18446744073709551615
                    v21 = v21
                    if true {
                        switch -1 {
                        case 128:
                            var v22 bool = false
                            _ = v22
                        case 727:
                            v23 := GenS1{F0: 1, F1: 255}
                            _ = v23
                        }
                        var v24 float64 = float64(1e10)
                        var v25 float32 = 0.1
                        _ = v24
                        _ = v25
                    }
                    var v26 uint64 = 58
                    _ = fi16
                    _ = v17
                    _ = v21
                    _ = v26
                }
            case 622:
                for fi27 := 0; fi27 < 8; fi27++ {
                    delete(v4, 65535)
                    v2 += (v7 & fi10)
                    if (9223372036854775807 >= -1) {
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v4[1] = (-(v7))
                        } else {
                            v28 := GenS2{F0: 1}
                            delete(v4, 2147483647)
                            _ = v28
                        }
                        delete(v4, 2147483647)
                        if (42 < 255) {
                            var v29 uint16 = 2
                            v30 := GenS1{F0: 42, F1: v29}
                            var v31 int64 = -9223372036854775808
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            fmt.Println(fi27, v9)
                            var v32 int32 = 2147483647
                            if false {
                                v33 := GenS1{F0: 1, F1: 404}
                                var v34 int32 = 2147483647
                                var v35 int64 = 42
                                var v36 uint = 0
                                var v37 int16 = 1
                                _ = len(v4)
                                var v38 uint64 = 576
                                _ = v33
                                _ = v34
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            _ = v32
                        }
                    }
                    var v39 int32 = 0
                    _ = fi27
                    _ = v39
                }
            default:
                if (2147483647 > 42) {
                    for fi40 := 0; fi40 < 18; fi40++ {
                        v2 = fi40
                        v41 := GenS3{F0: v9}
                        _ = fi40
                        _ = v41
                    }
                    delete(v4, 0)
                    fmt.Println(v7, fi10, v2)
                }
            }
            var v42 string = v9
            if (((0.1 * -2.25) + 0.0) == 3.14159) {
                var v43 int64 = 1
                _ = v43
            }
            _ = v42
        }
        _ = fi10
        _ = v11
    }
    v4[797] = (v2 % ((2147483647) | 1))
    v44 := GenS2{F0: 42}
    var v45 int16 = 1000
    v46 := GenS4{F0: 146}
    switch v7 {
    case 2147483647:
        _ = len(v1)
    }
    switch 709 {
    case -1:
        v7 = v2
    default:
        v9 = "hello"
    }
    _ = len(v1)
    v47 := GenS5{F0: 59}
    delete(v4, 524)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
