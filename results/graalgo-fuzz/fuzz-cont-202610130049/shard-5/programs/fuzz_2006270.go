package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 string = "a"
    v2 := make(map[string]uint32)
    var v3 uint8 = 18
    v4 := []int{127, 42}
    v3 -= (v3 ^ (v3 | (v3 + (^(v3)))))
    if ((^(v3)) < (v3 & v3)) {
        fmt.Println(v3, v3)
    } else {
        v1 = "hello"
    }
    if (int32(v3) < 2147483647) {
        v1 = ("foo bar" + (("foo bar" + "hello") + "\n\t"))
        var v5 bool = false
        _ = v5
    } else {
        if (!(2 == 244)) {
            fmt.Println(v1, v3)
            var v6 uint8 = 15
            _ = v6
        }
        v7 := GenS1{F0: ("a" + "\n\t")}
        _ = v7
    }
    v8 := v2[(v1 + v1)]
    v9 := GenS2{F0: int8(v8), F1: 1}
    for fi10 := 0; fi10 < 10; fi10++ {
        if len(v4) > 0 {
            v4[0] = fi10
        }
        var v11 uint64 = 2
        _ = fi10
        _ = v11
    }
    v8 = (v8 * v8)
    v12 := GenS2{F0: 0, F1: 861}
    if (!(42 <= 2)) {
        if false {
            fmt.Println(v3, v3)
        } else {
            v3 = (v3 | 127)
            if false {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v1 = ((v1 + ((v1 + "hello") + v1)) + (v1 + v1))
                    var v14 string = v1
                    _ = fi13
                    _ = v14
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v17 := GenS1{F0: ""}
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
            }
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18, v3)
            v1 = (((v1 + "") + "a") + (v1 + v1))
            _ = fi18
        }
    }
    var v19 float64 = -2.25
    var v20 int32 = 2147483647
    switch (v1 + v1) {
    case "\n\t":
        v2[(v1 + v1)] = v8
    }
    var v21 int32 = v20
    fmt.Println(v3, v3)
    _ = len(v4)
    for fi22 := 0; fi22 < 16; fi22++ {
        v8 -= 1
        for fi23 := 0; fi23 < 13; fi23++ {
            switch v1 {
            case "":
                if true {
                    var v24 int = (fi23 >> ((65535) & 63))
                    fmt.Println(fi23, v19, v3)
                    _ = v24
                } else {
                    fmt.Println(v21, v19)
                    var v25 uint = 570
                    v8 += v8
                    var v26 int8 = -1
                    var v27 string = "foo bar"
                    v28 := GenS3{F0: v20}
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            default:
                v29 := GenS4{F0: v8, F1: 32767}
                _ = v29
            }
            fmt.Println(v3, v19, v8)
            var v30 uint = 765
            if true {
                var v31 int64 = 9223372036854775807
                _ = v31
            } else {
                switch (fi23 / (((fi22 % ((fi22) | 1))) | 1)) {
                case -41:
                    var v32 int64 = 1
                    _ = v32
                default:
                    var v33 int16 = -32768
                    _ = v33
                }
                v8 += v8
            }
            _ = fi23
            _ = v30
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
