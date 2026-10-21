package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int; F2 uint }

func main() {
    var v1 int32 = -923
    v2 := []string{"foo bar"}
    var v3 uint16 = 144
    v4 := []uint32{65535, 4294967295, 2}
    v4 = append(v4, 2147483647)
    var v5 int16 = -32768
    v6 := GenS1{F0: 32767, F1: false, F2: v3}
    switch 32767 {
    case 887:
        if false {
            switch 1 {
            case 32767:
                fmt.Println(v3, v1, v3)
            case 0:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                v7 := GenS2{F0: 255}
                _ = v7
            }
            v8 := GenS3{F0: 297, F1: -836, F2: 241}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: (^(fi9)), F1: false, F2: v3}
                switch ("" + "hello") {
                case "":
                    v1 = v1
                case "foo bar":
                    var v11 string = (("\n\t" + "\n\t") + "foo bar")
                    _ = v11
                default:
                    v3 = v3
                }
                var v12 uint32 = 333
                _ = fi9
                _ = v10
                _ = v12
            }
        }
    case 4294967295:
        var v13 uint64 = 647
        _ = v13
    case 128:
        fmt.Println(v3)
    default:
        v5 = 607
    }
    v5 = v5
    var v14 uint32 = 2
    var v15 int64 = 42
    var v16 uint32 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
