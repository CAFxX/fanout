package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 string; F2 uint16 }
type GenS5 struct { F0 int8; F1 int8; F2 int64 }
func genFn0() uint8 {
    if true {
        var v6 uint = 2
        v7 := []int{int(v6), -1, -180, 32768}
        _ = v6
        _ = v7
    } else {
        v8 := []int{32768}
        _ = v8
    }
    var v9 int16 = 572
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: 32767, F1: 31}
        fmt.Println(v9, v9)
        var v12 uint32 = 65535
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 255}
    _ = v13
    _ = v9
    return 42
}
func genFn1(p0 int16, p1 uint16) uint64 {
    var v14 int16 = 1000
    var v15 uint64 = 1
    switch -2147483648 {
    case 42:
        v15 += v15
    case 9223372036854775807:
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                if false {
                    switch ("" + "") {
                    case "a":
                        if false {
                            v14 = v14
                            v14 = -1
                        }
                    case "\n\t":
                        var v18 uint8 = 227
                        _ = v18
                    }
                    v15 -= 7
                    switch (^(fi16)) {
                    case -624:
                        v19 := GenS3{F0: 42}
                        _ = v19
                    }
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v20 int32 = -1
                        var v21 int8 = 0
                        v21 -= (v21 << ((255) & 63))
                        v22 := []int{fi17, (fi16 / (((-(fi17))) | 1)), fi16, (^(fi17))}
                        _ = v20
                        _ = v21
                        _ = v22
                    case "foo bar":
                        if true {
                            v23 := GenS3{F0: 42}
                            fmt.Println(p0, v15)
                            _ = v23
                        }
                    case "hello":
                        if (42 < -9223372036854775808) {
                            v24 := GenS4{F0: "\n\t", F1: "hello", F2: (p1 % ((p1) | 1))}
                            _ = v24
                        } else {
                            v14 += p0
                            var v25 uint64 = (v15 | 1)
                            _ = v25
                        }
                    default:
                        switch (-(fi17)) {
                        case 551:
                            if ((p1 & (p1 >> ((405) & 63))) > p1) {
                                switch 127 {
                                case 948:
                                    var v26 bool = ("hello" == "hello")
                                    _ = v26
                                case 42:
                                    fmt.Println(fi17, fi17, fi17)
                                case 32768:
                                    var v27 uint16 = p1
                                    _ = v27
                                default:
                                    var v28 int16 = v14
                                    _ = v28
                                }
                            } else {
                                v29 := []int{}
                                _ = v29
                            }
                        }
                    }
                    v14 -= p0
                }
                v30 := GenS5{F0: -84, F1: 100, F2: 606}
                var v31 bool = true
                _ = fi17
                _ = v30
                _ = v31
            }
            _ = fi16
        }
    default:
        p1 = (p1 - p1)
    }
    var v32 float32 = -2.25
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v32
    return (v15 ^ (v15 >> ((692) & 63)))
}

func main() {
    var v1 uint32 = 42
    v2 := []string{"\n\t", ("\n\t" + ((("hello" + "a") + "a") + "foo bar")), "foo bar", ""}
    v3 := make(map[string]string)
    var v4 uint16 = 831
    var v5 float32 = (0.0 * ((0.1 - (1e10 * 3.14159)) * ((0.0 * 0.1) / 1e10)))
    if true {
        var v33 int32 = -1
        fmt.Println(v33, v1, v33)
        _ = v33
    }
    v1 = v1
    var v34 uint = 65535
    for fi35 := 0; fi35 < 20; fi35++ {
        v34 = v34
        var v36 int8 = 0
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 int16 = -892
        v4 -= v4
        v4 = v4
        _ = fi37
        _ = v38
    }
    var v39 int64 = -9223372036854775808
    v5 -= v5
    var v40 int32 = 2147483647
    var v41 int = -990
    fmt.Println(v1, v34)
    if ((("foo bar" + "hello") + ("\n\t" + ("" + ("" + "\n\t")))) <= ("hello" + "a")) {
        v34 -= v34
        for fi42 := 0; fi42 < 7; fi42++ {
            v40 -= 42
            v39 = 42
            v1 = v1
            _ = fi42
        }
        v2 = append(v2, "\n\t")
    } else {
        v3["foo bar"] = ""
        fmt.Println(v4, v34, v41)
    }
    if ((v4 % ((v4) | 1)) > (v4 + v4)) {
        fmt.Println(v1, v39)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
