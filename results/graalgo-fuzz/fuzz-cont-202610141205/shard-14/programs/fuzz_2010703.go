package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint8, p1 string, p2 int64) uint {
    var v4 string = (("a" + p1) + (("foo bar" + p1) + p1))
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{(fi5 / (((^(fi5))) | 1)), (fi5 + fi5), (fi5 >> ((635) & 63)), fi5}
        if true {
            if true {
                if true {
                    p0 += (p0 / ((208) | 1))
                    switch p1 {
                    case "hello":
                        v7 := GenS1{F0: 65535, F1: (!true), F2: 87}
                        _ = v7
                    default:
                        _ = len(v6)
                    }
                }
            }
            v4 = "foo bar"
            v8 := GenS1{F0: 4294967295, F1: (false || false), F2: int32(p0)}
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS2{F0: 386, F1: "foo bar"}
    for fi10 := 0; fi10 < 3; fi10++ {
        switch v4 {
        case "":
            p2 = (-(p2))
        case "\n\t":
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int32 = 0
                fmt.Println(p0, p2)
                _ = fi11
                _ = v12
            }
        case "foo bar":
            var v13 int64 = p2
            _ = v13
        default:
            switch v4 {
            case "\n\t":
                switch "" {
                case "a":
                    var v14 bool = true
                    _ = v14
                case "\n\t":
                    v15 := []int{fi10}
                    _ = v15
                case "hello":
                    p2 += -9223372036854775808
                default:
                    for fi16 := 0; fi16 < 6; fi16++ {
                        var v17 int16 = -32768
                        v18 := GenS3{F0: 268}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            case "a":
                var v19 int = (-(fi10))
                _ = v19
            }
        }
        switch v4 {
        case "hello":
            fmt.Println(p2, p1)
        case "a":
            v20 := GenS4{F0: 704}
            _ = v20
        case "":
            var v21 int16 = 42
            _ = v21
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 255
    var v3 string = "\n\t"
    v22 := []int{0, 128}
    v2 = v2
    v3 = v3
    delete(v1, 975)
    fmt.Println(v3, v2)
    var v23 uint16 = 42
    v23 = (v23 >> ((255) & 63))
    var v24 uint32 = 1
    v23 = v23
    v23 = v23
    var v25 uint16 = (v23 + v23)
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
