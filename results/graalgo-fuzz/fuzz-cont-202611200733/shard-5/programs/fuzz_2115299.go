package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 32767
    var v3 int64 = -977
    var v4 int64 = (v3 << ((4294967295) & 63))
    v5 := make(map[uint32]bool)
    var v6 uint32 = 802
    var v7 float32 = float32(1.5)
    v2 = v2
    var v8 int64 = v3
    v9 := GenS1{F0: (v6 << ((366) & 63)), F1: v2, F2: 16}
    v8 -= (v4 % (((-(v3))) | 1))
    fmt.Println(v4, v3, v2)
    if ((("" + "") + "") >= "a") {
        var v10 int16 = 1000
        switch -641 {
        case -1:
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 uint8 = 128
                switch ("foo bar" + ("\n\t" + (("foo bar" + "a") + ((("" + ("" + "foo bar")) + "hello") + "\n\t")))) {
                case "foo bar":
                    for fi13 := 0; fi13 < 16; fi13++ {
                        if true {
                            v14 := GenS2{F0: -1, F1: "hello"}
                            delete(v5, v6)
                            _ = v14
                        }
                        v15 := []int{-930, (-(fi11))}
                        v7 = ((-2.25 - v7) / v7)
                        if true {
                            fmt.Println(v12, v2, v7)
                            var v16 int32 = -1
                            if true {
                                switch fi11 {
                                case 32767:
                                    fmt.Println(fi11)
                                default:
                                    v15 = append(v15, fi11)
                                    v4 = v3
                                }
                            }
                            var v17 int = fi11
                            _ = v16
                            _ = v17
                        } else {
                            var v18 int8 = int8(fi11)
                            v3 = v3
                            _ = len(v15)
                            _ = v18
                        }
                        v6 = 65535
                        switch "" {
                        case "":
                            fmt.Println(v8)
                        case "hello":
                            var v19 int16 = (v10 * v10)
                            _ = v19
                        case "foo bar":
                            if true {
                                var v20 int16 = 619
                                if len(v15) > 0 {
                                    v15[0] = fi11
                                }
                                _ = v20
                            }
                        }
                        v10 = 32767
                        _ = fi13
                        _ = v15
                    }
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        if (!true) {
                            var v21 int32 = -1
                            for fi22 := 0; fi22 < 9; fi22++ {
                                v2 -= v2
                                var v23 int32 = v21
                                _ = fi22
                                _ = v23
                            }
                            _ = v21
                        }
                    default:
                        var v24 float32 = (1.5 * float32((3.14159 + (float64(0.1) - 3.14159))))
                        _ = v24
                    }
                }
                for fi25 := 0; fi25 < 13; fi25++ {
                    if true {
                        v2 -= (v2 * v2)
                        var v26 uint64 = 4294967295
                        _ = v26
                    }
                    var v27 int16 = v10
                    delete(v5, 1)
                    fmt.Println(fi25)
                    _ = fi25
                    _ = v27
                }
                _ = fi11
                _ = v12
            }
        }
        _ = v10
    } else {
        v28 := []int{4294967295, -9223372036854775808}
        _ = v28
    }
    var v29 bool = true
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            switch fi30 {
            case 127:
                var v32 float32 = (((v7 * v7) / (v7 + 0.1)) * v7)
                _ = v32
            case 32767:
                switch fi31 {
                case 32767:
                    var v33 bool = (v29 || (v6 < (v6 ^ (v6 & v6))))
                    _ = v33
                }
            default:
                v34 := GenS2{F0: 100, F1: "a"}
                _ = v34
            }
            v7 += ((v7 / v7) - float32((0.0 * 1.5)))
            if (!(!v29)) {
                v35 := GenS2{F0: 42, F1: ""}
                _ = v35
            } else {
                delete(v5, v6)
            }
            _ = fi31
        }
        var v36 int32 = 408
        _ = fi30
        _ = v36
    }
    switch "foo bar" {
    case "\n\t":
        v3 = int64(v6)
    case "a":
        v37 := GenS3{F0: v8}
        _ = v37
    case "":
        var v38 int8 = -128
        _ = v38
    }
    v39 := GenS3{F0: (v4 % (((v8 ^ -955)) | 1))}
    v40 := GenS4{F0: 1, F1: -61, F2: 65535}
    if ("" < "\n\t") {
        fmt.Println(v3, v7, v3)
        var v41 uint = 18446744073709551615
        _ = v41
    } else {
        v2 = (v2 >> ((466) & 63))
        for fi42 := 0; fi42 < 13; fi42++ {
            v43 := []int{}
            v5[v6] = v29
            switch ("a" + "hello") {
            case "":
                fmt.Println(v8, v2)
            }
            _ = fi42
            _ = v43
        }
    }
    v4 += v4
    v2 -= (v2 + uint16(v6))
    if ((v29 && (42 <= -2147483648)) || (42 != 131)) {
        var v44 float64 = 0.1
        _ = v44
    } else {
        for fi45 := 0; fi45 < 20; fi45++ {
            if (!((v2 > (v2 >> ((2) & 63))) && v29)) {
                v46 := GenS1{F0: v6, F1: 681, F2: 42}
                v6 = (^(v6))
                fmt.Println(v29, v2)
                _ = v46
            }
            fmt.Println(v6)
            _ = fi45
        }
    }
    var v47 uint = 1
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
