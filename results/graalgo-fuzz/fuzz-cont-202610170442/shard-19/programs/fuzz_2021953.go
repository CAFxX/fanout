package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{0, 2147483648}
    v3 := []uint32{144, 816}
    var v4 uint = 1
    if ((v4 % ((v4) | 1)) <= (v4 >> ((488) & 63))) {
        switch -1 {
        case 32767:
            var v5 uint32 = 0
            _ = v5
        case 313:
            var v6 int16 = 1
            _ = v6
        }
        switch "" {
        case "hello":
            v7 := GenS1{F0: v4, F1: -32768}
            _ = v7
        }
        v4 = (v4 + 521)
    } else {
        if (2 > 2147483648) {
            v4 = v4
        } else {
            if true {
                if ((!(("a" + ("foo bar" + "\n\t")) == ("hello" + "hello"))) && (!(127 <= 128))) {
                    v4 += v4
                    v4 -= 742
                }
            } else {
                v4 = 255
            }
        }
    }
    v2 = append(v2, 609)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 string = (("" + "hello") + "foo bar")
        _ = fi8
        _ = v9
    }
    fmt.Println(v4)
    if false {
        var v10 bool = (0 <= 128)
        var v11 int32 = -928
        _ = v10
        _ = v11
    }
    v4 = (^(v4))
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint32 = 0
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: (^(v4)), F1: 1}
    v1[2] = "\n\t"
    v4 = (v4 >> ((42) & 63))
    fmt.Println(v4)
    var v15 float32 = 0.1
    v2 = append(v2, 2147483647)
    v16 := GenS1{F0: v4, F1: -811}
    v3 = append(v3, 2147483647)
    fmt.Println(v4, v15)
    if true {
        switch "hello" {
        case "foo bar":
            if (false || false) {
                var v17 string = "hello"
                _ = v17
            } else {
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 int64 = 42
                    _ = fi18
                    _ = v19
                }
            }
        case "hello":
            v20 := GenS2{F0: 65535, F1: 2, F2: 92}
            _ = v20
        case "":
            if true {
                v4 += 42
            }
        }
        var v21 int16 = 1000
        var v22 bool = false
        _ = v21
        _ = v22
    }
    switch 127 {
    case -326:
        v23 := GenS1{F0: (v4 & (v4 << ((v4) & 63))), F1: 648}
        _ = v23
    case 940:
        v4 += v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
