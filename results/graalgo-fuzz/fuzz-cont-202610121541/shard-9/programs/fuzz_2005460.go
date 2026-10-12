package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 42
    var v3 string = (((((("foo bar" + v1) + ("foo bar" + v1)) + (v1 + (("" + v1) + (((v1 + v1) + "a") + "hello")))) + v1) + v1) + (v1 + v1))
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 11; fi5++ {
        v2 = 399
        switch (fi5 & fi5) {
        case 9223372036854775807:
            var v6 int = (fi5 ^ 4294967295)
            _ = v6
        }
        var v7 string = "\n\t"
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{(fi8 | fi8), fi8, (fi8 | (fi8 / ((-730) | 1))), (fi8 % ((2147483647) | 1))}
        _ = fi8
        _ = v9
    }
    _ = len(v4)
    v10 := []int{}
    v2 += 2
    fmt.Println(v3, v2, v3)
    var v11 string = v3
    var v12 int = 917
    v2 = (v2 / ((v2) | 1))
    if false {
        var v13 int16 = 32767
        _ = v13
    } else {
        v14 := GenS1{F0: (v2 << ((1) & 63)), F1: 0}
        var v15 bool = false
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 32767, F1: 4294967295}
    v1 = ((((v1 + ((v11 + "hello") + v11)) + "") + (v11 + v1)) + v3)
    v11 = v11
    v2 += (v2 - 953)
    for fi17 := 0; fi17 < 15; fi17++ {
        v3 = ("" + (v3 + ("a" + v3)))
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 bool = false
            _ = fi18
            _ = v19
        }
        v20 := GenS2{F0: 912}
        _ = fi17
        _ = v20
    }
    switch v1 {
    case "\n\t":
        if false {
            var v21 float32 = 1e10
            _ = v21
        } else {
            _ = len(v4)
            if true {
                v22 := GenS3{F0: 2147483647, F1: -1}
                v1 = "foo bar"
                _ = v22
            } else {
                _ = len(v4)
            }
        }
    case "a":
        if false {
            fmt.Println(v2)
            var v23 int32 = -795
            switch (v12 & (v12 + -2147483648)) {
            case 128:
                switch (v12 - -1) {
                case -340:
                    _ = len(v10)
                default:
                    var v24 int = v12
                    _ = v24
                }
            }
            _ = v23
        }
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v3 = ("foo bar" + v3)
        fmt.Println(v1, fi25)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
