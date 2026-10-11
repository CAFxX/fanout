package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint = 503
    var v2 uint = (v1 & (v1 << ((235) & 63)))
    v3 := []uint32{2147483648, 1, 0, 952}
    for fi4 := 0; fi4 < 10; fi4++ {
        v1 = v1
        fmt.Println(fi4)
        for fi5 := 0; fi5 < 1; fi5++ {
            switch (fi5 | 9223372036854775807) {
            case 42:
                var v6 uint8 = 42
                _ = v6
            case 842:
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            }
            fmt.Println(fi5)
            var v7 uint64 = 2
            _ = fi5
            _ = v7
        }
        _ = fi4
    }
    if (127 < -128) {
        v1 -= 1
        v2 = v2
        if true {
            var v8 bool = true
            _ = v8
        } else {
            var v9 float32 = 0.1
            _ = v9
        }
    }
    v2 = 65535
    switch "\n\t" {
    case "\n\t":
        v3 = append(v3, 982)
    default:
        fmt.Println(v1, v2)
    }
    v10 := GenS1{F0: 42, F1: -80}
    var v11 string = ("hello" + "a")
    fmt.Println(v2)
    v1 = (^(v1))
    for fi12 := 0; fi12 < 17; fi12++ {
        if true {
            if true {
                v13 := GenS2{F0: (v1 + v2), F1: 42}
                _ = v13
            } else {
                if false {
                    v3 = append(v3, 1)
                }
                v11 = (v11 + "a")
            }
            v1 -= (v1 << (((^(v1))) & 63))
            var v14 uint64 = 696
            _ = v14
        } else {
            for fi15 := 0; fi15 < 17; fi15++ {
                switch v11 {
                case "foo bar":
                    v16 := GenS1{F0: 411, F1: -128}
                    _ = v16
                case "":
                    switch v11 {
                    case "hello":
                        var v17 int32 = 686
                        _ = v17
                    case "\n\t":
                        switch fi12 {
                        case -380:
                            var v18 int64 = 42
                            _ = v18
                        }
                    case "foo bar":
                        v11 = (v11 + ("a" + v11))
                    default:
                        fmt.Println(v1, fi15)
                    }
                }
                var v19 int16 = 32767
                var v20 string = "a"
                _ = fi15
                _ = v19
                _ = v20
            }
            if true {
                v2 = v1
                _ = len(v3)
            } else {
                v11 = "hello"
            }
        }
        _ = fi12
    }
    var v21 int8 = -1
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS1{F0: 279, F1: 1}
        var v24 uint = v2
        switch "" {
        case "":
            var v25 int = -1
            _ = v25
        case "foo bar":
            switch (v11 + v11) {
            case "a":
                var v26 float32 = ((1e10 + float32(1.5)) - float32(float64(3.14159)))
                _ = v26
            case "hello":
                switch (fi22 / ((fi22) | 1)) {
                case -2147483648:
                    var v27 uint16 = 2
                    _ = v27
                case 2147483647:
                    v28 := GenS2{F0: 317, F1: 618}
                    _ = v28
                case 32767:
                    fmt.Println(v11)
                }
            default:
                fmt.Println(fi22, v11)
            }
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    var v29 uint64 = 1
    v21 -= v21
    if (!((v11 + (v11 + "\n\t")) != (((v11 + v11) + v11) + "hello"))) {
        if len(v3) > 0 {
            v3[0] = uint32(v1)
        }
    }
    v30 := GenS2{F0: 4294967295, F1: 18446744073709551615}
    v31 := GenS2{F0: v1, F1: v29}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
