package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 string }
func genFn0() uint8 {
    var v5 float64 = 3.14159
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 = v5
        switch -45 {
        case 128:
            var v7 int16 = -1
            _ = v7
        case -503:
            v8 := GenS1{F0: 1}
            _ = v8
        }
        v5 = (v5 - v5)
        _ = fi6
    }
    _ = v5
    return 255
}
func genFn1() uint8 {
    switch "\n\t" {
    case "\n\t":
        var v9 int8 = 1
        _ = v9
    case "foo bar":
        var v10 string = "foo bar"
        _ = v10
    case "":
        v11 := GenS1{F0: 0}
        _ = v11
    }
    if false {
        var v12 int64 = 444
        v13 := GenS1{F0: -436}
        v14 := GenS2{F0: 128}
        _ = v12
        _ = v13
        _ = v14
    }
    return 128
}

func main() {
    var v1 int16 = 1
    var v2 int = 382
    v3 := []string{"hello", "foo bar", "\n\t", "foo bar"}
    v4 := []int{v2, v2, (v2 % (((v2 & (^(v2)))) | 1))}
    v15 := GenS3{F0: v2, F1: true, F2: "foo bar"}
    v2 -= -327
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint16 = 65535
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int16 = (v1 - (-(v1)))
        if len(v4) > 0 {
            v4[0] = (v2 / (((fi18 & fi18)) | 1))
        }
        if len(v4) > 0 {
            v4[0] = (fi18 - (^(fi18)))
        }
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 131}
    var v21 int64 = -1
    for fi22 := 0; fi22 < 17; fi22++ {
        if len(v4) > 0 {
            v4[0] = -1
        }
        v3 = append(v3, "a")
        _ = fi22
    }
    if true {
        if (!((v1 / ((v1) | 1)) > v1)) {
            if true {
                fmt.Println(v2)
                fmt.Println(v21, v2)
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v1 = (v1 >> ((255) & 63))
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(v2, v1)
            v4 = append(v4, fi23)
            _ = fi23
        }
    }
    v21 = (v21 >> ((65535) & 63))
    if ((-2.25 < (0.1 / 3.14159)) || false) {
        fmt.Println(v21, v21)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
