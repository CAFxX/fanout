package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 32768
    var v3 int = 127
    v4 := make(map[string]int)
    var v5 uint16 = 32767
    v6 := []string{(("a" + (("foo bar" + "a") + "")) + "")}
    delete(v4, "")
    fmt.Println(v2, v5)
    var v7 int8 = -128
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 562, F1: 65535}
        _ = fi8
        _ = v9
    }
    v6 = append(v6, "\n\t")
    var v10 uint16 = 2
    if (false || (978 == 0)) {
        v11 := GenS1{F0: 2147483647, F1: 4294967295}
        switch (v2 / (((v3 & (v2 ^ -836))) | 1)) {
        case -517:
            var v12 float32 = ((3.14159 - (0.0 * float32(3.14159))) / 1e10)
            _ = v12
        case -2147483648:
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int = (v2 ^ v3)
                _ = fi13
                _ = v14
            }
        case 0:
            v3 = v2
        }
        if len(v6) > 0 {
            v6[0] = ""
        }
        _ = v11
    }
    v3 = (v2 & 66)
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            switch 42 {
            case 128:
                v2 += (v2 + int(v10))
            case 9223372036854775807:
                v17 := GenS2{F0: 329, F1: 0}
                _ = v17
            case 32767:
                if (459 > v10) {
                    fmt.Println(v3)
                    v6 = append(v6, ("" + ("a" + ("hello" + ("\n\t" + ("" + (("hello" + "\n\t") + "\n\t")))))))
                }
                var v18 int16 = 42
                switch "hello" {
                case "":
                    fmt.Println(v3, v18, v2)
                case "hello":
                    var v19 int = (fi15 << ((42) & 63))
                    var v20 int32 = -558
                    if true {
                        v5 -= (v10 - 696)
                    } else {
                        switch fi15 {
                        case 1:
                            v3 -= (v2 >> ((1) & 63))
                        }
                        v21 := GenS1{F0: 42, F1: 2147483647}
                        _ = v21
                    }
                    switch "a" {
                    case "hello":
                        var v22 uint64 = 1
                        _ = v22
                    case "\n\t":
                        v19 = int(v10)
                    }
                    if false {
                        fmt.Println(v7, v2, v19)
                        v23 := GenS2{F0: 4294967295, F1: 18446744073709551615}
                        _ = v23
                    }
                    if true {
                        var v24 float32 = 3.14159
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v25 int8 = int8(v19)
                        _ = v24
                        _ = v25
                    } else {
                        var v26 float32 = 0.0
                        _ = v26
                    }
                    _ = v19
                    _ = v20
                }
                _ = v18
            }
            v7 -= v7
            var v27 float64 = 0.1
            _ = fi16
            _ = v27
        }
        fmt.Println(v7, v2)
        fmt.Println(v7, v7)
        _ = fi15
    }
    v10 = (v10 ^ 251)
    v2 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
