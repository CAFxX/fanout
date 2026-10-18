package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool; F1 int16; F2 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 string }
type GenS4 struct { F0 int8 }

func main() {
    v1 := []string{(("\n\t" + "hello") + "\n\t"), "a"}
    var v2 int8 = -88
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    v5 := []uint32{42}
    v6 := v4[2147483647]
    var v7 int16 = 42
    var v8 int32 = 0
    if false {
        _ = len(v3)
    } else {
        if false {
            v8 = (^(v8))
        }
        switch (v6 + v6) {
        case "foo bar":
            v9 := GenS1{F0: 480, F1: 128}
            _ = v9
        case "hello":
            v10 := GenS2{F0: ((-1 > 0) && false), F1: -1, F2: v7}
            _ = v10
        default:
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int8 = (v2 - -43)
                if (0 == 17) {
                    var v13 uint32 = 4294967295
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
        }
    }
    switch "" {
    case "a":
        if (42 >= 966) {
            for fi14 := 0; fi14 < 8; fi14++ {
                v6 = "a"
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi14
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                _ = len(v1)
                _ = fi15
            }
            v16 := GenS3{F0: 42, F1: v2, F2: v6}
            _ = v16
        }
    case "":
        for fi17 := 0; fi17 < 10; fi17++ {
            v2 = 111
            _ = fi17
        }
    default:
        switch 128 {
        case -1:
            v18 := v4[0]
            _ = v18
        default:
            var v19 uint8 = 128
            _ = v19
        }
    }
    switch 9223372036854775807 {
    case 127:
        v20 := GenS2{F0: ((!(!(v6 <= v6))) || false), F1: 32767, F2: v7}
        _ = v20
    case 0:
        for fi21 := 0; fi21 < 14; fi21++ {
            if len(v5) > 0 {
                v5[0] = 842
            }
            v2 = v2
            _ = fi21
        }
    }
    v7 -= (v7 / ((v7) | 1))
    v7 -= 1000
    if (!(!(int(v7) != 42))) {
        v7 = v7
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(v8, v8, v7)
            _ = fi22
        }
    } else {
        for fi23 := 0; fi23 < 13; fi23++ {
            v8 = v8
            _ = fi23
        }
    }
    var v24 int8 = v2
    if (127 >= 941) {
        delete(v3, 0)
    }
    v6 = v6
    if (true && false) {
        v25 := GenS4{F0: (v24 % (((v24 % ((100) | 1))) | 1))}
        if (!(!(433 != (v7 >> ((0) & 63))))) {
            switch v6 {
            case "hello":
                var v26 int8 = 42
                _ = v26
            default:
                v24 = v24
            }
            v8 += -767
        }
        if (v6 >= v6) {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 bool = false
                _ = fi27
                _ = v28
            }
        }
        _ = v25
    } else {
        for fi29 := 0; fi29 < 19; fi29++ {
            for fi30 := 0; fi30 < 12; fi30++ {
                v7 = (v7 >> ((1) & 63))
                v31 := GenS2{F0: (false || true), F1: v7, F2: (v7 % ((v7) | 1))}
                _ = fi30
                _ = v31
            }
            v2 -= (v2 << ((457) & 63))
            _ = fi29
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 float64 = float64(0.0)
            v8 += 0
            var v34 uint8 = 255
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 bool = (false || true)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v36 uint32 = 4294967295
    var v37 uint64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
