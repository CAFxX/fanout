package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 bool = (!false)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    var v5 uint32 = 65
    v6 := []string{"\n\t", "", (("" + "foo bar") + "\n\t"), ""}
    var v7 float32 = 3.14159
    v1 = (378 > 42)
    if (!(v1 || (0 >= 0))) {
        v5 -= (v5 / (((v5 % (((^(v5))) | 1))) | 1))
        if true {
            if v1 {
                v1 = (-553 == 2147483647)
                if (-144 < 1000) {
                    _ = len(v4)
                    var v8 uint32 = v5
                    _ = v8
                }
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 int64 = 517
                    var v11 bool = false
                    fmt.Println(v10, fi9)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            } else {
                if (!(!v1)) {
                    var v12 uint8 = 128
                    if false {
                        v3[v5] = 9223372036854775807
                        v13 := GenS1{F0: (v12 ^ (v12 + 0)), F1: -432}
                        v1 = v1
                        _ = v13
                    } else {
                        var v14 int64 = 9223372036854775807
                        _ = v14
                    }
                    switch "" {
                    case "foo bar":
                        var v15 int = 2147483647
                        _ = v15
                    }
                    _ = v12
                }
                var v16 int64 = -916
                _ = v16
            }
            var v17 uint32 = 780
            _ = v17
        }
    }
    if true {
        switch 9223372036854775807 {
        case -1:
            for fi18 := 0; fi18 < 11; fi18++ {
                if false {
                    var v19 bool = v1
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v6 = append(v6, ("foo bar" + "foo bar"))
                        v7 = float32(float64(v7))
                        _ = len(v6)
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 uint16 = 255
                        _ = fi21
                        _ = v22
                    }
                    _ = v19
                }
                _ = fi18
            }
        case 42:
            for fi23 := 0; fi23 < 17; fi23++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v24 int = 32767
                _ = fi23
                _ = v24
            }
        }
        v1 = false
    } else {
        switch "\n\t" {
        case "a":
            v5 = v5
        case "":
            fmt.Println(v5, v1)
        case "hello":
            v5 = (v5 >> ((902) & 63))
        }
        fmt.Println(v5, v5, v1)
    }
    var v25 float32 = (-2.25 - float32((0.0 - -2.25)))
    var v26 int32 = 903
    v27 := GenS1{F0: 255, F1: 0}
    var v28 int64 = 307
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
