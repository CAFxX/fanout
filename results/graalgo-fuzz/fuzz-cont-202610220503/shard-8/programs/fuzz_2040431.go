package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint; F1 int16; F2 int32 }
type GenS5 struct { F0 int64; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]bool)
    var v4 uint = 655
    v5 := make(map[string]int)
    var v6 int = int(v2)
    v4 -= (v4 * v4)
    v4 = (v4 / ((1) | 1))
    v7 := GenS1{F0: 0}
    v4 -= (^(v4))
    var v8 int64 = (v2 - 244)
    var v9 int16 = -721
    var v10 uint64 = 1
    switch v6 {
    case 32768:
        var v11 int16 = (v9 | -589)
        _ = v11
    case 174:
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int32 = -462
                if true {
                    v15 := GenS2{F0: "\n\t", F1: (v10 % ((v10) | 1))}
                    _ = v15
                } else {
                    v16 := GenS1{F0: v14}
                    var v17 bool = true
                    _ = v16
                    _ = v17
                }
                var v18 float32 = (1.5 * 1.5)
                _ = fi13
                _ = v14
                _ = v18
            }
            switch (v6 ^ 32767) {
            case 9223372036854775807:
                v10 = (v10 + (v10 ^ 4294967295))
            case -2147483648:
                switch (fi12 - (v6 | fi12)) {
                case 127:
                    v4 = v4
                }
            default:
                var v19 uint16 = 32767
                _ = v19
            }
            fmt.Println(v6)
            _ = fi12
        }
    case -1:
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS2{F0: "hello", F1: v10}
            v22 := GenS3{F0: "\n\t", F1: -1}
            var v23 int8 = -43
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    default:
        fmt.Println(v10, v4, v6)
    }
    var v24 uint64 = v10
    for fi25 := 0; fi25 < 13; fi25++ {
        v2 = (v2 & v2)
        switch 114 {
        case 32768:
            v26 := GenS3{F0: ("foo bar" + ("\n\t" + ("a" + "hello"))), F1: (v6 % (((v6 | (v6 - (fi25 % ((v6) | 1))))) | 1))}
            _ = v26
        case 1:
            if true {
                v27 := GenS4{F0: (v4 % (((v4 % (((v4 ^ 1)) | 1))) | 1)), F1: v9, F2: 42}
                _ = v27
            } else {
                v4 = v4
                switch "a" {
                case "a":
                    for fi28 := 0; fi28 < 19; fi28++ {
                        var v29 uint32 = 42
                        if (82 <= 42) {
                            v10 += 18446744073709551615
                        } else {
                            _ = len(v3)
                            v30 := GenS5{F0: (v8 % ((v8) | 1)), F1: 243, F2: -128}
                            _ = v30
                        }
                        _ = fi28
                        _ = v29
                    }
                }
                v31 := []int{(fi25 ^ -524)}
                _ = v31
            }
        case 128:
            for fi32 := 0; fi32 < 17; fi32++ {
                v33 := GenS6{F0: v4, F1: (^(v10))}
                if (!false) {
                    var v34 int8 = 100
                    if false {
                        v4 = (^(v4))
                        v2 += 9223372036854775807
                    } else {
                        if false {
                            v35 := GenS6{F0: v4, F1: (v24 >> (((v4 >> ((v4) & 63))) & 63))}
                            _ = len(v1)
                            _ = v35
                        }
                        if (4294967295 < 0) {
                            fmt.Println(v34, v2, v24)
                            delete(v3, "\n\t")
                        }
                        switch "\n\t" {
                        case "hello":
                            v36 := GenS4{F0: v4, F1: v9, F2: -1}
                            _ = v36
                        case "":
                            var v37 int16 = (v9 << ((v4) & 63))
                            _ = v37
                        }
                    }
                    _ = v34
                }
                var v38 string = ((("hello" + "\n\t") + (("a" + "") + "foo bar")) + "a")
                _ = fi32
                _ = v33
                _ = v38
            }
        default:
            var v39 float32 = 0.0
            _ = v39
        }
        _ = fi25
    }
    v10 = 988
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
