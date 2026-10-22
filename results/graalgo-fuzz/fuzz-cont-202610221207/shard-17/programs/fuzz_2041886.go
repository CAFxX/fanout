package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 uint; F2 string }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint64) int32 {
    var v5 int = -2147483648
    v6 := GenS1{F0: 2, F1: p0}
    switch "hello" {
    case "\n\t":
        if false {
            v7 := []int{230, (v5 << ((18446744073709551615) & 63))}
            if false {
                var v8 int32 = -1
                _ = len(v7)
                var v9 uint = 4294967295
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 14, F1: p0}
            _ = v10
            _ = v7
        } else {
            p0 = (p0 / ((p0) | 1))
            v11 := GenS1{F0: 127, F1: p0}
            if true {
                p0 = (p0 / ((953) | 1))
            } else {
                p0 += (p0 & 564)
            }
            _ = v11
        }
    case "foo bar":
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS2{F0: (fi12 & 821)}
            _ = fi12
            _ = v13
        }
    }
    if false {
        v14 := GenS3{F0: (p0 * 58), F1: 900}
        _ = v14
    } else {
        var v15 uint64 = 18446744073709551615
        v16 := GenS4{F0: 382}
        fmt.Println(v15, v15, v15)
        _ = v15
        _ = v16
    }
    p0 = (p0 & 1)
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{571, 65535, 216}
    var v2 uint32 = 951
    var v3 bool = (false || false)
    var v4 string = "a"
    if (847 <= genFn0(4294967295)) {
        fmt.Println(v2, v3)
        v2 = 2147483647
    }
    var v17 int = 85
    for fi18 := 0; fi18 < 5; fi18++ {
        if len(v1) > 0 {
            v1[0] = (v2 ^ 0)
        }
        _ = fi18
    }
    if true {
        v17 -= 4294967295
    }
    v3 = ((v2 | (v2 - (v2 & 1))) >= (^(v2)))
    var v19 uint16 = 255
    if ((v17 - (v17 << ((667) & 63))) >= 42) {
        switch "" {
        case "a":
            switch "\n\t" {
            case "a":
                if v3 {
                    switch "" {
                    case "":
                        v2 = 0
                    case "a":
                        if (v4 >= v4) {
                            if len(v1) > 0 {
                                v1[0] = (v2 % (((v2 * v2)) | 1))
                            }
                        } else {
                            v20 := GenS5{F0: 32767, F1: 65535, F2: "a"}
                            if (!((42 <= -32768) || (0 >= -1))) {
                                v21 := GenS6{F0: v2}
                                if v3 {
                                    var v22 uint = 42
                                    _ = v22
                                }
                                var v23 int = (v17 / ((32767) | 1))
                                _ = v21
                                _ = v23
                            }
                            if false {
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    var v25 int8 = 0
                                    var v26 int8 = v25
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            }
                            _ = v20
                        }
                    case "foo bar":
                        var v27 int = 127
                        _ = v27
                    }
                    var v28 uint8 = 255
                    _ = v28
                }
            }
        }
        var v29 string = v4
        v30 := GenS6{F0: v2}
        _ = v29
        _ = v30
    }
    v4 = "hello"
    for fi31 := 0; fi31 < 7; fi31++ {
        v2 -= (v2 - 4294967295)
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
