package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    var v5 uint64 = 233
    v6 := v3["\n\t"]
    for fi7 := 0; fi7 < 16; fi7++ {
        switch (("hello" + "\n\t") + ("a" + "a")) {
        case "\n\t":
            switch ("" + "hello") {
            case "":
                var v8 float32 = (3.14159 + ((1.5 - (3.14159 + 3.14159)) - float32(float64(-2.25))))
                _ = v8
            default:
                for fi9 := 0; fi9 < 6; fi9++ {
                    v5 -= (v5 >> ((255) & 63))
                    var v10 float32 = 1e10
                    v11 := v3[("foo bar" + "")]
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        case "foo bar":
            var v12 uint32 = (v6 >> ((65535) & 63))
            _ = v12
        default:
            fmt.Println(fi7, fi7)
        }
        switch ("\n\t" + "foo bar") {
        case "":
            v13 := GenS1{F0: 0}
            _ = v13
        default:
            if false {
                var v14 string = ""
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 uint8 = 42
                    if (false && (true || (1.5 != 1e10))) {
                        switch v14 {
                        case "foo bar":
                            v16 += v16
                        }
                        v17 := GenS1{F0: -1}
                        _ = v17
                    }
                    v18 := []int{fi7, -301, (fi15 << ((1) & 63))}
                    _ = fi15
                    _ = v16
                    _ = v18
                }
                _ = v14
            }
        }
        _ = fi7
    }
    v19 := []int{int(v5), 4294967295}
    _ = len(v19)
    v20 := GenS2{F0: "hello", F1: -485}
    v5 -= (^(v5))
    if true {
        v21 := GenS3{F0: 335, F1: -272, F2: (v6 / ((v6) | 1))}
        var v22 uint8 = 128
        fmt.Println(v1, v5)
        _ = v21
        _ = v22
    }
    v1 = v1
    if len(v19) > 0 {
        v19[0] = 4294967295
    }
    var v23 int16 = 42
    v4["foo bar"] = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
