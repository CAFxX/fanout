package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int8; F1 uint8; F2 int32 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint16, p1 int32, p2 string) int32 {
    p2 = p2
    fmt.Println(p0, p1, p2)
    if true {
        var v7 int8 = 127
        v7 += v7
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint64 = 4294967295
        if (!(-1 >= 32767)) {
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := []int{fi8, fi8}
                var v12 string = p2
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: (p1 | p1), F1: 1, F2: -1}
            v14 := GenS1{F0: p1, F1: -93, F2: int16(fi8)}
            _ = v13
            _ = v14
        } else {
            v15 := []int{fi8, fi8, (fi8 & (fi8 ^ fi8))}
            var v16 int8 = 121
            v17 := GenS2{F0: 4294967295, F1: 479}
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi8
        _ = v9
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        p2 = p2
        if false {
            v19 := GenS3{F0: (p0 * p0), F1: 62, F2: 255}
            _ = v19
        } else {
            v20 := []int{}
            for fi21 := 0; fi21 < 8; fi21++ {
                v20 = append(v20, (fi21 * (fi21 >> ((255) & 63))))
                var v22 int16 = 42
                var v23 int32 = (p1 >> ((255) & 63))
                _ = fi21
                _ = v22
                _ = v23
            }
            p2 = "foo bar"
            _ = v20
        }
        var v24 uint8 = 205
        _ = fi18
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 int16 = 963
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v4 bool = false
    v5 := []string{"\n\t", "a", "foo bar"}
    var v6 uint = 1
    var v25 uint32 = 2147483648
    v1 += (v1 * -32768)
    _ = len(v2)
    v26 := GenS4{F0: 1, F1: 2}
    fmt.Println(v4)
    var v27 int = 42
    var v28 int64 = 1
    switch ("\n\t" + "\n\t") {
    case "hello":
        v6 = v6
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint64 = 18446744073709551615
        var v31 uint8 = 171
        var v32 int64 = (^(v28))
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 int64 = v28
    v33 += (v28 >> (((^(v6))) & 63))
    if false {
        fmt.Println(v27)
        fmt.Println(v4, v1)
    } else {
        v4 = (float32(3.14159) < (1e10 - float32(0.0)))
    }
    v34 := GenS1{F0: 1, F1: 127, F2: 32767}
    var v35 uint8 = 127
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 int8 = 1
        for fi38 := 0; fi38 < 11; fi38++ {
            switch "\n\t" {
            case "foo bar":
                if v4 {
                    v28 += int64(v6)
                    switch -2147483648 {
                    case 127:
                        if false {
                            var v39 int32 = genFn0(573, 2147483647, "hello")
                            switch "a" {
                            case "a":
                                if (!v4) {
                                    v40 := GenS5{F0: "hello", F1: (v33 << ((0) & 63))}
                                    var v41 int = fi38
                                    _ = v40
                                    _ = v41
                                }
                            case "\n\t":
                                var v42 int8 = int8(v35)
                                var v43 uint8 = uint8(v25)
                                _ = v42
                                _ = v43
                            case "":
                                var v44 uint32 = 860
                                _ = v44
                            default:
                                var v45 string = ""
                                _ = v45
                            }
                            var v46 uint16 = 42
                            _ = v39
                            _ = v46
                        }
                    case -1:
                        v33 = (^(v33))
                    case 32767:
                        v47 := v2[(-(v27))]
                        _ = v47
                    default:
                        fmt.Println(v27, v25)
                    }
                    v37 -= v37
                    v48 := v3[(fi38 / ((fi38) | 1))]
                    _ = len(v5)
                    _ = v48
                }
            }
            _ = fi38
        }
        var v49 int64 = (v28 & (v33 & (v28 * (v33 | v33))))
        _ = fi36
        _ = v37
        _ = v49
    }
    for fi50 := 0; fi50 < 19; fi50++ {
        if true {
            for fi51 := 0; fi51 < 17; fi51++ {
                v6 += v6
                v52 := GenS6{F0: 0, F1: 128, F2: -2147483648}
                _ = fi51
                _ = v52
            }
            v53 := v3[42]
            var v54 int8 = 0
            _ = v53
            _ = v54
        } else {
            if v4 {
                fmt.Println(v4)
                var v55 uint16 = 907
                _ = v55
            }
            switch (fi50 << (((v6 + v6)) & 63)) {
            case 42:
                for fi56 := 0; fi56 < 12; fi56++ {
                    v57 := GenS7{F0: v33}
                    _ = fi56
                    _ = v57
                }
            }
            for fi58 := 0; fi58 < 8; fi58++ {
                var v59 uint64 = 698
                _ = fi58
                _ = v59
            }
        }
        v27 = 2147483647
        v27 += int(v35)
        _ = fi50
    }
    if (false && (100 <= -128)) {
        var v60 string = "\n\t"
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
