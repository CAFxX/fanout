package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 bool; F2 int64 }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 int = 241
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 100
    v1 = (v1 >> ((842) & 63))
    _ = len(v2)
    var v5 int64 = 9223372036854775807
    var v6 uint = 997
    var v7 uint64 = 42
    v8 := GenS1{F0: 1}
    if (42 != (v1 * v1)) {
        for fi9 := 0; fi9 < 4; fi9++ {
            switch fi9 {
            case 32767:
                var v10 uint8 = uint8(v7)
                _ = v10
            case 420:
                v11 := []int{fi9}
                _ = v11
            }
            v12 := []int{v1}
            _ = fi9
            _ = v12
        }
        switch "\n\t" {
        case "foo bar":
            if true {
                delete(v2, (v1 / ((2147483647) | 1)))
                v6 += 101
                var v13 bool = true
                _ = v13
            } else {
                switch v1 {
                case 2147483647:
                    var v14 float32 = float32(float64((1e10 / (float32(1e10) * (0.1 * (0.1 * 1e10))))))
                    _ = v14
                }
                if (1000 < 42) {
                    v6 = (^(v6))
                } else {
                    v15 := GenS1{F0: 424}
                    v16 := []int{(v1 % ((v1) | 1)), 32768, v1}
                    _ = len(v16)
                    _ = v15
                    _ = v16
                }
            }
        case "":
            var v17 float32 = 1e10
            _ = v17
        default:
            for fi18 := 0; fi18 < 16; fi18++ {
                v5 -= v5
                for fi19 := 0; fi19 < 2; fi19++ {
                    v20 := GenS1{F0: 42}
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
        v1 = (v1 | (^(v1)))
    }
    v21 := []int{4294967295, v1, 9223372036854775807}
    switch ("" + "hello") {
    case "foo bar":
        v22 := GenS2{F0: "\n\t", F1: (1 <= 4294967295), F2: v5}
        _ = v22
    case "hello":
        var v23 bool = (2147483647 <= 2147483647)
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 7; fi24++ {
            if true {
                delete(v2, (v1 >> (((v6 & (v6 & v6))) & 63)))
                for fi25 := 0; fi25 < 18; fi25++ {
                    var v26 uint64 = 742
                    var v27 int = (fi25 - -535)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                var v28 int16 = 929
                _ = v28
            }
            _ = fi24
        }
    }
    if false {
        v29 := GenS3{F0: true, F1: 42}
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := GenS1{F0: 194}
            var v32 string = "hello"
            for fi33 := 0; fi33 < 12; fi33++ {
                if false {
                    v5 -= v5
                } else {
                    var v34 uint16 = 137
                    _ = v34
                }
                _ = fi33
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        v1 -= int(v6)
        _ = v29
    }
    var v35 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
