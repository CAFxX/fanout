package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 int64 }

func main() {
    var v1 string = "a"
    var v2 int64 = -9223372036854775808
    var v3 float32 = 3.14159
    var v4 string = v1
    var v5 int32 = -2147483648
    if true {
        v5 += 42
        var v6 int16 = 42
        _ = v6
    } else {
        v7 := []int{828, -443, 806}
        fmt.Println(v3, v3)
        fmt.Println(v5, v3)
        _ = v7
    }
    v4 = ((v4 + ("\n\t" + "\n\t")) + (("a" + ((v4 + (("hello" + ("" + v4)) + "")) + v4)) + "a"))
    v1 = v1
    switch (v1 + "\n\t") {
    case "foo bar":
        v1 = (v1 + (((v4 + v4) + v4) + v4))
    case "\n\t":
        if false {
            fmt.Println(v1, v3, v4)
            fmt.Println(v4, v2, v2)
            var v8 uint64 = 0
            _ = v8
        } else {
            v5 += v5
        }
    case "hello":
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int64 = v2
            v11 := GenS1{F0: 9223372036854775807, F1: 64}
            v12 := GenS1{F0: v10, F1: 167}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        switch ((v1 + (v4 + v4)) + v1) {
        case "":
            for fi13 := 0; fi13 < 19; fi13++ {
                switch fi13 {
                case 32767:
                    if false {
                        if false {
                            switch fi13 {
                            case 288:
                                v14 := []int{}
                                _ = v14
                            }
                            if false {
                                var v15 string = "hello"
                                _ = v15
                            }
                        }
                        v2 = (v2 / (((v2 >> ((264) & 63))) | 1))
                    } else {
                        var v16 int32 = (v5 >> ((262) & 63))
                        var v17 uint32 = 16
                        var v18 int32 = (v5 << ((0) & 63))
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                case -9223372036854775808:
                    var v19 int16 = 42
                    _ = v19
                }
                _ = fi13
            }
        default:
            if (true && (((false && true) || false) || false)) {
                if true {
                    var v20 int64 = (v2 ^ v2)
                    v21 := []int{}
                    switch (((v1 + v1) + ((((v4 + v1) + ((v1 + "\n\t") + v4)) + "a") + v1)) + (v4 + v4)) {
                    case "a":
                        v5 -= -47
                    }
                    _ = v20
                    _ = v21
                }
                v4 = v4
            } else {
                switch (("a" + ((v1 + "") + "\n\t")) + "\n\t") {
                case "":
                    v3 = (float32(-2.25) + v3)
                case "foo bar":
                    var v22 uint8 = 2
                    _ = v22
                case "\n\t":
                    if (v5 < v5) {
                        v23 := []int{229}
                        _ = v23
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if (-9223372036854775808 <= 42) {
                                for fi24 := 0; fi24 < 13; fi24++ {
                                    if true {
                                        if true {
                                            for fi25 := 0; fi25 < 9; fi25++ {
                                                v26 := []int{fi24, (-(fi24)), fi25}
                                                v27 := GenS2{F0: 4294967295, F1: 172}
                                                _ = fi25
                                                _ = v26
                                                _ = v27
                                            }
                                            v2 += 1
                                            v28 := []int{846, fi24}
                                            _ = v28
                                        } else {
                                            switch v1 {
                                            case "\n\t":
                                                for fi29 := 0; fi29 < 16; fi29++ {
                                                    if false {
                                                        var v30 bool = false
                                                        v31 := GenS2{F0: 65535, F1: 1}
                                                        v32 := GenS1{F0: 9223372036854775807, F1: 2147483648}
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    _ = fi29
                                                }
                                            case "a":
                                                v1 = v4
                                            case "foo bar":
                                                var v33 uint64 = 326
                                                _ = v33
                                            }
                                            var v34 uint8 = 1
                                            fmt.Println(v3, v4)
                                            _ = v34
                                        }
                                        v3 += (v3 * 1.5)
                                    }
                                    _ = fi24
                                }
                                v4 = (v1 + ("a" + v1))
                            }
                        }
                    }
                }
            }
        }
    }
    v35 := GenS3{F0: 18446744073709551615, F1: 2147483647, F2: 1}
    v36 := []int{}
    v5 = (v5 >> ((181) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
