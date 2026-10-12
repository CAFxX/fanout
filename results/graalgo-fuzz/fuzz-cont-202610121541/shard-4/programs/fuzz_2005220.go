package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 string; F1 int8; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 1
    var v3 uint16 = 1000
    var v4 int8 = 42
    v5 := make(map[uint32]string)
    var v6 int32 = -2147483648
    v7 := GenS1{F0: v6, F1: 42}
    var v8 int64 = -1
    v9 := GenS2{F0: "a", F1: -1, F2: (!("a" > ""))}
    v10 := []int{254}
    v6 = int32(v3)
    var v11 bool = true
    if ((^(v8)) >= (^(v8))) {
        var v12 int8 = 100
        if (!v11) {
            for fi13 := 0; fi13 < 14; fi13++ {
                v4 = (v12 + (v12 * v4))
                _ = fi13
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                switch "hello" {
                case "\n\t":
                    for fi15 := 0; fi15 < 6; fi15++ {
                        delete(v5, 2147483647)
                        v11 = v11
                        _ = fi15
                    }
                case "":
                    for fi16 := 0; fi16 < 11; fi16++ {
                        var v17 int8 = v12
                        _ = fi16
                        _ = v17
                    }
                }
                _ = fi14
            }
            v12 = 0
        }
        _ = v12
    } else {
        var v18 uint = (v2 >> (((^(v2))) & 63))
        if v11 {
            v6 = -2147483648
        } else {
            if len(v10) > 0 {
                v10[0] = -979
            }
            delete(v1, "\n\t")
        }
        _ = v18
    }
    switch 32767 {
    case 32767:
        v8 += -9223372036854775808
    default:
        var v19 int64 = (v8 & -124)
        _ = v19
    }
    v8 -= -9223372036854775808
    for fi20 := 0; fi20 < 14; fi20++ {
        if (v3 >= (v3 - (v3 - 920))) {
            for fi21 := 0; fi21 < 13; fi21++ {
                if true {
                    var v22 int = (fi21 ^ (fi21 & -2147483648))
                    v23 := GenS2{F0: "a", F1: (v4 * (v4 * (^(v4)))), F2: false}
                    fmt.Println(v2, v6, v11)
                    _ = v22
                    _ = v23
                }
                v24 := v5[65535]
                var v25 int16 = 1
                _ = fi21
                _ = v24
                _ = v25
            }
            if (int16(v4) > 1000) {
                v26 := GenS3{F0: 1}
                _ = v26
            } else {
                v8 = v8
                v3 += 255
            }
            v27 := GenS4{F0: v2, F1: 639, F2: -32768}
            _ = v27
        }
        v28 := v5[4294967295]
        _ = fi20
        _ = v28
    }
    v29 := GenS4{F0: 65535, F1: 4294967295, F2: 0}
    for fi30 := 0; fi30 < 12; fi30++ {
        v31 := v1[("\n\t" + "a")]
        _ = fi30
        _ = v31
    }
    v3 -= (^(v3))
    v32 := v1["hello"]
    if ("foo bar" > ("hello" + ("foo bar" + "foo bar"))) {
        switch "\n\t" {
        case "":
            v3 += v3
        default:
            switch "a" {
            case "":
                for fi33 := 0; fi33 < 18; fi33++ {
                    v3 -= v3
                    var v34 uint64 = 18446744073709551615
                    _ = fi33
                    _ = v34
                }
            case "foo bar":
                var v35 bool = true
                _ = v35
            }
        }
        fmt.Println(v3)
        v4 = 0
    } else {
        v36 := v1[""]
        fmt.Println(v6, v11)
        v4 += (v4 - -125)
        _ = v36
    }
    v1["a"] = (v32 ^ (v32 / (((v32 & v32)) | 1)))
    v11 = (v3 < (v3 >> ((v2) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
