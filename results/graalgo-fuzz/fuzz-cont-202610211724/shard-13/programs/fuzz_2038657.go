package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 bool = false
    v4 := make(map[int]bool)
    v5 := make(map[string]int)
    v6 := []uint32{4294967295, 1, 432}
    var v7 float32 = 3.14159
    v3 = (1 > 42)
    v8 := GenS1{F0: 1}
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(v7, fi10)
            _ = fi10
        }
        var v11 int64 = -9223372036854775808
        if ((false && (-128 <= -128)) && (v3 || true)) {
            switch (fi9 | (fi9 ^ fi9)) {
            case -268:
                var v12 float32 = ((v7 + 0.0) + (float32((((-2.25 + 3.14159) + 0.0) * 3.14159)) * 0.0))
                _ = v12
            case 4294967295:
                switch (fi9 + fi9) {
                case 32767:
                    v11 -= 137
                case 911:
                    switch 32767 {
                    case -2147483648:
                        v3 = false
                    case -990:
                        v13 := GenS1{F0: 280}
                        _ = v13
                    }
                case 0:
                    switch "\n\t" {
                    case "hello":
                        for fi14 := 0; fi14 < 14; fi14++ {
                            v15 := v5["a"]
                            _ = fi14
                            _ = v15
                        }
                    case "":
                        for fi16 := 0; fi16 < 18; fi16++ {
                            var v17 float32 = (v7 - float32(3.14159))
                            _ = fi16
                            _ = v17
                        }
                    }
                default:
                    v11 = (v11 - v11)
                }
            case -181:
                var v18 int64 = -9223372036854775808
                _ = v18
            }
            v19 := GenS1{F0: 65535}
            var v20 int = fi9
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 2147483648}
            _ = v21
        }
        _ = fi9
        _ = v11
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(fi22)
        v23 := GenS1{F0: 4294967295}
        if (v7 <= v7) {
            var v24 int16 = 1
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    v25 := GenS1{F0: 119}
    fmt.Println(v7, v3)
    fmt.Println(v7, v7)
    v26 := GenS2{F0: (v3 || (42 <= 42)), F1: -1, F2: 0}
    if (v3 || ((v7 + float32(-2.25)) >= v7)) {
        switch 630 {
        case 31:
            fmt.Println(v7, v3)
        }
        if v3 {
            if v3 {
                var v27 uint = 18446744073709551615
                _ = v27
            }
            fmt.Println(v7, v7)
            var v28 string = "\n\t"
            _ = v28
        }
    }
    if v3 {
        var v29 float64 = ((0.1 / float64((v7 / 1.5))) * 3.14159)
        _ = v29
    } else {
        switch 32767 {
        case -9223372036854775808:
            var v30 bool = ((!(!false)) || (v3 || (!v3)))
            _ = v30
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
