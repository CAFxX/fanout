package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: -32768}
        _ = fi4
        _ = v5
    }
    fmt.Println(v1)
    v1 += v1
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint64 = 0
            _ = fi7
            _ = v8
        }
        if (!true) {
            var v9 uint32 = (v1 + 42)
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS2{F0: 1, F1: -501, F2: -1}
                var v12 uint32 = (^(v9))
                var v13 int16 = 32767
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            switch ("a" + "a") {
            case "foo bar":
                v14 := []int{fi6, 128}
                _ = v14
            case "hello":
                for fi15 := 0; fi15 < 14; fi15++ {
                    v9 = v1
                    var v16 float32 = -2.25
                    v17 := GenS2{F0: -1, F1: (fi15 / ((-923) | 1)), F2: 42}
                    if true {
                        v16 = (3.14159 - v16)
                        v18 := v2[""]
                        v19 := GenS2{F0: 42, F1: 32767, F2: -1}
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = v9
        } else {
            for fi20 := 0; fi20 < 1; fi20++ {
                switch "a" {
                case "a":
                    switch 42 {
                    case 191:
                        switch "\n\t" {
                        case "hello":
                            var v21 uint64 = 62
                            _ = v21
                        case "foo bar":
                            v1 += (v1 << ((0) & 63))
                        case "a":
                            v1 = (v1 >> ((18446744073709551615) & 63))
                        }
                        var v22 bool = (1 < 18446744073709551615)
                        _ = v22
                    }
                }
                var v23 int32 = 373
                _ = fi20
                _ = v23
            }
            fmt.Println(v1, v1)
        }
        delete(v3, "hello")
        _ = fi6
    }
    v24 := v3["\n\t"]
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 float64 = (((float64((((float32(0.1) * (0.0 - 1.5)) - (0.0 / 1.5)) * 1.5)) - 3.14159) * 1e10) + 0.1)
        _ = fi25
        _ = v26
    }
    v24 = v24
    var v27 uint = 42
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS2{F0: -32768, F1: (fi28 ^ -993), F2: -597}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v27 = v27
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
}
