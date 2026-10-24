package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint8 = 183
    v2 := []int{-2147483648, 42, 0}
    var v3 uint = 501
    v4 := []uint32{1, 0, 2147483647, 165}
    var v5 uint64 = 42
    var v6 uint32 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 float32 = 1e10
    switch "a" {
    case "foo bar":
        switch ((("hello" + ("hello" + "foo bar")) + "hello") + "\n\t") {
        case "\n\t":
            v6 = v6
        case "":
            var v8 uint16 = 255
            _ = v8
        case "hello":
            v9 := GenS1{F0: 1}
            _ = v9
        default:
            for fi10 := 0; fi10 < 3; fi10++ {
                if (!true) {
                    v3 -= v3
                    var v11 uint64 = v5
                    if true {
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 bool = (v6 < (v6 - v6))
                            var v14 int = 32767
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    }
                    _ = v11
                } else {
                    fmt.Println(v1)
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v6 = v6
                        _ = fi15
                    }
                }
                _ = fi10
            }
        }
    case "":
        fmt.Println(v6)
    }
    if (127 <= 59) {
        switch "" {
        case "":
            for fi16 := 0; fi16 < 7; fi16++ {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 uint = (v3 << (((v3 & (v3 ^ v3))) & 63))
                    if true {
                        v7 += v7
                    }
                    _ = fi17
                    _ = v18
                }
                switch fi16 {
                case -2147483648:
                    v5 = v5
                }
                var v19 uint8 = 127
                _ = fi16
                _ = v19
            }
        default:
            var v20 int16 = 32767
            _ = v20
        }
        v21 := GenS2{F0: 0, F1: (v5 << ((184) & 63))}
        var v22 uint64 = (^(v5))
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: -1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v1 = (v1 & (v1 | v1))
        var v25 uint64 = v5
        _ = fi24
        _ = v25
    }
    v7 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
