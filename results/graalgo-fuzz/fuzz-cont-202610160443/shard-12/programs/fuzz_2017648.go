package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }

func main() {
    var v1 int8 = 0
    var v2 uint8 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    v5 := []uint32{2147483648}
    var v6 int8 = -78
    switch "hello" {
    case "\n\t":
        var v7 uint16 = 32767
        _ = v7
    }
    if (false || true) {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int8 = v1
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(fi10, v9)
                var v11 uint8 = (v2 / (((v2 + v2)) | 1))
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        var v12 int8 = 91
        _ = v12
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 string = "foo bar"
            fmt.Println(v6, fi13, v6)
            _ = fi13
            _ = v14
        }
        switch 212 {
        case 1:
            if ("hello" < "a") {
                fmt.Println(v6)
            }
        case 42:
            fmt.Println(v6)
        }
    }
    fmt.Println(v1)
    v15 := GenS1{F0: v2, F1: uint64(v2), F2: "foo bar"}
    var v16 bool = (!(1 != 0))
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS1{F0: (v2 * 1), F1: 389, F2: "\n\t"}
        var v19 bool = (1 <= 513)
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: v2, F1: 4294967295, F2: ("" + ("\n\t" + ("" + "\n\t")))}
    fmt.Println(v2, v2)
    var v21 uint = 388
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
