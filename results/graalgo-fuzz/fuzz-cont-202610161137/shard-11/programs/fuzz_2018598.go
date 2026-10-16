package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint16, p1 uint) uint64 {
    p1 += 65535
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: (fi5 & (fi5 + (fi5 / ((fi5) | 1))))}
        fmt.Println(p1)
        p0 += 347
        _ = fi5
        _ = v6
    }
    p0 = p0
    p0 = 32767
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 uint16) bool {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(fi8, fi8, p0)
            var v9 string = "\n\t"
            _ = fi8
            _ = v9
        }
        p0 = p0
        var v10 uint16 = (p0 ^ 42)
        _ = fi7
        _ = v10
    }
    var v11 bool = (-2147483648 > 2147483647)
    switch "" {
    case "\n\t":
        v12 := GenS1{F0: 9223372036854775807}
        _ = v12
    }
    _ = p0
    _ = v11
    return (v11 && (65535 >= 835))
}

func main() {
    var v1 int = -695
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 42
    v13 := []int{(v1 | v1)}
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS2{F0: 862, F1: "foo bar"}
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 127
    switch -16 {
    case -9223372036854775808:
        v13 = append(v13, v1)
    case 128:
        switch "\n\t" {
        case "\n\t":
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 string = "\n\t"
                _ = len(v3)
                _ = fi17
                _ = v18
            }
        default:
            v19 := GenS3{F0: "a"}
            _ = v19
        }
    }
    if len(v13) > 0 {
        v13[0] = v1
    }
    v20 := GenS1{F0: 72}
    if false {
        if (("" + "a") > ("foo bar" + "foo bar")) {
            fmt.Println(v1, v16, v4)
            if ("hello" > "") {
                var v21 uint64 = 578
                v22 := GenS2{F0: -365, F1: "\n\t"}
                _ = v21
                _ = v22
            }
        } else {
            if true {
                var v23 int = (v1 + v1)
                delete(v3, 2147483647)
                if true {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        var v25 int16 = -32768
                        v26 := GenS3{F0: ("a" + "")}
                        for fi27 := 0; fi27 < 16; fi27++ {
                            v28 := GenS2{F0: v23, F1: ("hello" + ("hello" + ""))}
                            v13 = append(v13, (fi27 % (((fi24 << ((0) & 63))) | 1)))
                            _ = fi27
                            _ = v28
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                _ = v23
            }
            var v29 uint = 42
            switch ("" + ("foo bar" + "")) {
            case "hello":
                switch v1 {
                case -895:
                    v30 := GenS3{F0: (("\n\t" + "\n\t") + ("foo bar" + "a"))}
                    _ = v30
                case 211:
                    v29 = v29
                default:
                    v2[(-(v1))] = v1
                }
            case "":
                fmt.Println(v29, v29, v29)
            }
            _ = v29
        }
    }
    v4 += v4
    if true {
        fmt.Println(v16)
        if false {
            var v31 int64 = 1
            _ = v31
        }
        var v32 int64 = 1
        _ = v32
    } else {
        for fi33 := 0; fi33 < 20; fi33++ {
            for fi34 := 0; fi34 < 18; fi34++ {
                if false {
                    switch (fi33 * v1) {
                    case 9223372036854775807:
                        v16 = (v16 << ((1) & 63))
                    default:
                        fmt.Println(fi33, v4)
                    }
                }
                v35 := GenS1{F0: v1}
                if (!true) {
                    v36 := GenS4{F0: int64(v4), F1: 240}
                    if false {
                        switch v1 {
                        case 4294967295:
                            v2[-9223372036854775808] = v1
                        default:
                            v4 += (^(v4))
                        }
                    } else {
                        var v37 int8 = -1
                        _ = v37
                    }
                    v2[(-(fi33))] = (v1 ^ (fi33 & fi33))
                    _ = v36
                }
                v16 = v16
                fmt.Println(fi34, fi34)
                fmt.Println(fi34)
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        v4 += v4
        _ = fi38
    }
    var v39 uint16 = 140
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v40 := GenS4{F0: -1, F1: (v39 >> ((uint(v4)) & 63))}
    for fi41 := 0; fi41 < 8; fi41++ {
        if false {
            v42 := GenS5{F0: uint64(v16), F1: 1}
            _ = v42
        } else {
            delete(v2, (fi41 & v1))
            if len(v13) > 0 {
                v13[0] = fi41
            }
        }
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
