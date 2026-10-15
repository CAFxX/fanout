package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
func genFn0() uint {
    v7 := GenS1{F0: 300, F1: -32768, F2: 322}
    var v8 uint16 = 1
    var v9 int16 = -32768
    fmt.Println(v9, v9)
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 uint = 4294967295
    var v4 int16 = 0
    v5 := []uint32{}
    var v6 bool = (-763 != -1)
    v4 = v4
    var v10 bool = (127 == 50)
    if v6 {
        v11 := GenS2{F0: ""}
        fmt.Println(v6)
        switch ("a" + ("foo bar" + "")) {
        case "foo bar":
            if (!true) {
                for fi12 := 0; fi12 < 14; fi12++ {
                    for fi13 := 0; fi13 < 6; fi13++ {
                        if ((v10 && (v10 || (!true))) || v6) {
                            var v14 uint32 = 65535
                            v15 := GenS2{F0: ("a" + "foo bar")}
                            v4 += 1
                            _ = v14
                            _ = v15
                        }
                        v3 -= (v3 ^ 64)
                        v3 = v3
                        _ = fi13
                    }
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 int32 = 42
                        var v18 float32 = ((-2.25 / (1.5 + 0.1)) * 1e10)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = fi12
                }
            } else {
                fmt.Println(v6)
                v4 += v4
                v10 = (!false)
            }
        }
        _ = v11
    } else {
        v19 := GenS1{F0: 9223372036854775807, F1: (v4 % (((v4 + v4)) | 1)), F2: 95}
        v3 += 2
        _ = v19
    }
    v3 = (v3 >> ((v3) & 63))
    fmt.Println(v4)
    var v20 uint8 = 2
    v21 := GenS3{F0: 11}
    fmt.Println(v4)
    fmt.Println(v20, v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
