package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint64; F2 bool }
func genFn0(p0 string) int8 {
    v7 := GenS1{F0: -904}
    p0 = p0
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{fi8, fi8}
        if len(v9) > 0 {
            v9[0] = (fi8 << ((42) & 63))
        }
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v10
    _ = v7
    return -11
}
func genFn1() string {
    var v11 uint = 4294967295
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int = fi12
        var v14 uint8 = 1
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = ("a" + "hello")
    v11 += (v11 >> (((^(v11))) & 63))
    _ = v11
    _ = v15
    return (v15 + (((v15 + v15) + v15) + v15))
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 808
    var v3 string = "\n\t"
    var v4 uint32 = 126
    v5 := make(map[uint32]bool)
    v6 := []int{9223372036854775807}
    v16 := GenS2{F0: 2147483647, F1: v2}
    v17 := GenS1{F0: 381}
    v3 = (v3 + ("a" + ((v3 + "\n\t") + v3)))
    _ = len(v6)
    for fi18 := 0; fi18 < 5; fi18++ {
        if true {
            v4 = v4
            v19 := GenS3{F0: false}
            if true {
                var v20 uint32 = (v4 & (v4 / (((v4 % ((173) | 1))) | 1)))
                switch (v3 + "hello") {
                case "hello":
                    var v21 uint64 = 1
                    _ = v21
                case "foo bar":
                    v2 = v2
                }
                _ = v20
            }
            _ = v19
        } else {
            v4 += (v4 << ((4294967295) & 63))
            if len(v6) > 0 {
                v6[0] = fi18
            }
        }
        v22 := GenS3{F0: true}
        _ = fi18
        _ = v22
    }
    v4 -= v4
    v3 = v3
    if true {
        v23 := GenS4{F0: -2147483648, F1: 0, F2: (!((2147483647 <= -1) || (false && (65535 >= 705))))}
        _ = v23
    }
    fmt.Println(v2)
    var v24 int = 378
    if ((v3 + (v3 + ("a" + v3))) > "a") {
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(fi25, v3, v24)
            v2 -= (v2 ^ (^(v2)))
            switch v3 {
            case "hello":
                var v26 string = ""
                _ = v26
            case "a":
                fmt.Println(v24, v4, v3)
            case "":
                v3 = v3
            default:
                fmt.Println(v3, v3, v24)
            }
            _ = fi25
        }
    } else {
        if (!(float32(1e10) <= float32(0.1))) {
            fmt.Println(v24, v24)
            switch "a" {
            case "\n\t":
                v27 := GenS4{F0: 0, F1: 4294967295, F2: (!(1.5 != -2.25))}
                _ = v27
            case "a":
                v28 := GenS3{F0: ("hello" >= v3)}
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 9; fi29++ {
                for fi30 := 0; fi30 < 12; fi30++ {
                    v1[(fi29 - (v24 + 0))] = -1
                    _ = fi30
                }
                v4 = 0
                _ = fi29
            }
            v31 := GenS2{F0: (^(v24)), F1: v2}
            var v32 uint64 = uint64(v4)
            _ = v31
            _ = v32
        }
        var v33 bool = true
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
