package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 string) string {
    v7 := GenS1{F0: 65535, F1: 42, F2: 9223372036854775807}
    v8 := []int{-202, 42}
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint8 = 128
        v11 := GenS1{F0: 255, F1: 0, F2: 0}
        v12 := GenS1{F0: uint(v10), F1: -22, F2: 9223372036854775807}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return ("\n\t" + (p0 + p0))
}

func main() {
    v1 := []int{}
    v2 := make(map[string]int)
    var v3 int64 = -140
    v4 := []uint32{948}
    var v5 int = -272
    var v6 float32 = 1.5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ((("hello" + ("hello" + "\n\t")) + "a") + (("hello" + "\n\t") + ("a" + "foo bar"))) {
    case "hello":
        fmt.Println(v3, v6)
    case "foo bar":
        fmt.Println(v6, v6)
    case "a":
        if true {
            v13 := GenS1{F0: 931, F1: 1, F2: v3}
            var v14 int = 1
            for fi15 := 0; fi15 < 7; fi15++ {
                v2[""] = (fi15 % ((fi15) | 1))
                _ = fi15
            }
            _ = v13
            _ = v14
        }
    default:
        v16 := GenS2{F0: 42, F1: 0}
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            switch -125 {
            case 4294967295:
                if ((v3 << ((133) & 63)) == v3) {
                    v19 := GenS1{F0: 42, F1: 127, F2: (v3 * -1)}
                    if true {
                        if false {
                            var v20 uint8 = 42
                            var v21 string = "\n\t"
                            switch (v21 + v21) {
                            case "":
                                var v22 float32 = 3.14159
                                _ = v22
                            }
                            _ = v20
                            _ = v21
                        }
                        var v23 int = (v5 | 4294967295)
                        fmt.Println(v3, fi17)
                        _ = v23
                    } else {
                        switch ("\n\t" + ("hello" + "hello")) {
                        case "\n\t":
                            var v24 uint32 = 42
                            _ = v24
                        }
                        v25 := GenS3{F0: (fi17 / ((fi18) | 1))}
                        _ = v25
                    }
                    fmt.Println(fi18, fi17)
                    _ = v19
                }
            case -2147483648:
                fmt.Println(v6, v3)
            }
            var v26 uint8 = 2
            _ = fi18
            _ = v26
        }
        _ = fi17
    }
    fmt.Println(v3, v5, v6)
    v27 := GenS1{F0: 1, F1: 127, F2: 696}
    fmt.Println(v5, v6, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
