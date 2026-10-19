package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 string }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []int{32767, -1, -115}
    var v2 bool = ("\n\t" >= ("\n\t" + ("a" + ("a" + "hello"))))
    v3 := []int{4294967295, 0, 4294967295, 32768}
    v4 := []uint32{42, 2147483648, 2}
    v5 := make(map[uint32]bool)
    var v6 int32 = 0
    if v2 {
        if (v2 && v2) {
            fmt.Println(v6, v2)
        }
        var v7 uint = 65535
        if v2 {
            var v8 uint8 = 1
            v6 = (v6 % (((v6 * v6)) | 1))
            _ = v8
        } else {
            switch "hello" {
            case "":
                var v9 bool = v2
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: 333, F1: 1, F2: false}
                _ = v10
            case "hello":
                var v11 int32 = v6
                _ = v11
            default:
                switch "foo bar" {
                case "a":
                    var v12 string = ""
                    _ = v12
                case "foo bar":
                    var v13 uint64 = 173
                    _ = v13
                }
            }
            v6 += v6
            v14 := GenS2{F0: 127, F1: v6, F2: "a"}
            _ = v14
        }
        _ = v7
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS2{F0: 30, F1: -1, F2: ("hello" + (("\n\t" + "hello") + "\n\t"))}
            _ = fi15
            _ = v16
        }
        switch "foo bar" {
        case "foo bar":
            v2 = (1 > -32768)
        default:
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 float64 = 0.1
                _ = fi17
                _ = v18
            }
        }
    }
    _ = len(v5)
    v19 := GenS3{F0: 0}
    v5[2147483647] = (-461 == 9223372036854775807)
    var v20 uint32 = 2
    v21 := GenS3{F0: 905}
    v22 := GenS3{F0: 42}
    var v23 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
