package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 string; F2 uint8 }

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 42
    v3 := []uint32{58, 305}
    v4 := make(map[int]uint32)
    v5 := []string{}
    v4[-1] = 42
    var v6 uint64 = 4294967295
    v6 += (v2 ^ (v6 - v2))
    if (65535 < 703) {
        v3 = append(v3, 50)
        fmt.Println(v2, v1, v2)
        switch ("" + "foo bar") {
        case "\n\t":
            var v7 uint32 = 0
            _ = v7
        case "":
            _ = len(v3)
        case "foo bar":
            fmt.Println(v6, v2)
        }
    }
    v8 := GenS1{F0: 9223372036854775807}
    var v9 bool = ("" < (("foo bar" + (((((("foo bar" + ("" + ("\n\t" + ("hello" + "foo bar")))) + ("" + ("\n\t" + ("foo bar" + "hello")))) + "foo bar") + ("\n\t" + (("\n\t" + ("a" + "")) + ""))) + "foo bar") + ((("foo bar" + "a") + "foo bar") + "hello"))) + "hello"))
    delete(v4, 32767)
    fmt.Println(v1)
    v10 := GenS1{F0: -1}
    v2 -= (v2 / ((908) | 1))
    for fi11 := 0; fi11 < 5; fi11++ {
        switch fi11 {
        case 128:
            fmt.Println(v2, v2, v2)
        case 0:
            v4[fi11] = 65535
        case 4294967295:
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS2{F0: -650}
                var v14 uint8 = 79
                if (-32768 == 311) {
                    v3 = append(v3, 4294967295)
                    fmt.Println(v1, v9)
                    v15 := GenS1{F0: int(v14)}
                    _ = v15
                } else {
                    if true {
                        v6 = v2
                    }
                    if v9 {
                        if v9 {
                            v16 := GenS1{F0: 1}
                            _ = v16
                        }
                        v17 := GenS3{F0: "\n\t", F1: "", F2: v14}
                        _ = v17
                    }
                    v6 = 4294967295
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi11
    }
    delete(v4, -674)
    v1 = (v1 | (^(v1)))
    if v9 {
        v6 = (v6 << ((296) & 63))
    } else {
        if v9 {
            v6 = v6
            if (18446744073709551615 == (v2 / ((v6) | 1))) {
                v18 := GenS2{F0: 612}
                _ = v18
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            v19 := v4[1]
            _ = v19
        }
    }
    fmt.Println(v2, v1)
    var v20 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
