package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint32 = 71
    var v2 string = ""
    var v3 uint64 = 0
    var v4 int64 = -9223372036854775808
    v2 = (v2 + v2)
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{fi5, 4294967295, fi5}
        var v7 string = v2
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = int(v1)
    for fi9 := 0; fi9 < 3; fi9++ {
        v8 = (fi9 >> ((158) & 63))
        var v10 int64 = int64(v8)
        _ = fi9
        _ = v10
    }
    v3 = 4294967295
    var v11 int8 = -36
    v12 := GenS1{F0: -965}
    v3 += (v3 * (v3 >> ((1) & 63)))
    v13 := []int{(v8 / ((v8) | 1)), (v8 - v8)}
    var v14 int = (v8 | (v8 << ((483) & 63)))
    if ((v3 >> ((42) & 63)) < v3) {
        fmt.Println(v3, v14, v14)
        var v15 int = (v8 | (^(v8)))
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            switch -504 {
            case 32768:
                var v18 int32 = 0
                _ = v18
            default:
                _ = len(v13)
            }
            v19 := GenS2{F0: 413, F1: 906}
            _ = fi17
            _ = v19
        }
        if false {
            v13 = append(v13, (v14 - (fi16 << ((42) & 63))))
            var v20 int64 = v4
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 uint16 = 65535
                var v23 int = (fi21 & v8)
                switch v2 {
                case "a":
                    v1 += v1
                case "foo bar":
                    var v24 uint64 = (v3 << ((42) & 63))
                    _ = v24
                case "hello":
                    var v25 int16 = -32768
                    _ = v25
                default:
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        } else {
            switch (v14 / ((v8) | 1)) {
            case 4294967295:
                v2 = ""
            }
        }
        _ = fi16
    }
    switch -9223372036854775808 {
    case 127:
        v1 += v1
    case -575:
        var v26 int16 = 496
        _ = v26
    default:
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                switch (v14 + fi27) {
                case 0:
                    var v29 int8 = 12
                    _ = v29
                default:
                    var v30 int32 = -1
                    _ = v30
                }
                if true {
                    var v31 float32 = 1.5
                    v32 := GenS1{F0: 357}
                    _ = v31
                    _ = v32
                } else {
                    v3 = 0
                    v33 := GenS1{F0: 244}
                    v1 += uint32(v3)
                    _ = v33
                }
                var v34 int16 = 1
                _ = fi28
                _ = v34
            }
            v35 := GenS1{F0: 494}
            _ = fi27
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
