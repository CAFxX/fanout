package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint }
type GenS3 struct { F0 int32; F1 uint32 }

func main() {
    var v1 string = (("\n\t" + "foo bar") + ("a" + "\n\t"))
    var v2 uint32 = 1
    v3 := make(map[string]int)
    var v4 float64 = 0.0
    switch v1 {
    case "foo bar":
        v1 = ("" + v1)
    case "":
        var v5 uint32 = (v2 % (((v2 >> ((42) & 63))) | 1))
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 10; fi6++ {
            v1 = ""
            var v7 uint8 = 212
            v7 -= 127
            _ = fi6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            _ = len(v3)
            fmt.Println(fi8, v2)
            var v9 int32 = -1
            _ = v9
        }
        var v10 uint16 = 32767
        v11 := GenS1{F0: 4294967295, F1: 255}
        _ = fi8
        _ = v10
        _ = v11
    }
    v12 := []int{1, 32768, -2147483648}
    fmt.Println(v4, v1)
    v13 := v3[v1]
    var v14 uint = 0
    if false {
        var v15 uint = 4294967295
        switch (-(v13)) {
        case 32767:
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS2{F0: 915, F1: v15, F2: v14}
                if true {
                    fmt.Println(v14, v4)
                }
                var v18 uint16 = uint16(v2)
                _ = fi16
                _ = v17
                _ = v18
            }
        case 1:
            fmt.Println(v14)
        case 9223372036854775807:
            switch (v13 | 121) {
            case -478:
                if (false && (float32(v4) == 1.5)) {
                    fmt.Println(v13, v14)
                    v14 = v14
                    var v19 float32 = 0.1
                    _ = v19
                } else {
                    fmt.Println(v4, v13)
                    switch int(v15) {
                    case 9223372036854775807:
                        fmt.Println(v15, v15, v2)
                    case -417:
                        var v20 int8 = -17
                        _ = v20
                    case -563:
                        v21 := v3["\n\t"]
                        _ = v21
                    }
                }
            case -1:
                switch 9223372036854775807 {
                case -364:
                    v22 := GenS3{F0: 2147483647, F1: v2}
                    _ = v22
                case 127:
                    fmt.Println(v2)
                }
            case 9223372036854775807:
                v1 = v1
            }
        }
        var v23 int64 = 0
        _ = v15
        _ = v23
    }
    var v24 int = (v13 | (v13 >> (((^(v14))) & 63)))
    var v25 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
