package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 bool; F2 bool }

func main() {
    var v1 int16 = -388
    var v2 uint8 = 255
    v3 := []string{(("" + "foo bar") + ("foo bar" + "a"))}
    v4 := []string{(("" + ("foo bar" + "a")) + ("a" + ("hello" + "foo bar")))}
    v5 := make(map[uint32]uint32)
    delete(v5, 42)
    for fi6 := 0; fi6 < 20; fi6++ {
        v2 = (^(v2))
        delete(v5, 110)
        _ = fi6
    }
    delete(v5, 1)
    v1 += (v1 - v1)
    v2 += 42
    v7 := GenS1{F0: 209}
    fmt.Println(v1)
    v5[42] = 894
    v1 = (v1 - 0)
    var v8 uint32 = 65535
    var v9 uint64 = 0
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    var v10 uint64 = (^(v9))
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := v5[v8]
        v9 = (v10 + 825)
        if true {
            var v13 int32 = -493
            v14 := GenS2{F0: (v13 - v13)}
            v15 := GenS2{F0: (v13 / ((v13) | 1))}
            _ = v13
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 int16 = (v1 / ((-780) | 1))
                v8 -= 42
                _ = fi16
                _ = v17
            }
            var v18 string = "a"
            if false {
                if ((v9 > uint64(v12)) && ((42 < 355) || (-1 <= 42))) {
                    switch (((("" + v18) + "hello") + v18) + ((v18 + "foo bar") + "")) {
                    case "a":
                        v4 = append(v4, v18)
                    case "foo bar":
                        v8 = 4294967295
                    }
                    var v19 uint8 = 157
                    _ = v19
                }
            } else {
                fmt.Println(v9, v9)
            }
            _ = v18
        }
        _ = fi11
        _ = v12
    }
    switch 981 {
    case 32768:
        v8 += 0
    case 9223372036854775807:
        v20 := GenS3{F0: 0, F1: (uint32(v2) <= (v8 / ((v8) | 1))), F2: true}
        _ = v20
    default:
        var v21 uint16 = 559
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
