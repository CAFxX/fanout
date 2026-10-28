package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8; F1 string; F2 string }

func main() {
    var v1 int8 = 77
    v2 := make(map[uint32]string)
    var v3 uint = 113
    v4 := GenS1{F0: (42 != 42), F1: (v3 + (v3 & 2))}
    var v5 uint16 = 255
    v1 += v1
    v5 += 2
    var v6 string = "\n\t"
    v3 -= 42
    var v7 int16 = -1
    v6 = ("hello" + "hello")
    if false {
        switch 128 {
        case -2147483648:
            fmt.Println(v7, v7)
        case 0:
            v8 := []int{}
            _ = v8
        default:
            var v9 int32 = 649
            _ = v9
        }
    } else {
        v5 = (^(v5))
    }
    v10 := GenS1{F0: (("\n\t" + "a") == ("a" + "")), F1: v3}
    if true {
        var v11 uint = v3
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 uint64 = 4294967295
                v15 := []int{}
                v16 := GenS2{F0: v1, F1: (v6 + v6), F2: "hello"}
                var v17 float32 = ((1e10 + (1.5 - 1e10)) - 0.0)
                v18 := GenS1{F0: ("a" != ((((("foo bar" + "") + (v6 + "\n\t")) + v6) + v6) + "foo bar")), F1: (v3 | v3)}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            }
            switch v6 {
            case "foo bar":
                if true {
                    v19 := GenS2{F0: (^(v1)), F1: v6, F2: "hello"}
                    _ = v19
                }
            case "a":
                var v20 int = 520
                _ = v20
            case "\n\t":
                switch (fi12 << ((v3) & 63)) {
                case 0:
                    v11 -= (v3 >> (((v3 - v3)) & 63))
                }
            }
            _ = fi12
        }
        switch (v6 + v6) {
        case "a":
            if false {
                fmt.Println(v1)
                v6 = v6
                v3 += (v3 >> (((v11 + v11)) & 63))
            } else {
                var v21 int8 = v1
                _ = v21
            }
        case "hello":
            var v22 int32 = 2147483647
            _ = v22
        default:
            for fi23 := 0; fi23 < 11; fi23++ {
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 float64 = 1e10
                    _ = len(v2)
                    v26 := []int{fi24, (^(fi24)), fi24}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        }
        _ = v11
    } else {
        v1 = (v1 << (((^(v3))) & 63))
    }
    if (v5 <= 255) {
        var v27 uint = (v3 | v3)
        var v28 float32 = 1.5
        v28 = (v28 / (1e10 / 1.0))
        _ = v27
        _ = v28
    }
    var v29 uint = (v3 * v3)
    v1 += (v1 + v1)
    v30 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
