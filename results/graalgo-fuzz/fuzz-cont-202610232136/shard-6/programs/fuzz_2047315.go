package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint }
func genFn0() int64 {
    v4 := GenS1{F0: 154, F1: 42}
    if (177 == 1) {
        var v5 uint64 = 433
        _ = v5
    } else {
        var v6 uint32 = 42
        _ = v6
    }
    if (1 > -2147483648) {
        v7 := GenS1{F0: 1, F1: 1}
        _ = v7
    }
    _ = v4
    return 42
}

func main() {
    v1 := []uint32{156, 1, 2147483647, 1}
    v2 := make(map[int]bool)
    var v3 uint = 42
    v3 += 293
    var v8 int = 4294967295
    var v9 int8 = 1
    for fi10 := 0; fi10 < 12; fi10++ {
        v3 -= 65535
        _ = fi10
    }
    v3 += (^(v3))
    switch "" {
    case "a":
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int8 = (v9 + (v9 ^ 0))
            v3 = 0
            _ = fi11
            _ = v12
        }
    case "":
        var v13 float32 = -2.25
        _ = v13
    }
    if true {
        v9 += v9
        v14 := GenS1{F0: 2, F1: -585}
        _ = v14
    } else {
        if false {
            switch (v8 / (((v8 % ((v8) | 1))) | 1)) {
            case 9223372036854775807:
                var v15 int32 = -2147483648
                _ = v15
            case 4294967295:
                var v16 int = 128
                _ = v16
            }
            if len(v1) > 0 {
                v1[0] = 821
            }
        } else {
            var v17 int8 = v9
            _ = v17
        }
        v18 := v2[v8]
        _ = v18
    }
    var v19 string = ("hello" + "")
    if ((!true) && false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (float64((-2.25 / 0.1)) != 1.5) {
        var v20 uint16 = 255
        switch ((v19 + v19) + (("foo bar" + v19) + v19)) {
        case "foo bar":
            switch 2147483647 {
            case 4294967295:
                switch (v8 | (v8 & v8)) {
                case 804:
                    var v21 uint16 = v20
                    _ = v21
                case -1:
                    v20 += (v20 >> (((v3 & 0)) & 63))
                }
            default:
                for fi22 := 0; fi22 < 13; fi22++ {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        delete(v2, (fi22 - (fi23 * -1)))
                        _ = fi23
                    }
                    v20 = 2
                    v9 += -89
                    _ = fi22
                }
            }
        case "hello":
            v2[(v8 % ((v8) | 1))] = false
        case "":
            fmt.Println(v19)
        }
        switch v8 {
        case 32768:
            for fi24 := 0; fi24 < 5; fi24++ {
                for fi25 := 0; fi25 < 4; fi25++ {
                    fmt.Println(fi25, v8, v20)
                    if (v19 == "a") {
                        delete(v2, (fi24 + fi24))
                        v2[(-(v8))] = true
                        v26 := GenS1{F0: 2147483647, F1: genFn0()}
                        _ = v26
                    }
                    delete(v2, 567)
                    v9 = (v9 - -108)
                    _ = fi25
                }
                _ = fi24
            }
        case 4294967295:
            var v27 string = v19
            _ = v27
        case 127:
            var v28 int16 = 671
            _ = v28
        default:
            v3 -= (^(v3))
        }
        _ = v20
    }
    v19 = v19
    v2[32767] = false
    v19 = v19
    for fi29 := 0; fi29 < 20; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            v3 = 42
            v31 := GenS2{F0: (^(v3))}
            _ = fi30
            _ = v31
        }
        v3 += (v3 & v3)
        v19 = v19
        _ = fi29
    }
    var v32 string = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v8
    _ = v9
}
