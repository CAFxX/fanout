package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    v1 := []string{}
    v2 := []string{("a" + "a"), "\n\t"}
    var v3 int16 = -1
    if true {
        if false {
            if false {
                var v4 int = 0
                _ = v4
            } else {
                _ = len(v2)
            }
            for fi5 := 0; fi5 < 2; fi5++ {
                fmt.Println(fi5, fi5)
                var v6 int64 = -128
                _ = fi5
                _ = v6
            }
        } else {
            fmt.Println(v3)
            v3 += (v3 << ((255) & 63))
        }
    } else {
        v7 := GenS1{F0: "foo bar"}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v7
    }
    if (348 >= 255) {
        switch int(v3) {
        case 9223372036854775807:
            if (2 == 4294967295) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v8 bool = true
                _ = v8
            } else {
                var v9 int32 = 0
                var v10 int16 = (v3 << ((200) & 63))
                _ = v10
                _ = v9
            }
        case 127:
            v11 := GenS2{F0: (v3 | 26), F1: false}
            _ = v11
        }
        v3 -= 0
    } else {
        var v12 uint32 = 2147483648
        v13 := GenS3{F0: "hello", F1: 0}
        switch "hello" {
        case "foo bar":
            var v14 float64 = (1e10 + 1e10)
            _ = v14
        case "hello":
            var v15 int64 = 1
            _ = v15
        case "a":
            v3 = v3
        }
        _ = v12
        _ = v13
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v3 = v3
        switch "\n\t" {
        case "\n\t":
            v1 = append(v1, "foo bar")
        case "foo bar":
            var v17 string = "a"
            _ = v17
        default:
            var v18 uint8 = 2
            _ = v18
        }
        v3 += 0
        _ = fi16
    }
    var v19 int = -416
    fmt.Println(v19)
    var v20 bool = false
    fmt.Println(v19, v20, v19)
    v3 += (-(v3))
    var v21 uint8 = 88
    var v22 uint = 255
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
