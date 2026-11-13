package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    var v3 uint64 = 4294967295
    var v4 uint = 42
    v1 = "foo bar"
    for fi5 := 0; fi5 < 6; fi5++ {
        if true {
            v6 := []int{(fi5 - fi5)}
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 int16 = 423
                switch 306 {
                case 42:
                    fmt.Println(fi7, fi7)
                case -369:
                    v9 := GenS1{F0: 32767, F1: 0}
                    _ = v9
                case 846:
                    var v10 uint8 = 42
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
            delete(v2, 688)
            _ = v6
        }
        _ = fi5
    }
    v4 += (v4 >> (((v4 | (^(v4)))) & 63))
    switch 9223372036854775807 {
    case 24:
        v11 := GenS1{F0: 32767, F1: 65535}
        _ = v11
    case 32768:
        v3 += 689
    case 957:
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                v1 = v1
                var v14 uint8 = 32
                _ = fi13
                _ = v14
            }
            v1 = v1
            _ = fi12
        }
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        if (false && (false || (0.1 != 1.5))) {
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS1{F0: -93, F1: 2}
                _ = fi16
                _ = v17
            }
            if false {
                v18 := GenS1{F0: 588, F1: 32767}
                v19 := GenS2{F0: 128, F1: (fi15 << (((v4 ^ v4)) & 63)), F2: 550}
                v20 := GenS2{F0: 2, F1: 0, F2: 42}
                _ = v18
                _ = v19
                _ = v20
            } else {
                var v21 int8 = 100
                var v22 uint16 = 255
                _ = v21
                _ = v22
            }
        }
        var v23 int = (fi15 % (((-(fi15))) | 1))
        v24 := v2[0]
        _ = fi15
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 1, F1: 255}
    _ = len(v2)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
