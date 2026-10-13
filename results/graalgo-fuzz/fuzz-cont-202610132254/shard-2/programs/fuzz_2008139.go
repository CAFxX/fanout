package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []int{9223372036854775807, -9223372036854775808, 370}
    v2 := make(map[string]bool)
    var v3 int8 = 100
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 bool = ((false || false) && ((!true) || false))
            _ = fi6
            _ = v7
        }
        fmt.Println(fi5, v3, fi5)
        if true {
            delete(v2, "hello")
            var v8 string = ("" + "hello")
            fmt.Println(v4)
            _ = v8
        } else {
            v9 := GenS1{F0: "hello", F1: 42, F2: 201}
            v3 = -128
            _ = v9
        }
        _ = fi5
    }
    v3 = v3
    fmt.Println(v4, v3)
    v3 += (v3 * v3)
    v3 = (v3 * (v3 / ((v3) | 1)))
    v3 -= v3
    if true {
        if true {
            v3 = v3
        } else {
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 17; fi11++ {
                    switch fi10 {
                    case 128:
                        var v12 bool = false
                        _ = v12
                    default:
                        v13 := GenS1{F0: "", F1: -210, F2: 2}
                        _ = v13
                    }
                    fmt.Println(v4)
                    _ = fi11
                }
                v3 = (v3 * (v3 / (((^(v3))) | 1)))
                v1 = append(v1, fi10)
                _ = fi10
            }
            var v14 int8 = (v3 << ((65535) & 63))
            if (!(true || true)) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 uint64 = 42
                    _ = fi15
                    _ = v16
                }
                var v17 uint8 = v4
                _ = v17
            } else {
                v14 = -7
                if (-811 <= 1) {
                    v14 = (v3 % ((v3) | 1))
                    fmt.Println(v4, v14, v14)
                }
            }
            _ = v14
        }
        v4 = 128
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    }
    v18 := GenS1{F0: "a", F1: 0, F2: 0}
    var v19 int = 32767
    var v20 string = (("hello" + "foo bar") + "")
    if (!false) {
        for fi21 := 0; fi21 < 6; fi21++ {
            if (9223372036854775807 > 624) {
                var v22 int8 = v3
                var v23 int8 = 127
                if (!((v20 + "\n\t") <= (v20 + v20))) {
                    v24 := GenS1{F0: (("a" + v20) + v20), F1: -746, F2: 354}
                    v20 = v20
                    fmt.Println(fi21)
                    _ = v24
                }
                _ = v22
                _ = v23
            } else {
                for fi25 := 0; fi25 < 9; fi25++ {
                    switch v20 {
                    case "foo bar":
                        var v26 int64 = 42
                        _ = v26
                    }
                    fmt.Println(v20, v4, fi21)
                    v20 = "foo bar"
                    _ = fi25
                }
                var v27 uint32 = 2147483647
                var v28 int16 = -1
                _ = v27
                _ = v28
            }
            v29 := GenS1{F0: v20, F1: 42, F2: 0}
            v3 = 6
            _ = fi21
            _ = v29
        }
    } else {
        var v30 int32 = -2147483648
        switch 32768 {
        case -769:
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 = 1
                switch (v20 + v20) {
                case "a":
                    var v31 float32 = -2.25
                    _ = v31
                case "hello":
                    var v32 int32 = v30
                    _ = v32
                case "":
                    var v33 bool = false
                    _ = v33
                default:
                    if ((v30 >= (v30 << ((535) & 63))) && ((0.1 - float64(-2.25)) == (0.1 * 3.14159))) {
                        v30 = int32(v4)
                        v34 := GenS2{F0: 2}
                        _ = v34
                    }
                }
            }
        }
        var v35 uint64 = 1
        _ = v30
        _ = v35
    }
    var v36 uint32 = 404
    v3 += (v3 % ((-1) | 1))
    v19 = v19
    if (!(42 < 921)) {
        if len(v1) > 0 {
            v1[0] = v19
        }
        v2[((v20 + v20) + v20)] = (!(v36 < (^(v36))))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v20, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
}
