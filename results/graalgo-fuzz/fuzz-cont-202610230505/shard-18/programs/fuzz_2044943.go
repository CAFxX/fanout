package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 int32, p1 uint16) uint {
    if (-656 < -686) {
        v6 := GenS1{F0: 42}
        _ = v6
    } else {
        var v7 int16 = 332
        var v8 int = 0
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 545}
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{240}
    var v2 uint32 = 0
    v3 := make(map[string]int)
    var v4 int64 = -860
    var v5 int16 = -1
    for fi10 := 0; fi10 < 12; fi10++ {
        v4 -= v4
        var v11 uint8 = 1
        switch (^(fi10)) {
        case -9223372036854775808:
            v5 += (v5 * v5)
        case 128:
            v12 := GenS1{F0: v2}
            _ = v12
        case 32767:
            v13 := v3["\n\t"]
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v4)
    if (!false) {
        if true {
            v5 = int16(v2)
            v3[("hello" + "foo bar")] = 9223372036854775807
        }
        var v14 uint8 = 127
        _ = v14
    } else {
        for fi15 := 0; fi15 < 4; fi15++ {
            v5 = 42
            v3[(((("a" + (((("hello" + ("" + "")) + "") + "") + "foo bar")) + "a") + ("\n\t" + ("hello" + (("a" + "") + "\n\t")))) + ((("a" + "hello") + "hello") + (("a" + "hello") + "foo bar")))] = fi15
            _ = fi15
        }
        fmt.Println(v4)
        for fi16 := 0; fi16 < 16; fi16++ {
            switch 9223372036854775807 {
            case 157:
                if (0 < -1) {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v18 := GenS1{F0: 0}
                        _ = fi17
                        _ = v18
                    }
                    var v19 float32 = ((1e10 + float32(((0.0 * 3.14159) / ((3.14159 - -2.25) - 3.14159)))) + 0.1)
                    v3[("\n\t" + "hello")] = (fi16 ^ 1)
                    _ = v19
                }
            case 32768:
                if len(v1) > 0 {
                    v1[0] = (^(v2))
                }
            default:
                v2 = v2
            }
            v20 := GenS1{F0: (v2 >> ((0) & 63))}
            v21 := GenS1{F0: (v2 >> ((65535) & 63))}
            _ = fi16
            _ = v20
            _ = v21
        }
    }
    v22 := GenS1{F0: v2}
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS1{F0: v2}
            if false {
                var v26 uint16 = 80
                var v27 int16 = (v5 % ((1) | 1))
                _ = v26
                _ = v27
            }
            var v28 uint = 360
            _ = fi24
            _ = v25
            _ = v28
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS1{F0: (v2 * (v2 / ((1) | 1)))}
            v4 = 0
            v31 := GenS2{F0: 35}
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 bool = true
        _ = fi23
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 int = (fi33 & -2147483648)
        var v35 int64 = (v4 / ((-526) | 1))
        if false {
            switch fi33 {
            case 42:
                var v36 uint64 = 1
                _ = v36
            default:
                v37 := GenS2{F0: 128}
                _ = v37
            }
            if (0 == 42) {
                v34 = (fi33 % ((fi33) | 1))
                if true {
                    var v38 float32 = -2.25
                    v39 := GenS1{F0: (v2 * (v2 << ((1) & 63)))}
                    _ = v38
                    _ = v39
                }
            }
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    v4 = (v4 - v4)
    v40 := GenS3{F0: 32768, F1: 212, F2: -1}
    v41 := GenS3{F0: -604, F1: 255, F2: int32(v2)}
    var v42 int = -2147483648
    var v43 int32 = -722
    var v44 uint32 = 4294967295
    v45 := GenS4{F0: 42, F1: (42 == 100)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
