package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int32, p1 uint32, p2 uint16) uint {
    for fi4 := 0; fi4 < 18; fi4++ {
        p2 = (p2 << ((797) & 63))
        _ = fi4
    }
    switch 1 {
    case 32768:
        if (0.1 != 1.5) {
            fmt.Println(p1, p2, p2)
        } else {
            for fi5 := 0; fi5 < 9; fi5++ {
                if true {
                    if (-172 >= -32768) {
                        fmt.Println(p1, p2, p1)
                    } else {
                        p1 = p1
                        p2 = (p2 - p2)
                        p1 = p1
                    }
                }
                if true {
                    p0 = (p0 << ((2) & 63))
                    fmt.Println(p0)
                } else {
                    for fi6 := 0; fi6 < 14; fi6++ {
                        if (127 != 0) {
                            switch "foo bar" {
                            case "a":
                                p2 -= (p2 + (p2 / ((p2) | 1)))
                            default:
                                v7 := []int{625}
                                _ = v7
                            }
                            fmt.Println(p0, p0)
                            var v8 int = fi6
                            fmt.Println(v8, v8, v8)
                            var v9 uint = 0
                            _ = v8
                            _ = v9
                        }
                        p0 += 1
                        _ = fi6
                    }
                    v10 := GenS1{F0: uint64(p1), F1: 255, F2: 32767}
                    _ = v10
                }
                switch (fi5 - fi5) {
                case 128:
                    if true {
                        p1 = (p1 & p1)
                        p1 = 0
                        v11 := GenS1{F0: 0, F1: 255, F2: -32768}
                        _ = v11
                    } else {
                        v12 := []int{fi5}
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = v12
                    }
                }
                _ = fi5
            }
        }
    default:
        if (!(!false)) {
            var v13 uint8 = 2
            _ = v13
        } else {
            p0 = p0
        }
    }
    v14 := GenS2{F0: p2, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 4294967295
}

func main() {
    var v1 uint8 = 1
    var v2 uint16 = 0
    v3 := make(map[string]string)
    v15 := []int{32768, 4294967295, -1}
    v1 += (v1 + (v1 - v1))
    var v16 int32 = 0
    for fi17 := 0; fi17 < 4; fi17++ {
        if (("" < "\n\t") && (true && false)) {
            if false {
                for fi18 := 0; fi18 < 19; fi18++ {
                    delete(v3, "\n\t")
                    _ = fi18
                }
            }
            var v19 float32 = (0.0 - (float32(1e10) + float32(1.5)))
            fmt.Println(v19, v19, fi17)
            _ = v19
        }
        v20 := GenS1{F0: 18446744073709551615, F1: 0, F2: 42}
        v21 := GenS3{F0: -9223372036854775808}
        _ = fi17
        _ = v20
        _ = v21
    }
    v16 = (v16 | int32(v2))
    var v22 float32 = float32((-2.25 - 1.5))
    switch ("\n\t" + ("hello" + "hello")) {
    case "a":
        v23 := GenS4{F0: "\n\t", F1: v2}
        _ = v23
    case "hello":
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(fi24, fi24, fi24)
            _ = fi24
        }
    case "foo bar":
        var v25 int = 9223372036854775807
        _ = v25
    }
    v26 := GenS4{F0: "hello", F1: (v2 / (((v2 % ((v2) | 1))) | 1))}
    v27 := GenS1{F0: 1, F1: v1, F2: 1}
    var v28 int = 32768
    var v29 uint = uint(v2)
    var v30 int8 = -33
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
