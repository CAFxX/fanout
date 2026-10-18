package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint64; F1 string }
func genFn0(p0 int, p1 uint16) int {
    v6 := GenS1{F0: 255, F1: 100, F2: (p1 - 1)}
    v7 := []int{p0, (p0 >> ((18446744073709551615) & 63)), 32768, p0}
    p1 = (^(p1))
    var v8 int = (p0 & 128)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (v8 >> ((1) & 63))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 uint = 65535
    var v4 int = 42
    v5 := make(map[string]int)
    var v9 int32 = 0
    v10 := GenS1{F0: (^(v3)), F1: 0, F2: 694}
    var v11 uint = 1
    switch "foo bar" {
    case "a":
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS2{F0: 1, F1: 163}
            for fi14 := 0; fi14 < 4; fi14++ {
                v11 = v3
                if false {
                    v15 := v2["\n\t"]
                    fmt.Println(v4, v11)
                    _ = v15
                }
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
    case "hello":
        var v16 string = "a"
        _ = v16
    }
    if (false && (("hello" + ((("" + "foo bar") + "foo bar") + ("" + "a"))) >= "hello")) {
        switch v4 {
        case -9223372036854775808:
            var v17 int = -159
            _ = v17
        case 312:
            var v18 uint64 = 0
            _ = v18
        }
        v4 = (v4 << ((2) & 63))
        v3 -= 141
    } else {
        if (2 > v11) {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint8 = 32
                switch (fi19 << (((v3 % ((v3) | 1))) & 63)) {
                case 654:
                    delete(v5, "hello")
                case -9223372036854775808:
                    for fi21 := 0; fi21 < 2; fi21++ {
                        var v22 uint16 = 0
                        _ = fi21
                        _ = v22
                    }
                case 337:
                    v11 += v3
                default:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v3 -= (v3 & (v3 * (v3 / (((v11 ^ 42)) | 1))))
                        delete(v5, "a")
                        if (-1 < -9223372036854775808) {
                            var v24 int = v4
                            v3 += v11
                            v25 := v2["hello"]
                            v26 := GenS1{F0: v3, F1: 0, F2: 870}
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            v9 -= 0
                            switch "" {
                            case "\n\t":
                                delete(v5, "\n\t")
                            default:
                                var v27 bool = (((!true) || (!(true && false))) && true)
                                _ = v27
                            }
                        }
                        _ = fi23
                    }
                }
                _ = fi19
                _ = v20
            }
            v11 = v11
        }
        switch (v4 + v4) {
        case 1:
            v28 := GenS2{F0: 978, F1: 0}
            _ = v28
        default:
            switch -2147483648 {
            case 32768:
                v9 = (v9 >> ((812) & 63))
            case 2147483647:
                var v29 int8 = 42
                _ = v29
            }
        }
    }
    if false {
        v30 := GenS3{F0: 4294967295, F1: v9}
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        fmt.Println(v11, v9)
        var v32 bool = (42 >= v9)
        _ = fi31
        _ = v32
    }
    switch genFn0(128, 0) {
    case -1:
        if (-15 < 1) {
            if true {
                v11 = v3
                delete(v2, "foo bar")
            } else {
                if false {
                    var v33 uint64 = 141
                    _ = v33
                }
                fmt.Println(v4)
            }
            v4 -= 406
        }
    case 894:
        _ = len(v2)
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        switch fi34 {
        case 9223372036854775807:
            fmt.Println(v4, v3, v9)
        }
        _ = fi34
    }
    var v35 uint = v11
    v36 := GenS1{F0: 739, F1: 42, F2: 255}
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 float64 = (float64(((0.0 * 0.0) - -2.25)) + 1.5)
            if true {
                fmt.Println(v4, v39, fi37)
            }
            _ = fi38
            _ = v39
        }
        var v40 float32 = 0.1
        v4 = v4
        _ = fi37
        _ = v40
    }
    v41 := GenS4{F0: 4294967295, F1: ""}
    var v42 float32 = (1e10 + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v9
}
