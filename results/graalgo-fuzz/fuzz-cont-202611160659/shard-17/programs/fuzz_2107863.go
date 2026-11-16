package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint8 }
func genFn0(p0 string, p1 int8, p2 bool) int64 {
    fmt.Println(p0, p1)
    if p2 {
        fmt.Println(p2, p1, p1)
        if ((p2 || ((-2.25 == 0.1) || ((true || p2) || p2))) || (!true)) {
            fmt.Println(p1, p0, p0)
            var v6 uint = 18446744073709551615
            var v7 uint16 = uint16(v6)
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: 1}
            var v9 bool = p2
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 647
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "hello"
    var v3 uint8 = 1
    v4 := []uint32{993, 175, 2147483647}
    var v5 int = 0
    if false {
        v10 := GenS1{F0: -939}
        if true {
            var v11 float32 = -2.25
            switch v5 {
            case 42:
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS2{F0: -107, F1: 1, F2: (v3 << ((1) & 63))}
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    v3 = v3
                    _ = fi12
                    _ = v13
                }
            }
            _ = v11
        } else {
            fmt.Println(v3, v5)
        }
        _ = v10
    }
    var v14 float32 = -2.25
    for fi15 := 0; fi15 < 11; fi15++ {
        if true {
            fmt.Println(v2, v14)
            switch v2 {
            case "a":
                for fi16 := 0; fi16 < 2; fi16++ {
                    switch v2 {
                    case "a":
                        switch v2 {
                        case "hello":
                            if false {
                                var v17 int16 = 0
                                fmt.Println(fi16)
                                if true {
                                    v4 = append(v4, 4294967295)
                                }
                                _ = v17
                            }
                        case "foo bar":
                            if (true && false) {
                                v14 += (v14 + v14)
                                v5 = fi16
                                v18 := v1[2147483648]
                                _ = v18
                            } else {
                                v1[0] = (fi15 << ((767) & 63))
                                fmt.Println(v3, fi16, v2)
                                switch fi16 {
                                case 804:
                                    var v19 uint32 = 2147483647
                                    _ = v19
                                }
                                v4 = append(v4, 2147483648)
                                v3 += (v3 | (^(v3)))
                            }
                        case "":
                            v5 += fi15
                        }
                    case "foo bar":
                        v20 := GenS2{F0: 127, F1: 42, F2: v3}
                        _ = v20
                    case "":
                        v21 := GenS2{F0: 42, F1: 991, F2: 144}
                        _ = v21
                    default:
                        var v22 uint8 = 2
                        _ = v22
                    }
                    _ = fi16
                }
            case "foo bar":
                v5 += (fi15 << ((413) & 63))
            }
            var v23 int8 = 1
            _ = v23
        } else {
            if (42 != 1000) {
                for fi24 := 0; fi24 < 5; fi24++ {
                    if ((v3 >> ((4294967295) & 63)) > v3) {
                        _ = len(v4)
                        fmt.Println(fi24, v3, v3)
                        v3 -= 111
                    }
                    var v25 uint8 = 0
                    v25 = v3
                    _ = fi24
                    _ = v25
                }
            }
        }
        _ = fi15
    }
    v5 = 127
    v1[2] = v5
    v2 = (v2 + v2)
    var v26 float32 = (v14 + v14)
    v27 := GenS2{F0: 110, F1: 0, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
