package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
func genFn0(p0 uint) uint16 {
    var v5 float32 = -2.25
    p0 += p0
    if false {
        v6 := GenS1{F0: 2}
        _ = v6
    } else {
        v7 := []int{-346, 42, 127}
        var v8 float64 = (float64(v5) / 1e10)
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return 65535
}
func genFn1(p0 uint32) string {
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 bool = (42 != 42)
        switch ((("" + "foo bar") + "") + ((("" + "a") + "foo bar") + ("" + "a"))) {
        case "":
            v10 = v10
        default:
            p0 = p0
        }
        var v11 int8 = 37
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    _ = p0
    return ""
}

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 24
    var v3 uint64 = 1
    v4 := make(map[string]uint32)
    v3 -= (v3 - (^(v3)))
    var v12 float64 = 0.1
    v2 += 461
    var v13 int64 = -996
    v4[("foo bar" + "foo bar")] = v2
    var v14 int64 = -1
    v2 += (v2 / (((v2 + (^(v2)))) | 1))
    v15 := GenS2{F0: v13}
    v16 := []int{-949, 4294967295}
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS2{F0: int64(v2)}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        switch "" {
        case "hello":
            v14 = v14
        case "\n\t":
            fmt.Println(fi19, v3, v1)
        case "foo bar":
            for fi20 := 0; fi20 < 11; fi20++ {
                if (32767 <= 65535) {
                    fmt.Println(fi19, fi19, fi19)
                }
                var v21 float32 = 1e10
                _ = fi20
                _ = v21
            }
        default:
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        _ = fi19
    }
    v22 := GenS3{F0: v3, F1: false, F2: 32767}
    if true {
        for fi23 := 0; fi23 < 20; fi23++ {
            v24 := GenS2{F0: int64(v3)}
            if (609 < 2147483647) {
                fmt.Println(v13, fi23)
                if (true || true) {
                    var v25 int16 = 0
                    switch (fi23 >> ((42) & 63)) {
                    case 32768:
                        var v26 string = "foo bar"
                        _ = v26
                    case 0:
                        if true {
                            for fi27 := 0; fi27 < 14; fi27++ {
                                v3 = v3
                                _ = fi27
                            }
                            if false {
                                for fi28 := 0; fi28 < 15; fi28++ {
                                    var v29 uint = 1
                                    if true {
                                        v2 += (v2 >> (((v29 / ((379) | 1))) & 63))
                                        v2 -= 2147483647
                                        fmt.Println(fi28, v29, v14)
                                    } else {
                                        if false {
                                            fmt.Println(v1, v29, v25)
                                            v14 = (v13 & int64(v2))
                                            var v30 bool = true
                                            _ = v30
                                        }
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                fmt.Println(v2, fi23)
                            } else {
                                var v31 int32 = 0
                                _ = v31
                            }
                        }
                    }
                    for fi32 := 0; fi32 < 3; fi32++ {
                        var v33 uint = 18446744073709551615
                        v34 := GenS1{F0: v3}
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    _ = v25
                }
            } else {
                fmt.Println(v3)
            }
            v35 := v4["\n\t"]
            _ = fi23
            _ = v24
            _ = v35
        }
        v14 = v14
        for fi36 := 0; fi36 < 3; fi36++ {
            switch fi36 {
            case 9223372036854775807:
                var v37 uint8 = 128
                _ = v37
            }
            for fi38 := 0; fi38 < 17; fi38++ {
                var v39 int32 = int32(v13)
                _ = fi38
                _ = v39
            }
            v1 = v1
            _ = fi36
        }
    } else {
        var v40 string = "foo bar"
        _ = v40
    }
    var v41 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v41
}
