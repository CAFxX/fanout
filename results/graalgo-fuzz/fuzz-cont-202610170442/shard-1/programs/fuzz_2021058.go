package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 bool; F2 string }
type GenS5 struct { F0 uint64; F1 int }

func main() {
    v1 := []uint32{2147483647}
    v2 := make(map[int]bool)
    var v3 uint32 = 71
    v4 := make(map[uint32]string)
    var v5 uint = 255
    var v6 bool = (!(-106 >= 100))
    v5 = v5
    v7 := GenS1{F0: 65535}
    v8 := GenS2{F0: (!v6)}
    v9 := GenS1{F0: 65535}
    v10 := GenS3{F0: 51}
    v2[341] = ("foo bar" > "hello")
    if false {
        switch 9223372036854775807 {
        case 295:
            v1 = append(v1, v3)
        case 32768:
            var v11 bool = true
            _ = v11
        default:
            v6 = (!v6)
        }
        v6 = true
        v12 := GenS1{F0: 2}
        _ = v12
    }
    v6 = (!false)
    var v13 int64 = -1
    v14 := GenS3{F0: -221}
    v6 = (false || (true && true))
    if (700 == -395) {
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS4{F0: 0, F1: v6, F2: ""}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint64 = uint64(v3)
            _ = fi17
            _ = v18
        }
        var v19 bool = v6
        _ = v19
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int = (fi20 & fi20)
            var v22 bool = (4294967295 > 4294967295)
            var v23 int8 = 42
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        if (827 <= 0) {
            if len(v1) > 0 {
                v1[0] = 0
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                for fi25 := 0; fi25 < 5; fi25++ {
                    v13 += v13
                    v26 := GenS5{F0: 196, F1: 4294967295}
                    _ = fi25
                    _ = v26
                }
                v27 := GenS1{F0: 993}
                _ = fi24
                _ = v27
            }
        } else {
            if false {
                for fi28 := 0; fi28 < 5; fi28++ {
                    fmt.Println(v13, v13)
                    var v29 int8 = 42
                    _ = fi28
                    _ = v29
                }
            }
        }
        _ = len(v1)
    }
    v30 := GenS3{F0: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
