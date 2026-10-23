package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint16 = 451
    var v2 uint16 = v1
    v3 := make(map[uint32]uint32)
    v4 := []string{"\n\t"}
    var v5 uint8 = 128
    v6 := []string{"hello", ("\n\t" + "\n\t")}
    v7 := GenS1{F0: "", F1: v2, F2: 491}
    if ((4294967295 == 0) && (!true)) {
        var v8 int32 = -2147483648
        _ = v8
    }
    if (!false) {
        for fi9 := 0; fi9 < 4; fi9++ {
            v5 -= uint8(v1)
            fmt.Println(fi9)
            _ = fi9
        }
        fmt.Println(v1)
    } else {
        v3[65535] = 2147483647
        if false {
            v10 := GenS1{F0: "hello", F1: (v1 / (((v2 * v2)) | 1)), F2: -9223372036854775808}
            _ = v10
        } else {
            v11 := GenS1{F0: "\n\t", F1: 20, F2: -1}
            for fi12 := 0; fi12 < 11; fi12++ {
                v3[2147483647] = 2147483648
                _ = fi12
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 uint = 255
                v4 = append(v4, "\n\t")
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 float64 = (1.5 / 1e10)
                    var v17 string = "hello"
                    if (fi15 > (fi15 / ((32768) | 1))) {
                        var v18 uint32 = 269
                        _ = v18
                    } else {
                        v2 = 950
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi13
                _ = v14
            }
            _ = v11
        }
    }
    var v19 int8 = -1
    v20 := GenS1{F0: "foo bar", F1: 1, F2: 414}
    for fi21 := 0; fi21 < 20; fi21++ {
        switch fi21 {
        case 32768:
            v19 = (v19 >> ((1) & 63))
        case 128:
            for fi22 := 0; fi22 < 6; fi22++ {
                v5 = 157
                if true {
                    if true {
                        switch ("a" + ("hello" + "\n\t")) {
                        case "foo bar":
                            var v23 int16 = 1
                            _ = v23
                        default:
                            if false {
                                v5 = v5
                                var v24 uint = 4294967295
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                _ = v24
                            }
                        }
                    } else {
                        var v25 uint64 = 4294967295
                        v5 = v5
                        _ = v25
                    }
                    v26 := GenS2{F0: 65535, F1: 0}
                    v2 = v2
                    _ = v26
                }
                var v27 int16 = 0
                _ = fi22
                _ = v27
            }
        }
        _ = fi21
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        if (!(true || true)) {
            v29 := GenS3{F0: 65535}
            _ = v29
        } else {
            v30 := GenS3{F0: 762}
            _ = v30
        }
        var v31 int64 = 0
        var v32 int16 = -5
        _ = fi28
        _ = v31
        _ = v32
    }
    if (-2147483648 >= 313) {
        v5 = 255
        var v33 int = -670
        _ = v33
    }
    v34 := GenS2{F0: 1, F1: v2}
    var v35 uint = 505
    var v36 uint = v35
    v1 -= uint16(v36)
    var v37 int64 = int64(v19)
    for fi38 := 0; fi38 < 7; fi38++ {
        _ = len(v3)
        _ = fi38
    }
    var v39 uint32 = 2147483647
    v40 := v3[v39]
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
