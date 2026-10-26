package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]string)
    v3 := []uint32{76, 65535}
    var v4 string = v1
    if true {
        var v5 uint8 = 189
        v4 = v1
        var v6 uint8 = 146
        _ = v5
        _ = v6
    } else {
        if ((1 == 1) && (558 < 2)) {
            v1 = v1
        } else {
            fmt.Println(v1, v4)
            if (-105 >= 26) {
                var v7 uint = 606
                fmt.Println(v7)
                _ = v7
            } else {
                var v8 int32 = 725
                v2[-994] = v1
                _ = v8
            }
        }
    }
    if true {
        var v9 uint8 = 238
        delete(v2, -2147483648)
        v4 = "hello"
        _ = v9
    }
    v4 = ""
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: 0, F1: -216}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    delete(v2, -1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = ((v4 + v4) + "foo bar")
    var v13 bool = false
    var v14 uint8 = 255
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 uint32 = 0
            var v18 uint64 = 2
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 string = v1
            if v13 {
                var v21 float64 = (1.5 + ((-2.25 + ((0.1 - 0.1) / (0.1 * (0.0 - (-2.25 * (0.0 + ((-2.25 + 3.14159) + float64(0.0)))))))) / -2.25))
                v14 -= (v14 ^ v14)
                fmt.Println(v4, v1, v1)
                _ = v21
            }
            v13 = (65535 > 417)
            var v22 int = fi15
            _ = fi19
            _ = v20
            _ = v22
        }
        v23 := GenS2{F0: (790 != 32767), F1: 0}
        _ = fi15
        _ = v23
    }
    var v24 bool = (42 <= -1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
