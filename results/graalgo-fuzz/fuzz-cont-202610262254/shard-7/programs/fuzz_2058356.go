package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
type GenS3 struct { F0 int32; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 65535
    var v3 uint8 = 255
    if (65535 <= 255) {
        for fi4 := 0; fi4 < 3; fi4++ {
            fmt.Println(v2, v3, v2)
            _ = fi4
        }
        switch 32768 {
        case 570:
            switch "" {
            case "":
                if true {
                    v3 += 9
                }
            case "foo bar":
                for fi5 := 0; fi5 < 4; fi5++ {
                    var v6 int32 = -528
                    if false {
                        delete(v1, fi5)
                        switch (("" + "hello") + "\n\t") {
                        case "foo bar":
                            v7 := GenS1{F0: (("\n\t" + ("foo bar" + "")) + "\n\t"), F1: 946}
                            _ = v7
                        case "hello":
                            if true {
                                fmt.Println(v3)
                                var v8 bool = ((fi5 | fi5) >= fi5)
                                _ = v8
                            }
                        case "":
                            v6 = v6
                        default:
                            v9 := GenS1{F0: "\n\t", F1: (fi5 + 2147483647)}
                            _ = v9
                        }
                        var v10 uint = 0
                        _ = v10
                    }
                    _ = fi5
                    _ = v6
                }
            case "hello":
                v2 += (^(v2))
            default:
                var v11 uint32 = 2
                _ = v11
            }
        case 128:
            if false {
                var v12 int64 = 0
                _ = v12
            } else {
                var v13 uint64 = 364
                v14 := GenS1{F0: "a", F1: 1}
                v13 -= (^(v13))
                _ = v13
                _ = v14
            }
        default:
            _ = len(v1)
        }
    }
    var v15 int8 = -128
    var v16 int = 127
    v17 := GenS2{F0: int8(v16), F1: 255, F2: v16}
    v3 += 189
    if true {
        v18 := GenS1{F0: "\n\t", F1: (v16 & 32768)}
        _ = v18
    }
    if ((((v3 & 42) < v3) || false) && (true || (true || (true && (("\n\t" + "") >= "foo bar"))))) {
        fmt.Println(v2, v3)
        v2 += 18446744073709551615
    }
    switch (v16 | v16) {
    case -167:
        v19 := GenS3{F0: 0, F1: (v3 / ((v3) | 1)), F2: v3}
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
