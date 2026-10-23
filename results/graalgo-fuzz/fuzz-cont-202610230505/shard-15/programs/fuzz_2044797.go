package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    var v5 uint = 352
    v6 := GenS1{F0: 0, F1: 1000}
    v7 := GenS2{F0: 1, F1: -271, F2: 42}
    var v8 bool = false
    if (!(float64(-2.25) <= -2.25)) {
        var v9 uint16 = 32767
        if v8 {
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(v5)
                fmt.Println(v9)
                v11 := []int{fi10}
                _ = fi10
                _ = v11
            }
        }
        _ = v9
    } else {
        v8 = false
        if (1.5 > 3.14159) {
            var v12 uint = (^(v5))
            v13 := GenS1{F0: 42, F1: -32768}
            if v8 {
                v8 = v8
            }
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 0, F1: -1}
            _ = v14
        }
    }
    fmt.Println(v5, v5)
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            switch ("\n\t" + ((("\n\t" + ("" + ("foo bar" + ("\n\t" + "foo bar")))) + "foo bar") + "a")) {
            case "":
                switch fi15 {
                case -1:
                    fmt.Println(fi15, fi15, fi16)
                case 2147483647:
                    var v17 int8 = -128
                    fmt.Println(v8, fi16)
                    v17 += (v17 & (v17 + (^(v17))))
                    _ = v17
                default:
                    v8 = (32767 <= 65535)
                }
            }
            var v18 int8 = 100
            switch (fi16 / (((fi16 - 128)) | 1)) {
            case -42:
                fmt.Println(v5, v8)
            case 32767:
                v19 := GenS2{F0: 1, F1: 2147483647, F2: -396}
                _ = v19
            case 2147483647:
                v5 = v5
            }
            _ = fi16
            _ = v18
        }
        _ = fi15
    }
    var v20 int64 = -9223372036854775808
    delete(v4, 541)
    v21 := GenS1{F0: 1, F1: -351}
    v8 = v8
    fmt.Println(v5)
    if v8 {
        var v22 float64 = 3.14159
        var v23 bool = ((v5 ^ v5) >= 2)
        v20 -= (v20 ^ (v20 >> (((v5 + (^(v5)))) & 63)))
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
