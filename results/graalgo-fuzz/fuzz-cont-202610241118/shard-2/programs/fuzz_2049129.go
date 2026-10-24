package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 int16; F2 int64 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint8; F2 string }
func genFn0(p0 int32, p1 int16) int32 {
    fmt.Println(p0, p0)
    v7 := GenS1{F0: true}
    if (!true) {
        var v8 string = (("hello" + "hello") + "")
        if (!true) {
            p1 += 32767
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 float32 = 0.0
                _ = fi9
                _ = v10
            }
        }
        _ = v8
    } else {
        var v11 int = -186
        v12 := []int{(v11 >> ((1) & 63)), (v11 + 553), v11, 42}
        switch "foo bar" {
        case "\n\t":
            if false {
                var v13 bool = false
                _ = len(v12)
                p0 = (^(p0))
                _ = v13
            } else {
                var v14 uint64 = 1
                switch 782 {
                case 1:
                    fmt.Println(v14, p1, p1)
                case 0:
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v14 -= 4294967295
                        _ = fi15
                    }
                case 32768:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
                p1 += -32768
                _ = v14
            }
        case "hello":
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        case "a":
            v16 := GenS2{F0: "\n\t", F1: -128}
            _ = v16
        default:
            switch v11 {
            case 2147483647:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            case 1:
                for fi17 := 0; fi17 < 16; fi17++ {
                    v12 = append(v12, v11)
                    _ = fi17
                }
            case 935:
                for fi18 := 0; fi18 < 19; fi18++ {
                    switch (-(fi18)) {
                    case -2147483648:
                        for fi19 := 0; fi19 < 3; fi19++ {
                            p0 = (p0 >> ((747) & 63))
                            _ = fi19
                        }
                    }
                    _ = fi18
                }
            default:
                fmt.Println(v11)
            }
        }
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    v1 := []int{700, 1, 760}
    var v2 uint32 = 2147483647
    v3 := make(map[string]int)
    var v4 string = ("foo bar" + ("a" + (("hello" + ("foo bar" + "a")) + "foo bar")))
    v5 := make(map[string]string)
    v6 := make(map[string]int)
    var v20 string = v4
    v2 = 0
    v4 = v20
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint = 18446744073709551615
        for fi23 := 0; fi23 < 12; fi23++ {
            fmt.Println(v2, v20, fi23)
            v24 := GenS3{F0: -32768}
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v25 := GenS4{F0: 955}
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        v26 := v3["\n\t"]
        _ = fi21
        _ = v22
        _ = v26
    }
    v27 := GenS5{F0: -9223372036854775808, F1: 42, F2: -475}
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := GenS6{F0: 569, F1: v2}
        _ = fi28
        _ = v29
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v30 uint64 = uint64(v2)
        _ = v30
    }
    v31 := GenS7{F0: -2147483648, F1: 255, F2: v4}
    v4 = "a"
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := GenS5{F0: 0, F1: 32767, F2: -884}
        _ = fi32
        _ = v33
    }
    v2 = (v2 << ((1) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
