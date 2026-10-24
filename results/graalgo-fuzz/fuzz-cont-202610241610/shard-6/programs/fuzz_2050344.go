package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 int }

func main() {
    v1 := []uint32{755}
    var v2 string = (("a" + ("" + "\n\t")) + (("" + "") + ("foo bar" + ("foo bar" + ""))))
    var v3 uint8 = 42
    var v4 float32 = (3.14159 / -2.25)
    var v5 uint = 18446744073709551615
    switch ("hello" + "\n\t") {
    case "":
        var v6 uint16 = 1
        _ = v6
    }
    var v7 int = -249
    var v8 uint8 = (v3 / ((v3) | 1))
    if true {
        var v9 int = (^(v7))
        var v10 int16 = 1
        _ = v10
        _ = v9
    }
    var v11 uint = v5
    if false {
        var v12 int32 = -717
        _ = v12
    } else {
        v3 -= 127
        v13 := GenS1{F0: 146, F1: v5, F2: -128}
        _ = v13
    }
    v14 := GenS1{F0: (^(v7)), F1: v5, F2: -21}
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: -108}
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: (v7 % (((v7 * v7)) | 1))}
    fmt.Println(v11, v5)
    var v18 uint64 = 500
    if (1 > (v18 | v18)) {
        v19 := GenS2{F0: (v7 | -636)}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v5 = (v11 >> (((v5 >> (((v11 - v11)) & 63))) & 63))
            _ = fi20
        }
        fmt.Println(v7, v4)
        if false {
            var v21 float32 = -2.25
            v18 = uint64(v5)
            var v22 int = 1
            _ = v21
            _ = v22
        } else {
            v3 -= (v8 / (((v8 & (v8 << ((v11) & 63)))) | 1))
            var v23 uint32 = 834
            _ = v23
        }
    }
    v8 -= v3
    fmt.Println(v3)
    if false {
        v7 -= v7
    } else {
        fmt.Println(v3)
        v18 = v18
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 15; fi25++ {
                v18 += (v18 ^ 2)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi25
            }
            v3 = (^(v8))
            var v26 uint = 18446744073709551615
            _ = fi24
            _ = v26
        }
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v18 = v18
        fmt.Println(v3)
        var v28 uint16 = 0
        _ = fi27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
