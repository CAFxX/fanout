package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0() uint16 {
    var v6 uint8 = 128
    v6 = v6
    v6 = (v6 & v6)
    v7 := GenS1{F0: 32767, F1: 1}
    _ = v6
    _ = v7
    return 832
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    v3 := make(map[uint32]bool)
    v4 := []int{}
    var v5 float32 = 1e10
    if true {
        for fi8 := 0; fi8 < 8; fi8++ {
            v2 = 0
            v9 := v3[0]
            _ = fi8
            _ = v9
        }
        if ((^(v2)) == (v2 * 156)) {
            var v10 string = "hello"
            for fi11 := 0; fi11 < 2; fi11++ {
                if (42 < 0) {
                    var v12 int32 = 658
                    _ = v12
                }
                v2 = 2
                v13 := GenS2{F0: 0}
                _ = fi11
                _ = v13
            }
            _ = v10
        }
        v14 := GenS2{F0: 42}
        _ = v14
    }
    _ = len(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v2 -= v2
            v17 := GenS3{F0: "foo bar", F1: 9223372036854775807}
            _ = fi16
            _ = v17
        }
        var v18 uint64 = 2
        _ = fi15
        _ = v18
    }
    var v19 uint = 65535
    v2 = (v2 & 1)
    v5 = v5
    var v20 uint8 = 0
    v20 -= 255
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    _ = len(v1)
    var v21 uint64 = (v2 * (^(v2)))
    var v22 int64 = int64(v19)
    v23 := GenS3{F0: ("\n\t" + "foo bar"), F1: (^(v22))}
    v24 := GenS3{F0: ("hello" + ("a" + ("hello" + "hello"))), F1: (v22 + (v22 >> ((18446744073709551615) & 63)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
