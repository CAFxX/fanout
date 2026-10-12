package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 bool; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{65535}
    var v3 uint = 164
    if true {
        _ = len(v2)
        fmt.Println(v3)
    }
    v3 = (v3 >> ((v3) & 63))
    var v4 int64 = -1
    fmt.Println(v4, v3)
    if true {
        v3 += 707
        fmt.Println(v4, v4)
        fmt.Println(v3, v3)
    } else {
        if false {
            switch (("" + "foo bar") + "hello") {
            case "a":
                v5 := GenS1{F0: (v3 >> ((v3) & 63)), F1: 2147483647}
                _ = v5
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                for fi6 := 0; fi6 < 1; fi6++ {
                    if (!((130 == 65535) || (!(((2147483647 > 2147483647) || false) || false)))) {
                        v3 = 341
                        v3 = (v3 + v3)
                    } else {
                        v7 := GenS2{F0: 255}
                        _ = v7
                    }
                    v4 -= v4
                    _ = fi6
                }
            default:
                v4 += (v4 - v4)
            }
        }
    }
    var v8 int = 0
    v8 += 764
    v9 := GenS3{F0: 1}
    v10 := GenS3{F0: v4}
    v11 := GenS3{F0: v4}
    var v12 int = 779
    v13 := GenS4{F0: 128, F1: (42 > (v3 ^ 400)), F2: -1}
    for fi14 := 0; fi14 < 15; fi14++ {
        delete(v1, "hello")
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int32 = 460
            _ = len(v1)
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
