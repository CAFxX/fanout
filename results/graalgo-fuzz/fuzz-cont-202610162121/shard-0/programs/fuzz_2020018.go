package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{1, 32768, 0, 821}
    var v3 int = 199
    var v4 uint16 = 666
    var v5 float32 = 1e10
    v6 := make(map[int]bool)
    if true {
        v4 += (v4 / ((v4) | 1))
        for fi7 := 0; fi7 < 5; fi7++ {
            v5 += 0.1
            v8 := GenS1{F0: v4, F1: -1}
            v9 := GenS2{F0: -32768}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    switch 32768 {
    case -815:
        v10 := GenS1{F0: (v4 * v4), F1: -277}
        _ = v10
    case 32767:
        v4 = v4
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int8 = -1
        _ = fi11
        _ = v12
    }
    switch "a" {
    case "foo bar":
        if true {
            _ = len(v1)
            fmt.Println(v5)
        } else {
            if true {
                if false {
                    var v13 float32 = v5
                    v14 := v6[(-(v3))]
                    _ = v13
                    _ = v14
                }
                v6[(v3 - -835)] = (2147483647 != v3)
                switch 0 {
                case 32768:
                    var v15 int8 = -1
                    _ = v15
                case 1:
                    var v16 int8 = 0
                    _ = v16
                default:
                    switch "\n\t" {
                    case "\n\t":
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v18 := GenS2{F0: -198}
                            for fi19 := 0; fi19 < 4; fi19++ {
                                v20 := GenS3{F0: (2147483648 != 915), F1: 2, F2: 1000}
                                _ = fi19
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 6; fi21++ {
                                var v22 uint8 = 127
                                _ = fi21
                                _ = v22
                            }
                            _ = fi17
                            _ = v18
                        }
                    case "":
                        var v23 float32 = (1e10 * v5)
                        _ = v23
                    }
                }
            }
        }
    default:
        for fi24 := 0; fi24 < 10; fi24++ {
            v3 = -2147483648
            switch "hello" {
            case "hello":
                switch v3 {
                case 127:
                    var v25 int64 = -1
                    _ = v25
                default:
                    fmt.Println(fi24)
                }
            case "\n\t":
                v26 := GenS2{F0: -15}
                _ = v26
            case "foo bar":
                v4 += 2
            }
            _ = fi24
        }
    }
    var v27 uint32 = 781
    v3 -= 1
    v5 = v5
    if true {
        v3 += (-(v3))
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        if (976 != 4294967295) {
            var v29 bool = false
            v30 := GenS1{F0: 32767, F1: 42}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v29
            _ = v30
        }
        var v31 uint64 = 250
        v1[v3] = v27
        _ = fi28
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
