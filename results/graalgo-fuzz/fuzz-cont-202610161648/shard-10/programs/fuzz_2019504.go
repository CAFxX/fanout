package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = (v1 | v1)
    var v4 int8 = -23
    v5 := make(map[int]bool)
    var v6 string = ""
    var v7 float32 = 1.5
    v8 := []int{}
    for fi9 := 0; fi9 < 14; fi9++ {
        if true {
            fmt.Println(v4, fi9, v4)
            var v10 uint32 = 65535
            _ = v10
        } else {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint32 = 958
                var v13 uint64 = v2
                _ = fi11
                _ = v12
                _ = v13
            }
            v2 -= (v2 ^ (^(v2)))
            var v14 uint16 = v3
            _ = v14
        }
        if true {
            if (v6 < v6) {
                var v15 int = fi9
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 uint = 112
                    v18 := GenS1{F0: 2147483647}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if false {
                    if false {
                        if true {
                            v4 -= (v4 / ((42) | 1))
                            var v19 bool = false
                            var v20 uint32 = 399
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v7, v3, fi9)
                            var v21 int = (v15 << ((18446744073709551615) & 63))
                            _ = v21
                        }
                        v22 := GenS1{F0: 0}
                        if (302 < 255) {
                            var v23 float64 = float64(float32(1e10))
                            _ = v23
                        } else {
                            var v24 float32 = v7
                            var v25 uint32 = 2
                            if (9223372036854775807 > 32768) {
                                v4 = 127
                            }
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    }
                } else {
                    var v26 int = (fi9 % ((int(v4)) | 1))
                    _ = v26
                }
                _ = v15
            }
        } else {
            v27 := v5[fi9]
            v28 := GenS1{F0: 42}
            fmt.Println(v27, v4, fi9)
            _ = v27
            _ = v28
        }
        var v29 int16 = 32767
        _ = fi9
        _ = v29
    }
    fmt.Println(v2, v4, v6)
    v1 = v3
    v30 := v5[32768]
    if true {
        var v31 float64 = (0.0 + 0.0)
        fmt.Println(v2)
        v32 := GenS1{F0: 645}
        _ = v31
        _ = v32
    }
    var v33 uint32 = 4294967295
    var v34 float64 = -2.25
    v6 = (("foo bar" + "\n\t") + (v6 + ""))
    switch v6 {
    case "a":
        _ = len(v5)
    case "":
        switch 9223372036854775807 {
        case 128:
            v6 = "a"
        case -17:
            if v30 {
                if (981 <= 0) {
                    fmt.Println(v3)
                } else {
                    if v30 {
                        v35 := GenS2{F0: true}
                        _ = v35
                    } else {
                        if true {
                            var v36 uint8 = 89
                            v37 := v5[-490]
                            _ = v36
                            _ = v37
                        } else {
                            v5[42] = (v30 && (true || (true && v30)))
                        }
                        var v38 float64 = -2.25
                        _ = v38
                    }
                    if v30 {
                        if false {
                            fmt.Println(v2, v2)
                        } else {
                            var v39 uint8 = 128
                            for fi40 := 0; fi40 < 16; fi40++ {
                                v41 := GenS1{F0: v33}
                                _ = fi40
                                _ = v41
                            }
                            var v42 uint8 = (v39 - uint8(v33))
                            _ = v39
                            _ = v42
                        }
                        v30 = (v30 && false)
                    } else {
                        v2 += v2
                        v34 = v34
                    }
                    for fi43 := 0; fi43 < 16; fi43++ {
                        var v44 uint8 = 155
                        switch fi43 {
                        case 0:
                            fmt.Println(v1, v3)
                        case 2147483647:
                            v45 := GenS1{F0: 2}
                            _ = v45
                        case -1:
                            var v46 bool = v30
                            _ = v46
                        }
                        v7 += 1.5
                        _ = fi43
                        _ = v44
                    }
                }
                var v47 int32 = 941
                _ = v47
            } else {
                if (!(42 != 9223372036854775807)) {
                    var v48 uint = 42
                    _ = v48
                }
            }
        case -1:
            v49 := GenS1{F0: (v33 * (v33 / (((v33 ^ 647)) | 1)))}
            _ = v49
        }
    case "\n\t":
        v6 = v6
    default:
        v30 = v30
    }
    v34 = 3.14159
    var v50 int8 = (v4 + v4)
    v3 = (v1 << ((787) & 63))
    var v51 string = "a"
    if v30 {
        fmt.Println(v2, v51)
        var v52 int8 = -105
        _ = v52
    }
    v53 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v6
    _ = v7
    _ = v8
}
