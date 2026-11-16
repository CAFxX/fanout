package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 4294967295
    var v2 string = "hello"
    var v3 uint32 = 390
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 1; fi5++ {
        if ((v3 % (((v3 + v3)) | 1)) <= 2147483647) {
            v6 := GenS1{F0: 65535, F1: 766, F2: 597}
            _ = v6
        }
        var v7 uint32 = 42
        if (355 < 65535) {
            for fi8 := 0; fi8 < 16; fi8++ {
                if (291 >= 206) {
                    if (!(0 != v3)) {
                        var v9 uint32 = 199
                        if ((true && (42 <= -1)) || false) {
                            v10 := GenS2{F0: 1}
                            _ = v10
                        }
                        _ = v9
                    }
                }
                v1 = v1
                _ = fi8
            }
        }
        _ = fi5
        _ = v7
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v2, v2, v1)
        _ = fi11
    }
    v12 := GenS1{F0: 255, F1: 65535, F2: 664}
    v3 -= v3
    v2 = v2
    switch (((v2 + v2) + "hello") + "a") {
    case "\n\t":
        var v13 int16 = 32767
        _ = v13
    case "foo bar":
        if false {
            var v14 bool = true
            var v15 float32 = 0.1
            var v16 int32 = 1
            _ = v14
            _ = v15
            _ = v16
        }
    default:
        v1 = v1
    }
    fmt.Println(v3)
    v17 := GenS1{F0: 602, F1: 255, F2: 175}
    if false {
        var v18 int32 = 744
        delete(v4, 127)
        fmt.Println(v2, v2, v1)
        _ = v18
    } else {
        var v19 string = (v2 + ("\n\t" + v2))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
