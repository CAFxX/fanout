package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint; F2 int64 }
type GenS5 struct { F0 string; F1 int }
func genFn0(p0 uint) string {
    p0 = (^(p0))
    v4 := []int{}
    switch int(p0) {
    case 618:
        var v5 uint64 = 18446744073709551615
        _ = v5
    case 32768:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -985:
        var v6 uint = p0
        _ = v6
    }
    p0 -= (p0 - 4294967295)
    _ = p0
    _ = v4
    return "a"
}

func main() {
    var v1 int8 = 42
    v2 := []int{32767}
    var v3 int32 = -1
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    v3 = (v3 - (v3 - 42))
    fmt.Println(v1)
    var v9 int32 = v3
    v10 := GenS1{F0: "hello"}
    v1 = (v1 | v1)
    _ = len(v2)
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int64 = 0
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: ""}
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v15 := GenS2{F0: 1}
    v9 -= 2147483647
    v16 := GenS3{F0: 0, F1: v1, F2: 4294967295}
    if (false || false) {
        fmt.Println(v9)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v1 -= v1
    switch "foo bar" {
    case "a":
        v17 := GenS3{F0: -615, F1: (v1 ^ (v1 + 1)), F2: 0}
        _ = v17
    case "foo bar":
        switch "foo bar" {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "foo bar":
            v18 := GenS4{F0: 1000, F1: 550, F2: -827}
            _ = v18
        case "":
            v9 -= (v3 << ((733) & 63))
        }
    default:
        for fi19 := 0; fi19 < 16; fi19++ {
            v1 += v1
            switch (fi19 + 288) {
            case 127:
                v1 = v1
            case 191:
                var v20 int16 = 0
                _ = v20
            case -278:
                switch 127 {
                case -262:
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if (2 > 1) {
                            if true {
                                if (1 == v1) {
                                    var v22 int8 = (v1 << ((936) & 63))
                                    var v23 string = ("a" + ("a" + ""))
                                    _ = v22
                                    _ = v23
                                }
                                v9 += -2147483648
                            }
                        } else {
                            if (!true) {
                                v24 := GenS5{F0: "", F1: 2147483647}
                                _ = v24
                            }
                        }
                        _ = fi21
                    }
                }
            }
            v25 := GenS2{F0: 42}
            _ = fi19
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
