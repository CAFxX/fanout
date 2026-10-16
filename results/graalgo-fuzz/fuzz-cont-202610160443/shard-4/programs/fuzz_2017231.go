package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    var v1 int8 = 100
    v2 := []string{"hello", "\n\t"}
    var v3 int8 = (v1 / ((-48) | 1))
    var v4 float32 = 0.1
    var v5 uint = 402
    v6 := make(map[int]uint32)
    var v7 int = -2147483648
    fmt.Println(v5, v3)
    if true {
        v5 -= v5
        switch 1 {
        case -142:
            if (1000 != 32767) {
                var v8 uint8 = 194
                _ = v8
            }
        case 127:
            v9 := GenS1{F0: v5, F1: 255, F2: 17}
            _ = v9
        case 32768:
            for fi10 := 0; fi10 < 6; fi10++ {
                switch ((("foo bar" + "\n\t") + ("foo bar" + "")) + "foo bar") {
                case "":
                    v11 := GenS1{F0: v5, F1: 1, F2: 32767}
                    _ = v11
                case "hello":
                    var v12 uint64 = 4294967295
                    _ = v12
                default:
                    var v13 bool = false
                    _ = v13
                }
                v14 := v6[fi10]
                var v15 int = 994
                _ = fi10
                _ = v14
                _ = v15
            }
        default:
            var v16 bool = (42 > 360)
            _ = v16
        }
        v17 := GenS1{F0: 4294967295, F1: 32767, F2: 1}
        _ = v17
    }
    _ = len(v2)
    if true {
        v6[v7] = 65535
        fmt.Println(v4)
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            switch (v7 >> (((v5 + (v5 / ((289) | 1)))) & 63)) {
            case -206:
                v4 = (((v4 - v4) + (v4 / (0.1 + v4))) * v4)
            default:
                var v19 int16 = 32767
                _ = v19
            }
            v2 = append(v2, "a")
            _ = fi18
        }
    }
    v6[v7] = 2
    switch (v7 - 2147483647) {
    case 42:
        fmt.Println(v3, v7, v3)
    default:
        for fi20 := 0; fi20 < 14; fi20++ {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            switch (v7 | (fi20 * v7)) {
            case -11:
                v3 -= v1
            }
            _ = fi20
        }
    }
    if true {
        if (816 > 42) {
            for fi21 := 0; fi21 < 3; fi21++ {
                for fi22 := 0; fi22 < 8; fi22++ {
                    switch (fi22 + -9223372036854775808) {
                    case -2147483648:
                        var v23 uint32 = 2
                        _ = v23
                    case 4294967295:
                        switch (v7 & (-(fi22))) {
                        case 806:
                            var v24 uint64 = 42
                            _ = v24
                        case 4294967295:
                            var v25 uint64 = 1
                            _ = v25
                        default:
                            if (0 <= 873) {
                                fmt.Println(v3, v3)
                            }
                            v3 = v3
                            _ = len(v2)
                        }
                    default:
                        v26 := GenS1{F0: (^(v5)), F1: 32767, F2: -32768}
                        _ = v26
                    }
                    v27 := GenS1{F0: v5, F1: 0, F2: -594}
                    _ = fi22
                    _ = v27
                }
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 uint = v5
                    v30 := GenS2{F0: v3, F1: 4294967295}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v4 = v4
                _ = fi21
            }
        }
        var v31 int8 = v1
        _ = v31
    } else {
        switch 32767 {
        case 9223372036854775807:
            v32 := GenS2{F0: v1, F1: 42}
            _ = v32
        }
    }
    var v33 bool = (!(4294967295 >= 526))
    v2 = append(v2, (("" + "\n\t") + "hello"))
    v3 = (v3 + (v3 * v1))
    v34 := GenS3{F0: 142, F1: int64(v5)}
    var v35 string = "hello"
    v35 = v35
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
