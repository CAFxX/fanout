package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint32 }

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: 823, F1: ("" + "\n\t")}
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int64 = -1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    var v8 string = ("foo bar" + ("" + ("foo bar" + "foo bar")))
    switch v8 {
    case "":
        v8 = "foo bar"
    case "hello":
        v8 = v8
    default:
        v9 := GenS1{F0: 7, F1: (v8 + "foo bar")}
        _ = v9
    }
    v8 = ("\n\t" + "foo bar")
    delete(v3, v8)
    v8 = "a"
    v8 = v8
    for fi10 := 0; fi10 < 16; fi10++ {
        if true {
            var v11 string = v8
            _ = v11
        } else {
            _ = len(v3)
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: 18446744073709551615, F1: v8}
                if (-377 <= 0) {
                    v14 := GenS1{F0: 0, F1: v8}
                    _ = v14
                } else {
                    fmt.Println(fi12, fi10)
                    fmt.Println(v8, fi10, fi12)
                    v15 := GenS1{F0: 530, F1: (v8 + ("\n\t" + v8))}
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        }
        var v16 uint = 65535
        fmt.Println(v16, v8)
        _ = fi10
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: 18446744073709551615, F1: "hello"}
            if false {
                v8 = (v8 + (("" + "hello") + ""))
                v20 := GenS1{F0: 1, F1: (("hello" + "\n\t") + v8)}
                _ = v20
            }
            _ = fi18
            _ = v19
        }
        v21 := GenS1{F0: 288, F1: v8}
        _ = fi17
        _ = v21
    }
    var v22 uint16 = 42
    for fi23 := 0; fi23 < 13; fi23++ {
        if false {
            v24 := GenS2{F0: 2147483647}
            _ = v24
        }
        v25 := GenS3{F0: -2147483648, F1: 755}
        _ = fi23
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v8
}
