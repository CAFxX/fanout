package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v5 := []int{0, 4294967295}
    v6 := GenS1{F0: 42, F1: 42, F2: 2}
    if (false && false) {
        var v7 int = 32768
        if false {
            if ((0 >= 496) || true) {
                fmt.Println(v7)
                if (2 < 42) {
                    switch (v7 << ((42) & 63)) {
                    case 1:
                        fmt.Println(v7)
                    case -9223372036854775808:
                        _ = len(v5)
                    case 195:
                        for fi8 := 0; fi8 < 11; fi8++ {
                            switch -9223372036854775808 {
                            case 2147483647:
                                fmt.Println(v7, v7)
                            case -1:
                                v7 = (fi8 >> ((2) & 63))
                            case 32767:
                                switch "hello" {
                                case "\n\t":
                                    v7 += (fi8 & v7)
                                case "foo bar":
                                    var v9 int64 = 1
                                    _ = v9
                                }
                            }
                            v10 := GenS1{F0: 255, F1: 42, F2: 1}
                            _ = fi8
                            _ = v10
                        }
                    default:
                        v11 := GenS1{F0: 413, F1: -2147483648, F2: 0}
                        _ = v11
                    }
                }
                switch 32768 {
                case 9223372036854775807:
                    var v12 string = "hello"
                    _ = v12
                default:
                    v1["\n\t"] = "hello"
                }
            }
        } else {
            switch (v7 % ((v7) | 1)) {
            case 32767:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                fmt.Println(v7)
            }
            delete(v3, "\n\t")
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            switch (("a" + ("hello" + "")) + "\n\t") {
            case "a":
                var v14 uint32 = 738
                _ = v14
            case "foo bar":
                var v15 uint8 = 1
                _ = v15
            }
            var v16 uint16 = 617
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(v7)
                fmt.Println(v16)
                var v18 uint64 = 0
                v7 = fi17
                _ = fi17
                _ = v18
            }
            _ = fi13
            _ = v16
        }
        _ = v7
    }
    _ = len(v1)
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS2{F0: -34, F1: 2147483648, F2: fi19}
        _ = fi19
        _ = v20
    }
    if false {
        if true {
            v21 := GenS2{F0: 1, F1: 65535, F2: -2147483648}
            v22 := GenS1{F0: 985, F1: 0, F2: 65535}
            _ = v21
            _ = v22
        } else {
            var v23 bool = false
            v24 := GenS2{F0: 63, F1: 706, F2: -162}
            fmt.Println(v23)
            _ = v23
            _ = v24
        }
        v25 := GenS3{F0: 42, F1: 0}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 uint64 = 482
            var v28 uint32 = 0
            v27 += (v27 - 4294967295)
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 uint = 444
        _ = v29
    }
    var v30 uint8 = 61
    v31 := GenS4{F0: -609, F1: 255, F2: (v30 << ((42) & 63))}
    v30 = v30
    for fi32 := 0; fi32 < 1; fi32++ {
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 uint64 = 4294967295
            var v35 uint32 = 13
            v34 -= (v34 / ((18446744073709551615) | 1))
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    fmt.Println(v30)
    var v36 uint16 = 255
    var v37 int8 = -128
    var v38 int = -543
    v38 -= v38
    var v39 float64 = 1e10
    _ = len(v5)
    v40 := GenS2{F0: v37, F1: 2147483647, F2: v38}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
