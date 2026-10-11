package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int32 = 42
    var v2 int64 = 205
    var v3 bool = (40 < 2)
    fmt.Println(v1, v3)
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 bool = (-545 != (v1 % (((v1 ^ v1)) | 1)))
        fmt.Println(v3)
        _ = fi4
        _ = v5
    }
    fmt.Println(v2)
    v3 = ((v1 | (-(v1))) > v1)
    v6 := GenS1{F0: 1, F1: 347, F2: 18446744073709551615}
    v7 := []int{}
    var v8 uint = 4294967295
    if false {
        if (("a" + "") != "hello") {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint8 = 0
                _ = fi9
                _ = v10
            }
        } else {
            if true {
                fmt.Println(v8, v3)
                v11 := GenS2{F0: 9223372036854775807}
                v2 = v2
                _ = v11
            } else {
                if (-2.25 == -2.25) {
                    var v12 uint32 = 65535
                    _ = len(v7)
                    _ = v12
                }
                v2 += (v2 ^ v2)
            }
            _ = len(v7)
            v1 = (v1 << ((v8) & 63))
        }
        var v13 float32 = -2.25
        v2 = v2
        _ = v13
    }
    switch (("\n\t" + "foo bar") + ((("foo bar" + "\n\t") + "foo bar") + "")) {
    case "hello":
        var v14 uint = v8
        _ = v14
    }
    if v3 {
        v8 = 42
        v15 := GenS3{F0: 665}
        _ = v15
    }
    v2 = (-(v2))
    v16 := GenS2{F0: int(v1)}
    var v17 int = 32767
    v1 += v1
    v18 := GenS2{F0: (v17 << ((v8) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
