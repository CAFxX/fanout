package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []string{("" + "")}
    var v2 int64 = 42
    var v3 int32 = 42
    var v4 bool = (-128 <= 1)
    v5 := make(map[int]int)
    v6 := make(map[string]int)
    var v7 uint32 = 42
    var v8 int64 = (v2 % ((-1) | 1))
    switch ("" + "") {
    case "a":
        if len(v1) > 0 {
            v1[0] = ("a" + "")
        }
    case "\n\t":
        for fi9 := 0; fi9 < 12; fi9++ {
            switch 32768 {
            case 4294967295:
                _ = len(v1)
            case 664:
                for fi10 := 0; fi10 < 20; fi10++ {
                    fmt.Println(v7)
                    _ = fi10
                }
            case -1:
                v11 := GenS1{F0: 700, F1: (fi9 ^ fi9), F2: 240}
                _ = v11
            }
            v7 = v7
            _ = fi9
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v12 := GenS1{F0: 32767, F1: 697, F2: 42}
    v13 := v6[("hello" + "a")]
    fmt.Println(v7, v8)
    for fi14 := 0; fi14 < 19; fi14++ {
        v3 += 0
        _ = fi14
    }
    v15 := GenS1{F0: 65535, F1: (v13 / ((v13) | 1)), F2: 0}
    if true {
        v4 = false
        v2 -= -309
    }
    delete(v5, v13)
    fmt.Println(v4, v13)
    v16 := GenS2{F0: ((((v3 % (((v3 - v3)) | 1)) <= (v3 >> ((99) & 63))) && v4) && v4)}
    v17 := GenS1{F0: 17, F1: v13, F2: 241}
    v18 := GenS3{F0: (^(v3))}
    for fi19 := 0; fi19 < 13; fi19++ {
        if (("hello" > "") && v4) {
            fmt.Println(v3, fi19)
            for fi20 := 0; fi20 < 14; fi20++ {
                v1 = append(v1, ("\n\t" + ""))
                var v21 int8 = 42
                _ = fi20
                _ = v21
            }
        } else {
            var v22 string = ""
            _ = v22
        }
        v13 -= int(v7)
        for fi23 := 0; fi23 < 13; fi23++ {
            _ = len(v1)
            _ = fi23
        }
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
