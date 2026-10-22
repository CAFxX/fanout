package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64) uint16 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{fi7}
        if (p0 <= (p0 | (p0 | p0))) {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if true {
                v9 := GenS1{F0: "foo bar", F1: 0}
                v10 := GenS1{F0: "foo bar", F1: 15}
                v11 := GenS1{F0: "foo bar", F1: 2147483647}
                _ = v10
                _ = v11
                _ = v9
            }
        }
        p0 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    v12 := []int{}
    switch ("hello" + ("hello" + "foo bar")) {
    case "foo bar":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    case "a":
        p0 = (^(p0))
    case "hello":
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(fi13, p0)
            _ = fi13
        }
    default:
        var v14 int16 = 1000
        _ = v14
    }
    v12 = append(v12, -1)
    _ = p0
    _ = v12
    return 1
}

func main() {
    v1 := []uint32{176, 42, 1, 2147483647}
    var v2 uint32 = 2147483648
    var v3 int16 = -1
    var v4 int32 = 2147483647
    var v5 int64 = -9223372036854775808
    var v6 bool = (!true)
    var v15 bool = v6
    if (!false) {
        var v16 int = -2147483648
        _ = v16
    }
    var v17 uint8 = 42
    v15 = false
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint = 669
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: "foo bar", F1: -427}
    switch "\n\t" {
    case "foo bar":
        v21 := GenS1{F0: "hello", F1: (v4 / (((v4 | v4)) | 1))}
        _ = v21
    case "hello":
        if (!(-119 >= 127)) {
            if v6 {
                if ((42 > 32767) && v6) {
                    if len(v1) > 0 {
                        v1[0] = v2
                    }
                    var v22 uint = 4294967295
                    _ = v22
                }
            } else {
                var v23 int8 = 127
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                for fi25 := 0; fi25 < 4; fi25++ {
                    v1 = append(v1, (^(v2)))
                    switch fi24 {
                    case -9223372036854775808:
                        switch "" {
                        case "foo bar":
                            fmt.Println(v15, v3, v17)
                        default:
                            var v26 int32 = v4
                            _ = v26
                        }
                    default:
                        var v27 int64 = v5
                        _ = v27
                    }
                    _ = fi25
                }
                v28 := GenS2{F0: (v5 >> ((1) & 63))}
                _ = fi24
                _ = v28
            }
            var v29 uint32 = v2
            v17 += (v17 | v17)
            _ = v29
        }
    default:
        var v30 int8 = 127
        _ = v30
    }
    if v15 {
        var v31 int64 = (v5 << ((515) & 63))
        fmt.Println(v5, v6, v15)
        _ = v31
    } else {
        fmt.Println(v5)
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 int = 32768
            for fi34 := 0; fi34 < 11; fi34++ {
                v35 := GenS2{F0: -9223372036854775808}
                v36 := GenS2{F0: v5}
                var v37 uint64 = 0
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            _ = fi32
            _ = v33
        }
    }
    if v6 {
        fmt.Println(v15, v2, v6)
        fmt.Println(v15, v6)
        v38 := GenS1{F0: "foo bar", F1: v4}
        _ = v38
    } else {
        v2 += v2
    }
    var v39 float64 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v40 uint = 168
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
