package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    var v1 int = 0
    var v2 uint64 = 2
    v3 := []uint32{}
    var v4 int = (v1 >> ((0) & 63))
    if false {
        var v5 uint = 65535
        var v6 bool = false
        _ = len(v3)
        _ = v5
        _ = v6
    }
    switch "\n\t" {
    case "hello":
        for fi7 := 0; fi7 < 5; fi7++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi7
        }
    case "\n\t":
        v8 := GenS1{F0: -1}
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: -1}
        _ = v9
    }
    var v10 int64 = 867
    var v11 float64 = 0.1
    v2 = 1
    v1 = v4
    var v12 uint16 = 774
    switch (v4 & v1) {
    case 32767:
        v13 := GenS2{F0: 848, F1: 4294967295, F2: 128}
        _ = v13
    default:
        v4 -= 32768
    }
    _ = len(v3)
    var v14 int64 = (v10 | -594)
    for fi15 := 0; fi15 < 10; fi15++ {
        if false {
            v16 := GenS3{F0: false, F1: true}
            switch "\n\t" {
            case "hello":
                for fi17 := 0; fi17 < 12; fi17++ {
                    if false {
                        var v18 uint32 = 2147483647
                        _ = v18
                    }
                    _ = fi17
                }
            case "":
                var v19 uint16 = (v12 >> ((1) & 63))
                _ = v19
            default:
                _ = len(v3)
            }
            _ = v16
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS2{F0: 2147483647, F1: (v1 << ((400) & 63)), F2: 23}
            var v22 int64 = v10
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 uint64 = v2
        _ = fi15
        _ = v23
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
