package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 bool }
type GenS3 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 bool = false
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v5 int64 = -197
    v5 += v5
    v6 := GenS1{F0: "", F1: v5, F2: 42}
    v7 := GenS1{F0: ("\n\t" + ("\n\t" + "hello")), F1: v5, F2: 255}
    if ((!v2) && (!(v5 == 441))) {
        var v8 string = "hello"
        _ = v8
    } else {
        v5 -= v5
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := v3[1]
                v12 := GenS1{F0: "\n\t", F1: v5, F2: 80}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    if v2 {
        if (((float32(0.0) * 3.14159) / 1.0) == 0.1) {
            if ("\n\t" > "") {
                _ = len(v4)
                v13 := GenS2{F0: -64, F1: 13, F2: v2}
                for fi14 := 0; fi14 < 10; fi14++ {
                    if (0.1 > 0.0) {
                        if false {
                            fmt.Println(v5, fi14)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v15 := GenS1{F0: "\n\t", F1: (-(v5)), F2: 201}
                            _ = v15
                        } else {
                            v16 := GenS2{F0: 0, F1: -1, F2: (42 <= fi14)}
                            _ = v16
                        }
                    } else {
                        switch (fi14 / ((32767) | 1)) {
                        case -738:
                            var v17 int = 404
                            _ = v17
                        case 1:
                            v4[966] = 1
                        case 42:
                            var v18 int16 = -748
                            _ = v18
                        }
                    }
                    v4[4294967295] = 1
                    v5 = v5
                    _ = fi14
                }
                _ = v13
            }
            v19 := GenS3{F0: 2, F1: 883}
            _ = v19
        }
        switch (("a" + ("\n\t" + ("" + "foo bar"))) + "") {
        case "\n\t":
            var v20 uint = 42
            _ = v20
        case "":
            v2 = (!("" >= "hello"))
        }
        var v21 uint16 = 0
        _ = v21
    } else {
        var v22 uint8 = 6
        _ = v22
    }
    v5 = v5
    v4[627] = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
