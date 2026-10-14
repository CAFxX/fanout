package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    var v1 string = "hello"
    var v2 int32 = 611
    var v3 string = ""
    v4 := []int{2147483647, -9223372036854775808, 42, 4294967295}
    fmt.Println(v2, v2)
    v3 = (("hello" + (v3 + (v3 + "a"))) + v1)
    if true {
        for fi5 := 0; fi5 < 5; fi5++ {
            if true {
                v4 = append(v4, (fi5 >> ((42) & 63)))
            } else {
                var v6 uint16 = 63
                var v7 int64 = 247
                _ = v6
                _ = v7
            }
            _ = fi5
        }
    } else {
        v3 = "\n\t"
        if true {
            v8 := GenS1{F0: 42, F1: -1}
            var v9 uint16 = 42
            switch (v3 + "\n\t") {
            case "":
                v10 := GenS2{F0: 0, F1: 2, F2: 975}
                _ = v10
            case "hello":
                for fi11 := 0; fi11 < 15; fi11++ {
                    v12 := GenS2{F0: 1, F1: 593, F2: 725}
                    var v13 int64 = 0
                    if len(v4) > 0 {
                        v4[0] = -822
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = v8
            _ = v9
        } else {
            fmt.Println(v1, v3, v3)
            v3 = v1
        }
        fmt.Println(v1, v1)
    }
    v14 := GenS1{F0: 65535, F1: -1}
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            v3 = ""
            v2 = (v2 % ((v2) | 1))
            _ = fi16
        }
        _ = fi15
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS1{F0: 255, F1: 42}
            _ = fi18
            _ = v19
        }
        var v20 uint = 538
        _ = fi17
        _ = v20
    }
    if false {
        v21 := GenS3{F0: 431, F1: v2}
        v1 = "hello"
        _ = v21
    } else {
        switch 456 {
        case 4294967295:
            v2 += (v2 * -775)
        case 128:
            v2 += (^(v2))
        default:
            v1 = ""
        }
        v4 = append(v4, -9223372036854775808)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
