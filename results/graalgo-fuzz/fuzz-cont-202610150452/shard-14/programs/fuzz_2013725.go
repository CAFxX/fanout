package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int }

func main() {
    var v1 uint = 4294967295
    var v2 uint32 = 65535
    var v3 uint16 = 615
    var v4 uint64 = 228
    v1 -= 0
    switch "" {
    case "\n\t":
        if ((v2 & v2) <= (v2 + 2147483648)) {
            v5 := GenS1{F0: 661}
            for fi6 := 0; fi6 < 1; fi6++ {
                v7 := GenS1{F0: int64(v3)}
                fmt.Println(v3, fi6)
                var v8 int16 = 104
                _ = fi6
                _ = v7
                _ = v8
            }
            v4 += 1
            _ = v5
        } else {
            fmt.Println(v1)
        }
    case "a":
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{}
            v11 := GenS1{F0: 526}
            var v12 uint32 = 340
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        switch 988 {
        case -1:
            var v13 uint8 = 128
            _ = v13
        }
    }
    v14 := GenS2{F0: (^(v2)), F1: false, F2: 32767}
    v4 += (^(v4))
    var v15 float32 = 0.1
    v16 := []int{int(v1), 197, -1}
    v3 -= (v3 | v3)
    var v17 float64 = 0.1
    fmt.Println(v3, v3, v2)
    fmt.Println(v1)
    if false {
        switch "foo bar" {
        case "foo bar":
            if (!(9223372036854775807 == -336)) {
                switch -597 {
                case 42:
                    fmt.Println(v15, v15)
                }
                v18 := GenS2{F0: v2, F1: true, F2: 32767}
                if false {
                    var v19 int16 = int16(v1)
                    _ = v19
                } else {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    switch "foo bar" {
                    case "":
                        v3 = (v3 * (v3 + (v3 - v3)))
                    }
                }
                _ = v18
            } else {
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 int64 = int64(v4)
                    var v22 int8 = 1
                    v23 := GenS1{F0: (v21 ^ v21)}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
        case "a":
            v24 := GenS2{F0: (v2 - (v2 ^ (v2 / (((v2 ^ (v2 + (v2 >> (((v1 * 65535)) & 63))))) | 1)))), F1: (!((!(!(!((-230 >= 42) || false)))) || true)), F2: 128}
            _ = v24
        case "hello":
            v1 += v1
        }
        v4 -= v4
    } else {
        for fi25 := 0; fi25 < 8; fi25++ {
            v4 = v4
            fmt.Println(v17)
            _ = fi25
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS1{F0: 9223372036854775807}
            var v28 int8 = 100
            _ = fi26
            _ = v27
            _ = v28
        }
        v16 = append(v16, int(v3))
    }
    v17 = v17
    var v29 bool = (true || true)
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 string = ""
        _ = fi30
        _ = v31
    }
    v29 = ((9223372036854775807 > 42) && (!(!true)))
    v32 := GenS1{F0: 42}
    if false {
        v17 += float64(v15)
        for fi33 := 0; fi33 < 2; fi33++ {
            if v29 {
                var v34 uint8 = 127
                var v35 string = "a"
                v36 := GenS2{F0: 2147483648, F1: v29, F2: fi33}
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi33
        }
        var v37 uint64 = v4
        _ = v37
    } else {
        for fi38 := 0; fi38 < 16; fi38++ {
            var v39 int32 = int32(v3)
            if (1000 > -701) {
                v3 = v3
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            } else {
                if len(v16) > 0 {
                    v16[0] = -9223372036854775808
                }
                _ = len(v16)
            }
            v40 := GenS2{F0: v2, F1: ((false || v29) || (v29 && false)), F2: fi38}
            _ = fi38
            _ = v39
            _ = v40
        }
        var v41 bool = true
        _ = v41
    }
    var v42 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v42
}
