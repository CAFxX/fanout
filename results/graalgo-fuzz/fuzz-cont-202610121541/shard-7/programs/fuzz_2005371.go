package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 string) bool {
    if true {
        p0 = (p0 + p0)
        p0 = p0
        if true {
            v7 := []int{}
            fmt.Println(p0)
            if (!((!(false && (-9223372036854775808 <= 128))) || (2 < 677))) {
                var v8 int64 = 122
                _ = v8
            } else {
                var v9 uint16 = 294
                switch 128 {
                case -9223372036854775808:
                    var v10 int16 = -16
                    _ = v10
                case 32767:
                    for fi11 := 0; fi11 < 20; fi11++ {
                        fmt.Println(v9)
                        _ = fi11
                    }
                case 127:
                    var v12 float32 = 1e10
                    _ = v12
                default:
                    if (1 >= 697) {
                        if (464 >= 0) {
                            fmt.Println(v9, v9)
                            var v13 uint32 = 65535
                            var v14 uint16 = v9
                            _ = v13
                            _ = v14
                        }
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                }
                _ = v9
            }
            _ = v7
        }
    } else {
        v15 := []int{}
        fmt.Println(p0)
        fmt.Println(p0)
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: 32767}
    v17 := GenS2{F0: ("\n\t" + p0)}
    var v18 uint32 = 42
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return true
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 bool = true
    var v5 int16 = -245
    var v6 int32 = 2147483647
    v5 = (v5 | (v5 + (v5 << ((272) & 63))))
    fmt.Println(v5, v4)
    v19 := GenS1{F0: 0, F1: 32767}
    v20 := GenS2{F0: "a"}
    switch "hello" {
    case "a":
        var v21 uint = 585
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 string = ("hello" + "hello")
        switch v23 {
        case "\n\t":
            if (true && (36 < 0)) {
                for fi24 := 0; fi24 < 18; fi24++ {
                    v4 = (v4 || false)
                    _ = fi24
                }
            }
        }
        _ = fi22
        _ = v23
    }
    var v25 uint = 42
    v26 := GenS2{F0: ("hello" + "")}
    v4 = (4294967295 != 18446744073709551615)
    var v27 uint16 = 32767
    var v28 uint64 = 18446744073709551615
    if (true && (v4 || ((((v28 ^ (v28 / ((v28) | 1))) < (v28 << (((v25 & (v25 / ((851) | 1)))) & 63))) && (uint(v28) != (^(v25)))) || false))) {
        var v29 int32 = v6
        _ = v29
    } else {
        switch -2147483648 {
        case -1:
            v27 = (v27 << ((v25) & 63))
        case 66:
            if (128 <= 1) {
                var v30 uint32 = 1
                for fi31 := 0; fi31 < 6; fi31++ {
                    if ((v6 - 0) == (v6 / (((v6 * v6)) | 1))) {
                        if false {
                            var v32 int8 = (v1 / (((v1 & int8(fi31))) | 1))
                            var v33 float32 = 1e10
                            _ = v32
                            _ = v33
                        } else {
                            v34 := GenS1{F0: 9223372036854775807, F1: v27}
                            _ = v34
                        }
                    }
                    fmt.Println(v30, v25)
                    _ = fi31
                }
                v1 = (v1 + int8(v5))
                _ = v30
            }
        }
        v35 := GenS3{F0: ("hello" + ((("" + "") + "hello") + "a"))}
        _ = v35
    }
    v28 = v28
    v36 := GenS4{F0: 65535, F1: -1}
    fmt.Println(v27)
    switch 9223372036854775807 {
    case 42:
        for fi37 := 0; fi37 < 8; fi37++ {
            v28 -= v28
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 string = "hello"
                _ = fi38
                _ = v39
            }
            var v40 float64 = 0.1
            _ = fi37
            _ = v40
        }
    case 32767:
        var v41 float32 = (1.5 + float32(1e10))
        _ = v41
    }
    v42 := GenS4{F0: 2147483647, F1: 0}
    v43 := GenS2{F0: ("" + ("hello" + "foo bar"))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
