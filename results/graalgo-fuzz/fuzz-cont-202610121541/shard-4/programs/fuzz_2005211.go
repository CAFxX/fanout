package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint8 = 2
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    v4 := []int{128, 127, 621}
    fmt.Println(v2, v1)
    var v5 uint8 = v1
    var v6 uint64 = 18446744073709551615
    v1 -= v1
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint32 = 4294967295
        fmt.Println(fi7, v8, v2)
        v9 := GenS1{F0: -1, F1: (("\n\t" + "a") + "hello")}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            if (0 != 42) {
                var v11 uint32 = uint32(v6)
                _ = v11
            } else {
                if len(v4) > 0 {
                    v4[0] = 761
                }
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 uint8 = (^(v1))
                    _ = fi12
                    _ = v13
                }
            }
            v6 -= v6
        } else {
            if (true || (-1 < -2147483648)) {
                v14 := GenS1{F0: -1, F1: "hello"}
                _ = v14
            }
        }
        v2 = (v2 & (v2 & v2))
        _ = fi10
    }
    var v15 uint32 = 527
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v6)
        for fi17 := 0; fi17 < 18; fi17++ {
            switch "" {
            case "a":
                if ((v1 != (v1 + 127)) && (32767 <= 0)) {
                    switch "" {
                    case "foo bar":
                        v18 := GenS2{F0: 255}
                        _ = v18
                    case "\n\t":
                        var v19 uint = 42
                        _ = v19
                    case "":
                        v15 += 2
                    }
                    switch fi16 {
                    case 196:
                        _ = len(v3)
                    case 202:
                        v20 := GenS3{F0: 9223372036854775807, F1: v5}
                        _ = v20
                    case 127:
                        v15 = v15
                        fmt.Println(v1, v1, v6)
                        var v21 int64 = 435
                        _ = v21
                    }
                }
            }
            var v22 string = ("hello" + "a")
            _ = fi17
            _ = v22
        }
        _ = fi16
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
