package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint32) int {
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            if (642 >= fi7) {
                switch "" {
                case "foo bar":
                    v8 := GenS1{F0: 1000}
                    _ = v8
                case "\n\t":
                    p0 += p0
                case "a":
                    p0 += p0
                }
            } else {
                fmt.Println(fi7)
            }
            var v9 uint16 = 1
            v10 := GenS2{F0: "\n\t"}
            _ = fi7
            _ = v10
            _ = v9
        }
        fmt.Println(p0)
        _ = fi6
    }
    var v11 uint16 = 717
    var v12 uint8 = uint8(v11)
    v13 := GenS3{F0: int32(v12), F1: 1}
    if (4294967295 >= 18446744073709551615) {
        var v14 float32 = 3.14159
        v15 := GenS1{F0: -32768}
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 340
}
func genFn1(p0 int64, p1 string, p2 string) uint16 {
    switch (((p2 + p2) + "\n\t") + ((p1 + p1) + (("" + p2) + "hello"))) {
    case "foo bar":
        if (42 == -592) {
            switch "\n\t" {
            case "foo bar":
                p1 = (p2 + p2)
            case "hello":
                p0 = 0
            default:
                if false {
                    var v16 int64 = (p0 | (p0 / (((p0 + (p0 / (((p0 << ((186) & 63))) | 1)))) | 1)))
                    var v17 uint = 909
                    _ = v16
                    _ = v17
                }
            }
            var v18 int32 = -898
            _ = v18
        } else {
            p1 = ""
            p0 = 42
            v19 := []int{4294967295, 42, 32768, 1}
            _ = v19
        }
    }
    p2 = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int16 = -950
    var v2 int8 = -1
    v3 := []string{"foo bar", "hello", "hello"}
    v4 := []uint32{65535}
    var v5 int = 128
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "hello")
    }
    if false {
        var v20 int16 = 32767
        if len(v4) > 0 {
            v4[0] = 0
        }
        _ = v20
    }
    fmt.Println(v2, v5)
    v21 := GenS3{F0: 42, F1: 42}
    v2 = (v2 >> ((222) & 63))
    var v22 uint = 65535
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 string = ("foo bar" + "\n\t")
            v26 := GenS1{F0: -662}
            var v27 int32 = 2147483647
            var v28 int32 = (-(v27))
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            v5 += (fi23 / (((v5 | v5)) | 1))
            _ = fi29
        }
        _ = fi23
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
