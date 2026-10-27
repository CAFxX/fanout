package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 float32 = (1e10 * float32((3.14159 / 1e10)))
    var v2 float32 = v1
    var v3 uint = 2
    var v4 uint64 = 1
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := v5["a"]
        if true {
            var v8 int8 = -128
            switch fi6 {
            case 127:
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: 0, F1: 42}
                    v11 := GenS1{F0: 1, F1: int64(v4)}
                    v12 := GenS1{F0: 9223372036854775807, F1: -578}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case -575:
                var v13 uint32 = 650
                _ = v13
            }
            _ = v8
        } else {
            v5[v7] = v7
            v1 = float32(0.1)
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int = (fi6 << (((^(v3))) & 63))
            v16 := GenS2{F0: 4294967295, F1: 1}
            if (68 <= 4294967295) {
                switch (fi6 | -2147483648) {
                case -2147483648:
                    var v17 float32 = 0.1
                    _ = v17
                case 416:
                    fmt.Println(v3)
                case 167:
                    var v18 bool = (false && (false || (908 == 2)))
                    _ = v18
                default:
                    if true {
                        v7 = "foo bar"
                        if (32767 != 65535) {
                            var v19 uint64 = 964
                            _ = v19
                        }
                    } else {
                        v3 += v3
                        fmt.Println(v2, fi6, v3)
                        var v20 int16 = -1
                        _ = v20
                    }
                }
                var v21 bool = false
                if true {
                    v7 = v7
                }
                _ = v21
            } else {
                v22 := GenS1{F0: -9223372036854775808, F1: 288}
                v3 = 4294967295
                fmt.Println(v2, v2, v3)
                _ = v22
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi6
        _ = v7
    }
    v3 -= 255
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            delete(v5, "hello")
            fmt.Println(fi24, v4, fi24)
            v25 := GenS2{F0: 541, F1: -811}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v26 := GenS3{F0: 65535}
    if (639 >= -2147483648) {
        switch 22 {
        case 401:
            var v27 int64 = 9223372036854775807
            _ = v27
        default:
            switch ("\n\t" + (("" + "a") + "\n\t")) {
            case "":
                var v28 uint8 = 42
                _ = v28
            }
        }
        var v29 uint16 = 695
        v30 := GenS1{F0: int64(v4), F1: 1}
        _ = v29
        _ = v30
    } else {
        for fi31 := 0; fi31 < 15; fi31++ {
            v4 += 407
            v4 = 902
            _ = fi31
        }
        var v32 int = 1
        switch (v32 | int(v3)) {
        case 1:
            v2 += v1
        default:
            var v33 uint = v3
            _ = v33
        }
        _ = v32
    }
    v4 += (^(v4))
    v4 = v4
    v34 := GenS2{F0: 42, F1: -2147483648}
    var v35 int64 = 891
    v1 += v2
    v36 := []int{}
    var v37 float64 = 0.1
    v5["hello"] = ("hello" + "hello")
    var v38 int64 = v35
    var v39 int64 = (v35 / ((v38) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
