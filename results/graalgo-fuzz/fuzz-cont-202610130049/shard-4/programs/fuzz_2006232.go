package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    v1 := []uint32{651}
    var v2 uint = 219
    var v3 uint16 = 32767
    if len(v1) > 0 {
        v1[0] = 1
    }
    var v4 string = "\n\t"
    var v5 uint = v2
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint16 = (v3 + v3)
        fmt.Println(v7, v3)
        _ = fi6
        _ = v7
    }
    var v8 string = ""
    switch v8 {
    case "a":
        for fi9 := 0; fi9 < 17; fi9++ {
            switch (v4 + ((v4 + "foo bar") + "\n\t")) {
            case "foo bar":
                var v10 uint32 = 65535
                _ = v10
            case "\n\t":
                v3 -= 417
            default:
                v11 := GenS1{F0: (v2 ^ v5), F1: -1}
                _ = v11
            }
            if ((1 == 0) || true) {
                v4 = ""
            } else {
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int32 = 0
                    _ = fi12
                    _ = v13
                }
                var v14 int = -637
                _ = v14
            }
            _ = fi9
        }
    case "hello":
        v2 += 2
    default:
        v5 = v2
    }
    v15 := GenS2{F0: 128, F1: ((("" + v8) + v8) + "foo bar")}
    v2 = (v5 * v5)
    if (true && (!(v4 <= v4))) {
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(fi16, v8)
            if len(v1) > 0 {
                v1[0] = 1
            }
            _ = fi16
        }
    }
    var v17 int32 = -1
    _ = len(v1)
    switch -652 {
    case 32767:
        var v18 uint16 = v3
        _ = v18
    case 450:
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v8, v4)
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v20 uint32 = 0
                _ = v20
            } else {
                var v21 bool = true
                _ = v21
            }
            _ = len(v1)
            _ = fi19
        }
    default:
        for fi22 := 0; fi22 < 5; fi22++ {
            v8 = (v4 + v8)
            v4 = (("a" + v4) + (v8 + (("foo bar" + v8) + "a")))
            v5 += 923
            _ = fi22
        }
    }
    if ((0 <= 0) && false) {
        switch v8 {
        case "foo bar":
            v8 = ("\n\t" + v4)
        }
    }
    v1 = append(v1, 2)
    fmt.Println(v8, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
