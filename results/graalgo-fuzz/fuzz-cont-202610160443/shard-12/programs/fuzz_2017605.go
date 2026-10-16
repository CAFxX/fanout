package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 string; F2 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -128
    var v3 uint64 = 984
    var v4 uint16 = 572
    v3 += (v3 * 415)
    switch (("a" + "") + ("hello" + ("\n\t" + ""))) {
    case "\n\t":
        v3 = (v3 >> ((4294967295) & 63))
    }
    v3 = 716
    var v5 uint = 1
    v6 := []int{978, 1, 0, 709}
    fmt.Println(v4)
    var v7 int32 = int32(v5)
    fmt.Println(v2, v4, v3)
    var v8 int8 = (v2 + v2)
    if (v3 <= (v3 << (((v5 + v5)) & 63))) {
        var v9 uint = 486
        var v10 uint = (v9 >> (((^(v9))) & 63))
        v11 := v1[42]
        _ = v10
        _ = v11
        _ = v9
    }
    _ = len(v1)
    switch 1 {
    case 128:
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: 42}
            v3 = 42
            _ = fi12
            _ = v13
        }
    case 9223372036854775807:
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                if ((float32(-2.25) + (0.0 * (-2.25 + 1.5))) != (0.1 + 0.0)) {
                    if (-32768 <= 700) {
                        if true {
                            if len(v6) > 0 {
                                v6[0] = 0
                            }
                        } else {
                            var v16 float64 = (1.5 + 1e10)
                            var v17 int = (fi15 / ((42) | 1))
                            var v18 uint8 = 1
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        var v19 float64 = 1e10
                        _ = v19
                    }
                    v5 -= v5
                } else {
                    if false {
                        v7 -= (v7 - v7)
                        v4 -= v4
                    }
                    v20 := GenS1{F0: 119}
                    _ = v20
                }
                _ = fi15
            }
            _ = fi14
        }
    }
    if (true && ((true && (!false)) && (32767 > (v4 >> ((uint(v4)) & 63))))) {
        var v21 int16 = 0
        for fi22 := 0; fi22 < 8; fi22++ {
            fmt.Println(v4)
            v23 := GenS2{F0: fi22, F1: "", F2: "\n\t"}
            switch (fi22 % ((int(v5)) | 1)) {
            case -82:
                var v24 uint16 = v4
                _ = v24
            case -622:
                var v25 int64 = 38
                _ = v25
            case 32768:
                var v26 uint64 = v3
                _ = v26
            }
            _ = fi22
            _ = v23
        }
        var v27 bool = (3.14159 < 1e10)
        _ = v21
        _ = v27
    } else {
        if len(v6) > 0 {
            v6[0] = 127
        }
        v28 := GenS2{F0: 0, F1: "", F2: ("foo bar" + "\n\t")}
        fmt.Println(v2)
        _ = v28
    }
    var v29 int16 = 32767
    var v30 string = ("hello" + "foo bar")
    if (-2.25 >= 1.5) {
        var v31 bool = ((!true) || true)
        v32 := GenS1{F0: 0}
        _ = v31
        _ = v32
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v6 = append(v6, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
