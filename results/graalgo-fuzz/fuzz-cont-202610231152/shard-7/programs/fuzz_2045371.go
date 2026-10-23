package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint; F2 string }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{955}
    v3 := make(map[string]string)
    v4 := GenS1{F0: -874, F1: 42}
    for fi5 := 0; fi5 < 8; fi5++ {
        _ = len(v1)
        if true {
            fmt.Println(fi5)
            for fi6 := 0; fi6 < 19; fi6++ {
                v7 := GenS1{F0: 0, F1: -838}
                var v8 uint8 = 128
                v9 := GenS1{F0: 42, F1: 42}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 uint16 = 42
            _ = v10
        } else {
            v11 := GenS1{F0: -2147483648, F1: -601}
            if ((1.5 + (0.1 + float32((float64(float32(1.5)) * 0.1)))) >= (((0.1 / 3.14159) + 3.14159) + 1e10)) {
                delete(v3, "\n\t")
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 int64 = -323
                    var v14 int16 = 1000
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            fmt.Println(fi5)
            _ = v11
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint8 = 1
            delete(v3, "hello")
            v17 := GenS2{F0: 1, F1: true, F2: -1}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi5
    }
    delete(v3, "foo bar")
    var v18 int16 = -32768
    v18 += 717
    for fi19 := 0; fi19 < 11; fi19++ {
        v18 += (-(v18))
        _ = fi19
    }
    delete(v3, "\n\t")
    v20 := GenS1{F0: -2147483648, F1: 0}
    if false {
        if true {
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := GenS3{F0: 2147483648, F1: 889, F2: ""}
                _ = fi21
                _ = v22
            }
        }
        if ((true && (58 != 62)) || false) {
            v23 := v3[("\n\t" + "a")]
            v24 := GenS3{F0: 65535, F1: 65535, F2: "a"}
            _ = v23
            _ = v24
        } else {
            v18 = (v18 / ((v18) | 1))
        }
        v18 += (^(v18))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
