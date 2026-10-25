package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    v1 := []string{"", "hello", ("\n\t" + ("\n\t" + "hello")), (("\n\t" + "") + ("" + (((("\n\t" + ("\n\t" + "foo bar")) + "\n\t") + "a") + "foo bar")))}
    var v2 int8 = 100
    v3 := []string{}
    var v4 int16 = 1
    switch 127 {
    case 0:
        v4 = (v4 - v4)
    }
    var v5 bool = false
    v6 := GenS1{F0: 32767, F1: -1}
    var v7 uint = 484
    v8 := GenS2{F0: 4294967295}
    v4 = (v4 >> ((137) & 63))
    if v5 {
        v4 = (v4 >> (((v7 % ((v7) | 1))) & 63))
    } else {
        for fi9 := 0; fi9 < 13; fi9++ {
            if (1 < 292) {
                if v5 {
                    var v10 uint = (v7 >> ((208) & 63))
                    for fi11 := 0; fi11 < 4; fi11++ {
                        if v5 {
                            if (!false) {
                                v12 := GenS3{F0: 0, F1: int64(fi11), F2: 127}
                                var v13 int16 = (v4 / ((v4) | 1))
                                _ = v12
                                _ = v13
                            } else {
                                v4 += v4
                                _ = len(v1)
                                var v14 int64 = 414
                                _ = v14
                            }
                            var v15 uint32 = 0
                            _ = len(v3)
                            _ = v15
                        }
                        v5 = v5
                        _ = fi11
                    }
                    v2 = 0
                    _ = v10
                } else {
                    v1 = append(v1, "foo bar")
                    if (-128 == v2) {
                        if ("\n\t" <= "a") {
                            v4 = (v4 * -32768)
                        } else {
                            var v16 int = 32767
                            fmt.Println(v4, fi9)
                            _ = v16
                        }
                    }
                }
                fmt.Println(v2)
                v5 = false
            }
            fmt.Println(v2, v5, v2)
            switch fi9 {
            case -9223372036854775808:
                v2 += 1
            case -549:
                v17 := GenS4{F0: 42, F1: 0, F2: v5}
                _ = v17
            default:
                fmt.Println(fi9, v5)
            }
            _ = fi9
        }
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    var v18 uint32 = 0
    for fi19 := 0; fi19 < 10; fi19++ {
        switch "" {
        case "foo bar":
            v20 := GenS3{F0: 0, F1: 42, F2: 255}
            _ = v20
        }
        v7 -= uint(v18)
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
