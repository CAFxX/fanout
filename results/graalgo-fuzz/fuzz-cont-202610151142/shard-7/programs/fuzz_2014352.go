package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []string{}
    var v2 string = ""
    v3 := make(map[int]string)
    v4 := []uint32{}
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(v2, fi6, fi6)
            v2 = ((v2 + "foo bar") + "hello")
            _ = fi6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            v8 := v3[(fi7 << ((412) & 63))]
            switch 9223372036854775807 {
            case 821:
                switch "a" {
                case "":
                    v1 = append(v1, ((v2 + v2) + ("\n\t" + v8)))
                }
            case 9223372036854775807:
                var v9 int = (fi7 | -1)
                _ = v9
            case 125:
                v10 := GenS1{F0: 9223372036854775807, F1: 1, F2: (true || (4294967295 < 42))}
                _ = v10
            }
            _ = len(v4)
            _ = v8
        } else {
            if (("" == "") || (false || false)) {
                v2 = ""
            } else {
                var v11 bool = false
                _ = v11
            }
            v2 = (v2 + (v2 + v2))
            v12 := GenS2{F0: 252}
            _ = v12
        }
        _ = fi7
    }
    v2 = v2
    var v13 int32 = 42
    v2 = (((v2 + "hello") + v2) + v2)
    v14 := GenS3{F0: 620}
    var v15 bool = (true && (v13 >= 1))
    var v16 uint8 = 255
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint16 = 32767
        _ = fi17
        _ = v18
    }
    _ = len(v1)
    if false {
        if v15 {
            v16 += v16
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
