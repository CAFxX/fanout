package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 int, p1 uint8, p2 int64) uint32 {
    fmt.Println(p1)
    if false {
        if false {
            var v5 int32 = 255
            var v6 string = ""
            p0 = (-(p0))
            _ = v5
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{-2147483648, (fi7 & 9223372036854775807)}
            switch -268 {
            case 32768:
                for fi9 := 0; fi9 < 9; fi9++ {
                    p1 = (p1 % ((p1) | 1))
                    _ = fi9
                }
            default:
                p1 = p1
            }
            _ = fi7
            _ = v8
        }
        var v10 uint64 = 42
        var v11 bool = true
        _ = v10
        _ = v11
    }
    fmt.Println(p1)
    if (0 > 0) {
        p0 = p0
        if (65535 > 927) {
            p0 -= int(p2)
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            switch (fi12 | (fi12 - 42)) {
            case -394:
                for fi13 := 0; fi13 < 14; fi13++ {
                    fmt.Println(p0, fi12, p2)
                    v14 := GenS1{F0: p1, F1: (p0 * -359), F2: (p1 >> ((42) & 63))}
                    _ = fi13
                    _ = v14
                }
            case 32767:
                v15 := []int{}
                _ = v15
            default:
                if false {
                    var v16 int8 = 1
                    _ = v16
                }
            }
            _ = fi12
        }
    }
    switch p0 {
    case 127:
        var v17 uint = 42
        _ = v17
    case -952:
        for fi18 := 0; fi18 < 10; fi18++ {
            p2 += (p2 / (((p2 & p2)) | 1))
            if ((p0 & (fi18 | p0)) >= (-(fi18))) {
                var v19 int32 = -341
                fmt.Println(p0, p2, p1)
                _ = v19
            } else {
                v20 := GenS1{F0: 66, F1: p0, F2: p1}
                _ = v20
            }
            var v21 int = (p0 + 127)
            _ = fi18
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 293
}

func main() {
    var v1 int64 = 42
    var v2 int = 0
    v3 := []uint32{788, 813, 65535, 2147483648}
    var v4 uint64 = 2
    v4 -= v4
    var v22 int64 = v1
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS1{F0: 184, F1: fi23, F2: 128}
        v22 -= (v22 >> ((65535) & 63))
        _ = fi23
        _ = v24
    }
    _ = len(v3)
    if false {
        for fi25 := 0; fi25 < 1; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                if (v1 < 42) {
                    if true {
                        v27 := GenS2{F0: 42}
                        var v28 uint32 = genFn0(42, 193, v22)
                        _ = len(v3)
                        _ = v27
                        _ = v28
                    } else {
                        v1 -= 1
                        var v29 int = fi25
                        _ = v29
                    }
                    v4 = (v4 * v4)
                }
                _ = fi26
            }
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS2{F0: 1000}
                _ = fi30
                _ = v31
            }
            _ = fi25
        }
        if true {
            v1 = (v22 * -466)
        } else {
            v32 := GenS3{F0: 255, F1: "a"}
            v1 -= (v1 * v22)
            _ = v32
        }
    }
    var v33 uint8 = 42
    if (genFn0(769, (v33 + 127), v1) >= 65535) {
        v1 += v1
        if ("" == (("\n\t" + ("foo bar" + (("" + "hello") + ""))) + "hello")) {
            switch v2 {
            case -266:
                if (true || true) {
                    fmt.Println(v33, v4)
                    var v34 string = ("a" + "foo bar")
                    _ = v34
                }
            }
        }
        v35 := GenS3{F0: 32767, F1: "a"}
        _ = v35
    }
    v22 = v22
    _ = len(v3)
    var v36 int8 = 127
    v33 = v33
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v36
    _ = v4
}
