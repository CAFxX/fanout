package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS5 struct { F0 int; F1 int; F2 string }
func genFn0(p0 string, p1 string, p2 int8) uint32 {
    p1 = p1
    var v6 uint32 = 65535
    var v7 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v6 | v6)
}
func genFn1(p0 int, p1 uint16) int16 {
    v8 := GenS1{F0: (p0 / ((p0) | 1))}
    switch p0 {
    case -147:
        var v9 uint16 = p1
        _ = v9
    case 128:
        fmt.Println(p1)
    case 9223372036854775807:
        switch "foo bar" {
        case "hello":
            switch p0 {
            case -9223372036854775808:
                if (648 < 65535) {
                    var v10 int16 = -1
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v12 := GenS1{F0: (p0 % ((p0) | 1))}
                        var v13 uint = 18446744073709551615
                        if false {
                            if true {
                                p0 -= fi11
                                switch (("hello" + "") + "hello") {
                                case "a":
                                    var v14 uint8 = 2
                                    _ = v14
                                }
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    var v16 uint64 = 18446744073709551615
                                    switch 410 {
                                    case 42:
                                        if false {
                                            fmt.Println(p0)
                                        } else {
                                            v13 += (^(v13))
                                            v17 := GenS2{F0: ("hello" + "\n\t")}
                                            v18 := []int{(-(p0)), p0, 971, 32767}
                                            _ = v17
                                            _ = v18
                                        }
                                    case -9223372036854775808:
                                        v19 := GenS3{F0: 18446744073709551615, F1: 0}
                                        _ = v19
                                    case 127:
                                        var v20 int16 = (v10 ^ v10)
                                        _ = v20
                                    }
                                    var v21 int64 = -1
                                    _ = fi15
                                    _ = v16
                                    _ = v21
                                }
                            }
                            var v22 uint16 = 300
                            if (!false) {
                                var v23 int64 = 42
                                var v24 bool = (-2147483648 >= 2147483647)
                                v25 := GenS4{F0: 143, F1: 0, F2: v22}
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                var v26 int32 = 0
                                v13 = 316
                                _ = v26
                            }
                            _ = v22
                        } else {
                            v27 := GenS1{F0: p0}
                            v28 := GenS1{F0: fi11}
                            _ = v27
                            _ = v28
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v10 += (v10 % ((v10) | 1))
                    _ = v10
                } else {
                    fmt.Println(p0)
                    p1 = 2
                    var v29 uint16 = (p1 >> ((2) & 63))
                    _ = v29
                }
            default:
                p1 += p1
            }
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (!(!(0 > 2)))
    var v3 int32 = 374
    v4 := []string{"\n\t", "a"}
    var v5 uint16 = 65535
    var v30 float32 = 1e10
    var v31 uint32 = 2147483647
    v32 := GenS2{F0: ("\n\t" + "\n\t")}
    v33 := GenS5{F0: 128, F1: 0, F2: ""}
    _ = len(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        var v34 int16 = 1
        if ((v30 / float32(1e10)) >= v30) {
            var v35 uint8 = 0
            v34 += -609
            _ = v35
        } else {
            var v36 string = "hello"
            _ = v36
        }
        _ = v34
    } else {
        switch 611 {
        case 9223372036854775807:
            for fi37 := 0; fi37 < 16; fi37++ {
                v3 = v3
                switch 1 {
                case -620:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v31)
                    }
                case 42:
                    fmt.Println(v2, fi37, v2)
                }
                _ = fi37
            }
        }
        fmt.Println(v2, v31)
    }
    _ = len(v1)
    var v38 float32 = 3.14159
    switch "hello" {
    case "\n\t":
        switch -10 {
        case -2147483648:
            var v39 int32 = v3
            _ = v39
        default:
            fmt.Println(v5, v38)
        }
    }
    var v40 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
