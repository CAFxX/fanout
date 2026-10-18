package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint16, p2 int16) uint16 {
    var v7 uint = 745
    fmt.Println(p2)
    if ((p2 % (((p2 % ((p2) | 1))) | 1)) >= (p2 << (((v7 / ((945) | 1))) & 63))) {
        var v8 int32 = -2147483648
        v9 := []int{32768, 2147483647}
        v10 := GenS1{F0: (v7 + (^(v7))), F1: 1, F2: 9223372036854775807}
        _ = v10
        _ = v8
        _ = v9
    }
    p1 += p0
    fmt.Println(p2, v7, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    v1 := []int{}
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 string = ""
    var v5 float64 = -2.25
    var v6 bool = (true || ((("" + v4) + "\n\t") > "foo bar"))
    var v11 uint = 4294967295
    var v12 uint = 746
    var v13 uint64 = 4294967295
    v13 = (v13 % ((v13) | 1))
    v14 := GenS1{F0: 1, F1: int16(v13), F2: 1}
    v2[-444] = 2147483647
    delete(v2, 862)
    v15 := GenS1{F0: v12, F1: 835, F2: -9223372036854775808}
    var v16 string = v4
    var v17 uint = uint(v13)
    v3[((v4 + "hello") + (v4 + v16))] = ((128 < 25) || v6)
    if false {
        delete(v3, (((v16 + "foo bar") + (("a" + v4) + "foo bar")) + v4))
        var v18 float64 = 1e10
        _ = v18
    }
    switch 42 {
    case 147:
        var v19 float32 = 0.0
        _ = v19
    case 4294967295:
        if (42 < 42) {
            v20 := GenS2{F0: 0, F1: v6}
            _ = v20
        } else {
            if v6 {
                v21 := GenS1{F0: (^(v12)), F1: 120, F2: -1}
                v22 := GenS3{F0: 109}
                if len(v1) > 0 {
                    v1[0] = -1
                }
                _ = v21
                _ = v22
            } else {
                var v23 int64 = -1
                _ = v23
            }
            delete(v2, -335)
        }
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        switch v4 {
        case "a":
            v1 = append(v1, -1)
        case "\n\t":
            switch ((v4 + v16) + (("\n\t" + v4) + v4)) {
            case "hello":
                var v25 uint16 = 0
                _ = v25
            }
        case "foo bar":
            v26 := GenS1{F0: (^(v12)), F1: 1000, F2: 42}
            _ = v26
        }
        v13 = 18446744073709551615
        if v6 {
            v6 = true
            for fi27 := 0; fi27 < 13; fi27++ {
                if (-384 < 999) {
                    var v28 int8 = 42
                    switch v16 {
                    case "a":
                        v12 -= (v12 / (((v17 * (v11 - (^(v12))))) | 1))
                    case "\n\t":
                        var v29 uint16 = 0
                        _ = v29
                    case "foo bar":
                        if v6 {
                            if (!v6) {
                                if len(v1) > 0 {
                                    v1[0] = fi27
                                }
                                var v30 int = fi24
                                delete(v3, (v4 + v16))
                                _ = v30
                            }
                            if false {
                                switch fi27 {
                                case -505:
                                    if v6 {
                                        var v31 string = v4
                                        if true {
                                            if ((42 >= 9223372036854775807) && (0 >= 163)) {
                                                v17 = (v12 - v11)
                                                var v32 int16 = 1
                                                _ = v32
                                            }
                                            fmt.Println(v12)
                                            v17 += v11
                                        }
                                        var v33 uint64 = (v13 - v13)
                                        _ = v31
                                        _ = v33
                                    }
                                }
                            } else {
                                v13 -= 42
                                fmt.Println(v17, v5, v16)
                                var v34 int8 = v28
                                _ = v34
                            }
                        }
                    }
                    v3[(v4 + ("\n\t" + (v4 + v4)))] = (false || ((^(fi24)) == fi24))
                    _ = v28
                } else {
                    v13 = v13
                    v6 = v6
                }
                v5 = 0.1
                _ = fi27
            }
            var v35 string = v4
            _ = v35
        }
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
