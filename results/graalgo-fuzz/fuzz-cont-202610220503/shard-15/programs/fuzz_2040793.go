package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    v1 := []string{"foo bar", "foo bar", "", "a"}
    var v2 int = 0
    var v3 string = "foo bar"
    v4 := []string{"\n\t"}
    var v5 int = (v2 << ((330) & 63))
    v2 += 32768
    var v6 uint8 = 0
    switch "\n\t" {
    case "":
        switch (v5 | (v5 / ((-786) | 1))) {
        case 0:
            var v7 string = "a"
            _ = v7
        case -177:
            if len(v1) > 0 {
                v1[0] = (((v3 + "\n\t") + v3) + (v3 + v3))
            }
        }
    case "\n\t":
        v8 := GenS1{F0: 1, F1: 944}
        _ = v8
    case "hello":
        if false {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 string = v3
                _ = fi9
                _ = v10
            }
        } else {
            v11 := GenS2{F0: 42}
            _ = v11
        }
    }
    switch ("hello" + ("a" + "hello")) {
    case "\n\t":
        var v12 int = v5
        _ = v12
    }
    switch (v5 % ((v2) | 1)) {
    case -439:
        var v13 int8 = 42
        _ = v13
    }
    v1 = append(v1, (((v3 + v3) + (("" + "\n\t") + v3)) + (v3 + v3)))
    var v14 uint64 = 0
    var v15 int = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 int64 = 42
            var v19 uint8 = 255
            if true {
                if ((v2 | fi16) > int(v19)) {
                    var v20 uint16 = 42
                    switch "foo bar" {
                    case "a":
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v21 uint32 = 2
                            _ = v21
                        } else {
                            v22 := GenS1{F0: (v14 - (v14 - (v14 / ((v14) | 1)))), F1: v20}
                            _ = v22
                        }
                    }
                    v23 := GenS3{F0: 732, F1: (0 != int16(v14))}
                    _ = v20
                    _ = v23
                }
                v14 = (v14 | uint64(v19))
            } else {
                var v24 float64 = (0.0 + (1.5 - -2.25))
                fmt.Println(v24)
                _ = v24
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    v2 += 127
    var v25 uint8 = 92
    if (42 >= -199) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi26 := 0; fi26 < 20; fi26++ {
            switch 42 {
            case 127:
                if false {
                    v14 = (v14 / (((v14 % ((2) | 1))) | 1))
                    var v27 int32 = 439
                    if len(v4) > 0 {
                        v4[0] = (v3 + "\n\t")
                    }
                    _ = v27
                } else {
                    v2 = v5
                }
            default:
                _ = len(v4)
            }
            if true {
                v6 -= v25
                v28 := GenS1{F0: 4294967295, F1: 32767}
                _ = v28
            }
            _ = fi26
        }
        var v29 int16 = -473
        _ = v29
    }
    v14 = (v14 + v14)
    for fi30 := 0; fi30 < 18; fi30++ {
        v2 += v15
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
