package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int64 }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{42}
    v3 := make(map[int]uint32)
    var v4 int32 = 42
    v5 := make(map[uint32]bool)
    if (42 == -72) {
        v4 = (v4 | (^(v4)))
        if (42 < -99) {
            v6 := GenS1{F0: -514, F1: ("a" + "hello"), F2: 1}
            for fi7 := 0; fi7 < 7; fi7++ {
                v4 += (v4 * (v4 >> ((795) & 63)))
                _ = fi7
            }
            _ = v6
        }
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            if true {
                var v9 int = fi8
                v10 := GenS1{F0: 1, F1: "foo bar", F2: 1}
                fmt.Println(fi8, fi8, v9)
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 uint = 18446744073709551615
                _ = fi11
                _ = v12
            }
        }
        var v13 uint = 2
        _ = fi8
        _ = v13
    }
    if (216 <= 128) {
        var v14 int = 0
        _ = v14
    }
    fmt.Println(v4)
    var v15 float32 = 1.5
    var v16 int16 = -32768
    v4 += (v4 % (((v4 + -107)) | 1))
    v4 = (v4 & v4)
    var v17 uint8 = 30
    v18 := GenS1{F0: -9223372036854775808, F1: "a", F2: 0}
    var v19 int = -9223372036854775808
    delete(v3, v19)
    v20 := v5[2]
    var v21 float64 = 0.1
    if ((v20 || (0 >= 2)) || (!v20)) {
        if (0 < (v17 * v17)) {
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint16 = 0
                _ = fi22
                _ = v23
            }
            var v24 int = (v19 << ((2) & 63))
            _ = v24
        } else {
            fmt.Println(v20)
            v2 = append(v2, 2147483648)
        }
        switch (^(v19)) {
        case 0:
            delete(v5, 591)
        case 71:
            if true {
                var v25 int16 = v16
                var v26 int = -1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v25
                _ = v26
            }
        case 128:
            v27 := GenS2{F0: v17, F1: "foo bar"}
            _ = v27
        }
        var v28 int8 = 8
        _ = v28
    }
    if v20 {
        v29 := GenS1{F0: -1, F1: ("hello" + ((((("a" + ("\n\t" + ("foo bar" + "hello"))) + "hello") + "") + ("\n\t" + (((("hello" + ((("a" + "") + "\n\t") + ("\n\t" + "\n\t"))) + "") + "\n\t") + "\n\t"))) + "")), F2: -9223372036854775808}
        for fi30 := 0; fi30 < 14; fi30++ {
            if false {
                if (214 != v4) {
                    var v31 uint64 = 0
                    _ = v31
                } else {
                    switch (v19 & 1) {
                    case 746:
                        v32 := GenS1{F0: 9223372036854775807, F1: "", F2: -506}
                        _ = v32
                    case 32768:
                        v15 += v15
                    case 430:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        var v33 string = "a"
                        _ = v33
                    }
                    fmt.Println(v17, v16, v15)
                }
                for fi34 := 0; fi34 < 4; fi34++ {
                    switch "\n\t" {
                    case "hello":
                        switch ("\n\t" + ("hello" + ("foo bar" + "\n\t"))) {
                        case "foo bar":
                            v16 += v16
                        case "a":
                            var v35 int64 = -782
                            _ = v35
                        }
                    }
                    _ = fi34
                }
            } else {
                var v36 bool = (!v20)
                _ = v36
            }
            _ = fi30
        }
        _ = v29
    } else {
        var v37 uint32 = 779
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
