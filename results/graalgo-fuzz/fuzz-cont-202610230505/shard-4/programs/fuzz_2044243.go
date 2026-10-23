package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int64; F1 uint; F2 string }
type GenS6 struct { F0 uint; F1 int8; F2 bool }
func genFn0(p0 int64, p1 int64) string {
    var v6 string = "\n\t"
    var v7 string = ((v6 + (v6 + "\n\t")) + v6)
    v8 := GenS1{F0: 1}
    v6 = (("\n\t" + (v7 + "hello")) + v7)
    p0 += (p0 & (p1 * p0))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1() int16 {
    var v9 int8 = 0
    v9 = v9
    if true {
        switch -161 {
        case 9223372036854775807:
            fmt.Println(v9)
        case 2147483647:
            switch -1 {
            case 127:
                fmt.Println(v9)
            case 32768:
                v10 := []int{}
                _ = v10
            }
        }
    }
    _ = v9
    return 32767
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    v3 := []uint32{65535, 198, 1, 4294967295}
    var v4 string = "hello"
    v5 := []uint32{}
    delete(v2, 128)
    if (0.1 <= (1.5 + -2.25)) {
        v11 := GenS1{F0: 104}
        if (42 > 187) {
            v4 = ""
            _ = len(v2)
        } else {
            if (0 < 102) {
                for fi12 := 0; fi12 < 9; fi12++ {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        delete(v2, fi13)
                        v4 = "\n\t"
                        v14 := GenS1{F0: 42}
                        _ = fi13
                        _ = v14
                    }
                    v1 += v1
                    _ = fi12
                }
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS2{F0: (v1 << ((84) & 63))}
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                    _ = fi15
                    _ = v16
                }
            }
            var v17 float32 = 1e10
            v18 := GenS3{F0: 543, F1: 9223372036854775807}
            _ = v17
            _ = v18
        }
        v1 += (^(v1))
        _ = v11
    } else {
        var v19 uint64 = 42
        v20 := GenS4{F0: 625, F1: -2147483648}
        v4 = v4
        _ = v19
        _ = v20
    }
    if (v4 >= (((("" + ("hello" + "\n\t")) + v4) + v4) + v4)) {
        var v21 uint16 = 65535
        _ = v21
    } else {
        fmt.Println(v4, v4)
        v22 := GenS5{F0: -9223372036854775808, F1: 42, F2: "a"}
        switch 4294967295 {
        case -2147483648:
            for fi23 := 0; fi23 < 6; fi23++ {
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 uint64 = 0
                    if (325 != 42) {
                        var v26 string = "\n\t"
                        v27 := v2[(-(fi24))]
                        var v28 uint16 = 325
                        v28 += uint16(v25)
                        if (42 > 114) {
                            if (100 != 100) {
                                if true {
                                    var v29 uint = 1
                                    _ = v29
                                }
                                v25 = 1
                                fmt.Println(fi24, v26, fi23)
                                if v27 {
                                    v4 = ((v4 + v26) + "a")
                                    v26 = v4
                                } else {
                                    v30 := v2[fi24]
                                    var v31 bool = (((v4 + "") + v4) <= v26)
                                    _ = len(v2)
                                    _ = v30
                                    _ = v31
                                }
                            }
                        } else {
                            if ((0.1 * -2.25) < 1.5) {
                                delete(v2, (fi24 % ((fi24) | 1)))
                                var v32 uint16 = v28
                                var v33 string = "hello"
                                _ = v32
                                _ = v33
                            } else {
                                fmt.Println(v4, v1, fi23)
                                delete(v2, fi23)
                                var v34 int32 = v1
                                _ = v34
                            }
                        }
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                }
                var v35 int8 = 106
                _ = fi23
                _ = v35
            }
        }
        _ = v22
    }
    fmt.Println(v1)
    var v36 uint32 = 2147483647
    v4 = v4
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := v2[-1]
        for fi39 := 0; fi39 < 7; fi39++ {
            v36 = 4294967295
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    v40 := GenS6{F0: 2, F1: -1, F2: (true && (1 > 2))}
    switch v4 {
    case "a":
        _ = len(v3)
    case "\n\t":
        var v41 uint16 = 2
        _ = v41
    }
    v4 = v4
    v4 = "hello"
    var v42 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
