package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int8; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    var v5 uint64 = 462
    var v6 uint16 = 475
    v5 -= (v5 - (v5 + 2))
    if false {
        v7 := []int{-1, 32767, 9223372036854775807}
        if (!false) {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint32 = 42
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 8; fi10++ {
                switch "\n\t" {
                case "a":
                    v11 := GenS1{F0: 2}
                    _ = v11
                case "":
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v6 = v6
                        if (-1 == (v2 * v2)) {
                            if false {
                                fmt.Println(fi12, v2, fi12)
                                fmt.Println(v6, fi12)
                                fmt.Println(v5)
                            }
                            v13 := GenS2{F0: 1, F1: 32767, F2: ""}
                            v2 = v2
                            _ = v13
                        }
                        var v14 int = fi12
                        _ = fi12
                        _ = v14
                    }
                }
                _ = fi10
            }
            v15 := GenS2{F0: -2147483648, F1: 0, F2: ("\n\t" + "foo bar")}
            _ = v15
        } else {
            var v16 int16 = -1
            _ = v16
        }
        fmt.Println(v2)
        _ = v7
    } else {
        v17 := []int{4294967295, -1, 2147483647, -1}
        switch -2147483648 {
        case 128:
            var v18 int8 = -128
            _ = v18
        case 626:
            if ((v2 ^ v2) == (v2 >> ((0) & 63))) {
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 uint = 65535
                    _ = fi19
                    _ = v20
                }
                v6 = v6
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 bool = (float64(0.0) < float64((1e10 + 3.14159)))
                    _ = fi21
                    _ = v22
                }
            }
        case 1:
            if true {
                v6 -= (v6 % ((v6) | 1))
                if len(v17) > 0 {
                    v17[0] = 32767
                }
                var v23 string = ("\n\t" + "hello")
                _ = v23
            } else {
                v6 -= v6
                var v24 int64 = -21
                _ = v24
            }
        }
        fmt.Println(v2)
        _ = v17
    }
    if ((false || ((v6 % (((v6 % ((v6) | 1))) | 1)) != v6)) || true) {
        for fi25 := 0; fi25 < 3; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS3{F0: 255}
                _ = fi26
                _ = v27
            }
            if false {
                v2 = v2
            } else {
                v6 = (v6 % ((v6) | 1))
            }
            _ = fi25
        }
    }
    v2 += (-(v2))
    v2 = (v2 + -1)
    v28 := GenS4{F0: "foo bar", F1: v2}
    switch -2147483648 {
    case -2147483648:
        v29 := GenS5{F0: v2, F1: 1, F2: (v5 | v5)}
        _ = v29
    }
    v2 = 1
    if false {
        for fi30 := 0; fi30 < 2; fi30++ {
            _ = len(v3)
            fmt.Println(v5, v2)
            var v31 uint8 = 1
            _ = fi30
            _ = v31
        }
        if true {
            if ("" >= ("" + "a")) {
                var v32 string = ("\n\t" + ("a" + ("hello" + "")))
                var v33 float32 = 0.0
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 13; fi34++ {
                for fi35 := 0; fi35 < 2; fi35++ {
                    v6 += (v6 << ((34) & 63))
                    _ = fi35
                }
                fmt.Println(v6, v5)
                _ = fi34
            }
            if false {
                switch int(v5) {
                case -9223372036854775808:
                    switch "hello" {
                    case "\n\t":
                        v36 := []int{-971, -265, 0}
                        _ = v36
                    default:
                        for fi37 := 0; fi37 < 5; fi37++ {
                            v5 -= v5
                            v38 := GenS6{F0: 42, F1: -9223372036854775808}
                            var v39 uint16 = v6
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                case 0:
                    fmt.Println(v6)
                case 42:
                    v40 := GenS4{F0: "", F1: -1}
                    _ = v40
                }
                fmt.Println(v2, v2, v5)
            }
        }
        if true {
            v41 := []int{32768, 4294967295, 127, 2147483647}
            v42 := GenS4{F0: "", F1: v2}
            _ = v41
            _ = v42
        } else {
            v6 = v6
        }
    }
    fmt.Println(v5, v2, v2)
    var v43 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
