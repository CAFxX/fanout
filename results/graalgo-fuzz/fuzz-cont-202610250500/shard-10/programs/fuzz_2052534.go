package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }

func main() {
    var v1 int = 4294967295
    var v2 int8 = 100
    v3 := []string{"a"}
    var v4 float32 = 1e10
    v5 := make(map[string]string)
    v6 := make(map[uint32]bool)
    var v7 int = (v1 + (v1 >> ((4294967295) & 63)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS1{F0: (true || (!(true || (true || (!true))))), F1: -308}
    v9 := v5["hello"]
    v10 := GenS1{F0: false, F1: 376}
    v11 := GenS2{F0: 1000}
    if true {
        var v12 float64 = 3.14159
        v13 := GenS2{F0: -492}
        if (int64(v2) != 42) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 -= v2
        }
        _ = v12
        _ = v13
    }
    v2 += 42
    switch v9 {
    case "":
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS3{F0: 0}
            v7 += 32767
            v6[2] = ((v9 + v9) >= v9)
            _ = fi14
            _ = v15
        }
    case "foo bar":
        var v16 uint16 = 65535
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int32 = 0
            _ = fi17
            _ = v18
        }
    default:
        v7 = (v7 / (((^(v1))) | 1))
    }
    v5[""] = v9
    v9 = v9
    v19 := GenS4{F0: v9}
    switch (v1 % ((v1) | 1)) {
    case -905:
        if false {
            v3 = append(v3, "hello")
        } else {
            if (42 == 9223372036854775807) {
                if (true || false) {
                    v7 += -295
                } else {
                    if false {
                        var v20 float64 = 0.1
                        v6[65535] = (42 != -32768)
                        for fi21 := 0; fi21 < 4; fi21++ {
                            var v22 int64 = 655
                            v22 -= v22
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                }
                fmt.Println(v9, v9, v7)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
