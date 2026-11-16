package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int32; F1 int32; F2 int16 }

func main() {
    var v1 string = ""
    var v2 int64 = -1
    v3 := make(map[string]bool)
    var v4 uint16 = 1
    v2 += 0
    var v5 uint16 = 65535
    v6 := []int{-2147483648, 4294967295}
    v7 := GenS1{F0: 2147483647, F1: 383}
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS2{F0: (42 > 0), F1: 127, F2: 431}
        v10 := GenS3{F0: 298, F1: (v2 + 9223372036854775807)}
        v11 := GenS4{F0: -353, F1: 888, F2: 42}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch "hello" {
        case "\n\t":
            for fi13 := 0; fi13 < 20; fi13++ {
                if (!false) {
                    v5 = v4
                    v14 := v3[(v1 + (v1 + v1))]
                    _ = v14
                } else {
                    v3[(v1 + v1)] = ((2 <= 42) && (255 < 0))
                    v3["foo bar"] = true
                    v15 := v3[((((v1 + "\n\t") + "a") + "") + "\n\t")]
                    _ = v15
                }
                if false {
                    v1 = (("" + v1) + v1)
                    v4 = v5
                    fmt.Println(v1, fi12, v1)
                }
                _ = fi13
            }
        default:
            var v16 int8 = 0
            _ = v16
        }
        if true {
            v1 = (("\n\t" + "hello") + (v1 + "\n\t"))
        } else {
            var v17 string = ((v1 + "") + v1)
            v18 := GenS1{F0: int(v4), F1: 4294967295}
            switch (((v1 + (v17 + (("foo bar" + (v17 + (("" + ((((v1 + ("hello" + (v1 + v1))) + (v1 + ("\n\t" + v17))) + "hello") + (v17 + v1))) + (v17 + v1)))) + "hello"))) + v17) + v1) {
            case "\n\t":
                for fi19 := 0; fi19 < 20; fi19++ {
                    fmt.Println(fi12, fi12)
                    _ = len(v3)
                    v2 = v2
                    _ = fi19
                }
            }
            _ = v17
            _ = v18
        }
        _ = fi12
    }
    if false {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint8 = 56
            v6 = append(v6, 9223372036854775807)
            _ = fi20
            _ = v21
        }
        v4 = v4
        v1 = (v1 + v1)
    }
    v1 = (v1 + "a")
    var v22 uint16 = v4
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
