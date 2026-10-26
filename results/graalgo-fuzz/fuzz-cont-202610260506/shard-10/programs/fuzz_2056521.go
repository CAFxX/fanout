package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }

func main() {
    v1 := []string{((("hello" + "foo bar") + "hello") + ""), "\n\t"}
    var v2 uint = 0
    var v3 int16 = 0
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    var v6 int32 = 42
    var v7 int = 896
    if false {
        delete(v4, 1)
        for fi8 := 0; fi8 < 16; fi8++ {
            v3 -= -633
            _ = fi8
        }
        var v9 int = (v7 % ((v7) | 1))
        _ = v9
    }
    v10 := GenS1{F0: 2147483648}
    v2 = 629
    _ = len(v4)
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            switch 9223372036854775807 {
            case 1:
                var v13 uint = v2
                _ = v13
            default:
                v6 -= v6
            }
            _ = fi12
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v6 -= (v6 + 79)
        _ = fi14
    }
    v7 = v7
    for fi15 := 0; fi15 < 17; fi15++ {
        if false {
            var v16 bool = (1e10 <= float64(1e10))
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v7 = 2147483647
            v7 = (fi15 ^ (fi17 >> ((509) & 63)))
            var v18 uint32 = 65535
            _ = fi17
            _ = v18
        }
        _ = fi15
    }
    var v19 bool = (true && ((^(v2)) != v2))
    v20 := GenS1{F0: 465}
    var v21 int64 = 9223372036854775807
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := GenS2{F0: ("foo bar" + "\n\t")}
        if false {
            for fi24 := 0; fi24 < 10; fi24++ {
                if (2 >= 32) {
                    v19 = (!false)
                    var v25 bool = v19
                    _ = v25
                } else {
                    v26 := GenS2{F0: ""}
                    var v27 float32 = 0.1
                    if (v6 < v6) {
                        if ((v19 || (42 > 42)) && v19) {
                            var v28 bool = true
                            var v29 uint64 = 0
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS1{F0: 1}
                        _ = v30
                    } else {
                        v7 += -787
                    }
                    _ = v26
                    _ = v27
                }
                v2 = (v2 << (((v2 - (v2 << ((v2) & 63)))) & 63))
                v31 := GenS1{F0: 2147483648}
                _ = fi24
                _ = v31
            }
            v32 := GenS2{F0: ("a" + "hello")}
            _ = v32
        }
        _ = fi22
        _ = v23
    }
    var v33 int64 = 42
    if (!v19) {
        v34 := GenS3{F0: (^(v7))}
        var v35 int64 = (v21 >> ((v2) & 63))
        _ = v34
        _ = v35
    }
    var v36 string = ""
    switch v36 {
    case "":
        v37 := GenS1{F0: 65535}
        _ = v37
    default:
        var v38 uint16 = 520
        _ = v38
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
