package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int }
func genFn0() int64 {
    v4 := GenS1{F0: 255, F1: 2}
    v5 := []int{32767}
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    v3 := []int{9223372036854775807, 571, 32768}
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint8 = 1
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: -128, F1: 1}
    switch 676 {
    case 2147483647:
        if (0 != 65535) {
            var v9 string = ""
            v3 = append(v3, 4294967295)
            var v10 uint = 2
            _ = v10
            _ = v9
        } else {
            switch 32768 {
            case 4294967295:
                var v11 uint = 885
                _ = v11
            case -9223372036854775808:
                switch 32767 {
                case 32767:
                    var v12 string = "\n\t"
                    _ = v12
                default:
                    var v13 uint8 = 0
                    _ = v13
                }
            case 1:
                var v14 string = ("" + "foo bar")
                _ = v14
            }
            if (!((float32(((0.1 * (1.5 - (-2.25 + -2.25))) * 1e10)) - 1e10) != 0.0)) {
                var v15 uint32 = 4294967295
                for fi16 := 0; fi16 < 8; fi16++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 bool = false
                        var v19 int64 = 1
                        v20 := GenS2{F0: -1, F1: 32768}
                        v21 := GenS1{F0: 0, F1: (v15 + (v15 * v15))}
                        switch (("hello" + "hello") + "") {
                        case "\n\t":
                            v22 := v1["\n\t"]
                            _ = v22
                        default:
                            var v23 int16 = -32768
                            _ = v23
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    if false {
                        v15 = (v15 * v15)
                        if false {
                            v15 = v15
                        }
                        v15 = v15
                    }
                    _ = fi16
                }
                _ = v15
            }
        }
    case 32768:
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS1{F0: 201, F1: 65535}
            var v26 string = "\n\t"
            v27 := GenS3{F0: 42}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    case -690:
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 uint64 = 18446744073709551615
            var v30 int16 = 1
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    switch "" {
    case "\n\t":
        var v31 bool = (((0.1 != 1e10) && false) && (true && (false || (42 >= 541))))
        _ = v31
    case "a":
        var v32 uint64 = 1
        _ = v32
    default:
        switch "foo bar" {
        case "\n\t":
            if true {
                var v33 uint32 = 65535
                _ = len(v2)
                v34 := GenS2{F0: 127, F1: 282}
                _ = v33
                _ = v34
            }
        default:
            var v35 int16 = -32768
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        var v37 int = (fi36 >> ((42) & 63))
        if ((0.0 < float64((float32((1.5 + 3.14159)) / 3.14159))) || false) {
            fmt.Println(v37, v37)
            fmt.Println(fi36, fi36)
        } else {
            var v38 int16 = 0
            var v39 bool = true
            var v40 bool = (210 != 700)
            _ = v38
            _ = v39
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    var v41 uint16 = 65535
    for fi42 := 0; fi42 < 3; fi42++ {
        switch ("" + "foo bar") {
        case "hello":
            v41 = v41
        case "":
            v43 := GenS3{F0: (v41 + v41)}
            _ = v43
        default:
            var v44 int8 = 0
            _ = v44
        }
        if (0 < 42) {
            fmt.Println(fi42, fi42)
            v41 += 0
        } else {
            var v45 uint32 = 0
            _ = v45
        }
        var v46 uint8 = 22
        _ = fi42
        _ = v46
    }
    if (!(0 < 667)) {
        fmt.Println(v41)
        _ = len(v3)
    }
    v47 := GenS4{F0: 0, F1: 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v47
    _ = v8
}
