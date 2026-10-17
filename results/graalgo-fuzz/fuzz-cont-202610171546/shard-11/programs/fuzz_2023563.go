package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0() int16 {
    if true {
        var v5 int64 = -1
        var v6 int32 = 232
        v5 = (v5 << ((65535) & 63))
        _ = v5
        _ = v6
    }
    var v7 uint = 255
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ("" + "\n\t")
    var v3 uint = 686
    v4 := make(map[uint32]bool)
    v2 = "hello"
    v8 := []int{2147483647, 42, 2147483647, 395}
    switch -457 {
    case 1:
        v9 := GenS1{F0: 419, F1: ((("" + (v2 + (v2 + (v2 + ("\n\t" + v2))))) + ((v2 + v2) + (v2 + (("a" + (v2 + ((((v2 + "hello") + v2) + (v2 + "\n\t")) + ("" + v2)))) + v2)))) + ("\n\t" + "a"))}
        _ = v9
    default:
        _ = len(v8)
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        if (!true) {
            fmt.Println(v3, fi10)
        } else {
            v11 := GenS1{F0: (fi10 << ((4294967295) & 63)), F1: v2}
            if (42 > 42) {
                v12 := v4[42]
                switch fi10 {
                case -2147483648:
                    v12 = ("hello" != ((v2 + (((v2 + v2) + v2) + "a")) + ("\n\t" + ("\n\t" + v2))))
                case 9223372036854775807:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case 32767:
                    var v13 uint = 1
                    _ = v13
                }
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v12)
                case "":
                    fmt.Println(fi10, v3)
                default:
                    var v14 uint16 = 32767
                    _ = v14
                }
                _ = v12
            }
            _ = v11
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            switch (^(fi10)) {
            case -292:
                if true {
                    switch 234 {
                    case 1:
                        v16 := GenS1{F0: fi10, F1: (v2 + v2)}
                        _ = v16
                    case 0:
                        v17 := GenS1{F0: (fi10 % ((fi15) | 1)), F1: ((v2 + v2) + v2)}
                        _ = v17
                    }
                    var v18 uint16 = 255
                    _ = v18
                } else {
                    var v19 float32 = 0.1
                    _ = v19
                }
            case 0:
                v20 := v4[42]
                _ = v20
            }
            _ = fi15
        }
        var v21 int8 = 106
        _ = fi10
        _ = v21
    }
    var v22 int = 0
    v2 = v2
    v2 = (v2 + v2)
    var v23 string = ""
    var v24 int8 = int8(v22)
    for fi25 := 0; fi25 < 20; fi25++ {
        fmt.Println(v24, fi25, v23)
        if false {
            var v26 bool = true
            v24 -= v24
            _ = v26
        } else {
            var v27 int32 = -162
            _ = v27
        }
        _ = fi25
    }
    switch v2 {
    case "foo bar":
        var v28 float64 = (1e10 - -2.25)
        _ = v28
    case "hello":
        var v29 int16 = 32767
        _ = v29
    }
    var v30 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
